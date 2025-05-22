package main

import (
	"context"
	"os"

	"github.com/Layr-Labs/eigensdk-go/aggregator"
	taskprocessor "github.com/Layr-Labs/eigensdk-go/aggregator/task-processor"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	avtaskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"

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

func main() {
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Get the config from the aggregator config file, reading the required information to create the aggregator
	config, err := GetConfigFromPath("config/aggregator_config.toml")
	if err != nil {
		logger.Errorf("Failed to read config file: %w", err)
		return
	}

	// 3. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := avtaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	// 4. Create the ethereum client that will send the RPC messages to the node
	ethClient, err := ethclient.Dial(config.EthHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	// 5. Create the transaction manager, that will manage the transaction sending

	// Depends on the aggregatorPrivateKey passed to TaskManager.initialize
	// To change it, modify aggregator_addr in
	// examples/awesome-vault-service/contracts/config/avs/incredible_dot_product_config.json
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

	// 6. Create the task responder, which will send the aggregated responses to the on-chain TaskManager contract.
	// Here we use an SDK implementation that satisfies the TaskResponder interface, but you can create your own wrapper
	// which implements the interface and provide it to the Indexing Task Processor.
	taskManagerAddr := gethcommon.HexToAddress(config.TaskManagerAddress)

	// Note that in this step we define the input and output types that we are using on our AVS. In this case
	// the TaskInput struct (a key-value pair) and 32 bytes.
	taskResponder, err := taskmanager.NewTaskManagerFromAbi[examplecommon.TaskInput, [32]byte](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Responder: %w", err)
		return
	}

	// 7. Create the Task Processor, which will manage the processing of the tasks and the aggregated responses.
	// Here we use the IndexingTaskProcessor, a generic implementation provided by the SDK that saves the tasks
	// in a map and delegates the sending of aggregated responses to the on-chain TaskManager contract.
	taskProcessor, err := taskprocessor.NewIndexingTaskProcessor(logger, taskResponder)
	if err != nil {
		logger.Errorf("Failed to create Task Processor: %w", err)
		return
	}

	// 8. Build the aggregator, providing aggregator config, logger, task processor and the task manager ABI.
	aggregator, err := aggregator.NewAggregator(aggConfig, logger, taskProcessor, taskManagerAbi)
	if err != nil {
		logger.Errorf("Failed to create aggregator: %w", err)
		return
	}

	// 9. Run the created aggregator
	err = aggregator.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running aggregator: %w", err)
		return
	}
}
