package main

import (
	"context"
	"slices"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/operator"
	"github.com/ethereum/go-ethereum/crypto"

	examplecommon "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/common"
	taskmanager "github.com/Layr-Labs/eigensdk-go/examples/awesome-vault-service/contracts/bindings/AwesomeVaultTaskManager"
)

func main() {
	logger, err := logging.NewZapLogger(logging.Production)
	if err != nil {
		println("Failure creating logger")
		return
	}

	taskManagerAbi, err := taskmanager.ContractAwesomeVaultTaskManagerMetaData.GetAbi()
	if err != nil {
		logger.Errorf("Failed to get task manager abi: %w", err)
		return
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

	cmpFn := func(vault examplecommon.TaskInput, key string) int {
		return strings.Compare(vault.Key, key)
	}

	vaults := make([]examplecommon.TaskInput, 0)

	computeFn := func(taskIndex uint32, input examplecommon.TaskInput) ([32]byte, error) {
		index, wasFound := slices.BinarySearchFunc(vaults, input.Key, cmpFn)
		if wasFound {
			vaults[index].Value = input.Value
		} else {
			vaults = slices.Insert(vaults, index, input)
		}
		return computeVaultsRoot(vaults), nil
	}

	possibleFailureFunction, err := operator.ComputeWithFailures(computeFn, failingVaultSet, 50)
	if err != nil {
		logger.Fatalf("Failed to create the possible failure function: %v", err.Error())
	}

	// Setting the TaskResponseHashFn parameter in nil because I'm using the abi default encoding function
	operator, err := operator.NewOperatorFromConfig(operatorConfig, possibleFailureFunction, nil)
	if err != nil {
		logger.Fatalf("Failed to create operator: %w", err)
	}

	err = operator.Start(context.Background())
	if err != nil {
		logger.Fatalf("Failure while running operator: %w", err)
	}
}

func failingVaultSet(taskIndex uint32, input examplecommon.TaskInput) ([32]byte, error) {
	return [32]byte{0}, nil
}

func computeVaultsRoot(vaults []examplecommon.TaskInput) [32]byte {
	leaves := make([][32]byte, len(vaults))
	for i, vault := range vaults {
		leaves[i] = hashVault(vault)
	}
	for len(leaves) > 1 {
		halfLength := (len(leaves) + 1) / 2
		for i := range halfLength {
			rightIdx := i*2 + 1
			if rightIdx >= len(leaves) {
				rightIdx = i * 2
			}
			leaves[i] = hashNodes(leaves[i*2], leaves[rightIdx])
		}
		leaves = leaves[:halfLength]
	}
	return leaves[0]
}

func hashVault(input examplecommon.TaskInput) [32]byte {
	return crypto.Keccak256Hash([]byte(input.Key + input.Value))
}

func hashNodes(leftNode [32]byte, rightNode [32]byte) [32]byte {
	if slices.Compare(leftNode[:], rightNode[:]) > 0 {
		leftNode, rightNode = rightNode, leftNode
	}
	return crypto.Keccak256Hash(leftNode[:], rightNode[:])
}
