package taskprocessor

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskManagerContract[Input any, Output any] interface {
	RespondToTask(task sdktypes.Task[Input], taskResponse sdktypes.TaskResponse[Output], nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature) error
}
