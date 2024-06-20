package testutils

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	contractreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ContractsRegistry"
)

func StartAnvilContainer(anvilStateFileName string) (testcontainers.Container, error) {

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:nightly-3abac322efdb69e27b6fe8748b72754ae878f64d@sha256:871b66957335636a02c6c324c969db9adb1d6d64f148753c4a986cf32a40dc3c",
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--base-fee", "0", "--gas-price", "0"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	if anvilStateFileName != "" {
		fmt.Println("Starting Anvil container with state file: ", anvilStateFileName)
		req.Cmd = append(req.Cmd, "--load-state", "/root/.anvil/state.json")
		_, curFilePath, _, _ := runtime.Caller(0)
		req.Files = []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(curFilePath, "../../contracts/anvil/", anvilStateFileName),
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

type ContractAddresses struct {
	ServiceManager         common.Address
	RegistryCoordinator    common.Address
	OperatorStateRetriever common.Address
	DelegationManager      common.Address
	Erc20MockStrategy      common.Address
}

func GetContractAddressesFromContractRegistry(ethHttpUrl string) (mockAvsContracts ContractAddresses) {
	ethHttpClient, err := eth.NewClient(ethHttpUrl)
	if err != nil {
		panic(err)
	}
	// The ContractsRegistry contract should always be deployed at this address on anvil
	// it's the address of the contract created at nonce 0 by the first anvil account (0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)
	contractsRegistry, err := contractreg.NewContractContractsRegistry(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), ethHttpClient)
	if err != nil {
		panic(err)
	}

	mockAvsServiceManagerAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "mockAvsServiceManager")
	if err != nil {
		panic(err)
	}
	if mockAvsServiceManagerAddr == (common.Address{}) {
		panic("mockAvsServiceManagerAddr is empty")
	}
	mockAvsRegistryCoordinatorAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "mockAvsRegistryCoordinator")
	if err != nil {
		panic(err)
	}
	if mockAvsRegistryCoordinatorAddr == (common.Address{}) {
		panic("mockAvsRegistryCoordinatorAddr is empty")
	}
	mockAvsOperatorStateRetrieverAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "mockAvsOperatorStateRetriever")
	if err != nil {
		panic(err)
	}
	if mockAvsOperatorStateRetrieverAddr == (common.Address{}) {
		panic("mockAvsOperatorStateRetrieverAddr is empty")
	}
	delegationManagerAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "delegationManager")
	if err != nil {
		panic(err)
	}
	if delegationManagerAddr == (common.Address{}) {
		panic("delegationManagerAddr is empty")
	}
	erc20MockStrategyAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "erc20MockStrategy")
	if err != nil {
		panic(err)
	}
	if erc20MockStrategyAddr == (common.Address{}) {
		panic("erc20MockStrategyAddr is empty")
	}
	mockAvsContracts = ContractAddresses{
		ServiceManager:         mockAvsServiceManagerAddr,
		RegistryCoordinator:    mockAvsRegistryCoordinatorAddr,
		OperatorStateRetriever: mockAvsOperatorStateRetrieverAddr,
		DelegationManager:      delegationManagerAddr,
		Erc20MockStrategy:      erc20MockStrategyAddr,
	}
	return mockAvsContracts
}
