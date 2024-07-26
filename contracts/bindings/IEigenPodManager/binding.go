// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIEigenPodManager

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

// ContractIEigenPodManagerMetaData contains all meta data concerning the ContractIEigenPodManager contract.
var ContractIEigenPodManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beaconChainOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeaconChainOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"denebForkTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockRootAtTimestamp\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numPods\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwnerShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordBeaconChainETHBalanceUpdate\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDenebForkTimestamp\",\"inputs\":[{\"name\":\"newDenebForkTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISlasher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateBeaconChainOracle\",\"inputs\":[{\"name\":\"newBeaconChainOracle\",\"type\":\"address\",\"internalType\":\"contractIBeaconChainOracle\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconChainETHDeposited\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconChainETHWithdrawalCompleted\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconOracleUpdated\",\"inputs\":[{\"name\":\"newOracleAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DenebForkTimestampUpdated\",\"inputs\":[{\"name\":\"newValue\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodDeployed\",\"inputs\":[{\"name\":\"eigenPod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodSharesUpdated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// ContractIEigenPodManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIEigenPodManagerMetaData.ABI instead.
var ContractIEigenPodManagerABI = ContractIEigenPodManagerMetaData.ABI

// ContractIEigenPodManagerMethods is an auto generated interface around an Ethereum contract.
type ContractIEigenPodManagerMethods interface {
	ContractIEigenPodManagerCalls
	ContractIEigenPodManagerTransacts
	ContractIEigenPodManagerFilters
}

// ContractIEigenPodManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIEigenPodManagerCalls interface {
	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	BeaconChainOracle(opts *bind.CallOpts) (common.Address, error)

	DenebForkTimestamp(opts *bind.CallOpts) (uint64, error)

	EigenPodBeacon(opts *bind.CallOpts) (common.Address, error)

	EthPOS(opts *bind.CallOpts) (common.Address, error)

	GetBlockRootAtTimestamp(opts *bind.CallOpts, timestamp uint64) ([32]byte, error)

	GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error)

	HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error)

	NumPods(opts *bind.CallOpts) (*big.Int, error)

	OwnerToPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PodOwnerShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error)

	Slasher(opts *bind.CallOpts) (common.Address, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)
}

// ContractIEigenPodManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIEigenPodManagerTransacts interface {
	AddShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error)

	CreatePod(opts *bind.TransactOpts) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error)

	RemoveShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error)

	SetDenebForkTimestamp(opts *bind.TransactOpts, newDenebForkTimestamp uint64) (*types.Transaction, error)

	SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	UpdateBeaconChainOracle(opts *bind.TransactOpts, newBeaconChainOracle common.Address) (*types.Transaction, error)

	WithdrawSharesAsTokens(opts *bind.TransactOpts, podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error)
}

// ContractIEigenPodManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIEigenPodManagerFilters interface {
	FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodManagerBeaconChainETHDepositedIterator, error)
	WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error)
	ParseBeaconChainETHDeposited(log types.Log) (*ContractIEigenPodManagerBeaconChainETHDeposited, error)

	FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator, error)
	WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error)
	ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted, error)

	FilterBeaconOracleUpdated(opts *bind.FilterOpts, newOracleAddress []common.Address) (*ContractIEigenPodManagerBeaconOracleUpdatedIterator, error)
	WatchBeaconOracleUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerBeaconOracleUpdated, newOracleAddress []common.Address) (event.Subscription, error)
	ParseBeaconOracleUpdated(log types.Log) (*ContractIEigenPodManagerBeaconOracleUpdated, error)

	FilterDenebForkTimestampUpdated(opts *bind.FilterOpts) (*ContractIEigenPodManagerDenebForkTimestampUpdatedIterator, error)
	WatchDenebForkTimestampUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerDenebForkTimestampUpdated) (event.Subscription, error)
	ParseDenebForkTimestampUpdated(log types.Log) (*ContractIEigenPodManagerDenebForkTimestampUpdated, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIEigenPodManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractIEigenPodManagerPaused, error)

	FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIEigenPodManagerPauserRegistrySetIterator, error)
	WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPauserRegistrySet) (event.Subscription, error)
	ParsePauserRegistrySet(log types.Log) (*ContractIEigenPodManagerPauserRegistrySet, error)

	FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*ContractIEigenPodManagerPodDeployedIterator, error)
	WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error)
	ParsePodDeployed(log types.Log) (*ContractIEigenPodManagerPodDeployed, error)

	FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodManagerPodSharesUpdatedIterator, error)
	WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPodSharesUpdated, podOwner []common.Address) (event.Subscription, error)
	ParsePodSharesUpdated(log types.Log) (*ContractIEigenPodManagerPodSharesUpdated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIEigenPodManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractIEigenPodManagerUnpaused, error)
}

