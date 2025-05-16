package taskprocessor

import (
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

// The task processor is the responible for processing new tasks, task responses and aggregated responses.
// A default implementation for this trait can be the Indexing Task Processor, that can be found in
// aggregator/task-processor/indexing_task_processor.go
type TaskProcessor[Input any, Output any] interface {
	// Processes a new task, and returns the task metadata to be sent to the BLS aggregation service
	ProcessNewTask(taskIndex sdktypes.TaskIndex, task taskmanager.Task[Input]) (blsagg.TaskMetadata, error)

	// Processes a task response, and returns the hashed bytes of the response
	ProcessTaskResponse(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error)

	// Processes an aggregated response, returning an error if the process fails
	ProcessAggregatedResponse(response blsagg.BlsAggregationServiceResponse) error
}
