// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractBLSApkRegistry

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

// IBLSApkRegistryTypesApkUpdate is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesApkUpdate struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}

// IBLSApkRegistryTypesPubkeyRegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryTypesPubkeyRegistrationParams struct {
	PubkeyRegistrationSignature BN254G1Point
	PubkeyG1                    BN254G1Point
	PubkeyG2                    BN254G2Point
}

// ContractBLSApkRegistryMetaData contains all meta data concerning the ContractBLSApkRegistry contract.
var ContractBLSApkRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_slashingRegistryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"apkHistory\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApk\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHashAtBlockNumberAndIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkHistoryLength\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkIndicesAtBlockNumber\",\"inputs\":[{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApkUpdateAtIndex\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.ApkUpdate\",\"components\":[{\"name\":\"apkHash\",\"type\":\"bytes24\",\"internalType\":\"bytes24\"},{\"name\":\"updateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"nextUpdateBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorFromPubkeyHash\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorId\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorPubkeyG2\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRegisteredPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeQuorum\",\"inputs\":[{\"name\":\"quorumNumber\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorToPubkey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorToPubkeyHash\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pubkeyHashToOperator\",\"inputs\":[{\"name\":\"pubkeyHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistryTypes.PubkeyRegistrationParams\",\"components\":[{\"name\":\"pubkeyRegistrationSignature\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}]},{\"name\":\"pubkeyRegistrationMessageHash\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryCoordinator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyAndRegisterG2PubkeyForOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewG2PubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewPubkeyRegistration\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkeyG1\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"pubkeyG2\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorAddedToQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRemovedFromQuorums\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BLSPubkeyAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberBeforeFirstUpdate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberNotLatest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BlockNumberTooRecent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECAddFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECMulFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECPairingFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"G2PubkeyAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBLSSignatureOrPrivateKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyRegistryCoordinatorOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"QuorumDoesNotExist\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroPubKey\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f5ffd5b50604051611e9e380380611e9e83398101604081905261002e91610107565b6001600160a01b0381166080528061004461004b565b5050610134565b5f54610100900460ff16156100b65760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610105575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b5f60208284031215610117575f5ffd5b81516001600160a01b038116811461012d575f5ffd5b9392505050565b608051611d4461015a5f395f818161033901528181610e5e01526114620152611d445ff3fe608060405234801561000f575f5ffd5b5060043610610126575f3560e01c80636d14a987116100a9578063d1a646501161006e578063d1a6465014610408578063d5254a8c1461041b578063de29fac01461043b578063e8bb9ae61461045a578063f4e24fe514610482575f5ffd5b80636d14a987146103345780637916cea61461035b5780637ff81a871461039c578063a3db80e2146103cf578063bf79ce58146103f5575f5ffd5b806347b314e8116100ef57806347b314e8146101ff5780635f61a8841461023f578063605747d51461029957806367169911146102e757806368bccaac14610307575f5ffd5b8062a1f4cb1461012a57806313542a4e1461016a57806326d941f2146101a0578063377ed99d146101b55780633fb27952146101ec575b5f5ffd5b610150610138366004611701565b60036020525f90815260409020805460019091015482565b604080519283526020830191909152015b60405180910390f35b610192610178366004611701565b6001600160a01b03165f9081526001602052604090205490565b604051908152602001610161565b6101b36101ae36600461172c565b610495565b005b6101d76101c336600461172c565b60ff165f9081526004602052604090205490565b60405163ffffffff9091168152602001610161565b6101b36101fa3660046117b3565b610555565b61022761020d36600461185a565b5f908152600260205260409020546001600160a01b031690565b6040516001600160a01b039091168152602001610161565b61028c61024d36600461172c565b604080518082019091525f80825260208201525060ff165f90815260056020908152604091829020825180840190935280548352600101549082015290565b6040516101619190611871565b6102ac6102a7366004611888565b6105d1565b60408051825167ffffffffffffffff1916815260208084015163ffffffff908116918301919091529282015190921690820152606001610161565b6102fa6102f5366004611701565b610662565b60405161016191906118d2565b61031a6103153660046118fd565b6106f6565b60405167ffffffffffffffff199091168152602001610161565b6102277f000000000000000000000000000000000000000000000000000000000000000081565b61036e610369366004611888565b6107db565b6040805167ffffffffffffffff19909416845263ffffffff9283166020850152911690820152606001610161565b6103af6103aa366004611701565b610822565b604080518351815260209384015193810193909352820152606001610161565b6101506103dd36600461172c565b60056020525f90815260409020805460019091015482565b610192610403366004611941565b610898565b6101b361041636600461199b565b610b6d565b61042e6104293660046119d9565b610c54565b6040516101619190611a4b565b610192610449366004611701565b60016020525f908152604090205481565b61022761046836600461185a565b60026020525f90815260409020546001600160a01b031681565b6101b36104903660046117b3565b610dec565b61049d610e53565b60ff81165f90815260046020526040902054156104cd576040516310cda51760e21b815260040160405180910390fd5b60ff165f908152600460209081526040808320815160608101835284815263ffffffff4381168286019081528285018781528454600181018655948852959096209151919092018054955194518316600160e01b026001600160e01b0395909316600160c01b026001600160e01b03199096169190931c179390931791909116919091179055565b61055d610e53565b5f61056783610822565b5090506105748282610e9e565b7f73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e836105b4856001600160a01b03165f9081526001602052604090205490565b846040516105c493929190611a93565b60405180910390a1505050565b604080516060810182525f808252602080830182905282840182905260ff86168252600490529190912080548390811061060d5761060d611ade565b5f91825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b8204811694830194909452600160e01b90049092169082015290505b92915050565b61066a61162e565b6001600160a01b0382165f9081526006602052604090819020815160808101835291829081018260028282826020028201915b81548152602001906001019080831161069d57505050918352505060408051808201918290526020909201919060028481019182845b8154815260200190600101908083116106d3575050505050815250509050919050565b60ff83165f90815260046020526040812080548291908490811061071c5761071c611ade565b5f91825260209182902060408051606081018252919092015467ffffffffffffffff1981841b16825263ffffffff600160c01b82048116948301859052600160e01b90910481169282019290925292508516101561078d57604051633d22884160e01b815260040160405180910390fd5b604081015163ffffffff1615806107b35750806040015163ffffffff168463ffffffff16105b6107d057604051636fe02d4b60e01b815260040160405180910390fd5b5190505b9392505050565b6004602052815f5260405f2081815481106107f4575f80fd5b5f91825260209091200154604081901b925063ffffffff600160c01b820481169250600160e01b9091041683565b604080518082019091525f80825260208201526001600160a01b0382165f8181526003602090815260408083208151808301835281548152600191820154818501529484529091528120549091908061088e576040516325ec6c1f60e01b815260040160405180910390fd5b9094909350915050565b5f6108a1610e53565b5f6108cd6108b736869003860160408701611af2565b80515f9081526020918201519091526040902090565b90507fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5810361090f57604051630cc7509160e01b815260040160405180910390fd5b6001600160a01b0385165f9081526001602052604081205414610945576040516342ee68b560e01b815260040160405180910390fd5b5f818152600260205260409020546001600160a01b03161561097a57604051634c334c9760e11b815260040160405180910390fd5b604080515f917f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001916109d2918835916020808b0135928b01359160608c01359160808d019160c08e01918d35918e8201359101611b23565b604051602081830303815290604052805190602001205f1c6109f49190611b65565b9050610a8d610a2d610a1883610a12368a90038a0160408b01611af2565b90611085565b610a2736899003890189611af2565b906110f5565b610a35611169565b610a76610a6785610a126040805180820182525f80825260209182015281518083019092526001825260029082015290565b610a27368a90038a018a611af2565b610a88368a90038a0160808b01611bc6565b611229565b610aaa5760405163a72d026360e01b815260040160405180910390fd5b6001600160a01b0386165f9081526003602090815260408083208882013581556060890135600190910155600690915290206080860190610aeb8282611c30565b50506001600160a01b0386165f81815260016020908152604080832086905585835260029091529081902080546001600160a01b0319168317905580517fe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba382804191610b5c919089019060808a0190611c8b565b60405180910390a250949350505050565b610b75611460565b5f610b7f83610822565b509050610b8b83611511565b610bcb81610b97611169565b6040805180820182525f808252602091820152815180830190925260018252600290820152610a8836879003870187611bc6565b610be85760405163a72d026360e01b815260040160405180910390fd5b6001600160a01b0383165f9081526006602052604090208290610c0b8282611c30565b905050826001600160a01b03167f5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a83604051610c479190611caa565b60405180910390a2505050565b60605f8367ffffffffffffffff811115610c7057610c70611745565b604051908082528060200260200182016040528015610c99578160200160208202803683370190505b5090505f5b84811015610de3575f868683818110610cb957610cb9611ade565b919091013560f81c5f818152600460205260409020549092509050801580610d19575060ff82165f9081526004602052604081208054909190610cfe57610cfe611ade565b5f91825260209091200154600160c01b900463ffffffff1686105b15610d3757604051633f4cb70f60e01b815260040160405180910390fd5b805b8015610dd85760ff83165f9081526004602052604090208790610d5d600184611cb8565b81548110610d6d57610d6d611ade565b5f91825260209091200154600160c01b900463ffffffff1611610dc657610d95600182611cb8565b858581518110610da757610da7611ade565b602002602001019063ffffffff16908163ffffffff1681525050610dd8565b80610dd081611ccb565b915050610d39565b505050600101610c9e565b50949350505050565b610df4610e53565b5f610dfe83610822565b509050610e1382610e0e83611572565b610e9e565b7ff843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e836105b4856001600160a01b03165f9081526001602052604090205490565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e9c57604051637070f3b160e11b815260040160405180910390fd5b565b604080518082019091525f80825260208201525f5b835181101561107f575f848281518110610ecf57610ecf611ade565b0160209081015160f81c5f8181526004909252604082205490925090819003610f0b57604051637310cff560e11b815260040160405180910390fd5b60ff82165f908152600560209081526040918290208251808401909352805483526001015490820152610f3e90866110f5565b60ff83165f818152600560209081526040808320855180825586840180516001938401559085525183528184209484526004909252822093975091929091610f869085611cb8565b81548110610f9657610f96611ade565b5f918252602090912001805490915063ffffffff438116600160c01b9092041603610fd45780546001600160c01b031916604083901c17815561106f565b805463ffffffff438116600160e01b8181026001600160e01b0394851617855560ff88165f908152600460209081526040808320815160608101835267ffffffffffffffff198b16815280840196875280830185815282546001810184559286529390942093519301805495519251871690940291909516600160c01b026001600160e01b0319949094169190941c17919091179092161790555b505060019092019150610eb39050565b50505050565b604080518082019091525f80825260208201526110a0611653565b835181526020808501519082015260408082018490525f908360608460076107d05a03fa905080806110ce57fe5b50806110ed57604051632319df1960e11b815260040160405180910390fd5b505092915050565b604080518082019091525f8082526020820152611110611671565b835181526020808501518183015283516040808401919091529084015160608301525f908360808460066107d05a03fa9050808061114a57fe5b50806110ed5760405163d4b68fd760e01b815260040160405180910390fd5b61117161162e565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b6040805180820182528581526020808201859052825180840190935285835282018390525f9161125761168f565b5f5b600281101561140e575f61126e826006611c19565b905084826002811061128257611282611ade565b60200201515183611293835f611ce0565b600c81106112a3576112a3611ade565b60200201528482600281106112ba576112ba611ade565b602002015160200151838260016112d19190611ce0565b600c81106112e1576112e1611ade565b60200201528382600281106112f8576112f8611ade565b602002015151518361130b836002611ce0565b600c811061131b5761131b611ade565b602002015283826002811061133257611332611ade565b602002015151600160200201518361134b836003611ce0565b600c811061135b5761135b611ade565b602002015283826002811061137257611372611ade565b6020020151602001515f6002811061138c5761138c611ade565b60200201518361139d836004611ce0565b600c81106113ad576113ad611ade565b60200201528382600281106113c4576113c4611ade565b6020020151602001516001600281106113df576113df611ade565b6020020151836113f0836005611ce0565b600c811061140057611400611ade565b602002015250600101611259565b506114176116ae565b5f6020826101808560086107d05a03fa9050808061143157fe5b5080611450576040516324ccc79360e21b815260040160405180910390fd5b5051151598975050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114bc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114e09190611cf3565b6001600160a01b0316336001600160a01b031614610e9c57604051637070f3b160e11b815260040160405180910390fd5b5f61151b82610662565b8051519091501580156115315750805160200151155b80156115405750602081015151155b801561155157506020818101510151155b61156e57604051630849e5cf60e41b815260040160405180910390fd5b5050565b604080518082019091525f8082526020820152815115801561159657506020820151155b156115b3575050604080518082019091525f808252602082015290565b6040518060400160405280835f015181526020017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4784602001516115f79190611b65565b611621907f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47611cb8565b905292915050565b919050565b60405180604001604052806116416116cc565b815260200161164e6116cc565b905290565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b03811681146116fe575f5ffd5b50565b5f60208284031215611711575f5ffd5b81356107d4816116ea565b803560ff81168114611629575f5ffd5b5f6020828403121561173c575f5ffd5b6107d48261171c565b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff8111828210171561177c5761177c611745565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156117ab576117ab611745565b604052919050565b5f5f604083850312156117c4575f5ffd5b82356117cf816116ea565b9150602083013567ffffffffffffffff8111156117ea575f5ffd5b8301601f810185136117fa575f5ffd5b803567ffffffffffffffff81111561181457611814611745565b611827601f8201601f1916602001611782565b81815286602083850101111561183b575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b5f6020828403121561186a575f5ffd5b5035919050565b81518152602080830151908201526040810161065c565b5f5f60408385031215611899575f5ffd5b6118a28361171c565b946020939093013593505050565b805f5b600281101561107f5781518452602093840193909101906001016118b3565b5f6080820190506118e48284516118b0565b60208301516118f660408401826118b0565b5092915050565b5f5f5f6060848603121561190f575f5ffd5b6119188461171c565b9250602084013563ffffffff81168114611930575f5ffd5b929592945050506040919091013590565b5f5f5f838503610160811215611955575f5ffd5b8435611960816116ea565b9350610100601f1982011215611974575f5ffd5b602085019250604061011f198201121561198c575f5ffd5b50610120840190509250925092565b5f5f82840360a08112156119ad575f5ffd5b83356119b8816116ea565b92506080601f19820112156119cb575f5ffd5b506020830190509250929050565b5f5f5f604084860312156119eb575f5ffd5b833567ffffffffffffffff811115611a01575f5ffd5b8401601f81018613611a11575f5ffd5b803567ffffffffffffffff811115611a27575f5ffd5b866020828401011115611a38575f5ffd5b6020918201979096509401359392505050565b602080825282518282018190525f918401906040840190835b81811015611a8857835163ffffffff16835260209384019390920191600101611a64565b509095945050505050565b60018060a01b0384168152826020820152606060408201525f82518060608401528060208501608085015e5f608082850101526080601f19601f830116840101915050949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f6040828403128015611b03575f5ffd5b50611b0c611759565b823581526020928301359281019290925250919050565b888152876020820152866040820152856060820152604085608083013760408460c0830137610100810192909252610120820152610140019695505050505050565b5f82611b7f57634e487b7160e01b5f52601260045260245ffd5b500690565b5f82601f830112611b93575f5ffd5b611b9b611759565b806040840185811115611bac575f5ffd5b845b81811015611a88578035845260209384019301611bae565b5f6080828403128015611bd7575f5ffd5b50611be0611759565b611bea8484611b84565b8152611bf98460408501611b84565b60208201529392505050565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761065c5761065c611c05565b815f5b6002811015611c5057813583820155602090910190600101611c33565b5050604082015f5b600281101561107f57813583820160020155602090910190600101611c58565b6040818337604080820160408401375050565b823581526020808401359082015260c081016107d46040830184611c78565b6080810161065c8284611c78565b8181038181111561065c5761065c611c05565b5f81611cd957611cd9611c05565b505f190190565b8082018082111561065c5761065c611c05565b5f60208284031215611d03575f5ffd5b81516107d4816116ea56fea2646970667358221220fda4083afac56ebcb94a7445aa6daed5ae10bd623695be7cf879026018473c7164736f6c634300081b0033",
}

// ContractBLSApkRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractBLSApkRegistryMetaData.ABI instead.
var ContractBLSApkRegistryABI = ContractBLSApkRegistryMetaData.ABI

// ContractBLSApkRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractBLSApkRegistryMetaData.Bin instead.
var ContractBLSApkRegistryBin = ContractBLSApkRegistryMetaData.Bin

// DeployContractBLSApkRegistry deploys a new Ethereum contract, binding an instance of ContractBLSApkRegistry to it.
func DeployContractBLSApkRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _slashingRegistryCoordinator common.Address) (common.Address, *types.Transaction, *ContractBLSApkRegistry, error) {
	parsed, err := ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBLSApkRegistryBin), backend, _slashingRegistryCoordinator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractBLSApkRegistry{ContractBLSApkRegistryCaller: ContractBLSApkRegistryCaller{contract: contract}, ContractBLSApkRegistryTransactor: ContractBLSApkRegistryTransactor{contract: contract}, ContractBLSApkRegistryFilterer: ContractBLSApkRegistryFilterer{contract: contract}}, nil
}

// ContractBLSApkRegistryMethods is an auto generated interface around an Ethereum contract.
type ContractBLSApkRegistryMethods interface {
	ContractBLSApkRegistryCalls
	ContractBLSApkRegistryTransacts
	ContractBLSApkRegistryFilters
}

// ContractBLSApkRegistryCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractBLSApkRegistryCalls interface {
	ApkHistory(opts *bind.CallOpts, quorumNumber uint8, arg1 *big.Int) (struct {
		ApkHash               [24]byte
		UpdateBlockNumber     uint32
		NextUpdateBlockNumber uint32
	}, error)

	CurrentApk(opts *bind.CallOpts, quorumNumber uint8) (struct {
		X *big.Int
		Y *big.Int
	}, error)

	GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error)

	GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error)

	GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error)

	GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error)

	GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error)

	GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	GetOperatorPubkeyG2(opts *bind.CallOpts, operator common.Address) (BN254G2Point, error)

	GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error)

	OperatorToPubkey(opts *bind.CallOpts, operator common.Address) (struct {
		X *big.Int
		Y *big.Int
	}, error)

	OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error)

	PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error)

	RegistryCoordinator(opts *bind.CallOpts) (common.Address, error)
}

// ContractBLSApkRegistryTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractBLSApkRegistryTransacts interface {
	DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error)

	RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error)

	RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error)

	VerifyAndRegisterG2PubkeyForOperator(opts *bind.TransactOpts, operator common.Address, pubkeyG2 BN254G2Point) (*types.Transaction, error)
}

// ContractBLSApkRegistryFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractBLSApkRegistryFilters interface {
	FilterInitialized(opts *bind.FilterOpts) (*ContractBLSApkRegistryInitializedIterator, error)
	WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryInitialized) (event.Subscription, error)
	ParseInitialized(log types.Log) (*ContractBLSApkRegistryInitialized, error)

	FilterNewG2PubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewG2PubkeyRegistrationIterator, error)
	WatchNewG2PubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewG2PubkeyRegistration, operator []common.Address) (event.Subscription, error)
	ParseNewG2PubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewG2PubkeyRegistration, error)

	FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewPubkeyRegistrationIterator, error)
	WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error)
	ParseNewPubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewPubkeyRegistration, error)

	FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorAddedToQuorumsIterator, error)
	WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error)
	ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSApkRegistryOperatorAddedToQuorums, error)

	FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator, error)
	WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error)
	ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSApkRegistryOperatorRemovedFromQuorums, error)
}

