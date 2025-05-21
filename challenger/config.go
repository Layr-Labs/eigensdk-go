package challenger

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Challenger configuration struct
type Config struct {
	// Ethereum WebSocket RPC URL to use for subscribing to on-chain events
	EthWsUrl string

	// The logger where the loggs will appear
	Logger logging.Logger

	// The ABI of the task manager contract
	TaskManagerAbi *abi.ABI

	// The client used to communicate with the anvil node
	EthClient *ethclient.Client
}
