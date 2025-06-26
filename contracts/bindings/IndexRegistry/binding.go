// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIndexRegistry

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

// IIndexRegistryTypesOperatorUpdate is an auto generated low-level Go binding around an user-defined struct.
type IIndexRegistryTypesOperatorUpdate struct {
	FromBlockNumber uint32
	OperatorId      [32]byte
}

// IIndexRegistryTypesQuorumUpdate is an auto generated low-level Go binding around an user-defined struct.
type IIndexRegistryTypesQuorumUpdate struct {
	FromBlockNumber uint32
	NumOperators    uint32
}

// ContractIndexRegistryMetaData contains all meta data concerning the ContractIndexRegistry contract.
var ContractIndexRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slashingRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"OPERATOR_DOES_NOT_EXIST_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentOperatorIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestOperatorUpdate\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistryTypes.OperatorUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestQuorumUpdate\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistryTypes.QuorumUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numOperators\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorListAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"operatorIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"arrayIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistryTypes.OperatorUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"quorumIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIIndexRegistryTypes.QuorumUpdate\",\"components\":[{\"name\":\"fromBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"numOperators\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalOperatorsForQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalOperatorsForQuorumAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"QuorumIndexUpdate\",\"inputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"newOperatorIndex\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorIdDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b5060405161116338038061116383398101604081905261002e91610107565b6001600160a01b0381166080528061004461004b565b5050610134565b5f54610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610105575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610117575f5ffd5b81516001600160a01b038116811461012d575f5ffd5b9392505050565b6080516110106101535f395f8181610158015261078d01526110105ff3fe608060405234801561000f575f5ffd5b50600436106100ca575f3560e01c80638121906f11610088578063bd29b8cd11610063578063bd29b8cd14610224578063caa3cd7614610237578063e2e685801461024c578063f34109221461027c575f5ffd5b80638121906f146101ba57806389026245146101f1578063a48bb0ac14610211575f5ffd5b8062bff04d146100ce57806312d1d74d146100f757806326d941f21461012b5780632ed583e5146101405780636d14a987146101535780637b5d26be14610192575b5f5ffd5b6100e16100dc366004610d6e565b61028f565b6040516100ee9190610de5565b60405180910390f35b61010a610105366004610e55565b610398565b60408051825163ffffffff16815260209283015192810192909252016100ee565b61013e610139366004610e86565b6103dd565b005b61010a61014e366004610e9f565b610476565b61017a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ee565b6101a56101a0366004610e55565b6104f9565b60405163ffffffff90911681526020016100ee565b6101cd6101c8366004610e86565b610504565b60408051825163ffffffff90811682526020938401511692810192909252016100ee565b6102046101ff366004610e55565b61054a565b6040516100ee9190610edf565b6101cd61021f366004610e55565b61062a565b61013e610232366004610d6e565b61069e565b61023e5f81565b6040519081526020016100ee565b6101a561025a366004610f16565b600160209081525f928352604080842090915290825290205463ffffffff1681565b6101a561028a366004610e86565b610764565b6060610299610782565b5f8267ffffffffffffffff8111156102b3576102b3610f3e565b6040519080825280602002602001820160405280156102dc578160200160208202803683370190505b5090505f5b8381101561038d575f8585838181106102fc576102fc610f52565b919091013560f81c5f8181526003602052604081205491935090915081900361033857604051637310cff560e11b815260040160405180910390fd5b5f610342836107cd565b90506103598984610354600185610f7a565b6108c4565b8085858151811061036c5761036c610f52565b63ffffffff92909216602092830291909101909101525050506001016102e1565b5090505b9392505050565b604080518082019091525f80825260208201526103b5838361094c565b60408051808201909152815463ffffffff168152600190910154602082015290505b92915050565b6103e5610782565b60ff81165f908152600360205260409020541561041557604051637310cff560e11b815260040160405180910390fd5b60ff165f908152600360209081526040808320815180830190925263ffffffff438116835282840185815282546001810184559286529390942091519101805492518416600160201b0267ffffffffffffffff199093169190931617179055565b604080518082019091525f808252602082015260ff84165f90815260026020908152604080832063ffffffff808816855292529091208054909184169081106104c1576104c1610f52565b5f91825260209182902060408051808201909152600290920201805463ffffffff168252600101549181019190915290509392505050565b5f61039183836109a1565b604080518082019091525f808252602082015261052082610ad7565b60408051808201909152905463ffffffff8082168352600160201b90910416602082015292915050565b60605f61055784846109a1565b90505f8163ffffffff1667ffffffffffffffff81111561057957610579610f3e565b6040519080825280602002602001820160405280156105a2578160200160208202803683370190505b5090505f5b8263ffffffff16811015610621576105c0868287610b16565b8282815181106105d2576105d2610f52565b6020026020010181815250505f5f1b8282815181106105f3576105f3610f52565b60200260200101510361061957604051637f12098d60e11b815260040160405180910390fd5b6001016105a7565b50949350505050565b604080518082019091525f808252602082015260ff83165f908152600360205260409020805463ffffffff841690811061066657610666610f52565b5f9182526020918290206040805180820190915291015463ffffffff8082168352600160201b90910416918101919091529392505050565b6106a6610782565b5f5b8181101561075e575f8383838181106106c3576106c3610f52565b919091013560f81c5f818152600360205260408120549193509091508190036106ff57604051637310cff560e11b815260040160405180910390fd5b60ff82165f90815260016020908152604080832089845290915281205463ffffffff169061072c84610be9565b90505f6107398583610c21565b905080891461074d5761074d8186856108c4565b5050600190930192506106a8915050565b50505050565b5f61076e82610ad7565b54600160201b900463ffffffff1692915050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107cb57604051634394dbdf60e11b815260040160405180910390fd5b565b5f5f6107d883610ad7565b80549091505f906107f790600160201b900463ffffffff166001610f96565b9050610804848383610c49565b60ff84165f90815260026020526040812090610821600184610f7a565b63ffffffff16815260208101919091526040015f9081205490036103915760ff84165f9081526002602052604081209061085c600184610f7a565b63ffffffff908116825260208083019390935260409182015f908120835180850190945243831684528385018281528154600180820184559284529590922093516002909502909301805463ffffffff19169490921693909317815591519101559392505050565b5f6108cf838361094c565b90506108dd83838387610ce6565b60ff83165f818152600160209081526040808320888452825291829020805463ffffffff191663ffffffff871690811790915582519384529083015285917f6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6910160405180910390a250505050565b60ff82165f90815260026020908152604080832063ffffffff85168452909152812080549061097c600183610fb2565b8154811061098c5761098c610f52565b905f5260205f20906002020191505092915050565b60ff82165f90815260036020526040812054805b8015610a465760ff85165f9081526003602052604081206109d7600184610fb2565b815481106109e7576109e7610f52565b5f9182526020918290206040805180820190915291015463ffffffff808216808452600160201b90920481169383019390935290925090861610610a33576020015192506103d7915050565b5080610a3e81610fc5565b9150506109b5565b5060405162461bcd60e51b815260206004820152605560248201527f496e64657852656769737472792e5f6f70657261746f72436f756e744174426c60448201527f6f636b4e756d6265723a2071756f72756d20646964206e6f742065786973742060648201527430ba1033b4bb32b710313637b1b590373ab6b132b960591b608482015260a40160405180910390fd5b60ff81165f908152600360205260408120805490610af6600183610fb2565b81548110610b0657610b06610f52565b905f5260205f2001915050919050565b60ff83165f90815260026020908152604080832063ffffffff86168452909152812054805b8015610bde5760ff86165f90815260026020908152604080832063ffffffff891684529091528120610b6e600184610fb2565b81548110610b7e57610b7e610f52565b5f91825260209182902060408051808201909152600290920201805463ffffffff9081168084526001909201549383019390935290925090861610610bcb57602001519250610391915050565b5080610bd681610fc5565b915050610b3b565b505f95945050505050565b5f5f610bf483610ad7565b80549091505f90610c1490600190600160201b900463ffffffff16610f7a565b9050610391848383610c49565b5f5f610c2d848461094c565b6001810154909150610c418585845f610ce6565b949350505050565b815463ffffffff438116911603610c7e57815463ffffffff8216600160201b0267ffffffff0000000019909116178255505050565b60ff83165f908152600360209081526040808320815180830190925263ffffffff438116835285811683850190815282546001810184559286529390942091519101805492518416600160201b0267ffffffffffffffff199093169190931617179055505050565b815463ffffffff438116911603610d03576001820181905561075e565b60ff939093165f90815260026020818152604080842063ffffffff968716855282528084208151808301909252438716825281830197885280546001808201835591865292909420905191909202909101805463ffffffff1916919094161783559251919092015550565b5f5f5f60408486031215610d80575f5ffd5b83359250602084013567ffffffffffffffff811115610d9d575f5ffd5b8401601f81018613610dad575f5ffd5b803567ffffffffffffffff811115610dc3575f5ffd5b866020828401011115610dd4575f5ffd5b939660209190910195509293505050565b602080825282518282018190525f918401906040840190835b81811015610e2257835163ffffffff16835260209384019390920191600101610dfe565b509095945050505050565b803560ff81168114610e3d575f5ffd5b919050565b803563ffffffff81168114610e3d575f5ffd5b5f5f60408385031215610e66575f5ffd5b610e6f83610e2d565b9150610e7d60208401610e42565b90509250929050565b5f60208284031215610e96575f5ffd5b61039182610e2d565b5f5f5f60608486031215610eb1575f5ffd5b610eba84610e2d565b9250610ec860208501610e42565b9150610ed660408501610e42565b90509250925092565b602080825282518282018190525f918401906040840190835b81811015610e22578351835260209384019390920191600101610ef8565b5f5f60408385031215610f27575f5ffd5b610f3083610e2d565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b63ffffffff82811682821603908111156103d7576103d7610f66565b63ffffffff81811683821601908111156103d7576103d7610f66565b818103818111156103d7576103d7610f66565b5f81610fd357610fd3610f66565b505f19019056fea26469706673582212202ffa0f0ada7c4746a01956769e2fa4c9c2f3737998fe9dbb10c5a0d97f31b81c64736f6c634300081b0033",
}

// ContractIndexRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIndexRegistryMetaData.ABI instead.
var ContractIndexRegistryABI = ContractIndexRegistryMetaData.ABI

// ContractIndexRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIndexRegistryMetaData.Bin instead.
var ContractIndexRegistryBin = ContractIndexRegistryMetaData.Bin

// DeployContractIndexRegistry deploys a new Ethereum contract, binding an instance of ContractIndexRegistry to it.
func DeployContractIndexRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _slashingRegistryCoordinator common.Address) (common.Address, *types.Transaction, *ContractIndexRegistry, error) {
	parsed, err := ContractIndexRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIndexRegistryBin), backend, _slashingRegistryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIndexRegistry{ContractIndexRegistryCaller: ContractIndexRegistryCaller{contract: contract}, ContractIndexRegistryTransactor: ContractIndexRegistryTransactor{contract: contract}, ContractIndexRegistryFilterer: ContractIndexRegistryFilterer{contract: contract}}, nil
}

// ContractIndexRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractIndexRegistryMethods interface {
	ContractIndexRegistryCalls
	ContractIndexRegistryTransacts
	ContractIndexRegistryFilters
}

// ContractIndexRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractIndexRegistryCalls interface {
	OPERATORDOESNOTEXISTID(opts *bind.CallOpts) ([32]byte, error)

	CurrentOperatorIndex(opts *bind.CallOpts, arg0 uint8, arg1 [32]byte) (uint32, error)

	GetLatestOperatorUpdate(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32) (IIndexRegistryTypesOperatorUpdate, error)

	GetLatestQuorumUpdate(opts *bind.CallOpts, quorumNumber uint8) (IIndexRegistryTypesQuorumUpdate, error)

	GetOperatorListAtBlockNumber(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32) ([][32]byte, error)

	GetOperatorUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryTypesOperatorUpdate, error)

	GetQuorumUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, quorumIndex uint32) (IIndexRegistryTypesQuorumUpdate, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)

	TotalOperatorsForQuorum(opts *bind.CallOpts, quorumNumber uint8) (uint32, error)

	TotalOperatorsForQuorumAtBlockNumber(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32) (uint32, error)
}

// ContractIndexRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractIndexRegistryTransacts interface {
	DeregisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error)

	InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error)
}

// ContractIndexRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractIndexRegistryFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractIndexRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractIndexRegistryInitialized, error)

	FilterQuorumIndexUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractIndexRegistryQuorumIndexUpdateIterator, error)
	WatchQuorumIndexUpdate(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryQuorumIndexUpdate, operatorId [][32]byte) (event.Subscription, error)
	ParseQuorumIndexUpdate(log types.Log) (*ContractIndexRegistryQuorumIndexUpdate, error)
}

// ContractIndexRegistry is an auto generated Go binding around an Ethereum contract.
type ContractIndexRegistry struct {
	ContractIndexRegistryCaller     // Read-only binding to the contract
	ContractIndexRegistryTransactor // Write-only binding to the contract
	ContractIndexRegistryFilterer   // Log filterer for contract events
}

// ContractIndexRegistry implements the ContractIndexRegistryMethods interface.
var _ ContractIndexRegistryMethods = (*ContractIndexRegistry)(nil)

// ContractIndexRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIndexRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIndexRegistryCaller implements the ContractIndexRegistryCalls interface.
var _ ContractIndexRegistryCalls = (*ContractIndexRegistryCaller)(nil)

// ContractIndexRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIndexRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIndexRegistryTransactor implements the ContractIndexRegistryTransacts interface.
var _ ContractIndexRegistryTransacts = (*ContractIndexRegistryTransactor)(nil)

// ContractIndexRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIndexRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIndexRegistryFilterer implements the ContractIndexRegistryFilters interface.
var _ ContractIndexRegistryFilters = (*ContractIndexRegistryFilterer)(nil)

// ContractIndexRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIndexRegistrySession struct {
	Contract     *ContractIndexRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ContractIndexRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIndexRegistryCallerSession struct {
	Contract *ContractIndexRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ContractIndexRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIndexRegistryTransactorSession struct {
	Contract     *ContractIndexRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ContractIndexRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIndexRegistryRaw struct {
	Contract *ContractIndexRegistry // Generic contract binding to access the raw methods on
}

// ContractIndexRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIndexRegistryCallerRaw struct {
	Contract *ContractIndexRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIndexRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIndexRegistryTransactorRaw struct {
	Contract *ContractIndexRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIndexRegistry creates a new instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistry(address common.Address, backend bind.ContractBackend) (*ContractIndexRegistry, error) {
	contract, err := bindContractIndexRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistry{ContractIndexRegistryCaller: ContractIndexRegistryCaller{contract: contract}, ContractIndexRegistryTransactor: ContractIndexRegistryTransactor{contract: contract}, ContractIndexRegistryFilterer: ContractIndexRegistryFilterer{contract: contract}}, nil
}

// NewContractIndexRegistryCaller creates a new read-only instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractIndexRegistryCaller, error) {
	contract, err := bindContractIndexRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryCaller{contract: contract}, nil
}

// NewContractIndexRegistryTransactor creates a new write-only instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIndexRegistryTransactor, error) {
	contract, err := bindContractIndexRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryTransactor{contract: contract}, nil
}

// NewContractIndexRegistryFilterer creates a new log filterer instance of ContractIndexRegistry, bound to a specific deployed contract.
func NewContractIndexRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIndexRegistryFilterer, error) {
	contract, err := bindContractIndexRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryFilterer{contract: contract}, nil
}

// bindContractIndexRegistry binds a generic wrapper to an already deployed contract.
func bindContractIndexRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIndexRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIndexRegistry *ContractIndexRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIndexRegistry.Contract.ContractIndexRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIndexRegistry *ContractIndexRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.ContractIndexRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIndexRegistry *ContractIndexRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.ContractIndexRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIndexRegistry *ContractIndexRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIndexRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIndexRegistry *ContractIndexRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIndexRegistry *ContractIndexRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.contract.Transact(opts, method, params...)
}

