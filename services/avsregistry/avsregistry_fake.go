package avsregistry

import (
	"context"
	"errors"
	"math/big"

	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type FakeAvsRegistryService struct {
	operators map[types.BlockNum]map[types.OperatorId]types.OperatorAvsState
}

func NewFakeAvsRegistryService(blockNum types.BlockNum, operators []types.TestOperator) *FakeAvsRegistryService {
	fakeAvsRegistryService := &FakeAvsRegistryService{
		operators: map[types.BlockNum]map[types.OperatorId]types.OperatorAvsState{
			blockNum: {},
		},
	}
	for _, operator := range operators {
		fakeAvsRegistryService.operators[blockNum][operator.OperatorId] = types.OperatorAvsState{
			OperatorId:     operator.OperatorId,
			Pubkeys:        types.OperatorPubkeys{G1Pubkey: operator.BlsKeypair.GetPubKeyG1(), G2Pubkey: operator.BlsKeypair.GetPubKeyG2()},
			StakePerQuorum: operator.StakePerQuorum,
			BlockNumber:    blockNum,
		}
	}
	return fakeAvsRegistryService
}

var _ AvsRegistryService = (*FakeAvsRegistryService)(nil)

func (f *FakeAvsRegistryService) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.OperatorId]types.OperatorAvsState, error) {
	operatorsAvsState, ok := f.operators[blockNumber]
	if !ok {
		return nil, errors.New("block number not found")
	}
	return operatorsAvsState, nil
}

func (f *FakeAvsRegistryService) GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.QuorumNum]types.QuorumAvsState, error) {
	operatorsAvsState, ok := f.operators[blockNumber]
	if !ok {
		return nil, errors.New("block number not found")
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

func (f *FakeAvsRegistryService) GetCheckSignaturesIndices(ctx context.Context, referenceBlockNumber types.BlockNum, quorumNumbers []types.QuorumNum, nonSignerOperatorIds []types.OperatorId) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	return blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices{}, nil
}
