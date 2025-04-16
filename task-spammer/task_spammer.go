package taskspammer

import (
	"context"
	"iter"
	"time"
)

// Interface for creating new tasks.
//
// Interface is generic over the task input type.
type TaskCreator[Input any] interface {
	// Creates a new task with the given input.
	CreateNewTask(ctx context.Context, input Input) error
}

type TaskSpammer[Input any] struct {
	taskCreator TaskCreator[Input]
	config      Config
}

func NewTaskSpammer[Input any](taskCreator TaskCreator[Input], config Config) (*TaskSpammer[Input], error) {
	// TODO: validate config
	return &TaskSpammer[Input]{
		taskCreator,
		config,
	}, nil
}

func (taskGen *TaskSpammer[Input]) Start(ctx context.Context, inputGen iter.Seq[Input]) error {
	logger := taskGen.config.Logger

	logger.Info("Starting Task Spammer.")

	ticker := time.NewTicker(taskGen.config.TimeBetweenTasks)
	defer ticker.Stop()
	logger.Infof("Task Spammer set to send new task every %v seconds...", taskGen.config.TimeBetweenTasks)

	taskIndex := int64(0)

	nextInput, stop := iter.Pull(inputGen)
	defer stop()

	for {
		// Submit new task
		taskIndex++
		logger.Infof("Task Spammer sending new task, task index: %v", taskIndex)
		value, ok := nextInput()
		if !ok {
			logger.Info("Task Spammer finished sending tasks")
			return nil
		}
		err := taskGen.taskCreator.CreateNewTask(ctx, value)
		if err != nil {
			logger.Error("Task Spammer failed to send new task", "err", err)
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