// OPERATORDOESNOTEXISTID is a free data retrieval call binding the contract method 0xcaa3cd76.
//
// Solidity: function OPERATOR_DOES_NOT_EXIST_ID() view returns(bytes32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) OPERATORDOESNOTEXISTID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "OPERATOR_DOES_NOT_EXIST_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORDOESNOTEXISTID is a free data retrieval call binding the contract method 0xcaa3cd76.
//
// Solidity: function OPERATOR_DOES_NOT_EXIST_ID() view returns(bytes32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) OPERATORDOESNOTEXISTID() ([32]byte, error) {
	return _ContractIndexRegistry.Contract.OPERATORDOESNOTEXISTID(&_ContractIndexRegistry.CallOpts)
}

// OPERATORDOESNOTEXISTID is a free data retrieval call binding the contract method 0xcaa3cd76.
//
// Solidity: function OPERATOR_DOES_NOT_EXIST_ID() view returns(bytes32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) OPERATORDOESNOTEXISTID() ([32]byte, error) {
	return _ContractIndexRegistry.Contract.OPERATORDOESNOTEXISTID(&_ContractIndexRegistry.CallOpts)
}

// CurrentOperatorIndex is a free data retrieval call binding the contract method 0xe2e68580.
//
// Solidity: function currentOperatorIndex(uint8 , bytes32 ) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) CurrentOperatorIndex(opts *bind.CallOpts, arg0 uint8, arg1 [32]byte) (uint32, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "currentOperatorIndex", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurrentOperatorIndex is a free data retrieval call binding the contract method 0xe2e68580.
//
// Solidity: function currentOperatorIndex(uint8 , bytes32 ) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) CurrentOperatorIndex(arg0 uint8, arg1 [32]byte) (uint32, error) {
	return _ContractIndexRegistry.Contract.CurrentOperatorIndex(&_ContractIndexRegistry.CallOpts, arg0, arg1)
}

// CurrentOperatorIndex is a free data retrieval call binding the contract method 0xe2e68580.
//
// Solidity: function currentOperatorIndex(uint8 , bytes32 ) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) CurrentOperatorIndex(arg0 uint8, arg1 [32]byte) (uint32, error) {
	return _ContractIndexRegistry.Contract.CurrentOperatorIndex(&_ContractIndexRegistry.CallOpts, arg0, arg1)
}

// GetLatestOperatorUpdate is a free data retrieval call binding the contract method 0x12d1d74d.
//
// Solidity: function getLatestOperatorUpdate(uint8 quorumNumber, uint32 operatorIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetLatestOperatorUpdate(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32) (IIndexRegistryTypesOperatorUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getLatestOperatorUpdate", quorumNumber, operatorIndex)

	if err != nil {
		return *new(IIndexRegistryTypesOperatorUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryTypesOperatorUpdate)).(*IIndexRegistryTypesOperatorUpdate)

	return out0, err

}

// GetLatestOperatorUpdate is a free data retrieval call binding the contract method 0x12d1d74d.
//
// Solidity: function getLatestOperatorUpdate(uint8 quorumNumber, uint32 operatorIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetLatestOperatorUpdate(quorumNumber uint8, operatorIndex uint32) (IIndexRegistryTypesOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestOperatorUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex)
}

// GetLatestOperatorUpdate is a free data retrieval call binding the contract method 0x12d1d74d.
//
// Solidity: function getLatestOperatorUpdate(uint8 quorumNumber, uint32 operatorIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetLatestOperatorUpdate(quorumNumber uint8, operatorIndex uint32) (IIndexRegistryTypesOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestOperatorUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex)
}

// GetLatestQuorumUpdate is a free data retrieval call binding the contract method 0x8121906f.
//
// Solidity: function getLatestQuorumUpdate(uint8 quorumNumber) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetLatestQuorumUpdate(opts *bind.CallOpts, quorumNumber uint8) (IIndexRegistryTypesQuorumUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getLatestQuorumUpdate", quorumNumber)

	if err != nil {
		return *new(IIndexRegistryTypesQuorumUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryTypesQuorumUpdate)).(*IIndexRegistryTypesQuorumUpdate)

	return out0, err

}

