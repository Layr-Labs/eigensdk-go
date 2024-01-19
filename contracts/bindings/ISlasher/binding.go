// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractISlasher

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

// ISlasherMiddlewareTimes is an auto generated low-level Go binding around an user-defined struct.
type ISlasherMiddlewareTimes struct {
	StalestUpdateBlock    uint32
	LatestServeUntilBlock uint32
}

// ContractISlasherMetaData contains all meta data concerning the ContractISlasher contract.
var ContractISlasherMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previouslySlashedAddress\",\"type\":\"address\"}],\"name\":\"FrozenStatusReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalestUpdateBlock\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"latestServeUntilBlock\",\"type\":\"uint32\"}],\"name\":\"MiddlewareTimesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slashedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"slashingContract\",\"type\":\"address\"}],\"name\":\"OperatorFrozen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"OptedIntoSlashing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"contractCanSlashOperatorUntilBlock\",\"type\":\"uint32\"}],\"name\":\"SlashingAbilityRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toBeSlashed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashingContract\",\"type\":\"address\"}],\"name\":\"canSlash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"withdrawalStartBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"middlewareTimesIndex\",\"type\":\"uint256\"}],\"name\":\"canWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"serviceContract\",\"type\":\"address\"}],\"name\":\"contractCanSlashOperatorUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toBeFrozen\",\"type\":\"address\"}],\"name\":\"freezeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"updateBlock\",\"type\":\"uint32\"}],\"name\":\"getCorrectValueForInsertAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getMiddlewareTimesIndexServeUntilBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"getMiddlewareTimesIndexStalestUpdateBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"serviceContract\",\"type\":\"address\"}],\"name\":\"latestUpdateBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"middlewareTimesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"arrayIndex\",\"type\":\"uint256\"}],\"name\":\"operatorToMiddlewareTimes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"stalestUpdateBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"latestServeUntilBlock\",\"type\":\"uint32\"}],\"internalType\":\"structISlasher.MiddlewareTimes\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"operatorWhitelistedContractsLinkedListEntry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"operatorWhitelistedContractsLinkedListSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"optIntoSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"recordFirstStakeUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"}],\"name\":\"recordLastStakeUpdateAndRevokeSlashingAbility\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"updateBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"serveUntilBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"insertAfter\",\"type\":\"uint256\"}],\"name\":\"recordStakeUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"frozenAddresses\",\"type\":\"address[]\"}],\"name\":\"resetFrozenStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractISlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractISlasherMetaData.ABI instead.
var ContractISlasherABI = ContractISlasherMetaData.ABI

// ContractISlasherMethods is an auto generated interface around an Ethereum contract.
type ContractISlasherMethods interface {
	ContractISlasherCalls
	ContractISlasherTransacts
	ContractISlasherFilters
}

// ContractISlasherCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractISlasherCalls interface {
	CanSlash(opts *bind.CallOpts, toBeSlashed common.Address, slashingContract common.Address) (bool, error)

	ContractCanSlashOperatorUntilBlock(opts *bind.CallOpts, operator common.Address, serviceContract common.Address) (uint32, error)

	Delegation(opts *bind.CallOpts) (common.Address, error)

	GetCorrectValueForInsertAfter(opts *bind.CallOpts, operator common.Address, updateBlock uint32) (*big.Int, error)

	GetMiddlewareTimesIndexServeUntilBlock(opts *bind.CallOpts, operator common.Address, index uint32) (uint32, error)

	GetMiddlewareTimesIndexStalestUpdateBlock(opts *bind.CallOpts, operator common.Address, index uint32) (uint32, error)

	IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error)

	LatestUpdateBlock(opts *bind.CallOpts, operator common.Address, serviceContract common.Address) (uint32, error)

	MiddlewareTimesLength(opts *bind.CallOpts, operator common.Address) (*big.Int, error)

	OperatorToMiddlewareTimes(opts *bind.CallOpts, operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error)

	OperatorWhitelistedContractsLinkedListEntry(opts *bind.CallOpts, operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error)

	OperatorWhitelistedContractsLinkedListSize(opts *bind.CallOpts, operator common.Address) (*big.Int, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)
}

// ContractISlasherTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractISlasherTransacts interface {
	CanWithdraw(opts *bind.TransactOpts, operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error)

	FreezeOperator(opts *bind.TransactOpts, toBeFrozen common.Address) (*types.Transaction, error)

	OptIntoSlashing(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error)

	RecordFirstStakeUpdate(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error)

	RecordLastStakeUpdateAndRevokeSlashingAbility(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error)

	RecordStakeUpdate(opts *bind.TransactOpts, operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error)

	ResetFrozenStatus(opts *bind.TransactOpts, frozenAddresses []common.Address) (*types.Transaction, error)
}

// ContractISlasherFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractISlasherFilters interface {
	FilterFrozenStatusReset(opts *bind.FilterOpts, previouslySlashedAddress []common.Address) (*ContractISlasherFrozenStatusResetIterator, error)
	WatchFrozenStatusReset(opts *bind.WatchOpts, sink chan<- *ContractISlasherFrozenStatusReset, previouslySlashedAddress []common.Address) (event.Subscription, error)
	ParseFrozenStatusReset(log types.Log) (*ContractISlasherFrozenStatusReset, error)

	FilterMiddlewareTimesAdded(opts *bind.FilterOpts) (*ContractISlasherMiddlewareTimesAddedIterator, error)
	WatchMiddlewareTimesAdded(opts *bind.WatchOpts, sink chan<- *ContractISlasherMiddlewareTimesAdded) (event.Subscription, error)
	ParseMiddlewareTimesAdded(log types.Log) (*ContractISlasherMiddlewareTimesAdded, error)

	FilterOperatorFrozen(opts *bind.FilterOpts, slashedOperator []common.Address, slashingContract []common.Address) (*ContractISlasherOperatorFrozenIterator, error)
	WatchOperatorFrozen(opts *bind.WatchOpts, sink chan<- *ContractISlasherOperatorFrozen, slashedOperator []common.Address, slashingContract []common.Address) (event.Subscription, error)
	ParseOperatorFrozen(log types.Log) (*ContractISlasherOperatorFrozen, error)

	FilterOptedIntoSlashing(opts *bind.FilterOpts, operator []common.Address, contractAddress []common.Address) (*ContractISlasherOptedIntoSlashingIterator, error)
	WatchOptedIntoSlashing(opts *bind.WatchOpts, sink chan<- *ContractISlasherOptedIntoSlashing, operator []common.Address, contractAddress []common.Address) (event.Subscription, error)
	ParseOptedIntoSlashing(log types.Log) (*ContractISlasherOptedIntoSlashing, error)

	FilterSlashingAbilityRevoked(opts *bind.FilterOpts, operator []common.Address, contractAddress []common.Address) (*ContractISlasherSlashingAbilityRevokedIterator, error)
	WatchSlashingAbilityRevoked(opts *bind.WatchOpts, sink chan<- *ContractISlasherSlashingAbilityRevoked, operator []common.Address, contractAddress []common.Address) (event.Subscription, error)
	ParseSlashingAbilityRevoked(log types.Log) (*ContractISlasherSlashingAbilityRevoked, error)
}

