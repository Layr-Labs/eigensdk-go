package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/aggregator/task-processor"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	idptaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/common"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := idptaskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	ethHttpUrl := "http://localhost:8545"
	ethClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	aggregatorPrivateKey := "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"
	ecdsaPrivateKey, err := crypto.HexToECDSA(aggregatorPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create tx manager from private key: %w", err)
		return
	}

	aggConfig := aggregator.Config{
		Logger:             logger,
		TaskManagerAbi:     taskManagerAbi,
		TaskResponseHashFn: nil,

		EthHttpUrl:                 ethHttpUrl,
		EthWsUrl:                   "ws://localhost:8545",
		AggregatorServerIpPortAddr: "localhost:8090",

		RegistryCoordinatorAddress:    gethcommon.HexToAddress("0xfd471836031dc5108809d173a067e8486b9047a3"),
		OperatorStateRetrieverAddress: gethcommon.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		ServiceManagerAddress:         gethcommon.HexToAddress("0xcd8a1c3ba11cf5ecfa6267617243239504a98d90"),

		EthHttpClient:   ethClient,
		EcdsaPrivateKey: ecdsaPrivateKey,
	}

	taskManagerAddr := gethcommon.HexToAddress("0x2bdcc0de6be1f7d2ee689a0342d76f52e8efaba3")
	taskResponder, err := taskmanager.NewTaskManagerFromAbi[examplecommon.DotProductInput, *big.Int, any](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Responder: %w", err)
		return
	}

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Errorf("Failed to create Task Processor: %w", err)
		return
	}

	aggregator, err := aggregator.NewAggregator(aggConfig, taskProcessor)
	if err != nil {
		logger.Errorf("Failed to create aggregator: %w", err)
		return
	}

	err = aggregator.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running aggregator: %w", err)
		return
	}
}
