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

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	avsregistryservice "github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
)

// The aggregator is responsible for aggregating signed task responses from operators and posting them on chain. This includes:
//   - Listening to new task created events.
//   - Receiving signed responses from the operators.
//   - Sending the aggregated responses to the `TaskManager` contract
//
// Most of these things are delegated to the `TaskProcessor` interface, that processes
// tasks and communicates with the on-chain `TaskManager` contract.
type Aggregator[Input any, Output any] struct {
	logger logging.Logger

	// BLS aggregation service
	blsAggregationService blsagg.BlsAggregationService

	// Channel for receiving new task created event logs
	newTaskCreatedLogs chan types.Log

	// ABI of the task manager contract
	taskManagerAbi *abi.ABI

	aggregatorRpcServer *AggregatorRpcServer[Input, Output]

	taskProcessor taskprocessor.TaskProcessor[Input, Output]
}

// NewAggregator creates a new Aggregator with the provided config, a logger, a task processor and the task
// manager contract's ABI.
func NewAggregator[Input any, Output any](
	logger logging.Logger,
	c Config,
	taskManagerAbi *abi.ABI,
	taskProcessor taskprocessor.TaskProcessor[Input, Output],
) (*Aggregator[Input, Output], error) {
	avsRegistryConfig := avsregistry.Config{
		RegistryCoordinatorAddress:    c.RegistryCoordinatorAddress,
		OperatorStateRetrieverAddress: c.OperatorStateRetrieverAddress,
	}

	wsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil {
		logger.Fatal("error connecting to web socket", "err", err)
	}

	avsRegistrySubscriber, err := avsregistry.NewSubscriberFromConfig(avsRegistryConfig, wsClient, logger)
	if err != nil {
		logger.Fatal("Failed to create avs registry subscriber", "err", err)
	}

	httpClient, err := ethclient.Dial(c.EthHttpUrl)
	if err != nil {
		logger.Fatal("error connecting to http client", "err", err)
	}

	avsRegistryReader, err := avsregistry.NewReaderFromConfig(avsRegistryConfig, httpClient, logger)
	if err != nil {
		logger.Fatal("Failed to create avs registry reader", "err", err)
	}

	operatorPubkeysService := oprsinfoserv.NewOperatorsInfoServiceInMemory(
		context.Background(),
		avsRegistrySubscriber,
		avsRegistryReader,
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

	avsRegistryService := avsregistryservice.NewAvsRegistryServiceChainCaller(avsRegistryReader, operatorPubkeysService, logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, taskResponseHashFn, logger)

	newTaskCreatedEventHash := taskManagerAbi.Events["NewTaskCreated"].ID
	query := ethereum.FilterQuery{
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{{newTaskCreatedEventHash}},
	}

	newTaskCreatedLogs := make(chan types.Log)
	_, err = wsClient.SubscribeFilterLogs(context.Background(), query, newTaskCreatedLogs)
	if err != nil {
		logger.Fatal("error subscribing to newTaskCreated events", "err", err)
	}

	rpcServer := NewAggregatorRpcServer[Input, Output](logger, c.AggregatorServerIpPortAddr, blsAggregationService)

	return &Aggregator[Input, Output]{
		logger:                logger,
		blsAggregationService: blsAggregationService,
		newTaskCreatedLogs:    newTaskCreatedLogs,
		taskManagerAbi:        taskManagerAbi,
		taskProcessor:         taskProcessor,
		aggregatorRpcServer:   rpcServer,
	}, nil
}

// Runs the Aggregator in a separate goroutine. This should be called only one time per Aggregator.
// Will return an error if execution fails or nil in case the context is cancelled.
func (agg *Aggregator[Input, Output]) Start(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		errChan <- agg.run(ctx)
	}()

	return errChan
}

// The run method contains the main loop of the Aggregator, that has 2 main events:
//   - Get a response from the BLS aggregation service: In this case the response is processed and sent to
//     the Task Manager on-chain contract.
//   - Receive a new task created event log: In this case the aggregator processes that event, and sends to
//     the BLS aggregation service the new task created metadata.
func (agg *Aggregator[Input, Output]) run(ctx context.Context) error {
	agg.logger.Info("Starting aggregator.")
	agg.logger.Info("Starting aggregator rpc server.")

	serverErrorChannel := make(chan error)
	go func() {
		serverErrorChannel <- agg.aggregatorRpcServer.StartServer()
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-serverErrorChannel:
			return err
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

// When processing a new task event, the aggregator unpacks the log data into the new task created event and
// sends it to the task processor
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

// When processing an aggregated response, the aggregator delegates the processing to the task processor
func (agg *Aggregator[Input, Output]) processAggregatedResponse(
	response blsagg.BlsAggregationServiceResponse,
) error {
	if response.Err != nil {
		return utils.WrapError("BlsAggregationServiceResponse contains an error", response.Err)
	}

	nonSignerPubkeys := []sdktypes.BN254G1Point{}
	for _, nonSignerPubkey := range response.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, sdktypes.ConvertToBN254G1Point(nonSignerPubkey))
	}
	quorumApks := []sdktypes.BN254G1Point{}
	for _, quorumApk := range response.QuorumApksG1 {
		quorumApks = append(quorumApks, sdktypes.ConvertToBN254G1Point(quorumApk))
	}
	nonSignerStakesAndSignature := sdktypes.NonSignerStakesAndSignature{
		NonSignerPubkeys:             nonSignerPubkeys,
		QuorumApks:                   quorumApks,
		ApkG2:                        sdktypes.ConvertToBN254G2Point(response.SignersApkG2),
		Sigma:                        sdktypes.ConvertToBN254G1Point(response.SignersAggSigG1.G1Point),
		NonSignerQuorumBitmapIndices: response.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             response.QuorumApkIndices,
		TotalStakeIndices:            response.TotalStakeIndices,
		NonSignerStakeIndices:        response.NonSignerStakeIndices,
	}

	taskResponse, ok := response.TaskResponse.(taskmanager.TaskResponse[Output])
	if !ok {
		agg.logger.Error("task Response could not be converted to sdk aggregator's Task Response type")
	}

	err := agg.taskProcessor.ProcessAggregatedResponse(response.TaskIndex, taskResponse, nonSignerStakesAndSignature)
	if err != nil {
		return utils.WrapError("Aggregator failed to respond to task", err)
	}
	return nil
}
