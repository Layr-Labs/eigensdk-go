// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractEigenPodManager

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

// ContractEigenPodManagerMetaData contains all meta data concerning the ContractEigenPodManager contract.
var ContractEigenPodManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_ethPOS\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"},{\"name\":\"_eigenPodBeacon\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"},{\"name\":\"_strategyManager\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"},{\"name\":\"_delegationManager\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOwnedShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"OwnedShares\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beaconChainETHStrategy\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenPodBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethPOS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIETHPOSDeposit\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPod\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"_initPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"numPods\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToPod\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPod\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAll\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauserRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"podOwnerShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordBeaconChainETHBalanceUpdate\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"proportionOfOldBalance\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeShares\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"Shares\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPauserRegistry\",\"inputs\":[{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"internalType\":\"contractIPauserRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakerStrategyShares\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"}],\"outputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"Shares\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"strategyManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStrategyManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSharesAsTokens\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"strategy\",\"type\":\"address\",\"internalType\":\"contractIStrategy\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"OwnedShares\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BeaconChainETHDeposited\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconChainETHWithdrawalCompleted\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"delegatedAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"withdrawalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewTotalShares\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTotalShares\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PauserRegistrySet\",\"inputs\":[{\"name\":\"pauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"},{\"name\":\"newPauserRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"contractIPauserRegistry\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodDeployed\",\"inputs\":[{\"name\":\"eigenPod\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PodSharesUpdated\",\"inputs\":[{\"name\":\"podOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sharesDelta\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newPausedStatus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CurrentlyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EigenPodAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InputAddressZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNewPausedStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidStrategy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyWithdrawalsNotCompleted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyDelegationManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyEigenPod\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyPauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyUnpauser\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesNegative\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharesNotMultipleOfGwei\",\"inputs\":[]}]",
	Bin: "0x61010060405234801561001157600080fd5b5060405161276f38038061276f83398101604081905261003091610137565b6001600160a01b0380851660805280841660a05280831660c052811660e052610057610060565b50505050610196565b600054610100900460ff16156100cc5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff9081161461011d576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461013457600080fd5b50565b6000806000806080858703121561014d57600080fd5b84516101588161011f565b60208601519094506101698161011f565b604086015190935061017a8161011f565b606086015190925061018b8161011f565b939692955090935050565b60805160a05160c05160e05161256f61020060003960008181610539015281816106ea01528181610a9001528181610cb101528181610fbc015261130b015260006102c101526000818161025001528181610f3c01526115de015260006103ab015261256f6000f3fe6080604052600436106101b75760003560e01c806384d81062116100ec578063a38406a31161008a578063ea4d3c9b11610064578063ea4d3c9b14610527578063f2fde38b1461055b578063f6848d241461057b578063fabc1cbc146105b657600080fd5b8063a38406a3146104d1578063a6a509be146104f1578063a9095e941461050757600080fd5b80638da5cb5b116100c65780638da5cb5b146104425780639104c319146104605780639b4e4634146104885780639ba062751461049b57600080fd5b806384d81062146103ed578063886f1195146104025780638c80d4e51461042257600080fd5b8063595c6a671161015957806360f4062b1161013357806360f4062b14610357578063715018a61461038457806374cdd798146103995780637a7e0d92146103cd57600080fd5b8063595c6a67146102e35780635ac86ab7146102f85780635c975abb1461033857600080fd5b80631794bb3c116101955780631794bb3c1461021e578063292b7b2b1461023e5780632eae418c1461028f57806339b70e38146102af57600080fd5b8063095e210c146101bc57806310d67a2f146101de578063136439dd146101fe575b600080fd5b3480156101c857600080fd5b506101dc6101d736600461184b565b6105d6565b005b3480156101ea57600080fd5b506101dc6101f936600461189a565b6107c3565b34801561020a57600080fd5b506101dc6102193660046118b7565b610877565b34801561022a57600080fd5b506101dc6102393660046118d0565b610962565b34801561024a57600080fd5b506102727f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561029b57600080fd5b506101dc6102aa366004611911565b610a85565b3480156102bb57600080fd5b506102727f000000000000000000000000000000000000000000000000000000000000000081565b3480156102ef57600080fd5b506101dc610ace565b34801561030457600080fd5b50610328610313366004611962565b606654600160ff9092169190911b9081161490565b6040519015158152602001610286565b34801561034457600080fd5b506066545b604051908152602001610286565b34801561036357600080fd5b5061034961037236600461189a565b609b6020526000908152604090205481565b34801561039057600080fd5b506101dc610b96565b3480156103a557600080fd5b506102727f000000000000000000000000000000000000000000000000000000000000000081565b3480156103d957600080fd5b506103496103e8366004611985565b610baa565b3480156103f957600080fd5b50610272610c33565b34801561040e57600080fd5b50606554610272906001600160a01b031681565b34801561042e57600080fd5b506101dc61043d3660046118d0565b610ca6565b34801561044e57600080fd5b506033546001600160a01b0316610272565b34801561046c57600080fd5b5061027273beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac081565b6101dc610496366004611a07565b610e1d565b3480156104a757600080fd5b506102726104b636600461189a565b6098602052600090815260409020546001600160a01b031681565b3480156104dd57600080fd5b506102726104ec36600461189a565b610ee0565b3480156104fd57600080fd5b5061034960995481565b34801561051357600080fd5b506101dc610522366004611911565b610fb1565b34801561053357600080fd5b506102727f000000000000000000000000000000000000000000000000000000000000000081565b34801561056757600080fd5b506101dc61057636600461189a565b611041565b34801561058757600080fd5b5061032861059636600461189a565b6001600160a01b0390811660009081526098602052604090205416151590565b3480156105c257600080fd5b506101dc6105d13660046118b7565b6110b7565b6001600160a01b038084166000908152609860205260409020548491163314610612576040516312e16d7160e11b815260040160405180910390fd5b61061a6111bf565b6001600160a01b038416610641576040516339b190bb60e11b815260040160405180910390fd5b61064f633b9aca0084611a96565b1561066d576040516347d072bb60e11b815260040160405180910390fd5b6001600160a01b0384166000908152609b602052604081205412156106a557604051634b692bcf60e01b815260040160405180910390fd5b60008313156106bd576106b88484611218565b6107b3565b6000831280156106e357506001600160a01b0384166000908152609b6020526040812054135b156107b3577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635d9aed2385610744609b6000896001600160a01b03166001600160a01b031681526020019081526020016000205490565b6040516001600160e01b031960e085901b1681526001600160a01b039092166004830152602482015267ffffffffffffffff85166044820152606401600060405180830381600087803b15801561079a57600080fd5b505af11580156107ae573d6000803e3d6000fd5b505050505b6107bd600160c955565b50505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610816573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083a9190611aaa565b6001600160a01b0316336001600160a01b03161461086b5760405163794821ff60e01b815260040160405180910390fd5b610874816113d5565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156108bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108e39190611ac7565b61090057604051631d77d47760e21b815260040160405180910390fd5b606654818116146109245760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b600054610100900460ff16158080156109825750600054600160ff909116105b8061099c5750303b15801561099c575060005460ff166001145b610a045760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610a27576000805461ff0019166101001790555b610a3084611465565b610a3a83836114b7565b80156107bd576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107bd5760405163f739589b60e01b815260040160405180910390fd5b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa158015610b16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3a9190611ac7565b610b5757604051631d77d47760e21b815260040160405180910390fd5b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b610b9e61153c565b610ba86000611465565b565b60006001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014610be957604051632711b74d60e11b815260040160405180910390fd5b6001600160a01b0383166000908152609b6020526040812054610c2c9113610c27576001600160a01b0384166000908152609b602052604090205490565b600090565b9392505050565b6066546000908190600190811603610c5e5760405163840a48d560e01b815260040160405180910390fd5b336000908152609860205260409020546001600160a01b031615610c955760405163031a852160e21b815260040160405180910390fd5b6000610c9f611596565b9250505090565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610cef5760405163f739589b60e01b815260040160405180910390fd5b6001600160a01b03821673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac014610d2c57604051632711b74d60e11b815260040160405180910390fd5b6000811215610d4e5760405163ef147de160e01b815260040160405180910390fd5b610d5c633b9aca0082611ae9565b15610d7a576040516347d072bb60e11b815260040160405180910390fd5b6001600160a01b0383166000908152609b6020526040812054610d9e908390611b13565b90506000811215610dc25760405163ef147de160e01b815260040160405180910390fd5b6001600160a01b0384166000818152609b602052604090819020839055517fd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe07709890610e0f9084815260200190565b60405180910390a250505050565b606654600090600190811603610e465760405163840a48d560e01b815260040160405180910390fd5b336000908152609860205260409020546001600160a01b031680610e6f57610e6c611596565b90505b6040516326d3918d60e21b81526001600160a01b03821690639b4e4634903490610ea5908b908b908b908b908b90600401611b63565b6000604051808303818588803b158015610ebe57600080fd5b505af1158015610ed2573d6000803e3d6000fd5b505050505050505050505050565b6001600160a01b0380821660009081526098602052604081205490911680610fab57610c2c836001600160a01b031660001b60405180610940016040528061090e8152602001611c2c61090e9139604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166020820152808201919091526000606082015260800160408051601f1981840301815290829052610f909291602001611bcd565b604051602081830303815290604052805190602001206116fb565b92915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610ffa5760405163f739589b60e01b815260040160405180910390fd5b6001600160a01b03831673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac01461103757604051632711b74d60e11b815260040160405180910390fd5b6107bd8482611218565b61104961153c565b6001600160a01b0381166110ae5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016109fb565b61087481611465565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561110a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061112e9190611aaa565b6001600160a01b0316336001600160a01b03161461115f5760405163794821ff60e01b815260040160405180910390fd5b6066541981196066541916146111885760405163c61dca5d60e01b815260040160405180910390fd5b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c90602001610957565b600260c954036112115760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016109fb565b600260c955565b6001600160a01b03821661123f576040516339b190bb60e11b815260040160405180910390fd5b6001600160a01b0382166000908152609b602052604081205482916112648383611bea565b6001600160a01b0386166000818152609b60205260409081902083905551919250907f4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193906112b59086815260200190565b60405180910390a2846001600160a01b03167fd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098826040516112f891815260200190565b60405180910390a260008113156113ce577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633c651cf28673beac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac06000861261135c578561135f565b60005b6040516001600160e01b031960e086901b1681526001600160a01b039384166004820152929091166024830152604482015260648101879052608401600060405180830381600087803b1580156113b557600080fd5b505af11580156113c9573d6000803e3d6000fd5b505050505b5050505050565b6001600160a01b0381166113fc576040516339b190bb60e11b815260040160405180910390fd5b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6065546001600160a01b03161580156114d857506001600160a01b03821615155b6114f5576040516339b190bb60e11b815260040160405180910390fd5b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2611538826113d5565b5050565b6033546001600160a01b03163314610ba85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109fb565b60006099600081546115a790611c12565b9091555060408051610940810190915261090e8082526000916116469183913391611c2c6020830139604080516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166020820152808201919091526000606082015260800160408051601f19818403018152908290526116329291602001611bcd565b604051602081830303815290604052611708565b60405163189acdbd60e31b81523360048201529091506001600160a01b0382169063c4d66de890602401600060405180830381600087803b15801561168a57600080fd5b505af115801561169e573d6000803e3d6000fd5b50503360008181526098602052604080822080546001600160a01b0319166001600160a01b038816908117909155905192945092507f21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a91a3919050565b6000610c2c83833061180c565b60008347101561175a5760405162461bcd60e51b815260206004820152601d60248201527f437265617465323a20696e73756666696369656e742062616c616e636500000060448201526064016109fb565b81516000036117ab5760405162461bcd60e51b815260206004820181905260248201527f437265617465323a2062797465636f6465206c656e677468206973207a65726f60448201526064016109fb565b8282516020840186f590506001600160a01b038116610c2c5760405162461bcd60e51b815260206004820152601960248201527f437265617465323a204661696c6564206f6e206465706c6f790000000000000060448201526064016109fb565b6000604051836040820152846020820152828152600b8101905060ff815360559020949350505050565b6001600160a01b038116811461087457600080fd5b60008060006060848603121561186057600080fd5b833561186b81611836565b925060208401359150604084013567ffffffffffffffff8116811461188f57600080fd5b809150509250925092565b6000602082840312156118ac57600080fd5b8135610c2c81611836565b6000602082840312156118c957600080fd5b5035919050565b6000806000606084860312156118e557600080fd5b83356118f081611836565b9250602084013561190081611836565b929592945050506040919091013590565b6000806000806080858703121561192757600080fd5b843561193281611836565b9350602085013561194281611836565b9250604085013561195281611836565b9396929550929360600135925050565b60006020828403121561197457600080fd5b813560ff81168114610c2c57600080fd5b6000806040838503121561199857600080fd5b82356119a381611836565b915060208301356119b381611836565b809150509250929050565b60008083601f8401126119d057600080fd5b50813567ffffffffffffffff8111156119e857600080fd5b602083019150836020828501011115611a0057600080fd5b9250929050565b600080600080600060608688031215611a1f57600080fd5b853567ffffffffffffffff811115611a3657600080fd5b611a42888289016119be565b909650945050602086013567ffffffffffffffff811115611a6257600080fd5b611a6e888289016119be565b96999598509660400135949350505050565b634e487b7160e01b600052601260045260246000fd5b600082611aa557611aa5611a80565b500790565b600060208284031215611abc57600080fd5b8151610c2c81611836565b600060208284031215611ad957600080fd5b81518015158114610c2c57600080fd5b600082611af857611af8611a80565b500690565b634e487b7160e01b600052601160045260246000fd5b8181036000831280158383131683831282161715611b3357611b33611afd565b5092915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b606081526000611b77606083018789611b3a565b8281036020840152611b8a818688611b3a565b9150508260408301529695505050505050565b6000815160005b81811015611bbe5760208185018101518683015201611ba4565b50600093019283525090919050565b6000611be2611bdc8386611b9d565b84611b9d565b949350505050565b8082018281126000831280158216821582161715611c0a57611c0a611afd565b505092915050565b600060018201611c2457611c24611afd565b506001019056fe608060405260405161090e38038061090e83398101604081905261002291610460565b61002e82826000610035565b505061058a565b61003e83610100565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a260008251118061007f5750805b156100fb576100f9836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e99190610520565b836102a360201b6100291760201c565b505b505050565b610113816102cf60201b6100551760201c565b6101725760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101e6816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101d79190610520565b6102cf60201b6100551760201c565b61024b5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b6064820152608401610169565b806102827fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102de60201b6100641760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102c883836040518060600160405280602781526020016108e7602791396102e1565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b0316856040516102fe919061053b565b600060405180830381855af49150503d8060008114610339576040519150601f19603f3d011682016040523d82523d6000602084013e61033e565b606091505b5090925090506103508683838761035a565b9695505050505050565b606083156103c65782516103bf576001600160a01b0385163b6103bf5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610169565b50816103d0565b6103d083836103d8565b949350505050565b8151156103e85781518083602001fd5b8060405162461bcd60e51b81526004016101699190610557565b80516001600160a01b038116811461041957600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561044f578181015183820152602001610437565b838111156100f95750506000910152565b6000806040838503121561047357600080fd5b61047c83610402565b60208401519092506001600160401b038082111561049957600080fd5b818501915085601f8301126104ad57600080fd5b8151818111156104bf576104bf61041e565b604051601f8201601f19908116603f011681019083821181831017156104e7576104e761041e565b8160405282815288602084870101111561050057600080fd5b610511836020830160208801610434565b80955050505050509250929050565b60006020828403121561053257600080fd5b6102c882610402565b6000825161054d818460208701610434565b9190910192915050565b6020815260008251806020840152610576816040850160208701610434565b601f01601f19169190910160400192915050565b61034e806105996000396000f3fe60806040523661001357610011610017565b005b6100115b610027610022610067565b610100565b565b606061004e83836040518060600160405280602781526020016102f260279139610124565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100fb9190610249565b905090565b3660008037600080366000845af43d6000803e80801561011f573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161014191906102a2565b600060405180830381855af49150503d806000811461017c576040519150601f19603f3d011682016040523d82523d6000602084013e610181565b606091505b50915091506101928683838761019c565b9695505050505050565b6060831561020d578251610206576001600160a01b0385163b6102065760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610217565b610217838361021f565b949350505050565b81511561022f5781518083602001fd5b8060405162461bcd60e51b81526004016101fd91906102be565b60006020828403121561025b57600080fd5b81516001600160a01b038116811461004e57600080fd5b60005b8381101561028d578181015183820152602001610275565b8381111561029c576000848401525b50505050565b600082516102b4818460208701610272565b9190910192915050565b60208152600082518060208401526102dd816040850160208701610272565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d51e81d3bc5ed20a26aeb05dce7e825c503b2061aa78628027300c8d65b9d89a64736f6c634300080c0033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e37f39584c7f15b7a034232e593ee3508e12f4b9c86800e991fbaefd381bd9e864736f6c634300081b0033",
}

// ContractEigenPodManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractEigenPodManagerMetaData.ABI instead.
var ContractEigenPodManagerABI = ContractEigenPodManagerMetaData.ABI

// ContractEigenPodManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractEigenPodManagerMetaData.Bin instead.
var ContractEigenPodManagerBin = ContractEigenPodManagerMetaData.Bin

// DeployContractEigenPodManager deploys a new Ethereum contract, binding an instance of ContractEigenPodManager to it.
func DeployContractEigenPodManager(auth *bind.TransactOpts, backend bind.ContractBackend, _ethPOS common.Address, _eigenPodBeacon common.Address, _strategyManager common.Address, _delegationManager common.Address) (common.Address, *types.Transaction, *ContractEigenPodManager, error) {
	parsed, err := ContractEigenPodManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractEigenPodManagerBin), backend, _ethPOS, _eigenPodBeacon, _strategyManager, _delegationManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractEigenPodManager{ContractEigenPodManagerCaller: ContractEigenPodManagerCaller{contract: contract}, ContractEigenPodManagerTransactor: ContractEigenPodManagerTransactor{contract: contract}, ContractEigenPodManagerFilterer: ContractEigenPodManagerFilterer{contract: contract}}, nil
}

// ContractEigenPodManagerMethods is an auto generated interface around an Ethereum contract.
type ContractEigenPodManagerMethods interface {
	ContractEigenPodManagerCalls
	ContractEigenPodManagerTransacts
	ContractEigenPodManagerFilters
}

// ContractEigenPodManagerCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractEigenPodManagerCalls interface {
	BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error)

	DelegationManager(opts *bind.CallOpts) (common.Address, error)

	EigenPodBeacon(opts *bind.CallOpts) (common.Address, error)

	EthPOS(opts *bind.CallOpts) (common.Address, error)

	GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error)

	HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error)

	NumPods(opts *bind.CallOpts) (*big.Int, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	OwnerToPod(opts *bind.CallOpts, arg0 common.Address) (common.Address, error)

	Paused(opts *bind.CallOpts, index uint8) (bool, error)

	Paused0(opts *bind.CallOpts) (*big.Int, error)

	PauserRegistry(opts *bind.CallOpts) (common.Address, error)

	PodOwnerShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error)

	StakerStrategyShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error)

	StrategyManager(opts *bind.CallOpts) (common.Address, error)
}

// ContractEigenPodManagerTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractEigenPodManagerTransacts interface {
	AddOwnedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error)

	CreatePod(opts *bind.TransactOpts) (*types.Transaction, error)

	Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, _initPausedStatus *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	PauseAll(opts *bind.TransactOpts) (*types.Transaction, error)

	RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, sharesDelta *big.Int, proportionOfOldBalance uint64) (*types.Transaction, error)

	RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error)

	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error)

	Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error)

	WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error)
}

// ContractEigenPodManagerFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractEigenPodManagerFilters interface {
	FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerBeaconChainETHDepositedIterator, error)
	WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error)
	ParseBeaconChainETHDeposited(log types.Log) (*ContractEigenPodManagerBeaconChainETHDeposited, error)

	FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator, error)
	WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error)
	ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*ContractEigenPodManagerBeaconChainETHWithdrawalCompleted, error)

	FilterInitialized(opts *bind.FilterOpts) (*ContractEigenPodManagerInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractEigenPodManagerInitialized, error)

	FilterNewTotalShares(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerNewTotalSharesIterator, error)
	WatchNewTotalShares(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerNewTotalShares, podOwner []common.Address) (event.Subscription, error)
	ParseNewTotalShares(log types.Log) (*ContractEigenPodManagerNewTotalShares, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractEigenPodManagerOwnershipTransferredIterator, error)
	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error)
	ParseOwnershipTransferred(log types.Log) (*ContractEigenPodManagerOwnershipTransferred, error)

	FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractEigenPodManagerPausedIterator, error)
	WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPaused, account []common.Address) (event.Subscription, error)
	ParsePaused(log types.Log) (*ContractEigenPodManagerPaused, error)

	FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractEigenPodManagerPauserRegistrySetIterator, error)
	WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPauserRegistrySet) (event.Subscription, error)
	ParsePauserRegistrySet(log types.Log) (*ContractEigenPodManagerPauserRegistrySet, error)

	FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*ContractEigenPodManagerPodDeployedIterator, error)
	WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error)
	ParsePodDeployed(log types.Log) (*ContractEigenPodManagerPodDeployed, error)

	FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerPodSharesUpdatedIterator, error)
	WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPodSharesUpdated, podOwner []common.Address) (event.Subscription, error)
	ParsePodSharesUpdated(log types.Log) (*ContractEigenPodManagerPodSharesUpdated, error)

	FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractEigenPodManagerUnpausedIterator, error)
	WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerUnpaused, account []common.Address) (event.Subscription, error)
	ParseUnpaused(log types.Log) (*ContractEigenPodManagerUnpaused, error)
}

// ContractEigenPodManager is an auto generated Go binding around an Ethereum contract.
type ContractEigenPodManager struct {
	ContractEigenPodManagerCaller     // Read-only binding to the contract
	ContractEigenPodManagerTransactor // Write-only binding to the contract
	ContractEigenPodManagerFilterer   // Log filterer for contract events
}

// ContractEigenPodManager implements the ContractEigenPodManagerMethods interface.
var _ ContractEigenPodManagerMethods = (*ContractEigenPodManager)(nil)

// ContractEigenPodManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractEigenPodManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEigenPodManagerCaller implements the ContractEigenPodManagerCalls interface.
var _ ContractEigenPodManagerCalls = (*ContractEigenPodManagerCaller)(nil)

// ContractEigenPodManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractEigenPodManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEigenPodManagerTransactor implements the ContractEigenPodManagerTransacts interface.
var _ ContractEigenPodManagerTransacts = (*ContractEigenPodManagerTransactor)(nil)

// ContractEigenPodManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractEigenPodManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractEigenPodManagerFilterer implements the ContractEigenPodManagerFilters interface.
var _ ContractEigenPodManagerFilters = (*ContractEigenPodManagerFilterer)(nil)

// ContractEigenPodManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractEigenPodManagerSession struct {
	Contract     *ContractEigenPodManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ContractEigenPodManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractEigenPodManagerCallerSession struct {
	Contract *ContractEigenPodManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ContractEigenPodManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractEigenPodManagerTransactorSession struct {
	Contract     *ContractEigenPodManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ContractEigenPodManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractEigenPodManagerRaw struct {
	Contract *ContractEigenPodManager // Generic contract binding to access the raw methods on
}

// ContractEigenPodManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractEigenPodManagerCallerRaw struct {
	Contract *ContractEigenPodManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractEigenPodManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractEigenPodManagerTransactorRaw struct {
	Contract *ContractEigenPodManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractEigenPodManager creates a new instance of ContractEigenPodManager, bound to a specific deployed contract.
func NewContractEigenPodManager(address common.Address, backend bind.ContractBackend) (*ContractEigenPodManager, error) {
	contract, err := bindContractEigenPodManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManager{ContractEigenPodManagerCaller: ContractEigenPodManagerCaller{contract: contract}, ContractEigenPodManagerTransactor: ContractEigenPodManagerTransactor{contract: contract}, ContractEigenPodManagerFilterer: ContractEigenPodManagerFilterer{contract: contract}}, nil
}

// NewContractEigenPodManagerCaller creates a new read-only instance of ContractEigenPodManager, bound to a specific deployed contract.
func NewContractEigenPodManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractEigenPodManagerCaller, error) {
	contract, err := bindContractEigenPodManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerCaller{contract: contract}, nil
}

// NewContractEigenPodManagerTransactor creates a new write-only instance of ContractEigenPodManager, bound to a specific deployed contract.
func NewContractEigenPodManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractEigenPodManagerTransactor, error) {
	contract, err := bindContractEigenPodManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerTransactor{contract: contract}, nil
}

// NewContractEigenPodManagerFilterer creates a new log filterer instance of ContractEigenPodManager, bound to a specific deployed contract.
func NewContractEigenPodManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractEigenPodManagerFilterer, error) {
	contract, err := bindContractEigenPodManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerFilterer{contract: contract}, nil
}

// bindContractEigenPodManager binds a generic wrapper to an already deployed contract.
func bindContractEigenPodManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractEigenPodManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractEigenPodManager *ContractEigenPodManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractEigenPodManager.Contract.ContractEigenPodManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractEigenPodManager *ContractEigenPodManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.ContractEigenPodManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractEigenPodManager *ContractEigenPodManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.ContractEigenPodManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractEigenPodManager *ContractEigenPodManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractEigenPodManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.contract.Transact(opts, method, params...)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) BeaconChainETHStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "beaconChainETHStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractEigenPodManager.Contract.BeaconChainETHStrategy(&_ContractEigenPodManager.CallOpts)
}

// BeaconChainETHStrategy is a free data retrieval call binding the contract method 0x9104c319.
//
// Solidity: function beaconChainETHStrategy() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) BeaconChainETHStrategy() (common.Address, error) {
	return _ContractEigenPodManager.Contract.BeaconChainETHStrategy(&_ContractEigenPodManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) DelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "delegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) DelegationManager() (common.Address, error) {
	return _ContractEigenPodManager.Contract.DelegationManager(&_ContractEigenPodManager.CallOpts)
}

// DelegationManager is a free data retrieval call binding the contract method 0xea4d3c9b.
//
// Solidity: function delegationManager() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) DelegationManager() (common.Address, error) {
	return _ContractEigenPodManager.Contract.DelegationManager(&_ContractEigenPodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) EigenPodBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "eigenPodBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) EigenPodBeacon() (common.Address, error) {
	return _ContractEigenPodManager.Contract.EigenPodBeacon(&_ContractEigenPodManager.CallOpts)
}

// EigenPodBeacon is a free data retrieval call binding the contract method 0x292b7b2b.
//
// Solidity: function eigenPodBeacon() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) EigenPodBeacon() (common.Address, error) {
	return _ContractEigenPodManager.Contract.EigenPodBeacon(&_ContractEigenPodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) EthPOS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "ethPOS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) EthPOS() (common.Address, error) {
	return _ContractEigenPodManager.Contract.EthPOS(&_ContractEigenPodManager.CallOpts)
}

// EthPOS is a free data retrieval call binding the contract method 0x74cdd798.
//
// Solidity: function ethPOS() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) EthPOS() (common.Address, error) {
	return _ContractEigenPodManager.Contract.EthPOS(&_ContractEigenPodManager.CallOpts)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) GetPod(opts *bind.CallOpts, podOwner common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "getPod", podOwner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _ContractEigenPodManager.Contract.GetPod(&_ContractEigenPodManager.CallOpts, podOwner)
}

// GetPod is a free data retrieval call binding the contract method 0xa38406a3.
//
// Solidity: function getPod(address podOwner) view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) GetPod(podOwner common.Address) (common.Address, error) {
	return _ContractEigenPodManager.Contract.GetPod(&_ContractEigenPodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) HasPod(opts *bind.CallOpts, podOwner common.Address) (bool, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "hasPod", podOwner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) HasPod(podOwner common.Address) (bool, error) {
	return _ContractEigenPodManager.Contract.HasPod(&_ContractEigenPodManager.CallOpts, podOwner)
}

// HasPod is a free data retrieval call binding the contract method 0xf6848d24.
//
// Solidity: function hasPod(address podOwner) view returns(bool)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) HasPod(podOwner common.Address) (bool, error) {
	return _ContractEigenPodManager.Contract.HasPod(&_ContractEigenPodManager.CallOpts, podOwner)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) NumPods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "numPods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) NumPods() (*big.Int, error) {
	return _ContractEigenPodManager.Contract.NumPods(&_ContractEigenPodManager.CallOpts)
}

// NumPods is a free data retrieval call binding the contract method 0xa6a509be.
//
// Solidity: function numPods() view returns(uint256)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) NumPods() (*big.Int, error) {
	return _ContractEigenPodManager.Contract.NumPods(&_ContractEigenPodManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Owner() (common.Address, error) {
	return _ContractEigenPodManager.Contract.Owner(&_ContractEigenPodManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) Owner() (common.Address, error) {
	return _ContractEigenPodManager.Contract.Owner(&_ContractEigenPodManager.CallOpts)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address ) view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) OwnerToPod(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "ownerToPod", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address ) view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) OwnerToPod(arg0 common.Address) (common.Address, error) {
	return _ContractEigenPodManager.Contract.OwnerToPod(&_ContractEigenPodManager.CallOpts, arg0)
}

