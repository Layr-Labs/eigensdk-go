// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSocketRegistry

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

// ContractSocketRegistryMetaData contains all meta data concerning the ContractSocketRegistry contract.
var ContractSocketRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slashingRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorSocket\",\"inputs\":[{\"name\":\"_operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorIdToSocket\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setOperatorSocket\",\"inputs\":[{\"name\":\"_operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashingRegistryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"OnlySlashingRegistryCoordinator\",\"inputs\":[]}]",
	Bin: "0x60a0604052348015600e575f5ffd5b5060405161058e38038061058e833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f5ffd5b81516001600160a01b0381168114605f575f5ffd5b9392505050565b60805161050a6100845f395f8181608f015261021f015261050a5ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806310bea0d71461004e578063af65fdfc14610077578063cf1d6b421461008a578063f043367e146100c9575b5f5ffd5b61006161005c366004610279565b6100de565b60405161006e9190610290565b60405180910390f35b610061610085366004610279565b61017d565b6100b17f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161006e565b6100dc6100d73660046102d9565b610214565b005b5f8181526020819052604090208054606091906100fa90610396565b80601f016020809104026020016040519081016040528092919081815260200182805461012690610396565b80156101715780601f1061014857610100808354040283529160200191610171565b820191905f5260205f20905b81548152906001019060200180831161015457829003601f168201915b50505050509050919050565b5f602081905290815260409020805461019590610396565b80601f01602080910402602001604051908101604052809291908181526020018280546101c190610396565b801561020c5780601f106101e35761010080835404028352916020019161020c565b820191905f5260205f20905b8154815290600101906020018083116101ef57829003601f168201915b505050505081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461025d57604051632c01b20560e21b815260040160405180910390fd5b5f8281526020819052604090206102748282610419565b505050565b5f60208284031215610289575f5ffd5b5035919050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b634e487b7160e01b5f52604160045260245ffd5b5f5f604083850312156102ea575f5ffd5b82359150602083013567ffffffffffffffff811115610307575f5ffd5b8301601f81018513610317575f5ffd5b803567ffffffffffffffff811115610331576103316102c5565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610360576103606102c5565b604052818152828201602001871015610377575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b600181811c908216806103aa57607f821691505b6020821081036103c857634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561027457805f5260205f20601f840160051c810160208510156103f35750805b601f840160051c820191505b81811015610412575f81556001016103ff565b5050505050565b815167ffffffffffffffff811115610433576104336102c5565b610447816104418454610396565b846103ce565b6020601f821160018114610479575f83156104625750848201515b5f19600385901b1c1916600184901b178455610412565b5f84815260208120601f198516915b828110156104a85787850151825560209485019460019092019101610488565b50848210156104c557868401515f19600387901b60f8161c191681555b50505050600190811b0190555056fea2646970667358221220191eb54453ffb02ffa7a875dace13e4eab340f8cca6f18bd71de064f3c37ed8164736f6c634300081b0033",
}

// ContractSocketRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSocketRegistryMetaData.ABI instead.
var ContractSocketRegistryABI = ContractSocketRegistryMetaData.ABI

// ContractSocketRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSocketRegistryMetaData.Bin instead.
var ContractSocketRegistryBin = ContractSocketRegistryMetaData.Bin

// DeployContractSocketRegistry deploys a new Ethereum contract, binding an instance of ContractSocketRegistry to it.
func DeployContractSocketRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _slashingRegistryCoordinator common.Address) (common.Address, *types.Transaction, *ContractSocketRegistry, error) {
	parsed, err := ContractSocketRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSocketRegistryBin), backend, _slashingRegistryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSocketRegistry{ContractSocketRegistryCaller: ContractSocketRegistryCaller{contract: contract}, ContractSocketRegistryTransactor: ContractSocketRegistryTransactor{contract: contract}, ContractSocketRegistryFilterer: ContractSocketRegistryFilterer{contract: contract}}, nil
}

// ContractSocketRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractSocketRegistryMethods interface {
	ContractSocketRegistryCalls
	ContractSocketRegistryTransacts
	ContractSocketRegistryFilters
}

// ContractSocketRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractSocketRegistryCalls interface {
	GetOperatorSocket(opts *bind.CallOpts, _operatorId [32]byte) (string, error)

	OperatorIdToSocket(opts *bind.CallOpts, arg0 [32]byte) (string, error)

	SlashingRegistryCoordinator(opts *bind.CallOpts) (common.Address, error)
}

// ContractSocketRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractSocketRegistryTransacts interface {
	SetOperatorSocket(opts *bind.TransactOpts, _operatorId [32]byte, _socket string) (*types.Transaction, error)
}

// ContractSocketRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractSocketRegistryFilters interface {
}

// ContractSocketRegistry is an auto generated Go binding around an Ethereum contract.
type ContractSocketRegistry struct {
	ContractSocketRegistryCaller     // Read-only binding to the contract
	ContractSocketRegistryTransactor // Write-only binding to the contract
	ContractSocketRegistryFilterer   // Log filterer for contract events
}

// ContractSocketRegistry implements the ContractSocketRegistryMethods interface.
var _ ContractSocketRegistryMethods = (*ContractSocketRegistry)(nil)

// ContractSocketRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSocketRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSocketRegistryCaller implements the ContractSocketRegistryCalls interface.
var _ ContractSocketRegistryCalls = (*ContractSocketRegistryCaller)(nil)

// ContractSocketRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSocketRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSocketRegistryTransactor implements the ContractSocketRegistryTransacts interface.
var _ ContractSocketRegistryTransacts = (*ContractSocketRegistryTransactor)(nil)

// ContractSocketRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSocketRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSocketRegistryFilterer implements the ContractSocketRegistryFilters interface.
var _ ContractSocketRegistryFilters = (*ContractSocketRegistryFilterer)(nil)

// ContractSocketRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSocketRegistrySession struct {
	Contract     *ContractSocketRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractSocketRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSocketRegistryCallerSession struct {
	Contract *ContractSocketRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractSocketRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSocketRegistryTransactorSession struct {
	Contract     *ContractSocketRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractSocketRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSocketRegistryRaw struct {
	Contract *ContractSocketRegistry // Generic contract binding to access the raw methods on
}

// ContractSocketRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSocketRegistryCallerRaw struct {
	Contract *ContractSocketRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSocketRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSocketRegistryTransactorRaw struct {
	Contract *ContractSocketRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSocketRegistry creates a new instance of ContractSocketRegistry, bound to a specific deployed contract.
func NewContractSocketRegistry(address common.Address, backend bind.ContractBackend) (*ContractSocketRegistry, error) {
	contract, err := bindContractSocketRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSocketRegistry{ContractSocketRegistryCaller: ContractSocketRegistryCaller{contract: contract}, ContractSocketRegistryTransactor: ContractSocketRegistryTransactor{contract: contract}, ContractSocketRegistryFilterer: ContractSocketRegistryFilterer{contract: contract}}, nil
}

// NewContractSocketRegistryCaller creates a new read-only instance of ContractSocketRegistry, bound to a specific deployed contract.
func NewContractSocketRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractSocketRegistryCaller, error) {
	contract, err := bindContractSocketRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSocketRegistryCaller{contract: contract}, nil
}

// NewContractSocketRegistryTransactor creates a new write-only instance of ContractSocketRegistry, bound to a specific deployed contract.
func NewContractSocketRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSocketRegistryTransactor, error) {
	contract, err := bindContractSocketRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSocketRegistryTransactor{contract: contract}, nil
}

// NewContractSocketRegistryFilterer creates a new log filterer instance of ContractSocketRegistry, bound to a specific deployed contract.
func NewContractSocketRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSocketRegistryFilterer, error) {
	contract, err := bindContractSocketRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSocketRegistryFilterer{contract: contract}, nil
}

