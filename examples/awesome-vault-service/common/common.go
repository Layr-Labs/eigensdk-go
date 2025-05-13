package examplecommon

import (
	"slices"

	"github.com/ethereum/go-ethereum/crypto"
)

type TaskInput struct {
	Key   string
	Value string
}

func ComputeVaultsRoot(vaults []TaskInput) [32]byte {
	leaves := make([][32]byte, len(vaults))
	for i, vault := range vaults {
		leaves[i] = hashVault(vault)
	}
	for len(leaves) > 1 {
		halfLength := (len(leaves) + 1) / 2
		for i := range halfLength {
			rightIdx := i*2 + 1
			if rightIdx >= len(leaves) {
				rightIdx = i * 2
			}
			leaves[i] = hashNodes(leaves[i*2], leaves[rightIdx])
		}
		leaves = leaves[:halfLength]
	}
	return leaves[0]
}

func hashVault(input TaskInput) [32]byte {
	return crypto.Keccak256Hash([]byte(input.Key + input.Value))
}

func hashNodes(leftNode [32]byte, rightNode [32]byte) [32]byte {
	if slices.Compare(leftNode[:], rightNode[:]) > 0 {
		leftNode, rightNode = rightNode, leftNode
	}
	return crypto.Keccak256Hash(leftNode[:], rightNode[:])
}
