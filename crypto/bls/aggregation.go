package bls

import (
	"errors"
	"math/big"

	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const PercentMultiplier = 100

var (
	ErrPubKeysNotEqual     = errors.New("public keys are not equal")
	ErrInsufficientEthSigs = errors.New("insufficient eth signatures")
	ErrAggSigNotValid      = errors.New("aggregated signature is not valid")
)

type SignerMessage struct {
	Signature *Signature
	Operator  OperatorId
	Err       error
}

type SignatureAggregation struct {
	NonSigners       []*G1Point
	QuorumAggPubKeys []*G1Point
	AggPubKey        *G2Point
	AggSignature     *Signature
}

// SignatureAggregator is an interface for aggregating the signatures returned by DA nodes so that they can be verified
// by the DA contract
type SignatureAggregator interface {

	// AggregateSignatures blocks until it recieves a response for each operator in the operator state via messageChan,
	// and then returns the aggregated signature.
	// If the aggregated signature is invalid, an error is returned.
	AggregateSignatures(
		state *IndexedOperatorState,
		quorumParams []QuorumParam,
		message [32]byte,
		messageChan chan SignerMessage,
	) (*SignatureAggregation, error)
}

type StdSignatureAggregator struct {
	Logger logging.Logger
}

func NewStdSignatureAggregator(logger logging.Logger) *StdSignatureAggregator {
	return &StdSignatureAggregator{
		Logger: logger,
	}
}

var _ SignatureAggregator = (*StdSignatureAggregator)(nil)

func (a *StdSignatureAggregator) AggregateSignatures(
	state *IndexedOperatorState,
	quorumParams []QuorumParam,
	message [32]byte,
	messageChan chan SignerMessage,
) (*SignatureAggregation, error) {

	// TODO: Add logging

	// Ensure all quorums are found in state
	for _, quorum := range quorumParams {
		_, found := state.Operators[quorum.QuorumID]
		if !found {
			return nil, errors.New("quorum not found")
		}
	}

	stakeSigned := make([]*big.Int, len(quorumParams))
	for ind := range quorumParams {
		stakeSigned[ind] = big.NewInt(0)
	}
	aggSigs := make([]*Signature, len(quorumParams))
	aggPubKeys := make([]*G2Point, len(quorumParams))

	signerMap := make(map[OperatorId]bool)

	// Aggregate Signatures
	numOperators := len(state.IndexedOperators)

	for numReply := 0; numReply < numOperators; numReply++ {
		r := <-messageChan

		if r.Err != nil {
			a.Logger.Warn("Error returned from messageChan", "err", r.Err)
			continue
		}

		op, found := state.IndexedOperators[r.Operator]
		if !found {
			a.Logger.Error("Operator not found in state", "operator", r.Operator)
			continue
		}

		// Verify Signature
		sig := r.Signature
		ok, err := sig.Verify(op.PubkeyG2, message)
		if err != nil {
			a.Logger.Error("Error verifying signature", "err", err)
			continue
		}
		if !ok {
			a.Logger.Error("Signature is not valid", "pubkey", hexutil.Encode(op.PubkeyG2.Serialize()))
			continue
		}

		for ind, quorumParam := range quorumParams {

			// Get stake amounts for operator
			ops := state.Operators[quorumParam.QuorumID]
			opInfo, ok := ops[r.Operator]

			// If operator is not in quorum, skip
			if !ok {
				a.Logger.Error("Operator not found in quorum", "operator", r.Operator)
				continue
			}

			signerMap[r.Operator] = true

			// Add to stake signed
			stakeSigned[ind].Add(stakeSigned[ind], opInfo.Stake)

			// Add to agg signature
			if aggSigs[ind] == nil {
				aggSigs[ind] = &Signature{sig.Deserialize(sig.Serialize())}
				aggPubKeys[ind] = op.PubkeyG2.Deserialize(op.PubkeyG2.Serialize())
			} else {
				aggSigs[ind].Add(sig.G1Point)
				aggPubKeys[ind].Add(op.PubkeyG2)
			}

		}

	}

	// Aggregrate Non signer Pubkey Id
	nonSignerKeys := make([]*G1Point, 0)
	nonSignerOperatorIds := make([]OperatorId, 0)

	for id, op := range state.IndexedOperators {
		_, found := signerMap[id]
		if !found {
			nonSignerKeys = append(nonSignerKeys, op.PubkeyG1)
			nonSignerOperatorIds = append(nonSignerOperatorIds, id)
		}
	}

	quorumAggPubKeys := make([]*G1Point, len(quorumParams))

	// Validate the amount signed and aggregate signatures for each quorum
	for ind, quorum := range quorumParams {

		// Check that quorum has sufficient stake
		threshold := GetStakeThreshold(state.OperatorState, quorum.QuorumID, quorum.QuorumThreshold)
		if stakeSigned[ind].Cmp(threshold) == -1 {
			return nil, ErrInsufficientEthSigs
		}

		// Verify that the aggregated public key for the quorum matches the on-chain quorum aggregate public key sans
		// non-signers of the quorum
		quorumAggKey := state.AggKeys[quorum.QuorumID]
		quorumAggPubKeys[ind] = quorumAggKey

		signersAggKey := quorumAggKey.Deserialize(quorumAggKey.Serialize())
		for opInd, nsk := range nonSignerKeys {
			ops := state.Operators[quorum.QuorumID]
			if _, ok := ops[nonSignerOperatorIds[opInd]]; ok {
				signersAggKey.Sub(nsk)
			}
		}

		if aggPubKeys[ind] == nil {
			return nil, ErrAggSigNotValid
		}

		ok, err := signersAggKey.VerifyEquivalence(aggPubKeys[ind])
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, ErrPubKeysNotEqual
		}

		// Verify the aggregated signature for the quorum
		ok, err = aggSigs[ind].Verify(aggPubKeys[ind], message)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, ErrAggSigNotValid
		}

	}

	// Aggregate the aggregated signatures. We reuse the first aggregated signature as the accumulator
	for i := 1; i < len(aggSigs); i++ {
		aggSigs[0].Add(aggSigs[i].G1Point)
	}

	// Aggregate the aggregated public keys. We reuse the first aggregated public key as the accumulator
	for i := 1; i < len(aggPubKeys); i++ {
		aggPubKeys[0].Add(aggPubKeys[i])
	}

	return &SignatureAggregation{
		NonSigners:       nonSignerKeys,
		QuorumAggPubKeys: quorumAggPubKeys,
		AggPubKey:        aggPubKeys[0],
		AggSignature:     aggSigs[0],
	}, nil

}

func GetStakeThreshold(state *OperatorState, quorum QuorumID, quorumThreshold uint8) *big.Int {

	// Get stake threshold
	quorumThresholdBig := new(big.Int).SetUint64(uint64(quorumThreshold))
	stakeThreshold := new(big.Int)
	stakeThreshold.Mul(quorumThresholdBig, state.Totals[quorum].Stake)
	stakeThreshold = utils.RoundUpDivideBig(stakeThreshold, new(big.Int).SetUint64(PercentMultiplier))

	return stakeThreshold
}
