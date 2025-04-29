package operator

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type OperatorConfig[Input any, Output any] struct {
	OperatorAddress string

	// Avs Reader addresses
	OperatorStateRetrieverAddress string
	ServiceManagerAddress         string
	AVSRegistryCoordinatorAddress string

	EthRpcUrl string
	EthWsUrl  string

	BlsPrivateKeyStorePath        string
	AggregatorServerIpPortAddress string

	RegisterOnStartup bool

	Logger         logging.Logger
	TaskManagerAbi *abi.ABI

	ResponseCalculationFn ResponseCalculationFunction[Input, Output]
	TaskResponseHashFn    TaskResponseHashFunction[Output]
}
