// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractPermissionController

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractPermissionControllerMetaData contains all meta data concerning the ContractPermissionController contract.
var ContractPermissionControllerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canCall\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointeePermissions\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAppointees\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAdmins\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pendingAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removePendingAdmin\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAppointee\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AppointeeSet\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"appointee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":false,\"internalType\":\"bytes4\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminAdded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingAdminRemoved\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AdminAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AdminNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AppointeeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotHaveZeroAdmins\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60d3565b5f54610100900460ff161560835760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff9081161460d1575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b610f40806100e05f395ff3fe608060405234801561000f575f5ffd5b50600436106100cb575f3560e01c80639100674511610088578063ad8aca7711610063578063ad8aca77146101b0578063df595cb8146101c3578063eb5a4e87146101d6578063fddbdefd146101e9575f5ffd5b80639100674514610167578063950d806e1461018a578063ad5f22101461019d575f5ffd5b806306641201146100cf578063268959e5146100e45780634f906cf9146100f7578063628806ef1461010a5780636bddfa1f1461011d578063882a3b3814610146575b5f5ffd5b6100e26100dd366004610cfa565b6101fc565b005b6100e26100f2366004610d4b565b61031d565b6100e2610105366004610d4b565b6103f8565b6100e2610118366004610d7c565b61049b565b61013061012b366004610d7c565b610529565b60405161013d9190610dd8565b60405180910390f35b610159610154366004610d4b565b610552565b60405161013d929190610dea565b61017a610175366004610d4b565b6106b3565b604051901515815260200161013d565b6100e2610198366004610cfa565b610723565b6101306101ab366004610d7c565b610834565b61017a6101be366004610d4b565b6108da565b61017a6101d1366004610cfa565b6108fb565b6100e26101e4366004610d4b565b610950565b6101306101f7366004610e4c565b610a1e565b8361020781336106b3565b61022457604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0385165f908152600160205260408120906102468585610a5c565b6001600160a01b0387165f908152600484016020526040902090915061026c9082610a89565b6102895760405163262118cd60e01b815260040160405180910390fd5b6001600160a01b0386165f90815260048301602052604090206102ac9082610aa0565b505f81815260058301602052604090206102c69087610aab565b50856001600160a01b0316876001600160a01b03167f18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6878760405161030c929190610e8c565b60405180910390a350505050505050565b8161032881336106b3565b61034557604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602081905260409091206002019061036c82610abf565b1161038a576040516310ce892b60e31b815260040160405180910390fd5b6103948184610aab565b6103b157604051630716d81b60e51b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce906020015b60405180910390a250505050565b8161040381336106b3565b61042057604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f9081526001602052604090206104418184610aab565b61045e5760405163bed8295f60e01b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7906020016103ea565b6001600160a01b0381165f9081526001602052604090206104bc8133610aab565b6104d95760405163bed8295f60e01b815260040160405180910390fd5b6104e66002820133610ac8565b506040513381526001600160a01b038316907fbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff979060200160405180910390a25050565b6001600160a01b0381165f90815260016020526040902060609061054c90610adc565b92915050565b6001600160a01b038083165f90815260016020908152604080832093851683526004909301905290812060609182919061058b82610abf565b90505f8167ffffffffffffffff8111156105a7576105a7610eaf565b6040519080825280602002602001820160405280156105d0578160200160208202803683370190505b5090505f8267ffffffffffffffff8111156105ed576105ed610eaf565b604051908082528060200260200182016040528015610616578160200160208202803683370190505b5090505f5b838110156106a5576106496106308683610ae8565b606081901c9160a09190911b6001600160e01b03191690565b84838151811061065b5761065b610ec3565b6020026020010184848151811061067457610674610ec3565b6001600160e01b0319909316602093840291909101909201919091526001600160a01b03909116905260010161061b565b509097909650945050505050565b6001600160a01b0382165f9081526001602052604081206106d690600201610abf565b5f036106f857816001600160a01b0316836001600160a01b031614905061054c565b6001600160a01b0383165f90815260016020526040902061071c9060020183610af3565b9392505050565b8361072e81336106b3565b61074b57604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0385165f9081526001602052604081209061076d8585610a5c565b6001600160a01b0387165f90815260048401602052604090209091506107939082610a89565b156107b15760405163ad8efeb760e01b815260040160405180910390fd5b6001600160a01b0386165f90815260048301602052604090206107d49082610b14565b505f81815260058301602052604090206107ee9087610ac8565b50856001600160a01b0316876001600160a01b03167f037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169878760405161030c929190610e8c565b6001600160a01b0381165f90815260016020526040902060609061085a90600201610abf565b5f036108b2576040805160018082528183019092525f916020808301908036833701905050905082815f8151811061089457610894610ec3565b6001600160a01b039092166020928302919091019091015292915050565b6001600160a01b0382165f90815260016020526040902061054c90600201610adc565b919050565b6001600160a01b0382165f90815260016020526040812061071c9083610af3565b5f61090685856106b3565b8061094757506109476109198484610a5c565b6001600160a01b038088165f908152600160209081526040808320938a168352600490930190522090610a89565b95945050505050565b8161095b81336106b3565b61097857604051637bfa4b9f60e01b815260040160405180910390fd5b6001600160a01b0383165f90815260016020526040902061099c6002820184610af3565b156109ba5760405163130160e560e31b815260040160405180910390fd5b6109c48184610ac8565b6109e1576040516319abede360e11b815260040160405180910390fd5b6040516001600160a01b0384811682528516907fb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c906020016103ea565b60605f610a2b8484610a5c565b6001600160a01b0386165f908152600160209081526040808320848452600501909152902090915061094790610adc565b60609190911b6bffffffffffffffffffffffff191660a09190911c6bffffffff0000000000000000161790565b5f818152600183016020526040812054151561071c565b5f61071c8383610b1f565b5f61071c836001600160a01b038416610b1f565b5f61054c825490565b5f61071c836001600160a01b038416610c02565b60605f61071c83610c4e565b5f61071c8383610ca7565b6001600160a01b0381165f908152600183016020526040812054151561071c565b5f61071c8383610c02565b5f8181526001830160205260408120548015610bf9575f610b41600183610ed7565b85549091505f90610b5490600190610ed7565b9050818114610bb3575f865f018281548110610b7257610b72610ec3565b905f5260205f200154905080875f018481548110610b9257610b92610ec3565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080610bc457610bc4610ef6565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f90556001935050505061054c565b5f91505061054c565b5f818152600183016020526040812054610c4757508154600181810184555f84815260208082209093018490558454848252828601909352604090209190915561054c565b505f61054c565b6060815f01805480602002602001604051908101604052809291908181526020018280548015610c9b57602002820191905f5260205f20905b815481526020019060010190808311610c87575b50505050509050919050565b5f825f018281548110610cbc57610cbc610ec3565b905f5260205f200154905092915050565b80356001600160a01b03811681146108d5575f5ffd5b80356001600160e01b0319811681146108d5575f5ffd5b5f5f5f5f60808587031215610d0d575f5ffd5b610d1685610ccd565b9350610d2460208601610ccd565b9250610d3260408601610ccd565b9150610d4060608601610ce3565b905092959194509250565b5f5f60408385031215610d5c575f5ffd5b610d6583610ccd565b9150610d7360208401610ccd565b90509250929050565b5f60208284031215610d8c575f5ffd5b61071c82610ccd565b5f8151808452602084019350602083015f5b82811015610dce5781516001600160a01b0316865260209586019590910190600101610da7565b5093949350505050565b602081525f61071c6020830184610d95565b604081525f610dfc6040830185610d95565b82810360208401528084518083526020830191506020860192505f5b81811015610e405783516001600160e01b031916835260209384019390920191600101610e18565b50909695505050505050565b5f5f5f60608486031215610e5e575f5ffd5b610e6784610ccd565b9250610e7560208501610ccd565b9150610e8360408501610ce3565b90509250925092565b6001600160a01b039290921682526001600160e01b031916602082015260400190565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b8181038181111561054c57634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52603160045260245ffdfea2646970667358221220398327e64bb22b1896a4965ed1042c1c4193c2369ed4f6f24d879b8ae17eb45f64736f6c634300081b0033",
}

// ContractPermissionControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractPermissionControllerMetaData.ABI instead.
var ContractPermissionControllerABI = ContractPermissionControllerMetaData.ABI

// ContractPermissionControllerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractPermissionControllerMetaData.Bin instead.
var ContractPermissionControllerBin = ContractPermissionControllerMetaData.Bin

// DeployContractPermissionController deploys a new Ethereum contract, binding an instance of ContractPermissionController to it.
func DeployContractPermissionController(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractPermissionController, error) {
	parsed, err := ContractPermissionControllerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractPermissionControllerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractPermissionController{ContractPermissionControllerCaller: ContractPermissionControllerCaller{contract: contract}, ContractPermissionControllerTransactor: ContractPermissionControllerTransactor{contract: contract}, ContractPermissionControllerFilterer: ContractPermissionControllerFilterer{contract: contract}}, nil
}

// ContractPermissionControllerMethods is an auto generated interface around an Ethereum contract.
type ContractPermissionControllerMethods interface {
	ContractPermissionControllerCalls
	ContractPermissionControllerTransacts
	ContractPermissionControllerFilters
}

// ContractPermissionControllerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractPermissionControllerCalls interface {
	CanCall(opts *bind.CallOpts, account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error)

	GetAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error)

	GetAppointeePermissions(opts *bind.CallOpts, account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error)

	GetAppointees(opts *bind.CallOpts, account common.Address, target common.Address, selector [4]byte) ([]common.Address, error)

	GetPendingAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error)

	IsAdmin(opts *bind.CallOpts, account common.Address, caller common.Address) (bool, error)

	IsPendingAdmin(opts *bind.CallOpts, account common.Address, pendingAdmin common.Address) (bool, error)
}

// ContractPermissionControllerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractPermissionControllerTransacts interface {
	AcceptAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	AddPendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error)

	RemoveAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error)

	RemoveAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error)

	RemovePendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error)

	SetAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error)
}

// ContractPermissionControllerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractPermissionControllerFilters interface {
	FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerAdminRemovedIterator, error)
	WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAdminRemoved, account []common.Address) (event.Subscription, error)
	ParseAdminRemoved(log types.Log) (*ContractPermissionControllerAdminRemoved, error)

	FilterAdminSet(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerAdminSetIterator, error)
	WatchAdminSet(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAdminSet, account []common.Address) (event.Subscription, error)
	ParseAdminSet(log types.Log) (*ContractPermissionControllerAdminSet, error)

	FilterAppointeeRemoved(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*ContractPermissionControllerAppointeeRemovedIterator, error)
	WatchAppointeeRemoved(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAppointeeRemoved, account []common.Address, appointee []common.Address) (event.Subscription, error)
	ParseAppointeeRemoved(log types.Log) (*ContractPermissionControllerAppointeeRemoved, error)

	FilterAppointeeSet(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*ContractPermissionControllerAppointeeSetIterator, error)
	WatchAppointeeSet(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAppointeeSet, account []common.Address, appointee []common.Address) (event.Subscription, error)
	ParseAppointeeSet(log types.Log) (*ContractPermissionControllerAppointeeSet, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractPermissionControllerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractPermissionControllerInitialized, error)

	FilterPendingAdminAdded(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerPendingAdminAddedIterator, error)
	WatchPendingAdminAdded(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerPendingAdminAdded, account []common.Address) (event.Subscription, error)
	ParsePendingAdminAdded(log types.Log) (*ContractPermissionControllerPendingAdminAdded, error)

	FilterPendingAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerPendingAdminRemovedIterator, error)
	WatchPendingAdminRemoved(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerPendingAdminRemoved, account []common.Address) (event.Subscription, error)
	ParsePendingAdminRemoved(log types.Log) (*ContractPermissionControllerPendingAdminRemoved, error)
}

// ContractPermissionController is an auto generated Go binding around an Ethereum contract.
type ContractPermissionController struct {
	ContractPermissionControllerCaller     // Read-only binding to the contract
	ContractPermissionControllerTransactor // Write-only binding to the contract
	ContractPermissionControllerFilterer   // Log filterer for contract events
}

