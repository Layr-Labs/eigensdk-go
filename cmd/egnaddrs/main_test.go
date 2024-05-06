package main

import (
	"context"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/testutils"
)

const (
	// just copied this file from an eigencert deployment and hardcoded addrs below
	// TODO(samlaf): eventually we should make updating this file automated
	anvilStateFileName      = "eigenlayer-and-registries-deployed-anvil-state.json"
	serviceManagerAddr      = "0x7a2088a1bFc9d81c55368AE168C2C02570cB814F"
	registryCoordinatorAddr = "0x09635F643e140090A9A8Dcd712eD6285858ceBef"
)

func TestEgnAddrsWithServiceManagerFlag(t *testing.T) {

	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	if err != nil {
		t.Fatal(err)
	}
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	args := []string{"egnaddrs"}
	args = append(args, "--service-manager", serviceManagerAddr)
	args = append(args, "--rpc-url", "http://"+anvilEndpoint)
	// we just make sure it doesn't crash
	run(args)
}

func TestEgnAddrsWithRegistryCoordinatorFlag(t *testing.T) {

	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	if err != nil {
		t.Fatal(err)
	}
	anvilEndpoint, err := anvilC.Endpoint(context.Background(), "")
	if err != nil {
		t.Error(err)
	}

	args := []string{"egnaddrs"}
	args = append(args, "--registry-coordinator", registryCoordinatorAddr)
	args = append(args, "--rpc-url", "http://"+anvilEndpoint)
	// we just make sure it doesn't crash
	run(args)
}
