package aggregator

import (
	"github.com/ethereum/go-ethereum/common"
)

// Aggregator configuration struct
type Config struct {
	// Ethereum HTTP RPC URL to use for interacting with on-chain contracts
	EthHttpUrl string `toml:"eth_http_url"`

	// Ethereum WebSocket RPC URL to use for subscribing to on-chain events
	EthWsUrl string `toml:"eth_ws_url"`

	// IP address and port where the aggregator will listen to operator task responses
	AggregatorServerIpPortAddr string `toml:"aggregator_server_ip_port"`

	// These addresses are used to initialize the avs registry chain reader and subscriber
	RegistryCoordinatorAddress    common.Address `toml:"registry_coordinator_address"`
	OperatorStateRetrieverAddress common.Address `toml:"operator_state_retriever_address"`
}
