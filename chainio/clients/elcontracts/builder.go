package elcontracts

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
)

// Deprecated: Use BuildClientsFromConfig instead.
func BuildClients(
	config Config,
	client eth.Client,
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
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		client,
		logger,
		eigenMetrics,
		txMgr,
	)

	return elChainReader, elChainWriter, elContractBindings, nil
}
