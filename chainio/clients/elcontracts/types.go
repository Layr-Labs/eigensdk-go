package elcontracts

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type PendingDeallocation struct {
	MagnitudeDiff        uint64
	CompletableTimestamp uint32
}

type AllocationInfo struct {
	CurrentMagnitude     *big.Int
	PendingDiff          *big.Int
	CompletableTimestamp uint32
	OperatorSetId        uint32
	AvsAddress           common.Address
}
