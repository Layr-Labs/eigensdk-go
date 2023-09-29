// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractBLSPubkeyRegistry

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

// IBLSPubkeyRegistryApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSPubkeyRegistryApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// ContractBLSPubkeyRegistryMetaData contains all meta data concerning the ContractBLSPubkeyRegistry contract.
var ContractBLSPubkeyRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIBLSPublicKeyCompendium\",\"name\":\"_pubkeyCompendium\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"OperatorAddedToQuorums\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"OperatorRemovedFromQuorums\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"name\":\"deregisterOperator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getApkForQuorum\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApkHashForQuorumAtBlockNumberFromIndex\",\"outputs\":[{\"internalType\":\"bytes24\",\"name\":\"\",\"type\":\"bytes24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getApkIndicesForQuorumsAtBlockNumber\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getApkUpdateForQuorumByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes24\",\"name\":\"apkHash\",\"type\":\"bytes24\"},{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"}],\"internalType\":\"structIBLSPubkeyRegistry.ApkUpdate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"quorumNumber\",\"type\":\"uint8\"}],\"name\":\"getQuorumApkHistoryLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubkeyCompendium\",\"outputs\":[{\"internalType\":\"contractIBLSPublicKeyCompendium\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quorumApkUpdates\",\"outputs\":[{\"internalType\":\"bytes24\",\"name\":\"apkHash\",\"type\":\"bytes24\"},{\"internalType\":\"uint32\",\"name\":\"updateBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161149f38038061149f83398101604081905261002f9161005e565b6001600160a01b039182166080521660a052610098565b6001600160a01b038116811461005b57600080fd5b50565b6000806040838503121561007157600080fd5b825161007c81610046565b602084015190925061008d81610046565b809150509250929050565b60805160a0516113c76100d86000396000818160ce015281816103c2015261052301526000818161018a015281816102cd01526104cd01526113c76000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637225057e116100665780637225057e146101ac5780637f5eccbb146101fa578063c1af6b241461023b578063eda1076314610268578063fb81a7be1461028857600080fd5b806303ce4bad146100a3578063187548c8146100c957806324369b2a1461010857806363a945101461011b5780636d14a98714610185575b600080fd5b6100b66100b1366004610f51565b6102c0565b6040519081526020015b60405180910390f35b6100f07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c0565b6100b6610116366004610f51565b6104c0565b61016a61012936600461101b565b60408051808201909152600080825260208201525060ff16600090815260016020818152604092839020835180850190945280548452909101549082015290565b604080518251815260209283015192810192909252016100c0565b6100f07f000000000000000000000000000000000000000000000000000000000000000081565b6101bf6101ba36600461103d565b610619565b60408051825167ffffffffffffffff1916815260208084015163ffffffff9081169183019190915292820151909216908201526060016100c0565b61020d61020836600461103d565b6106aa565b6040805167ffffffffffffffff19909416845263ffffffff92831660208501529116908201526060016100c0565b61024e610249366004611067565b6106f5565b60405167ffffffffffffffff1990911681526020016100c0565b61027b6102763660046110af565b61077e565b6040516100c09190611127565b6102ab61029636600461101b565b60ff1660009081526020819052604090205490565b60405163ffffffff90911681526020016100c0565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103135760405162461bcd60e51b815260040161030a90611171565b60405180910390fd5b600061031e836109d5565b90507fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb58114156103b65760405162461bcd60e51b815260206004820152603f60248201527f424c535075626b657952656769737472792e72656769737465724f706572617460448201527f6f723a2063616e6e6f74207265676973746572207a65726f207075626b657900606482015260840161030a565b846001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6836040518263ffffffff1660e01b815260040161040e91815260200190565b602060405180830381865afa15801561042b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044f91906111e8565b6001600160a01b0316146104755760405162461bcd60e51b815260040161030a90611205565b61047f8484610a18565b7f5358c5b42179178c8fc757734ac2a3198f9071c765ee0d8389211525f500524685856040516104b0929190611263565b60405180910390a1949350505050565b6000336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461050a5760405162461bcd60e51b815260040161030a90611171565b6000610515836109d5565b9050846001600160a01b03167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6836040518263ffffffff1660e01b815260040161056f91815260200190565b602060405180830381865afa15801561058c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b091906111e8565b6001600160a01b0316146105d65760405162461bcd60e51b815260040161030a90611205565b6105e8846105e385610bd6565b610a18565b7f14a5172b312e9d2c22b8468f9c70ec2caa9de934fe380734fbc6f3beff2b14ba85856040516104b0929190611263565b604080516060810182526000808252602080830182905282840182905260ff8616825281905291909120805483908110610655576106556112c8565b600091825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b8204811694830194909452600160e01b900490921690820152905092915050565b600060205281600052604060002081815481106106c657600080fd5b600091825260209091200154604081901b925063ffffffff600160c01b820481169250600160e01b9091041683565b60ff8316600090815260208190526040812080548291908490811061071c5761071c6112c8565b600091825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b8204811694830194909452600160e01b90049092169082015290506107758185610c95565b51949350505050565b606060008367ffffffffffffffff81111561079b5761079b610ebb565b6040519080825280602002602001820160405280156107c4578160200160208202803683370190505b50905060005b848110156109cc5760008686838181106107e6576107e66112c8565b919091013560f81c600081815260208190526040902054909250905063ffffffff8116158061084f575060ff821660009081526020819052604081208054909190610833576108336112c8565b600091825260209091200154600160c01b900463ffffffff1686105b156108e85760405162461bcd60e51b815260206004820152605e60248201527f424c535075626b657952656769737472792e67657441706b496e64696365734660448201527f6f7251756f72756d734174426c6f636b4e756d6265723a20626c6f636b4e756d60648201527f626572206973206265666f726520746865206669727374207570646174650000608482015260a40161030a565b60005b8163ffffffff168163ffffffff1610156109b65760ff831660009081526020819052604090208790600161091f84866112f4565b61092991906112f4565b63ffffffff168154811061093f5761093f6112c8565b600091825260209091200154600160c01b900463ffffffff16116109a457600161096982846112f4565b61097391906112f4565b858581518110610985576109856112c8565b602002602001019063ffffffff16908163ffffffff16815250506109b6565b806109ae81611319565b9150506108eb565b50505080806109c49061133d565b9150506107ca565b50949350505050565b6000816000015182602001516040516020016109fb929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b604080518082019091526000808252602082015260005b8351811015610bd0576000848281518110610a4c57610a4c6112c8565b0160209081015160f81c6000818152918290526040909120549091508015610ac95760ff821660009081526020819052604090204390610a8d600184611358565b81548110610a9d57610a9d6112c8565b90600052602060002001600001601c6101000a81548163ffffffff021916908363ffffffff1602179055505b60ff82166000908152600160208181526040928390208351808501909452805484529091015490820152610afd9086610de2565b60ff831660009081526001602081815260408084208551815582860151930192909255815160608101835283815290810183905290810191909152909450610b44856109d5565b67ffffffffffffffff1916815263ffffffff438116602080840191825260ff90951660009081528086526040808220805460018181018355918452979092208551970180549351958201518516600160e01b026001600160e01b0396909516600160c01b026001600160e01b03199094169790911c969096179190911792909216179092555001610a2f565b50505050565b60408051808201909152600080825260208201528151158015610bfb57506020820151155b15610c19575050604080518082019091526000808252602082015290565b6040518060400160405280836000015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478460200151610c5e919061136f565b610c88907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47611358565b905292915050565b919050565b816020015163ffffffff168163ffffffff161015610d2e5760405162461bcd60e51b815260206004820152604a60248201527f424c535075626b657952656769737472792e5f76616c696461746541706b486160448201527f7368466f7251756f72756d4174426c6f636b4e756d6265723a20696e646578206064820152691d1bdbc81c9958d95b9d60b21b608482015260a40161030a565b604082015163ffffffff161580610d545750816040015163ffffffff168163ffffffff16105b610dde5760405162461bcd60e51b815260206004820152604f60248201527f424c535075626b657952656769737472792e5f76616c696461746541706b486160448201527f7368466f7251756f72756d4174426c6f636b4e756d6265723a206e6f74206c6160648201526e746573742061706b2075706461746560881b608482015260a40161030a565b5050565b6040805180820190915260008082526020820152610dfe610e85565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa9050808015610e3d57610e3f565bfe5b5080610e7d5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b604482015260640161030a565b505092915050565b60405180608001604052806004906020820280368337509192915050565b6001600160a01b0381168114610eb857600080fd5b50565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610efa57610efa610ebb565b604052919050565b600060408284031215610f1457600080fd5b6040516040810181811067ffffffffffffffff82111715610f3757610f37610ebb565b604052823581526020928301359281019290925250919050565b600080600060808486031215610f6657600080fd5b8335610f7181610ea3565b925060208481013567ffffffffffffffff80821115610f8f57600080fd5b818701915087601f830112610fa357600080fd5b813581811115610fb557610fb5610ebb565b610fc7601f8201601f19168501610ed1565b91508082528884828501011115610fdd57600080fd5b80848401858401376000848284010152508094505050506110018560408601610f02565b90509250925092565b803560ff81168114610c9057600080fd5b60006020828403121561102d57600080fd5b6110368261100a565b9392505050565b6000806040838503121561105057600080fd5b6110598361100a565b946020939093013593505050565b60008060006060848603121561107c57600080fd5b6110858461100a565b9250602084013563ffffffff8116811461109e57600080fd5b929592945050506040919091013590565b6000806000604084860312156110c457600080fd5b833567ffffffffffffffff808211156110dc57600080fd5b818601915086601f8301126110f057600080fd5b8135818111156110ff57600080fd5b87602082850101111561111157600080fd5b6020928301989097509590910135949350505050565b6020808252825182820181905260009190848201906040850190845b8181101561116557835163ffffffff1683529284019291840191600101611143565b50909695505050505050565b60208082526051908201527f424c535075626b657952656769737472792e6f6e6c795265676973747279436f60408201527f6f7264696e61746f723a2063616c6c6572206973206e6f74207468652072656760608201527034b9ba393c9031b7b7b93234b730ba37b960791b608082015260a00190565b6000602082840312156111fa57600080fd5b815161103681610ea3565b602080825260409082018190527f424c535075626b657952656769737472792e72656769737465724f7065726174908201527f6f723a206f70657261746f7220646f6573206e6f74206f776e207075626b6579606082015260800190565b60018060a01b038316815260006020604081840152835180604085015260005b8181101561129f57858101830151858201606001528201611283565b818111156112b1576000606083870101525b50601f01601f191692909201606001949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600063ffffffff83811690831681811015611311576113116112de565b039392505050565b600063ffffffff80831681811415611333576113336112de565b6001019392505050565b6000600019821415611351576113516112de565b5060010190565b60008282101561136a5761136a6112de565b500390565b60008261138c57634e487b7160e01b600052601260045260246000fd5b50069056fea2646970667358221220d86a14d6c1bc3c29952bc6ab848caedd2bbc34fd492e0cdaf5f4a087e629b55b64736f6c634300080c0033",
}

// ContractBLSPubkeyRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBLSPubkeyRegistryMetaData.ABI instead.
var ContractBLSPubkeyRegistryABI = ContractBLSPubkeyRegistryMetaData.ABI

// ContractBLSPubkeyRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBLSPubkeyRegistryMetaData.Bin instead.
var ContractBLSPubkeyRegistryBin = ContractBLSPubkeyRegistryMetaData.Bin

// DeployContractBLSPubkeyRegistry deploys a new Ethereum contract, binding an instance of ContractBLSPubkeyRegistry to it.
func DeployContractBLSPubkeyRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _pubkeyCompendium common.Address) (common.Address, *types.Transaction, *ContractBLSPubkeyRegistry, error) {
	parsed, err := ContractBLSPubkeyRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBLSPubkeyRegistryBin), backend, _registryCoordinator, _pubkeyCompendium)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBLSPubkeyRegistry{ContractBLSPubkeyRegistryCaller: ContractBLSPubkeyRegistryCaller{contract: contract}, ContractBLSPubkeyRegistryTransactor: ContractBLSPubkeyRegistryTransactor{contract: contract}, ContractBLSPubkeyRegistryFilterer: ContractBLSPubkeyRegistryFilterer{contract: contract}}, nil
}

// ContractBLSPubkeyRegistry is an auto generated Go binding around an Ethereum contract.
type ContractBLSPubkeyRegistry struct {
	ContractBLSPubkeyRegistryCaller     // Read-only binding to the contract
	ContractBLSPubkeyRegistryTransactor // Write-only binding to the contract
	ContractBLSPubkeyRegistryFilterer   // Log filterer for contract events
}

// ContractBLSPubkeyRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBLSPubkeyRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPubkeyRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBLSPubkeyRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPubkeyRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBLSPubkeyRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSPubkeyRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBLSPubkeyRegistrySession struct {
	Contract     *ContractBLSPubkeyRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractBLSPubkeyRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBLSPubkeyRegistryCallerSession struct {
	Contract *ContractBLSPubkeyRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractBLSPubkeyRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBLSPubkeyRegistryTransactorSession struct {
	Contract     *ContractBLSPubkeyRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractBLSPubkeyRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBLSPubkeyRegistryRaw struct {
	Contract *ContractBLSPubkeyRegistry // Generic contract binding to access the raw methods on
}

// ContractBLSPubkeyRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBLSPubkeyRegistryCallerRaw struct {
	Contract *ContractBLSPubkeyRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBLSPubkeyRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBLSPubkeyRegistryTransactorRaw struct {
	Contract *ContractBLSPubkeyRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBLSPubkeyRegistry creates a new instance of ContractBLSPubkeyRegistry, bound to a specific deployed contract.
func NewContractBLSPubkeyRegistry(address common.Address, backend bind.ContractBackend) (*ContractBLSPubkeyRegistry, error) {
	contract, err := bindContractBLSPubkeyRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPubkeyRegistry{ContractBLSPubkeyRegistryCaller: ContractBLSPubkeyRegistryCaller{contract: contract}, ContractBLSPubkeyRegistryTransactor: ContractBLSPubkeyRegistryTransactor{contract: contract}, ContractBLSPubkeyRegistryFilterer: ContractBLSPubkeyRegistryFilterer{contract: contract}}, nil
}

// NewContractBLSPubkeyRegistryCaller creates a new read-only instance of ContractBLSPubkeyRegistry, bound to a specific deployed contract.
func NewContractBLSPubkeyRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractBLSPubkeyRegistryCaller, error) {
	contract, err := bindContractBLSPubkeyRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPubkeyRegistryCaller{contract: contract}, nil
}

// NewContractBLSPubkeyRegistryTransactor creates a new write-only instance of ContractBLSPubkeyRegistry, bound to a specific deployed contract.
func NewContractBLSPubkeyRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBLSPubkeyRegistryTransactor, error) {
	contract, err := bindContractBLSPubkeyRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPubkeyRegistryTransactor{contract: contract}, nil
}

// NewContractBLSPubkeyRegistryFilterer creates a new log filterer instance of ContractBLSPubkeyRegistry, bound to a specific deployed contract.
func NewContractBLSPubkeyRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBLSPubkeyRegistryFilterer, error) {
	contract, err := bindContractBLSPubkeyRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBLSPubkeyRegistryFilterer{contract: contract}, nil
}

