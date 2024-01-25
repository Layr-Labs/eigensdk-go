package avsregistry

import (
	"context"
	"errors"
	"math/big"

	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type FakeAvsRegistryService struct {
	operators map[types.BlockNum]map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState
}

func NewFakeAvsRegistryService(blockNum types.BlockNum, operators []types.TestEcdsaOperator) *FakeAvsRegistryService {
	fakeAvsRegistryService := &FakeAvsRegistryService{
		operators: map[types.BlockNum]map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState{
			blockNum: {},
		},
	}
	for _, operator := range operators {
		fakeAvsRegistryService.operators[blockNum][operator.OperatorId] = types.OperatorEcdsaAvsState{
			OperatorId:     operator.OperatorId,
			StakePerQuorum: operator.StakePerQuorum,
			BlockNumber:    blockNum,
		}
	}
	return fakeAvsRegistryService
}

var _ AvsRegistryService = (*FakeAvsRegistryService)(nil)

func (f *FakeAvsRegistryService) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState, error) {
	operatorsAvsState, ok := f.operators[blockNumber]
	if !ok {
		return nil, errors.New("block number not found")
	}
	return operatorsAvsState, nil
}

func (f *FakeAvsRegistryService) GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.QuorumNum]types.QuorumEcdsaAvsState, error) {
	operatorsAvsState, ok := f.operators[blockNumber]
	if !ok {
		return nil, errors.New("block number not found")
	}
	quorumsAvsState := make(map[types.QuorumNum]types.QuorumEcdsaAvsState)
	for _, quorumNum := range quorumNumbers {
		operatorIds := make([]types.EcdsaOperatorId, 0, len(operatorsAvsState))
		totalStake := big.NewInt(0)
		for _, operator := range operatorsAvsState {
			// only include operators that have a stake in this quorum
			if stake, ok := operator.StakePerQuorum[quorumNum]; ok {
				operatorIds = append(operatorIds, operator.OperatorId)
				totalStake.Add(totalStake, stake)
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

func (f *FakeAvsRegistryService) GetCheckSignaturesIndices(
	opts *bind.CallOpts, referenceBlockNumber types.BlockNum,
	quorumNumbers []types.QuorumNum, signerOperatorIds []types.EcdsaOperatorId,
) (opstateretriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices, error) {
	return opstateretriever.ECDSAOperatorStateRetrieverCheckSignaturesIndices{}, nil
}
