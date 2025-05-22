package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/ethereum/go-ethereum/common"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

func main() {
	// 1. Create the logger where all the loggs will appear
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 2. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	// 3. Create the registration config, used to register an operator en startup. If you don't want to
	// register your operator, you can leave the RegistrationCfg field of operator config empty
	ethHttpUrl := "http://localhost:8545"

	// 4. Create the config passed to the operator.
	// The values from this config are extracted from an incredible squaring config file:
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	registrationConfig := operator.RegistrationConfig{
		RegisterOnStartup: true,

		AllocationManagerAddr: common.HexToAddress("0x2279b7a0a67db372996a5fab50d91eaa73d2ebe6"),
		AvsAddress:            common.HexToAddress("0xcd8a1c3ba11cf5ecfa6267617243239504a98d90"),
		StrategyAddrs:         []common.Address{common.HexToAddress("0x2b961e3959b79326a8e7f64ef0d2d825707669b5")},

		DelegationManagerAddress:    common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		RewardsCoordinatorAddress:   common.HexToAddress("0xa51c1fc2f0d1a1b8494ed1fe312d7c3a78ed91c0"),
		PermissionControllerAddress: common.HexToAddress("0x59b670e9fa9d0a427751af201d676719a970857b"),

		EcdsaKeyStorePath: "keys/test.ecdsa.key.json",

		AmountToMint:          amount,
		AllocatableMagnitudes: []uint64{1000000000000000},

		OperatorSetIds: []uint32{0},
	}

	operatorConfig := operator.Config{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,

		OperatorAddress: "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",

		RegistryCoordinatorAddress: "0xfd471836031dc5108809d173a067e8486b9047a3",

		EthWsUrl:                      "ws://localhost:8545",
		EthRpcUrl:                     ethHttpUrl,
		AggregatorServerIpPortAddress: "localhost:8090",

		BlsPrivateKeyStorePath: "keys/test.bls.key.json",

		RegistrationCfg: registrationConfig,
	}

	// 5. Create the response calculator with the AVS calculation logic. Note that here we create a Response
	// calculator with the NewFunctionResponseCalculator from the operator package. We also convert the
	// calculator to a failing one, to test that challenges work as expected.
	responseCalculator := operator.NewFunctionResponseCalculator(examplecommon.DotProduct)

	possibleFailureCalculator, err := operator.NewFailingResponseCalculator(responseCalculator, 50, big.NewInt(31234213443))
	if err != nil {
		logger.Fatalf("Failed to create the possible failure function: %v", err.Error())
	}

	// 6. Build the operator, providing operator config, the response calculator and a task response hashing
	// function. Note that we leave this last parameter as nil because we are using the default hashing
	// function provided by the SDK
	operator, err := operator.NewOperatorFromConfig(operatorConfig, possibleFailureCalculator, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	// 7. Run the created operator
	err = operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}
