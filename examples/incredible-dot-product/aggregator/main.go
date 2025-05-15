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
	"github.com/pelletier/go-toml"
)

type Config struct {
	aggregator.Config

	TaskManagerAddress string `toml:"task_manager_address"`
}

func GetConfigFromPath(path string) (*Config, error) {
	config := &Config{}
	tree, err := toml.LoadFile(path)
	if err != nil {
		return nil, err
	}
	err = tree.Unmarshal(config)
	return config, err
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	config, err := GetConfigFromPath("config/aggregator_config.toml")
	if err != nil {
		logger.Errorf("Failed to read config file: %w", err)
		return
	}

	taskManagerAbi, err := idptaskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	ethClient, err := ethclient.Dial(config.EthHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	// Depends on the aggregatorPrivateKey passed to TaskManager.initialize
	// To change it, modify aggregator_addr in
	// examples/incredible-dot-product/contracts/config/avs/incredible_dot_product_config.json
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

	aggConfig := config.Config
	aggConfig.EcdsaPrivateKey = ecdsaPrivateKey

	taskManagerAddr := gethcommon.HexToAddress(config.TaskManagerAddress)
	taskResponder, err := taskmanager.NewTaskManagerFromAbi[examplecommon.DotProductInput, *big.Int](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Responder: %w", err)
		return
	}

	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Errorf("Failed to create Task Processor: %w", err)
		return
	}

	aggregator, err := aggregator.NewAggregator(aggConfig, logger, taskProcessor, taskManagerAbi)
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
