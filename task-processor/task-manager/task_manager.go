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

type TaskResponder[Input any, Output any] interface {
	RespondToTask(task types.GenericInputTask[Input], taskResponse types.GenericOutputTaskResponse[Output], nonSignersStakesAndSig types.NonSignerStakesAndSignature) error
	ProcessTaskResponse(taskResponse types.GenericOutputTaskResponse[Output]) (types.TaskResponseDigest, error)
}

type ChallengeRaiser[Input any, Output any] interface {
	RaiseChallenge(task types.GenericInputTask[Input], taskResponse types.GenericOutputTaskResponse[Output], TaskResponseMetadata types.GenericTaskResponseMetadata, NonSigningOperatorPubKeys []types.BN254G1Point) error
}
