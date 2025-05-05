package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

type DotProductOutput struct {
	Result *big.Int
}

func main() {
	logger, err := logging.NewZapLogger(logging.Development)
	if err != nil {
		panic(err)
	}

	taskManagerAbi := abi.ABI{}

	operatorConfig := operator.OperatorConfig{
		Logger:         logger,
		TaskManagerAbi: &taskManagerAbi,

		RegisterOnStartup: true,
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

	// Setting the TaskResponseHashFn parameter in nil because I'm using the abi default encoding function
	operator, err := operator.NewOperatorFromConfig(operatorConfig, responseCalcFunction, nil)
	if err != nil {
		logger.Errorf("Failed to create operator: %w", err)
	}

	err = operator.Start(context.Background())
	if err != nil {
		logger.Errorf("Failure while running operator: %w", err)
	}
}
