package challenger

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// Challenger configuration struct
type Config struct {
	// Ethereum WebSocket RPC URL to use for subscribing to on-chain events
	EthWsUrl string

	// The client used to communicate with the anvil node
	EthClient *ethclient.Client
}
