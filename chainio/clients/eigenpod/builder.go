package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
)

func BuildEigenPodClients(
	address common.Address,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainReader, *ChainWriter, *ContractBindings, error) {
	eigenPodBindings, err := NewContractBindings(address, client)
	if err != nil {
		return nil, nil, nil, err
	}

	eigenPodChainReader := newChainReader(
		&eigenPodBindings.IEigenPodCaller,
		client,
		logger,
	)

	eigenPodChainWriter := newChainWriter(
		eigenPodBindings.IEigenPod,
		client,
		logger,
		txMgr,
	)

	return eigenPodChainReader, eigenPodChainWriter, eigenPodBindings, nil
}

func BuildEigenPodManagerClients(
	address common.Address,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ManagerChainReader, *ManagerChainWriter, *ManagerContractBindings, error) {
	eigenPodManagerBindings, err := NewManagerContractBindings(address, client)
	if err != nil {
		return nil, nil, nil, err
	}

	eigenPodManagerChainReader := newManagerChainReader(
		&eigenPodManagerBindings.IEigenPodManagerCaller,
		client,
		logger,
	)

	eigenPodManagerChainWriter := newManagerChainWriter(
		eigenPodManagerBindings.IEigenPodManager,
		client,
		logger,
		txMgr,
	)

	return eigenPodManagerChainReader, eigenPodManagerChainWriter, eigenPodManagerBindings, nil
}