// ContractIEigenPodManager is an auto generated Go binding around an Ethereum contract.
type ContractIEigenPodManager struct {
	ContractIEigenPodManagerCaller     // Read-only binding to the contract
	ContractIEigenPodManagerTransactor // Write-only binding to the contract
	ContractIEigenPodManagerFilterer   // Log filterer for contract events
}

// ContractIEigenPodManager implements the ContractIEigenPodManagerMethods interface.
var _ ContractIEigenPodManagerMethods = (*ContractIEigenPodManager)(nil)

// ContractIEigenPodManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIEigenPodManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodManagerCaller implements the ContractIEigenPodManagerCalls interface.
var _ ContractIEigenPodManagerCalls = (*ContractIEigenPodManagerCaller)(nil)

// ContractIEigenPodManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIEigenPodManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodManagerTransactor implements the ContractIEigenPodManagerTransacts interface.
var _ ContractIEigenPodManagerTransacts = (*ContractIEigenPodManagerTransactor)(nil)

// ContractIEigenPodManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIEigenPodManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIEigenPodManagerFilterer implements the ContractIEigenPodManagerFilters interface.
var _ ContractIEigenPodManagerFilters = (*ContractIEigenPodManagerFilterer)(nil)

// ContractIEigenPodManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIEigenPodManagerSession struct {
	Contract     *ContractIEigenPodManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ContractIEigenPodManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIEigenPodManagerCallerSession struct {
	Contract *ContractIEigenPodManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ContractIEigenPodManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIEigenPodManagerTransactorSession struct {
	Contract     *ContractIEigenPodManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ContractIEigenPodManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIEigenPodManagerRaw struct {
	Contract *ContractIEigenPodManager // Generic contract binding to access the raw methods on
}

// ContractIEigenPodManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIEigenPodManagerCallerRaw struct {
	Contract *ContractIEigenPodManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIEigenPodManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIEigenPodManagerTransactorRaw struct {
	Contract *ContractIEigenPodManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIEigenPodManager creates a new instance of ContractIEigenPodManager, bound to a specific deployed contract.
func NewContractIEigenPodManager(address common.Address, backend bind.ContractBackend) (*ContractIEigenPodManager, error) {
	contract, err := bindContractIEigenPodManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManager{ContractIEigenPodManagerCaller: ContractIEigenPodManagerCaller{contract: contract}, ContractIEigenPodManagerTransactor: ContractIEigenPodManagerTransactor{contract: contract}, ContractIEigenPodManagerFilterer: ContractIEigenPodManagerFilterer{contract: contract}}, nil
}

// NewContractIEigenPodManagerCaller creates a new read-only instance of ContractIEigenPodManager, bound to a specific deployed contract.
func NewContractIEigenPodManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIEigenPodManagerCaller, error) {
	contract, err := bindContractIEigenPodManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerCaller{contract: contract}, nil
}

// NewContractIEigenPodManagerTransactor creates a new write-only instance of ContractIEigenPodManager, bound to a specific deployed contract.
func NewContractIEigenPodManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIEigenPodManagerTransactor, error) {
	contract, err := bindContractIEigenPodManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerTransactor{contract: contract}, nil
}

// NewContractIEigenPodManagerFilterer creates a new log filterer instance of ContractIEigenPodManager, bound to a specific deployed contract.
func NewContractIEigenPodManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIEigenPodManagerFilterer, error) {
	contract, err := bindContractIEigenPodManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerFilterer{contract: contract}, nil
}

// bindContractIEigenPodManager binds a generic wrapper to an already deployed contract.
func bindContractIEigenPodManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIEigenPodManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIEigenPodManager *ContractIEigenPodManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIEigenPodManager.Contract.ContractIEigenPodManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIEigenPodManager *ContractIEigenPodManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.ContractIEigenPodManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIEigenPodManager *ContractIEigenPodManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.ContractIEigenPodManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIEigenPodManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.contract.Transact(opts, method, params...)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.BeaconChainETHStrategy(&_ContractIEigenPodManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.BeaconChainETHStrategy(&_ContractIEigenPodManager.CallOpts)
}

