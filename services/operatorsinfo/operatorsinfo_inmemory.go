package operatorsinfo

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

var defaultLogFilterQueryBlockRange = big.NewInt(10_000)

// OperatorsInfoServiceInMemory is a stateful goroutine (see https://gobyexample.com/stateful-goroutines)
// implementation of OperatorsInfoService that listen for the NewPubkeyRegistration and OperatorSocketUpdate events using a websocket connection
// to an eth client and stores the pubkeys/sockets in memory. Another possible implementation is using a mutex
// (https://gobyexample.com/mutexes) instead. We can switch to that if we ever find a good reason to.
//
// Warning: this service should probably not be used in production. Haven't done a thorough analysis of all the clients
// but there is still an open PR about an issue with ws subscription on geth: https://github.com/ethereum/go-ethereum/issues/23845
// Another reason to note for infra/devops engineer who would put this into production, is that this service crashes on
// websocket connection errors or when failing to query past events. The philosophy here is that hard crashing is
// better than silently failing, since it will be easier to debug. Naturally, this means that this aggregator using this service needs
// to be replicated and load-balanced, so that when it fails traffic can be switched to the other aggregator.
type OperatorsInfoServiceInMemory struct {
	logFilterQueryBlockRange *big.Int
	avsRegistrySubscriber    avsregistry.AvsRegistrySubscriber
	avsRegistryReader        avsregistry.AvsRegistryReader
	logger                   logging.Logger
	queryC                   chan<- query
	// queried via the queryC channel, so don't need mutex to access
	pubkeyDict       map[common.Address]types.OperatorPubkeys
	operatorAddrToId map[common.Address]types.OperatorId
	socketDict       map[types.OperatorId]types.Socket
}
type query struct {
	operatorAddr common.Address
	// channel through which to receive the resp
	respC chan<- resp
}
type resp struct {
	// TODO: possible for socket to be empty string if haven't received the event yet... would be a crazy race condition though.
	operatorInfo types.OperatorInfo
	// false if operators were not present in the pubkey dict
	operatorExists bool
}

var _ OperatorsInfoService = (*OperatorsInfoServiceInMemory)(nil)

// NewOperatorsInfoServiceInMemory constructs a OperatorsInfoServiceInMemory and starts it in a goroutine.
// It takes a context as argument because the "backfilling" of the database is done inside this constructor,
// so we wait for all past NewPubkeyRegistration/OperatorSocketUpdate events to be queried and the db to be filled before returning the service.
// The constructor is thus following a RAII-like pattern, of initializing the serving during construction.
// Using a separate initialize() function might lead to some users forgetting to call it and the service not behaving properly.
func NewOperatorsInfoServiceInMemory(
	ctx context.Context,
	avsRegistrySubscriber avsregistry.AvsRegistrySubscriber,
	avsRegistryReader avsregistry.AvsRegistryReader,
	logFilterQueryBlockRange *big.Int,
	logger logging.Logger,
) *OperatorsInfoServiceInMemory {
	queryC := make(chan query)
	if logFilterQueryBlockRange == nil {
		logFilterQueryBlockRange = defaultLogFilterQueryBlockRange
	}
	pkcs := &OperatorsInfoServiceInMemory{
		avsRegistrySubscriber:    avsRegistrySubscriber,
		avsRegistryReader:        avsRegistryReader,
		logFilterQueryBlockRange: logFilterQueryBlockRange,
		logger:                   logger,
		queryC:                   queryC,
		pubkeyDict:               make(map[common.Address]types.OperatorPubkeys),
		operatorAddrToId:         make(map[common.Address]types.OperatorId),
		socketDict:               make(map[types.OperatorId]types.Socket),
	}
	// We use this waitgroup to wait on the initialization of the inmemory pubkey dict,
	// which requires querying the past events of the pubkey registration contract
	wg := sync.WaitGroup{}
	wg.Add(1)
	pkcs.startServiceInGoroutine(ctx, queryC, &wg)
	wg.Wait()
	return pkcs
}

