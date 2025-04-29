package taskprocessor

import "github.com/Layr-Labs/eigensdk-go/types"

type TaskManagerContract[Input any, Output any] interface {
	RespondToTask(task types.GenericInputTask[Input], taskResponse types.GenericOutputTaskResponse[Output], nonSignersStakesAndSig types.NonSignerStakesAndSignature) error
}
