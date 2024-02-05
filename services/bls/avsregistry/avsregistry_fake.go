package avsregistry

import (
	"context"
	"errors"
	"math/big"

	opstateretriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type FakeAvsRegistryService struct {
	operators map[types.BlockNum]map[types.BlsOperatorId]types.OperatorBlsAvsState
}

func NewFakeAvsRegistryService(blockNum types.BlockNum, operators []types.TestBlsOperator) *FakeAvsRegistryService {
	fakeAvsRegistryService := &FakeAvsRegistryService{
		operators: map[types.BlockNum]map[types.BlsOperatorId]types.OperatorBlsAvsState{
			blockNum: {},
		},
	}
	for _, operator := range operators {
		fakeAvsRegistryService.operators[blockNum][operator.OperatorId] = types.OperatorBlsAvsState{
			OperatorId:     operator.OperatorId,
			Pubkeys:        types.OperatorPubkeys{G1Pubkey: operator.BlsKeypair.GetPubKeyG1(), G2Pubkey: operator.BlsKeypair.GetPubKeyG2()},
			StakePerQuorum: operator.StakePerQuorum,
			BlockNumber:    blockNum,
		}
	}
	return fakeAvsRegistryService
}

var _ AvsRegistryService = (*FakeAvsRegistryService)(nil)

func (f *FakeAvsRegistryService) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.BlsOperatorId]types.OperatorBlsAvsState, error) {
	operatorsAvsState, ok := f.operators[blockNumber]
	if !ok {
		return nil, errors.New("block number not found")
	}
	return operatorsAvsState, nil
}

func (f *FakeAvsRegistryService) GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.QuorumNum]types.QuorumBlsAvsState, error) {
	operatorsAvsState, ok := f.operators[blockNumber]
	if !ok {
		return nil, errors.New("block number not found")
	}
	quorumsAvsState := make(map[types.QuorumNum]types.QuorumBlsAvsState)
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
		quorumsAvsState[quorumNum] = types.QuorumBlsAvsState{
			QuorumNumber: quorumNum,
			AggPubkeyG1:  aggPubkeyG1,
			TotalStake:   totalStake,
			BlockNumber:  blockNumber,
		}
	}
	return quorumsAvsState, nil
}

func (f *FakeAvsRegistryService) GetCheckSignaturesIndices(
	opts *bind.CallOpts, referenceBlockNumber types.BlockNum,
	quorumNumbers []types.QuorumNum, nonSignerOperatorIds []types.BlsOperatorId,
) (opstateretriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	return opstateretriever.OperatorStateRetrieverCheckSignaturesIndices{}, nil
}
