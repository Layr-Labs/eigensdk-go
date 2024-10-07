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

// IAllocationManagerMagnitudeAllocation is an auto generated low-level Go binding around an user-defined struct.
type IAllocationManagerMagnitudeAllocation struct {
	Strategy               common.Address
	ExpectedTotalMagnitude uint64
	OperatorSets           []OperatorSet
	Magnitudes             []uint64
}

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// OperatorSet is an auto generated low-level Go binding around an user-defined struct.
type OperatorSet struct {
	Avs           common.Address
	OperatorSetId uint32
}

// ContractIAllocationManagerMetaData contains all meta data concerning the ContractIAllocationManager contract.
var ContractIAllocationManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"isSet\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateMagnitudeAllocationDigestHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManager.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelSalt\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completePendingDeallocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"numToComplete\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"domainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllocatableMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPendingDeallocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashableMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"\",\"type\":\"uint64[][]\",\"internalType\":\"uint64[][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitude\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudes\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalMagnitudesAtTimestamp\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"modifyAllocations\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocations\",\"type\":\"tuple[]\",\"internalType\":\"structIAllocationManager.MagnitudeAllocation[]\",\"components\":[{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"expectedTotalMagnitude\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"operatorSets\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorSet[]\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudes\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]},{\"name\":\"operatorSignature\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAllocationDelay\",\"inputs\":[{\"name\":\"delay\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllocationDelaySet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delay\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeAllocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToAllocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"effectTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeDeallocationCompleted\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"freeMagnitudeAdded\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MagnitudeQueueDeallocated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"magnitudeToDeallocate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"completableTimestamp\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSetCreated\",\"inputs\":[{\"name\":\"operatorSet\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOperatorSet\",\"components\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSlashed\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorSetId\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"strategy\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIStrategy\"},{\"name\":\"bipsToSlash\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InputArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAllocatableMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBipsToSlash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelay\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExpectedTotalMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperatorSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ModificationAlreadyPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SaltSpent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SameMagnitude\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UninitializedAllocationDelay\",\"inputs\":[]}]",
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
	AllocationDelay(opts *bind.CallOpts, operator common.Address) (struct {
		IsSet bool
		Delay uint32
	}, error)

	CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error)

	DomainSeparator(opts *bind.CallOpts) ([32]byte, error)

	GetAllocatableMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetPendingAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error)

	GetPendingDeallocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error)

	GetSlashableMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error)

	GetTotalMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error)

	GetTotalMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error)

	GetTotalMagnitudesAtTimestamp(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error)
}

// ContractIAllocationManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIAllocationManagerTransacts interface {
	CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error)

	CompletePendingDeallocations(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error)

	ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error)

	SetAllocationDelay(opts *bind.TransactOpts, operator common.Address, delay uint32) (*types.Transaction, error)

	SetAllocationDelay0(opts *bind.TransactOpts, delay uint32) (*types.Transaction, error)

	SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error)
}

// ContractIAllocationManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIAllocationManagerFilters interface {
	FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractIAllocationManagerAllocationDelaySetIterator, error)
	WatchAllocationDelaySet(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerAllocationDelaySet) (event.Subscription, error)
	ParseAllocationDelaySet(log types.Log) (*ContractIAllocationManagerAllocationDelaySet, error)

	FilterMagnitudeAllocated(opts *bind.FilterOpts) (*ContractIAllocationManagerMagnitudeAllocatedIterator, error)
	WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMagnitudeAllocated) (event.Subscription, error)
	ParseMagnitudeAllocated(log types.Log) (*ContractIAllocationManagerMagnitudeAllocated, error)

	FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*ContractIAllocationManagerMagnitudeDeallocationCompletedIterator, error)
	WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMagnitudeDeallocationCompleted) (event.Subscription, error)
	ParseMagnitudeDeallocationCompleted(log types.Log) (*ContractIAllocationManagerMagnitudeDeallocationCompleted, error)

	FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*ContractIAllocationManagerMagnitudeQueueDeallocatedIterator, error)
	WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMagnitudeQueueDeallocated) (event.Subscription, error)
	ParseMagnitudeQueueDeallocated(log types.Log) (*ContractIAllocationManagerMagnitudeQueueDeallocated, error)

	FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSetCreatedIterator, error)
	WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerOperatorSetCreated) (event.Subscription, error)
	ParseOperatorSetCreated(log types.Log) (*ContractIAllocationManagerOperatorSetCreated, error)

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

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) AllocationDelay(opts *bind.CallOpts, operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "allocationDelay", operator)

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

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerSession) AllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _ContractIAllocationManager.Contract.AllocationDelay(&_ContractIAllocationManager.CallOpts, operator)
}

