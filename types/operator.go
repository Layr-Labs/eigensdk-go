package types

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"

	apkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
)

const (
	ZeroAddress = "0x0000000000000000000000000000000000000000"
)

// Operator represents EigenLayer's view of an operator
type Operator struct {
	// Address of the operator
	Address string `yaml:"address" json:"address"`

	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/delegation-redesign/src/contracts/interfaces/IDelegationManager.sol#L18
	DelegationApproverAddress string `yaml:"delegation_approver_address"  json:"delegation_approver_address"`
	StakerOptOutWindowBlocks  uint32 `yaml:"staker_opt_out_window_blocks" json:"staker_opt_out_window_blocks"`

	// MetadataUrl URL where operator metadata is stored
	MetadataUrl string `yaml:"metadata_url" json:"metadata_url"`

	// AllocationDelay is the delay in seconds where an operator is allowed to change allocation
	// This can only be set once by the operator. Once set this can't be changed
	AllocationDelay uint32 `yaml:"allocation_delay" json:"allocation_delay"`
}

func (o Operator) Validate() error {
	if !utils.IsValidEthereumAddress(o.Address) {
		return ErrInvalidOperatorAddress
	}

	if o.DelegationApproverAddress != ZeroAddress && !utils.IsValidEthereumAddress(o.DelegationApproverAddress) {
		return ErrInvalidDelegationApproverAddress
	}

	err := utils.CheckIfUrlIsValid(o.MetadataUrl)
	if err != nil {
		return utils.WrapError(ErrInvalidMetadataUrl, err)
	}

	body, err := utils.ReadPublicURL(o.MetadataUrl)
	if err != nil {
		return utils.WrapError(ErrReadingMetadataUrlResponse, err)
	}

	operatorMetadata := OperatorMetadata{}
	err = json.Unmarshal(body, &operatorMetadata)
	if err != nil {
		return ErrUnmarshalOperatorMetadata
	}

	return operatorMetadata.Validate()
}

// Socket represents the operator's socket address, which is registered onchain
// TODO: this could have multiple formats... do we really want to use a custom type for this?
// it could be ip:port, or just port, or ip:port:port if 2 ports are needed (like in eigenda's cast)
// or whatever an avs decides to use
type Socket string

func (s Socket) String() string {
	return string(s)
}

type OperatorInfo struct {
	Socket  Socket
	Pubkeys OperatorPubkeys
}

type OperatorPubkeys struct {
	// G1 signatures are used to verify signatures onchain (since G1 is cheaper to verify onchain via precompiles)
	G1Pubkey *bls.G1Point
	// G2 is used to verify signatures offchain (signatures are on G1)
	G2Pubkey *bls.G2Point
}

func (op OperatorPubkeys) ToContractPubkeys() (apkreg.BN254G1Point, apkreg.BN254G2Point) {
	return apkreg.BN254G1Point{
			X: op.G1Pubkey.X.BigInt(new(big.Int)),
			Y: op.G1Pubkey.Y.BigInt(new(big.Int)),
		}, apkreg.BN254G2Point{
			X: [2]*big.Int{op.G2Pubkey.X.A0.BigInt(new(big.Int)), op.G2Pubkey.X.A1.BigInt(new(big.Int))},
			Y: [2]*big.Int{op.G2Pubkey.Y.A0.BigInt(new(big.Int)), op.G2Pubkey.Y.A1.BigInt(new(big.Int))},
		}
}

// ECDSA address of the operator
type OperatorAddr = common.Address
type StakeAmount = *big.Int

// OperatorId is the ID of an operator, defined by the AVS registry
// It is the hash of the operator's G1 pubkey
type OperatorId = Bytes32

func OperatorIdFromG1Pubkey(pubkey *bls.G1Point) OperatorId {
	x := pubkey.X.BigInt(new(big.Int))
	y := pubkey.Y.BigInt(new(big.Int))
	return OperatorId(crypto.Keccak256Hash(append(math.U256Bytes(x), math.U256Bytes(y)...)))
}

func OperatorIdFromContractG1Pubkey(pubkey apkreg.BN254G1Point) OperatorId {
	return OperatorIdFromG1Pubkey(G1PubkeyFromContractG1Pubkey(pubkey))
}

func OperatorIdFromKeyPair(keyPair *bls.KeyPair) OperatorId {
	return OperatorIdFromG1Pubkey(keyPair.GetPubKeyG1())
}

func G1PubkeyFromContractG1Pubkey(pubkey apkreg.BN254G1Point) *bls.G1Point {
	return bls.NewG1Point(pubkey.X, pubkey.Y)
}

type QuorumNums []QuorumNum

func (q QuorumNums) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("%v", q))
}

func (q QuorumNums) UnderlyingType() []uint8 {
	underlying := make([]uint8, len(q))
	for i, v := range q {
		underlying[i] = v.UnderlyingType()
	}
	return underlying
}

type QuorumNum uint8

func (q QuorumNum) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("%d", q))
}
func (q QuorumNum) UnderlyingType() uint8 {
	return uint8(q)
}

type QuorumThresholdPercentages []QuorumThresholdPercentage

func (q QuorumThresholdPercentages) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("%v", q))
}

func (q QuorumThresholdPercentages) UnderlyingType() []uint8 {
	underlying := make([]uint8, len(q))
	for i, v := range q {
		underlying[i] = uint8(v)
	}
	return underlying
}

type QuorumThresholdPercentage uint8

func (q QuorumThresholdPercentage) LogValue() slog.Value {
	return slog.StringValue(fmt.Sprintf("%d", q))
}

type BlockNum = uint32

// AvsOperator represents the operator state in AVS registries
type OperatorAvsState struct {
	OperatorId   OperatorId
	OperatorInfo OperatorInfo
	// Stake of the operator for each quorum
	StakePerQuorum map[QuorumNum]StakeAmount
	BlockNumber    BlockNum
}

var (
	maxNumberOfQuorums uint8 = 192
)

func BitmapToQuorumIds(bitmap *big.Int) []QuorumNum {
	// loop through each index in the bitmap to construct the array
	quorumIds := make([]QuorumNum, 0, maxNumberOfQuorums)
	for i := uint8(0); i < maxNumberOfQuorums; i++ {
		if bitmap.Bit(int(i)) == 1 {
			quorumIds = append(quorumIds, QuorumNum(i))
		}
	}
	return quorumIds
}

type QuorumAvsState struct {
	QuorumNumber QuorumNum
	TotalStake   StakeAmount
	AggPubkeyG1  *bls.G1Point
	BlockNumber  BlockNum
}
