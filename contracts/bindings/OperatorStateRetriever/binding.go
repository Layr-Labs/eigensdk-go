// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractOperatorStateRetriever

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

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	Operator   common.Address
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractOperatorStateRetrieverMetaData contains all meta data concerning the ContractOperatorStateRetriever contract.
var ContractOperatorStateRetrieverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getBatchOperatorFromId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBatchOperatorId\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckSignaturesIndices\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"referenceBlockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"components\":[{\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\",\"internalType\":\"uint32[][]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOperatorState\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"tuple[][]\",\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stake\",\"type\":\"uint96\",\"internalType\":\"uint96\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getQuorumBitmapsAtBlockNumber\",\"inputs\":[{\"name\":\"registryCoordinator\",\"type\":\"address\",\"internalType\":\"contractISlashingRegistryCoordinator\"},{\"name\":\"operatorIds\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"blockNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"OperatorNotRegistered\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611c088061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c806331b36bd9146100645780633563b0d11461008d5780634d2b57fe146100ad5780634f739f74146100cd5780635c155662146100ed578063cefdc1d41461010d575b5f5ffd5b6100776100723660046112df565b61012e565b60405161008491906113cc565b60405180910390f35b6100a061009b366004611406565b61023f565b6040516100849190611567565b6100c06100bb3660046115d4565b6106a7565b6040516100849190611620565b6100e06100db36600461166b565b6107b1565b604051610084919061175b565b6101006100fb366004611811565b610e2e565b6040516100849190611870565b61012061011b3660046118a7565b610fe3565b6040516100849291906118db565b606081516001600160401b0381111561014957610149611279565b604051908082528060200260200182016040528015610172578160200160208202803683370190505b5090505f5b825181101561023857836001600160a01b03166313542a4e8483815181106101a1576101a16118fb565b60200260200101516040518263ffffffff1660e01b81526004016101d491906001600160a01b0391909116815260200190565b602060405180830381865afa1580156101ef573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610213919061190f565b828281518110610225576102256118fb565b6020908102919091010152600101610177565b5092915050565b60605f846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561027e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102a29190611926565b90505f856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103059190611926565b90505f866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610344573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103689190611926565b90505f86516001600160401b0381111561038457610384611279565b6040519080825280602002602001820160405280156103b757816020015b60608152602001906001900390816103a25790505b5090505f5b875181101561069b575f8882815181106103d8576103d86118fb565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291505f906001600160a01b038716906389026245906044015f60405180830381865afa158015610435573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261045c9190810190611941565b905080516001600160401b0381111561047757610477611279565b6040519080825280602002602001820160405280156104c057816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816104955790505b508484815181106104d3576104d36118fb565b60209081029190910101525f5b8151811015610690576040518060600160405280876001600160a01b03166347b314e8858581518110610515576105156118fb565b60200260200101516040518263ffffffff1660e01b815260040161053b91815260200190565b602060405180830381865afa158015610556573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061057a9190611926565b6001600160a01b0316815260200183838151811061059a5761059a6118fb565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106105c8576105c86118fb565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610622573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061064691906119d1565b6001600160601b0316815250858581518110610664576106646118fb565b6020026020010151828151811061067d5761067d6118fb565b60209081029190910101526001016104e0565b5050506001016103bc565b50979650505050505050565b606081516001600160401b038111156106c2576106c2611279565b6040519080825280602002602001820160405280156106eb578160200160208202803683370190505b5090505f5b825181101561023857836001600160a01b031663296bb06484838151811061071a5761071a6118fb565b60200260200101516040518263ffffffff1660e01b815260040161074091815260200190565b602060405180830381865afa15801561075b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077f9190611926565b828281518110610791576107916118fb565b6001600160a01b03909216602092830291909101909101526001016106f0565b6107dc6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f866001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610819573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083d9190611926565b905061086a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b0389169063c391425e90610898908a9088906004016119f7565b5f60405180830381865afa1580156108b2573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108d99190810190611a15565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061090b908a908a908a90600401611acc565b5f60405180830381865afa158015610925573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261094c9190810190611a15565b6040820152846001600160401b0381111561096957610969611279565b60405190808252806020026020018201604052801561099c57816020015b60608152602001906001900390816109875790505b5060608201525f5b60ff8116861115610d475784515f906001600160401b038111156109ca576109ca611279565b6040519080825280602002602001820160405280156109f3578160200160208202803683370190505b5083606001518360ff1681518110610a0d57610a0d6118fb565b60209081029190910101525f5b8651811015610c53575f8b6001600160a01b03166304ec6351898481518110610a4557610a456118fb565b60200260200101518d885f01518681518110610a6357610a636118fb565b60200260200101516040518463ffffffff1660e01b8152600401610aa09392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610abb573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610adf9190611af4565b9050806001600160c01b03165f03610b0a576040516325ec6c1f60e01b815260040160405180910390fd5b89898560ff16818110610b1f57610b1f6118fb565b60016001600160c01b038516919093013560f81c1c82169091039050610c4a57856001600160a01b031663dd9846b9898481518110610b6057610b606118fb565b60200260200101518c8c8860ff16818110610b7d57610b7d6118fb565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8e166044820152606401602060405180830381865afa158015610bd1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bf59190611b1a565b85606001518560ff1681518110610c0e57610c0e6118fb565b60200260200101518481518110610c2757610c276118fb565b63ffffffff9092166020928302919091019091015282610c4681611b49565b9350505b50600101610a1a565b505f816001600160401b03811115610c6d57610c6d611279565b604051908082528060200260200182016040528015610c96578160200160208202803683370190505b5090505f5b82811015610d0c5784606001518460ff1681518110610cbc57610cbc6118fb565b60200260200101518181518110610cd557610cd56118fb565b6020026020010151828281518110610cef57610cef6118fb565b63ffffffff90921660209283029190910190910152600101610c9b565b508084606001518460ff1681518110610d2757610d276118fb565b602002602001018190525050508080610d3f90611b61565b9150506109a4565b505f886001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d85573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610da99190611926565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90610ddc908a908a908d90600401611b7f565b5f60405180830381865afa158015610df6573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e1d9190810190611a15565b602083015250979650505050505050565b60605f846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401610e5f9291906119f7565b5f60405180830381865afa158015610e79573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ea09190810190611a15565b90505f84516001600160401b03811115610ebc57610ebc611279565b604051908082528060200260200182016040528015610ee5578160200160208202803683370190505b5090505f5b8551811015610fd957866001600160a01b03166304ec6351878381518110610f1457610f146118fb565b602002602001015187868581518110610f2f57610f2f6118fb565b60200260200101516040518463ffffffff1660e01b8152600401610f6c9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610f87573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fab9190611af4565b6001600160c01b0316828281518110610fc657610fc66118fb565b6020908102919091010152600101610eea565b5095945050505050565b6040805160018082528183019092525f9160609183916020808301908036833701905050905084815f8151811061101c5761101c6118fb565b60209081029190910101526040516361c8a12f60e11b81525f906001600160a01b0388169063c391425e9061105790889086906004016119f7565b5f60405180830381865afa158015611071573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110989190810190611a15565b5f815181106110a9576110a96118fb565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291505f906001600160a01b038916906304ec635190606401602060405180830381865afa158015611112573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111369190611af4565b6001600160c01b031690505f61114b82611169565b9050816111598a838a61023f565b9550955050505050935093915050565b60605f5f61117684611232565b61ffff166001600160401b0381111561119157611191611279565b6040519080825280601f01601f1916602001820160405280156111bb576020820181803683370190505b5090505f805b8251821080156111d2575061010081105b15611228576001811b935085841615611218578060f81b8383815181106111fb576111fb6118fb565b60200101906001600160f81b03191690815f1a9053508160010191505b61122181611b49565b90506111c1565b5090949350505050565b5f805b821561125c57611246600184611ba8565b909216918061125481611bbb565b915050611235565b92915050565b6001600160a01b0381168114611276575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156112b5576112b5611279565b604052919050565b5f6001600160401b038211156112d5576112d5611279565b5060051b60200190565b5f5f604083850312156112f0575f5ffd5b82356112fb81611262565b915060208301356001600160401b03811115611315575f5ffd5b8301601f81018513611325575f5ffd5b8035611338611333826112bd565b61128d565b8082825260208201915060208360051b850101925087831115611359575f5ffd5b6020840193505b8284101561138457833561137381611262565b825260209384019390910190611360565b809450505050509250929050565b5f8151808452602084019350602083015f5b828110156113c25781518652602095860195909101906001016113a4565b5093949350505050565b602081525f6113de6020830184611392565b9392505050565b63ffffffff81168114611276575f5ffd5b8035611401816113e5565b919050565b5f5f5f60608486031215611418575f5ffd5b833561142381611262565b925060208401356001600160401b0381111561143d575f5ffd5b8401601f8101861361144d575f5ffd5b80356001600160401b0381111561146657611466611279565b611479601f8201601f191660200161128d565b81815287602083850101111561148d575f5ffd5b816020840160208301375f602083830101528094505050506114b1604085016113f6565b90509250925092565b5f82825180855260208501945060208160051b830101602085015f5b8381101561155b57848303601f19018852815180518085526020918201918501905f5b8181101561154257835180516001600160a01b03168452602080820151818601526040918201516001600160601b031691850191909152909301926060909201916001016114f9565b50506020998a01999094509290920191506001016114d6565b50909695505050505050565b602081525f6113de60208301846114ba565b5f82601f830112611588575f5ffd5b8135611596611333826112bd565b8082825260208201915060208360051b8601019250858311156115b7575f5ffd5b602085015b83811015610fd95780358352602092830192016115bc565b5f5f604083850312156115e5575f5ffd5b82356115f081611262565b915060208301356001600160401b0381111561160a575f5ffd5b61161685828601611579565b9150509250929050565b602080825282518282018190525f918401906040840190835b818110156116605783516001600160a01b0316835260209384019390920191600101611639565b509095945050505050565b5f5f5f5f5f6080868803121561167f575f5ffd5b853561168a81611262565b9450602086013561169a816113e5565b935060408601356001600160401b038111156116b4575f5ffd5b8601601f810188136116c4575f5ffd5b80356001600160401b038111156116d9575f5ffd5b8860208284010111156116ea575f5ffd5b6020919091019350915060608601356001600160401b0381111561170c575f5ffd5b61171888828901611579565b9150509295509295909350565b5f8151808452602084019350602083015f5b828110156113c257815163ffffffff16865260209586019590910190600101611737565b602081525f82516080602084015261177660a0840182611725565b90506020840151601f198483030160408501526117938282611725565b9150506040840151601f198483030160608501526117b18282611725565b6060860151858203601f190160808701528051808352919350602090810192508084019190600582901b8501015f5b8281101561069b57601f198683030184526117fc828651611725565b602095860195949094019391506001016117e0565b5f5f5f60608486031215611823575f5ffd5b833561182e81611262565b925060208401356001600160401b03811115611848575f5ffd5b61185486828701611579565b9250506040840135611865816113e5565b809150509250925092565b602080825282518282018190525f918401906040840190835b81811015611660578351835260209384019390920191600101611889565b5f5f5f606084860312156118b9575f5ffd5b83356118c481611262565b9250602084013591506040840135611865816113e5565b828152604060208201525f6118f360408301846114ba565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f6020828403121561191f575f5ffd5b5051919050565b5f60208284031215611936575f5ffd5b81516113de81611262565b5f60208284031215611951575f5ffd5b81516001600160401b03811115611966575f5ffd5b8201601f81018413611976575f5ffd5b8051611984611333826112bd565b8082825260208201915060208360051b8501019250868311156119a5575f5ffd5b6020840193505b828410156119c75783518252602093840193909101906119ac565b9695505050505050565b5f602082840312156119e1575f5ffd5b81516001600160601b03811681146113de575f5ffd5b63ffffffff83168152604060208201525f6118f36040830184611392565b5f60208284031215611a25575f5ffd5b81516001600160401b03811115611a3a575f5ffd5b8201601f81018413611a4a575f5ffd5b8051611a58611333826112bd565b8082825260208201915060208360051b850101925086831115611a79575f5ffd5b6020840193505b828410156119c7578351611a93816113e5565b825260209384019390910190611a80565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201525f611aeb604083018486611aa4565b95945050505050565b5f60208284031215611b04575f5ffd5b81516001600160c01b03811681146113de575f5ffd5b5f60208284031215611b2a575f5ffd5b81516113de816113e5565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611b5a57611b5a611b35565b5060010190565b5f60ff821660ff8103611b7657611b76611b35565b60010192915050565b604081525f611b92604083018587611aa4565b905063ffffffff83166020830152949350505050565b8181038181111561125c5761125c611b35565b5f61ffff821661ffff8103611b7657611b76611b3556fea2646970667358221220665dc0bb54347cd1e6f719d5a2f5e89ec80a6bd64c662badfb5f23281d2b402364736f6c634300081b0033",
}

// ContractOperatorStateRetrieverABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.ABI instead.
var ContractOperatorStateRetrieverABI = ContractOperatorStateRetrieverMetaData.ABI

// ContractOperatorStateRetrieverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractOperatorStateRetrieverMetaData.Bin instead.
var ContractOperatorStateRetrieverBin = ContractOperatorStateRetrieverMetaData.Bin

// DeployContractOperatorStateRetriever deploys a new Ethereum contract, binding an instance of ContractOperatorStateRetriever to it.
func DeployContractOperatorStateRetriever(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractOperatorStateRetriever, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractOperatorStateRetrieverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// ContractOperatorStateRetrieverMethods is an auto generated interface around an Ethereum contract.
type ContractOperatorStateRetrieverMethods interface {
	ContractOperatorStateRetrieverCalls
	ContractOperatorStateRetrieverTransacts
	ContractOperatorStateRetrieverFilters
}

// ContractOperatorStateRetrieverCalls is an auto generated interface that defines the call methods available for an Ethereum contract.
type ContractOperatorStateRetrieverCalls interface {
	GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error)

	GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error)

	GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error)

	GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error)

	GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error)

	GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error)
}

// ContractOperatorStateRetrieverTransacts is an auto generated interface that defines the transact methods available for an Ethereum contract.
type ContractOperatorStateRetrieverTransacts interface {
}

