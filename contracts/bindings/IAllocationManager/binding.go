// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIAllocationManager

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

// IAllocationManagerTypesMagnitudeAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesMagnitudeAllocation struct {
	Strategy             common.Address
	ExpectedMaxMagnitude uint64
	OperatorSets         []OperatorSet
	Magnitudes           []uint64
}

// IAllocationManagerTypesMagnitudeInfo is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesMagnitudeInfo struct {
	CurrentMagnitude uint64
	PendingDiff      *big.Int
	EffectTimestamp  uint32
}

// IAllocationManagerTypesSlashingParams is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerTypesSlashingParams struct {
	Operator      common.Address
	OperatorSetId uint32
	Strategies    []common.Address
	WadToSlash    *big.Int
	Description   string
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// ContractIAllocationManagerMetaData contains all meta data concerning the ContractIAllocationManager contract.
var ContractIAllocationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"clearDeallocationQueue\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSet\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationInfo\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structIAllocationManagerTypes.MagnitudeInfo[][]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationInfo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.MagnitudeInfo[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocationInfo\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.MagnitudeInfo[]\",\"components\":[{\"name\":\"currentMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"pendingDiff\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxMagnitudesAtTimestamp\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinDelegatedAndSlashableOperatorShares\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"beforeTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"},{\"name\":\"\",\"type\":\"uint256[][]\",\"internalType\":\"uint256[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"initialPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManagerTypes.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedMaxMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIAllocationManagerTypes.SlashingParams\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadToSlash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EncumberedMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"encumberedMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"totalMagnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetMagnitudeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"magnitude\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"strategies\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"contractIStrategy[]\"},{\"name\":\"wadSlashed\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"description\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadySlashedForTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllocatableMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAllocationDelay\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpectedTotalMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWadToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotAllocated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
}

// ContractIAllocationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIAllocationManagerMetaData.ABI instead.
var ContractIAllocationManagerABI = ContractIAllocationManagerMetaData.ABI

// ContractIAllocationManagerMethods is an auto generated interface around an Ethereum contract.
type ContractIAllocationManagerMethods interface {
	ContractIAllocationManagerCalls
	ContractIAllocationManagerTransacts
	ContractIAllocationManagerFilters
}

// ContractIAllocationManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIAllocationManagerCalls interface {
	GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (struct {
		IsSet bool
		Delay uint32
	}, error)

	GetAllocationInfo(opts *bind.CallOpts, operatorSet OperatorSet, strategies []common.Address, operators []common.Address) ([][]IAllocationManagerTypesMagnitudeInfo, error)

	GetAllocationInfo0(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]IAllocationManagerTypesMagnitudeInfo, error)

	GetAllocationInfo1(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesMagnitudeInfo, error)

	GetMaxMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error)

	GetMaxMagnitudesAtTimestamp(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error)

	GetMinDelegatedAndSlashableOperatorShares(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, beforeTimestamp uint32) ([][]*big.Int, [][]*big.Int, error)
}

// ContractIAllocationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIAllocationManagerTransacts interface {
	ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error)

	ModifyAllocations(opts *bind.TransactOpts, allocations []IAllocationManagerTypesMagnitudeAllocation) (*types.Transaction, error)

	SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error)

	SetAllocationDelay0(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error)
}

// ContractIAllocationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIAllocationManagerFilters interface {
	FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractIAllocationManagerAllocationDelaySetIterator, error)
	WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerAllocationDelaySet) (event.Subscription, error)
	ParseAllocationDelaySet(log types.Log) (*ContractIAllocationManagerAllocationDelaySet, error)

	FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator, error)
	WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error)
	ParseEncumberedMagnitudeUpdated(log types.Log) (*ContractIAllocationManagerEncumberedMagnitudeUpdated, error)

	FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*ContractIAllocationManagerMaxMagnitudeUpdatedIterator, error)
	WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMaxMagnitudeUpdated) (event.Subscription, error)
	ParseMaxMagnitudeUpdated(log types.Log) (*ContractIAllocationManagerMaxMagnitudeUpdated, error)

	FilterOperatorSetMagnitudeUpdated(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator, error)
	WatchOperatorSetMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerOperatorSetMagnitudeUpdated) (event.Subscription, error)
	ParseOperatorSetMagnitudeUpdated(log types.Log) (*ContractIAllocationManagerOperatorSetMagnitudeUpdated, error)

	FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSlashedIterator, error)
	WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerOperatorSlashed) (event.Subscription, error)
	ParseOperatorSlashed(log types.Log) (*ContractIAllocationManagerOperatorSlashed, error)
}

