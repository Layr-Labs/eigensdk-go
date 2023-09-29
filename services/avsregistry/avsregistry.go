package avsregistry

import (
	"context"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/types"
)

// AvsRegistryServicemService is a service that indexes the Avs Registry contracts and provides a way to query for operator state
// at certain blocks, including operatorIds, pubkeys, and staking status in each quorum.
type AvsRegistryService interface {
	// GetOperatorsAvsState returns the state of an avs wrt to a list of quorums at a certain block.
	// The state includes the operatorId, pubkey, and staking amount in each quorum.
	GetOperatorsAvsStateAtBlock(ctx context.Context, quorumNumbers []types.QuorumNum, blockNumber types.BlockNumber) (map[types.OperatorId]types.OperatorAvsState, map[types.QuorumNum]*bls.G1Point, error)
	GetOperatorPubkeys(ctx context.Context, operatorId types.OperatorId) (types.OperatorPubkeys, error)
}
