package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Generic Task, Task Response and Task Response Metadata
type GenericInputTask[Input any] struct {
	InputValue                Input
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	QuorumThresholdPercentage uint32
}

type GenericOutputTaskResponse[Output any] struct {
	ReferenceTaskIndex uint32
	OutputValue        Output
}

type GenericTaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

type TaskResponseData[Output any] struct {
	TaskResponse              GenericOutputTaskResponse[Output]
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

// New Task Created / Task Responded events
type NewTaskCreatedEvent[Input any] struct {
	TaskIndex uint32
	Task      GenericInputTask[Input]
	Raw       types.Log
}

type TaskRespondedEvent[Output any] struct {
	TaskResponse              GenericOutputTaskResponse[Output]
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}
