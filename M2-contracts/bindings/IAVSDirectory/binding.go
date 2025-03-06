// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIAVSDirectory

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractIAVSDirectoryMetaData contains all meta data concerning the ContractIAVSDirectory contract.
var ContractIAVSDirectoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"OPERATOR_AVS_REGISTRATION_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateOperatorAVSRegistrationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorSaltIsSpent\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVSMetadataURI\",\"inputs\":[{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AVSMetadataURIUpdated\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAVSRegistrationStatusUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIAVSDirectory.OperatorAVSRegistrationStatus\"}],\"anonymous\":false}]",
}

// ContractIAVSDirectoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIAVSDirectoryMetaData.ABI instead.
var ContractIAVSDirectoryABI = ContractIAVSDirectoryMetaData.ABI

// ContractIAVSDirectoryMethods is an auto generated interface around an Ethereum contract.
type ContractIAVSDirectoryMethods interface {
	ContractIAVSDirectoryCalls
	ContractIAVSDirectoryTransacts
	ContractIAVSDirectoryFilters
}

// ContractIAVSDirectoryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIAVSDirectoryCalls interface {
	OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error)

	CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error)
}

// ContractIAVSDirectoryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIAVSDirectoryTransacts interface {
	CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error)

	DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error)

	RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error)
}

// ContractIAVSDirectoryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIAVSDirectoryFilters interface {
	FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMetadataURIUpdatedIterator, error)
	WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error)
	ParseAVSMetadataURIUpdated(log types.Log) (*ContractIAVSDirectoryAVSMetadataURIUpdated, error)

	FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error)
	WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error)
	ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, error)
}

// ContractIAVSDirectory is an auto generated Go binding around an Ethereum contract.
type ContractIAVSDirectory struct {
	ContractIAVSDirectoryCaller     // Read-only binding to the contract
	ContractIAVSDirectoryTransactor // Write-only binding to the contract
	ContractIAVSDirectoryFilterer   // Log filterer for contract events
}

// ContractIAVSDirectory implements the ContractIAVSDirectoryMethods interface.
var _ ContractIAVSDirectoryMethods = (*ContractIAVSDirectory)(nil)

// ContractIAVSDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAVSDirectoryCaller implements the ContractIAVSDirectoryCalls interface.
var _ ContractIAVSDirectoryCalls = (*ContractIAVSDirectoryCaller)(nil)

// ContractIAVSDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAVSDirectoryTransactor implements the ContractIAVSDirectoryTransacts interface.
var _ ContractIAVSDirectoryTransacts = (*ContractIAVSDirectoryTransactor)(nil)

// ContractIAVSDirectoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIAVSDirectoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAVSDirectoryFilterer implements the ContractIAVSDirectoryFilters interface.
var _ ContractIAVSDirectoryFilters = (*ContractIAVSDirectoryFilterer)(nil)

// ContractIAVSDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIAVSDirectorySession struct {
	Contract     *ContractIAVSDirectory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractIAVSDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIAVSDirectoryCallerSession struct {
	Contract *ContractIAVSDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractIAVSDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIAVSDirectoryTransactorSession struct {
	Contract     *ContractIAVSDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractIAVSDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIAVSDirectoryRaw struct {
	Contract *ContractIAVSDirectory // Generic contract binding to access the raw methods on
}

// ContractIAVSDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryCallerRaw struct {
	Contract *ContractIAVSDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIAVSDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIAVSDirectoryTransactorRaw struct {
	Contract *ContractIAVSDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIAVSDirectory creates a new instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectory(address common.Address, backend bind.ContractBackend) (*ContractIAVSDirectory, error) {
	contract, err := bindContractIAVSDirectory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectory{ContractIAVSDirectoryCaller: ContractIAVSDirectoryCaller{contract: contract}, ContractIAVSDirectoryTransactor: ContractIAVSDirectoryTransactor{contract: contract}, ContractIAVSDirectoryFilterer: ContractIAVSDirectoryFilterer{contract: contract}}, nil
}

// NewContractIAVSDirectoryCaller creates a new read-only instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectoryCaller(address common.Address, caller bind.ContractCaller) (*ContractIAVSDirectoryCaller, error) {
	contract, err := bindContractIAVSDirectory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryCaller{contract: contract}, nil
}

// NewContractIAVSDirectoryTransactor creates a new write-only instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIAVSDirectoryTransactor, error) {
	contract, err := bindContractIAVSDirectory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryTransactor{contract: contract}, nil
}

// NewContractIAVSDirectoryFilterer creates a new log filterer instance of ContractIAVSDirectory, bound to a specific deployed contract.
func NewContractIAVSDirectoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIAVSDirectoryFilterer, error) {
	contract, err := bindContractIAVSDirectory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryFilterer{contract: contract}, nil
}