// bindContractBLSPubkeyRegistry binds a generic wrapper to an already deployed contract.
func bindContractBLSPubkeyRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBLSPubkeyRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSPubkeyRegistry.Contract.ContractBLSPubkeyRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.ContractBLSPubkeyRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.ContractBLSPubkeyRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSPubkeyRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetApkForQuorum is a free data retrieval call binding the contract method 0x63a94510.
//
// Solidity: function getApkForQuorum(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) GetApkForQuorum(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "getApkForQuorum", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApkForQuorum is a free data retrieval call binding the contract method 0x63a94510.
//
// Solidity: function getApkForQuorum(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) GetApkForQuorum(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkForQuorum(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber)
}

// GetApkForQuorum is a free data retrieval call binding the contract method 0x63a94510.
//
// Solidity: function getApkForQuorum(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) GetApkForQuorum(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkForQuorum(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber)
}

// GetApkHashForQuorumAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc1af6b24.
//
// Solidity: function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) GetApkHashForQuorumAtBlockNumberFromIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "getApkHashForQuorumAtBlockNumberFromIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashForQuorumAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc1af6b24.
//
// Solidity: function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) GetApkHashForQuorumAtBlockNumberFromIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkHashForQuorumAtBlockNumberFromIndex(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashForQuorumAtBlockNumberFromIndex is a free data retrieval call binding the contract method 0xc1af6b24.
//
// Solidity: function getApkHashForQuorumAtBlockNumberFromIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) GetApkHashForQuorumAtBlockNumberFromIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkHashForQuorumAtBlockNumberFromIndex(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkIndicesForQuorumsAtBlockNumber is a free data retrieval call binding the contract method 0xeda10763.
//
// Solidity: function getApkIndicesForQuorumsAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) GetApkIndicesForQuorumsAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "getApkIndicesForQuorumsAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesForQuorumsAtBlockNumber is a free data retrieval call binding the contract method 0xeda10763.
//
// Solidity: function getApkIndicesForQuorumsAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) GetApkIndicesForQuorumsAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkIndicesForQuorumsAtBlockNumber(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesForQuorumsAtBlockNumber is a free data retrieval call binding the contract method 0xeda10763.
//
// Solidity: function getApkIndicesForQuorumsAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) GetApkIndicesForQuorumsAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkIndicesForQuorumsAtBlockNumber(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateForQuorumByIndex is a free data retrieval call binding the contract method 0x7225057e.
//
// Solidity: function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) GetApkUpdateForQuorumByIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSPubkeyRegistryApkUpdate, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "getApkUpdateForQuorumByIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSPubkeyRegistryApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSPubkeyRegistryApkUpdate)).(*IBLSPubkeyRegistryApkUpdate)

	return out0, err

}

