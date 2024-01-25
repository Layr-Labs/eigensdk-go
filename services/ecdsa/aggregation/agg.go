package aggregation

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/ecdsa/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	TaskAlreadyInitializedErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d already initialized", taskIndex)
	}
	TaskExpiredError    = fmt.Errorf("task expired")
	TaskNotFoundErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	OperatorNotPartOfTaskQuorumErrorFn = func(operatorId types.EcdsaOperatorId, taskIndex types.TaskIndex) error {
		return fmt.Errorf("operator %x not part of task %d's quorum", operatorId, taskIndex)
	}
	SignatureVerificationError = func(err error) error {
		return fmt.Errorf("Failed to verify signature: %w", err)
	}
	IncorrectSignatureError = errors.New("Signature verification failed. Incorrect Signature.")
)

// EcdsaAggregationServiceResponse is the response from the ecdsa aggregation service
type EcdsaAggregationServiceResponse struct {
	Err                       error
	TaskIndex                 types.TaskIndex
	TaskResponseDigest        types.TaskResponseDigest
	SignerIds                 []common.Address
	Signatures                [][]byte
	SignerQuorumBitmapIndices []uint32
	TotalStakeIndices         []uint32
	SignerStakeIndices        [][]uint32
}

// aggregatedOperators is meant to be used as a value in a map
// map[taskResponseDigest]aggregatedOperators
type aggregatedOperators struct {
	// address of all operators who signed on this taskResponseDigest
	signersIds []common.Address
	// signatures of all operators who signed on this taskResponseDigest
	signatures [][]byte
	// aggregate stake of all operators who signed on this header for each quorum
	signersTotalStakePerQuorum map[types.QuorumNum]*big.Int
}

// EcdsaAggregationService is the interface provided to avs aggregator code for doing ecdsa aggregation
// Currently its only implementation is the EcdsaAggregatorService, so see the comment there for more details
type EcdsaAggregationService interface {
	// InitializeNewTask should be called whenever a new task is created. ProcessNewSignature will return an error
	// if the task it is trying to process has not been initialized yet.
	// quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered complete, which happens
	// when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers whose stake
	// in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in that quorum
	InitializeNewTask(
		taskIndex types.TaskIndex,
		taskCreatedBlock uint32,
		quorumNumbers []types.QuorumNum,
		quorumThresholdPercentages []types.QuorumThresholdPercentage,
		timeToExpiry time.Duration,
	) error

	// ProcessNewSignature processes a new signature over a taskResponseDigest for a particular taskIndex by a particular operator
	// It verifies that the signature is correct and returns an error if it is not, and then aggregates the signature and stake of
	// the operator with all other signatures for the same taskIndex and taskResponseDigest pair.
	// Note: This function currently only verifies signatures over the taskResponseDigest directly, so avs code needs to verify that the digest
	// passed to ProcessNewSignature is indeed the digest of a valid taskResponse (that is, EcdsaAggregationService does not verify semantic integrity of the taskResponses)
	ProcessNewSignature(
		ctx context.Context,
		taskIndex types.TaskIndex,
		taskResponseDigest types.TaskResponseDigest,
		ecdsaSignature []byte,
		operatorId common.Address,
	) error

	// GetResponseChannel returns the single channel that meant to be used as the response channel
	// Any task that is completed (see the completion criterion in the comment above InitializeNewTask)
	// will be sent on this channel along with all the necessary information to call ECDSASignatureChecker onchain
	GetResponseChannel() <-chan EcdsaAggregationServiceResponse
}

// EcdsaAggregatorService is a service that performs Ecdsa signature aggregation for an AVS' tasks
// Assumptions:
//  1. EcdsaAggregatorService only verifies digest signatures, so avs code needs to verify that the digest
//     passed to ProcessNewSignature is indeed the digest of a valid taskResponse
//     (see the comment above checkSignature for more details)
//  2. EcdsaAggregatorService is VERY generic and makes very few assumptions about the tasks structure or
//     the time at which operators will send their signatures. It is mostly suitable for offchain computation
//     oracle (a la truebit) type of AVS, where tasks are sent onchain by users sporadically, and where
//     new tasks can start even before the previous ones have finished aggregation.
//     AVSs like eigenDA that have a much more controlled task submission schedule and where new tasks are
//     only submitted after the previous one's response has been aggregated and responded onchain, could have
//     a much simpler AggregationService without all the complicated parallel goroutines.
type EcdsaAggregatorService struct {
	// aggregatedResponsesC is the channel which all goroutines share to send their responses back to the
	// main thread after they are done aggregating (either they reached the threshold, or timeout expired)
	aggregatedResponsesC chan EcdsaAggregationServiceResponse
	// signedTaskRespsCs are the channels to send the signed task responses to the goroutines processing them
	// each new task is assigned a new goroutine and a new channel
	signedTaskRespsCs map[types.TaskIndex]chan types.SignedEcdsaTaskResponseDigest
	// we add chans to taskChans from the main thread (InitializeNewTask) when we create new tasks,
	// we read them in ProcessNewSignature from the main thread when we receive new signed tasks,
	// and remove them from its respective goroutine when the task is completed or reached timeout
	// we thus need a mutex to protect taskChans
	taskChansMutex     sync.RWMutex
	avsRegistryService avsregistry.AvsRegistryService
	logger             logging.Logger
}

