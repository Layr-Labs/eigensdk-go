// bindings.go contains functions that create contract bindings for the Eigenlayer and AVS contracts.
// These functions are meant to be used by constructors of the chainio package.
package utils

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsoperatorstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	blspubkeyregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	blsregistrycoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type EigenlayerContractBindings struct {
	SlasherAddr             gethcommon.Address
	StrategyManagerAddr     gethcommon.Address
	DelegationManagerAddr   gethcommon.Address
	BlspubkeyCompendiumAddr gethcommon.Address
	Slasher                 *slasher.ContractSlasher
	DelegationManager       *delegationmanager.ContractDelegationManager
	StrategyManager         *strategymanager.ContractStrategyManager
	// TODO: should this really be here? or moved to avs bindings
	BlsPubkeyCompendium *blspubkeycompendium.ContractBLSPublicKeyCompendium
}

func NewEigenlayerContractBindings(
	slasherAddr gethcommon.Address,
	blsPubKeyCompendiumAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*EigenlayerContractBindings, error) {
	contractSlasher, err := slasher.NewContractSlasher(slasherAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch Slasher contract", "err", err)
		return nil, err
	}

	strategyManagerAddr, err := contractSlasher.StrategyManager(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch StrategyManager address", "err", err)
		return nil, err
	}
	contractStrategyManager, err := strategymanager.NewContractStrategyManager(strategyManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch StrategyManager contract", "err", err)
		return nil, err
	}
	delegationManagerAddr, err := contractSlasher.Delegation(&bind.CallOpts{})
	if err != nil {
		logger.Error("Failed to fetch DelegationManager address", "err", err)
		return nil, err
	}
	contractDelegationManager, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, ethclient)
	if err != nil {
		logger.Error("Failed to fetch DelegationManager contract", "err", err)
		return nil, err
	}

	contractBlsCompendium, err := blspubkeycompendium.NewContractBLSPublicKeyCompendium(
		blsPubKeyCompendiumAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSPublicKeyCompendium contract", "err", err)
		return nil, err
	}

	return &EigenlayerContractBindings{
		SlasherAddr:             slasherAddr,
		StrategyManagerAddr:     strategyManagerAddr,
		DelegationManagerAddr:   delegationManagerAddr,
		BlspubkeyCompendiumAddr: blsPubKeyCompendiumAddr,
		Slasher:                 contractSlasher,
		StrategyManager:         contractStrategyManager,
		DelegationManager:       contractDelegationManager,
		BlsPubkeyCompendium:     contractBlsCompendium,
	}, nil
}

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type AvsRegistryContractBindings struct {
	// contract addresses
	RegistryCoordinatorAddr gethcommon.Address
	StakeRegistryAddr       gethcommon.Address
	BlsPubkeyRegistryAddr   gethcommon.Address
	BlsPubkeyCompendiumAddr gethcommon.Address
	OperatorStateRetriever  gethcommon.Address
	// contract bindings
	RegistryCoordinator       *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices
	StakeRegistry             *stakeregistry.ContractStakeRegistry
	BlsPubkeyRegistry         *blspubkeyregistry.ContractBLSPubkeyRegistry
	BlsPubkeyCompendium       *blspubkeycompendium.ContractBLSPublicKeyCompendium
	BlsOperatorStateRetriever *blsoperatorstateretriever.ContractBLSOperatorStateRetriever
}

func NewAVSRegistryContractBindings(
	blsRegistryCoordinatorAddr gethcommon.Address,
	blsOperatorStateRetrieverAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*AvsRegistryContractBindings, error) {
	contractBlsRegistryCoordinator, err := blsregistrycoordinator.NewContractBLSRegistryCoordinatorWithIndices(
		blsRegistryCoordinatorAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSRegistryCoordinator contract", "err", err)
		return nil, err
	}

	stakeregistryAddr, err := contractBlsRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(
		stakeregistryAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch StakeRegistry contract", "err", err)
		return nil, err
	}

	blsPubkeyRegistryAddr, err := contractBlsRegistryCoordinator.BlsPubkeyRegistry(&bind.CallOpts{})
	contractBlsPubkeyRegistry, err := blspubkeyregistry.NewContractBLSPubkeyRegistry(
		blsPubkeyRegistryAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSPubkeyRegistry contract", "err", err)
		return nil, err
	}

	blsPubkeyCompendiumAddr, err := contractBlsPubkeyRegistry.PubkeyCompendium(&bind.CallOpts{})
	contractBlsPubkeyCompendium, err := blspubkeycompendium.NewContractBLSPublicKeyCompendium(
		blsPubkeyCompendiumAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSPublicKeyCompendium contract", "err", err)
		return nil, err
	}

	contractBlsOperatorStateRetriever, err := blsoperatorstateretriever.NewContractBLSOperatorStateRetriever(
		blsOperatorStateRetrieverAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch OperatorStateRetriever contract", "err", err)
		return nil, err
	}

	return &AvsRegistryContractBindings{
		RegistryCoordinatorAddr:   blsRegistryCoordinatorAddr,
		StakeRegistryAddr:         stakeregistryAddr,
		BlsPubkeyRegistryAddr:     blsPubkeyRegistryAddr,
		BlsPubkeyCompendiumAddr:   blsPubkeyCompendiumAddr,
		OperatorStateRetriever:    blsOperatorStateRetrieverAddr,
		RegistryCoordinator:       contractBlsRegistryCoordinator,
		StakeRegistry:             contractStakeRegistry,
		BlsPubkeyRegistry:         contractBlsPubkeyRegistry,
		BlsPubkeyCompendium:       contractBlsPubkeyCompendium,
		BlsOperatorStateRetriever: contractBlsOperatorStateRetriever,
	}, nil
}