// bindContractIAVSDirectory binds a generic wrapper to an already deployed contract.
func bindContractIAVSDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIAVSDirectoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIAVSDirectory *ContractIAVSDirectoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIAVSDirectory.Contract.ContractIAVSDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIAVSDirectory *ContractIAVSDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ContractIAVSDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIAVSDirectory *ContractIAVSDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.ContractIAVSDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIAVSDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.contract.Transact(opts, method, params...)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OPERATORAVSREGISTRATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "OPERATOR_AVS_REGISTRATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractIAVSDirectory.CallOpts)
}

// OPERATORAVSREGISTRATIONTYPEHASH is a free data retrieval call binding the contract method 0xd79aceab.
//
// Solidity: function OPERATOR_AVS_REGISTRATION_TYPEHASH() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OPERATORAVSREGISTRATIONTYPEHASH() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.OPERATORAVSREGISTRATIONTYPEHASH(&_ContractIAVSDirectory.CallOpts)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) CalculateOperatorAVSRegistrationDigestHash(opts *bind.CallOpts, operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "calculateOperatorAVSRegistrationDigestHash", operator, avs, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractIAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// CalculateOperatorAVSRegistrationDigestHash is a free data retrieval call binding the contract method 0xa1060c88.
//
// Solidity: function calculateOperatorAVSRegistrationDigestHash(address operator, address avs, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) CalculateOperatorAVSRegistrationDigestHash(operator common.Address, avs common.Address, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.CalculateOperatorAVSRegistrationDigestHash(&_ContractIAVSDirectory.CallOpts, operator, avs, salt, expiry)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) DomainSeparator() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.DomainSeparator(&_ContractIAVSDirectory.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractIAVSDirectory.Contract.DomainSeparator(&_ContractIAVSDirectory.CallOpts)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCaller) OperatorSaltIsSpent(opts *bind.CallOpts, operator common.Address, salt [32]byte) (bool, error) {
	var out []interface{}
	err := _ContractIAVSDirectory.contract.Call(opts, &out, "operatorSaltIsSpent", operator, salt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractIAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractIAVSDirectory.CallOpts, operator, salt)
}