// ContractPermissionController implements the ContractPermissionControllerMethods interface.
var _ ContractPermissionControllerMethods = (*ContractPermissionController)(nil)

// ContractPermissionControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractPermissionControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPermissionControllerCaller implements the ContractPermissionControllerCalls interface.
var _ ContractPermissionControllerCalls = (*ContractPermissionControllerCaller)(nil)

// ContractPermissionControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractPermissionControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPermissionControllerTransactor implements the ContractPermissionControllerTransacts interface.
var _ ContractPermissionControllerTransacts = (*ContractPermissionControllerTransactor)(nil)

// ContractPermissionControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractPermissionControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractPermissionControllerFilterer implements the ContractPermissionControllerFilters interface.
var _ ContractPermissionControllerFilters = (*ContractPermissionControllerFilterer)(nil)

// ContractPermissionControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractPermissionControllerSession struct {
	Contract     *ContractPermissionController // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractPermissionControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractPermissionControllerCallerSession struct {
	Contract *ContractPermissionControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ContractPermissionControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractPermissionControllerTransactorSession struct {
	Contract     *ContractPermissionControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ContractPermissionControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractPermissionControllerRaw struct {
	Contract *ContractPermissionController // Generic contract binding to access the raw methods on
}

// ContractPermissionControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractPermissionControllerCallerRaw struct {
	Contract *ContractPermissionControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractPermissionControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractPermissionControllerTransactorRaw struct {
	Contract *ContractPermissionControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractPermissionController creates a new instance of ContractPermissionController, bound to a specific deployed contract.
func NewContractPermissionController(address common.Address, backend bind.ContractBackend) (*ContractPermissionController, error) {
	contract, err := bindContractPermissionController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionController{ContractPermissionControllerCaller: ContractPermissionControllerCaller{contract: contract}, ContractPermissionControllerTransactor: ContractPermissionControllerTransactor{contract: contract}, ContractPermissionControllerFilterer: ContractPermissionControllerFilterer{contract: contract}}, nil
}

// NewContractPermissionControllerCaller creates a new read-only instance of ContractPermissionController, bound to a specific deployed contract.
func NewContractPermissionControllerCaller(address common.Address, caller bind.ContractCaller) (*ContractPermissionControllerCaller, error) {
	contract, err := bindContractPermissionController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerCaller{contract: contract}, nil
}

// NewContractPermissionControllerTransactor creates a new write-only instance of ContractPermissionController, bound to a specific deployed contract.
func NewContractPermissionControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractPermissionControllerTransactor, error) {
	contract, err := bindContractPermissionController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerTransactor{contract: contract}, nil
}

// NewContractPermissionControllerFilterer creates a new log filterer instance of ContractPermissionController, bound to a specific deployed contract.
func NewContractPermissionControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractPermissionControllerFilterer, error) {
	contract, err := bindContractPermissionController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerFilterer{contract: contract}, nil
}

// bindContractPermissionController binds a generic wrapper to an already deployed contract.
func bindContractPermissionController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractPermissionControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPermissionController *ContractPermissionControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPermissionController.Contract.ContractPermissionControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPermissionController *ContractPermissionControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.ContractPermissionControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPermissionController *ContractPermissionControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.ContractPermissionControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractPermissionController *ContractPermissionControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractPermissionController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractPermissionController *ContractPermissionControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractPermissionController *ContractPermissionControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerCaller) CanCall(opts *bind.CallOpts, account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "canCall", account, caller, target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _ContractPermissionController.Contract.CanCall(&_ContractPermissionController.CallOpts, account, caller, target, selector)
}

// CanCall is a free data retrieval call binding the contract method 0xdf595cb8.
//
// Solidity: function canCall(address account, address caller, address target, bytes4 selector) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerCallerSession) CanCall(account common.Address, caller common.Address, target common.Address, selector [4]byte) (bool, error) {
	return _ContractPermissionController.Contract.CanCall(&_ContractPermissionController.CallOpts, account, caller, target, selector)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerCaller) GetAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "getAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _ContractPermissionController.Contract.GetAdmins(&_ContractPermissionController.CallOpts, account)
}

