package challengerprocessor

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-processor/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type IndexingChallengerProcessor[Input any, Output any] struct {
	logger               logging.Logger
	responseValidationFn ResponseValidationFunction[Input, Output]
	challengerRaiser     taskmanager.ChallengeRaiser[Input, Output]
	tasks                map[uint32]taskmanager.Task[Input]
}

type ResponseValidationFunction[Input any, Output any] func(taskIndex uint32, input Input, output Output) (bool, error)

func NewIndexingChallengerProcessor[Input any, Output any](
	logger logging.Logger,
	responseValidationFn ResponseValidationFunction[Input, Output],
	challengerRaiser taskmanager.ChallengeRaiser[Input, Output],
) (IndexingChallengerProcessor[Input, Output], error) {
	return IndexingChallengerProcessor[Input, Output]{
		logger:               logger,
		responseValidationFn: responseValidationFn,
		challengerRaiser:     challengerRaiser,
		tasks:                make(map[uint32]taskmanager.Task[Input]),
	}, nil
}

func (icp IndexingChallengerProcessor[Input, Output]) ProcessNewTaskCreated(newTaskIndex uint32, newTask taskmanager.Task[Input]) error {
	icp.tasks[newTaskIndex] = newTask

	return nil
}

func (icp IndexingChallengerProcessor[Input, Output]) ProcessTaskResponded(taskIndex uint32, taskResponse sdktypes.TaskResponseData[Output]) error {
	task, found := icp.tasks[taskIndex]
	if !found {
		return fmt.Errorf("could not find the task for the received task index")
	}

	isResponseCorrect, err := icp.responseValidationFn(taskIndex, task.InputValue, taskResponse.TaskResponse.OutputValue)
	if err != nil {
		icp.logger.Errorf("Failure while validating response. Err: %w", err)
		return err
	}

	if !isResponseCorrect {
		icp.logger.Infof("Response was not correct, input was %v and output was %v", task.InputValue, taskResponse.TaskResponse.OutputValue)

		err = icp.challengerRaiser.RaiseChallenge(task, taskResponse.TaskResponse, taskResponse.TaskResponseMetadata, taskResponse.NonSigningOperatorPubKeys)
		if err != nil {
			icp.logger.Errorf("Failure while raising challenge to on-chain contract. Err: %w", err)
			return err
		}
	}

	return nil
}

// Takes a ResponseCalculator and an Equal function.
// Returns a function that receives a task input and output, calculates the expected
// output using the ResponseCalculator and returns whether it is equal to the given output,
// using the given Equal function.
func ResponseValidationFunctionFromResponseCalculator[Input any, Output any](
	responseCalculator operator.ResponseCalculator[Input, Output],
	equalFn func(a, b Output) bool,
) ResponseValidationFunction[Input, Output] {
	return func(taskIndex uint32, input Input, output Output) (bool, error) {
		computedResponse, err := responseCalculator.ComputeResponse(taskIndex, input)
		if err != nil {
			return false, utils.WrapError("failed to compute response", err)
		}
		return equalFn(computedResponse, output), nil
	}
}
