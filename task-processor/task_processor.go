package taskprocessor

import (
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskProcessor[Input any, Output any] interface {
	ProcessNewTask(taskIndex sdktypes.TaskIndex, task sdktypes.GenericInputTask[Input]) (blsagg.TaskMetadata, error)
	ProcessTaskResponse(taskResponse sdktypes.GenericOutputTaskResponse[Output]) ([32]byte, error)
	ProcessAggregatedResponse(response blsagg.BlsAggregationServiceResponse) error
}
