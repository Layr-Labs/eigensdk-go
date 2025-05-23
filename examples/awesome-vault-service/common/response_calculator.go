package examplecommon

import (
	"slices"
	"strings"
)

// We decided to create a custom Response Calculator because we have to save state between task responses, so
// the NewFunctionResponseCalculator from the operator package wont be useful.
type VaultServiceResponseCalculator struct {
	vaults []TaskInput
}

func NewVaultServiceResponseCalculator() *VaultServiceResponseCalculator {
	vaults := make([]TaskInput, 0)

	return &VaultServiceResponseCalculator{
		vaults: vaults,
	}
}

func (vsrc *VaultServiceResponseCalculator) ComputeResponse(taskIndex uint32, input TaskInput) ([32]byte, error) {
	cmpFn := func(vault TaskInput, key string) int {
		return strings.Compare(vault.Key, key)
	}
	index, wasFound := slices.BinarySearchFunc(vsrc.vaults, input.Key, cmpFn)
	if wasFound {
		vsrc.vaults[index].Value = input.Value
	} else {
		vsrc.vaults = slices.Insert(vsrc.vaults, index, input)
	}

	return ComputeVaultsRoot(vsrc.vaults), nil
}
