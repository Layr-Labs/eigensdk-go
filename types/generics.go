package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Generic Task, Task Response and Task Response Metadata
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

type TaskResponseMetadata struct {
	TaskRespondedBlock uint32
	HashOfNonSigners   [32]byte
}

type TaskResponseData[Output any] struct {
	TaskResponse              TaskResponse[Output]
	TaskResponseMetadata      TaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

// New Task Created / Task Responded events
type NewTaskCreatedEvent[Input any] struct {
	TaskIndex uint32
	Task      Task[Input]
	Raw       types.Log
}

type TaskRespondedEvent[Output any] struct {
	TaskResponse              TaskResponse[Output]
	TaskResponseMetadata      TaskResponseMetadata
	NonSigningOperatorPubKeys []BN254G1Point
}

type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

type NonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}
