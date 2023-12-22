package operatorpubkeys

import (
	"context"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

// OperatorPubkeysServiceInMemory is a stateful goroutine (see https://gobyexample.com/stateful-goroutines)
// implementation of PubkeyCompendiumService that listen for the NewPubkeyRegistration using a websocket connection
// to an eth client and stores the pubkeys in memory. Another possible implementation is using a mutex
// (https://gobyexample.com/mutexes) instead. We can switch to that if we ever find a good reason to.
//
// Warning: this service should probably not be used in production. Haven't done a thorough analysis of all the clients
// but there is still an open PR about an issue with ws subscription on geth: https://github.com/ethereum/go-ethereum/issues/23845
// Another reason to note for infra/devops engineer who would put this into production, is that this service crashes on
// websocket connection errors or when failing to query past events. The philosophy here is that hard crashing is
// better than silently failing, since it will be easier to debug. Naturally, this means that this aggregator using this service needs
// to be replicated and load-balanced, so that when it fails traffic can be switched to the other aggregator.
type OperatorPubkeysServiceInMemory struct {
	avsRegistrySubscriber avsregistry.AvsRegistrySubscriber
	avsRegistryReader     avsregistry.AvsRegistryReader
	logger                logging.Logger
	queryC                chan<- query
}
type query struct {
	operatorAddr common.Address
	// channel through which to receive the response (operator pubkeys)
	respC chan<- resp
}
type resp struct {
	operatorPubkeys types.OperatorPubkeys
	// false if operators were not present in the pubkey dict
	operatorExists bool
}

var _ OperatorPubkeysService = (*OperatorPubkeysServiceInMemory)(nil)

// NewPubkeyCompendiumInMemory constructs a PubkeyCompendiumServiceInMemory and starts it in a goroutine.
// It takes a context as argument because the "backfilling" of the database is done inside this constructor,
// so we wait for all past NewPubkeyRegistration events to be queried and the db to be filled before returning the service.
// The constructor is thus following a RAII-like pattern, of initializing the serving during construction.
// Using a separate initialize() function might lead to some users forgetting to call it and the service not behaving properly.
func NewPubkeyCompendiumInMemory(
	ctx context.Context,
	avsRegistrySubscriber avsregistry.AvsRegistrySubscriber,
	avsRegistryReader avsregistry.AvsRegistryReader,
	logger logging.Logger,
) *OperatorPubkeysServiceInMemory {
	queryC := make(chan query)
	pkcs := &OperatorPubkeysServiceInMemory{
		avsRegistrySubscriber: avsRegistrySubscriber,
		avsRegistryReader:     avsRegistryReader,
		logger:                logger,
		queryC:                queryC,
	}
	// We use this waitgroup to wait on the initialization of the inmemory pubkey dict,
	// which requires querying the past events of the pubkey registration contract
	wg := sync.WaitGroup{}
	wg.Add(1)
	pkcs.startServiceInGoroutine(ctx, queryC, &wg)
	wg.Wait()
	return pkcs
}

func (ops *OperatorPubkeysServiceInMemory) startServiceInGoroutine(ctx context.Context, queryC <-chan query, wg *sync.WaitGroup) {
	go func() {
		pubkeyDict := make(map[common.Address]types.OperatorPubkeys)

		// TODO(samlaf): we should probably save the service in the logger itself and add it automatically to all logs
		ops.logger.Debug("Subscribing to new pubkey registration event on pubkey compendium contract", "service", "PubkeyCompendiumServiceInMemory")
		newPubkeyRegistrationC, newPubkeyRegistrationSub, err := ops.avsRegistrySubscriber.SubscribeToNewPubkeyRegistrations()
		if err != nil {
			ops.logger.Error("Fatal error opening websocket subscription for new pubkey registrations", "err", err, "service", "PubkeyCompendiumServiceInMemory")
			// see the warning above the struct definition to understand why we panic here
			panic(err)
		}
		ops.queryPastRegisteredOperatorEventsAndFillDb(ctx, pubkeyDict)
		// The constructor can return after we have backfilled the db by querying the events of operators that have registered with the pubkey compendium
		// before the block at which we started the ws subscription above
		wg.Done()
		for {
			select {
			case <-ctx.Done():
				// TODO(samlaf): should we do anything here? Seems like this only happens when the aggregator is shutting down and we want graceful exit
				ops.logger.Infof("PubkeyCompendiumServiceInMemory: Context cancelled, exiting")
				return
			case err := <-newPubkeyRegistrationSub.Err():
				ops.logger.Error("Error in websocket subscription for new pubkey registration events. Attempting to reconnect...", "err", err, "service", "PubkeyCompendiumServiceInMemory")
				newPubkeyRegistrationSub.Unsubscribe()
				newPubkeyRegistrationC, newPubkeyRegistrationSub, err = ops.avsRegistrySubscriber.SubscribeToNewPubkeyRegistrations()
				if err != nil {
					ops.logger.Error("Error opening websocket subscription for new pubkey registrations", "err", err, "service", "PubkeyCompendiumServiceInMemory")
					// see the warning above the struct definition to understand why we panic here
					panic(err)
				}
			case newPubkeyRegistrationEvent := <-newPubkeyRegistrationC:
				operatorAddr := newPubkeyRegistrationEvent.Operator
				pubkeyDict[operatorAddr] = types.OperatorPubkeys{
					G1Pubkey: bls.NewG1Point(newPubkeyRegistrationEvent.PubkeyG1.X, newPubkeyRegistrationEvent.PubkeyG1.Y),
					G2Pubkey: bls.NewG2Point(newPubkeyRegistrationEvent.PubkeyG2.X, newPubkeyRegistrationEvent.PubkeyG2.Y),
				}
				ops.logger.Debug("Added operator pubkeys to pubkey dict",
					"service", "PubkeyCompendiumServiceInMemory", "operatorAddr", operatorAddr,
					"G1pubkey", pubkeyDict[operatorAddr].G1Pubkey, "G2pubkey", pubkeyDict[operatorAddr].G2Pubkey,
				)
			// Receive a query from GetOperatorPubkeys
			case operatorPubkeyQuery := <-queryC:
				pubkeys, ok := pubkeyDict[operatorPubkeyQuery.operatorAddr]
				operatorPubkeyQuery.respC <- resp{pubkeys, ok}
			}
		}
	}()
}

func (pkcs *OperatorPubkeysServiceInMemory) queryPastRegisteredOperatorEventsAndFillDb(ctx context.Context, pubkeydict map[common.Address]types.OperatorPubkeys) {
	// Querying with nil startBlock and stopBlock will return all events. It doesn't matter if we query some events that we will receive again in the websocket,
	// since we will just overwrite the pubkey dict with the same values.
	alreadyRegisteredOperatorAddrs, alreadyRegisteredOperatorPubkeys, err := pkcs.avsRegistryReader.QueryExistingRegisteredOperatorPubKeys(ctx, nil, nil)
	if err != nil {
		pkcs.logger.Error("Fatal error querying existing registered operators", "err", err, "service", "PubkeyCompendiumServiceInMemory")
		panic(err)
	}
	pkcs.logger.Debug("List of queried operator registration events in pubkey compendium", "alreadyRegisteredOperatorAddr", alreadyRegisteredOperatorAddrs, "service", "PubkeyCompendiumServiceInMemory")

	// Fill the pubkeydict db with the operators and pubkeys found
	for i, operatorAddr := range alreadyRegisteredOperatorAddrs {
		operatorPubkeys := alreadyRegisteredOperatorPubkeys[i]
		pubkeydict[operatorAddr] = operatorPubkeys
	}
}

// TODO(samlaf): we might want to also add an async version of this method that returns a channel of operator pubkeys?
func (pkcs *OperatorPubkeysServiceInMemory) GetOperatorPubkeys(ctx context.Context, operator common.Address) (operatorPubkeys types.OperatorPubkeys, operatorFound bool) {
	respC := make(chan resp)
	pkcs.queryC <- query{operator, respC}
	select {
	case <-ctx.Done():
		return types.OperatorPubkeys{}, false
	case resp := <-respC:
		return resp.operatorPubkeys, resp.operatorExists
	}
}
