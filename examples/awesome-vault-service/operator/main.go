package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/ethereum/go-ethereum/common"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
	"github.com/pelletier/go-toml"
)

type Config struct {
	EthHttpUrl string `toml:"eth_http_url"`
	EthWsUrl   string `toml:"eth_ws_url"`

	OperatorAddress        string `toml:"operator_address"`
	AggregatorServerIPPort string `toml:"aggregator_server_ip_port"`

	BlsKeyPath string `toml:"bls_key_path"`

	RegistryCoordinatorAddress    string `toml:"registry_coordinator_address"`
	OperatorStateRetrieverAddress string `toml:"operator_state_retriever_address"`
	ServiceManagerAddress         string `toml:"service_manager_address"`

	Registration struct {
		RegisterOnStartup        bool   `toml:"register_on_startup"`
		AllocationManagerAddress string `toml:"allocation_manager_address"`
		StrategyAddress          string `toml:"strategy_address"`

		DelegationManagerAddress    string `toml:"delegation_manager_address"`
		RewardsCoordinatorAddress   string `toml:"rewards_coordinator_address"`
		PermissionControllerAddress string `toml:"permission_controller_address"`

		EcdsaKeyPath string `toml:"ecdsa_key_path"`

		AmountToMint          string   `toml:"amount_to_mint"`
		AllocatableMagnitudes []uint64 `toml:"allocatable_magnitudes"`
		OperatorSetIds        []uint32 `toml:"operator_set_ids"`
	} `toml:"registration"`
}

func GetConfigFromPath(path string) (*Config, error) {
	config := &Config{}
	tree, err := toml.LoadFile(path)
	if err != nil {
		return nil, err
	}
	err = tree.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	config, err := GetConfigFromPath("config/operator_config.toml")
	if err != nil {
		logger.Errorf("Failed to read config file: %w", err)
		return
	}

	logger.Infof("config is: %#v", config)

	taskManagerAbi, err := taskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	amount := new(big.Int)
	amount.SetString(config.Registration.AmountToMint, 10)
	registrationConfig := operator.RegistrationConfig{
		RegisterOnStartup: true,

		OperatorAddr:            common.HexToAddress(config.OperatorAddress),
		AllocationManagerAddr:   common.HexToAddress(config.Registration.AllocationManagerAddress),
		AvsAddress:              common.HexToAddress(config.ServiceManagerAddress),
		RegistryCoordinatorAddr: common.HexToAddress(config.RegistryCoordinatorAddress),
		StrategyAddrs:           []common.Address{common.HexToAddress(config.Registration.StrategyAddress)},

		DelegationManagerAddress:    common.HexToAddress(config.Registration.DelegationManagerAddress),
		RewardsCoordinatorAddress:   common.HexToAddress(config.Registration.RewardsCoordinatorAddress),
		PermissionControllerAddress: common.HexToAddress(config.Registration.PermissionControllerAddress),

		EthRpcUrl: config.EthHttpUrl,

		EcdsaKeyStorePath: config.Registration.EcdsaKeyPath,
		BlsKeyStorePath:   config.BlsKeyPath,

		AmountToMint:          amount,
		AllocatableMagnitudes: config.Registration.AllocatableMagnitudes,

		OperatorSetIds: config.Registration.OperatorSetIds,
	}

	operatorConfig := operator.Config{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,

		OperatorAddress: config.OperatorAddress,

		AVSRegistryCoordinatorAddress: config.RegistryCoordinatorAddress,
		OperatorStateRetrieverAddress: config.OperatorStateRetrieverAddress,
		ServiceManagerAddress:         config.ServiceManagerAddress,

		EthWsUrl:                      config.EthWsUrl,
		EthRpcUrl:                     config.EthHttpUrl,
		AggregatorServerIpPortAddress: config.AggregatorServerIPPort,

		BlsPrivateKeyStorePath: config.BlsKeyPath,

		RegistrationCfg: registrationConfig,
	}

	vaultServiceResponseCalc := examplecommon.NewVaultServiceResponseCalculator()

	possibleFailureCalculator, err := operator.NewFailingResponseCalculator(vaultServiceResponseCalc, 35, [32]byte{0})
	if err != nil {
		logger.Fatalf("Failed to create the possible failure function: %v", err.Error())
	}

	// Setting the TaskResponseHashFn parameter in nil because I'm using the abi default encoding function
	operator, err := operator.NewOperatorFromConfig(operatorConfig, possibleFailureCalculator, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	err = operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}
