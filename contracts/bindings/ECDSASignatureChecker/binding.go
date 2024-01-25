// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractECDSASignatureChecker

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

// ECDSASignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type ECDSASignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// ECDSASignatureCheckerSignerStakeIndicesAndSignatures is an auto generated low-level Go binding around an user-defined struct.
type ECDSASignatureCheckerSignerStakeIndicesAndSignatures struct {
	SignerIds                 []common.Address
	Signatures                [][]byte
	SignerQuorumBitmapIndices []uint32
	TotalStakeIndices         []uint32
	SignerStakeIndices        [][]uint32
}

// ContractECDSASignatureCheckerMetaData contains all meta data concerning the ContractECDSASignatureChecker contract.
var ContractECDSASignatureCheckerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractECDSARegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"checkSignatures\",\"inputs\":[{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structECDSASignatureChecker.SignerStakeIndicesAndSignatures\",\"components\":[{\"name\":\"signerIds\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signatures\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"signerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"signerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structECDSASignatureChecker.QuorumStakeTotals\",\"components\":[{\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"},{\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\",\"internalType\":\"uint96[]\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractECDSARegistryCoordinator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setStaleStakesForbidden\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractECDSAStakeRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staleStakesForbidden\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"StaleStakesForbiddenUpdate\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false}]",
	Bin: "0x6",
}

// ContractECDSASignatureCheckerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractECDSASignatureCheckerMetaData.ABI instead.
var ContractECDSASignatureCheckerABI = ContractECDSASignatureCheckerMetaData.ABI

// ContractECDSASignatureCheckerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractECDSASignatureCheckerMetaData.Bin instead.
var ContractECDSASignatureCheckerBin = ContractECDSASignatureCheckerMetaData.Bin

// DeployContractECDSASignatureChecker deploys a new Ethereum contract, binding an instance of ContractECDSASignatureChecker to it.
func DeployContractECDSASignatureChecker(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address) (common.Address, *types.Transaction, *ContractECDSASignatureChecker, error) {
	parsed, err := ContractECDSASignatureCheckerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractECDSASignatureCheckerBin), backend, _registryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractECDSASignatureChecker{ContractECDSASignatureCheckerCaller: ContractECDSASignatureCheckerCaller{contract: contract}, ContractECDSASignatureCheckerTransactor: ContractECDSASignatureCheckerTransactor{contract: contract}, ContractECDSASignatureCheckerFilterer: ContractECDSASignatureCheckerFilterer{contract: contract}}, nil
}

// ContractECDSASignatureCheckerMethods is an auto generated interface around an Ethereum contract.
type ContractECDSASignatureCheckerMethods interface {
	ContractECDSASignatureCheckerCalls
	ContractECDSASignatureCheckerTransacts
	ContractECDSASignatureCheckerFilters
}

// ContractECDSASignatureCheckerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractECDSASignatureCheckerCalls interface {
	CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params ECDSASignatureCheckerSignerStakeIndicesAndSignatures) (ECDSASignatureCheckerQuorumStakeTotals, [32]byte, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)

	StakeRegistry(opts *bind.CallOpts) (common.Address, error)

	StaleStakesForbidden(opts *bind.CallOpts) (bool, error)
}

// ContractECDSASignatureCheckerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractECDSASignatureCheckerTransacts interface {
	SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error)
}

// ContractECDSASignatureCheckerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractECDSASignatureCheckerFilters interface {
	FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator, error)
	WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractECDSASignatureCheckerStaleStakesForbiddenUpdate) (event.Subscription, error)
	ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractECDSASignatureCheckerStaleStakesForbiddenUpdate, error)
}

// ContractECDSASignatureChecker is an auto generated Go binding around an Ethereum contract.
type ContractECDSASignatureChecker struct {
	ContractECDSASignatureCheckerCaller     // Read-only binding to the contract
	ContractECDSASignatureCheckerTransactor // Write-only binding to the contract
	ContractECDSASignatureCheckerFilterer   // Log filterer for contract events
}

// ContractECDSASignatureChecker implements the ContractECDSASignatureCheckerMethods interface.
var _ ContractECDSASignatureCheckerMethods = (*ContractECDSASignatureChecker)(nil)

// ContractECDSASignatureCheckerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractECDSASignatureCheckerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractECDSASignatureCheckerCaller implements the ContractECDSASignatureCheckerCalls interface.
var _ ContractECDSASignatureCheckerCalls = (*ContractECDSASignatureCheckerCaller)(nil)

// ContractECDSASignatureCheckerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractECDSASignatureCheckerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractECDSASignatureCheckerTransactor implements the ContractECDSASignatureCheckerTransacts interface.
var _ ContractECDSASignatureCheckerTransacts = (*ContractECDSASignatureCheckerTransactor)(nil)

// ContractECDSASignatureCheckerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractECDSASignatureCheckerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractECDSASignatureCheckerFilterer implements the ContractECDSASignatureCheckerFilters interface.
var _ ContractECDSASignatureCheckerFilters = (*ContractECDSASignatureCheckerFilterer)(nil)

// ContractECDSASignatureCheckerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractECDSASignatureCheckerSession struct {
	Contract     *ContractECDSASignatureChecker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ContractECDSASignatureCheckerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractECDSASignatureCheckerCallerSession struct {
	Contract *ContractECDSASignatureCheckerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ContractECDSASignatureCheckerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractECDSASignatureCheckerTransactorSession struct {
	Contract     *ContractECDSASignatureCheckerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ContractECDSASignatureCheckerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractECDSASignatureCheckerRaw struct {
	Contract *ContractECDSASignatureChecker // Generic contract binding to access the raw methods on
}

// ContractECDSASignatureCheckerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractECDSASignatureCheckerCallerRaw struct {
	Contract *ContractECDSASignatureCheckerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractECDSASignatureCheckerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractECDSASignatureCheckerTransactorRaw struct {
	Contract *ContractECDSASignatureCheckerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractECDSASignatureChecker creates a new instance of ContractECDSASignatureChecker, bound to a specific deployed contract.
func NewContractECDSASignatureChecker(address common.Address, backend bind.ContractBackend) (*ContractECDSASignatureChecker, error) {
	contract, err := bindContractECDSASignatureChecker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractECDSASignatureChecker{ContractECDSASignatureCheckerCaller: ContractECDSASignatureCheckerCaller{contract: contract}, ContractECDSASignatureCheckerTransactor: ContractECDSASignatureCheckerTransactor{contract: contract}, ContractECDSASignatureCheckerFilterer: ContractECDSASignatureCheckerFilterer{contract: contract}}, nil
}

// NewContractECDSASignatureCheckerCaller creates a new read-only instance of ContractECDSASignatureChecker, bound to a specific deployed contract.
func NewContractECDSASignatureCheckerCaller(address common.Address, caller bind.ContractCaller) (*ContractECDSASignatureCheckerCaller, error) {
	contract, err := bindContractECDSASignatureChecker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractECDSASignatureCheckerCaller{contract: contract}, nil
}

// NewContractECDSASignatureCheckerTransactor creates a new write-only instance of ContractECDSASignatureChecker, bound to a specific deployed contract.
func NewContractECDSASignatureCheckerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractECDSASignatureCheckerTransactor, error) {
	contract, err := bindContractECDSASignatureChecker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractECDSASignatureCheckerTransactor{contract: contract}, nil
}

// NewContractECDSASignatureCheckerFilterer creates a new log filterer instance of ContractECDSASignatureChecker, bound to a specific deployed contract.
func NewContractECDSASignatureCheckerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractECDSASignatureCheckerFilterer, error) {
	contract, err := bindContractECDSASignatureChecker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractECDSASignatureCheckerFilterer{contract: contract}, nil
}

// bindContractECDSASignatureChecker binds a generic wrapper to an already deployed contract.
func bindContractECDSASignatureChecker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractECDSASignatureCheckerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractECDSASignatureChecker.Contract.ContractECDSASignatureCheckerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.Contract.ContractECDSASignatureCheckerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.Contract.ContractECDSASignatureCheckerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractECDSASignatureChecker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.Contract.contract.Transact(opts, method, params...)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x5f62d642.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (address[],bytes[],uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params ECDSASignatureCheckerSignerStakeIndicesAndSignatures) (ECDSASignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractECDSASignatureChecker.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(ECDSASignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(ECDSASignatureCheckerQuorumStakeTotals)).(*ECDSASignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x5f62d642.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (address[],bytes[],uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params ECDSASignatureCheckerSignerStakeIndicesAndSignatures) (ECDSASignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractECDSASignatureChecker.Contract.CheckSignatures(&_ContractECDSASignatureChecker.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x5f62d642.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (address[],bytes[],uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params ECDSASignatureCheckerSignerStakeIndicesAndSignatures) (ECDSASignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractECDSASignatureChecker.Contract.CheckSignatures(&_ContractECDSASignatureChecker.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractECDSASignatureChecker.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerSession) Delegation() (common.Address, error) {
	return _ContractECDSASignatureChecker.Contract.Delegation(&_ContractECDSASignatureChecker.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCallerSession) Delegation() (common.Address, error) {
	return _ContractECDSASignatureChecker.Contract.Delegation(&_ContractECDSASignatureChecker.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractECDSASignatureChecker.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractECDSASignatureChecker.Contract.RegistryCoordinator(&_ContractECDSASignatureChecker.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractECDSASignatureChecker.Contract.RegistryCoordinator(&_ContractECDSASignatureChecker.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractECDSASignatureChecker.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerSession) StakeRegistry() (common.Address, error) {
	return _ContractECDSASignatureChecker.Contract.StakeRegistry(&_ContractECDSASignatureChecker.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractECDSASignatureChecker.Contract.StakeRegistry(&_ContractECDSASignatureChecker.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractECDSASignatureChecker.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerSession) StaleStakesForbidden() (bool, error) {
	return _ContractECDSASignatureChecker.Contract.StaleStakesForbidden(&_ContractECDSASignatureChecker.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractECDSASignatureChecker.Contract.StaleStakesForbidden(&_ContractECDSASignatureChecker.CallOpts)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.Contract.SetStaleStakesForbidden(&_ContractECDSASignatureChecker.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractECDSASignatureChecker.Contract.SetStaleStakesForbidden(&_ContractECDSASignatureChecker.TransactOpts, value)
}

// ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractECDSASignatureChecker contract.
type ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractECDSASignatureCheckerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractECDSASignatureCheckerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractECDSASignatureCheckerStaleStakesForbiddenUpdate)
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
func (it *ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractECDSASignatureCheckerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractECDSASignatureChecker contract.
type ContractECDSASignatureCheckerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractECDSASignatureChecker.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractECDSASignatureCheckerStaleStakesForbiddenUpdateIterator{contract: _ContractECDSASignatureChecker.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractECDSASignatureCheckerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractECDSASignatureChecker.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractECDSASignatureCheckerStaleStakesForbiddenUpdate)
				if err := _ContractECDSASignatureChecker.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractECDSASignatureChecker *ContractECDSASignatureCheckerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractECDSASignatureCheckerStaleStakesForbiddenUpdate, error) {
	event := new(ContractECDSASignatureCheckerStaleStakesForbiddenUpdate)
	if err := _ContractECDSASignatureChecker.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
