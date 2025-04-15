package testutils

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	contractreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ContractsRegistry"
)

const (
	ANVIL_FIRST_ADDRESS      = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	ANVIL_FIRST_PRIVATE_KEY  = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ANVIL_SECOND_ADDRESS     = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
	ANVIL_SECOND_PRIVATE_KEY = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	ANVIL_THIRD_ADDRESS      = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
	ANVIL_THIRD_PRIVATE_KEY  = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
)

func StartAnvilContainer(anvilStateFileName string) (testcontainers.Container, error) {

	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:stable",
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--base-fee", "0", "--gas-price", "0"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	if anvilStateFileName != "" {
		fmt.Println("Starting Anvil container with state file: ", anvilStateFileName)
		req.Cmd = append(req.Cmd, "--load-state", "/state.json")
		_, curFilePath, _, _ := runtime.Caller(0)
		req.Files = []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(curFilePath, "../../contracts/anvil/", anvilStateFileName),
				ContainerFilePath: "/state.json",
				FileMode:          0644, // Adjust the FileMode according to your requirements
			},
		}
	}
	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	// Still need to advance the chain by at least 1 block b/c some tests need to query the latest block,
	// and the blocks dumped/loaded by anvil don't contain full transactions, which leads to panics in tests.
	// See https://github.com/foundry-rs/foundry/issues/8213, which will hopefully get fixed soon.
	if anvilStateFileName != "" {
		AdvanceChainByNBlocksExecInContainer(ctx, 1, anvilC)
	}

	return anvilC, nil
}

// This function is for M2 functionality
func StartM2AnvilContainer(anvilStateFileName string) (testcontainers.Container, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "ghcr.io/foundry-rs/foundry:stable",
		Entrypoint:   []string{"anvil"},
		Cmd:          []string{"--host", "0.0.0.0", "--base-fee", "0", "--gas-price", "0"},
		ExposedPorts: []string{"8545/tcp"},
		WaitingFor:   wait.ForLog("Listening on"),
	}
	if anvilStateFileName != "" {
		fmt.Println("Starting Anvil container with state file: ", anvilStateFileName)
		req.Cmd = append(req.Cmd, "--load-state", "/state.json")
		_, curFilePath, _, _ := runtime.Caller(0)
		req.Files = []testcontainers.ContainerFile{
			{
				HostFilePath:      filepath.Join(curFilePath, "../../M2-contracts/anvil/", anvilStateFileName),
				ContainerFilePath: "/state.json",
				FileMode:          0644, // Adjust the FileMode according to your requirements
			},
		}
	}
	anvilC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	// Still need to advance the chain by at least 1 block b/c some tests need to query the latest block,
	// and the blocks dumped/loaded by anvil don't contain full transactions, which leads to panics in tests.
	// See https://github.com/foundry-rs/foundry/issues/8213, which will hopefully get fixed soon.
	if anvilStateFileName != "" {
		AdvanceChainByNBlocksExecInContainer(ctx, 1, anvilC)
	}

	return anvilC, nil
}

type ContractAddresses struct {
	ServiceManager         common.Address
	RegistryCoordinator    common.Address
	OperatorStateRetriever common.Address
	DelegationManager      common.Address
	Erc20MockStrategy      common.Address
	RewardsCoordinator     common.Address
	PermissionController   common.Address
}

func GetContractAddressesFromContractRegistry(ethHttpUrl string) (mockAvsContracts ContractAddresses) {
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		panic(err)
	}
	// The ContractsRegistry contract should always be deployed at this address on anvil
	// it's the address of the contract created at nonce 0 by the first anvil account
	// (0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)
	contractsRegistry, err := contractreg.NewContractContractsRegistry(
		common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		ethHttpClient,
	)
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
	mockAvsOperatorStateRetrieverAddr, err := contractsRegistry.Contracts(
		&bind.CallOpts{},
		"mockAvsOperatorStateRetriever",
	)
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
	rewardsCoordinatorAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "rewardsCoordinator")
	if err != nil {
		panic(err)
	}
	if rewardsCoordinatorAddr == (common.Address{}) {
		panic("rewardsCoordinatorAddr is empty")
	}
	permissionControllerAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "permissionController")
	if err != nil {
		panic(err)
	}
	if permissionControllerAddr == (common.Address{}) {
		panic("permissionControllerAddr is empty")
	}
	mockAvsContracts = ContractAddresses{
		ServiceManager:         mockAvsServiceManagerAddr,
		RegistryCoordinator:    mockAvsRegistryCoordinatorAddr,
		OperatorStateRetriever: mockAvsOperatorStateRetrieverAddr,
		DelegationManager:      delegationManagerAddr,
		Erc20MockStrategy:      erc20MockStrategyAddr,
		RewardsCoordinator:     rewardsCoordinatorAddr,
		PermissionController:   permissionControllerAddr,
	}
	return mockAvsContracts
}

// This function is for M2 functionality
func GetM2ContractAddressesFromContractRegistry(ethHttpUrl string) (mockAvsContracts ContractAddresses) {
	ethHttpClient, err := ethclient.Dial(ethHttpUrl)
	if err != nil {
		panic(err)
	}
	// The ContractsRegistry contract should always be deployed at this address on anvil
	// it's the address of the contract created at nonce 0 by the first anvil account
	// (0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266)
	contractsRegistry, err := contractreg.NewContractContractsRegistry(
		common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"),
		ethHttpClient,
	)
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
	mockAvsOperatorStateRetrieverAddr, err := contractsRegistry.Contracts(
		&bind.CallOpts{},
		"mockAvsOperatorStateRetriever",
	)
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
	rewardsCoordinatorAddr, err := contractsRegistry.Contracts(&bind.CallOpts{}, "rewardsCoordinator")
	if err != nil {
		panic(err)
	}
	if rewardsCoordinatorAddr == (common.Address{}) {
		panic("rewardsCoordinatorAddr is empty")
	}
	mockAvsContracts = ContractAddresses{
		ServiceManager:         mockAvsServiceManagerAddr,
		RegistryCoordinator:    mockAvsRegistryCoordinatorAddr,
		OperatorStateRetriever: mockAvsOperatorStateRetrieverAddr,
		DelegationManager:      delegationManagerAddr,
		Erc20MockStrategy:      erc20MockStrategyAddr,
		RewardsCoordinator:     rewardsCoordinatorAddr,
	}
	return mockAvsContracts
}

func AdvanceChainByNBlocks(n int, anvilEndpoint string) {
	// see https://book.getfoundry.sh/reference/anvil/#custom-methods
	cmd := exec.Command("cast", "rpc", "anvil_mine", fmt.Sprintf("%d", n), "--rpc-url", anvilEndpoint)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// Prefer this function over AdvanceChainByNBlocks b/c it doesn't require cast to be installed on the host machine,
// whereas this one doesn't.
func AdvanceChainByNBlocksExecInContainer(ctx context.Context, n int, anvilC testcontainers.Container) {
	c, _, err := anvilC.Exec(
		ctx,
		[]string{"cast", "rpc", "anvil_mine", fmt.Sprintf("%d", n), "--rpc-url", "http://localhost:8545"},
	)
	if err != nil {
		panic(err)
	}
	if c != 0 {
		log.Fatalf("Unable to advance anvil chain by n blocks. Expected return code 0, got %v", c)
	}
}

type TestConfig struct {
	AnvilStateFileName string
	LogLevel           slog.Level
}

func GetDefaultTestConfig() TestConfig {
	return TestConfig{
		AnvilStateFileName: "contracts-deployed-anvil-state.json",
		LogLevel:           slog.LevelDebug,
	}
}
