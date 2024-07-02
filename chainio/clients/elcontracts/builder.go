package elcontracts

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/metrics"
)

func BuildClients(
	config Config,
	ethHttpClient eth.Client,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*ELChainReader, *ELChainWriter, *ContractBindings, error) {
	elContractBindings, err := NewBindingsFromConfig(
		config,
		ethHttpClient,
		logger,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	elChainReader := NewELChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		logger,
		ethHttpClient,
	)

	elChainWriter := NewELChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		ethHttpClient,
		logger,
		eigenMetrics,
		txMgr,
	)

	return elChainReader, elChainWriter, elContractBindings, nil
}
