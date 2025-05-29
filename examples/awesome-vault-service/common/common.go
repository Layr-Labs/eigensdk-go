package examplecommon

import (
	"os"
	"slices"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/pelletier/go-toml/v2"
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

// This function reads the config from the .toml file at the path received as a parameter
// and returns a config with those values
func ReadTomlConfig(path string, config any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = toml.Unmarshal(data, config)
	return err
}
