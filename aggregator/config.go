package aggregator

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	// The url exposed by the anvil node to call the contract methods
	EthHttpUrl                 string `toml:"eth_http_url"`

	// The url exposed by the anvil node to listen to the contract events
	EthWsUrl                   string `toml:"eth_ws_url"`

	// The address where the aggregator will listen for rpc calls from the operators
	AggregatorServerIpPortAddr string `toml:"aggregator_server_ip_port"`

	// These addresses are used to initialize the avs registry chain reader and subscriber
	RegistryCoordinatorAddress    common.Address `toml:"registry_coordinator_address"`
	OperatorStateRetrieverAddress common.Address `toml:"operator_state_retriever_address"`
	ServiceManagerAddress         common.Address `toml:"service_manager_address"`

	// The private key used to create the avs registry chain reader and subscriber
	// TODO: This field can't be parsed from config files, we should add some way to parse it.
	// Maybe we should implement a wrapper using https://pkg.go.dev/encoding#TextUnmarshaler.
	EcdsaPrivateKey *ecdsa.PrivateKey
}
