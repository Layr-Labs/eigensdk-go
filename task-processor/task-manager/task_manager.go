package taskmanager

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/types"
)

type TaskManagerContract[Input any, Output any] interface {
	TaskCreator[Input]

	TaskResponder[Input, Output]

	ChallengeRaiser[Input, Output]
}

// Interface for creating new tasks.
//
// Interface is generic over the task input type.
type TaskCreator[Input any] interface {
	// Creates a new task with the given input.
	// Implementations usually send a transaction to a smart contract.
	CreateNewTask(ctx context.Context, input Input) error
}

// Interface for responding new tasks.
//
// Interface is generic over the task input and output types.
type TaskResponder[Input any, Output any] interface {
	// Saves the task response for a given task.
	// Implementations usually submits to the on chain task manager contract the response to a task.
	RespondToTask(task types.GenericInputTask[Input], taskResponse types.GenericOutputTaskResponse[Output], nonSignersStakesAndSig types.NonSignerStakesAndSignature) error
	HashTaskResponse(taskResponse types.GenericOutputTaskResponse[Output]) (types.TaskResponseDigest, error)
}

// Interface for raising challenges for responded tasks.
//
// Interface is generic over the task input and output types.
type ChallengeRaiser[Input any, Output any] interface {
	// Raises a challenge for a responded task.
	// Implementations usually sends the challenge to the on chain task manager contract.
	RaiseChallenge(task types.GenericInputTask[Input], taskResponse types.GenericOutputTaskResponse[Output], TaskResponseMetadata types.GenericTaskResponseMetadata, NonSigningOperatorPubKeys []types.BN254G1Point) error
}
