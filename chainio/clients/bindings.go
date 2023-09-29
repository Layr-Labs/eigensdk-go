package clients

import (
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	blspubkeyregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPubkeyRegistry"
	blspubkeycompendium "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSPublicKeyCompendium"
	blsregistrycoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSRegistryCoordinatorWithIndices"
	delegationmanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/DelegationManager"
	slasher "github.com/Layr-Labs/eigensdk-go/contracts/bindings/Slasher"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	strategymanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StrategyManager"
)

type eigenlayerContractBindings struct {
	Slasher             *slasher.ContractSlasher
	DelegationManager   *delegationmanager.ContractDelegationManager
	StrategyManager     *strategymanager.ContractStrategyManager
	BlsPubKeyCompendium *blspubkeycompendium.ContractBLSPublicKeyCompendium
	logger              logging.Logger
}

func newEigenlayerContractBindings(
	slasherAddr gethcommon.Address,
	blsPubKeyCompendiumAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*eigenlayerContractBindings, error) {
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

	return &eigenlayerContractBindings{
		Slasher:             contractSlasher,
		StrategyManager:     contractStrategyManager,
		DelegationManager:   contractDelegationManager,
		BlsPubKeyCompendium: contractBlsCompendium,
		logger:              logger,
	}, nil
}

type avsRegistryContractBindings struct {
	RegistryCoordinator       *blsregistrycoordinator.ContractBLSRegistryCoordinatorWithIndices
	BlsOperatorStateRetriever *blsoperatorstateretrievar.ContractBLSOperatorStateRetriever
	StakeRegistry             *stakeregistry.ContractStakeRegistry
	BlsPubkeyRegistry         *blspubkeyregistry.ContractBLSPubkeyRegistry
}

func newAVSRegistryContractBindings(
	blsRegistryCoordinatorAddr gethcommon.Address,
	blsOperatorStateRetrieverAddr gethcommon.Address,
	stakeregistryAddr gethcommon.Address,
	BlsPubkeyRegistryAddr gethcommon.Address,
	ethclient eth.EthClient,
	logger logging.Logger,
) (*avsRegistryContractBindings, error) {
	contractBlsRegistryCoordinator, err := blsregistrycoordinator.NewContractBLSRegistryCoordinatorWithIndices(
		blsRegistryCoordinatorAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSRegistryCoordinator contract", "err", err)
		return nil, err
	}

	contractBlsOperatorStateRetriever, err := blsoperatorstateretrievar.NewContractBLSOperatorStateRetriever(
		blsOperatorStateRetrieverAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch OperatorStateRetriever contract", "err", err)
		return nil, err
	}

	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(
		stakeregistryAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch StakeRegistry contract", "err", err)
		return nil, err
	}

	contractBlsPubkeyRegistry, err := blspubkeyregistry.NewContractBLSPubkeyRegistry(
		BlsPubkeyRegistryAddr,
		ethclient,
	)
	if err != nil {
		logger.Error("Failed to fetch BLSPubkeyRegistry contract", "err", err)
		return nil, err
	}

	return &avsRegistryContractBindings{
		RegistryCoordinator:       contractBlsRegistryCoordinator,
		BlsOperatorStateRetriever: contractBlsOperatorStateRetriever,
		StakeRegistry:             contractStakeRegistry,
		BlsPubkeyRegistry:         contractBlsPubkeyRegistry,
	}, nil
}
