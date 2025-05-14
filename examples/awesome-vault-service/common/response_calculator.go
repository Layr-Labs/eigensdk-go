package examplecommon

import (
	"fmt"
	"slices"
	"strings"
)

type VaultServiceResponseCalculator struct {
	vaults   []TaskInput
	oldValue *string
}

func NewVaultServiceResponseCalculator() *VaultServiceResponseCalculator {
	return &VaultServiceResponseCalculator{
		vaults:   make([]TaskInput, 0),
		oldValue: nil,
	}
}

func (vsrc *VaultServiceResponseCalculator) ComputeResponse(taskIndex uint32, input TaskInput) ([32]byte, error) {
	cmpFn := func(vault TaskInput, key string) int {
		return strings.Compare(vault.Key, key)
	}
	index, wasFound := slices.BinarySearchFunc(vsrc.vaults, input.Key, cmpFn)
	if wasFound {
		vsrc.oldValue = new(string)
		*vsrc.oldValue = vsrc.vaults[index].Value
		vsrc.vaults[index].Value = input.Value
	} else {
		vsrc.oldValue = nil
		vsrc.vaults = slices.Insert(vsrc.vaults, index, input)
	}

	return ComputeVaultsRoot(vsrc.vaults), nil
}

func (vsrc *VaultServiceResponseCalculator) GetPreviousState(lastInput TaskInput) ([]TaskInput, error) {
	oldLeaves := make([]TaskInput, len(vsrc.vaults))
	copy(oldLeaves, vsrc.vaults)

	cmpFn := func(vault TaskInput, key string) int {
		return strings.Compare(vault.Key, key)
	}
	index, wasFound := slices.BinarySearchFunc(vsrc.vaults, lastInput.Key, cmpFn)
	if !wasFound {
		return nil, fmt.Errorf("key %s not found", lastInput.Key)
	}

	if vsrc.oldValue != nil {
		vsrc.vaults[index].Value = *vsrc.oldValue
	} else {
		vsrc.vaults = slices.Delete(vsrc.vaults, index, index+1)
	}

	return oldLeaves, nil
}
