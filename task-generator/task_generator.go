package taskgenerator

import (
	"context"
	"iter"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

// Interface for creating new tasks.
//
// Interface is generic over the task input type.
type TaskCreator[Input any] interface {
	// Creates a new task with the given input.
	CreateNewTask(ctx context.Context, input Input) error
}

type TaskGenerator[Input any] struct {
	logger          logging.Logger
	taskSender      TaskCreator[Input]
	secondsInterval int
}

func NewTaskGenerator[Input any](logger logging.Logger, taskSender TaskCreator[Input], secondsInterval int) (*TaskGenerator[Input], error) {
	return &TaskGenerator[Input]{logger, taskSender, secondsInterval}, nil
}

func (taskGen *TaskGenerator[Input]) Start(ctx context.Context, inputGen iter.Seq[Input]) error {
	time.Sleep(time.Duration(2 * time.Second))

	taskGen.logger.Info("Starting Task Generator.")

	ticker := time.NewTicker(time.Duration(taskGen.secondsInterval) * time.Second)
	defer ticker.Stop()
	taskGen.logger.Infof("Task Generator set to send new task every %v seconds...", taskGen.secondsInterval)

	taskIndex := int64(0)

	nextInput, stop := iter.Pull(inputGen)
	defer stop()

	for {
		// Submit new task
		taskIndex++
		taskGen.logger.Infof("Task Generator sending new task, task index: %v", taskIndex)
		value, ok := nextInput()
		if !ok {
			taskGen.logger.Info("Task Generator finished sending tasks")
			return nil
		}
		err := taskGen.taskSender.CreateNewTask(ctx, value)
		if err != nil {
			taskGen.logger.Error("Task Generator failed to send new task", "err", err)
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