// GetApkUpdateForQuorumByIndex is a free data retrieval call binding the contract method 0x7225057e.
//
// Solidity: function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) GetApkUpdateForQuorumByIndex(quorumNumber uint8, index *big.Int) (IBLSPubkeyRegistryApkUpdate, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkUpdateForQuorumByIndex(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateForQuorumByIndex is a free data retrieval call binding the contract method 0x7225057e.
//
// Solidity: function getApkUpdateForQuorumByIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) GetApkUpdateForQuorumByIndex(quorumNumber uint8, index *big.Int) (IBLSPubkeyRegistryApkUpdate, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetApkUpdateForQuorumByIndex(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber, index)
}

// GetQuorumApkHistoryLength is a free data retrieval call binding the contract method 0xfb81a7be.
//
// Solidity: function getQuorumApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) GetQuorumApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "getQuorumApkHistoryLength", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetQuorumApkHistoryLength is a free data retrieval call binding the contract method 0xfb81a7be.
//
// Solidity: function getQuorumApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) GetQuorumApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetQuorumApkHistoryLength(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber)
}

// GetQuorumApkHistoryLength is a free data retrieval call binding the contract method 0xfb81a7be.
//
// Solidity: function getQuorumApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) GetQuorumApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSPubkeyRegistry.Contract.GetQuorumApkHistoryLength(&_ContractBLSPubkeyRegistry.CallOpts, quorumNumber)
}

// PubkeyCompendium is a free data retrieval call binding the contract method 0x187548c8.
//
// Solidity: function pubkeyCompendium() view returns(address)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) PubkeyCompendium(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "pubkeyCompendium")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyCompendium is a free data retrieval call binding the contract method 0x187548c8.
//
// Solidity: function pubkeyCompendium() view returns(address)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) PubkeyCompendium() (common.Address, error) {
	return _ContractBLSPubkeyRegistry.Contract.PubkeyCompendium(&_ContractBLSPubkeyRegistry.CallOpts)
}

// PubkeyCompendium is a free data retrieval call binding the contract method 0x187548c8.
//
// Solidity: function pubkeyCompendium() view returns(address)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) PubkeyCompendium() (common.Address, error) {
	return _ContractBLSPubkeyRegistry.Contract.PubkeyCompendium(&_ContractBLSPubkeyRegistry.CallOpts)
}

// QuorumApkUpdates is a free data retrieval call binding the contract method 0x7f5eccbb.
//
// Solidity: function quorumApkUpdates(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) QuorumApkUpdates(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "quorumApkUpdates", arg0, arg1)

	outstruct := new(struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApkHash = *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)
	outstruct.UpdateBlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.NextUpdateBlockNumber = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// QuorumApkUpdates is a free data retrieval call binding the contract method 0x7f5eccbb.
//
// Solidity: function quorumApkUpdates(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) QuorumApkUpdates(arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSPubkeyRegistry.Contract.QuorumApkUpdates(&_ContractBLSPubkeyRegistry.CallOpts, arg0, arg1)
}