// ContractIAllocationManager is an auto generated Go binding around an Ethereum contract.
type ContractIAllocationManager struct {
	ContractIAllocationManagerCaller     // Read-only binding to the contract
	ContractIAllocationManagerTransactor // Write-only binding to the contract
	ContractIAllocationManagerFilterer   // Log filterer for contract events
}

// ContractIAllocationManager implements the ContractIAllocationManagerMethods interface.
var _ ContractIAllocationManagerMethods = (*ContractIAllocationManager)(nil)

// ContractIAllocationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIAllocationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAllocationManagerCaller implements the ContractIAllocationManagerCalls interface.
var _ ContractIAllocationManagerCalls = (*ContractIAllocationManagerCaller)(nil)

// ContractIAllocationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIAllocationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAllocationManagerTransactor implements the ContractIAllocationManagerTransacts interface.
var _ ContractIAllocationManagerTransacts = (*ContractIAllocationManagerTransactor)(nil)

// ContractIAllocationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIAllocationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIAllocationManagerFilterer implements the ContractIAllocationManagerFilters interface.
var _ ContractIAllocationManagerFilters = (*ContractIAllocationManagerFilterer)(nil)

// ContractIAllocationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIAllocationManagerSession struct {
	Contract     *ContractIAllocationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractIAllocationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIAllocationManagerCallerSession struct {
	Contract *ContractIAllocationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ContractIAllocationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIAllocationManagerTransactorSession struct {
	Contract     *ContractIAllocationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ContractIAllocationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIAllocationManagerRaw struct {
	Contract *ContractIAllocationManager // Generic contract binding to access the raw methods on
}

// ContractIAllocationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIAllocationManagerCallerRaw struct {
	Contract *ContractIAllocationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIAllocationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIAllocationManagerTransactorRaw struct {
	Contract *ContractIAllocationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIAllocationManager creates a new instance of ContractIAllocationManager, bound to a specific deployed contract.
func NewContractIAllocationManager(address common.Address, backend bind.ContractBackend) (*ContractIAllocationManager, error) {
	contract, err := bindContractIAllocationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManager{ContractIAllocationManagerCaller: ContractIAllocationManagerCaller{contract: contract}, ContractIAllocationManagerTransactor: ContractIAllocationManagerTransactor{contract: contract}, ContractIAllocationManagerFilterer: ContractIAllocationManagerFilterer{contract: contract}}, nil
}

// NewContractIAllocationManagerCaller creates a new read-only instance of ContractIAllocationManager, bound to a specific deployed contract.
func NewContractIAllocationManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractIAllocationManagerCaller, error) {
	contract, err := bindContractIAllocationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerCaller{contract: contract}, nil
}

// NewContractIAllocationManagerTransactor creates a new write-only instance of ContractIAllocationManager, bound to a specific deployed contract.
func NewContractIAllocationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIAllocationManagerTransactor, error) {
	contract, err := bindContractIAllocationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerTransactor{contract: contract}, nil
}

// NewContractIAllocationManagerFilterer creates a new log filterer instance of ContractIAllocationManager, bound to a specific deployed contract.
func NewContractIAllocationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIAllocationManagerFilterer, error) {
	contract, err := bindContractIAllocationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerFilterer{contract: contract}, nil
}

