package aggregator

import (
	"context"
	"math/big"
	"net/http"
	"net/rpc"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	avsregistryservice "github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

const (
	// number of blocks after which a task is considered expired
	// this hardcoded here because it's also hardcoded in the contracts, but should
	// ideally be fetched from the contracts
	//taskChallengeWindowBlock = 100
	//blockTimeSeconds         = 12 * time.Second
	avsName = "incredible-squaring"
)

type Aggregator struct {
	logger           logging.Logger
	serverIpPortAddr string
	avsWriter        *avsregistry.ChainWriter
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	//tasks                 map[types.TaskIndex]cstaskmanager.IIncredibleSquaringTaskManagerTask
	//tasksMu               sync.RWMutex
	//avsSubscriber         chainio.AvsSubscriberer
	//newTaskCreatedChan    chan *cstaskmanager.ContractIncredibleSquaringTaskManagerNewTaskCreated
	taskProcessor      TaskProcessor
	newTaskCreatedLogs chan types.Log
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator(c AggregatorConfig, taskProcessor TaskProcessor, eventHash common.Hash) (*Aggregator, error) {
	avsConfig := avsregistry.Config{
		RegistryCoordinatorAddress:    c.RegistryCoordinatorAddress,
		OperatorStateRetrieverAddress: c.OperatorStateRetrieverAddress,
		ServiceManagerAddress:         c.ServiceManagerAddress,
	}
	avsReader, err := avsregistry.NewReaderFromConfig(avsConfig, c.EthHttpClient, c.Logger)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	avsWriter, err := avsregistry.NewWriterFromConfig(avsConfig, c.EthHttpClient, c.TxMgr, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create avsWriter", "err", err)
		return nil, err
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
		DontUseAllocationManager: true,
	}

	clients, err := sdkclients.BuildAll(chainioConfig, c.EcdsaPrivateKey, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(
		context.Background(),
		clients.AvsRegistryChainSubscriber,
		clients.AvsRegistryChainReader,
		nil,
		oprsinfoserv.Opts{},
		c.Logger,
	)

	// This is the same hash function used by the operator to hash the task response before signing it.
	hashFunction := func(taskResponse sdktypes.TaskResponse) (sdktypes.TaskResponseDigest, error) {
		// The order here has to match the field ordering of cstaskmanager.IIncredibleSquaringTaskManagerTaskResponse
		taskResponseType, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
			{
				Name: "referenceTaskIndex",
				Type: "uint32",
			},
			{
				Name: "numberSquared",
				Type: "uint256",
			},
		})
		if err != nil {
			c.Logger.Error("Error creating taskResponseType")
			return sdktypes.TaskResponseDigest{}, err
		}
		arguments := abi.Arguments{
			{
				Type: taskResponseType,
			},
		}

		encodeTaskResponseByte, err := arguments.Pack(taskResponse)
		if err != nil {
			c.Logger.Error("Error Packing taskResponse", err)
			return sdktypes.TaskResponseDigest{}, err
		}

		var taskResponseDigest [32]byte
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(encodeTaskResponseByte)
		copy(taskResponseDigest[:], hasher.Sum(nil)[:32])

		return taskResponseDigest, nil
	}

	avsRegistryService := avsregistryservice.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, hashFunction, c.Logger)

	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		c.Logger.Fatal("error connecting to web socket", "err", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{eventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		c.Logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		taskProcessor:         taskProcessor,
		newTaskCreatedLogs:    newTaskCreatedLogs,
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context) error {
	agg.logger.Info("Starting aggregator.")
	agg.logger.Info("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			err := agg.taskProcessor.ProcessAggregatedResponse(context.Background(), blsAggServiceResp)
			if err != nil {
				continue
			}
		case log := <-agg.newTaskCreatedLogs:
			metadata, err := agg.taskProcessor.ProcessNewTask(context.Background(), log)
			if err != nil {
				agg.logger.Fatal("Error processing the task", "err", err)
			}
			if err := agg.blsAggregationService.InitializeNewTask(metadata); err != nil {
				agg.logger.Fatal("Error initializing the task", "err", err)
			}
		}
	}
}

func (agg *Aggregator) startServer(ctx context.Context) {
	err := rpc.Register(agg)
	if err != nil {
		agg.logger.Fatal("Format of service TaskManager isn't correct. ", "err", err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(agg.serverIpPortAddr, nil)
	if err != nil {
		agg.logger.Fatal("ListenAndServe", "err", err)
	}
}

type NonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []regcoord.BN254G1Point
	QuorumApks                   []regcoord.BN254G1Point
	ApkG2                        regcoord.BN254G2Point
	Sigma                        regcoord.BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

type SignedTaskResponse struct {
	TaskResponse TaskResponse
	BlsSignature bls.Signature
	OperatorId   sdktypes.OperatorId
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *Aggregator) ProcessSignedTaskResponse(signedTaskResponse *SignedTaskResponse, reply *bool) error {
	agg.logger.Infof("Received signed task response: %#v", signedTaskResponse)
	taskIndex := signedTaskResponse.TaskResponse.ReferenceTaskIndex

	taskSignature := blsagg.NewTaskSignature(
		taskIndex,
		signedTaskResponse.TaskResponse,
		&signedTaskResponse.BlsSignature,
		signedTaskResponse.OperatorId,
	)

	err := agg.blsAggregationService.ProcessNewSignature(context.Background(), taskSignature)

	return err
}

type TaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

func (tr TaskResponse) TaskIndex() sdktypes.TaskIndex {
	return tr.ReferenceTaskIndex
}

func (tr TaskResponse) Digest() [256]byte {
	return [256]byte(tr.NumberSquared.Bytes())
}

type TaskProcessor interface {
	ProcessNewTask(ctx context.Context, event any) (blsagg.TaskMetadata, error)
	ProcessTaskResponse(ctx context.Context, event TPTaskResponse) ([256]byte, error)
	ProcessAggregatedResponse(ctx context.Context, response blsagg.BlsAggregationServiceResponse) error
}

type TPTaskResponse interface {
	TaskIndex() (sdktypes.TaskIndex)
	Digest() ([256]byte)
}
