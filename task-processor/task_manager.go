package taskprocessor

import (
	"context"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskManagerContract[Input any, Output any] interface {
	CreateNewTask(ctx context.Context, input Input) error

	RespondToTask(task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature) error
	ProcessTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) (sdktypes.TaskResponseDigest, error)

	RaiseChallenge(task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], TaskResponseMetadata sdktypes.GenericTaskResponseMetadata, NonSigningOperatorPubKeys []sdktypes.BN254G1Point) error
}
