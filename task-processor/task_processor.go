package taskprocessor

import (
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-processor/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskProcessor[Input any, Output any] interface {
	ProcessNewTask(taskIndex sdktypes.TaskIndex, task taskmanager.Task[Input]) (blsagg.TaskMetadata, error)
	ProcessTaskResponse(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error)
	ProcessAggregatedResponse(response blsagg.BlsAggregationServiceResponse) error
}
