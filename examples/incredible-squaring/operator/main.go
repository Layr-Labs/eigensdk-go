package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/ethereum/go-ethereum/common"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/common"
)

// The idea of this example is to show how to create a custom operator using the SDK generic implementation
func main() {
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 3. Create the registration config, used to register an operator en startup. If you don't want to
	// register your operator, you can leave the RegistrationCfg field of operator config empty
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	registrationConfig := operator.RegistrationConfig{
		RegisterOnStartup: true,

		AllocationManagerAddr: common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		AvsAddress:            common.HexToAddress("0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154"),
		StrategyAddrs:         []common.Address{common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5")},

		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		EcdsaKeyStorePath: "keys/test.ecdsa.key.json",

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},
	}

	// 4. Create the config passed to the operator.
	// The values from this config are extracted from an incredible squaring config file:
	// https://github.com/Layr-Labs/incredible-squaring-avs/blob/dev/config-files/operator.anvil.yaml
	operatorConfig := operator.Config{
		OperatorAddress:               "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		RegistryCoordinatorAddress:    "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
		EthRpcUrl:                     "http://localhost:8545",
		EthWsUrl:                      "ws://localhost:8545",
		BlsPrivateKeyStorePath:        "keys/test.bls.key.json",
		AggregatorServerIpPortAddress: "localhost:8090",
		Logger:                        logger,
		TaskManagerAbi:                taskManagerAbi,
		RegistrationCfg:               registrationConfig,
	}

	// 5. Create the response calculator with the AVS calculation logic. Note that here we create a Response
	// calculator with the NewFunctionResponseCalculator from the operator package. We also convert the
	// calculator to a failing one, to test that challenges work as expected.
	calculator := operator.NewFunctionResponseCalculator(examplecommon.Square)

	logic, err := operator.NewFailingResponseCalculator(calculator, 50, big.NewInt(0))
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// 6. Build the operator, providing operator config, the response calculator and a task response hashing
	// function. Note that we leave this last parameter as nil because we are using the default hashing
	// function provided by the SDK
	operator, err := operator.NewOperatorFromConfig(operatorConfig, logic, nil)
	if err != nil {
		logger.Errorf("Failed to create operator from config: %v", err)
		return
	}

	// 10. Run the created operator
	err = operator.Start(context.Background())
	if err != nil {
		logger.Errorf("Error while running operator: %v", err)
		return
	}
}
