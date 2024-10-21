// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIStrategy

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

// ContractIStrategyMetaData contains all meta data concerning the ContractIStrategy contract.
var ContractIStrategyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"explanation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sharesToUnderlying\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesToUnderlyingView\",\"inputs\":[{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToShares\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"underlyingToSharesView\",\"inputs\":[{\"name\":\"amountUnderlying\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"underlyingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userUnderlying\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userUnderlyingView\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"amountShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ExchangeRateEmitted\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StrategyTokenSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIERC20\"},{\"name\":\"decimals\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BalanceExceedsMaxTotalDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxPerDepositExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewSharesZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyStrategyManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnderlyingToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TotalSharesExceedsMax\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAmountExceedsTotalDeposits\",\"inputs\":[]}]",
}

// ContractIStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIStrategyMetaData.ABI instead.
var ContractIStrategyABI = ContractIStrategyMetaData.ABI

// ContractIStrategyMethods is an auto generated interface around an Ethereum contract.
type ContractIStrategyMethods interface {
	ContractIStrategyCalls
	ContractIStrategyTransacts
	ContractIStrategyFilters
}

// ContractIStrategyCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIStrategyCalls interface {
	Explanation(opts *bind.CallOpts) (string, error)

	Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error)

	SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error)

	TotalShares(opts *bind.CallOpts) (*big.Int, error)

	UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error)

	UnderlyingToken(opts *bind.CallOpts) (common.Address, error)

	UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error)
}

// ContractIStrategyTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIStrategyTransacts interface {
	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)

	SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error)

	UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error)

	UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error)

	Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error)
}

// ContractIStrategyFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIStrategyFilters interface {
	FilterExchangeRateEmitted(opts *bind.FilterOpts) (*ContractIStrategyExchangeRateEmittedIterator, error)
	WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *ContractIStrategyExchangeRateEmitted) (event.Subscription, error)
	ParseExchangeRateEmitted(log types.Log) (*ContractIStrategyExchangeRateEmitted, error)

	FilterStrategyTokenSet(opts *bind.FilterOpts) (*ContractIStrategyStrategyTokenSetIterator, error)
	WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *ContractIStrategyStrategyTokenSet) (event.Subscription, error)
	ParseStrategyTokenSet(log types.Log) (*ContractIStrategyStrategyTokenSet, error)
}

// ContractIStrategy is an auto generated Go binding around an Ethereum contract.
type ContractIStrategy struct {
	ContractIStrategyCaller     // Read-only binding to the contract
	ContractIStrategyTransactor // Write-only binding to the contract
	ContractIStrategyFilterer   // Log filterer for contract events
}

// ContractIStrategy implements the ContractIStrategyMethods interface.
var _ ContractIStrategyMethods = (*ContractIStrategy)(nil)

// ContractIStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyCaller implements the ContractIStrategyCalls interface.
var _ ContractIStrategyCalls = (*ContractIStrategyCaller)(nil)

// ContractIStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyTransactor implements the ContractIStrategyTransacts interface.
var _ ContractIStrategyTransacts = (*ContractIStrategyTransactor)(nil)

// ContractIStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIStrategyFilterer implements the ContractIStrategyFilters interface.
var _ ContractIStrategyFilters = (*ContractIStrategyFilterer)(nil)

// ContractIStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIStrategySession struct {
	Contract     *ContractIStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractIStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIStrategyCallerSession struct {
	Contract *ContractIStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractIStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIStrategyTransactorSession struct {
	Contract     *ContractIStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractIStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIStrategyRaw struct {
	Contract *ContractIStrategy // Generic contract binding to access the raw methods on
}

// ContractIStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIStrategyCallerRaw struct {
	Contract *ContractIStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIStrategyTransactorRaw struct {
	Contract *ContractIStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIStrategy creates a new instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategy(address common.Address, backend bind.ContractBackend) (*ContractIStrategy, error) {
	contract, err := bindContractIStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategy{ContractIStrategyCaller: ContractIStrategyCaller{contract: contract}, ContractIStrategyTransactor: ContractIStrategyTransactor{contract: contract}, ContractIStrategyFilterer: ContractIStrategyFilterer{contract: contract}}, nil
}

// NewContractIStrategyCaller creates a new read-only instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategyCaller(address common.Address, caller bind.ContractCaller) (*ContractIStrategyCaller, error) {
	contract, err := bindContractIStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyCaller{contract: contract}, nil
}

// NewContractIStrategyTransactor creates a new write-only instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIStrategyTransactor, error) {
	contract, err := bindContractIStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyTransactor{contract: contract}, nil
}

// NewContractIStrategyFilterer creates a new log filterer instance of ContractIStrategy, bound to a specific deployed contract.
func NewContractIStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIStrategyFilterer, error) {
	contract, err := bindContractIStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyFilterer{contract: contract}, nil
}

