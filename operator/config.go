package operator

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type OperatorConfig struct {
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

	// Percentage chance of randomly failing a task
	// This is used for testing purposes and defaults to not failing any tasks
	FailingPercentage uint

	Logger         logging.Logger
	TaskManagerAbi *abi.ABI
}
