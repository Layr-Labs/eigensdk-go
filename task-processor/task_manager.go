package taskprocessor

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskManagerContract[Input any, Output any] interface {
	RespondToTask(task sdktypes.GenericInputTask[Input], taskResponse sdktypes.GenericOutputTaskResponse[Output], nonSignersStakesAndSig sdktypes.NonSignerStakesAndSignature) error
}