// AllocationDelay is a free data retrieval call binding the contract method 0xbe4e1fd3.
//
// Solidity: function allocationDelay(address operator) view returns(bool isSet, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) AllocationDelay(operator common.Address) (struct {
	IsSet bool
	Delay uint32
}, error) {
	return _ContractIAllocationManager.Contract.AllocationDelay(&_ContractIAllocationManager.CallOpts, operator)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) CalculateMagnitudeAllocationDigestHash(opts *bind.CallOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "calculateMagnitudeAllocationDigestHash", operator, allocations, salt, expiry)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAllocationManager *ContractIAllocationManagerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAllocationManager.Contract.CalculateMagnitudeAllocationDigestHash(&_ContractIAllocationManager.CallOpts, operator, allocations, salt, expiry)
}

// CalculateMagnitudeAllocationDigestHash is a free data retrieval call binding the contract method 0x686b686e.
//
// Solidity: function calculateMagnitudeAllocationDigestHash(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, bytes32 salt, uint256 expiry) view returns(bytes32)
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) CalculateMagnitudeAllocationDigestHash(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, salt [32]byte, expiry *big.Int) ([32]byte, error) {
	return _ContractIAllocationManager.Contract.CalculateMagnitudeAllocationDigestHash(&_ContractIAllocationManager.CallOpts, operator, allocations, salt, expiry)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAllocationManager *ContractIAllocationManagerSession) DomainSeparator() ([32]byte, error) {
	return _ContractIAllocationManager.Contract.DomainSeparator(&_ContractIAllocationManager.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) DomainSeparator() ([32]byte, error) {
	return _ContractIAllocationManager.Contract.DomainSeparator(&_ContractIAllocationManager.CallOpts)
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

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetPendingAllocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getPendingAllocations", operator, strategy, operatorSets)

	if err != nil {
		return *new([]uint64), *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]uint32)).(*[]uint32)

	return out0, out1, err

}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetPendingAllocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _ContractIAllocationManager.Contract.GetPendingAllocations(&_ContractIAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetPendingAllocations is a free data retrieval call binding the contract method 0x67d973ef.
//
// Solidity: function getPendingAllocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetPendingAllocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _ContractIAllocationManager.Contract.GetPendingAllocations(&_ContractIAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetPendingDeallocations(opts *bind.CallOpts, operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getPendingDeallocations", operator, strategy, operatorSets)

	if err != nil {
		return *new([]uint64), *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	out1 := *abi.ConvertType(out[1], new([]uint32)).(*[]uint32)

	return out0, out1, err

}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetPendingDeallocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _ContractIAllocationManager.Contract.GetPendingDeallocations(&_ContractIAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetPendingDeallocations is a free data retrieval call binding the contract method 0x07053717.
//
// Solidity: function getPendingDeallocations(address operator, address strategy, (address,uint32)[] operatorSets) view returns(uint64[], uint32[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetPendingDeallocations(operator common.Address, strategy common.Address, operatorSets []OperatorSet) ([]uint64, []uint32, error) {
	return _ContractIAllocationManager.Contract.GetPendingDeallocations(&_ContractIAllocationManager.CallOpts, operator, strategy, operatorSets)
}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x25d8fad7.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetSlashableMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getSlashableMagnitudes", operator, strategies)

	if err != nil {
		return *new([]OperatorSet), *new([][]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]OperatorSet)).(*[]OperatorSet)
	out1 := *abi.ConvertType(out[1], new([][]uint64)).(*[][]uint64)

	return out0, out1, err

}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x25d8fad7.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetSlashableMagnitudes(operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	return _ContractIAllocationManager.Contract.GetSlashableMagnitudes(&_ContractIAllocationManager.CallOpts, operator, strategies)
}