// bindContractIAllocationManager binds a generic wrapper to an already deployed contract.
func bindContractIAllocationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIAllocationManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIAllocationManager *ContractIAllocationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIAllocationManager.Contract.ContractIAllocationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIAllocationManager *ContractIAllocationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ContractIAllocationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIAllocationManager *ContractIAllocationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ContractIAllocationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIAllocationManager *ContractIAllocationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIAllocationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.contract.Transact(opts, method, params...)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getAllocatableMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAllocationManager.Contract.GetAllocatableMagnitude(&_ContractIAllocationManager.CallOpts, operator, strategy)
}

// GetAllocatableMagnitude is a free data retrieval call binding the contract method 0x6cfb4481.
//
// Solidity: function getAllocatableMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetAllocatableMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAllocationManager.Contract.GetAllocatableMagnitude(&_ContractIAllocationManager.CallOpts, operator, strategy)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetAllocationDelay(opts *bind.CallOpts, operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getAllocationDelay", operator)

	outstruct := new(struct {
		IsSet bool
		Delay uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSet = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Delay = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetAllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _ContractIAllocationManager.Contract.GetAllocationDelay(&_ContractIAllocationManager.CallOpts, operator)
}

// GetAllocationDelay is a free data retrieval call binding the contract method 0xb9fbaed1.
//
// Solidity: function getAllocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetAllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _ContractIAllocationManager.Contract.GetAllocationDelay(&_ContractIAllocationManager.CallOpts, operator)
}

// GetAllocationInfo is a free data retrieval call binding the contract method 0x0b002119.
//
// Solidity: function getAllocationInfo((address,uint32) operatorSet, address[] strategies, address[] operators) view returns((uint64,int128,uint32)[][])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetAllocationInfo(opts *bind.CallOpts, operatorSet OperatorSet, strategies []common.Address, operators []common.Address) ([][]IAllocationManagerTypesMagnitudeInfo, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getAllocationInfo", operatorSet, strategies, operators)

	if err != nil {
		return *new([][]IAllocationManagerTypesMagnitudeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([][]IAllocationManagerTypesMagnitudeInfo)).(*[][]IAllocationManagerTypesMagnitudeInfo)

	return out0, err

}

// GetAllocationInfo is a free data retrieval call binding the contract method 0x0b002119.
//
// Solidity: function getAllocationInfo((address,uint32) operatorSet, address[] strategies, address[] operators) view returns((uint64,int128,uint32)[][])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetAllocationInfo(operatorSet OperatorSet, strategies []common.Address, operators []common.Address) ([][]IAllocationManagerTypesMagnitudeInfo, error) {
	return _ContractIAllocationManager.Contract.GetAllocationInfo(&_ContractIAllocationManager.CallOpts, operatorSet, strategies, operators)
}

// GetAllocationInfo is a free data retrieval call binding the contract method 0x0b002119.
//
// Solidity: function getAllocationInfo((address,uint32) operatorSet, address[] strategies, address[] operators) view returns((uint64,int128,uint32)[][])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetAllocationInfo(operatorSet OperatorSet, strategies []common.Address, operators []common.Address) ([][]IAllocationManagerTypesMagnitudeInfo, error) {
	return _ContractIAllocationManager.Contract.GetAllocationInfo(&_ContractIAllocationManager.CallOpts, operatorSet, strategies, operators)
}

// GetAllocationInfo0 is a free data retrieval call binding the contract method 0x35af054a.
//
// Solidity: function getAllocationInfo(address operator, address strategy, (address,uint32)[] operatorSets) view returns((uint64,int128,uint32)[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetAllocationInfo0(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]IAllocationManagerTypesMagnitudeInfo, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getAllocationInfo0", operator, strategy, operatorSets)

	if err != nil {
		return *new([]IAllocationManagerTypesMagnitudeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IAllocationManagerTypesMagnitudeInfo)).(*[]IAllocationManagerTypesMagnitudeInfo)

	return out0, err

}

