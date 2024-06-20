// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractContractsRegistry

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

// ContractContractsRegistryMetaData contains all meta data concerning the ContractContractsRegistry contract.
var ContractContractsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"contractCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractNames\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contracts\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerContract\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_contract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561001057600080fd5b506105f6806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633ca6bb92146100515780637f3c2c281461007a5780638736381a1461008f5780638c5b8385146100a6575b600080fd5b61006461005f366004610284565b6100f2565b60405161007191906102c1565b60405180910390f35b61008d610088366004610397565b61018c565b005b61009860025481565b604051908152602001610071565b6100da6100b43660046103f5565b80516020818301810180516000825292820191909301209152546001600160a01b031681565b6040516001600160a01b039091168152602001610071565b6001602052600090815260409020805461010b90610432565b80601f016020809104026020016040519081016040528092919081815260200182805461013790610432565b80156101845780601f1061015957610100808354040283529160200191610184565b820191906000526020600020905b81548152906001019060200180831161016757829003601f168201915b505050505081565b60006001600160a01b03166000836040516101a7919061046c565b908152604051908190036020019020546001600160a01b0316146102115760405162461bcd60e51b815260206004820152601b60248201527f636f6e747261637420616c726561647920726567697374657265640000000000604482015260640160405180910390fd5b80600083604051610222919061046c565b908152604080516020928190038301902080546001600160a01b0319166001600160a01b039490941693909317909255600254600090815260019091522061026a83826104d9565b506002805490600061027b83610599565b91905055505050565b60006020828403121561029657600080fd5b5035919050565b60005b838110156102b85781810151838201526020016102a0565b50506000910152565b60208152600082518060208401526102e081604085016020870161029d565b601f01601f19169190910160400192915050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261031b57600080fd5b813567ffffffffffffffff80821115610336576103366102f4565b604051601f8301601f19908116603f0116810190828211818310171561035e5761035e6102f4565b8160405283815286602085880101111561037757600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156103aa57600080fd5b823567ffffffffffffffff8111156103c157600080fd5b6103cd8582860161030a565b92505060208301356001600160a01b03811681146103ea57600080fd5b809150509250929050565b60006020828403121561040757600080fd5b813567ffffffffffffffff81111561041e57600080fd5b61042a8482850161030a565b949350505050565b600181811c9082168061044657607f821691505b60208210810361046657634e487b7160e01b600052602260045260246000fd5b50919050565b6000825161047e81846020870161029d565b9190910192915050565b601f8211156104d4576000816000526020600020601f850160051c810160208610156104b15750805b601f850160051c820191505b818110156104d0578281556001016104bd565b5050505b505050565b815167ffffffffffffffff8111156104f3576104f36102f4565b610507816105018454610432565b84610488565b602080601f83116001811461053c57600084156105245750858301515b600019600386901b1c1916600185901b1785556104d0565b600085815260208120601f198616915b8281101561056b5788860151825594840194600190910190840161054c565b50858210156105895787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000600182016105b957634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212201129bf50a6beb2fbed289c6e1c02d8506dcb138717d42c3d9cf57ad9ebfde11364736f6c63430008170033",
}

// ContractContractsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractContractsRegistryMetaData.ABI instead.
var ContractContractsRegistryABI = ContractContractsRegistryMetaData.ABI

// ContractContractsRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractContractsRegistryMetaData.Bin instead.
var ContractContractsRegistryBin = ContractContractsRegistryMetaData.Bin

// DeployContractContractsRegistry deploys a new Ethereum contract, binding an instance of ContractContractsRegistry to it.
func DeployContractContractsRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractContractsRegistry, error) {
	parsed, err := ContractContractsRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractContractsRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractContractsRegistry{ContractContractsRegistryCaller: ContractContractsRegistryCaller{contract: contract}, ContractContractsRegistryTransactor: ContractContractsRegistryTransactor{contract: contract}, ContractContractsRegistryFilterer: ContractContractsRegistryFilterer{contract: contract}}, nil
}

// ContractContractsRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractContractsRegistryMethods interface {
	ContractContractsRegistryCalls
	ContractContractsRegistryTransacts
	ContractContractsRegistryFilters
}

// ContractContractsRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractContractsRegistryCalls interface {
	ContractCount(opts *bind.CallOpts) (*big.Int, error)

	ContractNames(opts *bind.CallOpts, arg0 *big.Int) (string, error)

	Contracts(opts *bind.CallOpts, arg0 string) (common.Address, error)
}

// ContractContractsRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractContractsRegistryTransacts interface {
	RegisterContract(opts *bind.TransactOpts, name string, _contract common.Address) (*types.Transaction, error)
}

// ContractContractsRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractContractsRegistryFilters interface {
}

// ContractContractsRegistry is an auto generated Go binding around an Ethereum contract.
type ContractContractsRegistry struct {
	ContractContractsRegistryCaller     // Read-only binding to the contract
	ContractContractsRegistryTransactor // Write-only binding to the contract
	ContractContractsRegistryFilterer   // Log filterer for contract events
}

// ContractContractsRegistry implements the ContractContractsRegistryMethods interface.
var _ ContractContractsRegistryMethods = (*ContractContractsRegistry)(nil)

// ContractContractsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractContractsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractContractsRegistryCaller implements the ContractContractsRegistryCalls interface.
var _ ContractContractsRegistryCalls = (*ContractContractsRegistryCaller)(nil)

// ContractContractsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractContractsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractContractsRegistryTransactor implements the ContractContractsRegistryTransacts interface.
var _ ContractContractsRegistryTransacts = (*ContractContractsRegistryTransactor)(nil)

// ContractContractsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractContractsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractContractsRegistryFilterer implements the ContractContractsRegistryFilters interface.
var _ ContractContractsRegistryFilters = (*ContractContractsRegistryFilterer)(nil)

// ContractContractsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractContractsRegistrySession struct {
	Contract     *ContractContractsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractContractsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractContractsRegistryCallerSession struct {
	Contract *ContractContractsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractContractsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractContractsRegistryTransactorSession struct {
	Contract     *ContractContractsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractContractsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractContractsRegistryRaw struct {
	Contract *ContractContractsRegistry // Generic contract binding to access the raw methods on
}

// ContractContractsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractContractsRegistryCallerRaw struct {
	Contract *ContractContractsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractContractsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractContractsRegistryTransactorRaw struct {
	Contract *ContractContractsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractContractsRegistry creates a new instance of ContractContractsRegistry, bound to a specific deployed contract.
func NewContractContractsRegistry(address common.Address, backend bind.ContractBackend) (*ContractContractsRegistry, error) {
	contract, err := bindContractContractsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractContractsRegistry{ContractContractsRegistryCaller: ContractContractsRegistryCaller{contract: contract}, ContractContractsRegistryTransactor: ContractContractsRegistryTransactor{contract: contract}, ContractContractsRegistryFilterer: ContractContractsRegistryFilterer{contract: contract}}, nil
}

// NewContractContractsRegistryCaller creates a new read-only instance of ContractContractsRegistry, bound to a specific deployed contract.
func NewContractContractsRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractContractsRegistryCaller, error) {
	contract, err := bindContractContractsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractContractsRegistryCaller{contract: contract}, nil
}

// NewContractContractsRegistryTransactor creates a new write-only instance of ContractContractsRegistry, bound to a specific deployed contract.
func NewContractContractsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractContractsRegistryTransactor, error) {
	contract, err := bindContractContractsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractContractsRegistryTransactor{contract: contract}, nil
}

// NewContractContractsRegistryFilterer creates a new log filterer instance of ContractContractsRegistry, bound to a specific deployed contract.
func NewContractContractsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractContractsRegistryFilterer, error) {
	contract, err := bindContractContractsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractContractsRegistryFilterer{contract: contract}, nil
}

