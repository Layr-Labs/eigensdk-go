package eigenpod

import (
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eigenpod/bindings"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"

	"github.com/ethereum/go-ethereum/common"
)

type ContractBindings struct {
	Address common.Address
	*bindings.IEigenPod
}

type ContractCallerBindings struct {
	Address common.Address
	*bindings.IEigenPodCaller
}

type ManagerContractBindings struct {
	Address common.Address
	*bindings.IEigenPodManager
}

type ManagerContractCallerBindings struct {
	Address common.Address
	*bindings.IEigenPodManagerCaller
}

func NewContractBindings(
	address common.Address,
	ethClient eth.HttpBackend,
) (*ContractBindings, error) {
	pod, err := bindings.NewIEigenPod(address, ethClient)
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
	pod, err := bindings.NewIEigenPodCaller(address, ethClient)
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
	manager, err := bindings.NewIEigenPodManager(address, ethClient)
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
	manager, err := bindings.NewIEigenPodManagerCaller(address, ethClient)
	if err != nil {
		return nil, err
	}
	return &ManagerContractCallerBindings{
		Address:                address,
		IEigenPodManagerCaller: manager,
	}, nil
}