// QuorumApkUpdates is a free data retrieval call binding the contract method 0x7f5eccbb.
//
// Solidity: function quorumApkUpdates(uint8 , uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) QuorumApkUpdates(arg0 uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSPubkeyRegistry.Contract.QuorumApkUpdates(&_ContractBLSPubkeyRegistry.CallOpts, arg0, arg1)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSPubkeyRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSPubkeyRegistry.Contract.RegistryCoordinator(&_ContractBLSPubkeyRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSPubkeyRegistry.Contract.RegistryCoordinator(&_ContractBLSPubkeyRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x24369b2a.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers, (uint256,uint256) pubkey) returns(bytes32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte, pubkey BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers, pubkey)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x24369b2a.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers, (uint256,uint256) pubkey) returns(bytes32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte, pubkey BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.DeregisterOperator(&_ContractBLSPubkeyRegistry.TransactOpts, operator, quorumNumbers, pubkey)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x24369b2a.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers, (uint256,uint256) pubkey) returns(bytes32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte, pubkey BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.DeregisterOperator(&_ContractBLSPubkeyRegistry.TransactOpts, operator, quorumNumbers, pubkey)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x03ce4bad.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers, (uint256,uint256) pubkey) returns(bytes32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte, pubkey BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers, pubkey)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x03ce4bad.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers, (uint256,uint256) pubkey) returns(bytes32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte, pubkey BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.RegisterOperator(&_ContractBLSPubkeyRegistry.TransactOpts, operator, quorumNumbers, pubkey)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x03ce4bad.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers, (uint256,uint256) pubkey) returns(bytes32)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte, pubkey BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSPubkeyRegistry.Contract.RegisterOperator(&_ContractBLSPubkeyRegistry.TransactOpts, operator, quorumNumbers, pubkey)
}

// ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the ContractBLSPubkeyRegistry contract.
type ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator struct {
	Event *ContractBLSPubkeyRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSPubkeyRegistryOperatorAddedToQuorums)
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
		it.Event = new(ContractBLSPubkeyRegistryOperatorAddedToQuorums)
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
func (it *ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSPubkeyRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the ContractBLSPubkeyRegistry contract.
type ContractBLSPubkeyRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x5358c5b42179178c8fc757734ac2a3198f9071c765ee0d8389211525f5005246.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes quorumNumbers)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _ContractBLSPubkeyRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSPubkeyRegistryOperatorAddedToQuorumsIterator{contract: _ContractBLSPubkeyRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x5358c5b42179178c8fc757734ac2a3198f9071c765ee0d8389211525f5005246.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes quorumNumbers)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSPubkeyRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSPubkeyRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSPubkeyRegistryOperatorAddedToQuorums)
				if err := _ContractBLSPubkeyRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x5358c5b42179178c8fc757734ac2a3198f9071c765ee0d8389211525f5005246.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes quorumNumbers)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSPubkeyRegistryOperatorAddedToQuorums, error) {
	event := new(ContractBLSPubkeyRegistryOperatorAddedToQuorums)
	if err := _ContractBLSPubkeyRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the ContractBLSPubkeyRegistry contract.
type ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *ContractBLSPubkeyRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSPubkeyRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(ContractBLSPubkeyRegistryOperatorRemovedFromQuorums)
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
func (it *ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSPubkeyRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the ContractBLSPubkeyRegistry contract.
type ContractBLSPubkeyRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0x14a5172b312e9d2c22b8468f9c70ec2caa9de934fe380734fbc6f3beff2b14ba.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes quorumNumbers)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _ContractBLSPubkeyRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSPubkeyRegistryOperatorRemovedFromQuorumsIterator{contract: _ContractBLSPubkeyRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0x14a5172b312e9d2c22b8468f9c70ec2caa9de934fe380734fbc6f3beff2b14ba.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes quorumNumbers)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSPubkeyRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSPubkeyRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSPubkeyRegistryOperatorRemovedFromQuorums)
				if err := _ContractBLSPubkeyRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0x14a5172b312e9d2c22b8468f9c70ec2caa9de934fe380734fbc6f3beff2b14ba.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes quorumNumbers)
func (_ContractBLSPubkeyRegistry *ContractBLSPubkeyRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSPubkeyRegistryOperatorRemovedFromQuorums, error) {
	event := new(ContractBLSPubkeyRegistryOperatorRemovedFromQuorums)
	if err := _ContractBLSPubkeyRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
