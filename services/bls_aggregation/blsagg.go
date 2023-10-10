package blsagg

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
)

var (
	TaskAlreadyInitializedErrorF = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d already initialized", taskIndex)
	}
	TaskExpiredError   = fmt.Errorf("task expired")
	TaskNotFoundErrorF = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	OperatorNotPartOfTaskQuorumErrorF = func(operatorId types.OperatorId, taskIndex types.TaskIndex) error {
		return fmt.Errorf("operator %x not part of task %d's quorum", operatorId, taskIndex)
	}
	SignatureVerificationError = func(err error) error {
		return fmt.Errorf("Failed to verify signature: %w", err)
	}
	IncorrectSignatureError = errors.New("Signature verification failed. Incorrect Signature.")
)

// BlsAggregationServiceResponse is the response from the bls aggregation service
// it's half of the data needed to build the NonSignerStakesAndSignature struct
type BlsAggregationServiceResponse struct {
	Err                          error
	TaskIndex                    types.TaskIndex
	TaskResponseDigest           types.TaskResponseDigest
	NonSignersPubkeysG1          []*bls.G1Point
	QuorumApksG1                 []*bls.G1Point
	SignersApkG2                 *bls.G2Point
	SignersAggSigG1              *bls.Signature
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// aggregatedOperators is meant to be used as a value in a map
// map[taskResponseDigest]aggregatedOperators
type aggregatedOperators struct {
	// aggregate g2 pubkey of all operatos who signed on this taskResponseDigest
	signersApkG2 *bls.G2Point
	// aggregate signature of all operators who signed on this taskResponseDigest
	signersAggSigG1 *bls.Signature
	// aggregate stake of all operators who signed on this header for each quorum
	signersTotalStakePerQuorum map[types.QuorumNum]*big.Int
	// OperatorId of operators who signed on this header for each quorum
	signersOperatorIdsSet map[types.OperatorId]bool
}

type BlsAggregationService interface {
	InitializeNewTask(
		taskIndex types.TaskIndex,
		taskCreatedBlock uint32,
		quorumNumbers []types.QuorumNum,
		quorumThresholdPercentages []types.QuorumThresholdPercentage,
		timeToExpiry time.Duration,
	) error

	ProcessNewSignature(
		ctx context.Context,
		taskIndex types.TaskIndex,
		taskResponseDigest types.TaskResponseDigest,
		blsSignature *bls.Signature,
		operatorId bls.OperatorId,
	) error

	GetResponseChannel() <-chan BlsAggregationServiceResponse
}

// BlsAggregatorService is a service that performs BLS signature aggregation for an AVS' tasks
// Assumptions:
//  1. BlsAggregatorService only verifies digest signatures, so avs code needs to verify that the digest
//     passed to ProcessNewSignature is indeed the digest of a valid taskResponse
//     (see the comment above checkSignature for more details)
//  2. BlsAggregatorService is VERY generic and makes very few assumptions about the tasks structure or
//     the time at which operators will send their signatures. It is mostly suitable for offchain computation
//     oracle (a la truebit) type of AVS, where tasks are sent onchain by users sporadically, and where
//     new tasks can start even before the previous ones have finished aggregation.
//     AVSs like eigenDA that have a much more controlled task submission schedule and where new tasks are
//     only submitted after the previous one's response has been aggregated and responded onchain, could have
//     a much simpler AggregationService without all the complicated parallel goroutines.
type BlsAggregatorService struct {
	// ResponseC is the channel which all goroutines share to send their responses back to the
	// main thread after they are done aggregating (either they reached the threshold, or timeout expired)
	ResponseC chan BlsAggregationServiceResponse
	// taskChans are the channels to send the signed tasks to the goroutines processing them
	// each new task is assigned a new goroutine and a new channel
	taskChans map[types.TaskIndex]chan types.SignedTaskResponseDigest
	// we add chans to taskChans from the main thread (InitializeNewTask) when we create new tasks,
	// we read them in ProcessNewSignature from the main thread when we receive new signed tasks,
	// and remove them from its respective goroutine when the task is completed or reached timeout
	// we thus need a mutex to protect taskChans
	taskChansMutex     sync.RWMutex
	avsRegistryService avsregistry.AvsRegistryService
	logger             logging.Logger
}

var _ BlsAggregationService = (*BlsAggregatorService)(nil)

func NewBlsAggregatorService(avsRegistryService avsregistry.AvsRegistryService, logger logging.Logger) *BlsAggregatorService {
	return &BlsAggregatorService{
		ResponseC:          make(chan BlsAggregationServiceResponse),
		taskChans:          make(map[types.TaskIndex]chan types.SignedTaskResponseDigest),
		taskChansMutex:     sync.RWMutex{},
		avsRegistryService: avsRegistryService,
		logger:             logger,
	}
}

func (a *BlsAggregatorService) GetResponseChannel() <-chan BlsAggregationServiceResponse {
	return a.ResponseC
}

func (a *BlsAggregatorService) InitializeNewTask(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
) error {
	if _, taskExists := a.taskChans[taskIndex]; taskExists {
		return TaskAlreadyInitializedErrorF(taskIndex)
	}
	taskC := make(chan types.SignedTaskResponseDigest)
	a.taskChansMutex.Lock()
	a.taskChans[taskIndex] = taskC
	a.taskChansMutex.Unlock()
	go a.singleTaskAggregatorGoroutineFunc(taskIndex, taskCreatedBlock, quorumNumbers, quorumThresholdPercentages, timeToExpiry, taskC)
	return nil
}

func (a *BlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskResponseDigest types.TaskResponseDigest,
	blsSignature *bls.Signature,
	operatorId bls.OperatorId,
) error {
	a.taskChansMutex.Lock()
	taskC, taskInitialized := a.taskChans[taskIndex]
	a.taskChansMutex.Unlock()
	if !taskInitialized {
		return TaskNotFoundErrorF(taskIndex)
	}
	signatureVerificationErrorC := make(chan error)
	// send the task to the goroutine processing this task
	// and return the error (if any) returned by the signature verification routine
	select {
	case taskC <- types.SignedTaskResponseDigest{
		TaskResponseDigest:          taskResponseDigest,
		BlsSignature:                blsSignature,
		OperatorId:                  operatorId,
		SignatureVerificationErrorC: signatureVerificationErrorC,
	}:
		// note that we need to wait synchronously here for this response because we want to
		// send back an informative error message to the operator who sent his signature to the aggregator
		return <-signatureVerificationErrorC
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *BlsAggregatorService) singleTaskAggregatorGoroutineFunc(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	taskRespC <-chan types.SignedTaskResponseDigest,
) {
	defer a.closeTaskGoroutine(taskIndex)

	quorumThresholdPercentagesMap := make(map[types.QuorumNum]types.QuorumThresholdPercentage)
	for i, quorumNumber := range quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = quorumThresholdPercentages[i]
	}
	operatorsAvsStateDict, err := a.avsRegistryService.GetOperatorsAvsStateAtBlock(context.Background(), quorumNumbers, taskCreatedBlock)
	if err != nil {
		// TODO: how should we handle such an error?
		a.logger.Fatal("Aggregator failed to get operators state from avs registry", "err", err)
	}
	quorumsAvsStakeDict, err := a.avsRegistryService.GetQuorumsAvsStateAtBlock(context.Background(), quorumNumbers, taskCreatedBlock)
	if err != nil {
		a.logger.Fatal("Aggregator failed to get quorums state from avs registry", "err", err)
	}
	totalStakePerQuorum := make(map[types.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
	}
	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	// TODO(samlaf): instead of taking a TTE, we should take a block as input
	// and monitor the chain and only close the task goroutine when that block is reached
	taskExpiredTimer := time.NewTimer(timeToExpiry)

	aggregatedOperatorsDict := map[types.TaskResponseDigest]aggregatedOperators{}
	for {
		select {
		case signedTaskResponseDigest := <-taskRespC:
			signedTaskResponseDigest.SignatureVerificationErrorC <- a.verifySignature(taskIndex, signedTaskResponseDigest, operatorsAvsStateDict)
			// after verifying signature we aggregate its sig and pubkey, and update the signed stake amount
			digestAggregatedOperators, ok := aggregatedOperatorsDict[signedTaskResponseDigest.TaskResponseDigest]
			if !ok {
				// first operator to sign on this digest
				digestAggregatedOperators = aggregatedOperators{
					// we've already verified that the operator is part of the task's quorum, so we don't need checks here
					signersApkG2:               operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].Pubkeys.G2Pubkey,
					signersAggSigG1:            signedTaskResponseDigest.BlsSignature,
					signersOperatorIdsSet:      map[types.OperatorId]bool{signedTaskResponseDigest.OperatorId: true},
					signersTotalStakePerQuorum: operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum,
				}
			} else {
				digestAggregatedOperators.signersAggSigG1.Add(signedTaskResponseDigest.BlsSignature)
				digestAggregatedOperators.signersApkG2.Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].Pubkeys.G2Pubkey)
				digestAggregatedOperators.signersOperatorIdsSet[signedTaskResponseDigest.OperatorId] = true
				for quorumNum, stake := range operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum {
					if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
						// if we haven't seen this quorum before, initialize its signed stake to 0
						// possible if previous operators who sent us signatures were not part of this quorum
						digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
					}
					digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
				}
			}
			// update the aggregatedOperatorsDict. Note that we need to assign the whole struct value at once,
			// because of https://github.com/golang/go/issues/3117
			aggregatedOperatorsDict[signedTaskResponseDigest.TaskResponseDigest] = digestAggregatedOperators

			if checkIfStakeThresholdsMet(digestAggregatedOperators.signersTotalStakePerQuorum, totalStakePerQuorum, quorumThresholdPercentagesMap) {
				nonSignersOperatorIds := []types.OperatorId{}
				for operatorId := range operatorsAvsStateDict {
					if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
						nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
					}
				}
				indices, err := a.avsRegistryService.GetCheckSignaturesIndices(context.Background(), taskCreatedBlock, quorumNumbers, nonSignersOperatorIds)
				if err != nil {
					a.logger.Error("Failed to get check signatures indices", "err", err)
					a.ResponseC <- BlsAggregationServiceResponse{
						Err: err,
					}
					return
				}
				blsAggregationServiceResponse := BlsAggregationServiceResponse{
					Err:                          nil,
					TaskIndex:                    taskIndex,
					TaskResponseDigest:           signedTaskResponseDigest.TaskResponseDigest,
					NonSignersPubkeysG1:          getG1PubkeysOfNonSigners(digestAggregatedOperators.signersOperatorIdsSet, operatorsAvsStateDict),
					QuorumApksG1:                 quorumApksG1,
					SignersApkG2:                 digestAggregatedOperators.signersApkG2,
					SignersAggSigG1:              digestAggregatedOperators.signersAggSigG1,
					NonSignerQuorumBitmapIndices: indices.NonSignerQuorumBitmapIndices,
					QuorumApkIndices:             indices.QuorumApkIndices,
					TotalStakeIndices:            indices.TotalStakeIndices,
					NonSignerStakeIndices:        indices.NonSignerStakeIndices,
				}
				a.ResponseC <- blsAggregationServiceResponse
				return
			}
		case <-taskExpiredTimer.C:
			a.ResponseC <- BlsAggregationServiceResponse{
				Err: TaskExpiredError,
			}
			return
		}
	}

}