// ContractISlasher is an auto generated Go binding around an Ethereum contract.
type ContractISlasher struct {
	ContractISlasherCaller     // Read-only binding to the contract
	ContractISlasherTransactor // Write-only binding to the contract
	ContractISlasherFilterer   // Log filterer for contract events
}

// ContractISlasher implements the ContractISlasherMethods interface.
var _ ContractISlasherMethods = (*ContractISlasher)(nil)

// ContractISlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractISlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractISlasherCaller implements the ContractISlasherCalls interface.
var _ ContractISlasherCalls = (*ContractISlasherCaller)(nil)

// ContractISlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractISlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractISlasherTransactor implements the ContractISlasherTransacts interface.
var _ ContractISlasherTransacts = (*ContractISlasherTransactor)(nil)

// ContractISlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractISlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractISlasherFilterer implements the ContractISlasherFilters interface.
var _ ContractISlasherFilters = (*ContractISlasherFilterer)(nil)

// ContractISlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractISlasherSession struct {
	Contract     *ContractISlasher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractISlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractISlasherCallerSession struct {
	Contract *ContractISlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractISlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractISlasherTransactorSession struct {
	Contract     *ContractISlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractISlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractISlasherRaw struct {
	Contract *ContractISlasher // Generic contract binding to access the raw methods on
}

// ContractISlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractISlasherCallerRaw struct {
	Contract *ContractISlasherCaller // Generic read-only contract binding to access the raw methods on
}

// ContractISlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractISlasherTransactorRaw struct {
	Contract *ContractISlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractISlasher creates a new instance of ContractISlasher, bound to a specific deployed contract.
func NewContractISlasher(address common.Address, backend bind.ContractBackend) (*ContractISlasher, error) {
	contract, err := bindContractISlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractISlasher{ContractISlasherCaller: ContractISlasherCaller{contract: contract}, ContractISlasherTransactor: ContractISlasherTransactor{contract: contract}, ContractISlasherFilterer: ContractISlasherFilterer{contract: contract}}, nil
}

// NewContractISlasherCaller creates a new read-only instance of ContractISlasher, bound to a specific deployed contract.
func NewContractISlasherCaller(address common.Address, caller bind.ContractCaller) (*ContractISlasherCaller, error) {
	contract, err := bindContractISlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherCaller{contract: contract}, nil
}

// NewContractISlasherTransactor creates a new write-only instance of ContractISlasher, bound to a specific deployed contract.
func NewContractISlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractISlasherTransactor, error) {
	contract, err := bindContractISlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherTransactor{contract: contract}, nil
}

// NewContractISlasherFilterer creates a new log filterer instance of ContractISlasher, bound to a specific deployed contract.
func NewContractISlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractISlasherFilterer, error) {
	contract, err := bindContractISlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherFilterer{contract: contract}, nil
}

// bindContractISlasher binds a generic wrapper to an already deployed contract.
func bindContractISlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractISlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractISlasher *ContractISlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractISlasher.Contract.ContractISlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractISlasher *ContractISlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractISlasher.Contract.ContractISlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractISlasher *ContractISlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractISlasher.Contract.ContractISlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractISlasher *ContractISlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractISlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractISlasher *ContractISlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractISlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractISlasher *ContractISlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractISlasher.Contract.contract.Transact(opts, method, params...)
}

// CanSlash is a free data retrieval call binding the contract method 0xd98128c0.
//
// Solidity: function canSlash(address toBeSlashed, address slashingContract) view returns(bool)
func (_ContractISlasher *ContractISlasherCaller) CanSlash(opts *bind.CallOpts, toBeSlashed common.Address, slashingContract common.Address) (bool, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "canSlash", toBeSlashed, slashingContract)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanSlash is a free data retrieval call binding the contract method 0xd98128c0.
//
// Solidity: function canSlash(address toBeSlashed, address slashingContract) view returns(bool)
func (_ContractISlasher *ContractISlasherSession) CanSlash(toBeSlashed common.Address, slashingContract common.Address) (bool, error) {
	return _ContractISlasher.Contract.CanSlash(&_ContractISlasher.CallOpts, toBeSlashed, slashingContract)
}

// CanSlash is a free data retrieval call binding the contract method 0xd98128c0.
//
// Solidity: function canSlash(address toBeSlashed, address slashingContract) view returns(bool)
func (_ContractISlasher *ContractISlasherCallerSession) CanSlash(toBeSlashed common.Address, slashingContract common.Address) (bool, error) {
	return _ContractISlasher.Contract.CanSlash(&_ContractISlasher.CallOpts, toBeSlashed, slashingContract)
}

// ContractCanSlashOperatorUntilBlock is a free data retrieval call binding the contract method 0x6f0c2f74.
//
// Solidity: function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) view returns(uint32)
func (_ContractISlasher *ContractISlasherCaller) ContractCanSlashOperatorUntilBlock(opts *bind.CallOpts, operator common.Address, serviceContract common.Address) (uint32, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "contractCanSlashOperatorUntilBlock", operator, serviceContract)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ContractCanSlashOperatorUntilBlock is a free data retrieval call binding the contract method 0x6f0c2f74.
//
// Solidity: function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) view returns(uint32)
func (_ContractISlasher *ContractISlasherSession) ContractCanSlashOperatorUntilBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ContractISlasher.Contract.ContractCanSlashOperatorUntilBlock(&_ContractISlasher.CallOpts, operator, serviceContract)
}

