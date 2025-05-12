package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	sdkoperator "github.com/Layr-Labs/eigensdk-go/operator"

	cstaskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-squaring/bindings/taskManager"
)

// The idea of this example is to show how to create a custom operator using the SDK generic implementation
func main() {
	logger, err := logging.NewZapLogger(logging.Production) // Change here if want to change logging level
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := cstaskmanager.ContractIncredibleSquaringTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	// The values from this config are extracted from an incredible squaring config file:
	// https://github.com/Layr-Labs/incredible-squaring-avs/blob/dev/config-files/operator.anvil.yaml
	operatorConfig := sdkoperator.OperatorConfig{
		OperatorAddress:               "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		OperatorStateRetrieverAddress: "0x4c5859f0f772848b2d91f1d83e2fe57935348029",
		ServiceManagerAddress:         "0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154",
		AVSRegistryCoordinatorAddress: "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
		EthRpcUrl:                     "http://localhost:8545",
		EthWsUrl:                      "ws://localhost:8545",
		BlsPrivateKeyStorePath:        "tests/keys/test.bls.key.json",
		AggregatorServerIpPortAddress: "localhost:8090",
		Logger:                        logger,
		TaskManagerAbi:                taskManagerAbi,
	}

	calculator := sdkoperator.NewFunctionResponseCalculator(square)

	logic, err := sdkoperator.NewFailingResponseCalculator(calculator, 50, big.NewInt(0))
	if err != nil {
		logger.Fatalf(err.Error())
	}

	operator, err := sdkoperator.NewOperatorFromConfig(operatorConfig, logic, nil)
	if err != nil {
		logger.Errorf("Failed to create operator from config: %v", err)
		return
	}

	err = operator.Start(context.Background())
	if err != nil {
		logger.Errorf("Error while running operator: %v", err)
		return
	}
}

// This function computes the square of a number
func square(taskIndex uint32, numberToSquare *big.Int) (*big.Int, error) {
	numberSquared := big.NewInt(0).Exp(numberToSquare, big.NewInt(2), nil)

	return numberSquared, nil
}
