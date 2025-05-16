package aggregator

import (
	"context"
	"fmt"
	"math/big"

	taskprocessor "github.com/Layr-Labs/eigensdk-go/aggregator/task-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	avsregistryservice "github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
)

type Aggregator[Input any, Output any] struct {
	logger           logging.Logger

	// The port exposed by the aggregator to listen to operator task responses
	serverIpPortAddr string

	// bls aggregation service
	blsAggregationService blsagg.BlsAggregationService

	// channel that receives new task created event logs
	newTaskCreatedLogs chan types.Log

	// Abi of the task manager contract
	taskManagerAbi *abi.ABI

	taskProcessor taskprocessor.TaskProcessor[Input, Output]
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregator[Input any, Output any](
	c Config,
	logger logging.Logger,
	taskProcessor taskprocessor.TaskProcessor[Input, Output],
	taskManagerAbi *abi.ABI,
) (*Aggregator[Input, Output], error) {
	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpUrl,
		EthWsUrl:                   c.EthWsUrl,
		RegistryCoordinatorAddr:    c.RegistryCoordinatorAddress.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddress.String(),
		AvsName:                    "Aggregator",
		PromMetricsIpPortAddress:   ":9090",
		DontUseAllocationManager:   true,
	}

	clients, err := sdkclients.BuildAll(chainioConfig, c.EcdsaPrivateKey, logger)
	if err != nil {
		logger.Errorf("Cannot create sdk clients. Err: %w", err)
		return nil, err
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(
		context.Background(),
		clients.AvsRegistryChainSubscriber,
		clients.AvsRegistryChainReader,
		nil,
		oprsinfoserv.Opts{},
		logger,
	)

	taskResponseHashFn := func(response any) (sdktypes.TaskResponseDigest, error) {
		taskResponse, ok := response.(taskmanager.TaskResponse[Output])
		if !ok {
			logger.Error("task Response could not be converted to sdk aggregator's Task Response type")
		}

		return taskProcessor.ProcessTaskResponse(taskResponse)
	}

	avsRegistryService := avsregistryservice.NewAvsRegistryServiceChainCaller(clients.AvsRegistryChainReader, operatorPubkeysService, logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, taskResponseHashFn, logger)

	client, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatal("error connecting to web socket", "err", err)
	}

	newTaskCreatedEventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{newTaskCreatedEventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	return &Aggregator[Input, Output]{
		logger:                logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		blsAggregationService: blsAggregationService,
		newTaskCreatedLogs:    newTaskCreatedLogs,
		taskManagerAbi:        taskManagerAbi,
		taskProcessor:         taskProcessor,
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
			err := agg.processAggregatedResponse(blsAggServiceResp)
			if err != nil {
				agg.logger.Errorf("Failed to process aggregated response: %w", err)
				continue
			}
		case log := <-agg.newTaskCreatedLogs:
			metadata, err := agg.processNewTask(log)
			if err != nil {
				agg.logger.Fatal("Error processing the task", "err", err)
			}
			if err := agg.blsAggregationService.InitializeNewTask(metadata); err != nil {
				agg.logger.Fatal("Error initializing the task", "err", err)
			}
		}
	}
}

func (agg *Aggregator[Input, Output]) processNewTask(log types.Log) (blsagg.TaskMetadata, error) {
	var newTaskCreatedLog taskmanager.NewTaskCreatedEvent[Input]

	err := agg.taskManagerAbi.UnpackIntoInterface(&newTaskCreatedLog, "NewTaskCreated", log.Data)
	if err != nil {
		return blsagg.TaskMetadata{}, fmt.Errorf("error unpacking the log: %w", err)
	}

	// This is done this way because the taskIndex value in this event is indexed, so we take it from the log
	newTaskIndex := uint32(new(big.Int).SetBytes(log.Topics[1].Bytes()).Uint64())

	agg.logger.Infof("Aggregator received new task: %v: ", newTaskCreatedLog)

	newTask := newTaskCreatedLog.Task

	metadata, err := agg.taskProcessor.ProcessNewTask(newTaskIndex, newTask)
	if err != nil {
		return blsagg.TaskMetadata{}, err
	}

	return metadata, nil
}

func (agg *Aggregator[Input, Output]) processAggregatedResponse(
	response blsagg.BlsAggregationServiceResponse,
) error {
	if response.Err != nil {
		return utils.WrapError("BlsAggregationServiceResponse contains an error", response.Err)
	}

	err := agg.taskProcessor.ProcessAggregatedResponse(response)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}
	return nil
}