// GetAllocationInfo0 is a free data retrieval call binding the contract method 0x35af054a.
//
// Solidity: function getAllocationInfo(address operator, address strategy, (address,uint32)[] operatorSets) view returns((uint64,int128,uint32)[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetAllocationInfo0(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]IAllocationManagerTypesMagnitudeInfo, error) {
	return _ContractIAllocationManager.Contract.GetAllocationInfo0(&_ContractIAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetAllocationInfo0 is a free data retrieval call binding the contract method 0x35af054a.
//
// Solidity: function getAllocationInfo(address operator, address strategy, (address,uint32)[] operatorSets) view returns((uint64,int128,uint32)[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetAllocationInfo0(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]IAllocationManagerTypesMagnitudeInfo, error) {
	return _ContractIAllocationManager.Contract.GetAllocationInfo0(&_ContractIAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetAllocationInfo1 is a free data retrieval call binding the contract method 0x4d9dbde9.
//
// Solidity: function getAllocationInfo(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetAllocationInfo1(opts *bind.CallOpts, operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesMagnitudeInfo, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getAllocationInfo1", operator, strategy)

	if err != nil {
		return *new([]OperatorSet), *new([]IAllocationManagerTypesMagnitudeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([]IAllocationManagerTypesMagnitudeInfo)).(*[]IAllocationManagerTypesMagnitudeInfo)

	return out0, out1, err

}

// GetAllocationInfo1 is a free data retrieval call binding the contract method 0x4d9dbde9.
//
// Solidity: function getAllocationInfo(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetAllocationInfo1(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesMagnitudeInfo, error) {
	return _ContractIAllocationManager.Contract.GetAllocationInfo1(&_ContractIAllocationManager.CallOpts, operator, strategy)
}

// GetAllocationInfo1 is a free data retrieval call binding the contract method 0x4d9dbde9.
//
// Solidity: function getAllocationInfo(address operator, address strategy) view returns((address,uint32)[], (uint64,int128,uint32)[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetAllocationInfo1(operator common.Address, strategy common.Address) ([]OperatorSet, []IAllocationManagerTypesMagnitudeInfo, error) {
	return _ContractIAllocationManager.Contract.GetAllocationInfo1(&_ContractIAllocationManager.CallOpts, operator, strategy)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetMaxMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getMaxMagnitudes", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetMaxMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetMaxMagnitudes(&_ContractIAllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudes is a free data retrieval call binding the contract method 0x547afb87.
//
// Solidity: function getMaxMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetMaxMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetMaxMagnitudes(&_ContractIAllocationManager.CallOpts, operator, strategies)
}

// GetMaxMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x843b349f.
//
// Solidity: function getMaxMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetMaxMagnitudesAtTimestamp(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getMaxMagnitudesAtTimestamp", operator, strategies, timestamp)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetMaxMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x843b349f.
//
// Solidity: function getMaxMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetMaxMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetMaxMagnitudesAtTimestamp(&_ContractIAllocationManager.CallOpts, operator, strategies, timestamp)
}

// GetMaxMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x843b349f.
//
// Solidity: function getMaxMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetMaxMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetMaxMagnitudesAtTimestamp(&_ContractIAllocationManager.CallOpts, operator, strategies, timestamp)
}

// GetMinDelegatedAndSlashableOperatorShares is a free data retrieval call binding the contract method 0xe07d2b12.
//
// Solidity: function getMinDelegatedAndSlashableOperatorShares((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 beforeTimestamp) view returns(uint256[][], uint256[][])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetMinDelegatedAndSlashableOperatorShares(opts *bind.CallOpts, operatorSet OperatorSet, operators []common.Address, strategies []common.Address, beforeTimestamp uint32) ([][]*big.Int, [][]*big.Int, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getMinDelegatedAndSlashableOperatorShares", operatorSet, operators, strategies, beforeTimestamp)

	if err != nil {
		return *new([][]*big.Int), *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)
	out1 := *abi.ConvertType(out[1], new([][]*big.Int)).(*[][]*big.Int)

	return out0, out1, err

}

// GetMinDelegatedAndSlashableOperatorShares is a free data retrieval call binding the contract method 0xe07d2b12.
//
// Solidity: function getMinDelegatedAndSlashableOperatorShares((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 beforeTimestamp) view returns(uint256[][], uint256[][])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetMinDelegatedAndSlashableOperatorShares(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, beforeTimestamp uint32) ([][]*big.Int, [][]*big.Int, error) {
	return _ContractIAllocationManager.Contract.GetMinDelegatedAndSlashableOperatorShares(&_ContractIAllocationManager.CallOpts, operatorSet, operators, strategies, beforeTimestamp)
}