// GetAdmins is a free data retrieval call binding the contract method 0xad5f2210.
//
// Solidity: function getAdmins(address account) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerCallerSession) GetAdmins(account common.Address) ([]common.Address, error) {
	return _ContractPermissionController.Contract.GetAdmins(&_ContractPermissionController.CallOpts, account)
}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_ContractPermissionController *ContractPermissionControllerCaller) GetAppointeePermissions(opts *bind.CallOpts, account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "getAppointeePermissions", account, appointee)

	if err != nil {
		return *new([]common.Address), *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][4]byte)).(*[][4]byte)

	return out0, out1, err

}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_ContractPermissionController *ContractPermissionControllerSession) GetAppointeePermissions(account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	return _ContractPermissionController.Contract.GetAppointeePermissions(&_ContractPermissionController.CallOpts, account, appointee)
}

// GetAppointeePermissions is a free data retrieval call binding the contract method 0x882a3b38.
//
// Solidity: function getAppointeePermissions(address account, address appointee) view returns(address[], bytes4[])
func (_ContractPermissionController *ContractPermissionControllerCallerSession) GetAppointeePermissions(account common.Address, appointee common.Address) ([]common.Address, [][4]byte, error) {
	return _ContractPermissionController.Contract.GetAppointeePermissions(&_ContractPermissionController.CallOpts, account, appointee)
}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerCaller) GetAppointees(opts *bind.CallOpts, account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "getAppointees", account, target, selector)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	return _ContractPermissionController.Contract.GetAppointees(&_ContractPermissionController.CallOpts, account, target, selector)
}

// GetAppointees is a free data retrieval call binding the contract method 0xfddbdefd.
//
// Solidity: function getAppointees(address account, address target, bytes4 selector) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerCallerSession) GetAppointees(account common.Address, target common.Address, selector [4]byte) ([]common.Address, error) {
	return _ContractPermissionController.Contract.GetAppointees(&_ContractPermissionController.CallOpts, account, target, selector)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerCaller) GetPendingAdmins(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "getPendingAdmins", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _ContractPermissionController.Contract.GetPendingAdmins(&_ContractPermissionController.CallOpts, account)
}

// GetPendingAdmins is a free data retrieval call binding the contract method 0x6bddfa1f.
//
// Solidity: function getPendingAdmins(address account) view returns(address[])
func (_ContractPermissionController *ContractPermissionControllerCallerSession) GetPendingAdmins(account common.Address) ([]common.Address, error) {
	return _ContractPermissionController.Contract.GetPendingAdmins(&_ContractPermissionController.CallOpts, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerCaller) IsAdmin(opts *bind.CallOpts, account common.Address, caller common.Address) (bool, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "isAdmin", account, caller)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _ContractPermissionController.Contract.IsAdmin(&_ContractPermissionController.CallOpts, account, caller)
}

// IsAdmin is a free data retrieval call binding the contract method 0x91006745.
//
// Solidity: function isAdmin(address account, address caller) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerCallerSession) IsAdmin(account common.Address, caller common.Address) (bool, error) {
	return _ContractPermissionController.Contract.IsAdmin(&_ContractPermissionController.CallOpts, account, caller)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerCaller) IsPendingAdmin(opts *bind.CallOpts, account common.Address, pendingAdmin common.Address) (bool, error) {
	var out []interface{}
	err := _ContractPermissionController.contract.Call(opts, &out, "isPendingAdmin", account, pendingAdmin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _ContractPermissionController.Contract.IsPendingAdmin(&_ContractPermissionController.CallOpts, account, pendingAdmin)
}

// IsPendingAdmin is a free data retrieval call binding the contract method 0xad8aca77.
//
// Solidity: function isPendingAdmin(address account, address pendingAdmin) view returns(bool)
func (_ContractPermissionController *ContractPermissionControllerCallerSession) IsPendingAdmin(account common.Address, pendingAdmin common.Address) (bool, error) {
	return _ContractPermissionController.Contract.IsPendingAdmin(&_ContractPermissionController.CallOpts, account, pendingAdmin)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactor) AcceptAdmin(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.contract.Transact(opts, "acceptAdmin", account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_ContractPermissionController *ContractPermissionControllerSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.AcceptAdmin(&_ContractPermissionController.TransactOpts, account)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address account) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactorSession) AcceptAdmin(account common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.AcceptAdmin(&_ContractPermissionController.TransactOpts, account)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactor) AddPendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.contract.Transact(opts, "addPendingAdmin", account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.AddPendingAdmin(&_ContractPermissionController.TransactOpts, account, admin)
}

// AddPendingAdmin is a paid mutator transaction binding the contract method 0xeb5a4e87.
//
// Solidity: function addPendingAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactorSession) AddPendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.AddPendingAdmin(&_ContractPermissionController.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactor) RemoveAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.contract.Transact(opts, "removeAdmin", account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.RemoveAdmin(&_ContractPermissionController.TransactOpts, account, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactorSession) RemoveAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.RemoveAdmin(&_ContractPermissionController.TransactOpts, account, admin)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactor) RemoveAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractPermissionController.contract.Transact(opts, "removeAppointee", account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_ContractPermissionController *ContractPermissionControllerSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.RemoveAppointee(&_ContractPermissionController.TransactOpts, account, appointee, target, selector)
}

