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
	Bin: "0x608060405234801561001057600080fd5b50610f56806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063161a334d146100515780631f5ac1b214610066578063de29fac01461008f578063e8bb9ae6146100bd575b600080fd5b61006461005f366004610cd8565b6100fe565b005b610079610074366004610d51565b610430565b6040516100869190610d81565b60405180910390f35b6100af61009d366004610d51565b60006020819052908152604090205481565b604051908152602001610086565b6100e66100cb366004610d98565b6001602052600090815260409020546001600160a01b031681565b6040516001600160a01b039091168152602001610086565b6000610109836104cb565b33600090815260208190526040902054909150156101ac5760405162461bcd60e51b815260206004820152604f60248201527f424c535075626c69634b6579436f6d70656e6469756d2e72656769737465724260448201527f4c535075626c69634b65793a206f70657261746f7220616c726561647920726560648201526e6769737465726564207075626b657960881b608482015260a4015b60405180910390fd5b6000818152600160205260409020546001600160a01b03161561024a5760405162461bcd60e51b815260206004820152604a60248201527f424c535075626c69634b6579436f6d70656e6469756d2e72656769737465724260448201527f4c535075626c69634b65793a207075626c6963206b657920616c7265616479206064820152691c9959da5cdd195c995960b21b608482015260a4016101a3565b600061025533610430565b8551602080880151875188830151885189850151875186890151604051999a506000997f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001996102ad9990989796959493929101610dda565b6040516020818303038152906040528051906020012060001c6102d09190610e26565b90506103376102e96102e2878461050e565b88906105a5565b6102f1610639565b61033161032a85610324604080518082018252600080825260209182015281518083019092526001825260029082015290565b9061050e565b86906105a5565b876106f9565b6103c15760405162461bcd60e51b815260206004820152604f60248201527f424c535075626c69634b6579436f6d70656e6469756d2e72656769737465724260448201527f4c535075626c69634b65793a20473120616e642047322070726976617465206b60648201526e0caf240c8de40dcdee840dac2e8c6d608b1b608482015260a4016101a3565b3360008181526020818152604080832087905586835260019091529081902080546001600160a01b03191683179055517fe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041906104209088908890610e48565b60405180910390a2505050505050565b60408051808201909152600080825260208201526040516bffffffffffffffffffffffff19606084811b8216602084015230901b1660348201524660488201527f456967656e4c617965725f424e3235345f5075626b65795f52656769737472616068820152633a34b7b760e11b60888201526104c590608c0160405160208183030381529060405280519060200120610966565b92915050565b6000816000015182602001516040516020016104f1929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b604080518082019091526000808252602082015261052a610b20565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561055d5761055f565bfe5b508061059d5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b60448201526064016101a3565b505092915050565b60408051808201909152600080825260208201526105c1610b3e565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561055d57508061059d5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b60448201526064016101a3565b610641610b5c565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b604080518082018252858152602080820185905282518084019093528583528201839052600091610728610b81565b60005b60028110156108ed576000610741826006610eae565b905084826002811061075557610755610e82565b60200201515183610767836000610ecd565b600c811061077757610777610e82565b602002015284826002811061078e5761078e610e82565b602002015160200151838260016107a59190610ecd565b600c81106107b5576107b5610e82565b60200201528382600281106107cc576107cc610e82565b60200201515151836107df836002610ecd565b600c81106107ef576107ef610e82565b602002015283826002811061080657610806610e82565b602002015151600160200201518361081f836003610ecd565b600c811061082f5761082f610e82565b602002015283826002811061084657610846610e82565b60200201516020015160006002811061086157610861610e82565b602002015183610872836004610ecd565b600c811061088257610882610e82565b602002015283826002811061089957610899610e82565b6020020151602001516001600281106108b4576108b4610e82565b6020020151836108c5836005610ecd565b600c81106108d5576108d5610e82565b602002015250806108e581610ee5565b91505061072b565b506108f6610ba0565b60006020826101808560086107d05a03fa905080801561055d5750806109565760405162461bcd60e51b81526020600482015260156024820152741c185a5c9a5b99cb5bdc18dbd9194b59985a5b1959605a1b60448201526064016101a3565b5051151598975050505050505050565b604080518082019091526000808252602082015260008080610996600080516020610f0183398151915286610e26565b90505b6109a2816109f6565b9093509150600080516020610f018339815191528283098314156109dc576040805180820190915290815260208101919091529392505050565b600080516020610f01833981519152600182089050610999565b60008080600080516020610f018339815191526003600080516020610f0183398151915286600080516020610f01833981519152888909090890506000610a6c827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52600080516020610f01833981519152610a78565b91959194509092505050565b600080610a83610ba0565b610a8b610bbe565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561055d575082610b155760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c75726500000000000060448201526064016101a3565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b6040518060400160405280610b6f610bdc565b8152602001610b7c610bdc565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610c3357610c33610bfa565b60405290565b600060408284031215610c4b57600080fd5b6040516040810181811067ffffffffffffffff82111715610c6e57610c6e610bfa565b604052823581526020928301359281019290925250919050565b600082601f830112610c9957600080fd5b610ca1610c10565b806040840185811115610cb357600080fd5b845b81811015610ccd578035845260209384019301610cb5565b509095945050505050565b6000806000838503610100811215610cef57600080fd5b610cf98686610c39565b9350610d088660408701610c39565b92506080607f1982011215610d1c57600080fd5b50610d25610c10565b610d328660808701610c88565b8152610d418660c08701610c88565b6020820152809150509250925092565b600060208284031215610d6357600080fd5b81356001600160a01b0381168114610d7a57600080fd5b9392505050565b8151815260208083015190820152604081016104c5565b600060208284031215610daa57600080fd5b5035919050565b8060005b6002811015610dd4578151845260209384019390910190600101610db5565b50505050565b888152876020820152866040820152856060820152610dfc6080820186610db1565b610e0960c0820185610db1565b610100810192909252610120820152610140019695505050505050565b600082610e4357634e487b7160e01b600052601260045260246000fd5b500690565b825181526020808401519082015260c08101610e68604083018451610db1565b6020830151610e7a6080840182610db1565b509392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610ec857610ec8610e98565b500290565b60008219821115610ee057610ee0610e98565b500190565b6000600019821415610ef957610ef9610e98565b506001019056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47a2646970667358221220571b2cc632fe58248f560e92fcf23273c5bd065b6c8e83c7b38e775508b29c1864736f6c634300080c0033",
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

// ContractBLSPublicKeyCompendium is an auto generated Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendium struct {
	ContractBLSPublicKeyCompendiumCaller     // Read-only binding to the contract
	ContractBLSPublicKeyCompendiumTransactor // Write-only binding to the contract
	ContractBLSPublicKeyCompendiumFilterer   // Log filterer for contract events
}

// ContractBLSPublicKeyCompendiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPublicKeyCompendiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBLSPublicKeyCompendiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPublicKeyCompendiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBLSPublicKeyCompendiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

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
