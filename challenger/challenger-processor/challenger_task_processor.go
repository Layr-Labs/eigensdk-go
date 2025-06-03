package challengerprocessor

import (
	"fmt"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// The Indexing Challenger Processor is a generic implementation provided by the SDK
// Note: Do not confuse it with the aggregator module's IndexingProcessor struct
type IndexingChallengerProcessor[Input any, Output any] struct {
	logger logging.Logger

	// This function validates the task response submitted by the operators
	responseValidationFn ResponseValidationFunction[Input, Output]

	// The challenger raiser is responsible for communicating with the `TaskManager` contract to raise challenges
	challengerRaiser taskmanager.ChallengeRaiser[Input, Output]
	tasks            map[uint32]taskmanager.Task[Input]
}

type ResponseValidationFunction[Input any, Output any] func(taskIndex uint32, input Input, output Output) (bool, error)

// Creates an Indexing Challenger Processor from a logger, a response validation function and a challenger raiser.
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

// Processes a new task created event by saving the task in the tasks map
func (icp IndexingChallengerProcessor[Input, Output]) ProcessNewTaskCreated(newTaskIndex uint32, newTask taskmanager.Task[Input]) error {
	icp.tasks[newTaskIndex] = newTask

	return nil
}

// Processes a task responded event, by validating the received task response and raising a
// challenge to the tak manager contract if needed
func (icp IndexingChallengerProcessor[Input, Output]) ProcessTaskResponded(taskIndex uint32, taskResponse taskmanager.TaskResponse[Output], taskResponseMetadata sdktypes.TaskResponseMetadata, nonSigningOperatorPubKeys []sdktypes.BN254G1Point) error {
	task, found := icp.tasks[taskIndex]
	if !found {
		return fmt.Errorf("could not find the task for the received task index")
	}

	isResponseCorrect, err := icp.responseValidationFn(taskIndex, task.InputValue, taskResponse.OutputValue)
	if err != nil {
		icp.logger.Errorf("Failure while validating response. Err: %w", err)
		return err
	}

	if !isResponseCorrect {
		icp.logger.Infof("Response was not correct, input was %v and output was %v", task.InputValue, taskResponse.OutputValue)

		err = icp.challengerRaiser.RaiseChallenge(task, taskResponse, taskResponseMetadata, nonSigningOperatorPubKeys)
		if err != nil {
			icp.logger.Errorf("Failure while raising challenge to on-chain contract. Err: %w", err)
			return err
		}
	}

	delete(icp.tasks, taskIndex)

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
