package main

import (
	"context"
	"iter"
	"strconv"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	avtaskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
	gotoml "github.com/pelletier/go-toml"
)

type Config struct {
	EthHttpUrl            string `toml:"eth_http_url"`
	TaskSpammerPrivateKey string `toml:"task_spammer_private_key"`
	TaskManagerAddress    string `toml:"task_manager_address"`
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

	config, err := GetConfigFromPath("config/task_spammer_config.toml")
	if err != nil {
		logger.Errorf("Failed to read config file: %w", err)
		return
	}

	logger.Infof("config is: %#v", config)

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

	taskManagerAddr := common.HexToAddress(config.TaskManagerAddress)

	taskSpammerPrivateKey := config.TaskSpammerPrivateKey
	ecdsaPrivateKey, err := crypto.HexToECDSA(taskSpammerPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create ecdsa private key: %w", err)
		return
	}

	txMgr, err := txmgr.NewSimpleTxManagerFromPrivateKey(logger, ethClient, ecdsaPrivateKey)
	if err != nil {
		logger.Errorf("Failed to create tx manager from private key: %w", err)
		return
	}

	taskCreator, err := taskmanager.NewTaskManagerFromAbi[examplecommon.TaskInput, [32]byte](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Creator: %w", err)
		return
	}

	taskSpammerConfig := taskspammer.Config{
		Logger:                    logger,
		TimeBetweenTasks:          10 * time.Second,
		QuorumThresholdPercentage: 100,
		QuorumNumbers:             []uint8{0},
	}
	taskSpammer, err := taskspammer.NewTaskSpammer(taskCreator, taskSpammerConfig)
	if err != nil {
		logger.Errorf("Failed to create Task Spammer: %w", err)
		return
	}

	seq := LinearRangeSequence()

	err = taskSpammer.Start(context.Background(), seq)
	if err != nil {
		logger.Errorf("Failure while running Task Spammer: %w", err)
		return
	}
}

// Returns an iterator for the sequence 1, 2, 3, ...
// Inspired on the one from examples/incredible-squaring/task-spammer/task_spammer_use_example.go
// Returns [1, 2, ..., n]
func LinearRangeSequence() iter.Seq[examplecommon.TaskInput] {
	keys := []string{"foo", "bar", "baz"}
	n := 1
	return func(yield func(examplecommon.TaskInput) bool) {
		for {
			input := examplecommon.TaskInput{Key: keys[n%len(keys)], Value: strconv.Itoa(n)}
			if !yield(input) {
				break
			}
			n++
		}
	}
}
