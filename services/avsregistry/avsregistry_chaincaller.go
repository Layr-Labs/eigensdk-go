package avsregistry

import (
	"context"
	"math/big"

	avsregistry "github.com/Layr-Labs/eigensdk-go/chainio/avsregistry"
	elcontracts "github.com/Layr-Labs/eigensdk-go/chainio/elcontracts"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	pcservice "github.com/Layr-Labs/eigensdk-go/services/pubkeycompendium"
	"github.com/Layr-Labs/eigensdk-go/types"
)

type AvsRegistryServiceChainCaller struct {
	elReader                elcontracts.ELReader
	avsRegistryReader       avsregistry.AvsRegistryReader
	pubkeyCompendiumService pcservice.PubkeyCompendiumService
	logger                  logging.Logger
}

var _ AvsRegistryService = (*AvsRegistryServiceChainCaller)(nil)

func NewAvsRegistryServiceChainCaller(avsRegistryReader avsregistry.AvsRegistryReader, elReader elcontracts.ELReader, pubkeyCompendiumService pcservice.PubkeyCompendiumService, logger logging.Logger) *AvsRegistryServiceChainCaller {
	return &AvsRegistryServiceChainCaller{
		elReader:                elReader,
		avsRegistryReader:       avsRegistryReader,
		pubkeyCompendiumService: pubkeyCompendiumService,
		logger:                  logger,
	}
}

func (ar *AvsRegistryServiceChainCaller) GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNumber) (map[types.OperatorId]types.OperatorAvsState, map[types.QuorumNum]*bls.G1Point, error) {
	operatorsAvsState := make(map[types.OperatorId]types.OperatorAvsState)
	// Get operator state for each quorum by querying BLSOperatorStateRetriever (this call is why this service implementation is called ChainCaller)
	operatorsStakesInQuorums, err := ar.avsRegistryReader.GetOperatorsStakeInQuorumsAtBlock(ctx, quorumNumbers, blockNumber)
	if err != nil {
		ar.logger.Error("Failed to get operator state", "err", err, "service", "AvsRegistryServiceChainCaller")
		return nil, nil, err
	}
	numquorums := len(quorumNumbers)
	if len(operatorsStakesInQuorums) != numquorums {
		ar.logger.Fatal("Number of quorums returned from GetOperatorsStakeInQuorumsAtBlock does not match number of quorums requested. Probably pointing to old contract or wrong implementation.", "service", "AvsRegistryServiceChainCaller")
	}

	for quorumIdx, quorumNum := range quorumNumbers {
		for _, operator := range operatorsStakesInQuorums[quorumIdx] {
			pubkeys, err := ar.GetOperatorPubkeys(ctx, operator.OperatorId)
			if err != nil {
				ar.logger.Error("Failed find pubkeys for operator while building operatorsAvsState", "err", err, "service", "AvsRegistryServiceChainCaller")
				return nil, nil, err
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

	aggG1PubkeyPerQuorum := make(map[types.QuorumNum]*bls.G1Point)
	for _, quorumNum := range quorumNumbers {
		aggG1Pubkey := bls.NewG1Point(big.NewInt(0), big.NewInt(0))
		for _, operator := range operatorsAvsState {
			// only include operators that have a stake in this quorum
			if operator.StakePerQuorum != nil {
				aggG1Pubkey.Add(operator.Pubkeys.G1Pubkey)
			}
		}
		aggG1PubkeyPerQuorum[quorumNum] = aggG1Pubkey
	}
	return operatorsAvsState, aggG1PubkeyPerQuorum, nil
}

func (ar *AvsRegistryServiceChainCaller) GetOperatorPubkeys(ctx context.Context, operatorId types.OperatorId) (types.OperatorPubkeys, error) {
	// TODO(samlaf): This is a temporary hack until we implement GetOperatorAddr on the BLSpubkeyregistry contract (shortly)
	// we need operatorId -> operatorAddr so that we can query the pubkeyCompendiumService
	// this inverse mapping (only operatorAddr->operatorId is stored in registryCoordinator) is not stored,
	// but we know that the current implementation uses the hash of the G1 pubkey as the operatorId,
	// and the pubkeycompendium contract stores the mapping G1pubkeyHash -> operatorAddr
	// When the above PR is merged, we should change this to instead call GetOperatorAddressFromOperatorId on the avsRegistryReader
	// and not hardcode the definition of the operatorId here
	operatorAddr, err := ar.elReader.GetOperatorAddressFromPubkeyHash(ctx, operatorId)
	if err != nil {
		ar.logger.Error("Failed to get operator address from pubkey hash", "err", err, "service", "AvsRegistryServiceChainCaller")
		return types.OperatorPubkeys{}, err
	}
	pubkeys, ok := ar.pubkeyCompendiumService.GetOperatorPubkeys(ctx, operatorAddr)
	if !ok {
		ar.logger.Error("Failed to get operator pubkeys from pubkey compendium service", "service", "AvsRegistryServiceChainCaller", "operatorAddr", operatorAddr, "operatorId", operatorId)
		return types.OperatorPubkeys{}, err
	}
	return pubkeys, nil
}
