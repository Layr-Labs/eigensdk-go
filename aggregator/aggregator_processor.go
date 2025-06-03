package aggregator

import (
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

// The processor processes new tasks, task responses, and aggregated responses.
// A default implementation for this trait can be the Indexing Processor, that can be found in
// aggregator/indexing_processor.go
type Processor[Input any, Output any] interface {
	// Processes a new task, and returns the task metadata to be sent to the BLS aggregation service, or an error
	ProcessNewTask(taskIndex sdktypes.TaskIndex, task taskmanager.Task[Input]) (blsagg.TaskMetadata, error)

	// Processes a task response, and returns the hashed bytes of the response, or an error
	ProcessTaskResponse(taskResponse taskmanager.TaskResponse[Output]) ([32]byte, error)

	// Processes an aggregated response, returning an error if the process fails
	ProcessAggregatedResponse(taskIndex sdktypes.TaskIndex, taskResponse taskmanager.TaskResponse[Output], nonSignerStakesAndSignature sdktypes.NonSignerStakesAndSignature) error
}
