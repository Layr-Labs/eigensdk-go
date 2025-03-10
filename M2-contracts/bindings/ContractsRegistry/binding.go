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
	Bin: "0x608080604052346015576104c1908161001a8239f35b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c9081633ca6bb92146102f8575080637f3c2c28146100ca5780638736381a146100ad57638c5b838514610048575f80fd5b346100a95760203660031901126100a95760043567ffffffffffffffff81116100a95761007b6020913690600401610435565b8160405191805191829101835e5f90820190815281900382019020546040516001600160a01b039091168152f35b5f80fd5b346100a9575f3660031901126100a9576020600254604051908152f35b346100a95760403660031901126100a95760043567ffffffffffffffff81116100a9576100fb903690600401610435565b6024356001600160a01b038116908190036100a95760405182519060208401918083835e5f9082019081528190036020019020546001600160a01b03166102b3576020604051809285518091835e81015f815203019020906bffffffffffffffffffffffff60a01b8254161790556002545f52600160205260405f20815167ffffffffffffffff811161029f5761019282546103db565b601f811161025a575b50602092601f82116001146101fb57928192935f926101f0575b50508160011b915f199060031b1c19161790555b6002545f1981146101dc57600101600255005b634e487b7160e01b5f52601160045260245ffd5b0151905083806101b5565b601f19821693835f52805f20915f5b868110610242575083600195961061022a575b505050811b0190556101c9565b01515f1960f88460031b161c1916905583808061021d565b9192602060018192868501518155019401920161020a565b825f5260205f20601f830160051c81019160208410610295575b601f0160051c01905b81811061028a575061019b565b5f815560010161027d565b9091508190610274565b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152601b60248201527f636f6e747261637420616c7265616479207265676973746572656400000000006044820152606490fd5b346100a95760203660031901126100a9576004355f52600160205260405f20905f825492610325846103db565b9081845260208401946001811690815f146103be575060011461037e575b8460408561035381870382610413565b8151928391602083525180918160208501528484015e5f828201840152601f01601f19168101030190f35b5f90815260208120939250905b8082106103a45750909150810160200161035382610343565b91926001816020925483858801015201910190929161038b565b60ff191686525050151560051b8201602001905061035382610343565b90600182811c92168015610409575b60208310146103f557565b634e487b7160e01b5f52602260045260245ffd5b91607f16916103ea565b90601f8019910116810190811067ffffffffffffffff82111761029f57604052565b81601f820112156100a95780359067ffffffffffffffff821161029f576040519261046a601f8401601f191660200185610413565b828452602083830101116100a957815f92602080930183860137830101529056fea264697066735822122039bd1953a252e76f7a405125ede13adb1de2eba416c1ab9e500173de92189b3c64736f6c634300081b0033",
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