// RemoveAppointee is a paid mutator transaction binding the contract method 0x06641201.
//
// Solidity: function removeAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactorSession) RemoveAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.RemoveAppointee(&_ContractPermissionController.TransactOpts, account, appointee, target, selector)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactor) RemovePendingAdmin(opts *bind.TransactOpts, account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.contract.Transact(opts, "removePendingAdmin", account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.RemovePendingAdmin(&_ContractPermissionController.TransactOpts, account, admin)
}

// RemovePendingAdmin is a paid mutator transaction binding the contract method 0x4f906cf9.
//
// Solidity: function removePendingAdmin(address account, address admin) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactorSession) RemovePendingAdmin(account common.Address, admin common.Address) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.RemovePendingAdmin(&_ContractPermissionController.TransactOpts, account, admin)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactor) SetAppointee(opts *bind.TransactOpts, account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractPermissionController.contract.Transact(opts, "setAppointee", account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_ContractPermissionController *ContractPermissionControllerSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.SetAppointee(&_ContractPermissionController.TransactOpts, account, appointee, target, selector)
}

// SetAppointee is a paid mutator transaction binding the contract method 0x950d806e.
//
// Solidity: function setAppointee(address account, address appointee, address target, bytes4 selector) returns()
func (_ContractPermissionController *ContractPermissionControllerTransactorSession) SetAppointee(account common.Address, appointee common.Address, target common.Address, selector [4]byte) (*types.Transaction, error) {
	return _ContractPermissionController.Contract.SetAppointee(&_ContractPermissionController.TransactOpts, account, appointee, target, selector)
}

// ContractPermissionControllerAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the ContractPermissionController contract.
type ContractPermissionControllerAdminRemovedIterator struct {
	Event *ContractPermissionControllerAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerAdminRemoved represents a AdminRemoved event raised by the ContractPermissionController contract.
type ContractPermissionControllerAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerAdminRemovedIterator{contract: _ContractPermissionController.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "AdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerAdminRemoved)
				if err := _ContractPermissionController.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminRemoved is a log parse operation binding the contract event 0xdb9d5d31320daf5bc7181d565b6da4d12e30f0f4d5aa324a992426c14a1d19ce.
//
// Solidity: event AdminRemoved(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParseAdminRemoved(log types.Log) (*ContractPermissionControllerAdminRemoved, error) {
	event := new(ContractPermissionControllerAdminRemoved)
	if err := _ContractPermissionController.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPermissionControllerAdminSetIterator is returned from FilterAdminSet and is used to iterate over the raw logs and unpacked data for AdminSet events raised by the ContractPermissionController contract.
type ContractPermissionControllerAdminSetIterator struct {
	Event *ContractPermissionControllerAdminSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerAdminSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerAdminSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerAdminSet represents a AdminSet event raised by the ContractPermissionController contract.
type ContractPermissionControllerAdminSet struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminSet is a free log retrieval operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterAdminSet(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerAdminSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerAdminSetIterator{contract: _ContractPermissionController.contract, event: "AdminSet", logs: logs, sub: sub}, nil
}

// WatchAdminSet is a free log subscription operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchAdminSet(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAdminSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "AdminSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerAdminSet)
				if err := _ContractPermissionController.contract.UnpackLog(event, "AdminSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminSet is a log parse operation binding the contract event 0xbf265e8326285a2747e33e54d5945f7111f2b5edb826eb8c08d4677779b3ff97.
//
// Solidity: event AdminSet(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParseAdminSet(log types.Log) (*ContractPermissionControllerAdminSet, error) {
	event := new(ContractPermissionControllerAdminSet)
	if err := _ContractPermissionController.contract.UnpackLog(event, "AdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPermissionControllerAppointeeRemovedIterator is returned from FilterAppointeeRemoved and is used to iterate over the raw logs and unpacked data for AppointeeRemoved events raised by the ContractPermissionController contract.
type ContractPermissionControllerAppointeeRemovedIterator struct {
	Event *ContractPermissionControllerAppointeeRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerAppointeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerAppointeeRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerAppointeeRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerAppointeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerAppointeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerAppointeeRemoved represents a AppointeeRemoved event raised by the ContractPermissionController contract.
type ContractPermissionControllerAppointeeRemoved struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeRemoved is a free log retrieval operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterAppointeeRemoved(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*ContractPermissionControllerAppointeeRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerAppointeeRemovedIterator{contract: _ContractPermissionController.contract, event: "AppointeeRemoved", logs: logs, sub: sub}, nil
}

// WatchAppointeeRemoved is a free log subscription operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchAppointeeRemoved(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAppointeeRemoved, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "AppointeeRemoved", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerAppointeeRemoved)
				if err := _ContractPermissionController.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppointeeRemoved is a log parse operation binding the contract event 0x18242326b6b862126970679759169f01f646bd55ec5bfcab85ba9f337a74e0c6.
//
// Solidity: event AppointeeRemoved(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParseAppointeeRemoved(log types.Log) (*ContractPermissionControllerAppointeeRemoved, error) {
	event := new(ContractPermissionControllerAppointeeRemoved)
	if err := _ContractPermissionController.contract.UnpackLog(event, "AppointeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPermissionControllerAppointeeSetIterator is returned from FilterAppointeeSet and is used to iterate over the raw logs and unpacked data for AppointeeSet events raised by the ContractPermissionController contract.
type ContractPermissionControllerAppointeeSetIterator struct {
	Event *ContractPermissionControllerAppointeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerAppointeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerAppointeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerAppointeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerAppointeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerAppointeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerAppointeeSet represents a AppointeeSet event raised by the ContractPermissionController contract.
type ContractPermissionControllerAppointeeSet struct {
	Account   common.Address
	Appointee common.Address
	Target    common.Address
	Selector  [4]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppointeeSet is a free log retrieval operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterAppointeeSet(opts *bind.FilterOpts, account []common.Address, appointee []common.Address) (*ContractPermissionControllerAppointeeSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerAppointeeSetIterator{contract: _ContractPermissionController.contract, event: "AppointeeSet", logs: logs, sub: sub}, nil
}

// WatchAppointeeSet is a free log subscription operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchAppointeeSet(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerAppointeeSet, account []common.Address, appointee []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var appointeeRule []interface{}
	for _, appointeeItem := range appointee {
		appointeeRule = append(appointeeRule, appointeeItem)
	}

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "AppointeeSet", accountRule, appointeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerAppointeeSet)
				if err := _ContractPermissionController.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppointeeSet is a log parse operation binding the contract event 0x037f03a2ad6b967df4a01779b6d2b4c85950df83925d9e31362b519422fc0169.
//
// Solidity: event AppointeeSet(address indexed account, address indexed appointee, address target, bytes4 selector)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParseAppointeeSet(log types.Log) (*ContractPermissionControllerAppointeeSet, error) {
	event := new(ContractPermissionControllerAppointeeSet)
	if err := _ContractPermissionController.contract.UnpackLog(event, "AppointeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPermissionControllerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractPermissionController contract.
type ContractPermissionControllerInitializedIterator struct {
	Event *ContractPermissionControllerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerInitialized represents a Initialized event raised by the ContractPermissionController contract.
type ContractPermissionControllerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractPermissionControllerInitializedIterator, error) {

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerInitializedIterator{contract: _ContractPermissionController.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerInitialized)
				if err := _ContractPermissionController.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParseInitialized(log types.Log) (*ContractPermissionControllerInitialized, error) {
	event := new(ContractPermissionControllerInitialized)
	if err := _ContractPermissionController.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPermissionControllerPendingAdminAddedIterator is returned from FilterPendingAdminAdded and is used to iterate over the raw logs and unpacked data for PendingAdminAdded events raised by the ContractPermissionController contract.
type ContractPermissionControllerPendingAdminAddedIterator struct {
	Event *ContractPermissionControllerPendingAdminAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerPendingAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerPendingAdminAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerPendingAdminAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerPendingAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerPendingAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerPendingAdminAdded represents a PendingAdminAdded event raised by the ContractPermissionController contract.
type ContractPermissionControllerPendingAdminAdded struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminAdded is a free log retrieval operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterPendingAdminAdded(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerPendingAdminAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerPendingAdminAddedIterator{contract: _ContractPermissionController.contract, event: "PendingAdminAdded", logs: logs, sub: sub}, nil
}

// WatchPendingAdminAdded is a free log subscription operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchPendingAdminAdded(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerPendingAdminAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "PendingAdminAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerPendingAdminAdded)
				if err := _ContractPermissionController.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingAdminAdded is a log parse operation binding the contract event 0xb14b9a3d448c5b04f0e5b087b6f5193390db7955482a6ffb841e7b3ba61a460c.
//
// Solidity: event PendingAdminAdded(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParsePendingAdminAdded(log types.Log) (*ContractPermissionControllerPendingAdminAdded, error) {
	event := new(ContractPermissionControllerPendingAdminAdded)
	if err := _ContractPermissionController.contract.UnpackLog(event, "PendingAdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPermissionControllerPendingAdminRemovedIterator is returned from FilterPendingAdminRemoved and is used to iterate over the raw logs and unpacked data for PendingAdminRemoved events raised by the ContractPermissionController contract.
type ContractPermissionControllerPendingAdminRemovedIterator struct {
	Event *ContractPermissionControllerPendingAdminRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPermissionControllerPendingAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPermissionControllerPendingAdminRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPermissionControllerPendingAdminRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPermissionControllerPendingAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPermissionControllerPendingAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPermissionControllerPendingAdminRemoved represents a PendingAdminRemoved event raised by the ContractPermissionController contract.
type ContractPermissionControllerPendingAdminRemoved struct {
	Account common.Address
	Admin   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPendingAdminRemoved is a free log retrieval operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) FilterPendingAdminRemoved(opts *bind.FilterOpts, account []common.Address) (*ContractPermissionControllerPendingAdminRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.FilterLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractPermissionControllerPendingAdminRemovedIterator{contract: _ContractPermissionController.contract, event: "PendingAdminRemoved", logs: logs, sub: sub}, nil
}

// WatchPendingAdminRemoved is a free log subscription operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) WatchPendingAdminRemoved(opts *bind.WatchOpts, sink chan<- *ContractPermissionControllerPendingAdminRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractPermissionController.contract.WatchLogs(opts, "PendingAdminRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPermissionControllerPendingAdminRemoved)
				if err := _ContractPermissionController.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePendingAdminRemoved is a log parse operation binding the contract event 0xd706ed7ae044d795b49e54c9f519f663053951011985f663a862cd9ee72a9ac7.
//
// Solidity: event PendingAdminRemoved(address indexed account, address admin)
func (_ContractPermissionController *ContractPermissionControllerFilterer) ParsePendingAdminRemoved(log types.Log) (*ContractPermissionControllerPendingAdminRemoved, error) {
	event := new(ContractPermissionControllerPendingAdminRemoved)
	if err := _ContractPermissionController.contract.UnpackLog(event, "PendingAdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