func (ops *OperatorsInfoServiceInMemory) startServiceInGoroutine(ctx context.Context, queryC <-chan query, wg *sync.WaitGroup) {
	go func() {

		// TODO(samlaf): we should probably save the service in the logger itself and add it automatically to all logs
		ops.logger.Debug("Subscribing to new pubkey registration events on blsApkRegistry contract", "service", "OperatorPubkeysServiceInMemory")
		newPubkeyRegistrationC, newPubkeyRegistrationSub, err := ops.avsRegistrySubscriber.SubscribeToNewPubkeyRegistrations()
		if err != nil {
			ops.logger.Error("Fatal error opening websocket subscription for new pubkey registrations", "err", err, "service", "OperatorPubkeysServiceInMemory")
			// see the warning above the struct definition to understand why we panic here
			panic(err)
		}
		newSocketRegistrationC, newSocketRegistrationSub, err := ops.avsRegistrySubscriber.SubscribeToOperatorSocketUpdates()
		if err != nil {
			ops.logger.Error("Fatal error opening websocket subscription for new socket registrations", "err", err, "service", "OperatorPubkeysServiceInMemory")
			panic(err)
		}
		err = ops.queryPastRegisteredOperatorEventsAndFillDb(ctx)
		if err != nil {
			ops.logger.Error("Fatal error querying past registered operator events and filling db", "err", err, "service", "OperatorPubkeysServiceInMemory")
			panic(err)
		}
		// The constructor can return after we have backfilled the db by querying the events of operators that have registered with the blsApkRegistry
		// before the block at which we started the ws subscription above
		wg.Done()
		for {
			select {
			case <-ctx.Done():
				// TODO(samlaf): should we do anything here? Seems like this only happens when the aggregator is shutting down and we want graceful exit
				ops.logger.Infof("OperatorPubkeysServiceInMemory: Context cancelled, exiting")
				return
			case err := <-newPubkeyRegistrationSub.Err():
				ops.logger.Error("Error in websocket subscription for new pubkey registration events. Attempting to reconnect...", "err", err, "service", "OperatorPubkeysServiceInMemory")
				newPubkeyRegistrationSub.Unsubscribe()
				newPubkeyRegistrationC, newPubkeyRegistrationSub, err = ops.avsRegistrySubscriber.SubscribeToNewPubkeyRegistrations()
				if err != nil {
					ops.logger.Error("Error opening websocket subscription for new pubkey registrations", "err", err, "service", "OperatorPubkeysServiceInMemory")
					// see the warning above the struct definition to understand why we panic here
					panic(err)
				}
			case err := <-newSocketRegistrationSub.Err():
				ops.logger.Error("Error in websocket subscription for new socket registration events. Attempting to reconnect...", "err", err, "service", "OperatorPubkeysServiceInMemory")
				newSocketRegistrationSub.Unsubscribe()
				newSocketRegistrationC, newSocketRegistrationSub, err = ops.avsRegistrySubscriber.SubscribeToOperatorSocketUpdates()
				if err != nil {
					ops.logger.Error("Error opening websocket subscription for new socket registrations", "err", err, "service", "OperatorPubkeysServiceInMemory")
					panic(err)
				}
			case newPubkeyRegistrationEvent := <-newPubkeyRegistrationC:
				operatorAddr := newPubkeyRegistrationEvent.Operator
				ops.pubkeyDict[operatorAddr] = types.OperatorPubkeys{
					G1Pubkey: bls.NewG1Point(newPubkeyRegistrationEvent.PubkeyG1.X, newPubkeyRegistrationEvent.PubkeyG1.Y),
					G2Pubkey: bls.NewG2Point(newPubkeyRegistrationEvent.PubkeyG2.X, newPubkeyRegistrationEvent.PubkeyG2.Y),
				}

				operatorId := types.OperatorIdFromContractG1Pubkey(newPubkeyRegistrationEvent.PubkeyG1)
				ops.operatorAddrToId[operatorAddr] = operatorId
				ops.logger.Debug("Added operator pubkeys to pubkey dict from new pubkey registration event",
					"service", "OperatorPubkeysServiceInMemory", "block", newPubkeyRegistrationEvent.Raw.BlockNumber,
					"operatorAddr", operatorAddr, "operatorId", operatorId,
					"G1pubkey", ops.pubkeyDict[operatorAddr].G1Pubkey, "G2pubkey", ops.pubkeyDict[operatorAddr].G2Pubkey,
				)
			case newSocketRegistrationEvent := <-newSocketRegistrationC:
				ops.logger.Debug("Received new socket registration event", "service", "OperatorPubkeysServiceInMemory", "operatorId", types.OperatorId(newSocketRegistrationEvent.OperatorId), "socket", newSocketRegistrationEvent.Socket)
				ops.updateSocketMapping(newSocketRegistrationEvent.OperatorId, types.Socket(newSocketRegistrationEvent.Socket))
			// Receive a query from GetOperatorPubkeys
			case query := <-queryC:
				pubkeys, ok := ops.pubkeyDict[query.operatorAddr]
				operatorId := ops.operatorAddrToId[query.operatorAddr]
				socket := ops.socketDict[operatorId]
				operatorInfo := types.OperatorInfo{
					Socket:  socket,
					Pubkeys: pubkeys,
				}
				query.respC <- resp{operatorInfo, ok}
			}
		}
	}()
}

