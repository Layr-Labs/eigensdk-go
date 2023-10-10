package avsregistry

import (
	"context"

	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/types"
)

// AvsRegistryServicemService is a service that indexes the Avs Registry contracts and provides a way to query for operator state
// at certain blocks, including operatorIds, pubkeys, and staking status in each quorum.
type AvsRegistryService interface {
	// GetOperatorsAvsState returns the state of an avs wrt to a list of quorums at a certain block.
	// The state includes the operatorId, pubkey, and staking amount in each quorum.
	GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.OperatorId]types.OperatorAvsState, error)
	// GetQuorumsAvsStateAtBlock returns the aggregated data for a list of quorums at a certain block.
	// The aggregated data includes the aggregated pubkey and total stake in each quorum.
	// This information is derivable from the Operators Avs State (returned from GetOperatorsAvsStateAtBlock), but this function is provided for convenience.
	GetQuorumsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNum) (map[types.QuorumNum]types.QuorumAvsState, error)
	// GetCheckSignaturesIndices returns the registry indices of the nonsigner operators specified by nonSignerOperatorIds who were registered at referenceBlockNumber.
	GetCheckSignaturesIndices(ctx context.Context, referenceBlockNumber types.BlockNum, quorumNumbers []types.QuorumNum, nonSignerOperatorIds []types.OperatorId) (blsoperatorstateretrievar.BLSOperatorStateRetrieverCheckSignaturesIndices, error)
}
