package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"

	eigenpod "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IEigenPod"
	eigenpodmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IEigenPodManager"
)

// ContractBindings contains the addresses and bindings for the EigenPod contracts
type ContractBindings struct {
	EigenPodAddr        common.Address
	EigenPodManagerAddr common.Address
	EigenPod            *eigenpod.ContractIEigenPod
	EigenPodManager     *eigenpodmanager.ContractIEigenPodManager
}

// NewBindingsFromConfig creates a new ContractBindings struct from the provided config
func NewBindingsFromConfig(
	cfg Config,
	client eth.HttpBackend,
	logger logging.Logger,
) (*ContractBindings, error) {
	var (
		err error

		eigenPod        *eigenpod.ContractIEigenPod
		eigenPodManager *eigenpodmanager.ContractIEigenPodManager
	)

	if utils.IsZeroAddress(cfg.EigenPodAddress) {
		logger.Debug("EigenPod address not provided, the calls to the contract will not work")
	} else {
		eigenPod, err = eigenpod.NewContractIEigenPod(cfg.EigenPodAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to create EigenPod contract", err)
		}
	}

	if utils.IsZeroAddress(cfg.EigenPodManagerAddress) {
		logger.Debug("EigenPodManager address not provided, the calls to the contract will not work")
	} else {
		eigenPodManager, err = eigenpodmanager.NewContractIEigenPodManager(cfg.EigenPodManagerAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to create EigenPodManager contract", err)
		}
	}

	return &ContractBindings{
		EigenPodAddr:        cfg.EigenPodAddress,
		EigenPodManagerAddr: cfg.EigenPodManagerAddress,
		EigenPod:            eigenPod,
		EigenPodManager:     eigenPodManager,
	}, nil
}
