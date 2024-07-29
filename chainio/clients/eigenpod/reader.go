package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eigenpod/bindings"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
)

type ChainReader struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	*bindings.IEigenPodCaller
}

type ManagerChainReader struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	*bindings.IEigenPodManagerCaller
}

func newChainReader(
	eigenPod *bindings.IEigenPodCaller,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) *ChainReader {
	logger = logger.With(logging.ComponentKey, "eigenpod/reader")

	return &ChainReader{
		logger:          logger,
		ethClient:       ethClient,
		IEigenPodCaller: eigenPod,
	}
}

func newManagerChainReader(
	manager *bindings.IEigenPodManagerCaller,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) *ManagerChainReader {
	logger = logger.With(logging.ComponentKey, "eigenpodmanager/reader")

	return &ManagerChainReader{
		logger:                 logger,
		ethClient:              ethClient,
		IEigenPodManagerCaller: manager,
	}
}

func NewReader(
	address common.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	pod, err := NewContractCallerBindings(address, ethClient)
	if err != nil {
		return nil, err
	}

	return newChainReader(pod.IEigenPodCaller, ethClient, logger), nil
}

func NewManagerReader(
	address common.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ManagerChainReader, error) {
	manager, err := NewManagerContractCallerBindings(address, ethClient)
	if err != nil {
		return nil, err
	}

	return newManagerChainReader(manager.IEigenPodManagerCaller, ethClient, logger), nil
}
