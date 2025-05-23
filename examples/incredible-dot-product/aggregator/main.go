package main

import (
	"context"
	"math/big"
	"os"

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
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	aggregator.Config

	TaskManagerAddress string `toml:"task_manager_address"`
}

func GetConfigFromPath(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = toml.Unmarshal(data, config)
	return config, err
}

// This is the main function for the aggregator in the incredible dot product example. The steps followed are
// also explained in the aggregator module readme, which can be found at aggregator/README.md
func main() {
	// 0. Create the logger where all the logs will appear
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 1. Create the aggregator configuration (in this case we read it from aggregator config file)
	config, err := GetConfigFromPath("config/aggregator_config.toml")
	if err != nil {
		logger.Errorf("Failed to read config file: %w", err)
		return
	}
	aggConfig := config.Config

	// 2. Provide a Task Processor, first instantiating the things required for creating it

	// i. Create the ethereum client that will send the RPC messages to the node, and the transaction
	// manager, that will manage the transaction sending
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

	// ii. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := idptaskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	// iii. Provide a struct that implements the `TaskResponder` interface. Here we use an SDK implementation
	// that already satisfies it, but you can also provide your own type implementing the interface.
	// Note that in this step we define the input and output types that we are using on our AVS. In this case
	// the DotProductInput struct (a pair of vectors) and a big int.
	taskManagerAddr := gethcommon.HexToAddress(config.TaskManagerAddress)
	taskResponder, err := taskmanager.NewTaskManagerFromAbi[examplecommon.DotProductInput, *big.Int](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Responder: %w", err)
		return
	}

	// iv. Create the Task Processor. Here we use the IndexingTaskProcessor, you can see its implementation
	// in aggregator/task-processor/indexing_task_processor.go
	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Errorf("Failed to create Task Processor: %w", err)
		return
	}

	// 3. Build the aggregator, providing aggregator config, logger, task processor and the task manager ABI, and
	// then start it.
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