// bindContractSocketRegistry binds a generic wrapper to an already deployed contract.
func bindContractSocketRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSocketRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSocketRegistry *ContractSocketRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSocketRegistry.Contract.ContractSocketRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSocketRegistry *ContractSocketRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSocketRegistry.Contract.ContractSocketRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSocketRegistry *ContractSocketRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSocketRegistry.Contract.ContractSocketRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSocketRegistry *ContractSocketRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSocketRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSocketRegistry *ContractSocketRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSocketRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSocketRegistry *ContractSocketRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSocketRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x10bea0d7.
//
// Solidity: function getOperatorSocket(bytes32 _operatorId) view returns(string)
func (_ContractSocketRegistry *ContractSocketRegistryCaller) GetOperatorSocket(opts *bind.CallOpts, _operatorId [32]byte) (string, error) {
	var out []interface{}
	err := _ContractSocketRegistry.contract.Call(opts, &out, "getOperatorSocket", _operatorId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x10bea0d7.
//
// Solidity: function getOperatorSocket(bytes32 _operatorId) view returns(string)
func (_ContractSocketRegistry *ContractSocketRegistrySession) GetOperatorSocket(_operatorId [32]byte) (string, error) {
	return _ContractSocketRegistry.Contract.GetOperatorSocket(&_ContractSocketRegistry.CallOpts, _operatorId)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x10bea0d7.
//
// Solidity: function getOperatorSocket(bytes32 _operatorId) view returns(string)
func (_ContractSocketRegistry *ContractSocketRegistryCallerSession) GetOperatorSocket(_operatorId [32]byte) (string, error) {
	return _ContractSocketRegistry.Contract.GetOperatorSocket(&_ContractSocketRegistry.CallOpts, _operatorId)
}

// OperatorIdToSocket is a free data retrieval call binding the contract method 0xaf65fdfc.
//
// Solidity: function operatorIdToSocket(bytes32 ) view returns(string)
func (_ContractSocketRegistry *ContractSocketRegistryCaller) OperatorIdToSocket(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _ContractSocketRegistry.contract.Call(opts, &out, "operatorIdToSocket", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OperatorIdToSocket is a free data retrieval call binding the contract method 0xaf65fdfc.
//
// Solidity: function operatorIdToSocket(bytes32 ) view returns(string)
func (_ContractSocketRegistry *ContractSocketRegistrySession) OperatorIdToSocket(arg0 [32]byte) (string, error) {
	return _ContractSocketRegistry.Contract.OperatorIdToSocket(&_ContractSocketRegistry.CallOpts, arg0)
}

// OperatorIdToSocket is a free data retrieval call binding the contract method 0xaf65fdfc.
//
// Solidity: function operatorIdToSocket(bytes32 ) view returns(string)
func (_ContractSocketRegistry *ContractSocketRegistryCallerSession) OperatorIdToSocket(arg0 [32]byte) (string, error) {
	return _ContractSocketRegistry.Contract.OperatorIdToSocket(&_ContractSocketRegistry.CallOpts, arg0)
}

// SlashingRegistryCoordinator is a free data retrieval call binding the contract method 0xcf1d6b42.
//
// Solidity: function slashingRegistryCoordinator() view returns(address)
func (_ContractSocketRegistry *ContractSocketRegistryCaller) SlashingRegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractSocketRegistry.contract.Call(opts, &out, "slashingRegistryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashingRegistryCoordinator is a free data retrieval call binding the contract method 0xcf1d6b42.
//
// Solidity: function slashingRegistryCoordinator() view returns(address)
func (_ContractSocketRegistry *ContractSocketRegistrySession) SlashingRegistryCoordinator() (common.Address, error) {
	return _ContractSocketRegistry.Contract.SlashingRegistryCoordinator(&_ContractSocketRegistry.CallOpts)
}

// SlashingRegistryCoordinator is a free data retrieval call binding the contract method 0xcf1d6b42.
//
// Solidity: function slashingRegistryCoordinator() view returns(address)
func (_ContractSocketRegistry *ContractSocketRegistryCallerSession) SlashingRegistryCoordinator() (common.Address, error) {
	return _ContractSocketRegistry.Contract.SlashingRegistryCoordinator(&_ContractSocketRegistry.CallOpts)
}

// SetOperatorSocket is a paid mutator transaction binding the contract method 0xf043367e.
//
// Solidity: function setOperatorSocket(bytes32 _operatorId, string _socket) returns()
func (_ContractSocketRegistry *ContractSocketRegistryTransactor) SetOperatorSocket(opts *bind.TransactOpts, _operatorId [32]byte, _socket string) (*types.Transaction, error) {
	return _ContractSocketRegistry.contract.Transact(opts, "setOperatorSocket", _operatorId, _socket)
}

// SetOperatorSocket is a paid mutator transaction binding the contract method 0xf043367e.
//
// Solidity: function setOperatorSocket(bytes32 _operatorId, string _socket) returns()
func (_ContractSocketRegistry *ContractSocketRegistrySession) SetOperatorSocket(_operatorId [32]byte, _socket string) (*types.Transaction, error) {
	return _ContractSocketRegistry.Contract.SetOperatorSocket(&_ContractSocketRegistry.TransactOpts, _operatorId, _socket)
}

// SetOperatorSocket is a paid mutator transaction binding the contract method 0xf043367e.
//
// Solidity: function setOperatorSocket(bytes32 _operatorId, string _socket) returns()
func (_ContractSocketRegistry *ContractSocketRegistryTransactorSession) SetOperatorSocket(_operatorId [32]byte, _socket string) (*types.Transaction, error) {
	return _ContractSocketRegistry.Contract.SetOperatorSocket(&_ContractSocketRegistry.TransactOpts, _operatorId, _socket)
}