// bindContractIStrategy binds a generic wrapper to an already deployed contract.
func bindContractIStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIStrategy *ContractIStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIStrategy.Contract.ContractIStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIStrategy *ContractIStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.ContractIStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIStrategy *ContractIStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.ContractIStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIStrategy *ContractIStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIStrategy *ContractIStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIStrategy *ContractIStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_ContractIStrategy *ContractIStrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_ContractIStrategy *ContractIStrategySession) Explanation() (string, error) {
	return _ContractIStrategy.Contract.Explanation(&_ContractIStrategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() view returns(string)
func (_ContractIStrategy *ContractIStrategyCallerSession) Explanation() (string, error) {
	return _ContractIStrategy.Contract.Explanation(&_ContractIStrategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) Shares(user common.Address) (*big.Int, error) {
	return _ContractIStrategy.Contract.Shares(&_ContractIStrategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _ContractIStrategy.Contract.Shares(&_ContractIStrategy.CallOpts, user)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.SharesToUnderlyingView(&_ContractIStrategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.SharesToUnderlyingView(&_ContractIStrategy.CallOpts, amountShares)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) TotalShares() (*big.Int, error) {
	return _ContractIStrategy.Contract.TotalShares(&_ContractIStrategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) TotalShares() (*big.Int, error) {
	return _ContractIStrategy.Contract.TotalShares(&_ContractIStrategy.CallOpts)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.UnderlyingToSharesView(&_ContractIStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _ContractIStrategy.Contract.UnderlyingToSharesView(&_ContractIStrategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractIStrategy *ContractIStrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractIStrategy *ContractIStrategySession) UnderlyingToken() (common.Address, error) {
	return _ContractIStrategy.Contract.UnderlyingToken(&_ContractIStrategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_ContractIStrategy *ContractIStrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _ContractIStrategy.Contract.UnderlyingToken(&_ContractIStrategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIStrategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _ContractIStrategy.Contract.UserUnderlyingView(&_ContractIStrategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_ContractIStrategy *ContractIStrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _ContractIStrategy.Contract.UserUnderlyingView(&_ContractIStrategy.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Deposit(&_ContractIStrategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Deposit(&_ContractIStrategy.TransactOpts, token, amount)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) SharesToUnderlying(opts *bind.TransactOpts, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "sharesToUnderlying", amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.SharesToUnderlying(&_ContractIStrategy.TransactOpts, amountShares)
}

// SharesToUnderlying is a paid mutator transaction binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) SharesToUnderlying(amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.SharesToUnderlying(&_ContractIStrategy.TransactOpts, amountShares)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) UnderlyingToShares(opts *bind.TransactOpts, amountUnderlying *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "underlyingToShares", amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UnderlyingToShares(&_ContractIStrategy.TransactOpts, amountUnderlying)
}

// UnderlyingToShares is a paid mutator transaction binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) UnderlyingToShares(amountUnderlying *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UnderlyingToShares(&_ContractIStrategy.TransactOpts, amountUnderlying)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractIStrategy *ContractIStrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UserUnderlying(&_ContractIStrategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_ContractIStrategy *ContractIStrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.UserUnderlying(&_ContractIStrategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_ContractIStrategy *ContractIStrategyTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.contract.Transact(opts, "withdraw", recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_ContractIStrategy *ContractIStrategySession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Withdraw(&_ContractIStrategy.TransactOpts, recipient, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address recipient, address token, uint256 amountShares) returns()
func (_ContractIStrategy *ContractIStrategyTransactorSession) Withdraw(recipient common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _ContractIStrategy.Contract.Withdraw(&_ContractIStrategy.TransactOpts, recipient, token, amountShares)
}

// ContractIStrategyExchangeRateEmittedIterator is returned from FilterExchangeRateEmitted and is used to iterate over the raw logs and unpacked data for ExchangeRateEmitted events raised by the ContractIStrategy contract.
type ContractIStrategyExchangeRateEmittedIterator struct {
	Event *ContractIStrategyExchangeRateEmitted // Event containing the contract specifics and raw log

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
func (it *ContractIStrategyExchangeRateEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIStrategyExchangeRateEmitted)
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
		it.Event = new(ContractIStrategyExchangeRateEmitted)
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
func (it *ContractIStrategyExchangeRateEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIStrategyExchangeRateEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIStrategyExchangeRateEmitted represents a ExchangeRateEmitted event raised by the ContractIStrategy contract.
type ContractIStrategyExchangeRateEmitted struct {
	Rate *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateEmitted is a free log retrieval operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_ContractIStrategy *ContractIStrategyFilterer) FilterExchangeRateEmitted(opts *bind.FilterOpts) (*ContractIStrategyExchangeRateEmittedIterator, error) {

	logs, sub, err := _ContractIStrategy.contract.FilterLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyExchangeRateEmittedIterator{contract: _ContractIStrategy.contract, event: "ExchangeRateEmitted", logs: logs, sub: sub}, nil
}

// WatchExchangeRateEmitted is a free log subscription operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_ContractIStrategy *ContractIStrategyFilterer) WatchExchangeRateEmitted(opts *bind.WatchOpts, sink chan<- *ContractIStrategyExchangeRateEmitted) (event.Subscription, error) {

	logs, sub, err := _ContractIStrategy.contract.WatchLogs(opts, "ExchangeRateEmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIStrategyExchangeRateEmitted)
				if err := _ContractIStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
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

// ParseExchangeRateEmitted is a log parse operation binding the contract event 0xd2494f3479e5da49d386657c292c610b5b01df313d07c62eb0cfa49924a31be8.
//
// Solidity: event ExchangeRateEmitted(uint256 rate)
func (_ContractIStrategy *ContractIStrategyFilterer) ParseExchangeRateEmitted(log types.Log) (*ContractIStrategyExchangeRateEmitted, error) {
	event := new(ContractIStrategyExchangeRateEmitted)
	if err := _ContractIStrategy.contract.UnpackLog(event, "ExchangeRateEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIStrategyStrategyTokenSetIterator is returned from FilterStrategyTokenSet and is used to iterate over the raw logs and unpacked data for StrategyTokenSet events raised by the ContractIStrategy contract.
type ContractIStrategyStrategyTokenSetIterator struct {
	Event *ContractIStrategyStrategyTokenSet // Event containing the contract specifics and raw log

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
func (it *ContractIStrategyStrategyTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIStrategyStrategyTokenSet)
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
		it.Event = new(ContractIStrategyStrategyTokenSet)
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
func (it *ContractIStrategyStrategyTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIStrategyStrategyTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIStrategyStrategyTokenSet represents a StrategyTokenSet event raised by the ContractIStrategy contract.
type ContractIStrategyStrategyTokenSet struct {
	Token    common.Address
	Decimals uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyTokenSet is a free log retrieval operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_ContractIStrategy *ContractIStrategyFilterer) FilterStrategyTokenSet(opts *bind.FilterOpts) (*ContractIStrategyStrategyTokenSetIterator, error) {

	logs, sub, err := _ContractIStrategy.contract.FilterLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return &ContractIStrategyStrategyTokenSetIterator{contract: _ContractIStrategy.contract, event: "StrategyTokenSet", logs: logs, sub: sub}, nil
}

// WatchStrategyTokenSet is a free log subscription operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_ContractIStrategy *ContractIStrategyFilterer) WatchStrategyTokenSet(opts *bind.WatchOpts, sink chan<- *ContractIStrategyStrategyTokenSet) (event.Subscription, error) {

	logs, sub, err := _ContractIStrategy.contract.WatchLogs(opts, "StrategyTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIStrategyStrategyTokenSet)
				if err := _ContractIStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
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

// ParseStrategyTokenSet is a log parse operation binding the contract event 0x1c540707b00eb5427b6b774fc799d756516a54aee108b64b327acc55af557507.
//
// Solidity: event StrategyTokenSet(address token, uint8 decimals)
func (_ContractIStrategy *ContractIStrategyFilterer) ParseStrategyTokenSet(log types.Log) (*ContractIStrategyStrategyTokenSet, error) {
	event := new(ContractIStrategyStrategyTokenSet)
	if err := _ContractIStrategy.contract.UnpackLog(event, "StrategyTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
