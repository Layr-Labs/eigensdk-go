package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

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

type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

type TaskResponseData[Output any] struct {
	TaskResponse              GenericOutputTaskResponse[Output]
	TaskResponseMetadata      GenericTaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

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

func (taskRespEvent TaskRespondedEvent[Output]) TaskIndex() uint32 {
	return taskRespEvent.TaskResponse.ReferenceTaskIndex
}

func (taskRespEvent TaskRespondedEvent[Output]) GetTaskResponse() GenericOutputTaskResponse[Output] {
	return taskRespEvent.TaskResponse
}

func (taskRespEvent TaskRespondedEvent[Output]) GetTaskResponseMetadata() GenericTaskResponseMetadata {
	return taskRespEvent.TaskResponseMetadata
}