// GetSlashableMagnitudes is a free data retrieval call binding the contract method 0x25d8fad7.
//
// Solidity: function getSlashableMagnitudes(address operator, address[] strategies) view returns((address,uint32)[], uint64[][])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetSlashableMagnitudes(operator common.Address, strategies []common.Address) ([]OperatorSet, [][]uint64, error) {
	return _ContractIAllocationManager.Contract.GetSlashableMagnitudes(&_ContractIAllocationManager.CallOpts, operator, strategies)
}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetTotalMagnitude(opts *bind.CallOpts, operator common.Address, strategy common.Address) (uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getTotalMagnitude", operator, strategy)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAllocationManager.Contract.GetTotalMagnitude(&_ContractIAllocationManager.CallOpts, operator, strategy)
}

// GetTotalMagnitude is a free data retrieval call binding the contract method 0xb47265e2.
//
// Solidity: function getTotalMagnitude(address operator, address strategy) view returns(uint64)
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetTotalMagnitude(operator common.Address, strategy common.Address) (uint64, error) {
	return _ContractIAllocationManager.Contract.GetTotalMagnitude(&_ContractIAllocationManager.CallOpts, operator, strategy)
}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetTotalMagnitudes(opts *bind.CallOpts, operator common.Address, strategies []common.Address) ([]uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getTotalMagnitudes", operator, strategies)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetTotalMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetTotalMagnitudes(&_ContractIAllocationManager.CallOpts, operator, strategies)
}

// GetTotalMagnitudes is a free data retrieval call binding the contract method 0x39a9a3ed.
//
// Solidity: function getTotalMagnitudes(address operator, address[] strategies) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetTotalMagnitudes(operator common.Address, strategies []common.Address) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetTotalMagnitudes(&_ContractIAllocationManager.CallOpts, operator, strategies)
}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCaller) GetTotalMagnitudesAtTimestamp(opts *bind.CallOpts, operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	var out []interface{}
	err := _ContractIAllocationManager.contract.Call(opts, &out, "getTotalMagnitudesAtTimestamp", operator, strategies, timestamp)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerSession) GetTotalMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetTotalMagnitudesAtTimestamp(&_ContractIAllocationManager.CallOpts, operator, strategies, timestamp)
}

// GetTotalMagnitudesAtTimestamp is a free data retrieval call binding the contract method 0x858d0b47.
//
// Solidity: function getTotalMagnitudesAtTimestamp(address operator, address[] strategies, uint32 timestamp) view returns(uint64[])
func (_ContractIAllocationManager *ContractIAllocationManagerCallerSession) GetTotalMagnitudesAtTimestamp(operator common.Address, strategies []common.Address, timestamp uint32) ([]uint64, error) {
	return _ContractIAllocationManager.Contract.GetTotalMagnitudesAtTimestamp(&_ContractIAllocationManager.CallOpts, operator, strategies, timestamp)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) CancelSalt(opts *bind.TransactOpts, salt [32]byte) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "cancelSalt", salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.CancelSalt(&_ContractIAllocationManager.TransactOpts, salt)
}

// CancelSalt is a paid mutator transaction binding the contract method 0xec76f442.
//
// Solidity: function cancelSalt(bytes32 salt) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) CancelSalt(salt [32]byte) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.CancelSalt(&_ContractIAllocationManager.TransactOpts, salt)
}

// CompletePendingDeallocations is a paid mutator transaction binding the contract method 0x073c01ba.
//
// Solidity: function completePendingDeallocations(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) CompletePendingDeallocations(opts *bind.TransactOpts, operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "completePendingDeallocations", operator, strategies, numToComplete)
}

// CompletePendingDeallocations is a paid mutator transaction binding the contract method 0x073c01ba.
//
// Solidity: function completePendingDeallocations(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) CompletePendingDeallocations(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.CompletePendingDeallocations(&_ContractIAllocationManager.TransactOpts, operator, strategies, numToComplete)
}

