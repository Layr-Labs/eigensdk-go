package aggregator

import (
	"context"
	"errors"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	avsregistryservice "github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	oprsinfoserv "github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
)

type TaskProcessor interface {
	ProcessNewTask(ctx context.Context, log types.Log) (blsagg.TaskMetadata, error)
	ProcessTaskResponse(ctx context.Context, event TaskResponse) ([32]byte, error)
	ProcessAggregatedResponse(ctx context.Context, response blsagg.BlsAggregationServiceResponse) error
}

type Aggregator struct {
	logger           logging.Logger
	serverIpPortAddr string
	avsWriter        *avsregistry.ChainWriter
	// aggregation related fields
	blsAggregationService blsagg.BlsAggregationService
	taskProcessor         TaskProcessor
	newTaskCreatedLogs    chan types.Log
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

	return &Aggregator{
		logger:                c.Logger,
		serverIpPortAddr:      c.AggregatorServerIpPortAddr,
		avsWriter:             avsWriter,
		blsAggregationService: blsAggregationService,
		taskProcessor:         taskProcessor,
		newTaskCreatedLogs:    newTaskCreatedLogs,
	}, nil
}

func (agg *Aggregator) Start(ctx context.Context, taskResponseType interface{}) error {
	agg.logger.Info("Starting aggregator.")
	agg.logger.Info("Starting aggregator rpc server.")
	go agg.startServer(ctx, taskResponseType)

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