// GetMinDelegatedAndSlashableOperatorShares is a free data retrieval call binding the contract method 0xe07d2b12.
//
// Solidity: function getMinDelegatedAndSlashableOperatorShares((address,uint32) operatorSet, address[] operators, address[] strategies, uint32 beforeTimestamp) view returns(uint256[][], uint256[][])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetMinDelegatedAndSlashableOperatorShares(operatorSet OperatorSet, operators []common.Address, strategies []common.Address, beforeTimestamp uint32) ([][]*big.Int, [][]*big.Int, error) {
	return _ContractIAllocationManager.Contract.GetMinDelegatedAndSlashableOperatorShares(&_ContractIAllocationManager.CallOpts, operatorSet, operators, strategies, beforeTimestamp)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) ClearDeallocationQueue(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "clearDeallocationQueue", operator, strategies, numToComplete)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ClearDeallocationQueue(&_ContractIAllocationManager.TransactOpts, operator, strategies, numToComplete)
}

// ClearDeallocationQueue is a paid mutator transaction binding the contract method 0x4b5046ef.
//
// Solidity: function clearDeallocationQueue(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) ClearDeallocationQueue(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ClearDeallocationQueue(&_ContractIAllocationManager.TransactOpts, operator, strategies, numToComplete)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.Initialize(&_ContractIAllocationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 initialPausedStatus) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, initialPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.Initialize(&_ContractIAllocationManager.TransactOpts, initialOwner, _pauserRegistry, initialPausedStatus)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x1637b60f.
//
// Solidity: function modifyAllocations((address,uint64,(address,uint32)[],uint64[])[] allocations) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, allocations []IAllocationManagerTypesMagnitudeAllocation) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "modifyAllocations", allocations)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x1637b60f.
//
// Solidity: function modifyAllocations((address,uint64,(address,uint32)[],uint64[])[] allocations) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) ModifyAllocations(allocations []IAllocationManagerTypesMagnitudeAllocation) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ModifyAllocations(&_ContractIAllocationManager.TransactOpts, allocations)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0x1637b60f.
//
// Solidity: function modifyAllocations((address,uint64,(address,uint32)[],uint64[])[] allocations) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) ModifyAllocations(allocations []IAllocationManagerTypesMagnitudeAllocation) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ModifyAllocations(&_ContractIAllocationManager.TransactOpts, allocations)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "setAllocationDelay", operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SetAllocationDelay(&_ContractIAllocationManager.TransactOpts, operator, delay)
}

// SetAllocationDelay is a paid mutator transaction binding the contract method 0x56c483e6.
//
// Solidity: function setAllocationDelay(address operator, uint32 delay) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) SetAllocationDelay(operator common.Address, delay uint32) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SetAllocationDelay(&_ContractIAllocationManager.TransactOpts, operator, delay)
}

// SetAllocationDelay0 is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) SetAllocationDelay0(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "setAllocationDelay0", delay)
}

// SetAllocationDelay0 is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) SetAllocationDelay0(delay uint32) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SetAllocationDelay0(&_ContractIAllocationManager.TransactOpts, delay)
}

// SetAllocationDelay0 is a paid mutator transaction binding the contract method 0x5c489bb5.
//
// Solidity: function setAllocationDelay(uint32 delay) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) SetAllocationDelay0(delay uint32) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SetAllocationDelay0(&_ContractIAllocationManager.TransactOpts, delay)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x60db99a3.
//
// Solidity: function slashOperator((address,uint32,address[],uint256,string) params) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "slashOperator", params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x60db99a3.
//
// Solidity: function slashOperator((address,uint32,address[],uint256,string) params) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) SlashOperator(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SlashOperator(&_ContractIAllocationManager.TransactOpts, params)
}