// OwnerToPod is a free data retrieval call binding the contract method 0x9ba06275.
//
// Solidity: function ownerToPod(address ) view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) OwnerToPod(arg0 common.Address) (common.Address, error) {
	return _ContractEigenPodManager.Contract.OwnerToPod(&_ContractEigenPodManager.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Paused(index uint8) (bool, error) {
	return _ContractEigenPodManager.Contract.Paused(&_ContractEigenPodManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractEigenPodManager.Contract.Paused(&_ContractEigenPodManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Paused0() (*big.Int, error) {
	return _ContractEigenPodManager.Contract.Paused0(&_ContractEigenPodManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractEigenPodManager.Contract.Paused0(&_ContractEigenPodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractEigenPodManager.Contract.PauserRegistry(&_ContractEigenPodManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractEigenPodManager.Contract.PauserRegistry(&_ContractEigenPodManager.CallOpts)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address ) view returns(int256)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) PodOwnerShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "podOwnerShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address ) view returns(int256)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) PodOwnerShares(arg0 common.Address) (*big.Int, error) {
	return _ContractEigenPodManager.Contract.PodOwnerShares(&_ContractEigenPodManager.CallOpts, arg0)
}

// PodOwnerShares is a free data retrieval call binding the contract method 0x60f4062b.
//
// Solidity: function podOwnerShares(address ) view returns(int256)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) PodOwnerShares(arg0 common.Address) (*big.Int, error) {
	return _ContractEigenPodManager.Contract.PodOwnerShares(&_ContractEigenPodManager.CallOpts, arg0)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) StakerStrategyShares(opts *bind.CallOpts, user common.Address, strategy common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "stakerStrategyShares", user, strategy)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractEigenPodManager.Contract.StakerStrategyShares(&_ContractEigenPodManager.CallOpts, user, strategy)
}

// StakerStrategyShares is a free data retrieval call binding the contract method 0x7a7e0d92.
//
// Solidity: function stakerStrategyShares(address user, address strategy) view returns(uint256 shares)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) StakerStrategyShares(user common.Address, strategy common.Address) (*big.Int, error) {
	return _ContractEigenPodManager.Contract.StakerStrategyShares(&_ContractEigenPodManager.CallOpts, user, strategy)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractEigenPodManager.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) StrategyManager() (common.Address, error) {
	return _ContractEigenPodManager.Contract.StrategyManager(&_ContractEigenPodManager.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerCallerSession) StrategyManager() (common.Address, error) {
	return _ContractEigenPodManager.Contract.StrategyManager(&_ContractEigenPodManager.CallOpts)
}

// AddOwnedShares is a paid mutator transaction binding the contract method 0xa9095e94.
//
// Solidity: function addOwnedShares(address staker, address strategy, address , uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) AddOwnedShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "addOwnedShares", staker, strategy, arg2, shares)
}

// AddOwnedShares is a paid mutator transaction binding the contract method 0xa9095e94.
//
// Solidity: function addOwnedShares(address staker, address strategy, address , uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) AddOwnedShares(staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.AddOwnedShares(&_ContractEigenPodManager.TransactOpts, staker, strategy, arg2, shares)
}

// AddOwnedShares is a paid mutator transaction binding the contract method 0xa9095e94.
//
// Solidity: function addOwnedShares(address staker, address strategy, address , uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) AddOwnedShares(staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.AddOwnedShares(&_ContractEigenPodManager.TransactOpts, staker, strategy, arg2, shares)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) CreatePod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "createPod")
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerSession) CreatePod() (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.CreatePod(&_ContractEigenPodManager.TransactOpts)
}

// CreatePod is a paid mutator transaction binding the contract method 0x84d81062.
//
// Solidity: function createPod() returns(address)
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) CreatePod() (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.CreatePod(&_ContractEigenPodManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 _initPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _pauserRegistry common.Address, _initPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "initialize", initialOwner, _pauserRegistry, _initPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 _initPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, _initPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Initialize(&_ContractEigenPodManager.TransactOpts, initialOwner, _pauserRegistry, _initPausedStatus)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address initialOwner, address _pauserRegistry, uint256 _initPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) Initialize(initialOwner common.Address, _pauserRegistry common.Address, _initPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Initialize(&_ContractEigenPodManager.TransactOpts, initialOwner, _pauserRegistry, _initPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Pause(&_ContractEigenPodManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Pause(&_ContractEigenPodManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.PauseAll(&_ContractEigenPodManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.PauseAll(&_ContractEigenPodManager.TransactOpts)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0x095e210c.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta, uint64 proportionOfOldBalance) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) RecordBeaconChainETHBalanceUpdate(opts *bind.TransactOpts, podOwner common.Address, sharesDelta *big.Int, proportionOfOldBalance uint64) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "recordBeaconChainETHBalanceUpdate", podOwner, sharesDelta, proportionOfOldBalance)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0x095e210c.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta, uint64 proportionOfOldBalance) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int, proportionOfOldBalance uint64) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_ContractEigenPodManager.TransactOpts, podOwner, sharesDelta, proportionOfOldBalance)
}

// RecordBeaconChainETHBalanceUpdate is a paid mutator transaction binding the contract method 0x095e210c.
//
// Solidity: function recordBeaconChainETHBalanceUpdate(address podOwner, int256 sharesDelta, uint64 proportionOfOldBalance) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) RecordBeaconChainETHBalanceUpdate(podOwner common.Address, sharesDelta *big.Int, proportionOfOldBalance uint64) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.RecordBeaconChainETHBalanceUpdate(&_ContractEigenPodManager.TransactOpts, podOwner, sharesDelta, proportionOfOldBalance)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) RemoveShares(opts *bind.TransactOpts, staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "removeShares", staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.RemoveShares(&_ContractEigenPodManager.TransactOpts, staker, strategy, shares)
}

// RemoveShares is a paid mutator transaction binding the contract method 0x8c80d4e5.
//
// Solidity: function removeShares(address staker, address strategy, uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) RemoveShares(staker common.Address, strategy common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.RemoveShares(&_ContractEigenPodManager.TransactOpts, staker, strategy, shares)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.RenounceOwnership(&_ContractEigenPodManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.RenounceOwnership(&_ContractEigenPodManager.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.SetPauserRegistry(&_ContractEigenPodManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.SetPauserRegistry(&_ContractEigenPodManager.TransactOpts, newPauserRegistry)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) Stake(opts *bind.TransactOpts, pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "stake", pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Stake(&_ContractEigenPodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes pubkey, bytes signature, bytes32 depositDataRoot) payable returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) Stake(pubkey []byte, signature []byte, depositDataRoot [32]byte) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Stake(&_ContractEigenPodManager.TransactOpts, pubkey, signature, depositDataRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.TransferOwnership(&_ContractEigenPodManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.TransferOwnership(&_ContractEigenPodManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Unpause(&_ContractEigenPodManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.Unpause(&_ContractEigenPodManager.TransactOpts, newPausedStatus)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address , uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactor) WithdrawSharesAsTokens(opts *bind.TransactOpts, staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.contract.Transact(opts, "withdrawSharesAsTokens", staker, strategy, arg2, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address , uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.WithdrawSharesAsTokens(&_ContractEigenPodManager.TransactOpts, staker, strategy, arg2, shares)
}

