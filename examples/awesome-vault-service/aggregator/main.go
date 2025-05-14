package main

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/task-processor"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := taskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
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

	aggConfig := aggregator.AggregatorConfig{
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

	taskManagerAddr := gethcommon.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650")
	taskResponder, err := taskprocessor.NewTaskResponderFromAbi[examplecommon.TaskInput, [32]byte](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
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
