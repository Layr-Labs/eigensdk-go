package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type Config struct {
	operator.Config

	TaskManagerAddress string `toml:"task_manager_address"`
}

// This is the main function for the operator in the incredible dot product example. The steps followed are
// also explained in the operator module readme, which can be found at operator/README.md
func main() {
	// 0. Create the logger where all the logs will appear
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	// 1. Get the ABI of the task manager contract's binding
	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
	}

	// 2. Create the operator config, including the registration config. If you don't want to
	// register your operator, you can leave the RegistrationCfg field of operator config empty

	opConfig := &Config{}
	err = examplecommon.ReadTomlConfig("config/config.toml", opConfig)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	operatorConfig := opConfig.Config

	// 3. Implement the computation function that processes task inputs and produces outputs
	// (we do this in examples/incredible-dot-product/common/common.go)

	// 4. Create the response calculator with the AVS calculation logic. Note that here we create a
	// Response calculator with the NewFunctionResponseCalculator from the operator package
	responseCalculator := operator.NewFunctionResponseCalculator(examplecommon.DotProduct)

	// 5. We convert the calculator to a failing one, to test that challenges work as expected.
	possibleFailureCalculator, err := operator.NewFailingResponseCalculator(responseCalculator, 50, big.NewInt(31234213443))
	if err != nil {
		logger.Fatalf("Failed to create the possible failure function: %v", err.Error())
	}

	// 6. Build the operator, providing operator config, the response calculator and a task response hashing
	// function, and then start it.
	// We leave this last parameter as nil because we are using the default hashing function provided by the SDK
	operator, err := operator.NewOperatorFromConfig(logger, operatorConfig, taskManagerAbi, possibleFailureCalculator, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	err = <-operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}
