package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/types"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

type DotProductOutput struct {
	Result *big.Int
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		panic(err)
	}

	taskManagerAbi, err := taskmanager.ContractIncredibleDotProductTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
	}

	ethHttpUrl := "http://localhost:8545"

	operatorAddr := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	operatorConfig := operator.OperatorConfig{
		Logger:         logger,
		TaskManagerAbi: taskManagerAbi,

		OperatorAddress: operatorAddr,

		AVSRegistryCoordinatorAddress: "0x7bc06c482dead17c0e297afbc32f6e63d3846650",
		OperatorStateRetrieverAddress: "0x4c5859f0f772848b2d91f1d83e2fe57935348029",
		ServiceManagerAddress:         "0x5f3f1dbd7b74c6b46e8c44f98792a1daf8d69154",

		EthWsUrl:                      "ws://localhost:8545",
		EthRpcUrl:                     ethHttpUrl,
		AggregatorServerIpPortAddress: "localhost:8090",

		BlsPrivateKeyStorePath: "keys/test.bls.key.json",
	}

	// This function calculates the task response from a Task, in this case with the number to square
	responseCalcFunction := func(task types.GenericInputTask[DotProductInput], taskIndex uint32) (types.GenericOutputTaskResponse[DotProductOutput], error) {
		totalSum := big.NewInt(0)
		for i := range task.InputValue.X {
			currentSum := big.NewInt(0).Mul(task.InputValue.X[i], task.InputValue.Y[i])
			totalSum.Add(totalSum, currentSum)
		}

		taskResponse := types.GenericOutputTaskResponse[DotProductOutput]{
			ReferenceTaskIndex: taskIndex,
			OutputValue:        DotProductOutput{Result: totalSum},
		}

		return taskResponse, nil
	}

	err = RegisterOperatorOnStartup(logger)
	if err != nil {
		logger.Fatalf("Failed to register operator on startup: %v", err.Error())
	}

	// Setting the TaskResponseHashFn parameter in nil because I'm using the abi default encoding function
	operator, err := operator.NewOperatorFromConfig(operatorConfig, responseCalcFunction, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	err = operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}
