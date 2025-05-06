package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/task-processor"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

type DotProductOutput struct {
	Result *big.Int
}

func main() {
	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		panic(err)
	}

	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
	}

	ethHttpUrl := "127.0.0.1:8545"
	ethClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
	}

	aggConfig := aggregator.AggregatorConfig{
		Logger:             logger,
		TaskManagerAbi:     taskManagerAbi,
		TaskResponseHashFn: nil,
		EthHttpUrl:         ethHttpUrl,
		EthHttpClient:      ethClient,
	}

	// txMgr, err := txmgr.NewSimpleTxManager()

	taskManagerAddr := gethcommon.HexToAddress("0x")
	taskResponder, err := taskprocessor.NewTaskResponderFromAbi[DotProductInput, DotProductOutput](taskManagerAddr, taskManagerAbi, &txmgr.SimpleTxManager{}, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Responder: %w", err)
	}

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Errorf("Failed to create Task Processor: %w", err)
	}

	aggregator, err := aggregator.NewAggregator(aggConfig, taskProcessor)
	if err != nil {
		logger.Errorf("Failed to create aggregator: %w", err)
	}

	err = aggregator.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running aggregator: %w", err)
	}
}