// closeTaskGoroutine is run when the goroutine processing taskIndex's task responses ends (for whatever reason)
// it deletes the response channel for taskIndex from a.taskChans
// so that the main thread knows that this task goroutine is no longer running
// and doesn't try to send new signatures to it
func (a *BlsAggregatorService) closeTaskGoroutine(taskIndex types.TaskIndex) {
	a.taskChansMutex.Lock()
	delete(a.taskChans, taskIndex)
	a.taskChansMutex.Unlock()
}

// verifySignature verifies that a signature is valid against the operator pubkey stored in the
// operatorsAvsStateDict for that particular task
// TODO(samlaf): right now we are only checking that the *digest* is signed correctly!!
// we could be sent a signature of any kind of garbage and we would happily aggregate it
// this forces the avs code to verify that the digest is indeed the digest of a valid taskResponse
// we could take taskResponse as an interface{} and have avs code pass us a taskResponseHashFunction
// that we could use to hash and verify the taskResponse itself
func (a *BlsAggregatorService) verifySignature(
	taskIndex types.TaskIndex,
	signedTaskResponseDigest types.SignedTaskResponseDigest,
	operatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId]
	if !ok {
		a.logger.Warnf("Operator %#v not found. Skipping message.", signedTaskResponseDigest.OperatorId)
		return OperatorNotPartOfTaskQuorumErrorF(signedTaskResponseDigest.OperatorId, taskIndex)
	}

	// 0. verify that the msg actually came from the correct operator
	operatorG2Pubkey := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		a.logger.Fatalf("Operator G2 pubkey not found")
	}
	a.logger.Debug("Verifying signed task response digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"taskResponseDigest", signedTaskResponseDigest.TaskResponseDigest,
		"blsSignature", signedTaskResponseDigest.BlsSignature,
	)
	signatureVerified, err := signedTaskResponseDigest.BlsSignature.Verify(operatorG2Pubkey, signedTaskResponseDigest.TaskResponseDigest)
	if err != nil {
		a.logger.Errorf(SignatureVerificationError(err).Error())
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		a.logger.Errorf(IncorrectSignatureError.Error())
		return IncorrectSignatureError
	}
	return nil
}