// ContractCanSlashOperatorUntilBlock is a free data retrieval call binding the contract method 0x6f0c2f74.
//
// Solidity: function contractCanSlashOperatorUntilBlock(address operator, address serviceContract) view returns(uint32)
func (_ContractISlasher *ContractISlasherCallerSession) ContractCanSlashOperatorUntilBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ContractISlasher.Contract.ContractCanSlashOperatorUntilBlock(&_ContractISlasher.CallOpts, operator, serviceContract)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractISlasher *ContractISlasherCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractISlasher *ContractISlasherSession) Delegation() (common.Address, error) {
	return _ContractISlasher.Contract.Delegation(&_ContractISlasher.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractISlasher *ContractISlasherCallerSession) Delegation() (common.Address, error) {
	return _ContractISlasher.Contract.Delegation(&_ContractISlasher.CallOpts)
}

// GetCorrectValueForInsertAfter is a free data retrieval call binding the contract method 0x723e59c7.
//
// Solidity: function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) view returns(uint256)
func (_ContractISlasher *ContractISlasherCaller) GetCorrectValueForInsertAfter(opts *bind.CallOpts, operator common.Address, updateBlock uint32) (*big.Int, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "getCorrectValueForInsertAfter", operator, updateBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCorrectValueForInsertAfter is a free data retrieval call binding the contract method 0x723e59c7.
//
// Solidity: function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) view returns(uint256)
func (_ContractISlasher *ContractISlasherSession) GetCorrectValueForInsertAfter(operator common.Address, updateBlock uint32) (*big.Int, error) {
	return _ContractISlasher.Contract.GetCorrectValueForInsertAfter(&_ContractISlasher.CallOpts, operator, updateBlock)
}

// GetCorrectValueForInsertAfter is a free data retrieval call binding the contract method 0x723e59c7.
//
// Solidity: function getCorrectValueForInsertAfter(address operator, uint32 updateBlock) view returns(uint256)
func (_ContractISlasher *ContractISlasherCallerSession) GetCorrectValueForInsertAfter(operator common.Address, updateBlock uint32) (*big.Int, error) {
	return _ContractISlasher.Contract.GetCorrectValueForInsertAfter(&_ContractISlasher.CallOpts, operator, updateBlock)
}

// GetMiddlewareTimesIndexServeUntilBlock is a free data retrieval call binding the contract method 0x7259a45c.
//
// Solidity: function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) view returns(uint32)
func (_ContractISlasher *ContractISlasherCaller) GetMiddlewareTimesIndexServeUntilBlock(opts *bind.CallOpts, operator common.Address, index uint32) (uint32, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "getMiddlewareTimesIndexServeUntilBlock", operator, index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMiddlewareTimesIndexServeUntilBlock is a free data retrieval call binding the contract method 0x7259a45c.
//
// Solidity: function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) view returns(uint32)
func (_ContractISlasher *ContractISlasherSession) GetMiddlewareTimesIndexServeUntilBlock(operator common.Address, index uint32) (uint32, error) {
	return _ContractISlasher.Contract.GetMiddlewareTimesIndexServeUntilBlock(&_ContractISlasher.CallOpts, operator, index)
}

// GetMiddlewareTimesIndexServeUntilBlock is a free data retrieval call binding the contract method 0x7259a45c.
//
// Solidity: function getMiddlewareTimesIndexServeUntilBlock(address operator, uint32 index) view returns(uint32)
func (_ContractISlasher *ContractISlasherCallerSession) GetMiddlewareTimesIndexServeUntilBlock(operator common.Address, index uint32) (uint32, error) {
	return _ContractISlasher.Contract.GetMiddlewareTimesIndexServeUntilBlock(&_ContractISlasher.CallOpts, operator, index)
}

// GetMiddlewareTimesIndexStalestUpdateBlock is a free data retrieval call binding the contract method 0x1874e5ae.
//
// Solidity: function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) view returns(uint32)
func (_ContractISlasher *ContractISlasherCaller) GetMiddlewareTimesIndexStalestUpdateBlock(opts *bind.CallOpts, operator common.Address, index uint32) (uint32, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "getMiddlewareTimesIndexStalestUpdateBlock", operator, index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMiddlewareTimesIndexStalestUpdateBlock is a free data retrieval call binding the contract method 0x1874e5ae.
//
// Solidity: function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) view returns(uint32)
func (_ContractISlasher *ContractISlasherSession) GetMiddlewareTimesIndexStalestUpdateBlock(operator common.Address, index uint32) (uint32, error) {
	return _ContractISlasher.Contract.GetMiddlewareTimesIndexStalestUpdateBlock(&_ContractISlasher.CallOpts, operator, index)
}

// GetMiddlewareTimesIndexStalestUpdateBlock is a free data retrieval call binding the contract method 0x1874e5ae.
//
// Solidity: function getMiddlewareTimesIndexStalestUpdateBlock(address operator, uint32 index) view returns(uint32)
func (_ContractISlasher *ContractISlasherCallerSession) GetMiddlewareTimesIndexStalestUpdateBlock(operator common.Address, index uint32) (uint32, error) {
	return _ContractISlasher.Contract.GetMiddlewareTimesIndexStalestUpdateBlock(&_ContractISlasher.CallOpts, operator, index)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address staker) view returns(bool)
func (_ContractISlasher *ContractISlasherCaller) IsFrozen(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "isFrozen", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address staker) view returns(bool)
func (_ContractISlasher *ContractISlasherSession) IsFrozen(staker common.Address) (bool, error) {
	return _ContractISlasher.Contract.IsFrozen(&_ContractISlasher.CallOpts, staker)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address staker) view returns(bool)
func (_ContractISlasher *ContractISlasherCallerSession) IsFrozen(staker common.Address) (bool, error) {
	return _ContractISlasher.Contract.IsFrozen(&_ContractISlasher.CallOpts, staker)
}

// LatestUpdateBlock is a free data retrieval call binding the contract method 0xda16e29b.
//
// Solidity: function latestUpdateBlock(address operator, address serviceContract) view returns(uint32)
func (_ContractISlasher *ContractISlasherCaller) LatestUpdateBlock(opts *bind.CallOpts, operator common.Address, serviceContract common.Address) (uint32, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "latestUpdateBlock", operator, serviceContract)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestUpdateBlock is a free data retrieval call binding the contract method 0xda16e29b.
//
// Solidity: function latestUpdateBlock(address operator, address serviceContract) view returns(uint32)
func (_ContractISlasher *ContractISlasherSession) LatestUpdateBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ContractISlasher.Contract.LatestUpdateBlock(&_ContractISlasher.CallOpts, operator, serviceContract)
}

// LatestUpdateBlock is a free data retrieval call binding the contract method 0xda16e29b.
//
// Solidity: function latestUpdateBlock(address operator, address serviceContract) view returns(uint32)
func (_ContractISlasher *ContractISlasherCallerSession) LatestUpdateBlock(operator common.Address, serviceContract common.Address) (uint32, error) {
	return _ContractISlasher.Contract.LatestUpdateBlock(&_ContractISlasher.CallOpts, operator, serviceContract)
}

// MiddlewareTimesLength is a free data retrieval call binding the contract method 0xa49db732.
//
// Solidity: function middlewareTimesLength(address operator) view returns(uint256)
func (_ContractISlasher *ContractISlasherCaller) MiddlewareTimesLength(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "middlewareTimesLength", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiddlewareTimesLength is a free data retrieval call binding the contract method 0xa49db732.
//
// Solidity: function middlewareTimesLength(address operator) view returns(uint256)
func (_ContractISlasher *ContractISlasherSession) MiddlewareTimesLength(operator common.Address) (*big.Int, error) {
	return _ContractISlasher.Contract.MiddlewareTimesLength(&_ContractISlasher.CallOpts, operator)
}

// MiddlewareTimesLength is a free data retrieval call binding the contract method 0xa49db732.
//
// Solidity: function middlewareTimesLength(address operator) view returns(uint256)
func (_ContractISlasher *ContractISlasherCallerSession) MiddlewareTimesLength(operator common.Address) (*big.Int, error) {
	return _ContractISlasher.Contract.MiddlewareTimesLength(&_ContractISlasher.CallOpts, operator)
}

// OperatorToMiddlewareTimes is a free data retrieval call binding the contract method 0x282670fc.
//
// Solidity: function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) view returns((uint32,uint32))
func (_ContractISlasher *ContractISlasherCaller) OperatorToMiddlewareTimes(opts *bind.CallOpts, operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "operatorToMiddlewareTimes", operator, arrayIndex)

	if err != nil {
		return *new(ISlasherMiddlewareTimes), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlasherMiddlewareTimes)).(*ISlasherMiddlewareTimes)

	return out0, err

}

// OperatorToMiddlewareTimes is a free data retrieval call binding the contract method 0x282670fc.
//
// Solidity: function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) view returns((uint32,uint32))
func (_ContractISlasher *ContractISlasherSession) OperatorToMiddlewareTimes(operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error) {
	return _ContractISlasher.Contract.OperatorToMiddlewareTimes(&_ContractISlasher.CallOpts, operator, arrayIndex)
}

// OperatorToMiddlewareTimes is a free data retrieval call binding the contract method 0x282670fc.
//
// Solidity: function operatorToMiddlewareTimes(address operator, uint256 arrayIndex) view returns((uint32,uint32))
func (_ContractISlasher *ContractISlasherCallerSession) OperatorToMiddlewareTimes(operator common.Address, arrayIndex *big.Int) (ISlasherMiddlewareTimes, error) {
	return _ContractISlasher.Contract.OperatorToMiddlewareTimes(&_ContractISlasher.CallOpts, operator, arrayIndex)
}

// OperatorWhitelistedContractsLinkedListEntry is a free data retrieval call binding the contract method 0x855fcc4a.
//
// Solidity: function operatorWhitelistedContractsLinkedListEntry(address operator, address node) view returns(bool, uint256, uint256)
func (_ContractISlasher *ContractISlasherCaller) OperatorWhitelistedContractsLinkedListEntry(opts *bind.CallOpts, operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "operatorWhitelistedContractsLinkedListEntry", operator, node)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// OperatorWhitelistedContractsLinkedListEntry is a free data retrieval call binding the contract method 0x855fcc4a.
//
// Solidity: function operatorWhitelistedContractsLinkedListEntry(address operator, address node) view returns(bool, uint256, uint256)
func (_ContractISlasher *ContractISlasherSession) OperatorWhitelistedContractsLinkedListEntry(operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error) {
	return _ContractISlasher.Contract.OperatorWhitelistedContractsLinkedListEntry(&_ContractISlasher.CallOpts, operator, node)
}

// OperatorWhitelistedContractsLinkedListEntry is a free data retrieval call binding the contract method 0x855fcc4a.
//
// Solidity: function operatorWhitelistedContractsLinkedListEntry(address operator, address node) view returns(bool, uint256, uint256)
func (_ContractISlasher *ContractISlasherCallerSession) OperatorWhitelistedContractsLinkedListEntry(operator common.Address, node common.Address) (bool, *big.Int, *big.Int, error) {
	return _ContractISlasher.Contract.OperatorWhitelistedContractsLinkedListEntry(&_ContractISlasher.CallOpts, operator, node)
}

// OperatorWhitelistedContractsLinkedListSize is a free data retrieval call binding the contract method 0xe921d4fa.
//
// Solidity: function operatorWhitelistedContractsLinkedListSize(address operator) view returns(uint256)
func (_ContractISlasher *ContractISlasherCaller) OperatorWhitelistedContractsLinkedListSize(opts *bind.CallOpts, operator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "operatorWhitelistedContractsLinkedListSize", operator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OperatorWhitelistedContractsLinkedListSize is a free data retrieval call binding the contract method 0xe921d4fa.
//
// Solidity: function operatorWhitelistedContractsLinkedListSize(address operator) view returns(uint256)
func (_ContractISlasher *ContractISlasherSession) OperatorWhitelistedContractsLinkedListSize(operator common.Address) (*big.Int, error) {
	return _ContractISlasher.Contract.OperatorWhitelistedContractsLinkedListSize(&_ContractISlasher.CallOpts, operator)
}

// OperatorWhitelistedContractsLinkedListSize is a free data retrieval call binding the contract method 0xe921d4fa.
//
// Solidity: function operatorWhitelistedContractsLinkedListSize(address operator) view returns(uint256)
func (_ContractISlasher *ContractISlasherCallerSession) OperatorWhitelistedContractsLinkedListSize(operator common.Address) (*big.Int, error) {
	return _ContractISlasher.Contract.OperatorWhitelistedContractsLinkedListSize(&_ContractISlasher.CallOpts, operator)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractISlasher *ContractISlasherCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractISlasher.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractISlasher *ContractISlasherSession) StrategyManager() (common.Address, error) {
	return _ContractISlasher.Contract.StrategyManager(&_ContractISlasher.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractISlasher *ContractISlasherCallerSession) StrategyManager() (common.Address, error) {
	return _ContractISlasher.Contract.StrategyManager(&_ContractISlasher.CallOpts)
}

// CanWithdraw is a paid mutator transaction binding the contract method 0x8105e043.
//
// Solidity: function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) returns(bool)
func (_ContractISlasher *ContractISlasherTransactor) CanWithdraw(opts *bind.TransactOpts, operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "canWithdraw", operator, withdrawalStartBlock, middlewareTimesIndex)
}

// CanWithdraw is a paid mutator transaction binding the contract method 0x8105e043.
//
// Solidity: function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) returns(bool)
func (_ContractISlasher *ContractISlasherSession) CanWithdraw(operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _ContractISlasher.Contract.CanWithdraw(&_ContractISlasher.TransactOpts, operator, withdrawalStartBlock, middlewareTimesIndex)
}

// CanWithdraw is a paid mutator transaction binding the contract method 0x8105e043.
//
// Solidity: function canWithdraw(address operator, uint32 withdrawalStartBlock, uint256 middlewareTimesIndex) returns(bool)
func (_ContractISlasher *ContractISlasherTransactorSession) CanWithdraw(operator common.Address, withdrawalStartBlock uint32, middlewareTimesIndex *big.Int) (*types.Transaction, error) {
	return _ContractISlasher.Contract.CanWithdraw(&_ContractISlasher.TransactOpts, operator, withdrawalStartBlock, middlewareTimesIndex)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address toBeFrozen) returns()
func (_ContractISlasher *ContractISlasherTransactor) FreezeOperator(opts *bind.TransactOpts, toBeFrozen common.Address) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "freezeOperator", toBeFrozen)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address toBeFrozen) returns()
func (_ContractISlasher *ContractISlasherSession) FreezeOperator(toBeFrozen common.Address) (*types.Transaction, error) {
	return _ContractISlasher.Contract.FreezeOperator(&_ContractISlasher.TransactOpts, toBeFrozen)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address toBeFrozen) returns()
func (_ContractISlasher *ContractISlasherTransactorSession) FreezeOperator(toBeFrozen common.Address) (*types.Transaction, error) {
	return _ContractISlasher.Contract.FreezeOperator(&_ContractISlasher.TransactOpts, toBeFrozen)
}

// OptIntoSlashing is a paid mutator transaction binding the contract method 0xf73b7519.
//
// Solidity: function optIntoSlashing(address contractAddress) returns()
func (_ContractISlasher *ContractISlasherTransactor) OptIntoSlashing(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "optIntoSlashing", contractAddress)
}

