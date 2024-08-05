package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eigenpod/bindings"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"

	"github.com/ethereum/go-ethereum/common"
)

// ChainReader is a reader for the EigenPod contract.
// We want it to be different from the ManagerChainReader since as a user,
// I only need this reader to manage my EigenPod. So there's no need for ManagerChainReader
type ChainReader struct {
	logger    logging.Logger
	ethClient eth.HttpBackend
	*bindings.IEigenPodCaller
}

// ManagerChainReader is a reader for the EigenPodManager contract.
// We want it to be different from the ChainReader since as a user, this is only
// needed to get overall state of all the EigenPods in the system.
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
	eigenPodAddress common.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	pod, err := NewContractCallerBindings(eigenPodAddress, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create EigenPod contract", err)
	}

	return newChainReader(pod.IEigenPodCaller, ethClient, logger), nil
}

func NewManagerReader(
	eigenPodManagerAddress common.Address,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ManagerChainReader, error) {
	manager, err := NewManagerContractCallerBindings(eigenPodManagerAddress, ethClient)
	if err != nil {
		return nil, utils.WrapError("Failed to create EigenPodManager contract", err)
	}

	return newManagerChainReader(manager.IEigenPodManagerCaller, ethClient, logger), nil
}