// ContractBLSApkRegistry is an auto generated Go binding around an Ethereum contract.
type ContractBLSApkRegistry struct {
	ContractBLSApkRegistryCaller     // Read-only binding to the contract
	ContractBLSApkRegistryTransactor // Write-only binding to the contract
	ContractBLSApkRegistryFilterer   // Log filterer for contract events
}

// ContractBLSApkRegistry implements the ContractBLSApkRegistryMethods interface.
var _ ContractBLSApkRegistryMethods = (*ContractBLSApkRegistry)(nil)

// ContractBLSApkRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryCaller implements the ContractBLSApkRegistryCalls interface.
var _ ContractBLSApkRegistryCalls = (*ContractBLSApkRegistryCaller)(nil)

// ContractBLSApkRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryTransactor implements the ContractBLSApkRegistryTransacts interface.
var _ ContractBLSApkRegistryTransacts = (*ContractBLSApkRegistryTransactor)(nil)

// ContractBLSApkRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractBLSApkRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractBLSApkRegistryFilterer implements the ContractBLSApkRegistryFilters interface.
var _ ContractBLSApkRegistryFilters = (*ContractBLSApkRegistryFilterer)(nil)

// ContractBLSApkRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractBLSApkRegistrySession struct {
	Contract     *ContractBLSApkRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractBLSApkRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractBLSApkRegistryCallerSession struct {
	Contract *ContractBLSApkRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ContractBLSApkRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractBLSApkRegistryTransactorSession struct {
	Contract     *ContractBLSApkRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ContractBLSApkRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractBLSApkRegistryRaw struct {
	Contract *ContractBLSApkRegistry // Generic contract binding to access the raw methods on
}

// ContractBLSApkRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryCallerRaw struct {
	Contract *ContractBLSApkRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractBLSApkRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractBLSApkRegistryTransactorRaw struct {
	Contract *ContractBLSApkRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractBLSApkRegistry creates a new instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistry(address common.Address, backend bind.ContractBackend) (*ContractBLSApkRegistry, error) {
	contract, err := bindContractBLSApkRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistry{ContractBLSApkRegistryCaller: ContractBLSApkRegistryCaller{contract: contract}, ContractBLSApkRegistryTransactor: ContractBLSApkRegistryTransactor{contract: contract}, ContractBLSApkRegistryFilterer: ContractBLSApkRegistryFilterer{contract: contract}}, nil
}

// NewContractBLSApkRegistryCaller creates a new read-only instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractBLSApkRegistryCaller, error) {
	contract, err := bindContractBLSApkRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryCaller{contract: contract}, nil
}

// NewContractBLSApkRegistryTransactor creates a new write-only instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractBLSApkRegistryTransactor, error) {
	contract, err := bindContractBLSApkRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryTransactor{contract: contract}, nil
}

// NewContractBLSApkRegistryFilterer creates a new log filterer instance of ContractBLSApkRegistry, bound to a specific deployed contract.
func NewContractBLSApkRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractBLSApkRegistryFilterer, error) {
	contract, err := bindContractBLSApkRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryFilterer{contract: contract}, nil
}