// BeaconChainOracle is a free data retrieval call binding the contract method 0xc052bd61.
//
// Solidity: function beaconChainOracle() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) BeaconChainOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "beaconChainOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainOracle is a free data retrieval call binding the contract method 0xc052bd61.
//
// Solidity: function beaconChainOracle() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) BeaconChainOracle() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.BeaconChainOracle(&_ContractIEigenPodManager.CallOpts)
}

// BeaconChainOracle is a free data retrieval call binding the contract method 0xc052bd61.
//
// Solidity: function beaconChainOracle() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) BeaconChainOracle() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.BeaconChainOracle(&_ContractIEigenPodManager.CallOpts)
}

// DenebForkTimestamp is a free data retrieval call binding the contract method 0x44e71c80.
//
// Solidity: function denebForkTimestamp() view returns(uint64)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) DenebForkTimestamp(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "denebForkTimestamp")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DenebForkTimestamp is a free data retrieval call binding the contract method 0x44e71c80.
//
// Solidity: function denebForkTimestamp() view returns(uint64)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) DenebForkTimestamp() (uint64, error) {
	return _ContractIEigenPodManager.Contract.DenebForkTimestamp(&_ContractIEigenPodManager.CallOpts)
}

// DenebForkTimestamp is a free data retrieval call binding the contract method 0x44e71c80.
//
// Solidity: function denebForkTimestamp() view returns(uint64)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) DenebForkTimestamp() (uint64, error) {
	return _ContractIEigenPodManager.Contract.DenebForkTimestamp(&_ContractIEigenPodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) EigenPodBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "eigenPodBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) EigenPodBeacon() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.EigenPodBeacon(&_ContractIEigenPodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) EigenPodBeacon() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.EigenPodBeacon(&_ContractIEigenPodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) EthPOS() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.EthPOS(&_ContractIEigenPodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) EthPOS() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.EthPOS(&_ContractIEigenPodManager.CallOpts)
}

// GetBlockRootAtTimestamp is a free data retrieval call binding the contract method 0xd1c64cc9.
//
// Solidity: function getBlockRootAtTimestamp(uint64 timestamp) view returns(bytes32)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) GetBlockRootAtTimestamp(opts *bind.CallOpts, timestamp uint64) ([32]byte, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "getBlockRootAtTimestamp", timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockRootAtTimestamp is a free data retrieval call binding the contract method 0xd1c64cc9.
//
// Solidity: function getBlockRootAtTimestamp(uint64 timestamp) view returns(bytes32)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) GetBlockRootAtTimestamp(timestamp uint64) ([32]byte, error) {
	return _ContractIEigenPodManager.Contract.GetBlockRootAtTimestamp(&_ContractIEigenPodManager.CallOpts, timestamp)
}

