package avsregistry

import (
	"context"
	"fmt"
	"math/big"

	avsregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	oppubkeysservice "github.com/Layr-Labs/eigensdk-go/services/operatorpubkeys"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// AvsRegistryServiceChainCaller is a wrapper around AvsRegistryReader that transforms the data into
// nicer golang types that are easier to work with
type AvsRegistryServiceChainCaller struct {
	avsregistry.AvsRegistryReader
	operatorPubkeysService oppubkeysservice.OperatorPubkeysService
	logger                 logging.Logger
}

var _ AvsRegistryService = (*AvsRegistryServiceChainCaller)(nil)

func NewAvsRegistryServiceChainCaller(avsRegistryReader avsregistry.AvsRegistryReader, operatorPubkeysService oppubkeysservice.OperatorPubkeysService, logger logging.Logger) *AvsRegistryServiceChainCaller {
	return &AvsRegistryServiceChainCaller{
		AvsRegistryReader:      avsRegistryReader,
		operatorPubkeysService: operatorPubkeysService,
		logger:                 logger,
	}
}

func (ar *AvsRegistryServiceChainCaller) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers types.QuorumNums, blockNumber types.BlockNum) (map[types.OperatorId]types.OperatorAvsState, error) {
	operatorsAvsState := make(map[types.OperatorId]types.OperatorAvsState)
	// Get operator state for each quorum by querying BLSOperatorStateRetriever (this call is why this service implementation is called ChainCaller)
	operatorsStakesInQuorums, err := ar.AvsRegistryReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{Context: ctx}, quorumNumbers, blockNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get operator state", err)
	}
	numquorums := len(quorumNumbers)
	if len(operatorsStakesInQuorums) != numquorums {
		ar.logger.Fatal("Number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested. Probably pointing to old contract or wrong implementation.", "service", "AvsRegistryServiceChainCaller")
	}

	for quorumIdx, quorumNum := range quorumNumbers {
		for _, operator := range operatorsStakesInQuorums[quorumIdx] {
			pubkeys, err := ar.getOperatorPubkeys(ctx, operator.OperatorId)
			if err != nil {
				return nil, utils.WrapError("Failed to find pubkeys for operator while building operatorsAvsState", err)
			}
			if operatorAvsState, ok := operatorsAvsState[operator.OperatorId]; ok {
				operatorAvsState.StakePerQuorum[quorumNum] = operator.Stake
			} else {
				stakePerQuorum := make(map[types.QuorumNum]types.StakeAmount)
				stakePerQuorum[quorumNum] = operator.Stake
				operatorsAvsState[operator.OperatorId] = types.OperatorAvsState{
					OperatorId:     operator.OperatorId,
					Pubkeys:        pubkeys,
					StakePerQuorum: stakePerQuorum,
					BlockNumber:    blockNumber,
				}
				operatorsAvsState[operator.OperatorId].StakePerQuorum[quorumNum] = operator.Stake
			}
		}
	}

	return operatorsAvsState, nil
}

func (ar *AvsRegistryServiceChainCaller) GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers types.QuorumNums, blockNumber types.BlockNum) (map[types.QuorumNum]types.QuorumAvsState, error) {
	operatorsAvsState, err := ar.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, blockNumber)
	if err != nil {
		return nil, utils.WrapError("Failed to get quorum state", err)
	}
	quorumsAvsState := make(map[types.QuorumNum]types.QuorumAvsState)
	for _, quorumNum := range quorumNumbers {
		aggPubkeyG1 := bls.NewG1Point(big.NewInt(0), big.NewInt(0))
		totalStake := big.NewInt(0)
		for _, operator := range operatorsAvsState {
			// only include operators that have a stake in this quorum
			if stake, ok := operator.StakePerQuorum[quorumNum]; ok {
				aggPubkeyG1.Add(operator.Pubkeys.G1Pubkey)
				totalStake.Add(totalStake, stake)
			}
		}
		quorumsAvsState[quorumNum] = types.QuorumAvsState{
			QuorumNumber: quorumNum,
			AggPubkeyG1:  aggPubkeyG1,
			TotalStake:   totalStake,
			BlockNumber:  blockNumber,
		}
	}
	return quorumsAvsState, nil
}

// getOperatorPubkeys is a temporary hack until we implement GetOperatorAddr on the BLSpubkeyregistry contract
// TODO(samlaf): we need operatorId -> operatorAddr so that we can query the operatorPubkeysService
// this inverse mapping (only operatorAddr->operatorId is stored in registryCoordinator) is not stored,
// but we know that the current implementation uses the hash of the G1 pubkey as the operatorId,
// and the BLSApkRegistry contract stores the mapping G1pubkeyHash -> operatorAddr
// When the above PR is merged, we should change this to instead call GetOperatorAddressFromOperatorId on the avsRegistryReader
// and not hardcode the definition of the operatorId here
func (ar *AvsRegistryServiceChainCaller) getOperatorPubkeys(ctx context.Context, operatorId types.OperatorId) (types.OperatorPubkeys, error) {
	operatorAddr, err := ar.AvsRegistryReader.GetOperatorFromId(&bind.CallOpts{Context: ctx}, operatorId)
	if err != nil {
		return types.OperatorPubkeys{}, utils.WrapError("Failed to get operator address from pubkey hash", err)
	}
	pubkeys, ok := ar.operatorPubkeysService.GetOperatorPubkeys(ctx, operatorAddr)
	if !ok {
		return types.OperatorPubkeys{}, fmt.Errorf("Failed to get operator pubkeys from pubkey compendium service (operatorAddr: %v, operatorId: %v)", operatorAddr, operatorId)
	}
	return pubkeys, nil
}