// CompletePendingDeallocations is a paid mutator transaction binding the contract method 0x073c01ba.
//
// Solidity: function completePendingDeallocations(address operator, address[] strategies, uint16[] numToComplete) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) CompletePendingDeallocations(operator common.Address, strategies []common.Address, numToComplete []uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.CompletePendingDeallocations(&_ContractIAllocationManager.TransactOpts, operator, strategies, numToComplete)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) ModifyAllocations(opts *bind.TransactOpts, operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "modifyAllocations", operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) ModifyAllocations(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ModifyAllocations(&_ContractIAllocationManager.TransactOpts, operator, allocations, operatorSignature)
}

// ModifyAllocations is a paid mutator transaction binding the contract method 0xf8e91d16.
//
// Solidity: function modifyAllocations(address operator, (address,uint64,(address,uint32)[],uint64[])[] allocations, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) ModifyAllocations(operator common.Address, allocations []IAllocationManagerMagnitudeAllocation, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.ModifyAllocations(&_ContractIAllocationManager.TransactOpts, operator, allocations, operatorSignature)
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

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactor) SlashOperator(opts *bind.TransactOpts, operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.contract.Transact(opts, "slashOperator", operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SlashOperator(&_ContractIAllocationManager.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
}

// SlashOperator is a paid mutator transaction binding the contract method 0xbd74a06c.
//
// Solidity: function slashOperator(address operator, uint32 operatorSetId, address[] strategies, uint16 bipsToSlash) returns()
func (_ContractIAllocationManager *ContractIAllocationManagerTransactorSession) SlashOperator(operator common.Address, operatorSetId uint32, strategies []common.Address, bipsToSlash uint16) (*types.Transaction, error) {
	return _ContractIAllocationManager.Contract.SlashOperator(&_ContractIAllocationManager.TransactOpts, operator, operatorSetId, strategies, bipsToSlash)
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
	Operator common.Address
	Delay    uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllocationDelaySet is a free log retrieval operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterAllocationDelaySet(opts *bind.FilterOpts) (*ContractIAllocationManagerAllocationDelaySetIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "AllocationDelaySet")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerAllocationDelaySetIterator{contract: _ContractIAllocationManager.contract, event: "AllocationDelaySet", logs: logs, sub: sub}, nil
}

// WatchAllocationDelaySet is a free log subscription operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
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

// ParseAllocationDelaySet is a log parse operation binding the contract event 0xc17479bf29bcb9669d4dd3580ba716c0424d52c939d248d49b07efc02a32952d.
//
// Solidity: event AllocationDelaySet(address operator, uint32 delay)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseAllocationDelaySet(log types.Log) (*ContractIAllocationManagerAllocationDelaySet, error) {
	event := new(ContractIAllocationManagerAllocationDelaySet)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "AllocationDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerMagnitudeAllocatedIterator is returned from FilterMagnitudeAllocated and is used to iterate over the raw logs and unpacked data for MagnitudeAllocated events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMagnitudeAllocatedIterator struct {
	Event *ContractIAllocationManagerMagnitudeAllocated // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerMagnitudeAllocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerMagnitudeAllocated)
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
		it.Event = new(ContractIAllocationManagerMagnitudeAllocated)
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
func (it *ContractIAllocationManagerMagnitudeAllocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerMagnitudeAllocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerMagnitudeAllocated represents a MagnitudeAllocated event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMagnitudeAllocated struct {
	Operator            common.Address
	Strategy            common.Address
	OperatorSet         OperatorSet
	MagnitudeToAllocate uint64
	EffectTimestamp     uint32
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeAllocated is a free log retrieval operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterMagnitudeAllocated(opts *bind.FilterOpts) (*ContractIAllocationManagerMagnitudeAllocatedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerMagnitudeAllocatedIterator{contract: _ContractIAllocationManager.contract, event: "MagnitudeAllocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeAllocated is a free log subscription operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchMagnitudeAllocated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMagnitudeAllocated) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "MagnitudeAllocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerMagnitudeAllocated)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
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

// ParseMagnitudeAllocated is a log parse operation binding the contract event 0x6d7d0079582cb2c5e70d4135b37f36711415ee6c260778b716bd65e026eb4f1a.
//
// Solidity: event MagnitudeAllocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToAllocate, uint32 effectTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseMagnitudeAllocated(log types.Log) (*ContractIAllocationManagerMagnitudeAllocated, error) {
	event := new(ContractIAllocationManagerMagnitudeAllocated)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "MagnitudeAllocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerMagnitudeDeallocationCompletedIterator is returned from FilterMagnitudeDeallocationCompleted and is used to iterate over the raw logs and unpacked data for MagnitudeDeallocationCompleted events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMagnitudeDeallocationCompletedIterator struct {
	Event *ContractIAllocationManagerMagnitudeDeallocationCompleted // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerMagnitudeDeallocationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerMagnitudeDeallocationCompleted)
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
		it.Event = new(ContractIAllocationManagerMagnitudeDeallocationCompleted)
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
func (it *ContractIAllocationManagerMagnitudeDeallocationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerMagnitudeDeallocationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerMagnitudeDeallocationCompleted represents a MagnitudeDeallocationCompleted event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMagnitudeDeallocationCompleted struct {
	Operator           common.Address
	Strategy           common.Address
	OperatorSet        OperatorSet
	FreeMagnitudeAdded uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeDeallocationCompleted is a free log retrieval operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterMagnitudeDeallocationCompleted(opts *bind.FilterOpts) (*ContractIAllocationManagerMagnitudeDeallocationCompletedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerMagnitudeDeallocationCompletedIterator{contract: _ContractIAllocationManager.contract, event: "MagnitudeDeallocationCompleted", logs: logs, sub: sub}, nil
}

// WatchMagnitudeDeallocationCompleted is a free log subscription operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchMagnitudeDeallocationCompleted(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMagnitudeDeallocationCompleted) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "MagnitudeDeallocationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerMagnitudeDeallocationCompleted)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
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

// ParseMagnitudeDeallocationCompleted is a log parse operation binding the contract event 0x1e5c8e13c62c31d6252ac205e592477d643c7e95831d5b46d99a3c60c2fad8db.
//
// Solidity: event MagnitudeDeallocationCompleted(address operator, address strategy, (address,uint32) operatorSet, uint64 freeMagnitudeAdded)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseMagnitudeDeallocationCompleted(log types.Log) (*ContractIAllocationManagerMagnitudeDeallocationCompleted, error) {
	event := new(ContractIAllocationManagerMagnitudeDeallocationCompleted)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "MagnitudeDeallocationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerMagnitudeQueueDeallocatedIterator is returned from FilterMagnitudeQueueDeallocated and is used to iterate over the raw logs and unpacked data for MagnitudeQueueDeallocated events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMagnitudeQueueDeallocatedIterator struct {
	Event *ContractIAllocationManagerMagnitudeQueueDeallocated // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerMagnitudeQueueDeallocatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerMagnitudeQueueDeallocated)
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
		it.Event = new(ContractIAllocationManagerMagnitudeQueueDeallocated)
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
func (it *ContractIAllocationManagerMagnitudeQueueDeallocatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerMagnitudeQueueDeallocatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerMagnitudeQueueDeallocated represents a MagnitudeQueueDeallocated event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerMagnitudeQueueDeallocated struct {
	Operator              common.Address
	Strategy              common.Address
	OperatorSet           OperatorSet
	MagnitudeToDeallocate uint64
	CompletableTimestamp  uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMagnitudeQueueDeallocated is a free log retrieval operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterMagnitudeQueueDeallocated(opts *bind.FilterOpts) (*ContractIAllocationManagerMagnitudeQueueDeallocatedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerMagnitudeQueueDeallocatedIterator{contract: _ContractIAllocationManager.contract, event: "MagnitudeQueueDeallocated", logs: logs, sub: sub}, nil
}

// WatchMagnitudeQueueDeallocated is a free log subscription operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchMagnitudeQueueDeallocated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerMagnitudeQueueDeallocated) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "MagnitudeQueueDeallocated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerMagnitudeQueueDeallocated)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
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

// ParseMagnitudeQueueDeallocated is a log parse operation binding the contract event 0x2e68db1fe51107d7e451ae268d1631796989ab9d7925054e9b247854cb5be950.
//
// Solidity: event MagnitudeQueueDeallocated(address operator, address strategy, (address,uint32) operatorSet, uint64 magnitudeToDeallocate, uint32 completableTimestamp)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseMagnitudeQueueDeallocated(log types.Log) (*ContractIAllocationManagerMagnitudeQueueDeallocated, error) {
	event := new(ContractIAllocationManagerMagnitudeQueueDeallocated)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "MagnitudeQueueDeallocated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIAllocationManagerOperatorSetCreatedIterator is returned from FilterOperatorSetCreated and is used to iterate over the raw logs and unpacked data for OperatorSetCreated events raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerOperatorSetCreatedIterator struct {
	Event *ContractIAllocationManagerOperatorSetCreated // Event containing the contract specifics and raw log

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
func (it *ContractIAllocationManagerOperatorSetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIAllocationManagerOperatorSetCreated)
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
		it.Event = new(ContractIAllocationManagerOperatorSetCreated)
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
func (it *ContractIAllocationManagerOperatorSetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIAllocationManagerOperatorSetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIAllocationManagerOperatorSetCreated represents a OperatorSetCreated event raised by the ContractIAllocationManager contract.
type ContractIAllocationManagerOperatorSetCreated struct {
	OperatorSet OperatorSet
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorSetCreated is a free log retrieval operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterOperatorSetCreated(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSetCreatedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerOperatorSetCreatedIterator{contract: _ContractIAllocationManager.contract, event: "OperatorSetCreated", logs: logs, sub: sub}, nil
}

// WatchOperatorSetCreated is a free log subscription operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) WatchOperatorSetCreated(opts *bind.WatchOpts, sink chan<- *ContractIAllocationManagerOperatorSetCreated) (event.Subscription, error) {

	logs, sub, err := _ContractIAllocationManager.contract.WatchLogs(opts, "OperatorSetCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIAllocationManagerOperatorSetCreated)
				if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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

// ParseOperatorSetCreated is a log parse operation binding the contract event 0x31629285ead2335ae0933f86ed2ae63321f7af77b4e6eaabc42c057880977e6c.
//
// Solidity: event OperatorSetCreated((address,uint32) operatorSet)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseOperatorSetCreated(log types.Log) (*ContractIAllocationManagerOperatorSetCreated, error) {
	event := new(ContractIAllocationManagerOperatorSetCreated)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSetCreated", log); err != nil {
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
	Operator      common.Address
	OperatorSetId uint32
	Strategy      common.Address
	BipsToSlash   uint16
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorSlashed is a free log retrieval operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) FilterOperatorSlashed(opts *bind.FilterOpts) (*ContractIAllocationManagerOperatorSlashedIterator, error) {

	logs, sub, err := _ContractIAllocationManager.contract.FilterLogs(opts, "OperatorSlashed")
	if err != nil {
		return nil, err
	}
	return &ContractIAllocationManagerOperatorSlashedIterator{contract: _ContractIAllocationManager.contract, event: "OperatorSlashed", logs: logs, sub: sub}, nil
}

// WatchOperatorSlashed is a free log subscription operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
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

// ParseOperatorSlashed is a log parse operation binding the contract event 0xe672839d3c371691acdb52de9fefc94b3dbf407dc0920ef566c7c059ad575b1c.
//
// Solidity: event OperatorSlashed(address operator, uint32 operatorSetId, address strategy, uint16 bipsToSlash)
func (_ContractIAllocationManager *ContractIAllocationManagerFilterer) ParseOperatorSlashed(log types.Log) (*ContractIAllocationManagerOperatorSlashed, error) {
	event := new(ContractIAllocationManagerOperatorSlashed)
	if err := _ContractIAllocationManager.contract.UnpackLog(event, "OperatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
