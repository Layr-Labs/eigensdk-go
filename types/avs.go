package types

import (
	"log/slog"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
)

type TaskIndex = uint32
type TaskResponseDigest = Bytes32
type TaskResponse = interface{}

type TaskResponseHashFunction func(taskResponse TaskResponse) TaskResponseDigest

type SignedTaskResponseDigest struct {
	TaskResponse                TaskResponse
	BlsSignature                *bls.Signature
	OperatorId                  OperatorId
	SignatureVerificationErrorC chan error `json:"-"` // removed from json because channels are not marshallable
}

func (strd SignedTaskResponseDigest) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Any("taskResponse", strd.TaskResponse),
		slog.Any("blsSignature", strd.BlsSignature),
		slog.Any("operatorId", strd.OperatorId),
	)
}
