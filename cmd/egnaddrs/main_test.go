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
	// TODO: these are currently hardcoded for the contracts deployed in
	//       test_data/avs-and-eigenlayer-deplyed-anvil-state.json
	serviceManagerAddr      = "0xc3e53F4d16Ae77Db1c982e75a937B9f60FE63690"
	registryCoordinatorAddr = "0x9E545E3C0baAB3E08CdfD552C960A1050f373042"
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
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: filepath.Join(integrationDir, "test_data", "avs-and-eigenlayer-deployed-anvil-state.json"),
				},
				Target: "/root/.anvil/state.json",
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
