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
	Bin: "0x6080604052348015600e575f5ffd5b50611c988061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c806331b36bd9146100645780633563b0d11461008d5780634d2b57fe146100ad5780634f739f74146100cd5780635c155662146100ed578063cefdc1d41461010d575b5f5ffd5b6100776100723660046112dd565b61012e565b60405161008491906113ca565b60405180910390f35b6100a061009b366004611404565b61023f565b6040516100849190611565565b6100c06100bb3660046115d2565b6106a7565b604051610084919061161e565b6100e06100db3660046116b0565b6107b1565b60405161008491906117a6565b6101006100fb36600461185c565b610e2c565b60405161008491906118bb565b61012061011b3660046118f2565b610fe1565b604051610084929190611926565b606081516001600160401b0381111561014957610149611277565b604051908082528060200260200182016040528015610172578160200160208202803683370190505b5090505f5b825181101561023857836001600160a01b03166313542a4e8483815181106101a1576101a1611946565b60200260200101516040518263ffffffff1660e01b81526004016101d491906001600160a01b0391909116815260200190565b602060405180830381865afa1580156101ef573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610213919061195a565b82828151811061022557610225611946565b6020908102919091010152600101610177565b5092915050565b60605f846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa15801561027e573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102a29190611971565b90505f856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102e1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103059190611971565b90505f866001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610344573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103689190611971565b90505f86516001600160401b0381111561038457610384611277565b6040519080825280602002602001820160405280156103b757816020015b60608152602001906001900390816103a25790505b5090505f5b875181101561069b575f8882815181106103d8576103d8611946565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8a16602483015291505f906001600160a01b038716906389026245906044015f60405180830381865afa158015610435573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261045c919081019061198c565b905080516001600160401b0381111561047757610477611277565b6040519080825280602002602001820160405280156104c057816020015b604080516060810182525f80825260208083018290529282015282525f199092019101816104955790505b508484815181106104d3576104d3611946565b60209081029190910101525f5b8151811015610690576040518060600160405280876001600160a01b03166347b314e885858151811061051557610515611946565b60200260200101516040518263ffffffff1660e01b815260040161053b91815260200190565b602060405180830381865afa158015610556573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061057a9190611971565b6001600160a01b0316815260200183838151811061059a5761059a611946565b60200260200101518152602001896001600160a01b031663fa28c6278585815181106105c8576105c8611946565b60209081029190910101516040516001600160e01b031960e084901b168152600481019190915260ff8816602482015263ffffffff8f166044820152606401602060405180830381865afa158015610622573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106469190611a1c565b6001600160601b031681525085858151811061066457610664611946565b6020026020010151828151811061067d5761067d611946565b60209081029190910101526001016104e0565b5050506001016103bc565b50979650505050505050565b606081516001600160401b038111156106c2576106c2611277565b6040519080825280602002602001820160405280156106eb578160200160208202803683370190505b5090505f5b825181101561023857836001600160a01b031663296bb06484838151811061071a5761071a611946565b60200260200101516040518263ffffffff1660e01b815260040161074091815260200190565b602060405180830381865afa15801561075b573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061077f9190611971565b82828151811061079157610791611946565b6001600160a01b03909216602092830291909101909101526001016106f0565b6107dc6040518060800160405280606081526020016060815260200160608152602001606081525090565b5f876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610819573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083d9190611971565b905061086a6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e9061089a908b9089908990600401611a42565b5f60405180830381865afa1580156108b4573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108db9190810190611a87565b81526040516340e03a8160e11b81526001600160a01b038316906381c075029061090d908b908b908b90600401611b3e565b5f60405180830381865afa158015610927573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261094e9190810190611a87565b6040820152856001600160401b0381111561096b5761096b611277565b60405190808252806020026020018201604052801561099e57816020015b60608152602001906001900390816109895790505b5060608201525f5b60ff8116871115610d44575f856001600160401b038111156109ca576109ca611277565b6040519080825280602002602001820160405280156109f3578160200160208202803683370190505b5083606001518360ff1681518110610a0d57610a0d611946565b60209081029190910101525f5b86811015610c50575f8c6001600160a01b03166304ec63518a8a85818110610a4457610a44611946565b905060200201358e885f01518681518110610a6157610a61611946565b60200260200101516040518463ffffffff1660e01b8152600401610a9e9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610ab9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610add9190611b66565b9050806001600160c01b03165f03610b08576040516325ec6c1f60e01b815260040160405180910390fd5b8a8a8560ff16818110610b1d57610b1d611946565b60016001600160c01b038516919093013560f81c1c82169091039050610c4757856001600160a01b031663dd9846b98a8a85818110610b5e57610b5e611946565b905060200201358d8d8860ff16818110610b7a57610b7a611946565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015610bce573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610bf29190611b8c565b85606001518560ff1681518110610c0b57610c0b611946565b60200260200101518481518110610c2457610c24611946565b63ffffffff9092166020928302919091019091015282610c4381611bbb565b9350505b50600101610a1a565b505f816001600160401b03811115610c6a57610c6a611277565b604051908082528060200260200182016040528015610c93578160200160208202803683370190505b5090505f5b82811015610d095784606001518460ff1681518110610cb957610cb9611946565b60200260200101518181518110610cd257610cd2611946565b6020026020010151828281518110610cec57610cec611946565b63ffffffff90921660209283029190910190910152600101610c98565b508084606001518460ff1681518110610d2457610d24611946565b602002602001018190525050508080610d3c90611bd3565b9150506109a6565b505f896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d82573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610da69190611971565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c90610dd9908b908b908e90600401611bf1565b5f60405180830381865afa158015610df3573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e1a9190810190611a87565b60208301525098975050505050505050565b60605f846001600160a01b031663c391425e84866040518363ffffffff1660e01b8152600401610e5d929190611c1a565b5f60405180830381865afa158015610e77573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e9e9190810190611a87565b90505f84516001600160401b03811115610eba57610eba611277565b604051908082528060200260200182016040528015610ee3578160200160208202803683370190505b5090505f5b8551811015610fd757866001600160a01b03166304ec6351878381518110610f1257610f12611946565b602002602001015187868581518110610f2d57610f2d611946565b60200260200101516040518463ffffffff1660e01b8152600401610f6a9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015610f85573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fa99190611b66565b6001600160c01b0316828281518110610fc457610fc4611946565b6020908102919091010152600101610ee8565b5095945050505050565b6040805160018082528183019092525f9160609183916020808301908036833701905050905084815f8151811061101a5761101a611946565b60209081029190910101526040516361c8a12f60e11b81525f906001600160a01b0388169063c391425e906110559088908690600401611c1a565b5f60405180830381865afa15801561106f573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526110969190810190611a87565b5f815181106110a7576110a7611946565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291505f906001600160a01b038916906304ec635190606401602060405180830381865afa158015611110573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111349190611b66565b6001600160c01b031690505f61114982611167565b9050816111578a838a61023f565b9550955050505050935093915050565b60605f5f61117484611230565b61ffff166001600160401b0381111561118f5761118f611277565b6040519080825280601f01601f1916602001820160405280156111b9576020820181803683370190505b5090505f805b8251821080156111d0575061010081105b15611226576001811b935085841615611216578060f81b8383815181106111f9576111f9611946565b60200101906001600160f81b03191690815f1a9053508160010191505b61121f81611bbb565b90506111bf565b5090949350505050565b5f805b821561125a57611244600184611c38565b909216918061125281611c4b565b915050611233565b92915050565b6001600160a01b0381168114611274575f5ffd5b50565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156112b3576112b3611277565b604052919050565b5f6001600160401b038211156112d3576112d3611277565b5060051b60200190565b5f5f604083850312156112ee575f5ffd5b82356112f981611260565b915060208301356001600160401b03811115611313575f5ffd5b8301601f81018513611323575f5ffd5b8035611336611331826112bb565b61128b565b8082825260208201915060208360051b850101925087831115611357575f5ffd5b6020840193505b8284101561138257833561137181611260565b82526020938401939091019061135e565b809450505050509250929050565b5f8151808452602084019350602083015f5b828110156113c05781518652602095860195909101906001016113a2565b5093949350505050565b602081525f6113dc6020830184611390565b9392505050565b63ffffffff81168114611274575f5ffd5b80356113ff816113e3565b919050565b5f5f5f60608486031215611416575f5ffd5b833561142181611260565b925060208401356001600160401b0381111561143b575f5ffd5b8401601f8101861361144b575f5ffd5b80356001600160401b0381111561146457611464611277565b611477601f8201601f191660200161128b565b81815287602083850101111561148b575f5ffd5b816020840160208301375f602083830101528094505050506114af604085016113f4565b90509250925092565b5f82825180855260208501945060208160051b830101602085015f5b8381101561155957848303601f19018852815180518085526020918201918501905f5b8181101561154057835180516001600160a01b03168452602080820151818601526040918201516001600160601b031691850191909152909301926060909201916001016114f7565b50506020998a01999094509290920191506001016114d4565b50909695505050505050565b602081525f6113dc60208301846114b8565b5f82601f830112611586575f5ffd5b8135611594611331826112bb565b8082825260208201915060208360051b8601019250858311156115b5575f5ffd5b602085015b83811015610fd75780358352602092830192016115ba565b5f5f604083850312156115e3575f5ffd5b82356115ee81611260565b915060208301356001600160401b03811115611608575f5ffd5b61161485828601611577565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561165e5783516001600160a01b0316835260209384019390920191600101611637565b509095945050505050565b5f5f83601f840112611679575f5ffd5b5081356001600160401b0381111561168f575f5ffd5b6020830191508360208260051b85010111156116a9575f5ffd5b9250929050565b5f5f5f5f5f5f608087890312156116c5575f5ffd5b86356116d081611260565b955060208701356116e0816113e3565b945060408701356001600160401b038111156116fa575f5ffd5b8701601f8101891361170a575f5ffd5b80356001600160401b0381111561171f575f5ffd5b896020828401011115611730575f5ffd5b6020919091019450925060608701356001600160401b03811115611752575f5ffd5b61175e89828a01611669565b979a9699509497509295939492505050565b5f8151808452602084019350602083015f5b828110156113c057815163ffffffff16865260209586019590910190600101611782565b602081525f8251608060208401526117c160a0840182611770565b90506020840151601f198483030160408501526117de8282611770565b9150506040840151601f198483030160608501526117fc8282611770565b6060860151858203601f190160808701528051808352919350602090810192508084019190600582901b8501015f5b8281101561069b57601f19868303018452611847828651611770565b6020958601959490940193915060010161182b565b5f5f5f6060848603121561186e575f5ffd5b833561187981611260565b925060208401356001600160401b03811115611893575f5ffd5b61189f86828701611577565b92505060408401356118b0816113e3565b809150509250925092565b602080825282518282018190525f918401906040840190835b8181101561165e5783518352602093840193909201916001016118d4565b5f5f5f60608486031215611904575f5ffd5b833561190f81611260565b92506020840135915060408401356118b0816113e3565b828152604060208201525f61193e60408301846114b8565b949350505050565b634e487b7160e01b5f52603260045260245ffd5b5f6020828403121561196a575f5ffd5b5051919050565b5f60208284031215611981575f5ffd5b81516113dc81611260565b5f6020828403121561199c575f5ffd5b81516001600160401b038111156119b1575f5ffd5b8201601f810184136119c1575f5ffd5b80516119cf611331826112bb565b8082825260208201915060208360051b8501019250868311156119f0575f5ffd5b6020840193505b82841015611a125783518252602093840193909101906119f7565b9695505050505050565b5f60208284031215611a2c575f5ffd5b81516001600160601b03811681146113dc575f5ffd5b63ffffffff8416815260406020820181905281018290525f6001600160fb1b03831115611a6d575f5ffd5b8260051b8085606085013791909101606001949350505050565b5f60208284031215611a97575f5ffd5b81516001600160401b03811115611aac575f5ffd5b8201601f81018413611abc575f5ffd5b8051611aca611331826112bb565b8082825260208201915060208360051b850101925086831115611aeb575f5ffd5b6020840193505b82841015611a12578351611b05816113e3565b825260209384019390910190611af2565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b63ffffffff84168152604060208201525f611b5d604083018486611b16565b95945050505050565b5f60208284031215611b76575f5ffd5b81516001600160c01b03811681146113dc575f5ffd5b5f60208284031215611b9c575f5ffd5b81516113dc816113e3565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611bcc57611bcc611ba7565b5060010190565b5f60ff821660ff8103611be857611be8611ba7565b60010192915050565b604081525f611c04604083018587611b16565b905063ffffffff83166020830152949350505050565b63ffffffff83168152604060208201525f61193e6040830184611390565b8181038181111561125a5761125a611ba7565b5f61ffff821661ffff8103611be857611be8611ba756fea26469706673582212204b70de87e37d0443da5aa4bd66a94f7b1236832726a126631a2b32cb6ffd43ee64736f6c634300081b0033",
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
