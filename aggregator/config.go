package aggregator

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	EthHttpUrl                 string `toml:"eth_http_url"`
	EthWsUrl                   string `toml:"eth_ws_url"`
	AggregatorServerIpPortAddr string `toml:"aggregator_server_ip_port"`

	RegistryCoordinatorAddress    common.Address `toml:"registry_coordinator_address"`
	OperatorStateRetrieverAddress common.Address `toml:"operator_state_retriever_address"`
	ServiceManagerAddress         common.Address `toml:"service_manager_address"`

	EcdsaPrivateKey *ecdsa.PrivateKey
}