// GetBlockRootAtTimestamp is a free data retrieval call binding the contract method 0xd1c64cc9.
//
// Solidity: function getBlockRootAtTimestamp(uint64 timestamp) view returns(bytes32)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) GetBlockRootAtTimestamp(timestamp uint64) ([32]byte, error) {
	return _ContractIEigenPodManager.Contract.GetBlockRootAtTimestamp(&_ContractIEigenPodManager.CallOpts, timestamp)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "getPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _ContractIEigenPodManager.Contract.GetPod(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _ContractIEigenPodManager.Contract.GetPod(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "hasPod", podOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) HasPod(podOwner common.Address) (bool, error) {
	return _ContractIEigenPodManager.Contract.HasPod(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) HasPod(podOwner common.Address) (bool, error) {
	return _ContractIEigenPodManager.Contract.HasPod(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) NumPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "numPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) NumPods() (*big.Int, error) {
	return _ContractIEigenPodManager.Contract.NumPods(&_ContractIEigenPodManager.CallOpts)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) NumPods() (*big.Int, error) {
	return _ContractIEigenPodManager.Contract.NumPods(&_ContractIEigenPodManager.CallOpts)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) OwnerToPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "ownerToPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _ContractIEigenPodManager.Contract.OwnerToPod(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address podOwner) view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) OwnerToPod(podOwner common.Address) (common.Address, error) {
	return _ContractIEigenPodManager.Contract.OwnerToPod(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) Paused(index uint8) (bool, error) {
	return _ContractIEigenPodManager.Contract.Paused(&_ContractIEigenPodManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractIEigenPodManager.Contract.Paused(&_ContractIEigenPodManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) Paused0() (*big.Int, error) {
	return _ContractIEigenPodManager.Contract.Paused0(&_ContractIEigenPodManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractIEigenPodManager.Contract.Paused0(&_ContractIEigenPodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.PauserRegistry(&_ContractIEigenPodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.PauserRegistry(&_ContractIEigenPodManager.CallOpts)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address podOwner) view returns(int256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) PodOwnerShares(opts *bind.CallOpts, podOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "podOwnerShares", podOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address podOwner) view returns(int256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) PodOwnerShares(podOwner common.Address) (*big.Int, error) {
	return _ContractIEigenPodManager.Contract.PodOwnerShares(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address podOwner) view returns(int256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) PodOwnerShares(podOwner common.Address) (*big.Int, error) {
	return _ContractIEigenPodManager.Contract.PodOwnerShares(&_ContractIEigenPodManager.CallOpts, podOwner)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) Slasher() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.Slasher(&_ContractIEigenPodManager.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) Slasher() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.Slasher(&_ContractIEigenPodManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIEigenPodManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) StrategyManager() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.StrategyManager(&_ContractIEigenPodManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerCallerSession) StrategyManager() (common.Address, error) {
	return _ContractIEigenPodManager.Contract.StrategyManager(&_ContractIEigenPodManager.CallOpts)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) AddShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "addShares", podOwner, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) AddShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.AddShares(&_ContractIEigenPodManager.TransactOpts, podOwner, shares)
}

// AddShares is a paid mutator transaction binding the contract method 0x0e81073c.
//
// Solidity: function addShares(address podOwner, uint256 shares) returns(uint256)
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) AddShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.AddShares(&_ContractIEigenPodManager.TransactOpts, podOwner, shares)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) CreatePod() (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.CreatePod(&_ContractIEigenPodManager.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) CreatePod() (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.CreatePod(&_ContractIEigenPodManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.Pause(&_ContractIEigenPodManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.Pause(&_ContractIEigenPodManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.PauseAll(&_ContractIEigenPodManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.PauseAll(&_ContractIEigenPodManager.TransactOpts)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_ContractIEigenPodManager.TransactOpts, podOwner, sharesDelta)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0xc2c51c40.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_ContractIEigenPodManager.TransactOpts, podOwner, sharesDelta)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) RemoveShares(opts *bind.TransactOpts, podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "removeShares", podOwner, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) RemoveShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.RemoveShares(&_ContractIEigenPodManager.TransactOpts, podOwner, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0xbeffbb89.
//
// Solidity: function removeShares(address podOwner, uint256 shares) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) RemoveShares(podOwner common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.RemoveShares(&_ContractIEigenPodManager.TransactOpts, podOwner, shares)
}

// SetDenebForkTimestamp is a paid mutator transaction binding the contract method 0x463db038.
//
// Solidity: function setDenebForkTimestamp(uint64 newDenebForkTimestamp) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) SetDenebForkTimestamp(opts *bind.TransactOpts, newDenebForkTimestamp uint64) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "setDenebForkTimestamp", newDenebForkTimestamp)
}

// SetDenebForkTimestamp is a paid mutator transaction binding the contract method 0x463db038.
//
// Solidity: function setDenebForkTimestamp(uint64 newDenebForkTimestamp) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) SetDenebForkTimestamp(newDenebForkTimestamp uint64) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.SetDenebForkTimestamp(&_ContractIEigenPodManager.TransactOpts, newDenebForkTimestamp)
}

// SetDenebForkTimestamp is a paid mutator transaction binding the contract method 0x463db038.
//
// Solidity: function setDenebForkTimestamp(uint64 newDenebForkTimestamp) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) SetDenebForkTimestamp(newDenebForkTimestamp uint64) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.SetDenebForkTimestamp(&_ContractIEigenPodManager.TransactOpts, newDenebForkTimestamp)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.SetPauserRegistry(&_ContractIEigenPodManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.SetPauserRegistry(&_ContractIEigenPodManager.TransactOpts, newPauserRegistry)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.Stake(&_ContractIEigenPodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.Stake(&_ContractIEigenPodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.Unpause(&_ContractIEigenPodManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.Unpause(&_ContractIEigenPodManager.TransactOpts, newPausedStatus)
}

// UpdateBeaconChainOracle is a paid mutator transaction binding the contract method 0xc1de3aef.
//
// Solidity: function updateBeaconChainOracle(address newBeaconChainOracle) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) UpdateBeaconChainOracle(opts *bind.TransactOpts, newBeaconChainOracle common.Address) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "updateBeaconChainOracle", newBeaconChainOracle)
}

// UpdateBeaconChainOracle is a paid mutator transaction binding the contract method 0xc1de3aef.
//
// Solidity: function updateBeaconChainOracle(address newBeaconChainOracle) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) UpdateBeaconChainOracle(newBeaconChainOracle common.Address) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.UpdateBeaconChainOracle(&_ContractIEigenPodManager.TransactOpts, newBeaconChainOracle)
}

