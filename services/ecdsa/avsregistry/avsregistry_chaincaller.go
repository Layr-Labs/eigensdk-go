package avsregistry

import (
	"context"
	"math/big"

	avsecdsaregistry "github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// AvsRegistryServiceChainCaller is a wrapper around AvsRegistryReader that transforms the data into
// nicer golang types that are easier to work with
type AvsRegistryServiceChainCaller struct {
	avsecdsaregistry.AvsEcdsaRegistryReader
	logger logging.Logger
}

var _ AvsRegistryService = (*AvsRegistryServiceChainCaller)(nil)

func NewAvsRegistryServiceChainCaller(avsRegistryReader avsecdsaregistry.AvsEcdsaRegistryReader, logger logging.Logger) *AvsRegistryServiceChainCaller {
	return &AvsRegistryServiceChainCaller{
		AvsEcdsaRegistryReader: avsRegistryReader,
		logger:                 logger,
	}
}

func (ar *AvsRegistryServiceChainCaller) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState, error) {
	operatorsAvsState := make(map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState)
	// Get operator state for each quorum by querying ECDSAOperatorStateRetriever (this call is why this service implementation is called ChainCaller)
	operatorsStakesInQuorums, err := ar.AvsEcdsaRegistryReader.GetOperatorsStakeInQuorumsAtBlock(&bind.CallOpts{Context: ctx}, quorumNumbers, blockNumber)
	if err != nil {
		ar.logger.Error("Failed to get operator state", "err", err, "service", "AvsRegistryServiceChainCaller")
		return nil, err
	}
	numquorums := len(quorumNumbers)
	if len(operatorsStakesInQuorums) != numquorums {
		ar.logger.Fatal("Number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested. Probably pointing to old contract or wrong implementation.", "service", "AvsRegistryServiceChainCaller")
	}

	for quorumIdx, quorumNum := range quorumNumbers {
		for _, operator := range operatorsStakesInQuorums[quorumIdx] {
			if operatorAvsState, ok := operatorsAvsState[operator.OperatorId]; ok {
				operatorAvsState.StakePerQuorum[quorumNum] = operator.Stake
			} else {
				stakePerQuorum := make(map[types.QuorumNum]types.StakeAmount)
				stakePerQuorum[quorumNum] = operator.Stake
				operatorsAvsState[operator.OperatorId] = types.OperatorEcdsaAvsState{
					OperatorId:     operator.OperatorId,
					StakePerQuorum: stakePerQuorum,
					BlockNumber:    blockNumber,
				}
				operatorsAvsState[operator.OperatorId].StakePerQuorum[quorumNum] = operator.Stake
			}
		}
	}

	return operatorsAvsState, nil
}

func (ar *AvsRegistryServiceChainCaller) GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.QuorumNum]types.QuorumEcdsaAvsState, error) {
	operatorsAvsState, err := ar.GetOperatorsAvsStateAtBlock(ctx, quorumNumbers, blockNumber)
	if err != nil {
		ar.logger.Error("Failed to get quorum state", "err", err, "service", "AvsRegistryServiceChainCaller")
		return nil, err
	}
	quorumsAvsState := make(map[types.QuorumNum]types.QuorumEcdsaAvsState)
	for _, quorumNum := range quorumNumbers {
		totalStake := big.NewInt(0)
		operatorIds := make([]types.EcdsaOperatorId, 0, len(operatorsAvsState))
		for _, operator := range operatorsAvsState {
			// only include operators that have a stake in this quorum
			if stake, ok := operator.StakePerQuorum[quorumNum]; ok {
				totalStake.Add(totalStake, stake)
				operatorIds = append(operatorIds, operator.OperatorId)
			}
		}
		quorumsAvsState[quorumNum] = types.QuorumEcdsaAvsState{
			QuorumNumber: quorumNum,
			OperatorIds:  operatorIds,
			TotalStake:   totalStake,
			BlockNumber:  blockNumber,
		}
	}
	return quorumsAvsState, nil
}
