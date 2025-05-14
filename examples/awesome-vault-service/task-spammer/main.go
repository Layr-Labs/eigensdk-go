package main

import (
	"context"
	"iter"
	"strconv"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-processor/task-manager"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	avtaskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := avtaskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
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

	taskManagerAddr := common.HexToAddress("0x7bc06c482dead17c0e297afbc32f6e63d3846650")

	taskSpammerPrivateKey := "4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
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

	taskCreator, err := taskmanager.NewTaskManagerContractFromAbi[examplecommon.TaskInput, [32]byte](taskManagerAddr, taskManagerAbi, txMgr, ethClient)
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
