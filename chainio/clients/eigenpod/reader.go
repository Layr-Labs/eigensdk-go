package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	eigenpod "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IEigenPod"
	eigenpodmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IEigenPodManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	EigenPodAddress        common.Address
	EigenPodManagerAddress common.Address
}

type ChainReader struct {
	logger    logging.Logger
	ethClient eth.HttpBackend

	eigenpod.ContractIEigenPodCaller
	eigenpodmanager.ContractIEigenPodManagerCaller
}

func NewChainReader(
	cfg Config,
	ethClient eth.HttpBackend,
	logger logging.Logger,
) (*ChainReader, error) {
	logger = logger.With(logging.ComponentKey, "eigenpod/reader")

	bindings, err := NewBindingsFromConfig(cfg, ethClient, logger)
	if err != nil {
		return nil, err
	}

	return &ChainReader{
		logger:                         logger,
		ethClient:                      ethClient,
		ContractIEigenPodCaller:        bindings.EigenPod.ContractIEigenPodCaller,
		ContractIEigenPodManagerCaller: bindings.EigenPodManager.ContractIEigenPodManagerCaller,
	}, nil
}
