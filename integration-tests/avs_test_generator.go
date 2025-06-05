package integration_test

import (
	"iter"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TestConfig[Input any, Output any] struct {
	TaskManagerAddr common.Address

	TaskManagerAbi *abi.ABI

	// The path to the file with the anvil state
	//AnvilStateFileName string

	// A type that implements a task manager inteface would be quite much overhead,
	// maybe a bound contract like with the task manager wrapper

	// The function to calculate the logic
	LogicFn func(taskIndex uint32, input Input) (Output, error)

	// The function to compare the calculated and the received output in the challenger
	EqualFn func(a, b Output) bool

	// The sequence to generate the inputs sent to the task manager
	InputSequence iter.Seq[Input]

	// Avs Addresses
	RegistryCoordinatorAddress    string
	OperatorStateRetrieverAddress common.Address
	AvsAddress                    common.Address

	BlsKeyStorePath   string
	EcdsaKeyStorePath string

	TaskSpammerPrivateKey string

	// Check if really needed
	AggregatorServerIpPortAddr string
}
