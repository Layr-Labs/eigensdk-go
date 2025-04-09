package taskgenerator

import (
	"context"
	"time"

	"github.com/Layr-Labs/eigensdk-go/logging"
)

type TaskGenerator struct {
	logger          logging.Logger
	logic			TaskGeneratorLogic
	secondsBetweenTasks time.Duration
}

func BuildTaskGenerator(logger logging.Logger, logic TaskGeneratorLogic, timeBetweenTasks time.Duration) (*TaskGenerator, error) {
	return &TaskGenerator{
		logger,
		logic,
		timeBetweenTasks,
	}, nil
}

func (taskGen *TaskGenerator) Start(ctx context.Context) error {
	time.Sleep(time.Duration(2 * time.Second))

	taskGen.logger.Info("Starting Task Generator.")
	taskGen.logger.Info("Starting Task Generator rpc server.")

	ticker := time.NewTicker(taskGen.secondsBetweenTasks * time.Second)
	defer ticker.Stop()
	taskGen.logger.Info("Task Generator set to send new task every %v seconds...", taskGen.secondsBetweenTasks)

	taskNum := int64(0)

	// Send a task before looping
	err := taskGen.logic.sendNewTask(taskNum)
	if err != nil {
		taskGen.logger.Error("Aggregator failed to send number to square", "err", err)
		return err
	}
	taskNum++

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			taskGen.logger.Infof("Task Generator sending new task, number to square: %v", taskNum)
			err := taskGen.logic.sendNewTask(taskNum)
			if err != nil {
				taskGen.logger.Error("Aggregator failed to send number to square", "err", err)
				return err
			}
			taskNum++
		}
	}
}

type TaskGeneratorLogic interface{
	sendNewTask(taskNumber int64)(error)
}