// WithdrawSharesAsTokens is a paid mutator transaction binding the contract method 0x2eae418c.
//
// Solidity: function withdrawSharesAsTokens(address staker, address strategy, address , uint256 shares) returns()
func (_ContractEigenPodManager *ContractEigenPodManagerTransactorSession) WithdrawSharesAsTokens(staker common.Address, strategy common.Address, arg2 common.Address, shares *big.Int) (*types.Transaction, error) {
	return _ContractEigenPodManager.Contract.WithdrawSharesAsTokens(&_ContractEigenPodManager.TransactOpts, staker, strategy, arg2, shares)
}

// ContractEigenPodManagerBeaconChainETHDepositedIterator is returned from FilterBeaconChainETHDeposited and is used to iterate over the raw logs and unpacked data for BeaconChainETHDeposited events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerBeaconChainETHDepositedIterator struct {
	Event *ContractEigenPodManagerBeaconChainETHDeposited // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerBeaconChainETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerBeaconChainETHDeposited)
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
		it.Event = new(ContractEigenPodManagerBeaconChainETHDeposited)
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
func (it *ContractEigenPodManagerBeaconChainETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerBeaconChainETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerBeaconChainETHDeposited represents a BeaconChainETHDeposited event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerBeaconChainETHDeposited struct {
	PodOwner common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeaconChainETHDeposited is a free log retrieval operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterBeaconChainETHDeposited(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerBeaconChainETHDepositedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerBeaconChainETHDepositedIterator{contract: _ContractEigenPodManager.contract, event: "BeaconChainETHDeposited", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHDeposited is a free log subscription operation binding the contract event 0x35a85cabc603f48abb2b71d9fbd8adea7c449d7f0be900ae7a2986ea369c3d0d.
//
// Solidity: event BeaconChainETHDeposited(address indexed podOwner, uint256 amount)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchBeaconChainETHDeposited(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerBeaconChainETHDeposited, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "BeaconChainETHDeposited", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerBeaconChainETHDeposited)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParseBeaconChainETHDeposited(log types.Log) (*ContractEigenPodManagerBeaconChainETHDeposited, error) {
	event := new(ContractEigenPodManagerBeaconChainETHDeposited)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "BeaconChainETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator is returned from FilterBeaconChainETHWithdrawalCompleted and is used to iterate over the raw logs and unpacked data for BeaconChainETHWithdrawalCompleted events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator struct {
	Event *ContractEigenPodManagerBeaconChainETHWithdrawalCompleted // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerBeaconChainETHWithdrawalCompleted)
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
		it.Event = new(ContractEigenPodManagerBeaconChainETHWithdrawalCompleted)
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
func (it *ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerBeaconChainETHWithdrawalCompleted represents a BeaconChainETHWithdrawalCompleted event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerBeaconChainETHWithdrawalCompleted struct {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterBeaconChainETHWithdrawalCompleted(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerBeaconChainETHWithdrawalCompletedIterator{contract: _ContractEigenPodManager.contract, event: "BeaconChainETHWithdrawalCompleted", logs: logs, sub: sub}, nil
}

// WatchBeaconChainETHWithdrawalCompleted is a free log subscription operation binding the contract event 0xa6bab1d55a361fcea2eee2bc9491e4f01e6cf333df03c9c4f2c144466429f7d6.
//
// Solidity: event BeaconChainETHWithdrawalCompleted(address indexed podOwner, uint256 shares, uint96 nonce, address delegatedAddress, address withdrawer, bytes32 withdrawalRoot)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchBeaconChainETHWithdrawalCompleted(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerBeaconChainETHWithdrawalCompleted, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "BeaconChainETHWithdrawalCompleted", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerBeaconChainETHWithdrawalCompleted)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParseBeaconChainETHWithdrawalCompleted(log types.Log) (*ContractEigenPodManagerBeaconChainETHWithdrawalCompleted, error) {
	event := new(ContractEigenPodManagerBeaconChainETHWithdrawalCompleted)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "BeaconChainETHWithdrawalCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerInitializedIterator struct {
	Event *ContractEigenPodManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerInitialized)
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
		it.Event = new(ContractEigenPodManagerInitialized)
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
func (it *ContractEigenPodManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerInitialized represents a Initialized event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractEigenPodManagerInitializedIterator, error) {

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerInitializedIterator{contract: _ContractEigenPodManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerInitialized)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParseInitialized(log types.Log) (*ContractEigenPodManagerInitialized, error) {
	event := new(ContractEigenPodManagerInitialized)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerNewTotalSharesIterator is returned from FilterNewTotalShares and is used to iterate over the raw logs and unpacked data for NewTotalShares events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerNewTotalSharesIterator struct {
	Event *ContractEigenPodManagerNewTotalShares // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerNewTotalSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerNewTotalShares)
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
		it.Event = new(ContractEigenPodManagerNewTotalShares)
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
func (it *ContractEigenPodManagerNewTotalSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerNewTotalSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerNewTotalShares represents a NewTotalShares event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerNewTotalShares struct {
	PodOwner       common.Address
	NewTotalShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewTotalShares is a free log retrieval operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterNewTotalShares(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerNewTotalSharesIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerNewTotalSharesIterator{contract: _ContractEigenPodManager.contract, event: "NewTotalShares", logs: logs, sub: sub}, nil
}

// WatchNewTotalShares is a free log subscription operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchNewTotalShares(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerNewTotalShares, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "NewTotalShares", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerNewTotalShares)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
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

// ParseNewTotalShares is a log parse operation binding the contract event 0xd4def76d6d2bed6f14d5cd9af73cc2913d618d00edde42432e81c09bfe077098.
//
// Solidity: event NewTotalShares(address indexed podOwner, int256 newTotalShares)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParseNewTotalShares(log types.Log) (*ContractEigenPodManagerNewTotalShares, error) {
	event := new(ContractEigenPodManagerNewTotalShares)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "NewTotalShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerOwnershipTransferredIterator struct {
	Event *ContractEigenPodManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerOwnershipTransferred)
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
		it.Event = new(ContractEigenPodManagerOwnershipTransferred)
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
func (it *ContractEigenPodManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractEigenPodManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerOwnershipTransferredIterator{contract: _ContractEigenPodManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerOwnershipTransferred)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractEigenPodManagerOwnershipTransferred, error) {
	event := new(ContractEigenPodManagerOwnershipTransferred)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPausedIterator struct {
	Event *ContractEigenPodManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerPaused)
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
		it.Event = new(ContractEigenPodManagerPaused)
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
func (it *ContractEigenPodManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerPaused represents a Paused event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractEigenPodManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerPausedIterator{contract: _ContractEigenPodManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerPaused)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParsePaused(log types.Log) (*ContractEigenPodManagerPaused, error) {
	event := new(ContractEigenPodManagerPaused)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPauserRegistrySetIterator struct {
	Event *ContractEigenPodManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerPauserRegistrySet)
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
		it.Event = new(ContractEigenPodManagerPauserRegistrySet)
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
func (it *ContractEigenPodManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractEigenPodManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerPauserRegistrySetIterator{contract: _ContractEigenPodManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerPauserRegistrySet)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractEigenPodManagerPauserRegistrySet, error) {
	event := new(ContractEigenPodManagerPauserRegistrySet)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerPodDeployedIterator is returned from FilterPodDeployed and is used to iterate over the raw logs and unpacked data for PodDeployed events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPodDeployedIterator struct {
	Event *ContractEigenPodManagerPodDeployed // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerPodDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerPodDeployed)
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
		it.Event = new(ContractEigenPodManagerPodDeployed)
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
func (it *ContractEigenPodManagerPodDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerPodDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerPodDeployed represents a PodDeployed event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPodDeployed struct {
	EigenPod common.Address
	PodOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPodDeployed is a free log retrieval operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterPodDeployed(opts *bind.FilterOpts, eigenPod []common.Address, podOwner []common.Address) (*ContractEigenPodManagerPodDeployedIterator, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerPodDeployedIterator{contract: _ContractEigenPodManager.contract, event: "PodDeployed", logs: logs, sub: sub}, nil
}

// WatchPodDeployed is a free log subscription operation binding the contract event 0x21c99d0db02213c32fff5b05cf0a718ab5f858802b91498f80d82270289d856a.
//
// Solidity: event PodDeployed(address indexed eigenPod, address indexed podOwner)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchPodDeployed(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPodDeployed, eigenPod []common.Address, podOwner []common.Address) (event.Subscription, error) {

	var eigenPodRule []interface{}
	for _, eigenPodItem := range eigenPod {
		eigenPodRule = append(eigenPodRule, eigenPodItem)
	}
	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "PodDeployed", eigenPodRule, podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerPodDeployed)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParsePodDeployed(log types.Log) (*ContractEigenPodManagerPodDeployed, error) {
	event := new(ContractEigenPodManagerPodDeployed)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "PodDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerPodSharesUpdatedIterator is returned from FilterPodSharesUpdated and is used to iterate over the raw logs and unpacked data for PodSharesUpdated events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPodSharesUpdatedIterator struct {
	Event *ContractEigenPodManagerPodSharesUpdated // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerPodSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerPodSharesUpdated)
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
		it.Event = new(ContractEigenPodManagerPodSharesUpdated)
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
func (it *ContractEigenPodManagerPodSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerPodSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerPodSharesUpdated represents a PodSharesUpdated event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerPodSharesUpdated struct {
	PodOwner    common.Address
	SharesDelta *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPodSharesUpdated is a free log retrieval operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterPodSharesUpdated(opts *bind.FilterOpts, podOwner []common.Address) (*ContractEigenPodManagerPodSharesUpdatedIterator, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerPodSharesUpdatedIterator{contract: _ContractEigenPodManager.contract, event: "PodSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPodSharesUpdated is a free log subscription operation binding the contract event 0x4e2b791dedccd9fb30141b088cabf5c14a8912b52f59375c95c010700b8c6193.
//
// Solidity: event PodSharesUpdated(address indexed podOwner, int256 sharesDelta)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchPodSharesUpdated(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerPodSharesUpdated, podOwner []common.Address) (event.Subscription, error) {

	var podOwnerRule []interface{}
	for _, podOwnerItem := range podOwner {
		podOwnerRule = append(podOwnerRule, podOwnerItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "PodSharesUpdated", podOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerPodSharesUpdated)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParsePodSharesUpdated(log types.Log) (*ContractEigenPodManagerPodSharesUpdated, error) {
	event := new(ContractEigenPodManagerPodSharesUpdated)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "PodSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEigenPodManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerUnpausedIterator struct {
	Event *ContractEigenPodManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractEigenPodManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEigenPodManagerUnpaused)
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
		it.Event = new(ContractEigenPodManagerUnpaused)
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
func (it *ContractEigenPodManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEigenPodManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEigenPodManagerUnpaused represents a Unpaused event raised by the ContractEigenPodManager contract.
type ContractEigenPodManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractEigenPodManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractEigenPodManagerUnpausedIterator{contract: _ContractEigenPodManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractEigenPodManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractEigenPodManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEigenPodManagerUnpaused)
				if err := _ContractEigenPodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractEigenPodManager *ContractEigenPodManagerFilterer) ParseUnpaused(log types.Log) (*ContractEigenPodManagerUnpaused, error) {
	event := new(ContractEigenPodManagerUnpaused)
	if err := _ContractEigenPodManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