// ContractOperatorStateRetrieverFilterer is an auto generated interface that defines the log filtering methods available for an Ethereum contract.
type ContractOperatorStateRetrieverFilters interface {
}

// ContractOperatorStateRetriever is an auto generated Go binding around an Ethereum contract.
type ContractOperatorStateRetriever struct {
	ContractOperatorStateRetrieverCaller     // Read-only binding to the contract
	ContractOperatorStateRetrieverTransactor // Write-only binding to the contract
	ContractOperatorStateRetrieverFilterer   // Log filterer for contract events
}

// ContractOperatorStateRetriever implements the ContractOperatorStateRetrieverMethods interface.
var _ ContractOperatorStateRetrieverMethods = (*ContractOperatorStateRetriever)(nil)

// ContractOperatorStateRetrieverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverCaller implements the ContractOperatorStateRetrieverCalls interface.
var _ ContractOperatorStateRetrieverCalls = (*ContractOperatorStateRetrieverCaller)(nil)

// ContractOperatorStateRetrieverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverTransactor implements the ContractOperatorStateRetrieverTransacts interface.
var _ ContractOperatorStateRetrieverTransacts = (*ContractOperatorStateRetrieverTransactor)(nil)

// ContractOperatorStateRetrieverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractOperatorStateRetrieverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractOperatorStateRetrieverFilterer implements the ContractOperatorStateRetrieverFilters interface.
var _ ContractOperatorStateRetrieverFilters = (*ContractOperatorStateRetrieverFilterer)(nil)

// ContractOperatorStateRetrieverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractOperatorStateRetrieverSession struct {
	Contract     *ContractOperatorStateRetriever // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractOperatorStateRetrieverCallerSession struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// ContractOperatorStateRetrieverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractOperatorStateRetrieverTransactorSession struct {
	Contract     *ContractOperatorStateRetrieverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// ContractOperatorStateRetrieverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverRaw struct {
	Contract *ContractOperatorStateRetriever // Generic contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverCallerRaw struct {
	Contract *ContractOperatorStateRetrieverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractOperatorStateRetrieverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractOperatorStateRetrieverTransactorRaw struct {
	Contract *ContractOperatorStateRetrieverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractOperatorStateRetriever creates a new instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetriever(address common.Address, backend bind.ContractBackend) (*ContractOperatorStateRetriever, error) {
	contract, err := bindContractOperatorStateRetriever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetriever{ContractOperatorStateRetrieverCaller: ContractOperatorStateRetrieverCaller{contract: contract}, ContractOperatorStateRetrieverTransactor: ContractOperatorStateRetrieverTransactor{contract: contract}, ContractOperatorStateRetrieverFilterer: ContractOperatorStateRetrieverFilterer{contract: contract}}, nil
}

// NewContractOperatorStateRetrieverCaller creates a new read-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverCaller(address common.Address, caller bind.ContractCaller) (*ContractOperatorStateRetrieverCaller, error) {
	contract, err := bindContractOperatorStateRetriever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverCaller{contract: contract}, nil
}

