package challenger

import (
	taskmanager "github.com/Layr-Labs/eigensdk-go/task-manager"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

// The Challenger processor is responsible for processing the challenges
// Note: Do not confuse it with the aggregator module's Processor interface
type ChallengerProcessor[Input any, Output any] interface {
	// Processes new tasks, returns an error in case of failure
	ProcessNewTaskCreated(taskIndex uint32, task taskmanager.Task[Input]) error
	// Processes task responses, returns an error in case of failure
	ProcessTaskResponded(taskIndex uint32, taskResponse taskmanager.TaskResponse[Output], taskResponseMetadata sdktypes.TaskResponseMetadata, nonSigningOperatorPubKeys []sdktypes.BN254G1Point) error
}