// SlashOperator is a paid mutator transaction binding the contract method 0x60db99a3.
//
// Solidity: function slashOperator((address,uint32,address[],uint256,string) params) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) SlashOperator(params IAllocationManagerTypesSlashingParams) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SlashOperator(&_ContractIAllocationManager.TransactOpts, params)
}

// ContractIAllocationManagerAllocationDelaySetIterator is returned from FilterAllocationDelaySet and is used to iterate over the raw logs and unpacked data for AllocationDelaySet events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerAllocationDelaySetIterator struct {
	Event *ContractIAllocationManagerAllocationDelaySet // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerAllocationDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerAllocationDelaySet)
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
		it.Event = new(ContractIAllocationManagerAllocationDelaySet)
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
func (it *ContractIAllocationManagerAllocationDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerAllocationDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerAllocationDelaySet represents a AllocationDelaySet event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerAllocationDelaySet struct {
	Operator        common.Address
	Delay           uint32
	EffectTimestamp uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractIAllocationManagerAllocationDelaySetIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerAllocationDelaySetIterator{contract: _ContractIAllocationManager.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerAllocationDelaySet) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerAllocationDelaySet)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
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

// ParseAllocationDelaySet is a log parse operation binding the contract event 0x4e85751d6331506c6c62335f207eb31f12a61e570f34f5c17640308785c6d4db.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseAllocationDelaySet(log types.Log) (*ContractIAllocationManagerAllocationDelaySet, error) {
	event := new(ContractIAllocationManagerAllocationDelaySet)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator is returned from FilterEncumberedMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for EncumberedMagnitudeUpdated events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator struct {
	Event *ContractIAllocationManagerEncumberedMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerEncumberedMagnitudeUpdated)
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
		it.Event = new(ContractIAllocationManagerEncumberedMagnitudeUpdated)
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
func (it *ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerEncumberedMagnitudeUpdated represents a EncumberedMagnitudeUpdated event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerEncumberedMagnitudeUpdated struct {
	Operator            common.Address
	Strategy            common.Address
	EncumberedMagnitude uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterEncumberedMagnitudeUpdated is a free log retrieval operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterEncumberedMagnitudeUpdated(opts *bind.FilterOpts) (*ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerEncumberedMagnitudeUpdatedIterator{contract: _ContractIAllocationManager.contract, event: "EncumberedMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchEncumberedMagnitudeUpdated is a free log subscription operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchEncumberedMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerEncumberedMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "EncumberedMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerEncumberedMagnitudeUpdated)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
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

// ParseEncumberedMagnitudeUpdated is a log parse operation binding the contract event 0xacf9095feb3a370c9cf692421c69ef320d4db5c66e6a7d29c7694eb02364fc55.
//
// Solidity: event EncumberedMagnitudeUpdated(address operator, address strategy, uint64 encumberedMagnitude)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseEncumberedMagnitudeUpdated(log types.Log) (*ContractIAllocationManagerEncumberedMagnitudeUpdated, error) {
	event := new(ContractIAllocationManagerEncumberedMagnitudeUpdated)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "EncumberedMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerMaxMagnitudeUpdatedIterator is returned from FilterMaxMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for MaxMagnitudeUpdated events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMaxMagnitudeUpdatedIterator struct {
	Event *ContractIAllocationManagerMaxMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerMaxMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerMaxMagnitudeUpdated)
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
		it.Event = new(ContractIAllocationManagerMaxMagnitudeUpdated)
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
func (it *ContractIAllocationManagerMaxMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerMaxMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerMaxMagnitudeUpdated represents a MaxMagnitudeUpdated event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMaxMagnitudeUpdated struct {
	Operator       common.Address
	Strategy       common.Address
	TotalMagnitude uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaxMagnitudeUpdated is a free log retrieval operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterMaxMagnitudeUpdated(opts *bind.FilterOpts) (*ContractIAllocationManagerMaxMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerMaxMagnitudeUpdatedIterator{contract: _ContractIAllocationManager.contract, event: "MaxMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMagnitudeUpdated is a free log subscription operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchMaxMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMaxMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "MaxMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerMaxMagnitudeUpdated)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
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

// ParseMaxMagnitudeUpdated is a log parse operation binding the contract event 0x1c6458079a41077d003c11faf9bf097e693bd67979e4e6500bac7b29db779b5c.
//
// Solidity: event MaxMagnitudeUpdated(address operator, address strategy, uint64 totalMagnitude)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseMaxMagnitudeUpdated(log types.Log) (*ContractIAllocationManagerMaxMagnitudeUpdated, error) {
	event := new(ContractIAllocationManagerMaxMagnitudeUpdated)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "MaxMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator is returned from FilterOperatorSetMagnitudeUpdated and is used to iterate over the raw logs and unpacked data for OperatorSetMagnitudeUpdated events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator struct {
	Event *ContractIAllocationManagerOperatorSetMagnitudeUpdated // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerOperatorSetMagnitudeUpdated)
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
		it.Event = new(ContractIAllocationManagerOperatorSetMagnitudeUpdated)
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
func (it *ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerOperatorSetMagnitudeUpdated represents a OperatorSetMagnitudeUpdated event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerOperatorSetMagnitudeUpdated struct {
	Operator        common.Address
	OperatorSet     OperatorSet
	Strategy        common.Address
	Magnitude       uint64
	EffectTimestamp uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetMagnitudeUpdated is a free log retrieval operation binding the contract event 0x8b997e53d7b9e5d923d0a21c60df81e1740860d1a8c66b8c63c5047ae20eaaf6.
//
// Solidity: event OperatorSetMagnitudeUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterOperatorSetMagnitudeUpdated(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "OperatorSetMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerOperatorSetMagnitudeUpdatedIterator{contract: _ContractIAllocationManager.contract, event: "OperatorSetMagnitudeUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetMagnitudeUpdated is a free log subscription operation binding the contract event 0x8b997e53d7b9e5d923d0a21c60df81e1740860d1a8c66b8c63c5047ae20eaaf6.
//
// Solidity: event OperatorSetMagnitudeUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchOperatorSetMagnitudeUpdated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerOperatorSetMagnitudeUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "OperatorSetMagnitudeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerOperatorSetMagnitudeUpdated)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSetMagnitudeUpdated", log); err != nil {
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

// ParseOperatorSetMagnitudeUpdated is a log parse operation binding the contract event 0x8b997e53d7b9e5d923d0a21c60df81e1740860d1a8c66b8c63c5047ae20eaaf6.
//
// Solidity: event OperatorSetMagnitudeUpdated(address operator, (address,uint32) operatorSet, address strategy, uint64 magnitude, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseOperatorSetMagnitudeUpdated(log types.Log) (*ContractIAllocationManagerOperatorSetMagnitudeUpdated, error) {
	event := new(ContractIAllocationManagerOperatorSetMagnitudeUpdated)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSetMagnitudeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerOperatorSlashedIterator is returned from FilterOperatorSlashed and is used to iterate over the raw logs and unpacked data for OperatorSlashed events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerOperatorSlashedIterator struct {
	Event *ContractIAllocationManagerOperatorSlashed // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerOperatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerOperatorSlashed)
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
		it.Event = new(ContractIAllocationManagerOperatorSlashed)
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
func (it *ContractIAllocationManagerOperatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerOperatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerOperatorSlashed represents a OperatorSlashed event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerOperatorSlashed struct {
	Operator    common.Address
	OperatorSet OperatorSet
	Strategies  []common.Address
	WadSlashed  []*big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSlashedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerOperatorSlashedIterator{contract: _ContractIAllocationManager.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchOperatorSlashed(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerOperatorSlashed) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerOperatorSlashed)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0x80969ad29428d6797ee7aad084f9e4a42a82fc506dcd2ca3b6fb431f85ccebe5.
//
// Solidity: event OperatorSlashed(address operator, (address,uint32) operatorSet, address[] strategies, uint256[] wadSlashed, string description)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseOperatorSlashed(log types.Log) (*ContractIAllocationManagerOperatorSlashed, error) {
	event := new(ContractIAllocationManagerOperatorSlashed)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