// OperatorSaltIsSpent is a free data retrieval call binding the contract method 0x374823b5.
//
// Solidity: function operatorSaltIsSpent(address operator, bytes32 salt) view returns(bool)
func (_ContractIAVSDirectory *ContractIAVSDirectoryCallerSession) OperatorSaltIsSpent(operator common.Address, salt [32]byte) (bool, error) {
	return _ContractIAVSDirectory.Contract.OperatorSaltIsSpent(&_ContractIAVSDirectory.CallOpts, operator, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.CancelSalt(&_ContractIAVSDirectory.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.CancelSalt(&_ContractIAVSDirectory.TransactOpts, salt)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractIAVSDirectory.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.DeregisterOperatorFromAVS(&_ContractIAVSDirectory.TransactOpts, operator)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractIAVSDirectory.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.RegisterOperatorToAVS(&_ContractIAVSDirectory.TransactOpts, operator, operatorSignature)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactor) UpdateAVSMetadataURI(opts *bind.TransactOpts, metadataURI string) (*types.Transaction, error) {
	return _ContractIAVSDirectory.contract.Transact(opts, "updateAVSMetadataURI", metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectorySession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractIAVSDirectory.TransactOpts, metadataURI)
}

// UpdateAVSMetadataURI is a paid mutator transaction binding the contract method 0xa98fb355.
//
// Solidity: function updateAVSMetadataURI(string metadataURI) returns()
func (_ContractIAVSDirectory *ContractIAVSDirectoryTransactorSession) UpdateAVSMetadataURI(metadataURI string) (*types.Transaction, error) {
	return _ContractIAVSDirectory.Contract.UpdateAVSMetadataURI(&_ContractIAVSDirectory.TransactOpts, metadataURI)
}

// ContractIAVSDirectoryAVSMetadataURIUpdatedIterator is returned from FilterAVSMetadataURIUpdated and is used to iterate over the raw logs and unpacked data for AVSMetadataURIUpdated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryAVSMetadataURIUpdatedIterator struct {
	Event *ContractIAVSDirectoryAVSMetadataURIUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryAVSMetadataURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryAVSMetadataURIUpdated)
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
		it.Event = new(ContractIAVSDirectoryAVSMetadataURIUpdated)
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
func (it *ContractIAVSDirectoryAVSMetadataURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryAVSMetadataURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryAVSMetadataURIUpdated represents a AVSMetadataURIUpdated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryAVSMetadataURIUpdated struct {
	Avs         common.Address
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAVSMetadataURIUpdated is a free log retrieval operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterAVSMetadataURIUpdated(opts *bind.FilterOpts, avs []common.Address) (*ContractIAVSDirectoryAVSMetadataURIUpdatedIterator, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryAVSMetadataURIUpdatedIterator{contract: _ContractIAVSDirectory.contract, event: "AVSMetadataURIUpdated", logs: logs, sub: sub}, nil
}

// WatchAVSMetadataURIUpdated is a free log subscription operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchAVSMetadataURIUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryAVSMetadataURIUpdated, avs []common.Address) (event.Subscription, error) {

	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "AVSMetadataURIUpdated", avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryAVSMetadataURIUpdated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
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

// ParseAVSMetadataURIUpdated is a log parse operation binding the contract event 0xa89c1dc243d8908a96dd84944bcc97d6bc6ac00dd78e20621576be6a3c943713.
//
// Solidity: event AVSMetadataURIUpdated(address indexed avs, string metadataURI)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseAVSMetadataURIUpdated(log types.Log) (*ContractIAVSDirectoryAVSMetadataURIUpdated, error) {
	event := new(ContractIAVSDirectoryAVSMetadataURIUpdated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "AVSMetadataURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator is returned from FilterOperatorAVSRegistrationStatusUpdated and is used to iterate over the raw logs and unpacked data for OperatorAVSRegistrationStatusUpdated events raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator struct {
	Event *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
		it.Event = new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
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
func (it *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated represents a OperatorAVSRegistrationStatusUpdated event raised by the ContractIAVSDirectory contract.
type ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated struct {
	Operator common.Address
	Avs      common.Address
	Status   uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorAVSRegistrationStatusUpdated is a free log retrieval operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) FilterOperatorAVSRegistrationStatusUpdated(opts *bind.FilterOpts, operator []common.Address, avs []common.Address) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.FilterLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return &ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdatedIterator{contract: _ContractIAVSDirectory.contract, event: "OperatorAVSRegistrationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorAVSRegistrationStatusUpdated is a free log subscription operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) WatchOperatorAVSRegistrationStatusUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, operator []common.Address, avs []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsRule []interface{}
	for _, avsItem := range avs {
		avsRule = append(avsRule, avsItem)
	}

	logs, sub, err := _ContractIAVSDirectory.contract.WatchLogs(opts, "OperatorAVSRegistrationStatusUpdated", operatorRule, avsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
				if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
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

// ParseOperatorAVSRegistrationStatusUpdated is a log parse operation binding the contract event 0xf0952b1c65271d819d39983d2abb044b9cace59bcc4d4dd389f586ebdcb15b41.
//
// Solidity: event OperatorAVSRegistrationStatusUpdated(address indexed operator, address indexed avs, uint8 status)
func (_ContractIAVSDirectory *ContractIAVSDirectoryFilterer) ParseOperatorAVSRegistrationStatusUpdated(log types.Log) (*ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated, error) {
	event := new(ContractIAVSDirectoryOperatorAVSRegistrationStatusUpdated)
	if err := _ContractIAVSDirectory.contract.UnpackLog(event, "OperatorAVSRegistrationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
