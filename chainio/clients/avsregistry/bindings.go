package avsregistry

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	blsapkregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	indexregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/IndexRegistry"
	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	regcoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	servicemanager "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ServiceManagerBase"
	stakeregistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/StakeRegistry"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	gethcommon "github.com/ethereum/go-ethereum/common"
)

// ContractBindings Unclear to me why geth bindings don't store and expose the contract address...
// so we also store them here in case the different constructors that use this struct need them
type ContractBindings struct {
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
) (*ContractBindings, error) {
	contractBlsRegistryCoordinator, err := regcoordinator.NewContractRegistryCoordinator(
		registryCoordinatorAddr,
		ethclient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to create BLSRegistryCoordinator contract", err)
	}

	serviceManagerAddr, err := contractBlsRegistryCoordinator.ServiceManager(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch ServiceManager address", err)
	}
	contractServiceManager, err := servicemanager.NewContractServiceManagerBase(
		serviceManagerAddr,
		ethclient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch ServiceManager contract", err)
	}

	stakeregistryAddr, err := contractBlsRegistryCoordinator.StakeRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch StakeRegistry address", err)
	}
	contractStakeRegistry, err := stakeregistry.NewContractStakeRegistry(
		stakeregistryAddr,
		ethclient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch StakeRegistry contract", err)
	}

	blsApkRegistryAddr, err := contractBlsRegistryCoordinator.BlsApkRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch BLSPubkeyRegistry address", err)
	}
	contractBlsApkRegistry, err := blsapkregistry.NewContractBLSApkRegistry(
		blsApkRegistryAddr,
		ethclient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch BLSPubkeyRegistry contract", err)
	}

	indexRegistryAddr, err := contractBlsRegistryCoordinator.IndexRegistry(&bind.CallOpts{})
	if err != nil {
		return nil, utils.WrapError("Failed to fetch IndexRegistry address", err)
	}
	contractIndexRegistry, err := indexregistry.NewContractIndexRegistry(indexRegistryAddr, ethclient)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch IndexRegistry contract", err)
	}

	contractOperatorStateRetriever, err := opstateretriever.NewContractOperatorStateRetriever(
		operatorStateRetrieverAddr,
		ethclient,
	)
	if err != nil {
		return nil, utils.WrapError("Failed to fetch OperatorStateRetriever contract", err)
	}

	return &ContractBindings{
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