// OptIntoSlashing is a paid mutator transaction binding the contract method 0xf73b7519.
//
// Solidity: function optIntoSlashing(address contractAddress) returns()
func (_ContractISlasher *ContractISlasherSession) OptIntoSlashing(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractISlasher.Contract.OptIntoSlashing(&_ContractISlasher.TransactOpts, contractAddress)
}

// OptIntoSlashing is a paid mutator transaction binding the contract method 0xf73b7519.
//
// Solidity: function optIntoSlashing(address contractAddress) returns()
func (_ContractISlasher *ContractISlasherTransactorSession) OptIntoSlashing(contractAddress common.Address) (*types.Transaction, error) {
	return _ContractISlasher.Contract.OptIntoSlashing(&_ContractISlasher.TransactOpts, contractAddress)
}

// RecordFirstStakeUpdate is a paid mutator transaction binding the contract method 0x175d3205.
//
// Solidity: function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) returns()
func (_ContractISlasher *ContractISlasherTransactor) RecordFirstStakeUpdate(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "recordFirstStakeUpdate", operator, serveUntilBlock)
}

// RecordFirstStakeUpdate is a paid mutator transaction binding the contract method 0x175d3205.
//
// Solidity: function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) returns()
func (_ContractISlasher *ContractISlasherSession) RecordFirstStakeUpdate(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ContractISlasher.Contract.RecordFirstStakeUpdate(&_ContractISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordFirstStakeUpdate is a paid mutator transaction binding the contract method 0x175d3205.
//
// Solidity: function recordFirstStakeUpdate(address operator, uint32 serveUntilBlock) returns()
func (_ContractISlasher *ContractISlasherTransactorSession) RecordFirstStakeUpdate(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ContractISlasher.Contract.RecordFirstStakeUpdate(&_ContractISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_ContractISlasher *ContractISlasherTransactor) RecordLastStakeUpdateAndRevokeSlashingAbility(opts *bind.TransactOpts, operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "recordLastStakeUpdateAndRevokeSlashingAbility", operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_ContractISlasher *ContractISlasherSession) RecordLastStakeUpdateAndRevokeSlashingAbility(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ContractISlasher.Contract.RecordLastStakeUpdateAndRevokeSlashingAbility(&_ContractISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordLastStakeUpdateAndRevokeSlashingAbility is a paid mutator transaction binding the contract method 0x0ffabbce.
//
// Solidity: function recordLastStakeUpdateAndRevokeSlashingAbility(address operator, uint32 serveUntilBlock) returns()
func (_ContractISlasher *ContractISlasherTransactorSession) RecordLastStakeUpdateAndRevokeSlashingAbility(operator common.Address, serveUntilBlock uint32) (*types.Transaction, error) {
	return _ContractISlasher.Contract.RecordLastStakeUpdateAndRevokeSlashingAbility(&_ContractISlasher.TransactOpts, operator, serveUntilBlock)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) returns()
func (_ContractISlasher *ContractISlasherTransactor) RecordStakeUpdate(opts *bind.TransactOpts, operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "recordStakeUpdate", operator, updateBlock, serveUntilBlock, insertAfter)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) returns()
func (_ContractISlasher *ContractISlasherSession) RecordStakeUpdate(operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error) {
	return _ContractISlasher.Contract.RecordStakeUpdate(&_ContractISlasher.TransactOpts, operator, updateBlock, serveUntilBlock, insertAfter)
}

// RecordStakeUpdate is a paid mutator transaction binding the contract method 0xc747075b.
//
// Solidity: function recordStakeUpdate(address operator, uint32 updateBlock, uint32 serveUntilBlock, uint256 insertAfter) returns()
func (_ContractISlasher *ContractISlasherTransactorSession) RecordStakeUpdate(operator common.Address, updateBlock uint32, serveUntilBlock uint32, insertAfter *big.Int) (*types.Transaction, error) {
	return _ContractISlasher.Contract.RecordStakeUpdate(&_ContractISlasher.TransactOpts, operator, updateBlock, serveUntilBlock, insertAfter)
}

// ResetFrozenStatus is a paid mutator transaction binding the contract method 0x7cf72bba.
//
// Solidity: function resetFrozenStatus(address[] frozenAddresses) returns()
func (_ContractISlasher *ContractISlasherTransactor) ResetFrozenStatus(opts *bind.TransactOpts, frozenAddresses []common.Address) (*types.Transaction, error) {
	return _ContractISlasher.contract.Transact(opts, "resetFrozenStatus", frozenAddresses)
}

// ResetFrozenStatus is a paid mutator transaction binding the contract method 0x7cf72bba.
//
// Solidity: function resetFrozenStatus(address[] frozenAddresses) returns()
func (_ContractISlasher *ContractISlasherSession) ResetFrozenStatus(frozenAddresses []common.Address) (*types.Transaction, error) {
	return _ContractISlasher.Contract.ResetFrozenStatus(&_ContractISlasher.TransactOpts, frozenAddresses)
}

// ResetFrozenStatus is a paid mutator transaction binding the contract method 0x7cf72bba.
//
// Solidity: function resetFrozenStatus(address[] frozenAddresses) returns()
func (_ContractISlasher *ContractISlasherTransactorSession) ResetFrozenStatus(frozenAddresses []common.Address) (*types.Transaction, error) {
	return _ContractISlasher.Contract.ResetFrozenStatus(&_ContractISlasher.TransactOpts, frozenAddresses)
}

// ContractISlasherFrozenStatusResetIterator is returned from FilterFrozenStatusReset and is used to iterate over the raw logs and unpacked data for FrozenStatusReset events raised by the ContractISlasher contract.
type ContractISlasherFrozenStatusResetIterator struct {
	Event *ContractISlasherFrozenStatusReset // Event containing the contract specifics and raw log

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
func (it *ContractISlasherFrozenStatusResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractISlasherFrozenStatusReset)
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
		it.Event = new(ContractISlasherFrozenStatusReset)
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
func (it *ContractISlasherFrozenStatusResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractISlasherFrozenStatusResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractISlasherFrozenStatusReset represents a FrozenStatusReset event raised by the ContractISlasher contract.
type ContractISlasherFrozenStatusReset struct {
	PreviouslySlashedAddress common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterFrozenStatusReset is a free log retrieval operation binding the contract event 0xd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d8.
//
// Solidity: event FrozenStatusReset(address indexed previouslySlashedAddress)
func (_ContractISlasher *ContractISlasherFilterer) FilterFrozenStatusReset(opts *bind.FilterOpts, previouslySlashedAddress []common.Address) (*ContractISlasherFrozenStatusResetIterator, error) {

	var previouslySlashedAddressRule []interface{}
	for _, previouslySlashedAddressItem := range previouslySlashedAddress {
		previouslySlashedAddressRule = append(previouslySlashedAddressRule, previouslySlashedAddressItem)
	}

	logs, sub, err := _ContractISlasher.contract.FilterLogs(opts, "FrozenStatusReset", previouslySlashedAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherFrozenStatusResetIterator{contract: _ContractISlasher.contract, event: "FrozenStatusReset", logs: logs, sub: sub}, nil
}

// WatchFrozenStatusReset is a free log subscription operation binding the contract event 0xd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d8.
//
// Solidity: event FrozenStatusReset(address indexed previouslySlashedAddress)
func (_ContractISlasher *ContractISlasherFilterer) WatchFrozenStatusReset(opts *bind.WatchOpts, sink chan<- *ContractISlasherFrozenStatusReset, previouslySlashedAddress []common.Address) (event.Subscription, error) {

	var previouslySlashedAddressRule []interface{}
	for _, previouslySlashedAddressItem := range previouslySlashedAddress {
		previouslySlashedAddressRule = append(previouslySlashedAddressRule, previouslySlashedAddressItem)
	}

	logs, sub, err := _ContractISlasher.contract.WatchLogs(opts, "FrozenStatusReset", previouslySlashedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractISlasherFrozenStatusReset)
				if err := _ContractISlasher.contract.UnpackLog(event, "FrozenStatusReset", log); err != nil {
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

// ParseFrozenStatusReset is a log parse operation binding the contract event 0xd4cef0af27800d466fcacd85779857378b85cb61569005ff1464fa6e5ced69d8.
//
// Solidity: event FrozenStatusReset(address indexed previouslySlashedAddress)
func (_ContractISlasher *ContractISlasherFilterer) ParseFrozenStatusReset(log types.Log) (*ContractISlasherFrozenStatusReset, error) {
	event := new(ContractISlasherFrozenStatusReset)
	if err := _ContractISlasher.contract.UnpackLog(event, "FrozenStatusReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractISlasherMiddlewareTimesAddedIterator is returned from FilterMiddlewareTimesAdded and is used to iterate over the raw logs and unpacked data for MiddlewareTimesAdded events raised by the ContractISlasher contract.
type ContractISlasherMiddlewareTimesAddedIterator struct {
	Event *ContractISlasherMiddlewareTimesAdded // Event containing the contract specifics and raw log

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
func (it *ContractISlasherMiddlewareTimesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractISlasherMiddlewareTimesAdded)
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
		it.Event = new(ContractISlasherMiddlewareTimesAdded)
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
func (it *ContractISlasherMiddlewareTimesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractISlasherMiddlewareTimesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractISlasherMiddlewareTimesAdded represents a MiddlewareTimesAdded event raised by the ContractISlasher contract.
type ContractISlasherMiddlewareTimesAdded struct {
	Operator              common.Address
	Index                 *big.Int
	StalestUpdateBlock    uint32
	LatestServeUntilBlock uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMiddlewareTimesAdded is a free log retrieval operation binding the contract event 0x1b62ba64c72d01e41a2b8c46e6aeeff728ef3a4438cf1cac3d92ee12189d5649.
//
// Solidity: event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
func (_ContractISlasher *ContractISlasherFilterer) FilterMiddlewareTimesAdded(opts *bind.FilterOpts) (*ContractISlasherMiddlewareTimesAddedIterator, error) {

	logs, sub, err := _ContractISlasher.contract.FilterLogs(opts, "MiddlewareTimesAdded")
	if err != nil {
		return nil, err
	}
	return &ContractISlasherMiddlewareTimesAddedIterator{contract: _ContractISlasher.contract, event: "MiddlewareTimesAdded", logs: logs, sub: sub}, nil
}

// WatchMiddlewareTimesAdded is a free log subscription operation binding the contract event 0x1b62ba64c72d01e41a2b8c46e6aeeff728ef3a4438cf1cac3d92ee12189d5649.
//
// Solidity: event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
func (_ContractISlasher *ContractISlasherFilterer) WatchMiddlewareTimesAdded(opts *bind.WatchOpts, sink chan<- *ContractISlasherMiddlewareTimesAdded) (event.Subscription, error) {

	logs, sub, err := _ContractISlasher.contract.WatchLogs(opts, "MiddlewareTimesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractISlasherMiddlewareTimesAdded)
				if err := _ContractISlasher.contract.UnpackLog(event, "MiddlewareTimesAdded", log); err != nil {
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

// ParseMiddlewareTimesAdded is a log parse operation binding the contract event 0x1b62ba64c72d01e41a2b8c46e6aeeff728ef3a4438cf1cac3d92ee12189d5649.
//
// Solidity: event MiddlewareTimesAdded(address operator, uint256 index, uint32 stalestUpdateBlock, uint32 latestServeUntilBlock)
func (_ContractISlasher *ContractISlasherFilterer) ParseMiddlewareTimesAdded(log types.Log) (*ContractISlasherMiddlewareTimesAdded, error) {
	event := new(ContractISlasherMiddlewareTimesAdded)
	if err := _ContractISlasher.contract.UnpackLog(event, "MiddlewareTimesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractISlasherOperatorFrozenIterator is returned from FilterOperatorFrozen and is used to iterate over the raw logs and unpacked data for OperatorFrozen events raised by the ContractISlasher contract.
type ContractISlasherOperatorFrozenIterator struct {
	Event *ContractISlasherOperatorFrozen // Event containing the contract specifics and raw log

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
func (it *ContractISlasherOperatorFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractISlasherOperatorFrozen)
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
		it.Event = new(ContractISlasherOperatorFrozen)
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
func (it *ContractISlasherOperatorFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractISlasherOperatorFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractISlasherOperatorFrozen represents a OperatorFrozen event raised by the ContractISlasher contract.
type ContractISlasherOperatorFrozen struct {
	SlashedOperator  common.Address
	SlashingContract common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperatorFrozen is a free log retrieval operation binding the contract event 0x444a84f512816ae7be8ed8a66aa88e362eb54d0988e83acc9d81746622b3ba51.
//
// Solidity: event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract)
func (_ContractISlasher *ContractISlasherFilterer) FilterOperatorFrozen(opts *bind.FilterOpts, slashedOperator []common.Address, slashingContract []common.Address) (*ContractISlasherOperatorFrozenIterator, error) {

	var slashedOperatorRule []interface{}
	for _, slashedOperatorItem := range slashedOperator {
		slashedOperatorRule = append(slashedOperatorRule, slashedOperatorItem)
	}
	var slashingContractRule []interface{}
	for _, slashingContractItem := range slashingContract {
		slashingContractRule = append(slashingContractRule, slashingContractItem)
	}

	logs, sub, err := _ContractISlasher.contract.FilterLogs(opts, "OperatorFrozen", slashedOperatorRule, slashingContractRule)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherOperatorFrozenIterator{contract: _ContractISlasher.contract, event: "OperatorFrozen", logs: logs, sub: sub}, nil
}

// WatchOperatorFrozen is a free log subscription operation binding the contract event 0x444a84f512816ae7be8ed8a66aa88e362eb54d0988e83acc9d81746622b3ba51.
//
// Solidity: event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract)
func (_ContractISlasher *ContractISlasherFilterer) WatchOperatorFrozen(opts *bind.WatchOpts, sink chan<- *ContractISlasherOperatorFrozen, slashedOperator []common.Address, slashingContract []common.Address) (event.Subscription, error) {

	var slashedOperatorRule []interface{}
	for _, slashedOperatorItem := range slashedOperator {
		slashedOperatorRule = append(slashedOperatorRule, slashedOperatorItem)
	}
	var slashingContractRule []interface{}
	for _, slashingContractItem := range slashingContract {
		slashingContractRule = append(slashingContractRule, slashingContractItem)
	}

	logs, sub, err := _ContractISlasher.contract.WatchLogs(opts, "OperatorFrozen", slashedOperatorRule, slashingContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractISlasherOperatorFrozen)
				if err := _ContractISlasher.contract.UnpackLog(event, "OperatorFrozen", log); err != nil {
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

// ParseOperatorFrozen is a log parse operation binding the contract event 0x444a84f512816ae7be8ed8a66aa88e362eb54d0988e83acc9d81746622b3ba51.
//
// Solidity: event OperatorFrozen(address indexed slashedOperator, address indexed slashingContract)
func (_ContractISlasher *ContractISlasherFilterer) ParseOperatorFrozen(log types.Log) (*ContractISlasherOperatorFrozen, error) {
	event := new(ContractISlasherOperatorFrozen)
	if err := _ContractISlasher.contract.UnpackLog(event, "OperatorFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractISlasherOptedIntoSlashingIterator is returned from FilterOptedIntoSlashing and is used to iterate over the raw logs and unpacked data for OptedIntoSlashing events raised by the ContractISlasher contract.
type ContractISlasherOptedIntoSlashingIterator struct {
	Event *ContractISlasherOptedIntoSlashing // Event containing the contract specifics and raw log

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
func (it *ContractISlasherOptedIntoSlashingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractISlasherOptedIntoSlashing)
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
		it.Event = new(ContractISlasherOptedIntoSlashing)
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
func (it *ContractISlasherOptedIntoSlashingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractISlasherOptedIntoSlashingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractISlasherOptedIntoSlashing represents a OptedIntoSlashing event raised by the ContractISlasher contract.
type ContractISlasherOptedIntoSlashing struct {
	Operator        common.Address
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOptedIntoSlashing is a free log retrieval operation binding the contract event 0xefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d8.
//
// Solidity: event OptedIntoSlashing(address indexed operator, address indexed contractAddress)
func (_ContractISlasher *ContractISlasherFilterer) FilterOptedIntoSlashing(opts *bind.FilterOpts, operator []common.Address, contractAddress []common.Address) (*ContractISlasherOptedIntoSlashingIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractISlasher.contract.FilterLogs(opts, "OptedIntoSlashing", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherOptedIntoSlashingIterator{contract: _ContractISlasher.contract, event: "OptedIntoSlashing", logs: logs, sub: sub}, nil
}

// WatchOptedIntoSlashing is a free log subscription operation binding the contract event 0xefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d8.
//
// Solidity: event OptedIntoSlashing(address indexed operator, address indexed contractAddress)
func (_ContractISlasher *ContractISlasherFilterer) WatchOptedIntoSlashing(opts *bind.WatchOpts, sink chan<- *ContractISlasherOptedIntoSlashing, operator []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractISlasher.contract.WatchLogs(opts, "OptedIntoSlashing", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractISlasherOptedIntoSlashing)
				if err := _ContractISlasher.contract.UnpackLog(event, "OptedIntoSlashing", log); err != nil {
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

// ParseOptedIntoSlashing is a log parse operation binding the contract event 0xefa9fb38e813d53c15edf501e03852843a3fed691960523391d71a092b3627d8.
//
// Solidity: event OptedIntoSlashing(address indexed operator, address indexed contractAddress)
func (_ContractISlasher *ContractISlasherFilterer) ParseOptedIntoSlashing(log types.Log) (*ContractISlasherOptedIntoSlashing, error) {
	event := new(ContractISlasherOptedIntoSlashing)
	if err := _ContractISlasher.contract.UnpackLog(event, "OptedIntoSlashing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractISlasherSlashingAbilityRevokedIterator is returned from FilterSlashingAbilityRevoked and is used to iterate over the raw logs and unpacked data for SlashingAbilityRevoked events raised by the ContractISlasher contract.
type ContractISlasherSlashingAbilityRevokedIterator struct {
	Event *ContractISlasherSlashingAbilityRevoked // Event containing the contract specifics and raw log

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
func (it *ContractISlasherSlashingAbilityRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractISlasherSlashingAbilityRevoked)
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
		it.Event = new(ContractISlasherSlashingAbilityRevoked)
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
func (it *ContractISlasherSlashingAbilityRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractISlasherSlashingAbilityRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractISlasherSlashingAbilityRevoked represents a SlashingAbilityRevoked event raised by the ContractISlasher contract.
type ContractISlasherSlashingAbilityRevoked struct {
	Operator                           common.Address
	ContractAddress                    common.Address
	ContractCanSlashOperatorUntilBlock uint32
	Raw                                types.Log // Blockchain specific contextual infos
}

// FilterSlashingAbilityRevoked is a free log retrieval operation binding the contract event 0x9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14.
//
// Solidity: event SlashingAbilityRevoked(address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock)
func (_ContractISlasher *ContractISlasherFilterer) FilterSlashingAbilityRevoked(opts *bind.FilterOpts, operator []common.Address, contractAddress []common.Address) (*ContractISlasherSlashingAbilityRevokedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractISlasher.contract.FilterLogs(opts, "SlashingAbilityRevoked", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractISlasherSlashingAbilityRevokedIterator{contract: _ContractISlasher.contract, event: "SlashingAbilityRevoked", logs: logs, sub: sub}, nil
}

// WatchSlashingAbilityRevoked is a free log subscription operation binding the contract event 0x9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14.
//
// Solidity: event SlashingAbilityRevoked(address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock)
func (_ContractISlasher *ContractISlasherFilterer) WatchSlashingAbilityRevoked(opts *bind.WatchOpts, sink chan<- *ContractISlasherSlashingAbilityRevoked, operator []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _ContractISlasher.contract.WatchLogs(opts, "SlashingAbilityRevoked", operatorRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractISlasherSlashingAbilityRevoked)
				if err := _ContractISlasher.contract.UnpackLog(event, "SlashingAbilityRevoked", log); err != nil {
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

// ParseSlashingAbilityRevoked is a log parse operation binding the contract event 0x9aa1b1391f35c672ed1f3b7ece632f4513e618366bef7a2f67b7c6bc1f2d2b14.
//
// Solidity: event SlashingAbilityRevoked(address indexed operator, address indexed contractAddress, uint32 contractCanSlashOperatorUntilBlock)
func (_ContractISlasher *ContractISlasherFilterer) ParseSlashingAbilityRevoked(log types.Log) (*ContractISlasherSlashingAbilityRevoked, error) {
	event := new(ContractISlasherSlashingAbilityRevoked)
	if err := _ContractISlasher.contract.UnpackLog(event, "SlashingAbilityRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
