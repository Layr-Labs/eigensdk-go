package aggregator

import (
	"crypto/ecdsa"

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
	ServiceManagerAddress         common.Address `toml:"service_manager_address"`

	// The private key to use when signing transactions
	// TODO: This field can't be parsed from config files, we should add some way to parse it.
	// Maybe we should implement a wrapper using https://pkg.go.dev/encoding#TextUnmarshaler.
	EcdsaPrivateKey *ecdsa.PrivateKey
}
