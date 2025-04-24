package aggregator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/challenger"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	avsregistryservice "github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
)

const (
	// number of blocks after which a task is considered expired this hardcoded here because it's also
	//  hardcoded in the contracts, but should ideally be fetched from the contracts
	taskChallengeWindowBlock = 100
	blockTimeSeconds         = 12 * time.Second
)

type TaskProcessor[Input any] interface {
	ProcessAggregatedResponse(ctx context.Context, response blsagg.BlsAggregationServiceResponse, task challenger.GenericInputTask[Input]) error
}

type Aggregator[Input any, Output any] struct {
	logger           logging.Logger
	serverIpPortAddr string

	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	taskProcessor         TaskProcessor[Input]
	newTaskCreatedLogs    chan types.Log

	tasks         map[sdktypes.TaskIndex]challenger.GenericInputTask[Input]
	tasksMu       sync.RWMutex
	taskResponses map[uint32]challenger.TaskResponseData[Output]

	taskManagerAbi *abi.ABI
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator[Input any, Output any](
	c AggregatorConfig,
	taskProcessor TaskProcessor[Input],
	eventHash common.Hash,
	taskManagerAbi *abi.ABI,
) (*Aggregator[Input, Output], error) {
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

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress.String(),
		AvsName:                    "Aggregator",
		PromMetricsIpPortAddress:   ":9090",
		DontUseAllocationManager:   true,
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

	if c.TaskResponseHashFn == nil {
		return nil, errors.New("task response hash function not provided in aggregator config")
	}

	avsRegistryService := avsregistryservice.NewAvsRegistryServiceChainCaller(avsReader, operatorPubkeysService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.TaskResponseHashFn, c.Logger)

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

	return &Aggregator[Input, Output]{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		blsAggregationService: blsAggregationService,
		taskProcessor:         taskProcessor,
		newTaskCreatedLogs:    newTaskCreatedLogs,
		tasks:                 make(map[sdktypes.TaskIndex]challenger.GenericInputTask[Input]),
		taskResponses:         make(map[uint32]challenger.TaskResponseData[Output]),
		taskManagerAbi:        c.TaskManagerAbi,
	}, nil
}

func (agg *Aggregator[Input, Output]) Start(ctx context.Context) error {
	agg.logger.Info("Starting aggregator.")
	agg.logger.Info("Starting aggregator rpc server.")
	go agg.startServer(ctx)

	for {
		select {
		case <-ctx.Done():
			return nil
		case blsAggServiceResp := <-agg.blsAggregationService.GetResponseChannel():
			agg.logger.Info("Received response from blsAggregationService", "blsAggServiceResp", blsAggServiceResp)
			err := agg.processAggregatedResponse(context.Background(), blsAggServiceResp)
			if err != nil {
				continue
			}
		case log := <-agg.newTaskCreatedLogs:
			metadata, err := agg.processNewTask(context.Background(), log)
			if err != nil {
				agg.logger.Fatal("Error processing the task", "err", err)
			}
			if err := agg.blsAggregationService.InitializeNewTask(metadata); err != nil {
				agg.logger.Fatal("Error initializing the task", "err", err)
			}
		}
	}
}

func (agg *Aggregator[Input, Output]) processNewTask(ctx context.Context, log types.Log) (blsagg.TaskMetadata, error) {
	var newTaskCreatedLog challenger.NewTaskCreatedEvent[Input]

	err := agg.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return blsagg.TaskMetadata{}, fmt.Errorf("error unpacking the log: %w", err)
	}

	// This is done this way because the taskIndex value in this event is indexed, so we take it from the log
	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())

	agg.logger.Infof("Aggregator received new task: %v: ", newTaskCreatedLog)

	newTask := newTaskCreatedLog.Task
	agg.tasksMu.Lock()
	agg.tasks[newTaskIndex] = newTask
	agg.tasksMu.Unlock()

	quorumThresholdPercentages := make(sdktypes.QuorumThresholdPercentages, len(newTask.QuorumNumbers))
	for i := range newTask.QuorumNumbers {
		quorumThresholdPercentages[i] = sdktypes.QuorumThresholdPercentage(newTask.QuorumThresholdPercentage)
	}
	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block
	// number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeSeconds
	var quorumNums sdktypes.QuorumNums
	for _, quorumNum := range newTask.QuorumNumbers {
		quorumNums = append(quorumNums, sdktypes.QuorumNum(quorumNum))
	}
	metadata := blsagg.NewTaskMetadata(
		newTaskIndex,
		newTask.TaskCreatedBlock,
		quorumNums,
		quorumThresholdPercentages,
		taskTimeToExpiry,
	)

	return metadata, nil
}

func (agg *Aggregator[Input, Output]) processAggregatedResponse(
	ctx context.Context,
	response blsagg.BlsAggregationServiceResponse,
) error {
	if response.Err != nil {
		return utils.WrapError("BlsAggregationServiceResponse contains an error", response.Err)
	}

	agg.tasksMu.RLock()
	task := agg.tasks[response.TaskIndex]
	agg.tasksMu.RUnlock()

	err := agg.taskProcessor.ProcessAggregatedResponse(ctx, response, task)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}
	return nil
}
