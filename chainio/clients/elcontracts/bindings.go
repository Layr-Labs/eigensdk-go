// Package elcontracts bindings.go contains functions that create contract bindings for the Eigenlayer Core contracts
// These functions are meant to be used by constructors of the chainio package.
package elcontracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IAVSDirectory"
	rewardscoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IRewardsCoordinator"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
)

// ContractBindings contains the contract bindings for the EigenLayer Core contracts
//
// Unclear why geth bindings don't store and expose the contract address,
// so we also store them here in case the different constructors that use this struct need them
type ContractBindings struct {
	SlasherAddr               gethcommon.Address
	StrategyManagerAddr       gethcommon.Address
	DelegationManagerAddr     gethcommon.Address
	AvsDirectoryAddr          gethcommon.Address
	RewardsCoordinatorAddress gethcommon.Address
	Slasher                   *slasher.ContractISlasher
	DelegationManager         *delegationmanager.ContractDelegationManager
	StrategyManager           *strategymanager.ContractStrategyManager
	AvsDirectory              *avsdirectory.ContractIAVSDirectory
	RewardsCoordinator        *rewardscoordinator.ContractIRewardsCoordinator
}

func NewBindingsFromConfig(
	cfg Config,
	client eth.HttpBackend,
	logger logging.Logger,
) (*ContractBindings, error) {
	var (
		err error

		contractDelegationManager *delegationmanager.ContractDelegationManager
		contractSlasher           *slasher.ContractISlasher
		contractStrategyManager   *strategymanager.ContractStrategyManager
		slasherAddr               gethcommon.Address
		strategyManagerAddr       gethcommon.Address
		avsDirectory              *avsdirectory.ContractIAVSDirectory
		rewardsCoordinator        *rewardscoordinator.ContractIRewardsCoordinator
	)

	if isZeroAddress(cfg.DelegationManagerAddress) {
		logger.Debug("DelegationManager address not provided, the calls to the contract will not work")
	} else {
		contractDelegationManager, err = delegationmanager.NewContractDelegationManager(cfg.DelegationManagerAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to create DelegationManager contract", err)
		}

		slasherAddr, err = contractDelegationManager.Slasher(&bind.CallOpts{})
		if err != nil {
			return nil, utils.WrapError("Failed to fetch Slasher address", err)
		}
		contractSlasher, err = slasher.NewContractISlasher(slasherAddr, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch Slasher contract", err)
		}

		strategyManagerAddr, err = contractDelegationManager.StrategyManager(&bind.CallOpts{})
		if err != nil {
			return nil, utils.WrapError("Failed to fetch StrategyManager address", err)
		}
		contractStrategyManager, err = strategymanager.NewContractStrategyManager(strategyManagerAddr, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch StrategyManager contract", err)
		}
	}

	if isZeroAddress(cfg.AvsDirectoryAddress) {
		logger.Debug("AVSDirectory address not provided, the calls to the contract will not work")
	} else {
		avsDirectory, err = avsdirectory.NewContractIAVSDirectory(cfg.AvsDirectoryAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AVSDirectory contract", err)
		}
	}

	if isZeroAddress(cfg.RewardsCoordinatorAddress) {
		logger.Debug("RewardsCoordinator address not provided, the calls to the contract will not work")
	} else {
		rewardsCoordinator, err = rewardscoordinator.NewContractIRewardsCoordinator(cfg.RewardsCoordinatorAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch RewardsCoordinator contract", err)
		}
	}

	return &ContractBindings{
		SlasherAddr:               slasherAddr,
		StrategyManagerAddr:       strategyManagerAddr,
		DelegationManagerAddr:     cfg.DelegationManagerAddress,
		AvsDirectoryAddr:          cfg.AvsDirectoryAddress,
		RewardsCoordinatorAddress: cfg.RewardsCoordinatorAddress,
		Slasher:                   contractSlasher,
		StrategyManager:           contractStrategyManager,
		DelegationManager:         contractDelegationManager,
		AvsDirectory:              avsDirectory,
		RewardsCoordinator:        rewardsCoordinator,
	}, nil
}
func isZeroAddress(address gethcommon.Address) bool {
	return address == gethcommon.Address{}
}
