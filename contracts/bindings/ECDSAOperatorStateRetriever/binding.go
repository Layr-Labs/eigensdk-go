// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractECDSAOperatorStateRetriever

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

// ECDSAOperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type ECDSAOperatorStateRetrieverCheckSignaturesIndices struct {
	SignerQuorumBitmapIndices []uint32
	TotalStakeIndices         []uint32
	SignerStakeIndices        [][]uint32
}

// ECDSAOperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type ECDSAOperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId common.Address
	Stake      *big.Int
}

// ContractECDSAOperatorStateRetrieverMetaData contains all meta data concerning the ContractECDSAOperatorStateRetriever contract.
var ContractECDSAOperatorStateRetrieverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractECDSARegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signerOperatorIds\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structECDSAOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"signerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"signerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractECDSARegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structECDSAOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractECDSARegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structECDSAOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"}]",
	Bin: "0x6",
}

// ContractECDSAOperatorStateRetrieverABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractECDSAOperatorStateRetrieverMetaData.ABI instead.
var ContractECDSAOperatorStateRetrieverABI = ContractECDSAOperatorStateRetrieverMetaData.ABI

// ContractECDSAOperatorStateRetrieverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractECDSAOperatorStateRetrieverMetaData.Bin instead.
var ContractECDSAOperatorStateRetrieverBin = ContractECDSAOperatorStateRetrieverMetaData.Bin

// DeployContractECDSAOperatorStateRetriever deploys a new Ethereum contract, binding an instance of ContractECDSAOperatorStateRetriever to it.
func DeployContractECDSAOperatorStateRetriever(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractECDSAOperatorStateRetriever, error) {
	parsed, err := ContractECDSAOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractECDSAOperatorStateRetrieverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractECDSAOperatorStateRetriever{ContractECDSAOperatorStateRetrieverCaller: ContractECDSAOperatorStateRetrieverCaller{contract: contract}, ContractECDSAOperatorStateRetrieverTransactor: ContractECDSAOperatorStateRetrieverTransactor{contract: contract}, ContractECDSAOperatorStateRetrieverFilterer: ContractECDSAOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// ContractECDSAOperatorStateRetrieverMethods is an auto generated interface around an Ethereum contract.
type ContractECDSAOperatorStateRetrieverMethods interface {
	ContractECDSAOperatorStateRetrieverCalls
	ContractECDSAOperatorStateRetrieverTransacts
	ContractECDSAOperatorStateRetrieverFilters
}

// ContractECDSAOperatorStateRetrieverCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractECDSAOperatorStateRetrieverCalls interface {
	GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, signerOperatorIds []common.Address) (ECDSAOperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, operatorId common.Address, blockNumber uint32) (*big.Int, [][]ECDSAOperatorStateRetrieverOperator, error)

	GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]ECDSAOperatorStateRetrieverOperator, error)
}

// ContractECDSAOperatorStateRetrieverTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractECDSAOperatorStateRetrieverTransacts interface {
}

// ContractECDSAOperatorStateRetrieverFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractECDSAOperatorStateRetrieverFilters interface {
}

// ContractECDSAOperatorStateRetriever is an auto generated Go binding around an Ethereum contract.
type ContractECDSAOperatorStateRetriever struct {
	ContractECDSAOperatorStateRetrieverCaller     // Read-only binding to the contract
	ContractECDSAOperatorStateRetrieverTransactor // Write-only binding to the contract
	ContractECDSAOperatorStateRetrieverFilterer   // Log filterer for contract events
}

// ContractECDSAOperatorStateRetriever implements the ContractECDSAOperatorStateRetrieverMethods interface.
var _ ContractECDSAOperatorStateRetrieverMethods = (*ContractECDSAOperatorStateRetriever)(nil)

// ContractECDSAOperatorStateRetrieverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractECDSAOperatorStateRetrieverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractECDSAOperatorStateRetrieverCaller implements the ContractECDSAOperatorStateRetrieverCalls interface.
var _ ContractECDSAOperatorStateRetrieverCalls = (*ContractECDSAOperatorStateRetrieverCaller)(nil)

// ContractECDSAOperatorStateRetrieverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractECDSAOperatorStateRetrieverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractECDSAOperatorStateRetrieverTransactor implements the ContractECDSAOperatorStateRetrieverTransacts interface.
var _ ContractECDSAOperatorStateRetrieverTransacts = (*ContractECDSAOperatorStateRetrieverTransactor)(nil)

// ContractECDSAOperatorStateRetrieverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractECDSAOperatorStateRetrieverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractECDSAOperatorStateRetrieverFilterer implements the ContractECDSAOperatorStateRetrieverFilters interface.
var _ ContractECDSAOperatorStateRetrieverFilters = (*ContractECDSAOperatorStateRetrieverFilterer)(nil)

// ContractECDSAOperatorStateRetrieverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractECDSAOperatorStateRetrieverSession struct {
	Contract     *ContractECDSAOperatorStateRetriever // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                        // Call options to use throughout this session
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractECDSAOperatorStateRetrieverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractECDSAOperatorStateRetrieverCallerSession struct {
	Contract *ContractECDSAOperatorStateRetrieverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                              // Call options to use throughout this session
}

// ContractECDSAOperatorStateRetrieverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractECDSAOperatorStateRetrieverTransactorSession struct {
	Contract     *ContractECDSAOperatorStateRetrieverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                              // Transaction auth options to use throughout this session
}

// ContractECDSAOperatorStateRetrieverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractECDSAOperatorStateRetrieverRaw struct {
	Contract *ContractECDSAOperatorStateRetriever // Generic contract binding to access the raw methods on
}

// ContractECDSAOperatorStateRetrieverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractECDSAOperatorStateRetrieverCallerRaw struct {
	Contract *ContractECDSAOperatorStateRetrieverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractECDSAOperatorStateRetrieverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractECDSAOperatorStateRetrieverTransactorRaw struct {
	Contract *ContractECDSAOperatorStateRetrieverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractECDSAOperatorStateRetriever creates a new instance of ContractECDSAOperatorStateRetriever, bound to a specific deployed contract.
func NewContractECDSAOperatorStateRetriever(address common.Address, backend bind.ContractBackend) (*ContractECDSAOperatorStateRetriever, error) {
	contract, err := bindContractECDSAOperatorStateRetriever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractECDSAOperatorStateRetriever{ContractECDSAOperatorStateRetrieverCaller: ContractECDSAOperatorStateRetrieverCaller{contract: contract}, ContractECDSAOperatorStateRetrieverTransactor: ContractECDSAOperatorStateRetrieverTransactor{contract: contract}, ContractECDSAOperatorStateRetrieverFilterer: ContractECDSAOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// NewContractECDSAOperatorStateRetrieverCaller creates a new read-only instance of ContractECDSAOperatorStateRetriever, bound to a specific deployed contract.
func NewContractECDSAOperatorStateRetrieverCaller(address common.Address, caller bind.ContractCaller) (*ContractECDSAOperatorStateRetrieverCaller, error) {
	contract, err := bindContractECDSAOperatorStateRetriever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractECDSAOperatorStateRetrieverCaller{contract: contract}, nil
}

// NewContractECDSAOperatorStateRetrieverTransactor creates a new write-only instance of ContractECDSAOperatorStateRetriever, bound to a specific deployed contract.
func NewContractECDSAOperatorStateRetrieverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractECDSAOperatorStateRetrieverTransactor, error) {
	contract, err := bindContractECDSAOperatorStateRetriever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractECDSAOperatorStateRetrieverTransactor{contract: contract}, nil
}

// NewContractECDSAOperatorStateRetrieverFilterer creates a new log filterer instance of ContractECDSAOperatorStateRetriever, bound to a specific deployed contract.
func NewContractECDSAOperatorStateRetrieverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractECDSAOperatorStateRetrieverFilterer, error) {
	contract, err := bindContractECDSAOperatorStateRetriever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractECDSAOperatorStateRetrieverFilterer{contract: contract}, nil
}

// bindContractECDSAOperatorStateRetriever binds a generic wrapper to an already deployed contract.
func bindContractECDSAOperatorStateRetriever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractECDSAOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractECDSAOperatorStateRetriever.Contract.ContractECDSAOperatorStateRetrieverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.ContractECDSAOperatorStateRetrieverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.ContractECDSAOperatorStateRetrieverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractECDSAOperatorStateRetriever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.contract.Transact(opts, method, params...)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x93169f46.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, address[] signerOperatorIds) view returns((uint32[],uint32[],uint32[][]))
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, signerOperatorIds []common.Address) (ECDSAOperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractECDSAOperatorStateRetriever.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, signerOperatorIds)

	if err != nil {
		return *new(ECDSAOperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(ECDSAOperatorStateRetrieverCheckSignaturesIndices)).(*ECDSAOperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x93169f46.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, address[] signerOperatorIds) view returns((uint32[],uint32[],uint32[][]))
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, signerOperatorIds []common.Address) (ECDSAOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractECDSAOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, signerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x93169f46.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, address[] signerOperatorIds) view returns((uint32[],uint32[],uint32[][]))
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, signerOperatorIds []common.Address) (ECDSAOperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractECDSAOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, signerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x2617c130.
//
// Solidity: function getOperatorState(address registryCoordinator, address operatorId, uint32 blockNumber) view returns(uint256, (address,address,uint96)[][])
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, operatorId common.Address, blockNumber uint32) (*big.Int, [][]ECDSAOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractECDSAOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]ECDSAOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]ECDSAOperatorStateRetrieverOperator)).(*[][]ECDSAOperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x2617c130.
//
// Solidity: function getOperatorState(address registryCoordinator, address operatorId, uint32 blockNumber) view returns(uint256, (address,address,uint96)[][])
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverSession) GetOperatorState(registryCoordinator common.Address, operatorId common.Address, blockNumber uint32) (*big.Int, [][]ECDSAOperatorStateRetrieverOperator, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.GetOperatorState(&_ContractECDSAOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x2617c130.
//
// Solidity: function getOperatorState(address registryCoordinator, address operatorId, uint32 blockNumber) view returns(uint256, (address,address,uint96)[][])
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCallerSession) GetOperatorState(registryCoordinator common.Address, operatorId common.Address, blockNumber uint32) (*big.Int, [][]ECDSAOperatorStateRetrieverOperator, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.GetOperatorState(&_ContractECDSAOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,address,uint96)[][])
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]ECDSAOperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractECDSAOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]ECDSAOperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]ECDSAOperatorStateRetrieverOperator)).(*[][]ECDSAOperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,address,uint96)[][])
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverSession) GetOperatorState0(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]ECDSAOperatorStateRetrieverOperator, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.GetOperatorState0(&_ContractECDSAOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,address,uint96)[][])
func (_ContractECDSAOperatorStateRetriever *ContractECDSAOperatorStateRetrieverCallerSession) GetOperatorState0(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]ECDSAOperatorStateRetrieverOperator, error) {
	return _ContractECDSAOperatorStateRetriever.Contract.GetOperatorState0(&_ContractECDSAOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}
