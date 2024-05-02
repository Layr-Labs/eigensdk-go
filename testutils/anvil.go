package testutils

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func StartAnvilContainer(anvilStateFileName string) (testcontainers.Container, error) {
	integrationDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:latest",
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--base-fee", "0", "--gas-price", "0"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	if anvilStateFileName != "" {
		fmt.Println("Starting Anvil container with state file: ", anvilStateFileName)
		req.Cmd = append(req.Cmd, "--load-state", "/root/.anvil/state.json")
		req.Files = []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(integrationDir, "test_data", anvilStateFileName),
				ContainerFilePath: "/root/.anvil/state.json",
				FileMode:          0644, // Adjust the FileMode according to your requirements
			},
		}
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}