// NewContractOperatorStateRetrieverTransactor creates a new write-only instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractOperatorStateRetrieverTransactor, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverTransactor{contract: contract}, nil
}

// NewContractOperatorStateRetrieverFilterer creates a new log filterer instance of ContractOperatorStateRetriever, bound to a specific deployed contract.
func NewContractOperatorStateRetrieverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractOperatorStateRetrieverFilterer, error) {
	contract, err := bindContractOperatorStateRetriever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractOperatorStateRetrieverFilterer{contract: contract}, nil
}

// bindContractOperatorStateRetriever binds a generic wrapper to an already deployed contract.
func bindContractOperatorStateRetriever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractOperatorStateRetrieverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.ContractOperatorStateRetrieverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractOperatorStateRetriever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractOperatorStateRetriever.Contract.contract.Transact(opts, method, params...)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorFromId(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorFromId", registryCoordinator, operatorIds)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorFromId is a free data retrieval call binding the contract method 0x4d2b57fe.
//
// Solidity: function getBatchOperatorFromId(address registryCoordinator, bytes32[] operatorIds) view returns(address[] operators)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorFromId(registryCoordinator common.Address, operatorIds [][32]byte) ([]common.Address, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorFromId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetBatchOperatorId(opts *bind.CallOpts, registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getBatchOperatorId", registryCoordinator, operators)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetBatchOperatorId is a free data retrieval call binding the contract method 0x31b36bd9.
//
// Solidity: function getBatchOperatorId(address registryCoordinator, address[] operators) view returns(bytes32[] operatorIds)
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetBatchOperatorId(registryCoordinator common.Address, operators []common.Address) ([][32]byte, error) {
	return _ContractOperatorStateRetriever.Contract.GetBatchOperatorId(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operators)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractOperatorStateRetriever.Contract.GetCheckSignaturesIndices(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (address,bytes32,uint96)[][])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractOperatorStateRetriever.Contract.GetOperatorState0(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCaller) GetQuorumBitmapsAtBlockNumber(opts *bind.CallOpts, registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ContractOperatorStateRetriever.contract.Call(opts, &out, "getQuorumBitmapsAtBlockNumber", registryCoordinator, operatorIds, blockNumber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}

// GetQuorumBitmapsAtBlockNumber is a free data retrieval call binding the contract method 0x5c155662.
//
// Solidity: function getQuorumBitmapsAtBlockNumber(address registryCoordinator, bytes32[] operatorIds, uint32 blockNumber) view returns(uint256[])
func (_ContractOperatorStateRetriever *ContractOperatorStateRetrieverCallerSession) GetQuorumBitmapsAtBlockNumber(registryCoordinator common.Address, operatorIds [][32]byte, blockNumber uint32) ([]*big.Int, error) {
	return _ContractOperatorStateRetriever.Contract.GetQuorumBitmapsAtBlockNumber(&_ContractOperatorStateRetriever.CallOpts, registryCoordinator, operatorIds, blockNumber)
}
