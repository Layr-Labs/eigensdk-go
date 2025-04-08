package aggregator

import (
	"math/big"

	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
)

type TaskResponse struct {
	ReferenceTaskIndex uint32
	NumberSquared      *big.Int
}

func (tr TaskResponse) TaskIndex() sdktypes.TaskIndex {
	return tr.ReferenceTaskIndex
}

func (tr TaskResponse) Digest() [256]byte {
	return [256]byte(tr.NumberSquared.Bytes())
}