func (ops *OperatorsInfoServiceInMemory) queryPastRegisteredOperatorEventsAndFillDb(ctx context.Context) error {
	// Querying with nil startBlock and stopBlock will return all events. It doesn't matter if we query some events that we will receive again in the websocket,
	// since we will just overwrite the pubkey dict with the same values.
	wg := sync.WaitGroup{}
	var alreadyRegisteredOperatorAddrs []common.Address
	var alreadyRegisteredOperatorPubkeys []types.OperatorPubkeys
	var pubkeysErr error

	// we make both Queries in parallel because they take time and we don't want to wait for one to finish before starting the other
	wg.Add(2)
	go func() {
		alreadyRegisteredOperatorAddrs, alreadyRegisteredOperatorPubkeys, pubkeysErr = ops.avsRegistryReader.QueryExistingRegisteredOperatorPubKeys(ctx, nil, nil, ops.logFilterQueryBlockRange)
		wg.Done()
	}()
	var socketsMap map[types.OperatorId]types.Socket
	var socketsErr error
	go func() {
		socketsMap, socketsErr = ops.avsRegistryReader.QueryExistingRegisteredOperatorSockets(ctx, nil, nil, ops.logFilterQueryBlockRange)
		wg.Done()
	}()

	wg.Wait()
	if pubkeysErr != nil {
		return errors.Join(errors.New("error querying existing registered operators"), pubkeysErr)
	}
	if socketsErr != nil {
		return errors.Join(errors.New("error querying existing registered operator sockets"), socketsErr)
	}
	ops.logger.Debug("List of queried operator registration events in blsApkRegistry", "alreadyRegisteredOperatorAddr", alreadyRegisteredOperatorAddrs, "alreadyRegisteredOperatorPubkeys", alreadyRegisteredOperatorPubkeys, "service", "OperatorPubkeysServiceInMemory")
	ops.logger.Debug("List of queried operator socket registration events", "socketsMap", socketsMap, "service", "OperatorPubkeysServiceInMemory")

	// Fill the pubkeydict db with the operators and pubkeys found
	for i, operatorAddr := range alreadyRegisteredOperatorAddrs {
		operatorPubkeys := alreadyRegisteredOperatorPubkeys[i]
		ops.pubkeyDict[operatorAddr] = operatorPubkeys
		operatorId := types.OperatorIdFromG1Pubkey(operatorPubkeys.G1Pubkey)
		ops.operatorAddrToId[operatorAddr] = operatorId
		ops.updateSocketMapping(operatorId, socketsMap[operatorId])
	}
	return nil
}

// TODO(samlaf): we might want to also add an async version of this method that returns a channel of operator pubkeys?
func (ops *OperatorsInfoServiceInMemory) GetOperatorInfo(ctx context.Context, operator common.Address) (operatorPubkeys types.OperatorInfo, operatorFound bool) {
	respC := make(chan resp)
	ops.queryC <- query{operator, respC}
	select {
	case <-ctx.Done():
		return types.OperatorInfo{}, false
	case resp := <-respC:
		return resp.operatorInfo, resp.operatorExists
	}
}

func (ops *OperatorsInfoServiceInMemory) updateSocketMapping(operatorId types.OperatorId, socket types.Socket) {
	if socket == "" {
		ops.logger.Warn("Received empty socket for operator", "operatorId", operatorId)
		return
	}
	ops.socketDict[operatorId] = types.Socket(socket)
}
