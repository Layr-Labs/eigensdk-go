package main

import (
	"context"

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

	gotoml "github.com/pelletier/go-toml"
)

type Config struct {
	EthHttpUrl string `toml:"eth_http_url"`
	EthWsUrl   string `toml:"eth_ws_url"`

	AggregatorPrivateKey   string `toml:"aggregator_private_key"`
	AggregatorServerIPPort string `toml:"aggregator_server_ip_port"`

	RegistryCoordinatorAddress    string `toml:"registry_coordinator_address"`
	OperatorStateRetrieverAddress string `toml:"operator_state_retriever_address"`
	ServiceManagerAddress         string `toml:"service_manager_address"`
	TaskManagerAddress            string `toml:"task_manager_address"`
}

func GetConfigFromPath(path string) (*Config, error) {
	config := &Config{}
	tree, err := gotoml.LoadFile(path)
	if err != nil {
		return nil, err
	}
	err = tree.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
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

	taskManagerAbi, err := avtaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	ethClient, err := ethclient.Dial(config.EthHttpUrl)
	if err != nil {
		logger.Errorf("Failed to dial ethclient: %w", err)
		return
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.AggregatorPrivateKey)
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
		EthHttpUrl:                 config.EthHttpUrl,
		EthWsUrl:                   config.EthWsUrl,
		AggregatorServerIpPortAddr: config.AggregatorServerIPPort,

		RegistryCoordinatorAddress:    gethcommon.HexToAddress(config.RegistryCoordinatorAddress),
		OperatorStateRetrieverAddress: gethcommon.HexToAddress(config.OperatorStateRetrieverAddress),
		ServiceManagerAddress:         gethcommon.HexToAddress(config.ServiceManagerAddress),

		EcdsaPrivateKey: ecdsaPrivateKey,
	}

	taskManagerAddr := gethcommon.HexToAddress(config.TaskManagerAddress)
	taskResponder, err := taskmanager.NewTaskManagerFromAbi[examplecommon.TaskInput, [32]byte](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
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