// GetLatestQuorumUpdate is a free data retrieval call binding the contract method 0x8121906f.
//
// Solidity: function getLatestQuorumUpdate(uint8 quorumNumber) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetLatestQuorumUpdate(quorumNumber uint8) (IIndexRegistryTypesQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestQuorumUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// GetLatestQuorumUpdate is a free data retrieval call binding the contract method 0x8121906f.
//
// Solidity: function getLatestQuorumUpdate(uint8 quorumNumber) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetLatestQuorumUpdate(quorumNumber uint8) (IIndexRegistryTypesQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetLatestQuorumUpdate(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// GetOperatorListAtBlockNumber is a free data retrieval call binding the contract method 0x89026245.
//
// Solidity: function getOperatorListAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(bytes32[])
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetOperatorListAtBlockNumber(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32) ([][32]byte, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getOperatorListAtBlockNumber", quorumNumber, blockNumber)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetOperatorListAtBlockNumber is a free data retrieval call binding the contract method 0x89026245.
//
// Solidity: function getOperatorListAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(bytes32[])
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetOperatorListAtBlockNumber(quorumNumber uint8, blockNumber uint32) ([][32]byte, error) {
	return _ContractIndexRegistry.Contract.GetOperatorListAtBlockNumber(&_ContractIndexRegistry.CallOpts, quorumNumber, blockNumber)
}

// GetOperatorListAtBlockNumber is a free data retrieval call binding the contract method 0x89026245.
//
// Solidity: function getOperatorListAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(bytes32[])
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetOperatorListAtBlockNumber(quorumNumber uint8, blockNumber uint32) ([][32]byte, error) {
	return _ContractIndexRegistry.Contract.GetOperatorListAtBlockNumber(&_ContractIndexRegistry.CallOpts, quorumNumber, blockNumber)
}

// GetOperatorUpdateAtIndex is a free data retrieval call binding the contract method 0x2ed583e5.
//
// Solidity: function getOperatorUpdateAtIndex(uint8 quorumNumber, uint32 operatorIndex, uint32 arrayIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetOperatorUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryTypesOperatorUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getOperatorUpdateAtIndex", quorumNumber, operatorIndex, arrayIndex)

	if err != nil {
		return *new(IIndexRegistryTypesOperatorUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryTypesOperatorUpdate)).(*IIndexRegistryTypesOperatorUpdate)

	return out0, err

}

// GetOperatorUpdateAtIndex is a free data retrieval call binding the contract method 0x2ed583e5.
//
// Solidity: function getOperatorUpdateAtIndex(uint8 quorumNumber, uint32 operatorIndex, uint32 arrayIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetOperatorUpdateAtIndex(quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryTypesOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetOperatorUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex, arrayIndex)
}

// GetOperatorUpdateAtIndex is a free data retrieval call binding the contract method 0x2ed583e5.
//
// Solidity: function getOperatorUpdateAtIndex(uint8 quorumNumber, uint32 operatorIndex, uint32 arrayIndex) view returns((uint32,bytes32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetOperatorUpdateAtIndex(quorumNumber uint8, operatorIndex uint32, arrayIndex uint32) (IIndexRegistryTypesOperatorUpdate, error) {
	return _ContractIndexRegistry.Contract.GetOperatorUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, operatorIndex, arrayIndex)
}

// GetQuorumUpdateAtIndex is a free data retrieval call binding the contract method 0xa48bb0ac.
//
// Solidity: function getQuorumUpdateAtIndex(uint8 quorumNumber, uint32 quorumIndex) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCaller) GetQuorumUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, quorumIndex uint32) (IIndexRegistryTypesQuorumUpdate, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "getQuorumUpdateAtIndex", quorumNumber, quorumIndex)

	if err != nil {
		return *new(IIndexRegistryTypesQuorumUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IIndexRegistryTypesQuorumUpdate)).(*IIndexRegistryTypesQuorumUpdate)

	return out0, err

}

// GetQuorumUpdateAtIndex is a free data retrieval call binding the contract method 0xa48bb0ac.
//
// Solidity: function getQuorumUpdateAtIndex(uint8 quorumNumber, uint32 quorumIndex) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistrySession) GetQuorumUpdateAtIndex(quorumNumber uint8, quorumIndex uint32) (IIndexRegistryTypesQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetQuorumUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, quorumIndex)
}