// bindContractBLSApkRegistry binds a generic wrapper to an already deployed contract.
func bindContractBLSApkRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractBLSApkRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.ContractBLSApkRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractBLSApkRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.contract.Transact(opts, method, params...)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) ApkHistory(opts *bind.CallOpts, quorumNumber uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "apkHistory", quorumNumber, arg1)

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

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) ApkHistory(quorumNumber uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSApkRegistry.Contract.ApkHistory(&_ContractBLSApkRegistry.CallOpts, quorumNumber, arg1)
}

// ApkHistory is a free data retrieval call binding the contract method 0x7916cea6.
//
// Solidity: function apkHistory(uint8 quorumNumber, uint256 ) view returns(bytes24 apkHash, uint32 updateBlockNumber, uint32 nextUpdateBlockNumber)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) ApkHistory(quorumNumber uint8, arg1 *big.Int) (struct {
	ApkHash               [24]byte
	UpdateBlockNumber     uint32
	NextUpdateBlockNumber uint32
}, error) {
	return _ContractBLSApkRegistry.Contract.ApkHistory(&_ContractBLSApkRegistry.CallOpts, quorumNumber, arg1)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) CurrentApk(opts *bind.CallOpts, quorumNumber uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "currentApk", quorumNumber)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) CurrentApk(quorumNumber uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.CurrentApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// CurrentApk is a free data retrieval call binding the contract method 0xa3db80e2.
//
// Solidity: function currentApk(uint8 quorumNumber) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) CurrentApk(quorumNumber uint8) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.CurrentApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApk(opts *bind.CallOpts, quorumNumber uint8) (BN254G1Point, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApk", quorumNumber)

	if err != nil {
		return *new(BN254G1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)

	return out0, err

}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSApkRegistry.Contract.GetApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApk is a free data retrieval call binding the contract method 0x5f61a884.
//
// Solidity: function getApk(uint8 quorumNumber) view returns((uint256,uint256))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApk(quorumNumber uint8) (BN254G1Point, error) {
	return _ContractBLSApkRegistry.Contract.GetApk(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkHashAtBlockNumberAndIndex(opts *bind.CallOpts, quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkHashAtBlockNumberAndIndex", quorumNumber, blockNumber, index)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHashAtBlockNumberAndIndex is a free data retrieval call binding the contract method 0x68bccaac.
//
// Solidity: function getApkHashAtBlockNumberAndIndex(uint8 quorumNumber, uint32 blockNumber, uint256 index) view returns(bytes24)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkHashAtBlockNumberAndIndex(quorumNumber uint8, blockNumber uint32, index *big.Int) ([24]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHashAtBlockNumberAndIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, blockNumber, index)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkHistoryLength(opts *bind.CallOpts, quorumNumber uint8) (uint32, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkHistoryLength", quorumNumber)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHistoryLength(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkHistoryLength is a free data retrieval call binding the contract method 0x377ed99d.
//
// Solidity: function getApkHistoryLength(uint8 quorumNumber) view returns(uint32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkHistoryLength(quorumNumber uint8) (uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkHistoryLength(&_ContractBLSApkRegistry.CallOpts, quorumNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkIndicesAtBlockNumber(opts *bind.CallOpts, quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkIndicesAtBlockNumber", quorumNumbers, blockNumber)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkIndicesAtBlockNumber is a free data retrieval call binding the contract method 0xd5254a8c.
//
// Solidity: function getApkIndicesAtBlockNumber(bytes quorumNumbers, uint256 blockNumber) view returns(uint32[])
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkIndicesAtBlockNumber(quorumNumbers []byte, blockNumber *big.Int) ([]uint32, error) {
	return _ContractBLSApkRegistry.Contract.GetApkIndicesAtBlockNumber(&_ContractBLSApkRegistry.CallOpts, quorumNumbers, blockNumber)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetApkUpdateAtIndex(opts *bind.CallOpts, quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getApkUpdateAtIndex", quorumNumber, index)

	if err != nil {
		return *new(IBLSApkRegistryTypesApkUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSApkRegistryTypesApkUpdate)).(*IBLSApkRegistryTypesApkUpdate)

	return out0, err

}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	return _ContractBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetApkUpdateAtIndex is a free data retrieval call binding the contract method 0x605747d5.
//
// Solidity: function getApkUpdateAtIndex(uint8 quorumNumber, uint256 index) view returns((bytes24,uint32,uint32))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetApkUpdateAtIndex(quorumNumber uint8, index *big.Int) (IBLSApkRegistryTypesApkUpdate, error) {
	return _ContractBLSApkRegistry.Contract.GetApkUpdateAtIndex(&_ContractBLSApkRegistry.CallOpts, quorumNumber, index)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorFromPubkeyHash(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorFromPubkeyHash", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorFromPubkeyHash is a free data retrieval call binding the contract method 0x47b314e8.
//
// Solidity: function getOperatorFromPubkeyHash(bytes32 pubkeyHash) view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorFromPubkeyHash(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorFromPubkeyHash(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorId(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorId", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorId(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetOperatorId is a free data retrieval call binding the contract method 0x13542a4e.
//
// Solidity: function getOperatorId(address operator) view returns(bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorId(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorId(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetOperatorPubkeyG2 is a free data retrieval call binding the contract method 0x67169911.
//
// Solidity: function getOperatorPubkeyG2(address operator) view returns((uint256[2],uint256[2]))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetOperatorPubkeyG2(opts *bind.CallOpts, operator common.Address) (BN254G2Point, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getOperatorPubkeyG2", operator)

	if err != nil {
		return *new(BN254G2Point), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G2Point)).(*BN254G2Point)

	return out0, err

}

// GetOperatorPubkeyG2 is a free data retrieval call binding the contract method 0x67169911.
//
// Solidity: function getOperatorPubkeyG2(address operator) view returns((uint256[2],uint256[2]))
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetOperatorPubkeyG2(operator common.Address) (BN254G2Point, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorPubkeyG2(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetOperatorPubkeyG2 is a free data retrieval call binding the contract method 0x67169911.
//
// Solidity: function getOperatorPubkeyG2(address operator) view returns((uint256[2],uint256[2]))
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetOperatorPubkeyG2(operator common.Address) (BN254G2Point, error) {
	return _ContractBLSApkRegistry.Contract.GetOperatorPubkeyG2(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) GetRegisteredPubkey(opts *bind.CallOpts, operator common.Address) (BN254G1Point, [32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "getRegisteredPubkey", operator)

	if err != nil {
		return *new(BN254G1Point), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(BN254G1Point)).(*BN254G1Point)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// GetRegisteredPubkey is a free data retrieval call binding the contract method 0x7ff81a87.
//
// Solidity: function getRegisteredPubkey(address operator) view returns((uint256,uint256), bytes32)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) GetRegisteredPubkey(operator common.Address) (BN254G1Point, [32]byte, error) {
	return _ContractBLSApkRegistry.Contract.GetRegisteredPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) OperatorToPubkey(opts *bind.CallOpts, operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkey", operator)

	outstruct := new(struct {
		X *big.Int
		Y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) OperatorToPubkey(operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkey is a free data retrieval call binding the contract method 0x00a1f4cb.
//
// Solidity: function operatorToPubkey(address operator) view returns(uint256 X, uint256 Y)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) OperatorToPubkey(operator common.Address) (struct {
	X *big.Int
	Y *big.Int
}, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkey(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) OperatorToPubkeyHash(opts *bind.CallOpts, operator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "operatorToPubkeyHash", operator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractBLSApkRegistry.CallOpts, operator)
}

// OperatorToPubkeyHash is a free data retrieval call binding the contract method 0xde29fac0.
//
// Solidity: function operatorToPubkeyHash(address operator) view returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) OperatorToPubkeyHash(operator common.Address) ([32]byte, error) {
	return _ContractBLSApkRegistry.Contract.OperatorToPubkeyHash(&_ContractBLSApkRegistry.CallOpts, operator)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) PubkeyHashToOperator(opts *bind.CallOpts, pubkeyHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "pubkeyHashToOperator", pubkeyHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// PubkeyHashToOperator is a free data retrieval call binding the contract method 0xe8bb9ae6.
//
// Solidity: function pubkeyHashToOperator(bytes32 pubkeyHash) view returns(address operator)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) PubkeyHashToOperator(pubkeyHash [32]byte) (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.PubkeyHashToOperator(&_ContractBLSApkRegistry.CallOpts, pubkeyHash)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractBLSApkRegistry.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.RegistryCoordinator(&_ContractBLSApkRegistry.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractBLSApkRegistry.Contract.RegistryCoordinator(&_ContractBLSApkRegistry.CallOpts)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "deregisterOperator", operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.DeregisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0xf4e24fe5.
//
// Solidity: function deregisterOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) DeregisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.DeregisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) InitializeQuorum(opts *bind.TransactOpts, quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "initializeQuorum", quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.InitializeQuorum(&_ContractBLSApkRegistry.TransactOpts, quorumNumber)
}

// InitializeQuorum is a paid mutator transaction binding the contract method 0x26d941f2.
//
// Solidity: function initializeQuorum(uint8 quorumNumber) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) InitializeQuorum(quorumNumber uint8) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.InitializeQuorum(&_ContractBLSApkRegistry.TransactOpts, quorumNumber)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "registerBLSPublicKey", operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xbf79ce58.
//
// Solidity: function registerBLSPublicKey(address operator, ((uint256,uint256),(uint256,uint256),(uint256[2],uint256[2])) params, (uint256,uint256) pubkeyRegistrationMessageHash) returns(bytes32 operatorId)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) RegisterBLSPublicKey(operator common.Address, params IBLSApkRegistryTypesPubkeyRegistrationParams, pubkeyRegistrationMessageHash BN254G1Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterBLSPublicKey(&_ContractBLSApkRegistry.TransactOpts, operator, params, pubkeyRegistrationMessageHash)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "registerOperator", operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x3fb27952.
//
// Solidity: function registerOperator(address operator, bytes quorumNumbers) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) RegisterOperator(operator common.Address, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.RegisterOperator(&_ContractBLSApkRegistry.TransactOpts, operator, quorumNumbers)
}

// VerifyAndRegisterG2PubkeyForOperator is a paid mutator transaction binding the contract method 0xd1a64650.
//
// Solidity: function verifyAndRegisterG2PubkeyForOperator(address operator, (uint256[2],uint256[2]) pubkeyG2) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactor) VerifyAndRegisterG2PubkeyForOperator(opts *bind.TransactOpts, operator common.Address, pubkeyG2 BN254G2Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.contract.Transact(opts, "verifyAndRegisterG2PubkeyForOperator", operator, pubkeyG2)
}

// VerifyAndRegisterG2PubkeyForOperator is a paid mutator transaction binding the contract method 0xd1a64650.
//
// Solidity: function verifyAndRegisterG2PubkeyForOperator(address operator, (uint256[2],uint256[2]) pubkeyG2) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistrySession) VerifyAndRegisterG2PubkeyForOperator(operator common.Address, pubkeyG2 BN254G2Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.VerifyAndRegisterG2PubkeyForOperator(&_ContractBLSApkRegistry.TransactOpts, operator, pubkeyG2)
}

// VerifyAndRegisterG2PubkeyForOperator is a paid mutator transaction binding the contract method 0xd1a64650.
//
// Solidity: function verifyAndRegisterG2PubkeyForOperator(address operator, (uint256[2],uint256[2]) pubkeyG2) returns()
func (_ContractBLSApkRegistry *ContractBLSApkRegistryTransactorSession) VerifyAndRegisterG2PubkeyForOperator(operator common.Address, pubkeyG2 BN254G2Point) (*types.Transaction, error) {
	return _ContractBLSApkRegistry.Contract.VerifyAndRegisterG2PubkeyForOperator(&_ContractBLSApkRegistry.TransactOpts, operator, pubkeyG2)
}

// ContractBLSApkRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryInitializedIterator struct {
	Event *ContractBLSApkRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryInitialized)
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
		it.Event = new(ContractBLSApkRegistryInitialized)
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
func (it *ContractBLSApkRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryInitialized represents a Initialized event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractBLSApkRegistryInitializedIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryInitializedIterator{contract: _ContractBLSApkRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryInitialized)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseInitialized(log types.Log) (*ContractBLSApkRegistryInitialized, error) {
	event := new(ContractBLSApkRegistryInitialized)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryNewG2PubkeyRegistrationIterator is returned from FilterNewG2PubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewG2PubkeyRegistration events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewG2PubkeyRegistrationIterator struct {
	Event *ContractBLSApkRegistryNewG2PubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryNewG2PubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryNewG2PubkeyRegistration)
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
		it.Event = new(ContractBLSApkRegistryNewG2PubkeyRegistration)
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
func (it *ContractBLSApkRegistryNewG2PubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryNewG2PubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryNewG2PubkeyRegistration represents a NewG2PubkeyRegistration event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewG2PubkeyRegistration struct {
	Operator common.Address
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewG2PubkeyRegistration is a free log retrieval operation binding the contract event 0x5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a.
//
// Solidity: event NewG2PubkeyRegistration(address indexed operator, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterNewG2PubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewG2PubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "NewG2PubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryNewG2PubkeyRegistrationIterator{contract: _ContractBLSApkRegistry.contract, event: "NewG2PubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewG2PubkeyRegistration is a free log subscription operation binding the contract event 0x5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a.
//
// Solidity: event NewG2PubkeyRegistration(address indexed operator, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchNewG2PubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewG2PubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "NewG2PubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryNewG2PubkeyRegistration)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewG2PubkeyRegistration", log); err != nil {
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

// ParseNewG2PubkeyRegistration is a log parse operation binding the contract event 0x5c4f9f28153dbf3f00e69607a59e82ad806fffb78d09f179f62432f7e9d2511a.
//
// Solidity: event NewG2PubkeyRegistration(address indexed operator, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseNewG2PubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewG2PubkeyRegistration, error) {
	event := new(ContractBLSApkRegistryNewG2PubkeyRegistration)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewG2PubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryNewPubkeyRegistrationIterator is returned from FilterNewPubkeyRegistration and is used to iterate over the raw logs and unpacked data for NewPubkeyRegistration events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewPubkeyRegistrationIterator struct {
	Event *ContractBLSApkRegistryNewPubkeyRegistration // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryNewPubkeyRegistration)
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
		it.Event = new(ContractBLSApkRegistryNewPubkeyRegistration)
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
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryNewPubkeyRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryNewPubkeyRegistration represents a NewPubkeyRegistration event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryNewPubkeyRegistration struct {
	Operator common.Address
	PubkeyG1 BN254G1Point
	PubkeyG2 BN254G2Point
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewPubkeyRegistration is a free log retrieval operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterNewPubkeyRegistration(opts *bind.FilterOpts, operator []common.Address) (*ContractBLSApkRegistryNewPubkeyRegistrationIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryNewPubkeyRegistrationIterator{contract: _ContractBLSApkRegistry.contract, event: "NewPubkeyRegistration", logs: logs, sub: sub}, nil
}

// WatchNewPubkeyRegistration is a free log subscription operation binding the contract event 0xe3fb6613af2e8930cf85d47fcf6db10192224a64c6cbe8023e0eee1ba3828041.
//
// Solidity: event NewPubkeyRegistration(address indexed operator, (uint256,uint256) pubkeyG1, (uint256[2],uint256[2]) pubkeyG2)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchNewPubkeyRegistration(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryNewPubkeyRegistration, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "NewPubkeyRegistration", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryNewPubkeyRegistration)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
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
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseNewPubkeyRegistration(log types.Log) (*ContractBLSApkRegistryNewPubkeyRegistration, error) {
	event := new(ContractBLSApkRegistryNewPubkeyRegistration)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "NewPubkeyRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryOperatorAddedToQuorumsIterator is returned from FilterOperatorAddedToQuorums and is used to iterate over the raw logs and unpacked data for OperatorAddedToQuorums events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorAddedToQuorumsIterator struct {
	Event *ContractBLSApkRegistryOperatorAddedToQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryOperatorAddedToQuorums)
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
		it.Event = new(ContractBLSApkRegistryOperatorAddedToQuorums)
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
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryOperatorAddedToQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryOperatorAddedToQuorums represents a OperatorAddedToQuorums event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorAddedToQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorAddedToQuorums is a free log retrieval operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterOperatorAddedToQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorAddedToQuorumsIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryOperatorAddedToQuorumsIterator{contract: _ContractBLSApkRegistry.contract, event: "OperatorAddedToQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorAddedToQuorums is a free log subscription operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchOperatorAddedToQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorAddedToQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "OperatorAddedToQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryOperatorAddedToQuorums)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
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

// ParseOperatorAddedToQuorums is a log parse operation binding the contract event 0x73a2b7fb844724b971802ae9b15db094d4b7192df9d7350e14eb466b9b22eb4e.
//
// Solidity: event OperatorAddedToQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseOperatorAddedToQuorums(log types.Log) (*ContractBLSApkRegistryOperatorAddedToQuorums, error) {
	event := new(ContractBLSApkRegistryOperatorAddedToQuorums)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorAddedToQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator is returned from FilterOperatorRemovedFromQuorums and is used to iterate over the raw logs and unpacked data for OperatorRemovedFromQuorums events raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator struct {
	Event *ContractBLSApkRegistryOperatorRemovedFromQuorums // Event containing the contract specifics and raw log

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
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
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
		it.Event = new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
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
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBLSApkRegistryOperatorRemovedFromQuorums represents a OperatorRemovedFromQuorums event raised by the ContractBLSApkRegistry contract.
type ContractBLSApkRegistryOperatorRemovedFromQuorums struct {
	Operator      common.Address
	OperatorId    [32]byte
	QuorumNumbers []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemovedFromQuorums is a free log retrieval operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) FilterOperatorRemovedFromQuorums(opts *bind.FilterOpts) (*ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.FilterLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return &ContractBLSApkRegistryOperatorRemovedFromQuorumsIterator{contract: _ContractBLSApkRegistry.contract, event: "OperatorRemovedFromQuorums", logs: logs, sub: sub}, nil
}

// WatchOperatorRemovedFromQuorums is a free log subscription operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) WatchOperatorRemovedFromQuorums(opts *bind.WatchOpts, sink chan<- *ContractBLSApkRegistryOperatorRemovedFromQuorums) (event.Subscription, error) {

	logs, sub, err := _ContractBLSApkRegistry.contract.WatchLogs(opts, "OperatorRemovedFromQuorums")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
				if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
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

// ParseOperatorRemovedFromQuorums is a log parse operation binding the contract event 0xf843ecd53a563675e62107be1494fdde4a3d49aeedaf8d88c616d85346e3500e.
//
// Solidity: event OperatorRemovedFromQuorums(address operator, bytes32 operatorId, bytes quorumNumbers)
func (_ContractBLSApkRegistry *ContractBLSApkRegistryFilterer) ParseOperatorRemovedFromQuorums(log types.Log) (*ContractBLSApkRegistryOperatorRemovedFromQuorums, error) {
	event := new(ContractBLSApkRegistryOperatorRemovedFromQuorums)
	if err := _ContractBLSApkRegistry.contract.UnpackLog(event, "OperatorRemovedFromQuorums", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
