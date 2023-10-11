package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common"
)

const (
	ZeroAddress = "0x0000000000000000000000000000000000000000"
)

// Operator represents Eigenlayer's view of an operator
type Operator struct {
	// Address addres of the operator
	Address string `yaml:"address" json:"address"`

	// https://github.com/Layr-Labs/eigenlayer-contracts/blob/delegation-redesign/src/contracts/interfaces/IDelegationManager.sol#L18
	EarningsReceiverAddress   string `yaml:"earnings_receiver_address"    json:"earnings_receiver_address"`
	DelegationApproverAddress string `yaml:"delegation_approver_address"  json:"delegation_approver_address"`
	StakerOptOutWindowBlocks  uint32 `yaml:"staker_opt_out_window_blocks" json:"staker_opt_out_window_blocks"`

	// MetadataUrl URL where operator metadata is stored
	MetadataUrl string `yaml:"metadata_url" json:"metadata_url"`
}

func (o Operator) Validate() error {
	if !utils.IsValidEthereumAddress(o.Address) {
		return errors.New("invalid operator address")
	}

	if !utils.IsValidEthereumAddress(o.EarningsReceiverAddress) {
		return errors.New("invalid EarningsReceiverAddress address")
	}

	if o.DelegationApproverAddress != ZeroAddress && !utils.IsValidEthereumAddress(o.DelegationApproverAddress) {
		return fmt.Errorf(
			"invalid DelegationApproverAddress address, it should be either %s or a valid non zero ethereum address",
			ZeroAddress,
		)
	}

	err := checkIfUrlIsValid(o.MetadataUrl)
	if err != nil {
		return errors.New("invalid metadata url")
	}

	resp, err := http.Get(o.MetadataUrl)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading metadata url body")
		return err
	}

	operatorMetadata := OperatorMetadata{}
	err = json.Unmarshal(body, &operatorMetadata)
	if err != nil {
		fmt.Println("error unmarshalling operator metadata")
		return err
	}

	return operatorMetadata.Validate()
}

// Should we also add the operator address in here?
type OperatorPubkeys struct {
	// G1 signatures are used to verify signatures onchain (since G1 is cheaper to verify onchain via precompiles)
	G1Pubkey *bls.G1Point
	// G2 is used to verify signatures offchain (signatures are on G1)
	G2Pubkey *bls.G2Point
}

// ECDSA address of the operator
type OperatorAddr = common.Address
type StakeAmount = *big.Int

// OperatorId is the ID of an operator, defined by the AVS registry
// It is currently the hash of the operator's G1 pubkey (in the bls pubkey registry)
type OperatorId = [32]byte
type QuorumNum = uint8
type BlockNumber = uint32

// AvsOperator represents the operator state in AVS registries
type OperatorAvsState struct {
	OperatorId OperatorId
	Pubkeys    OperatorPubkeys
	// Stake of the operator for each quorum
	StakePerQuorum map[QuorumNum]StakeAmount
	BlockNumber    BlockNumber
}

var (
	maxNumberOfQuorums = 192
)

func BitmapToQuorumIds(bitmap *big.Int) []QuorumNum {
	// loop through each index in the bitmap to construct the array
	quorumIds := make([]QuorumNum, 0, maxNumberOfQuorums)
	for i := 0; i < maxNumberOfQuorums; i++ {
		if bitmap.Bit(i) == 1 {
			quorumIds = append(quorumIds, QuorumNum(i))
		}
	}
	return quorumIds
}
