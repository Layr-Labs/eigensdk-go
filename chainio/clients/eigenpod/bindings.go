package eigenpod

import (
	ieigenpod "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/IEigenPod"
	ieigenpodmanager "github.com/Layr-Labs/eigenlayer-contracts/pkg/bindings/IEigenPodManager"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"

	"github.com/ethereum/go-ethereum/common"
)

type ContractBindings struct {
	Address common.Address
	*ieigenpod.IEigenPod
}

type ContractCallerBindings struct {
	Address common.Address
	*ieigenpod.IEigenPodCaller
}

type ManagerContractBindings struct {
	Address common.Address
	*ieigenpodmanager.IEigenPodManager
}

type ManagerContractCallerBindings struct {
	Address common.Address
	*ieigenpodmanager.IEigenPodManagerCaller
}

func NewContractBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ContractBindings, error) {
	pod, err := ieigenpod.NewIEigenPod(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ContractBindings{
		Address:   address,
		IEigenPod: pod,
	}, nil
}

func NewContractCallerBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ContractCallerBindings, error) {
	pod, err := ieigenpod.NewIEigenPodCaller(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ContractCallerBindings{
		Address:         address,
		IEigenPodCaller: pod,
	}, nil
}

func NewManagerContractBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ManagerContractBindings, error) {
	manager, err := ieigenpodmanager.NewIEigenPodManager(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ManagerContractBindings{
		Address:          address,
		IEigenPodManager: manager,
	}, nil
}

func NewManagerContractCallerBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ManagerContractCallerBindings, error) {
	manager, err := ieigenpodmanager.NewIEigenPodManagerCaller(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ManagerContractCallerBindings{
		Address:                address,
		IEigenPodManagerCaller: manager,
	}, nil
}
