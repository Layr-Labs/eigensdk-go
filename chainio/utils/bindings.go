// bindings.go contains functions that create contract bindings for the Eigenlayer and AVS contracts.
// These functions are meant to be used by constructors of the chainio package.
package utils

import (
	"errors"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	avsdirectory "github.com/Layr-Labs/eigensdk-go/contracts/bindings/AVSDirectory"
	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ISlasher"
	indexregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IndexRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type EigenlayerContractBindings struct {
	SlasherAddr           gethcommon.Address
	StrategyManagerAddr   gethcommon.Address
	DelegationManagerAddr gethcommon.Address
	AvsDirectoryAddr      gethcommon.Address
	Slasher               *slasher.ContractISlasher
	DelegationManager     *delegationmanager.ContractDelegationManager
	StrategyManager       *strategymanager.ContractStrategyManager
	AvsDirectory          *avsdirectory.ContractAVSDirectory
}

func NewEigenlayerContractBindings(
	delegationManagerAddr gethcommon.Address,
	avsDirectoryAddr gethcommon.Address,
	ethclient eth.Client,
	logger logging.Logger,
) (*EigenlayerContractBindings, error) {
	contractDelegationManager, err := delegationmanager.NewContractDelegationManager(delegationManagerAddr, ethclient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create DelegationManager contract"), err)
	}

	slasherAddr, err := contractDelegationManager.Slasher(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch Slasher address"), err)
	}
	contractSlasher, err := slasher.NewContractISlasher(slasherAddr, ethclient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch Slasher contract"), err)
	}

	strategyManagerAddr, err := contractDelegationManager.StrategyManager(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch StrategyManager address"), err)
	}
	contractStrategyManager, err := strategymanager.NewContractStrategyManager(strategyManagerAddr, ethclient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch StrategyManager contract"), err)
	}

	avsDirectory, err := avsdirectory.NewContractAVSDirectory(avsDirectoryAddr, ethclient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch AVSDirectory contract"), err)
	}

	return &EigenlayerContractBindings{
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

// Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type AvsRegistryContractBindings struct {
	// contract addresses
	ServiceManagerAddr         gethcommon.Address
	RegistryCoordinatorAddr    gethcommon.Address
	StakeRegistryAddr          gethcommon.Address
	BlsApkRegistryAddr         gethcommon.Address
	OperatorStateRetrieverAddr gethcommon.Address
	IndexRegistryAddr          gethcommon.Address
	// contract bindings
	ServiceManager         *servicemanager.ContractServiceManagerBase
	RegistryCoordinator    *regcoordinator.ContractRegistryCoordinator
	StakeRegistry          *stakeregistry.ContractStakeRegistry
	BlsApkRegistry         *blsapkregistry.ContractBLSApkRegistry
	IndexRegistry          *indexregistry.ContractIndexRegistry
	OperatorStateRetriever *opstateretriever.ContractOperatorStateRetriever
}

func NewAVSRegistryContractBindings(
	registryCoordinatorAddr gethcommon.Address,
	operatorStateRetrieverAddr gethcommon.Address,
	ethclient eth.Client,
	logger logging.Logger,
) (*AvsRegistryContractBindings, error) {
	contractBlsRegistryCoordinator, err := regcoordinator.NewContractRegistryCoordinator(
		registryCoordinatorAddr,
		ethclient,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to create BLSRegistryCoordinator contract"), err)
	}

	serviceManagerAddr, err := contractBlsRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch ServiceManager address"), err)
	}
	contractServiceManager, err := servicemanager.NewContractServiceManagerBase(
		serviceManagerAddr,
		ethclient,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch ServiceManager contract"), err)
	}

	stakeregistryAddr, err := contractBlsRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch StakeRegistry address"), err)
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(
		stakeregistryAddr,
		ethclient,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch StakeRegistry contract"), err)
	}

	blsApkRegistryAddr, err := contractBlsRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch BLSPubkeyRegistry address"), err)
	}
	contractBlsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(
		blsApkRegistryAddr,
		ethclient,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch BLSPubkeyRegistry contract"), err)
	}

	indexRegistryAddr, err := contractBlsRegistryCoordinator.IndexRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch IndexRegistry address"), err)
	}
	contractIndexRegistry, err := indexregistry.NewContractIndexRegistry(indexRegistryAddr, ethclient)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch IndexRegistry contract"), err)
	}

	contractOperatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethclient,
	)
	if err != nil {
		return nil, types.WrapError(errors.New("Failed to fetch OperatorStateRetriever contract"), err)
	}

	return &AvsRegistryContractBindings{
		ServiceManagerAddr:         serviceManagerAddr,
		RegistryCoordinatorAddr:    registryCoordinatorAddr,
		StakeRegistryAddr:          stakeregistryAddr,
		BlsApkRegistryAddr:         blsApkRegistryAddr,
		IndexRegistryAddr:          indexRegistryAddr,
		OperatorStateRetrieverAddr: operatorStateRetrieverAddr,
		ServiceManager:             contractServiceManager,
		RegistryCoordinator:        contractBlsRegistryCoordinator,
		StakeRegistry:              contractStakeRegistry,
		BlsApkRegistry:             contractBlsApkRegistry,
		IndexRegistry:              contractIndexRegistry,
		OperatorStateRetriever:     contractOperatorStateRetriever,
	}, nil
}
