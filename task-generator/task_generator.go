package taskgenerator

import (
	"context"
	"iter"
	"time"

	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Type is generic over the task input type
type TaskManager[Input any] interface {
	CreateNewTask(opts *bind.TransactOpts, input Input, quorumThresholdPercentage uint32, quorumNumbers []byte) (*gethtypes.Transaction, error)
}

type TaskGenerator[Input any] struct {
	taskManager TaskManager[Input]
	txMgr       txmgr.TxManager
	config      Config
}

func NewTaskGenerator[Input any](taskManager TaskManager[Input], txMgr txmgr.TxManager, config Config) (*TaskGenerator[Input], error) {
	// TODO: validate config
	return &TaskGenerator[Input]{
		taskManager,
		txMgr,
		config,
	}, nil
}

func (taskGen *TaskGenerator[Input]) Start(ctx context.Context, inputGen iter.Seq[Input]) error {
	logger := taskGen.config.Logger

	logger.Info("Starting Task Generator.")

	ticker := time.NewTicker(taskGen.config.TimeBetweenTasks)
	defer ticker.Stop()
	logger.Infof("Task Generator set to send new task every %v seconds...", taskGen.config.TimeBetweenTasks)

	taskIndex := int64(0)

	nextInput, stop := iter.Pull(inputGen)
	defer stop()

	for {
		// Submit new task
		taskIndex++
		logger.Infof("Task Generator sending new task, task index: %v", taskIndex)
		value, ok := nextInput()
		if !ok {
			logger.Info("Task Generator finished sending tasks")
			return nil
		}
		err := taskGen.CreateNewTask(ctx, value)
		if err != nil {
			logger.Error("Task Generator failed to send new task", "err", err)
			return err
		}

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			continue
		}
	}
}

func (taskGen *TaskGenerator[Input]) CreateNewTask(
	ctx context.Context,
	input Input,
) error {
	txOpts, err := taskGen.txMgr.GetNoSendTxOpts()
	if err != nil {
		return utils.WrapError("Error getting tx opts", err)
	}

	tx, err := taskGen.taskManager.CreateNewTask(
		txOpts,
		input,
		taskGen.config.QuorumThresholdPercentage,
		taskGen.config.QuorumNumbers,
	)
	if err != nil {
		return utils.WrapError("Error assembling CreateNewTask tx", err)
	}
	receipt, err := taskGen.txMgr.Send(ctx, tx, true)
	if err != nil {
		return utils.WrapError("Error submitting CreateNewTask tx", err)
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return utils.WrapError("CreateNewTask tx failed", nil)
	}
	return nil
}