var _ EcdsaAggregationService = (*EcdsaAggregatorService)(nil)

func NewEcdsaAggregatorService(avsRegistryService avsregistry.AvsRegistryService, logger logging.Logger) *EcdsaAggregatorService {
	return &EcdsaAggregatorService{
		aggregatedResponsesC: make(chan EcdsaAggregationServiceResponse),
		signedTaskRespsCs:    make(map[types.TaskIndex]chan types.SignedEcdsaTaskResponseDigest),
		taskChansMutex:       sync.RWMutex{},
		avsRegistryService:   avsRegistryService,
		logger:               logger,
	}
}

func (a *EcdsaAggregatorService) GetResponseChannel() <-chan EcdsaAggregationServiceResponse {
	return a.aggregatedResponsesC
}

// InitializeNewTask creates a new task goroutine meant to process new signed task responses for that task
// (that are sent via ProcessNewSignature) and adds a channel to a.taskChans to send the signed task responses to it
// quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered complete, which happens
// when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers whose stake
// in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in that quorum
func (a *EcdsaAggregatorService) InitializeNewTask(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
) error {
	if _, taskExists := a.signedTaskRespsCs[taskIndex]; taskExists {
		return TaskAlreadyInitializedErrorFn(taskIndex)
	}
	signedTaskRespsC := make(chan types.SignedEcdsaTaskResponseDigest)
	a.taskChansMutex.Lock()
	a.signedTaskRespsCs[taskIndex] = signedTaskRespsC
	a.taskChansMutex.Unlock()
	go a.singleTaskAggregatorGoroutineFunc(taskIndex, taskCreatedBlock, quorumNumbers, quorumThresholdPercentages, timeToExpiry, signedTaskRespsC)
	return nil
}

