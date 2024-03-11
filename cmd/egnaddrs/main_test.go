package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	// just copied this file from an eigencert deployment and hardcoded addrs below
	// TODO(samlaf): eventually we should make updating this file automated
	anvilStateFileName      = "eigenlayer-and-registries-deployed-anvil-state.json"
	serviceManagerAddr      = "0x7a2088a1bFc9d81c55368AE168C2C02570cB814F"
	registryCoordinatorAddr = "0x09635F643e140090A9A8Dcd712eD6285858ceBef"
)

func TestEgnAddrsWithServiceManagerFlag(t *testing.T) {

	anvilC := startAnvilTestContainer()
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

	anvilC := startAnvilTestContainer()
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

func startAnvilTestContainer() testcontainers.Container {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image: "ghcr.io/foundry-rs/foundry:latest",
		Files: []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(integrationDir, "test_data", anvilStateFileName),
				ContainerFilePath: "/root/.anvil/state.json",
				FileMode:          0644, // Adjust the FileMode according to your requirements
			},
		},
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--load-state", "/root/.anvil/state.json"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	return anvilC
}
