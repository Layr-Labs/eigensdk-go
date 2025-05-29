package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/ethereum/go-ethereum/common"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

// TODO: add toml flags to the SDK operator config, removing most of this config attributes
type Config struct {
	OperatorAddress string `toml:"operator_address"`

	// Core deployment addresses
	AllocationManagerAddress    string `toml:"allocation_manager_address"`
	DelegationManagerAddress    string `toml:"delegation_manager_address"`
	RewardsCoordinatorAddress   string `toml:"rewards_coordinator_address"`
	PermissionControllerAddress string `toml:"permission_controller_address"`

	// Avs deployment addresses
	ServiceManagerAddress      string `toml:"service_manager_address"`
	RegistryCoordinatorAddress string `toml:"registry_coordinator_address"`
	TokenStrategyAddr          string `toml:"token_strategy_addr"`

	EcdsaPrivateKeyStorePath string `toml:"ecdsa_private_key_store_path"`
	BlsPrivateKeyStorePath   string `toml:"bls_private_key_store_path"`

	EthRpcUrl string `toml:"eth_http_url"`
	EthWsUrl  string `toml:"eth_ws_url"`

	AggregatorServerIpPortAddress string `toml:"aggregator_server_ip_port"`

	TaskManagerAddress string `toml:"task_manager_address"`
}

// This is the main function for the operator in the awesome vault service example. The steps followed are
// also explained in the operator module readme, which can be found at operator/README.md
func main() {
	// 0. Create the logger where all the logs will appear
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 1. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := taskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	// 2. Create the operator config, including the registration config. If you don't want to
	// register your operator, you can leave the RegistrationCfg field of operator config empty

	opConfig := &Config{}
	err = examplecommon.ReadTomlConfig("config/operator_config.toml", opConfig)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	registrationConfig := operator.RegistrationConfig{
		RegisterOnStartup: true,

		AllocationManagerAddr: common.HexToAddress(opConfig.AllocationManagerAddress),
		AvsAddress:            common.HexToAddress(opConfig.ServiceManagerAddress),
		StrategyAddrs:         []common.Address{common.HexToAddress(opConfig.TokenStrategyAddr)},

		DelegationManagerAddress:    common.HexToAddress(opConfig.DelegationManagerAddress),
		RewardsCoordinatorAddress:   common.HexToAddress(opConfig.RewardsCoordinatorAddress),
		PermissionControllerAddress: common.HexToAddress(opConfig.PermissionControllerAddress),

		EcdsaKeyStorePath: opConfig.EcdsaPrivateKeyStorePath,

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},
	}

	operatorConfig := operator.Config{
		OperatorAddress:               opConfig.OperatorAddress,
		RegistryCoordinatorAddress:    opConfig.RegistryCoordinatorAddress,
		EthRpcUrl:                     opConfig.EthRpcUrl,
		EthWsUrl:                      opConfig.EthWsUrl,
		BlsPrivateKeyStorePath:        opConfig.BlsPrivateKeyStorePath,
		AggregatorServerIpPortAddress: opConfig.AggregatorServerIpPortAddress,
		Logger:                        logger,
		TaskManagerAbi:                taskManagerAbi,
		RegistrationCfg:               registrationConfig,
	}

	// 3. Implement the computation function that processes task inputs and produces outputs
	// (we do this in examples/incredible-squaring/common/common.go)

	// 4. Create the response calculator with the AVS calculation logic. Note that here we create a
	// custom Response calculator, declared on examples/awesome-vault-service/common/response_calculator.go
	vaultServiceResponseCalc := examplecommon.NewVaultServiceResponseCalculator()

	// 5. We convert the calculator to a failing one, to test that challenges work as expected.
	possibleFailureCalculator, err := operator.NewFailingResponseCalculator(vaultServiceResponseCalc, 35, [32]byte{0})
	if err != nil {
		logger.Fatalf("Failed to create the possible failure function: %v", err.Error())
	}

	// 6. Build the operator, providing operator config, the response calculator and a task response hashing
	// function, and then start it.
	// We leave this last parameter as nil because we are using the default hashing function provided by the SDK
	operator, err := operator.NewOperatorFromConfig(operatorConfig, possibleFailureCalculator, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	err = <-operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}