func (a *EcdsaAggregatorService) ProcessNewSignature(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskResponseDigest types.TaskResponseDigest,
	ecdsaSignature []byte,
	operatorId common.Address,
) error {
	a.taskChansMutex.Lock()
	taskC, taskInitialized := a.signedTaskRespsCs[taskIndex]
	a.taskChansMutex.Unlock()
	if !taskInitialized {
		return TaskNotFoundErrorFn(taskIndex)
	}
	signatureVerificationErrorC := make(chan error)
	// send the task to the goroutine processing this task
	// and return the error (if any) returned by the signature verification routine
	select {
	// we need to send this as part of select because if the goroutine is processing another SignedTaskResponseDigest
	// and cannot receive this one, we want the context to be able to cancel the request
	case taskC <- types.SignedEcdsaTaskResponseDigest{
		TaskResponseDigest:          taskResponseDigest,
		EcdsaSignature:              ecdsaSignature,
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

func (a *EcdsaAggregatorService) singleTaskAggregatorGoroutineFunc(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	timeToExpiry time.Duration,
	signedTaskRespsC <-chan types.SignedEcdsaTaskResponseDigest,
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

	// TODO(samlaf): instead of taking a TTE, we should take a block as input
	// and monitor the chain and only close the task goroutine when that block is reached
	taskExpiredTimer := time.NewTimer(timeToExpiry)

	aggregatedOperatorsDict := map[types.TaskResponseDigest]aggregatedOperators{}
	for {
		select {
		case signedTaskResponseDigest := <-signedTaskRespsC:
			a.logger.Debug("Task goroutine received new signed task response digest",
				"taskIndex", taskIndex, "TaskResponseDigest", signedTaskResponseDigest.TaskResponseDigest,
				"OperatorId", signedTaskResponseDigest.OperatorId, "ecdsaSignature", signedTaskResponseDigest.EcdsaSignature,
			)
			signedTaskResponseDigest.SignatureVerificationErrorC <- a.verifySignature(taskIndex, signedTaskResponseDigest, operatorsAvsStateDict)
			// after verifying signature we aggregate its sig and pubkey, and update the signed stake amount
			digestAggregatedOperators, ok := aggregatedOperatorsDict[signedTaskResponseDigest.TaskResponseDigest]
			if !ok {
				// first operator to sign on this digest
				digestAggregatedOperators = aggregatedOperators{
					// we've already verified that the operator is part of the task's quorum, so we don't need checks here
					signersIds:                 []common.Address{signedTaskResponseDigest.OperatorId},
					signatures:                 [][]byte{signedTaskResponseDigest.EcdsaSignature},
					signersTotalStakePerQuorum: operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum,
				}
			} else {
				digestAggregatedOperators.signersIds = append(digestAggregatedOperators.signersIds, signedTaskResponseDigest.OperatorId)
				digestAggregatedOperators.signatures = append(digestAggregatedOperators.signatures, signedTaskResponseDigest.EcdsaSignature)
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
				indices, err := a.avsRegistryService.GetCheckSignaturesIndices(&bind.CallOpts{}, taskCreatedBlock, quorumNumbers, digestAggregatedOperators.signersIds)
				if err != nil {
					a.logger.Error("Failed to get check signatures indices", "err", err)
					a.aggregatedResponsesC <- EcdsaAggregationServiceResponse{
						Err: err,
					}
					return
				}
				ecdsaAggregationServiceResponse := EcdsaAggregationServiceResponse{
					Err:                nil,
					TaskIndex:          taskIndex,
					TaskResponseDigest: signedTaskResponseDigest.TaskResponseDigest,
					// TODO(samlaf): I think we'll have to order these by operatorId to pass the require in the contract iirc
					SignerIds:                 digestAggregatedOperators.signersIds,
					Signatures:                digestAggregatedOperators.signatures,
					SignerQuorumBitmapIndices: indices.SignerQuorumBitmapIndices,
					TotalStakeIndices:         indices.TotalStakeIndices,
					SignerStakeIndices:        indices.SignerStakeIndices,
				}
				a.aggregatedResponsesC <- ecdsaAggregationServiceResponse
				return
			}
		case <-taskExpiredTimer.C:
			a.aggregatedResponsesC <- EcdsaAggregationServiceResponse{
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
func (a *EcdsaAggregatorService) closeTaskGoroutine(taskIndex types.TaskIndex) {
	a.taskChansMutex.Lock()
	delete(a.signedTaskRespsCs, taskIndex)
	a.taskChansMutex.Unlock()
}

// verifySignature verifies that a signature is valid against the operator pubkey stored in the
// operatorsAvsStateDict for that particular task
// TODO(samlaf): right now we are only checking that the *digest* is signed correctly!!
// we could be sent a signature of any kind of garbage and we would happily aggregate it
// this forces the avs code to verify that the digest is indeed the digest of a valid taskResponse
// we could take taskResponse as an interface{} and have avs code pass us a taskResponseHashFunction
// that we could use to hash and verify the taskResponse itself
func (a *EcdsaAggregatorService) verifySignature(
	taskIndex types.TaskIndex,
	signedTaskResponseDigest types.SignedEcdsaTaskResponseDigest,
	operatorsAvsStateDict map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId]
	if !ok {
		a.logger.Warnf("Operator %#v not found. Skipping message.", signedTaskResponseDigest.OperatorId)
		return OperatorNotPartOfTaskQuorumErrorFn(signedTaskResponseDigest.OperatorId, taskIndex)
	}

	// 0. verify that the msg actually came from the correct operator
	a.logger.Debug("Verifying signed task response digest signature",
		"taskResponseDigest", signedTaskResponseDigest.TaskResponseDigest,
		"operatorId", signedTaskResponseDigest.OperatorId,
		"ecdsaSignature", signedTaskResponseDigest.EcdsaSignature,
	)
	signatureVerified, err := ecdsa.VerifySignature(signedTaskResponseDigest.TaskResponseDigest[:], signedTaskResponseDigest.EcdsaSignature, signedTaskResponseDigest.OperatorId)
	if err != nil {
		a.logger.Error(SignatureVerificationError(err).Error())
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		a.logger.Error(IncorrectSignatureError.Error())
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
