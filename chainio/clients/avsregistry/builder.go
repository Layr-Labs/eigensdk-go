package avsregistry

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
)

// Build an AVS registry client with the given configuration,
// HTTP and WS clients, and logger, but without a private key.
//
// This is useful for read-only operations, such as fetching metrics.
func BuildClientsForEcMetrics(
	config Config,
	client eth.HttpBackend,
	wsClient eth.WsBackend,
	logger logging.Logger,
) (*ChainReader, *ChainSubscriber, *ContractBindings, error) {
	avsBindings, err := NewBindingsFromConfig(
		config,
		client,
		logger,
	)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	chainReader := NewChainReader(
		avsBindings.RegistryCoordinatorAddr,
		avsBindings.BlsApkRegistryAddr,
		avsBindings.RegistryCoordinator,
		avsBindings.OperatorStateRetriever,
		avsBindings.StakeRegistry,
		logger,
		client,
	)

	chainSubscriber, err := NewSubscriberFromConfig(
		config,
		wsClient,
		logger,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	// This is ugly, but we need elReader to be able to create the AVS writer
	elChainReader, err := elcontracts.NewReaderFromConfig(
		elcontracts.Config{
			DelegationManagerAddress: avsBindings.DelegationManagerAddr,
			AvsDirectoryAddress:      avsBindings.AvsDirectoryAddr,
		},
		client,
		logger,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	return chainReader, chainSubscriber, avsBindings, nil
}

func BuildClients(
	config Config,
	client eth.HttpBackend,
	wsClient eth.WsBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainReader, *ChainSubscriber, *ChainWriter, *ContractBindings, error) {
	avsBindings, err := NewBindingsFromConfig(
		config,
		client,
		logger,
	)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	chainReader := NewChainReader(
		avsBindings.RegistryCoordinatorAddr,
		avsBindings.BlsApkRegistryAddr,
		avsBindings.RegistryCoordinator,
		avsBindings.OperatorStateRetriever,
		avsBindings.StakeRegistry,
		logger,
		client,
	)

	chainSubscriber, err := NewSubscriberFromConfig(
		config,
		wsClient,
		logger,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// This is ugly, but we need elReader to be able to create the AVS writer
	elChainReader, err := elcontracts.NewReaderFromConfig(
		elcontracts.Config{
			DelegationManagerAddress: avsBindings.DelegationManagerAddr,
			AvsDirectoryAddress:      avsBindings.AvsDirectoryAddr,
		},
		client,
		logger,
	)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	chainWriter := NewChainWriter(
		avsBindings.ServiceManagerAddr,
		avsBindings.RegistryCoordinator,
		avsBindings.OperatorStateRetriever,
		avsBindings.StakeRegistry,
		avsBindings.BlsApkRegistry,
		elChainReader,
		logger,
		client,
		txMgr,
	)

	return chainReader, chainSubscriber, chainWriter, avsBindings, nil
}
