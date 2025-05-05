package main

import (
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type DotProductInput struct{
	X []*big.Int 
	Y []*big.Int
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
	responseCalcFunction := func(task types.GenericInputTask[DotProductInput], taskIndex uint32) (types.GenericOutputTaskResponse[*big.Int], error) {
		totalSum := big.NewInt(0)
		for i :=0; i<len(task.InputValue.X); i++  {
			currentSum := big.NewInt(0).Mul(task.InputValue.X[i], task.InputValue.Y[i])
			totalSum.Add(totalSum, currentSum)
		}

		taskResponse := types.GenericOutputTaskResponse[*big.Int]{
			ReferenceTaskIndex: taskIndex,
			OutputValue:        totalSum,
		}

		return taskResponse, nil
	}

	// Setting the TaskResponseHashFn parameter in nil because I'm using the abi default encoding function
	operator.NewOperatorFromConfig(operatorConfig, responseCalcFunction, nil)
}
