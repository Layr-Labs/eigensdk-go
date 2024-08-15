package elcontracts

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
)

// Returns a tuple of reader clients with the given:
// configuration, HTTP client, logger and metrics.
func BuildReadClients(
	config Config,
	client eth.HttpBackend,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*ChainReader, *ContractBindings, error) {
	elContractBindings, err := NewBindingsFromConfig(
		config,
		client,
		logger,
	)
	if err != nil {
		return nil, nil, err
	}

	elChainReader := NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		client,
	)

	return elChainReader, elContractBindings, nil
}

func BuildClients(
	config Config,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*ChainReader, *ChainWriter, *ContractBindings, error) {
	elContractBindings, err := NewBindingsFromConfig(
		config,
		client,
		logger,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	elChainReader := NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		client,
	)

	elChainWriter := NewChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.AvsDirectory,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		client,
		logger,
		eigenMetrics,
		txMgr,
	)

	return elChainReader, elChainWriter, elContractBindings, nil
}