// GetQuorumUpdateAtIndex is a free data retrieval call binding the contract method 0xa48bb0ac.
//
// Solidity: function getQuorumUpdateAtIndex(uint8 quorumNumber, uint32 quorumIndex) view returns((uint32,uint32))
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) GetQuorumUpdateAtIndex(quorumNumber uint8, quorumIndex uint32) (IIndexRegistryTypesQuorumUpdate, error) {
	return _ContractIndexRegistry.Contract.GetQuorumUpdateAtIndex(&_ContractIndexRegistry.CallOpts, quorumNumber, quorumIndex)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIndexRegistry *ContractIndexRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractIndexRegistry.Contract.RegistryCoordinator(&_ContractIndexRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractIndexRegistry.Contract.RegistryCoordinator(&_ContractIndexRegistry.CallOpts)
}

// TotalOperatorsForQuorum is a free data retrieval call binding the contract method 0xf3410922.
//
// Solidity: function totalOperatorsForQuorum(uint8 quorumNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) TotalOperatorsForQuorum(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "totalOperatorsForQuorum", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalOperatorsForQuorum is a free data retrieval call binding the contract method 0xf3410922.
//
// Solidity: function totalOperatorsForQuorum(uint8 quorumNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) TotalOperatorsForQuorum(quorumNumber uint8) (uint32, error) {
	return _ContractIndexRegistry.Contract.TotalOperatorsForQuorum(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// TotalOperatorsForQuorum is a free data retrieval call binding the contract method 0xf3410922.
//
// Solidity: function totalOperatorsForQuorum(uint8 quorumNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) TotalOperatorsForQuorum(quorumNumber uint8) (uint32, error) {
	return _ContractIndexRegistry.Contract.TotalOperatorsForQuorum(&_ContractIndexRegistry.CallOpts, quorumNumber)
}

// TotalOperatorsForQuorumAtBlockNumber is a free data retrieval call binding the contract method 0x7b5d26be.
//
// Solidity: function totalOperatorsForQuorumAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCaller) TotalOperatorsForQuorumAtBlockNumber(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32) (uint32, error) {
	var out []interface{}
	err := _ContractIndexRegistry.contract.Call(opts, &out, "totalOperatorsForQuorumAtBlockNumber", quorumNumber, blockNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TotalOperatorsForQuorumAtBlockNumber is a free data retrieval call binding the contract method 0x7b5d26be.
//
// Solidity: function totalOperatorsForQuorumAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistrySession) TotalOperatorsForQuorumAtBlockNumber(quorumNumber uint8, blockNumber uint32) (uint32, error) {
	return _ContractIndexRegistry.Contract.TotalOperatorsForQuorumAtBlockNumber(&_ContractIndexRegistry.CallOpts, quorumNumber, blockNumber)
}

// TotalOperatorsForQuorumAtBlockNumber is a free data retrieval call binding the contract method 0x7b5d26be.
//
// Solidity: function totalOperatorsForQuorumAtBlockNumber(uint8 quorumNumber, uint32 blockNumber) view returns(uint32)
func (_ContractIndexRegistry *ContractIndexRegistryCallerSession) TotalOperatorsForQuorumAtBlockNumber(quorumNumber uint8, blockNumber uint32) (uint32, error) {
	return _ContractIndexRegistry.Contract.TotalOperatorsForQuorumAtBlockNumber(&_ContractIndexRegistry.CallOpts, quorumNumber, blockNumber)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.contract.Transact(opts, "deregisterOperator", operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_ContractIndexRegistry *ContractIndexRegistrySession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.DeregisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xbd29b8cd.
//
// Solidity: function deregisterOperator(bytes32 operatorId, bytes quorumNumbers) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactorSession) DeregisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.DeregisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIndexRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIndexRegistry *ContractIndexRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.InitializeQuorum(&_ContractIndexRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractIndexRegistry *ContractIndexRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.InitializeQuorum(&_ContractIndexRegistry.TransactOpts, quorumNumber)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x00bff04d.
//
// Solidity: function registerOperator(bytes32 operatorId, bytes quorumNumbers) returns(uint32[])
func (_ContractIndexRegistry *ContractIndexRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.contract.Transact(opts, "registerOperator", operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x00bff04d.
//
// Solidity: function registerOperator(bytes32 operatorId, bytes quorumNumbers) returns(uint32[])
func (_ContractIndexRegistry *ContractIndexRegistrySession) RegisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.RegisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x00bff04d.
//
// Solidity: function registerOperator(bytes32 operatorId, bytes quorumNumbers) returns(uint32[])
func (_ContractIndexRegistry *ContractIndexRegistryTransactorSession) RegisterOperator(operatorId [32]byte, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractIndexRegistry.Contract.RegisterOperator(&_ContractIndexRegistry.TransactOpts, operatorId, quorumNumbers)
}

// ContractIndexRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractIndexRegistry contract.
type ContractIndexRegistryInitializedIterator struct {
	Event *ContractIndexRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractIndexRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIndexRegistryInitialized)
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
		it.Event = new(ContractIndexRegistryInitialized)
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
func (it *ContractIndexRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIndexRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIndexRegistryInitialized represents a Initialized event raised by the ContractIndexRegistry contract.
type ContractIndexRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractIndexRegistryInitializedIterator, error) {

	logs, sub, err := _ContractIndexRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryInitializedIterator{contract: _ContractIndexRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractIndexRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIndexRegistryInitialized)
				if err := _ContractIndexRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) ParseInitialized(log types.Log) (*ContractIndexRegistryInitialized, error) {
	event := new(ContractIndexRegistryInitialized)
	if err := _ContractIndexRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIndexRegistryQuorumIndexUpdateIterator is returned from FilterQuorumIndexUpdate and is used to iterate over the raw logs and unpacked data for QuorumIndexUpdate events raised by the ContractIndexRegistry contract.
type ContractIndexRegistryQuorumIndexUpdateIterator struct {
	Event *ContractIndexRegistryQuorumIndexUpdate // Event containing the contract specifics and raw log

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
func (it *ContractIndexRegistryQuorumIndexUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIndexRegistryQuorumIndexUpdate)
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
		it.Event = new(ContractIndexRegistryQuorumIndexUpdate)
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
func (it *ContractIndexRegistryQuorumIndexUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIndexRegistryQuorumIndexUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIndexRegistryQuorumIndexUpdate represents a QuorumIndexUpdate event raised by the ContractIndexRegistry contract.
type ContractIndexRegistryQuorumIndexUpdate struct {
	OperatorId       [32]byte
	QuorumNumber     uint8
	NewOperatorIndex uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterQuorumIndexUpdate is a free log retrieval operation binding the contract event 0x6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6.
//
// Solidity: event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newOperatorIndex)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) FilterQuorumIndexUpdate(opts *bind.FilterOpts, operatorId [][32]byte) (*ContractIndexRegistryQuorumIndexUpdateIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIndexRegistry.contract.FilterLogs(opts, "QuorumIndexUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractIndexRegistryQuorumIndexUpdateIterator{contract: _ContractIndexRegistry.contract, event: "QuorumIndexUpdate", logs: logs, sub: sub}, nil
}

// WatchQuorumIndexUpdate is a free log subscription operation binding the contract event 0x6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6.
//
// Solidity: event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newOperatorIndex)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) WatchQuorumIndexUpdate(opts *bind.WatchOpts, sink chan<- *ContractIndexRegistryQuorumIndexUpdate, operatorId [][32]byte) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _ContractIndexRegistry.contract.WatchLogs(opts, "QuorumIndexUpdate", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIndexRegistryQuorumIndexUpdate)
				if err := _ContractIndexRegistry.contract.UnpackLog(event, "QuorumIndexUpdate", log); err != nil {
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

// ParseQuorumIndexUpdate is a log parse operation binding the contract event 0x6ee1e4f4075f3d067176140d34e87874244dd273294c05b2218133e49a2ba6f6.
//
// Solidity: event QuorumIndexUpdate(bytes32 indexed operatorId, uint8 quorumNumber, uint32 newOperatorIndex)
func (_ContractIndexRegistry *ContractIndexRegistryFilterer) ParseQuorumIndexUpdate(log types.Log) (*ContractIndexRegistryQuorumIndexUpdate, error) {
	event := new(ContractIndexRegistryQuorumIndexUpdate)
	if err := _ContractIndexRegistry.contract.UnpackLog(event, "QuorumIndexUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
