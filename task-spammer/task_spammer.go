package taskspammer

import (
	"context"
	"iter"
	"time"

	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
)

// The task spammer generates tasks every period of time, using the input generator received in the
// Start() method. To send the generated tasks to the TaskManager contract uses the task creator
// received on the NewTaskSpammer function.
type TaskSpammer[Input any] struct {
	// The task creator sends the tasks to the on-chain task manager contract
	taskCreator taskmanager.TaskCreator[Input]
	config      Config
	inputGen    iter.Seq[Input]
}

// Builds a task spammer from a task creator and the received task spammer config. Returns an error if the
// received config is invalid
func NewTaskSpammer[Input any](taskCreator taskmanager.TaskCreator[Input], config Config, inputGen iter.Seq[Input]) (*TaskSpammer[Input], error) {
	// TODO: validate config
	return &TaskSpammer[Input]{
		taskCreator,
		config,
		inputGen,
	}, nil
}

// Runs the Task Spammer in a separate goroutine. This should be called only one time per Task Spammer.
// Returns an error channel, that in case of an error in the run method will contain the received error.
func (taskSpam *TaskSpammer[Input]) Start(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		errChan <- taskSpam.run(ctx)
	}()

	return errChan
}

// The run method contains the main loop of the task spammer, that creates new tasks every period of time, and
// sends them to the on-chain task manager contract. The input sent to the task manager contract on each iteration
// depends on the inputGen received by parameter.
// Note that the taskIndex value is not sent to the task manager contract, so it may differ (will differ if shut
// down and raise another without reseting the anvil node), but it wont affect the workflow of the system.
func (taskSpam *TaskSpammer[Input]) run(ctx context.Context) error {
	logger := taskSpam.config.Logger

	logger.Info("Starting Task Spammer.")

	ticker := time.NewTicker(taskSpam.config.TimeBetweenTasks)
	defer ticker.Stop()
	logger.Infof("Task Spammer set to send new task every %v seconds...", taskSpam.config.TimeBetweenTasks)

	taskIndex := int64(0)

	nextInput, stop := iter.Pull(taskSpam.inputGen)
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
		err := taskSpam.taskCreator.CreateNewTask(ctx, value, taskSpam.config.QuorumThresholdPercentage, taskSpam.config.QuorumNumbers)
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
