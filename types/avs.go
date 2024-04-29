package types

import (
	"log/slog"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

type TaskIndex uint32
type TaskResponseDigest Bytes32

type SignedTaskResponseDigest struct {
	TaskResponseDigest          TaskResponseDigest
	BlsSignature                *bls.Signature
	OperatorId                  OperatorId
	SignatureVerificationErrorC chan error `json:"-"` // removed from json because channels are not marshallable
}

func (strd SignedTaskResponseDigest) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("taskResponseDigest", strd.TaskResponseDigest),
		slog.Any("blsSignature", strd.BlsSignature),
		slog.Any("operatorId", strd.OperatorId),
	)
}
