package challengerprocessor

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type IndexingChallengerProcessor[Input any, Output any] struct {
	logger               logging.Logger
	responseValidationFn ResponseValidationFunction[Input, Output]
	challengerRaiser     ChallengerRaiser[Input, Output]
	tasks                map[uint32]sdktypes.GenericInputTask[Input]
}

type ResponseValidationFunction[Input any, Output any] func(taskIndex uint32, input Input, output Output) (bool, error)

type ChallengerRaiser[Input any, Output any] interface {
	RaiseChallenge(task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], TaskResponseMetadata sdktypes.GenericTaskResponseMetadata, NonSigningOperatorPubKeys []sdktypes.BN254G1Point) error
}

func NewIndexingChallengerProcessor[Input any, Output any](
	logger logging.Logger,
	responseValidationFn ResponseValidationFunction[Input, Output],
	challengerRaiser ChallengerRaiser[Input, Output],
) (IndexingChallengerProcessor[Input, Output], error) {
	return IndexingChallengerProcessor[Input, Output]{
		logger:               logger,
		responseValidationFn: responseValidationFn,
		challengerRaiser:     challengerRaiser,
		tasks:                make(map[uint32]sdktypes.GenericInputTask[Input]),
	}, nil
}

func (icp IndexingChallengerProcessor[Input, Output]) ProcessNewTaskCreated(newTaskIndex uint32, newTask sdktypes.GenericInputTask[Input]) error {
	icp.tasks[newTaskIndex] = newTask

	return nil
}

func (icp IndexingChallengerProcessor[Input, Output]) ProcessTaskResponded(taskIndex uint32, taskResponse sdktypes.TaskResponseData[Output]) error {
	task, found := icp.tasks[taskIndex]
	if !found {
		return fmt.Errorf("could not find the task for the received task index")
	}

	shouldRaiseChallenge, err := icp.responseValidationFn(taskIndex, task.InputValue, taskResponse.TaskResponse.OutputValue)
	if err != nil {
		icp.logger.Errorf("Failure while validating response. Err: %w", err)
		return err
	}

	if shouldRaiseChallenge {
		icp.logger.Infof("Response was not correct, input was %v and output was %v", task.InputValue, taskResponse.TaskResponse.OutputValue)

		err = icp.challengerRaiser.RaiseChallenge(task, taskResponse.TaskResponse, taskResponse.TaskResponseMetadata, taskResponse.NonSigningOperatorPubKeys)
		if err != nil {
			icp.logger.Errorf("Failure while raising challenge to on-chain contract. Err: %w", err)
			return err
		}
	}

	return nil
}
