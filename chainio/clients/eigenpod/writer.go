package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eigenpod/bindings"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
)

type ChainWriter struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	eigenPod  *bindings.IEigenPod
	txMgr     txmgr.TxManager
}

type ManagerChainWriter struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	manager   *bindings.IEigenPodManager
	txMgr     txmgr.TxManager
}

func newChainWriter(
	eigenPod *bindings.IEigenPod,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *ChainWriter {
	logger = logger.With(logging.ComponentKey, "eigenpod/writer")

	return &ChainWriter{
		logger:    logger,
		ethClient: ethClient,
		eigenPod:  eigenPod,
		txMgr:     txMgr,
	}
}

func newManagerChainWriter(
	manager *bindings.IEigenPodManager,
	ethClient eth.HttpBackend,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *ManagerChainWriter {
	logger = logger.With(logging.ComponentKey, "eigenpodmanager/writer")

	return &ManagerChainWriter{
		logger:    logger,
		ethClient: ethClient,
		manager:   manager,
		txMgr:     txMgr,
	}
}

func NewWriter(
	address common.Address,
	ethClient eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ChainWriter, error) {
	pod, err := bindings.NewIEigenPod(address, ethClient)
	if err != nil {
		return nil, err
	}

	return newChainWriter(pod, ethClient, logger, txMgr), nil
}

func NewManagerWriter(
	address common.Address,
	ethClient eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
) (*ManagerChainWriter, error) {
	manager, err := bindings.NewIEigenPodManager(address, ethClient)
	if err != nil {
		return nil, err
	}

	return newManagerChainWriter(manager, ethClient, logger, txMgr), nil
}

// TODO(madhur): Add methods to ChainWriter and ManagerChainWriter to interact with the contracts.
