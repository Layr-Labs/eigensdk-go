package main

import (
	"context"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"

	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/incredible-dot-product/contracts/bindings/IncredibleDotProductTaskManager"
)

type DotProductInput struct {
	X []*big.Int
	Y []*big.Int
}

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
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

	err = RegisterOperatorOnStartup(logger)
	if err != nil {
		logger.Fatalf("Failed to register operator on startup: %v", err.Error())
	}

	// Setting the TaskResponseHashFn parameter in nil because I'm using the abi default encoding function
	operator, err := operator.NewOperatorFromConfig(operatorConfig, dotProduct, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	err = operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}

// This function computes the dot product of a pair of points
func dotProduct(taskIndex uint32, points DotProductInput) (*big.Int, error) {
	totalSum := big.NewInt(0)
	for i := range points.X {
		currentSum := big.NewInt(0).Mul(points.X[i], points.Y[i])
		totalSum.Add(totalSum, currentSum)
	}

	return totalSum, nil
}
