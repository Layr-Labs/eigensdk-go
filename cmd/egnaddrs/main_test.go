package main

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/testutils"
)

const (
	anvilStateFileName = "contracts-deployed-anvil-state.json" // in contracts/anvil/
)

func TestEgnAddrsWithServiceManagerFlag(t *testing.T) {

	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	if err != nil {
		t.Fatal(err)
	}
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	if err != nil {
		t.Error(err)
	}
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilEndpoint)

	args := []string{"egnaddrs"}
	args = append(args, "--service-manager", contractAddrs.ServiceManager.Hex())
	args = append(args, "--rpc-url", anvilEndpoint)
	// we just make sure it doesn't crash
	run(args)
}

func TestEgnAddrsWithRegistryCoordinatorFlag(t *testing.T) {

	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	if err != nil {
		t.Fatal(err)
	}
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	if err != nil {
		t.Error(err)
	}
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilEndpoint)

	args := []string{"egnaddrs"}
	args = append(args, "--registry-coordinator", contractAddrs.RegistryCoordinator.Hex())
	args = append(args, "--rpc-url", anvilEndpoint)
	// we just make sure it doesn't crash
	run(args)
}