// bindContractContractsRegistry binds a generic wrapper to an already deployed contract.
func bindContractContractsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractContractsRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractContractsRegistry *ContractContractsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractContractsRegistry.Contract.ContractContractsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractContractsRegistry *ContractContractsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractContractsRegistry.Contract.ContractContractsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractContractsRegistry *ContractContractsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractContractsRegistry.Contract.ContractContractsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractContractsRegistry *ContractContractsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractContractsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractContractsRegistry *ContractContractsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractContractsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractContractsRegistry *ContractContractsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractContractsRegistry.Contract.contract.Transact(opts, method, params...)
}

// ContractCount is a free data retrieval call binding the contract method 0x8736381a.
//
// Solidity: function contractCount() view returns(uint256)
func (_ContractContractsRegistry *ContractContractsRegistryCaller) ContractCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractContractsRegistry.contract.Call(opts, &out, "contractCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractCount is a free data retrieval call binding the contract method 0x8736381a.
//
// Solidity: function contractCount() view returns(uint256)
func (_ContractContractsRegistry *ContractContractsRegistrySession) ContractCount() (*big.Int, error) {
	return _ContractContractsRegistry.Contract.ContractCount(&_ContractContractsRegistry.CallOpts)
}

// ContractCount is a free data retrieval call binding the contract method 0x8736381a.
//
// Solidity: function contractCount() view returns(uint256)
func (_ContractContractsRegistry *ContractContractsRegistryCallerSession) ContractCount() (*big.Int, error) {
	return _ContractContractsRegistry.Contract.ContractCount(&_ContractContractsRegistry.CallOpts)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_ContractContractsRegistry *ContractContractsRegistryCaller) ContractNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _ContractContractsRegistry.contract.Call(opts, &out, "contractNames", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_ContractContractsRegistry *ContractContractsRegistrySession) ContractNames(arg0 *big.Int) (string, error) {
	return _ContractContractsRegistry.Contract.ContractNames(&_ContractContractsRegistry.CallOpts, arg0)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_ContractContractsRegistry *ContractContractsRegistryCallerSession) ContractNames(arg0 *big.Int) (string, error) {
	return _ContractContractsRegistry.Contract.ContractNames(&_ContractContractsRegistry.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x8c5b8385.
//
// Solidity: function contracts(string ) view returns(address)
func (_ContractContractsRegistry *ContractContractsRegistryCaller) Contracts(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _ContractContractsRegistry.contract.Call(opts, &out, "contracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Contracts is a free data retrieval call binding the contract method 0x8c5b8385.
//
// Solidity: function contracts(string ) view returns(address)
func (_ContractContractsRegistry *ContractContractsRegistrySession) Contracts(arg0 string) (common.Address, error) {
	return _ContractContractsRegistry.Contract.Contracts(&_ContractContractsRegistry.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x8c5b8385.
//
// Solidity: function contracts(string ) view returns(address)
func (_ContractContractsRegistry *ContractContractsRegistryCallerSession) Contracts(arg0 string) (common.Address, error) {
	return _ContractContractsRegistry.Contract.Contracts(&_ContractContractsRegistry.CallOpts, arg0)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x7f3c2c28.
//
// Solidity: function registerContract(string name, address _contract) returns()
func (_ContractContractsRegistry *ContractContractsRegistryTransactor) RegisterContract(opts *bind.TransactOpts, name string, _contract common.Address) (*types.Transaction, error) {
	return _ContractContractsRegistry.contract.Transact(opts, "registerContract", name, _contract)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x7f3c2c28.
//
// Solidity: function registerContract(string name, address _contract) returns()
func (_ContractContractsRegistry *ContractContractsRegistrySession) RegisterContract(name string, _contract common.Address) (*types.Transaction, error) {
	return _ContractContractsRegistry.Contract.RegisterContract(&_ContractContractsRegistry.TransactOpts, name, _contract)
}

// RegisterContract is a paid mutator transaction binding the contract method 0x7f3c2c28.
//
// Solidity: function registerContract(string name, address _contract) returns()
func (_ContractContractsRegistry *ContractContractsRegistryTransactorSession) RegisterContract(name string, _contract common.Address) (*types.Transaction, error) {
	return _ContractContractsRegistry.Contract.RegisterContract(&_ContractContractsRegistry.TransactOpts, name, _contract)
}
