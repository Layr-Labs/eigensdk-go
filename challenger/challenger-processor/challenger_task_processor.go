package challengerprocessor

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

type IndexingChallengerProcessor[Input any, Output any, Proof any] struct {
	logger               logging.Logger
	responseValidationFn ResponseValidationFunction[Input, Output, Proof]
	challengerRaiser     taskmanager.ChallengeRaiser[Input, Output, Proof]
	tasks                map[uint32]taskmanager.Task[Input]
}

type ResponseValidationFunction[Input any, Output any, Proof any] func(taskIndex uint32, input Input, output Output) (bool, Proof, error)

func NewIndexingChallengerProcessor[Input any, Output any, Proof any](
	logger logging.Logger,
	responseValidationFn ResponseValidationFunction[Input, Output, Proof],
	challengerRaiser taskmanager.ChallengeRaiser[Input, Output, Proof],
) (IndexingChallengerProcessor[Input, Output, Proof], error) {
	return IndexingChallengerProcessor[Input, Output, Proof]{
		logger:               logger,
		responseValidationFn: responseValidationFn,
		challengerRaiser:     challengerRaiser,
		tasks:                make(map[uint32]taskmanager.Task[Input]),
	}, nil
}

func (icp IndexingChallengerProcessor[Input, Output, Proof]) ProcessNewTaskCreated(newTaskIndex uint32, newTask taskmanager.Task[Input]) error {
	icp.tasks[newTaskIndex] = newTask

	return nil
}

func (icp IndexingChallengerProcessor[Input, Output, Proof]) ProcessTaskResponded(taskIndex uint32, taskResponse taskmanager.TaskResponse[Output], taskResponseMetadata sdktypes.TaskResponseMetadata, nonSigningOperatorPubKeys []sdktypes.BN254G1Point) error {
	task, found := icp.tasks[taskIndex]
	if !found {
		return fmt.Errorf("could not find the task for the received task index")
	}

	isResponseCorrect, proof, err := icp.responseValidationFn(taskIndex, task.InputValue, taskResponse.OutputValue)
	if err != nil {
		icp.logger.Errorf("Failure while validating response. Err: %w", err)
		return err
	}

	if !isResponseCorrect {
		icp.logger.Infof("Response was not correct, input was %v and output was %v", task.InputValue, taskResponse.OutputValue)

		err = icp.challengerRaiser.RaiseChallenge(task, taskResponse, taskResponseMetadata, nonSigningOperatorPubKeys, proof)
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
) ResponseValidationFunction[Input, Output, any] {
	return func(taskIndex uint32, input Input, output Output) (bool, any, error) {
		computedResponse, err := responseCalculator.ComputeResponse(taskIndex, input)
		if err != nil {
			return false, nil, utils.WrapError("failed to compute response", err)
		}
		return equalFn(computedResponse, output), nil, nil
	}
}
