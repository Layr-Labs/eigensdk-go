package types

import "github.com/Layr-Labs/eigensdk-go/crypto/bls"

type TaskIndex = uint32
type TaskDict struct {
	// We only use a single quorum for incredible squaring
	OperatorsState map[bls.OperatorId]OperatorAvsState
	// aggregate g1 pubkey of all operators in each quorum, at the blocknumber at which the task was started
	AggG1PubkeyPerQuorum map[QuorumNum]*bls.G1Point
}

type TaskResponseDigest [32]byte

type SignedTaskResponseDigest struct {
	TaskResponseDigest TaskResponseDigest
	BlsSignature       *bls.Signature
	OperatorId         bls.OperatorId
}
