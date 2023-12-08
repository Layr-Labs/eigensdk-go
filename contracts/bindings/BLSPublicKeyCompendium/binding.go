// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractBLSPublicKeyCompendium

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// ContractBLSPublicKeyCompendiumMetaData contains all meta data concerning the ContractBLSPublicKeyCompendium contract.
var ContractBLSPublicKeyCompendiumMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"name\":\"NewPubkeyRegistration\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getMessageHash\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorToPubkeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pubkeyHashToOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"signedMessageHash\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkeyG1\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"pubkeyG2\",\"type\":\"tuple\"}],\"name\":\"registerBLSPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6",
}

// ContractBLSPublicKeyCompendiumABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBLSPublicKeyCompendiumMetaData.ABI instead.
var ContractBLSPublicKeyCompendiumABI = ContractBLSPublicKeyCompendiumMetaData.ABI

// ContractBLSPublicKeyCompendiumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBLSPublicKeyCompendiumMetaData.Bin instead.
var ContractBLSPublicKeyCompendiumBin = ContractBLSPublicKeyCompendiumMetaData.Bin

// DeployContractBLSPublicKeyCompendium deploys a new Ethereum contract, binding an instance of ContractBLSPublicKeyCompendium to it.
func DeployContractBLSPublicKeyCompendium(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractBLSPublicKeyCompendium, error) {
	parsed, err := ContractBLSPublicKeyCompendiumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBLSPublicKeyCompendiumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBLSPublicKeyCompendium{ContractBLSPublicKeyCompendiumCaller: ContractBLSPublicKeyCompendiumCaller{contract: contract}, ContractBLSPublicKeyCompendiumTransactor: ContractBLSPublicKeyCompendiumTransactor{contract: contract}, ContractBLSPublicKeyCompendiumFilterer: ContractBLSPublicKeyCompendiumFilterer{contract: contract}}, nil
}

// ContractBLSPublicKeyCompendiumMethods is an auto generated interface around an Ethereum contract.
type ContractBLSPublicKeyCompendiumMethods interface {
	ContractBLSPublicKeyCompendiumCalls
	ContractBLSPublicKeyCompendiumTransacts
	ContractBLSPublicKeyCompendiumFilters
}

// ContractBLSPublicKeyCompendiumCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractBLSPublicKeyCompendiumCalls interface {
	GetMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error)

	OperatorToPubkeyHash(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error)

	PubkeyHashToOperator(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error)
}

// ContractBLSPublicKeyCompendiumTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractBLSPublicKeyCompendiumTransacts interface {
	RegisterBLSPublicKey(opts *bind.TransactOpts, signedMessageHash BN254G1Point, pubkeyG1 BN254G1Point, pubkeyG2 BN254G2Point) (*types.Transaction, error)
}

// ContractBLSPublicKeyCompendiumFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractBLSPublicKeyCompendiumFilters interface {
	FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator, error)
	WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSPublicKeyCompendiumNewPubkeyRegistration, operator []common.Address) (event.Subscription, error)
	ParseNewPubkeyRegistration(log types.Log) (*ContractBLSPublicKeyCompendiumNewPubkeyRegistration, error)
}

// ContractBLSPublicKeyCompendium is an auto generated Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendium struct {
	ContractBLSPublicKeyCompendiumCaller     // Read-only binding to the contract
	ContractBLSPublicKeyCompendiumTransactor // Write-only binding to the contract
	ContractBLSPublicKeyCompendiumFilterer   // Log filterer for contract events
}

// ContractBLSPublicKeyCompendium implements the ContractBLSPublicKeyCompendiumMethods interface.
var _ ContractBLSPublicKeyCompendiumMethods = (*ContractBLSPublicKeyCompendium)(nil)

// ContractBLSPublicKeyCompendiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPublicKeyCompendiumCaller implements the ContractBLSPublicKeyCompendiumCalls interface.
var _ ContractBLSPublicKeyCompendiumCalls = (*ContractBLSPublicKeyCompendiumCaller)(nil)

// ContractBLSPublicKeyCompendiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPublicKeyCompendiumTransactor implements the ContractBLSPublicKeyCompendiumTransacts interface.
var _ ContractBLSPublicKeyCompendiumTransacts = (*ContractBLSPublicKeyCompendiumTransactor)(nil)

// ContractBLSPublicKeyCompendiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBLSPublicKeyCompendiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPublicKeyCompendiumFilterer implements the ContractBLSPublicKeyCompendiumFilters interface.
var _ ContractBLSPublicKeyCompendiumFilters = (*ContractBLSPublicKeyCompendiumFilterer)(nil)

// ContractBLSPublicKeyCompendiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBLSPublicKeyCompendiumSession struct {
	Contract     *ContractBLSPublicKeyCompendium // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractBLSPublicKeyCompendiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBLSPublicKeyCompendiumCallerSession struct {
	Contract *ContractBLSPublicKeyCompendiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractBLSPublicKeyCompendiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBLSPublicKeyCompendiumTransactorSession struct {
	Contract     *ContractBLSPublicKeyCompendiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractBLSPublicKeyCompendiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumRaw struct {
	Contract *ContractBLSPublicKeyCompendium // Generic contract binding to access the raw methods on
}

// ContractBLSPublicKeyCompendiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumCallerRaw struct {
	Contract *ContractBLSPublicKeyCompendiumCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBLSPublicKeyCompendiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumTransactorRaw struct {
	Contract *ContractBLSPublicKeyCompendiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBLSPublicKeyCompendium creates a new instance of ContractBLSPublicKeyCompendium, bound to a specific deployed contract.
func NewContractBLSPublicKeyCompendium(address common.Address, backend bind.ContractBackend) (*ContractBLSPublicKeyCompendium, error) {
	contract, err := bindContractBLSPublicKeyCompendium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPublicKeyCompendium{ContractBLSPublicKeyCompendiumCaller: ContractBLSPublicKeyCompendiumCaller{contract: contract}, ContractBLSPublicKeyCompendiumTransactor: ContractBLSPublicKeyCompendiumTransactor{contract: contract}, ContractBLSPublicKeyCompendiumFilterer: ContractBLSPublicKeyCompendiumFilterer{contract: contract}}, nil
}

// NewContractBLSPublicKeyCompendiumCaller creates a new read-only instance of ContractBLSPublicKeyCompendium, bound to a specific deployed contract.
func NewContractBLSPublicKeyCompendiumCaller(address common.Address, caller bind.ContractCaller) (*ContractBLSPublicKeyCompendiumCaller, error) {
	contract, err := bindContractBLSPublicKeyCompendium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPublicKeyCompendiumCaller{contract: contract}, nil
}

// NewContractBLSPublicKeyCompendiumTransactor creates a new write-only instance of ContractBLSPublicKeyCompendium, bound to a specific deployed contract.
func NewContractBLSPublicKeyCompendiumTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBLSPublicKeyCompendiumTransactor, error) {
	contract, err := bindContractBLSPublicKeyCompendium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPublicKeyCompendiumTransactor{contract: contract}, nil
}

// NewContractBLSPublicKeyCompendiumFilterer creates a new log filterer instance of ContractBLSPublicKeyCompendium, bound to a specific deployed contract.
func NewContractBLSPublicKeyCompendiumFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBLSPublicKeyCompendiumFilterer, error) {
	contract, err := bindContractBLSPublicKeyCompendium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPublicKeyCompendiumFilterer{contract: contract}, nil
}

// bindContractBLSPublicKeyCompendium binds a generic wrapper to an already deployed contract.
func bindContractBLSPublicKeyCompendium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBLSPublicKeyCompendiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSPublicKeyCompendium.Contract.ContractBLSPublicKeyCompendiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.Contract.ContractBLSPublicKeyCompendiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.Contract.ContractBLSPublicKeyCompendiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSPublicKeyCompendium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.Contract.contract.Transact(opts, method, params...)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x1f5ac1b2.
//
// Solidity: function getMessageHash(address operator) view returns((uint256,uint256))
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCaller) GetMessageHash(opts *bind.CallOpts, operator common.Address) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractBLSPublicKeyCompendium.contract.Call(opts, &out, "getMessageHash", operator)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetMessageHash is a free data retrieval call binding the contract method 0x1f5ac1b2.
//
// Solidity: function getMessageHash(address operator) view returns((uint256,uint256))
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession) GetMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractBLSPublicKeyCompendium.Contract.GetMessageHash(&_ContractBLSPublicKeyCompendium.CallOpts, operator)
}

