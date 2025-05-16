package challenger

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	// The url exposed by the anvil node to listen to the contract events
	EthWsUrl string

	// The logger where the loggs will appear
	Logger logging.Logger

	// The abi of the task manager contract
	TaskManagerAbi *abi.ABI

	// The client used to communicate with the anvil node
	EthClient *ethclient.Client
}
