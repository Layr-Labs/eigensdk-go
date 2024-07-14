//go:build wireinject
// +build wireinject

package avsregistry

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/google/wire"
)

func NewReaderFromConfig(
	cfg Config,
	ethClient eth.Client,
	logger logging.Logger,
) (*ChainReader, error) {
	wire.Build(NewBindingsFromConfig, NewChainReaderFromBindings)
	return &ChainReader{}, nil
}

// including here because needed in NewWriterFromConfig
func NewElConfigFromAvsBindings(
	bindings *ContractBindings,
) elcontracts.Config {
	return elcontracts.Config{
		// TODO: missing RewardsCoordinatorAddr in bindings, because no other contract points to
		// RewardsCoordinator so bindings doesn't return that address...
		// Ways to fix this bug:
		// 1. Make some other contract point to RewardsCoordinator so we can get the address from bindings
		// 2. Pass RewardsCoordinatorAddr as a parameter to NewAVSRegistryContractBindings
		// 3. Pass RewardsCoordinatorAddr as a parameter to this function and NewWriterFromConfig
		DelegationManagerAddress: bindings.DelegationManagerAddr,
		AvsDirectoryAddress:      bindings.AvsDirectoryAddr,
	}
}

func NewWriterFromConfig(
	cfg Config,
	ethClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainWriter, error) {
	wire.Build(
		NewBindingsFromConfig,
		NewElConfigFromAvsBindings,
		elcontracts.NewReaderFromConfig,
		wire.Bind(new(elcontracts.Reader), new(*elcontracts.ChainReader)),
		NewChainWriterFromBindings,
	)
	return &ChainWriter{}, nil
}

// Note that unlike reader/writer, a websocket eth client must be provided
func NewSubscriberFromConfig(
	cfg Config,
	wsClient eth.Client,
	logger logging.Logger,
) (*ChainSubscriber, error) {
	wire.Build(NewBindingsFromConfig, NewChainSubscriberFromBindings)
	return &ChainSubscriber{}, nil
}

type ClientsAndBindings struct {
	ChainReader      *ChainReader
	ChainWriter      *ChainWriter
	ChainSubscriber  *ChainSubscriber
	ContractBindings *ContractBindings
}

// ============= Workaround Code =================
// this section is needed because wire does not allow a Build dependency graph
// with two dependencies of the same type. See:
// https://github.com/google/wire/blob/main/docs/faq.md#what-if-my-dependency-graph-has-two-dependencies-of-the-same-type
// In BuildClientsFromConfig since both ethHttpClient and ethWsClient have type eth.Client,
// we need to wrap one of them in a new type to avoid the error.
type ethWsClient eth.Client

func newEthWsClient(
	wsClient eth.Client,
) ethWsClient {
	return (ethWsClient)(wsClient)
}

func newSubscriberFromConfigWithEthWsClientType(
	cfg Config,
	wsClient ethWsClient,
	logger logging.Logger,
) (*ChainSubscriber, error) {
	return NewSubscriberFromConfig(cfg, eth.Client(wsClient), logger)
}

func buildClientsFromConfig(
	cfg Config,
	ethHttpClient eth.Client,
	ethWsClient ethWsClient,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ClientsAndBindings, error) {
	wire.Build(
		NewBindingsFromConfig,
		NewReaderFromConfig,
		NewWriterFromConfig,
		newSubscriberFromConfigWithEthWsClientType,
		wire.Struct(new(ClientsAndBindings), "*"),
	)
	return &ClientsAndBindings{}, nil
}

// ============= End of Workaround Code =================

func BuildClientsFromConfig(
	cfg Config,
	ethHttpClient eth.Client,
	ethWsClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ClientsAndBindings, error) {
	return buildClientsFromConfig(cfg, ethHttpClient, ethWsClient, txMgr, logger)
}