// GetMessageHash is a free data retrieval call binding the contract method 0x1f5ac1b2.
//
// Solidity: function getMessageHash(address operator) view returns((uint256,uint256))
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerSession) GetMessageHash(operator common.Address) (BN254G1Point, error) {
	return _ContractBLSPublicKeyCompendium.Contract.GetMessageHash(&_ContractBLSPublicKeyCompendium.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCaller) OperatorToPubkeyHash(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSPublicKeyCompendium.contract.Call(opts, &out, "operatorToPubkeyHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _ContractBLSPublicKeyCompendium.Contract.OperatorToPubkeyHash(&_ContractBLSPublicKeyCompendium.CallOpts, arg0)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address ) view returns(bytes32)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerSession) OperatorToPubkeyHash(arg0 common.Address) ([32]byte, error) {
	return _ContractBLSPublicKeyCompendium.Contract.OperatorToPubkeyHash(&_ContractBLSPublicKeyCompendium.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCaller) PubkeyHashToOperator(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSPublicKeyCompendium.contract.Call(opts, &out, "pubkeyHashToOperator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _ContractBLSPublicKeyCompendium.Contract.PubkeyHashToOperator(&_ContractBLSPublicKeyCompendium.CallOpts, arg0)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 ) view returns(address)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumCallerSession) PubkeyHashToOperator(arg0 [32]byte) (common.Address, error) {
	return _ContractBLSPublicKeyCompendium.Contract.PubkeyHashToOperator(&_ContractBLSPublicKeyCompendium.CallOpts, arg0)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x161a334d.
//
// Solidity: function registerBLSPublicKey((uint256,uint256) signedMessageHash, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2) returns()
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, signedMessageHash BN254G1Point, pubkeyG1 BN254G1Point, pubkeyG2 BN254G2Point) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.contract.Transact(opts, "registerBLSPublicKey", signedMessageHash, pubkeyG1, pubkeyG2)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x161a334d.
//
// Solidity: function registerBLSPublicKey((uint256,uint256) signedMessageHash, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2) returns()
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumSession) RegisterBLSPublicKey(signedMessageHash BN254G1Point, pubkeyG1 BN254G1Point, pubkeyG2 BN254G2Point) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.Contract.RegisterBLSPublicKey(&_ContractBLSPublicKeyCompendium.TransactOpts, signedMessageHash, pubkeyG1, pubkeyG2)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0x161a334d.
//
// Solidity: function registerBLSPublicKey((uint256,uint256) signedMessageHash, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2) returns()
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumTransactorSession) RegisterBLSPublicKey(signedMessageHash BN254G1Point, pubkeyG1 BN254G1Point, pubkeyG2 BN254G2Point) (*types.Transaction, error) {
	return _ContractBLSPublicKeyCompendium.Contract.RegisterBLSPublicKey(&_ContractBLSPublicKeyCompendium.TransactOpts, signedMessageHash, pubkeyG1, pubkeyG2)
}

// ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the ContractBLSPublicKeyCompendium contract.
type ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator struct {
	Event *ContractBLSPublicKeyCompendiumNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
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
		it.Event = new(ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
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
func (it *ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSPublicKeyCompendiumNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the ContractBLSPublicKeyCompendium contract.
type ContractBLSPublicKeyCompendiumNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSPublicKeyCompendium.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPublicKeyCompendiumNewPubkeyRegistrationIterator{contract: _ContractBLSPublicKeyCompendium.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSPublicKeyCompendiumNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSPublicKeyCompendium.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
				if err := _ContractBLSPublicKeyCompendium.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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

// ParseNewPubkeyRegistration is a log parse operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSPublicKeyCompendium *ContractBLSPublicKeyCompendiumFilterer) ParseNewPubkeyRegistration(log types.Log) (*ContractBLSPublicKeyCompendiumNewPubkeyRegistration, error) {
	event := new(ContractBLSPublicKeyCompendiumNewPubkeyRegistration)
	if err := _ContractBLSPublicKeyCompendium.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
