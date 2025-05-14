package taskmanager

import (
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/core/types"
)

type Task[Input any] struct {
	InputValue                Input
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

type TaskResponse[Output any] struct {
	ReferenceTaskIndex uint32
	OutputValue        Output
}

type NewTaskCreatedEvent[Input any] struct {
	TaskIndex uint32
	Task      Task[Input]
	Raw       types.Log
}

type TaskRespondedEvent[Output any] struct {
	TaskResponse              TaskResponse[Output]
	TaskResponseMetadata      sdktypes.TaskResponseMetadata
	NonSigningOperatorPubKeys []sdktypes.BN254G1Point
}
