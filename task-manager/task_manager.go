package taskmanager

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/types"
)

type TaskManager[Input any, Output any, Proof any] interface {
	TaskCreator[Input]

	TaskResponder[Input, Output]

	ChallengeRaiser[Input, Output, Proof]
}

// Interface for creating new tasks.
//
// Interface is generic over the task input type.
type TaskCreator[Input any] interface {
	// Creates a new task with the given input.
	// Implementations usually send a transaction to a smart contract.
	CreateNewTask(ctx context.Context, input Input, quorumThresholdPercentage uint32, quorumNumbers []uint8) error
}

// Interface for responding new tasks.
//
// Interface is generic over the task input and output types.
type TaskResponder[Input any, Output any] interface {
	// Saves the task response for a given task.
	// Implementations usually submits to the on chain task manager contract the response to a task.
	RespondToTask(task Task[Input], taskResponse TaskResponse[Output], nonSignersStakesAndSig types.NonSignerStakesAndSignature) error
	// Hashes a generic task response, first encoding it on the task manager abi and then hashing the encoded response.
	// Returns the 32 byte digest or an error in case the hashing process fails
	HashTaskResponse(taskResponse TaskResponse[Output]) (types.TaskResponseDigest, error)
}

// Interface for raising challenges for responded tasks.
//
// Interface is generic over the task input and output types.
type ChallengeRaiser[Input any, Output any, Proof any] interface {
	// Raises a challenge for a responded task.
	// Implementations usually sends the challenge to the on chain task manager contract.
	RaiseChallenge(task Task[Input], taskResponse TaskResponse[Output], taskResponseMetadata types.TaskResponseMetadata, nonSigningOperatorPubKeys []types.BN254G1Point, proof Proof) error
}