// checkIfStakeThresholdsMet checks at least quorumThresholdPercentage of stake
// has signed for each quorum.
func checkIfStakeThresholdsMet(
	signedStakePerQuorum map[types.QuorumNum]*big.Int,
	totalStakePerQuorum map[types.QuorumNum]*big.Int,
	quorumThresholdPercentagesMap map[types.QuorumNum]types.QuorumThresholdPercentage,
) bool {
	for quorumNum, quorumThresholdPercentage := range quorumThresholdPercentagesMap {
		// we check that signedStake >= totalStake * quorumThresholdPercentage / 100
		// to be exact (and do like the contracts), we actually check that
		// signedStake * 100 >= totalStake * quorumThresholdPercentage
		signedStake := big.NewInt(0).Mul(signedStakePerQuorum[quorumNum], big.NewInt(100))
		thresholdStake := big.NewInt(0).Mul(totalStakePerQuorum[quorumNum], big.NewInt(int64(quorumThresholdPercentage)))
		if signedStake.Cmp(thresholdStake) < 0 {
			return false
		}
	}
	return true
}

func getG1PubkeysOfNonSigners(signersOperatorIdsSet map[types.OperatorId]bool, operatorAvsStateDict map[[32]byte]types.OperatorAvsState) []*bls.G1Point {
	nonSignersG1Pubkeys := []*bls.G1Point{}
	for operatorId, operator := range operatorAvsStateDict {
		if _, operatorSigned := signersOperatorIdsSet[operatorId]; !operatorSigned {
			nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.Pubkeys.G1Pubkey)
		}
	}
	return nonSignersG1Pubkeys
}
