// bindings.go contains functions that create contract bindings for the Eigenlayer and AVS contracts.
// These functions are meant to be used by constructors of the chainio package.
package elcontracts

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

// ContractBindings Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type ContractBindings struct {
	SlasherAddr           gethcommon.Address
	StrategyManagerAddr   gethcommon.Address
	DelegationManagerAddr gethcommon.Address
	AvsDirectoryAddr      gethcommon.Address
	Slasher               *slasher.ContractISlasher
	DelegationManager     *delegationmanager.ContractDelegationManager
	StrategyManager       *strategymanager.ContractStrategyManager
	AvsDirectory          *avsdirectory.ContractAVSDirectory
}

func NewBindingsFromConfig(
	cfg types.ElChainReaderConfig,
	client eth.Client,
	logger logging.Logger,
) (*ContractBindings, error) {
	var contractDelegationManager *delegationmanager.ContractDelegationManager
	var contractSlasher *slasher.ContractISlasher
	var contractStrategyManager *strategymanager.ContractStrategyManager
	var slasherAddr gethcommon.Address
	var strategyManagerAddr gethcommon.Address
	var avsDirectory *avsdirectory.ContractAVSDirectory
	var err error

	if cfg.DelegationManagerAddress == gethcommon.HexToAddress("") {
		logger.Warn("DelegationManager address not provided, the calls to the contract will not work")
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

	if cfg.AvsDirectoryAddress == gethcommon.HexToAddress("") {
		logger.Warn("AVSDirectory address not provided, the calls to the contract will not work")
	} else {
		avsDirectory, err = avsdirectory.NewContractAVSDirectory(cfg.AvsDirectoryAddress, client)
		if err != nil {
			return nil, utils.WrapError("Failed to fetch AVSDirectory contract", err)
		}
	}

	return &ContractBindings{
		SlasherAddr:           slasherAddr,
		StrategyManagerAddr:   strategyManagerAddr,
		DelegationManagerAddr: cfg.DelegationManagerAddress,
		AvsDirectoryAddr:      cfg.AvsDirectoryAddress,
		Slasher:               contractSlasher,
		StrategyManager:       contractStrategyManager,
		DelegationManager:     contractDelegationManager,
		AvsDirectory:          avsDirectory,
	}, nil
}

// NewEigenlayerContractBindings creates a new ContractBindings struct with the provided contract addresses
// Deprecated: Use NewBindingsFromConfig instead
func NewEigenlayerContractBindings(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethclient eth.Client,
	logger logging.Logger,
) (*ContractBindings, error) {
	contractDelegationManager, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to create DelegationManager contract", err)
	}

	slasherAddr, err := contractDelegationManager.Slasher(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch Slasher address", err)
	}
	contractSlasher, err := slasher.NewContractISlasher(slasherAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch Slasher contract", err)
	}

	strategyManagerAddr, err := contractDelegationManager.StrategyManager(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch StrategyManager address", err)
	}
	contractStrategyManager, err := strategymanager.NewContractStrategyManager(strategyManagerAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch StrategyManager contract", err)
	}

	avsDirectory, err := avsdirectory.NewContractAVSDirectory(avsDirectoryAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch AVSDirectory contract", err)
	}

	return &ContractBindings{
		SlasherAddr:           slasherAddr,
		StrategyManagerAddr:   strategyManagerAddr,
		DelegationManagerAddr: delegationManagerAddr,
		AvsDirectoryAddr:      avsDirectoryAddr,
		Slasher:               contractSlasher,
		StrategyManager:       contractStrategyManager,
		DelegationManager:     contractDelegationManager,
		AvsDirectory:          avsDirectory,
	}, nil
}