// UpdateBeaconChainOracle is a paid mutator transaction binding the contract method 0xc1de3aef.
//
// Solidity: function updateBeaconChainOracle(address newBeaconChainOracle) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) UpdateBeaconChainOracle(newBeaconChainOracle common.Address) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.UpdateBeaconChainOracle(&_ContractIEigenPodManager.TransactOpts, newBeaconChainOracle)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.contract.Transact(opts, "withdrawSharesAsTokens", podOwner, destination, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerSession) WithdrawSharesAsTokens(podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.WithdrawSharesAsTokens(&_ContractIEigenPodManager.TransactOpts, podOwner, destination, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x387b1300.
//
// Solidity: function withdrawSharesAsTokens(address podOwner, address destination, uint256 shares) returns()
func (_ContractIEigenPodManager *ContractIEigenPodManagerTransactorSession) WithdrawSharesAsTokens(podOwner common.Address, destination common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractIEigenPodManager.Contract.WithdrawSharesAsTokens(&_ContractIEigenPodManager.TransactOpts, podOwner, destination, shares)
}

// ContractIEigenPodManagerBeaconChainETHDepositedIterator is returned from FilterBeaconChainETHDeposited and is used to iterate over the raw logs and unpacked data for BeaconChainETHDeposited events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerBeaconChainETHDepositedIterator struct {
	Event *ContractIEigenPodManagerBeaconChainETHDeposited // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerBeaconChainETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerBeaconChainETHDeposited)
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
		it.Event = new(ContractIEigenPodManagerBeaconChainETHDeposited)
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
func (it *ContractIEigenPodManagerBeaconChainETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerBeaconChainETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerBeaconChainETHDeposited represents a BeaconChainETHDeposited event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerBeaconChainETHDeposited struct {
	PodOwner common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHDeposited is a free log retrieval operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodManagerBeaconChainETHDepositedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerBeaconChainETHDepositedIterator{contract: _ContractIEigenPodManager.contract, event: "BeaconChainETHDeposited", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHDeposited is a free log subscription operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerBeaconChainETHDeposited)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
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

// ParseBeaconChainETHDeposited is a log parse operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParseBeaconChainETHDeposited(log types.Log) (*ContractIEigenPodManagerBeaconChainETHDeposited, error) {
	event := new(ContractIEigenPodManagerBeaconChainETHDeposited)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator is returned from FilterBeaconChainETHWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for BeaconChainETHWithdrawalCompleted events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator struct {
	Event *ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted)
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
		it.Event = new(ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted)
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
func (it *ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted represents a BeaconChainETHWithdrawalCompleted event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted struct {
	PodOwner         common.Address
	Shares           *big.Int
	Nonce            *big.Int
	DelegatedAddress common.Address
	Withdrawer       common.Address
	WithdrawalRoot   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHWithdrawalCompleted is a free log retrieval operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerBeaconChainETHWithdrawalCompletedIterator{contract: _ContractIEigenPodManager.contract, event: "BeaconChainETHWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHWithdrawalCompleted is a free log subscription operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
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

// ParseBeaconChainETHWithdrawalCompleted is a log parse operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted, error) {
	event := new(ContractIEigenPodManagerBeaconChainETHWithdrawalCompleted)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerBeaconOracleUpdatedIterator is returned from FilterBeaconOracleUpdated and is used to iterate over the raw logs and unpacked data for BeaconOracleUpdated events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerBeaconOracleUpdatedIterator struct {
	Event *ContractIEigenPodManagerBeaconOracleUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerBeaconOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerBeaconOracleUpdated)
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
		it.Event = new(ContractIEigenPodManagerBeaconOracleUpdated)
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
func (it *ContractIEigenPodManagerBeaconOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerBeaconOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerBeaconOracleUpdated represents a BeaconOracleUpdated event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerBeaconOracleUpdated struct {
	NewOracleAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBeaconOracleUpdated is a free log retrieval operation binding the contract event 0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6.
//
// Solidity: event BeaconOracleUpdated(address indexed newOracleAddress)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterBeaconOracleUpdated(opts *bind.FilterOpts, newOracleAddress []common.Address) (*ContractIEigenPodManagerBeaconOracleUpdatedIterator, error) {

	var newOracleAddressRule []interface{}
	for _, newOracleAddressItem := range newOracleAddress {
		newOracleAddressRule = append(newOracleAddressRule, newOracleAddressItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "BeaconOracleUpdated", newOracleAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerBeaconOracleUpdatedIterator{contract: _ContractIEigenPodManager.contract, event: "BeaconOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchBeaconOracleUpdated is a free log subscription operation binding the contract event 0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6.
//
// Solidity: event BeaconOracleUpdated(address indexed newOracleAddress)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchBeaconOracleUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerBeaconOracleUpdated, newOracleAddress []common.Address) (event.Subscription, error) {

	var newOracleAddressRule []interface{}
	for _, newOracleAddressItem := range newOracleAddress {
		newOracleAddressRule = append(newOracleAddressRule, newOracleAddressItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "BeaconOracleUpdated", newOracleAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerBeaconOracleUpdated)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "BeaconOracleUpdated", log); err != nil {
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

// ParseBeaconOracleUpdated is a log parse operation binding the contract event 0x08f0470754946ccfbb446ff7fd2d6ae6af1bbdae19f85794c0cc5ed5e8ceb4f6.
//
// Solidity: event BeaconOracleUpdated(address indexed newOracleAddress)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParseBeaconOracleUpdated(log types.Log) (*ContractIEigenPodManagerBeaconOracleUpdated, error) {
	event := new(ContractIEigenPodManagerBeaconOracleUpdated)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "BeaconOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerDenebForkTimestampUpdatedIterator is returned from FilterDenebForkTimestampUpdated and is used to iterate over the raw logs and unpacked data for DenebForkTimestampUpdated events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerDenebForkTimestampUpdatedIterator struct {
	Event *ContractIEigenPodManagerDenebForkTimestampUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerDenebForkTimestampUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerDenebForkTimestampUpdated)
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
		it.Event = new(ContractIEigenPodManagerDenebForkTimestampUpdated)
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
func (it *ContractIEigenPodManagerDenebForkTimestampUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerDenebForkTimestampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerDenebForkTimestampUpdated represents a DenebForkTimestampUpdated event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerDenebForkTimestampUpdated struct {
	NewValue uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDenebForkTimestampUpdated is a free log retrieval operation binding the contract event 0x19200b6fdad58f91b2f496b0c444fc4be3eff74a7e24b07770e04a7137bfd9db.
//
// Solidity: event DenebForkTimestampUpdated(uint64 newValue)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterDenebForkTimestampUpdated(opts *bind.FilterOpts) (*ContractIEigenPodManagerDenebForkTimestampUpdatedIterator, error) {

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "DenebForkTimestampUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerDenebForkTimestampUpdatedIterator{contract: _ContractIEigenPodManager.contract, event: "DenebForkTimestampUpdated", logs: logs, sub: sub}, nil
}

// WatchDenebForkTimestampUpdated is a free log subscription operation binding the contract event 0x19200b6fdad58f91b2f496b0c444fc4be3eff74a7e24b07770e04a7137bfd9db.
//
// Solidity: event DenebForkTimestampUpdated(uint64 newValue)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchDenebForkTimestampUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerDenebForkTimestampUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "DenebForkTimestampUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerDenebForkTimestampUpdated)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "DenebForkTimestampUpdated", log); err != nil {
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

// ParseDenebForkTimestampUpdated is a log parse operation binding the contract event 0x19200b6fdad58f91b2f496b0c444fc4be3eff74a7e24b07770e04a7137bfd9db.
//
// Solidity: event DenebForkTimestampUpdated(uint64 newValue)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParseDenebForkTimestampUpdated(log types.Log) (*ContractIEigenPodManagerDenebForkTimestampUpdated, error) {
	event := new(ContractIEigenPodManagerDenebForkTimestampUpdated)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "DenebForkTimestampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPausedIterator struct {
	Event *ContractIEigenPodManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerPaused)
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
		it.Event = new(ContractIEigenPodManagerPaused)
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
func (it *ContractIEigenPodManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerPaused represents a Paused event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractIEigenPodManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerPausedIterator{contract: _ContractIEigenPodManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerPaused)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParsePaused(log types.Log) (*ContractIEigenPodManagerPaused, error) {
	event := new(ContractIEigenPodManagerPaused)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPauserRegistrySetIterator struct {
	Event *ContractIEigenPodManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerPauserRegistrySet)
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
		it.Event = new(ContractIEigenPodManagerPauserRegistrySet)
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
func (it *ContractIEigenPodManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractIEigenPodManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerPauserRegistrySetIterator{contract: _ContractIEigenPodManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerPauserRegistrySet)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractIEigenPodManagerPauserRegistrySet, error) {
	event := new(ContractIEigenPodManagerPauserRegistrySet)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerPodDeployedIterator is returned from FilterPodDeployed and is used to iterate over the raw logs and unpacked data for PodDeployed events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPodDeployedIterator struct {
	Event *ContractIEigenPodManagerPodDeployed // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerPodDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerPodDeployed)
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
		it.Event = new(ContractIEigenPodManagerPodDeployed)
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
func (it *ContractIEigenPodManagerPodDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerPodDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerPodDeployed represents a PodDeployed event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPodDeployed struct {
	EigenPod common.Address
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPodDeployed is a free log retrieval operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*ContractIEigenPodManagerPodDeployedIterator, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerPodDeployedIterator{contract: _ContractIEigenPodManager.contract, event: "PodDeployed", logs: logs, sub: sub}, nil
}

// WatchPodDeployed is a free log subscription operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerPodDeployed)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
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

// ParsePodDeployed is a log parse operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParsePodDeployed(log types.Log) (*ContractIEigenPodManagerPodDeployed, error) {
	event := new(ContractIEigenPodManagerPodDeployed)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerPodSharesUpdatedIterator is returned from FilterPodSharesUpdated and is used to iterate over the raw logs and unpacked data for PodSharesUpdated events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPodSharesUpdatedIterator struct {
	Event *ContractIEigenPodManagerPodSharesUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerPodSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerPodSharesUpdated)
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
		it.Event = new(ContractIEigenPodManagerPodSharesUpdated)
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
func (it *ContractIEigenPodManagerPodSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerPodSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerPodSharesUpdated represents a PodSharesUpdated event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerPodSharesUpdated struct {
	PodOwner    common.Address
	SharesDelta *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPodSharesUpdated is a free log retrieval operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractIEigenPodManagerPodSharesUpdatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerPodSharesUpdatedIterator{contract: _ContractIEigenPodManager.contract, event: "PodSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPodSharesUpdated is a free log subscription operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerPodSharesUpdated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerPodSharesUpdated)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
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

// ParsePodSharesUpdated is a log parse operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParsePodSharesUpdated(log types.Log) (*ContractIEigenPodManagerPodSharesUpdated, error) {
	event := new(ContractIEigenPodManagerPodSharesUpdated)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIEigenPodManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerUnpausedIterator struct {
	Event *ContractIEigenPodManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractIEigenPodManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIEigenPodManagerUnpaused)
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
		it.Event = new(ContractIEigenPodManagerUnpaused)
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
func (it *ContractIEigenPodManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIEigenPodManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIEigenPodManagerUnpaused represents a Unpaused event raised by the ContractIEigenPodManager contract.
type ContractIEigenPodManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractIEigenPodManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractIEigenPodManagerUnpausedIterator{contract: _ContractIEigenPodManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractIEigenPodManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractIEigenPodManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIEigenPodManagerUnpaused)
				if err := _ContractIEigenPodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractIEigenPodManager *ContractIEigenPodManagerFilterer) ParseUnpaused(log types.Log) (*ContractIEigenPodManagerUnpaused, error) {
	event := new(ContractIEigenPodManagerUnpaused)
	if err := _ContractIEigenPodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
