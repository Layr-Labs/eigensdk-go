package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/challenger"
	challengerprocessor "github.com/Layr-Labs/eigensdk-go/challenger/challenger-processor"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	common "github.com/Layr-Labs/eigensdk-go/examples/common"
	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/common"
	idptaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
	gotoml "github.com/pelletier/go-toml"
)

type Config struct {
	EthHttpUrl string `toml:"eth_http_url"`
	EthWsUrl   string `toml:"eth_ws_url"`

	ChallengerPrivateKey string `toml:"challenger_private_key"`

	TaskManagerAddress string `toml:"task_manager_address"`
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

	config, err := GetConfigFromPath("config/challenger_config.toml")
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

	taskManagerAddr := gethcommon.HexToAddress(config.TaskManagerAddress)

	ecdsaPrivateKey, err := crypto.HexToECDSA(config.ChallengerPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create tx manager from private key: %w", err)
		return
	}

	challengerRaiser, err := taskmanager.NewTaskManagerFromAbi[examplecommon.DotProductInput, *big.Int](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create challenger raiser: %w", err)
		return
	}
	dotProductCalculator := operator.NewFunctionResponseCalculator(examplecommon.DotProduct)
	dotProductValidation := challengerprocessor.ResponseValidationFunctionFromResponseCalculator(dotProductCalculator, common.BigIntEqual)
	challengerProcessor, err := challengerprocessor.NewIndexingChallengerProcessor(logger, dotProductValidation, challengerRaiser)
	if err != nil {
		logger.Errorf("Failed to create challenger verifier: %w", err)
		return
	}

	challengerConfig := challenger.Config{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,
		EthClient:      ethClient,
		EthWsUrl:       "ws://localhost:8545",
	}
	challenger, err := challenger.NewChallenger(challengerConfig, challengerProcessor)
	if err != nil {
		logger.Errorf("Failed to create challenger: %w", err)
		return
	}

	err = challenger.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running challenger: %w", err)
		return
	}
}
