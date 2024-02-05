package types

import (
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/ethereum/go-ethereum/common"
)

type TaskIndex = uint32
type TaskResponseDigest [32]byte

type SignedBlsTaskResponseDigest struct {
	TaskResponseDigest          TaskResponseDigest
	BlsSignature                *bls.Signature
	OperatorId                  bls.OperatorId
	SignatureVerificationErrorC chan error
}

type SignedEcdsaTaskResponseDigest struct {
	TaskResponseDigest          TaskResponseDigest
	EcdsaSignature              []byte
	OperatorId                  common.Address
	SignatureVerificationErrorC chan error
}
