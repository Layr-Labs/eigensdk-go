package taskspammer

import (
	"context"
	"iter"
	"math/big"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	taskspammer "github.com/Layr-Labs/eigensdk-go/task-spammer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
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

	taskManagerAddr := common.HexToAddress("0x")

	taskCreator, err := taskspammer.NewTaskCreatorFromAbi[DotProductInput](taskManagerAddr, *taskManagerAbi, &txmgr.SimpleTxManager{}, ethClient)
	if err != nil {
		logger.Errorf("Failed to create Task Creator: %w", err)
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
	}

	seq := NewLinearSequence()

	err = taskSpammer.Start(context.Background(), seq)
	if err != nil {
		logger.Errorf("Failure while running Task Spammer: %w", err)
	}
}

// Returns an iterator for the sequence 1, 2, 3, ...
// Inspired on the one from examples/incredible-squaring/task-spammer/task_spammer_use_example.go
func NewLinearSequence() iter.Seq[DotProductInput] {
	acc := big.NewInt(1)
	delta := big.NewInt(1)
	return func(yield func(DotProductInput) bool) {
		for {
			newInput := DotProductInput{X: []*big.Int{acc}, Y: []*big.Int{acc}}
			if !yield(newInput) {
				break
			}
			acc.Add(acc, delta)
		}
	}
}
