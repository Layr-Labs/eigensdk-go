package blsagg

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/Layr-Labs/eigensdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	// TODO: refactor these errors to use a custom struct with taskIndex field instead of wrapping taskIndex in the
	// error string directly.
	//       see https://go.dev/blog/go1.13-errors
	TaskInitializationErrorFn = func(err error, taskIndex types.TaskIndex) error {
		text := fmt.Sprintf("failed to initialize task %d", taskIndex)
		return utils.WrapError(text, err)
	}
	TaskAlreadyInitializedErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d already initialized", taskIndex)
	}
	TaskExpiredErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d expired", taskIndex)
	}
	TaskNotFoundErrorFn = func(taskIndex types.TaskIndex) error {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	OperatorNotPartOfTaskQuorumErrorFn = func(operatorId types.OperatorId, taskIndex types.TaskIndex) error {
		return fmt.Errorf("operator %x not part of task %d's quorum", operatorId, taskIndex)
	}
	HashFunctionError = func(err error) error {
		return utils.WrapError("failed to hash task response", err)
	}
	SignatureVerificationError = func(err error) error {
		return utils.WrapError("failed to verify signature", err)
	}
	ErrIncorrectSignature = errors.New("signature verification failed. incorrect signature")
)

// BlsAggregationServiceResponse is the response from the bls aggregation service
type BlsAggregationServiceResponse struct {
	Err                error                    // if Err is not nil, the other fields are not valid
	TaskIndex          types.TaskIndex          // unique identifier of the task
	TaskResponse       types.TaskResponse       // the task response that was signed
	TaskResponseDigest types.TaskResponseDigest // digest of the task response that was signed
	// The below 8 fields are the data needed to build the IBLSSignatureChecker.NonSignerStakesAndSignature struct
	// users of this service will need to build the struct themselves by converting the bls points
	// into the BN254.G1/G2Point structs that the IBLSSignatureChecker expects
	// given that those are different for each AVS service manager that individually inherits BLSSignatureChecker
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
	// set of OperatorId of operators who signed on this header
	signersOperatorIdsSet map[types.OperatorId]bool
}

// TaskSignature contains the data required to process and verify a new signature for a task response.
type TaskSignature struct {
	// unique identifier of the task associated with this signature
	taskIndex types.TaskIndex
	// response data that has been signed
	taskResponse types.TaskResponse
	// BLS cryptographic signature for the task response
	blsSignature *bls.Signature
	// id of the operator who signed the task response
	operatorId types.OperatorId
}

// NewTaskSignature creates a new instance of TaskSignature
func NewTaskSignature(
	taskIndex types.TaskIndex,
	taskResponse types.TaskResponse,
	blsSignature *bls.Signature,
	operatorId types.OperatorId,
) TaskSignature {
	return TaskSignature{
		taskIndex:    taskIndex,
		taskResponse: taskResponse,
		blsSignature: blsSignature,
		operatorId:   operatorId,
	}
}

// TaskMetadata encapsulates the necessary parameters to initialize a task
type TaskMetadata struct {
	// index of the task
	taskIndex types.TaskIndex
	// block the task was created
	taskCreatedBlock uint32
	// quorum numbers which should respond to the task
	quorumNumbers types.QuorumNums
	// threshold percentages required for each quorum
	quorumThresholdPercentages types.QuorumThresholdPercentages
	// time before expiry of the task response aggregation
	timeToExpiry time.Duration
	// time window to collect signatures after reaching quorum, defaults to 0 if not specified
	windowDuration time.Duration
}

func NewTaskMetadata(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers types.QuorumNums,
	quorumThresholdPercentages types.QuorumThresholdPercentages,
	timeToExpiry time.Duration,
) TaskMetadata {
	return TaskMetadata{
		taskIndex:                  taskIndex,
		taskCreatedBlock:           taskCreatedBlock,
		quorumNumbers:              quorumNumbers,
		quorumThresholdPercentages: quorumThresholdPercentages,
		timeToExpiry:               timeToExpiry,
		windowDuration:             0, // Default to 0
	}
}

func (t TaskMetadata) WithWindowDuration(windowDuration time.Duration) TaskMetadata {
	t.windowDuration = windowDuration
	return t
}

// BlsAggregationService is the interface provided to avs aggregator code for doing bls aggregation
// Currently its only implementation is the BlsAggregatorService, so see the comment there for more details
type BlsAggregationService interface {
	// InitializeNewTask creates a new task goroutine meant to process new signed task responses for that task
	// (that are sent via ProcessNewSignature) and adds a channel to a.taskChans to send the signed task responses to
	// it. The quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered
	// complete, which
	// happens when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers
	// whose stake in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in
	// that quorum.
	// Once the quorum is reached, the task is still open for a window of `windowDuration` time (default 0) to receive
	// more signatures, before sending the aggregation response through the aggregatedResponsesC channel.
	// If the task expiration is reached before the window finishes, the task response will still be sent to the
	// aggregatedResponsesC channel.
	InitializeNewTask(
		metadata TaskMetadata,
	) error

	// ProcessNewSignature processes a new signature over a taskResponseDigest for a particular taskIndex by a
	// particular operator It verifies that the signature is correct and returns an error if it is not, and then
	// aggregates the signature and stake of
	// the operator with all other signatures for the same taskIndex and taskResponseDigest pair.
	// Note: This function currently only verifies signatures over the taskResponseDigest directly, so avs code needs to
	// verify that the digest passed to ProcessNewSignature is indeed the digest of a valid taskResponse (that is,
	// BlsAggregationService does not verify semantic integrity of the taskResponses)
	ProcessNewSignature(
		ctx context.Context,
		task TaskSignature,
	) error

	// GetResponseChannel returns the single channel that meant to be used as the response channel
	// Any task that is completed (see the completion criterion in the comment above InitializeNewTask)
	// will be sent on this channel along with all the necessary information to call BLSSignatureChecker onchain
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
	handler  ServiceHandler
	receiver <-chan BlsAggregationServiceResponse
}

var _ BlsAggregationService = (*BlsAggregatorService)(nil)

func (a *BlsAggregatorService) InitializeNewTask(
	metadata TaskMetadata,
) error {
	return a.handler.InitializeNewTask(metadata)
}

func (a *BlsAggregatorService) ProcessNewSignature(
	ctx context.Context,
	task TaskSignature,
) error {
	return a.handler.ProcessNewSignature(ctx, task)
}

func (a *BlsAggregatorService) GetResponseChannel() <-chan BlsAggregationServiceResponse {
	return a.receiver
}

type BlsAggregatorBuilder struct {
	avsRegistryService avsregistry.AvsRegistryService
	logger             logging.Logger

	hashFunction types.TaskResponseHashFunction
}

func NewBlsAggregatorBuilder(
	avsRegistryService avsregistry.AvsRegistryService,
	hashFunction types.TaskResponseHashFunction,
	logger logging.Logger,
) *BlsAggregatorBuilder {
	return &BlsAggregatorBuilder{
		avsRegistryService: avsRegistryService,
		logger:             logger,
		hashFunction:       hashFunction,
	}
}

// NewBlsAggregatorService creates a new BlsAggregatorService
// avsRegistryService is the AVS registry service to use
// hashFunction is the hash function to use to compute the taskResponseDigest from the taskResponse
// logger is the logger to use
//
// An example of hashFunction is the one defined in blsagg_test.go:
// ```go
//
//	hashFunction := func(taskResponse types.TaskResponse) (types.TaskResponseDigest, error) {
//		taskResponseBytes, err := json.Marshal(taskResponse)
//		if err != nil {
//			return types.TaskResponseDigest{}, err
//		}
//		return types.TaskResponseDigest(sha256.Sum256(taskResponseBytes)), nil
//	}
//
// ```
func NewBlsAggregatorService(
	avsRegistryService avsregistry.AvsRegistryService,
	hashFunction types.TaskResponseHashFunction,
	logger logging.Logger,
) *BlsAggregatorService {
	// Instantiate builder and save handler and receiver.
	builder := NewBlsAggregatorBuilder(avsRegistryService, hashFunction, logger)

	handler, receiver := builder.Start()

	return &BlsAggregatorService{
		handler:  handler,
		receiver: receiver,
	}
}

// The service handler is a structure used to use the service without the complexity of it.
type ServiceHandler struct {
	//This channels are used to send messages (requests) to the service.
	taskInitC         chan<- initializeTaskRequest
	processSignatureC chan<- processSignatureRequest
}

type initializeTaskRequest struct {
	metadata TaskMetadata
	errC     chan<- error
}

type processSignatureRequest struct {
	metadata TaskSignature
	errC     chan<- error
}

// This function starts the service thread, initializing the aggregateResponses, initializeTask and processSignature
// channels, passing them to the run method (where the main loop is executed) and returns the service handler an the
// aggregate receiver channel to interact with the service thread.
func (a *BlsAggregatorBuilder) Start() (ServiceHandler, <-chan BlsAggregationServiceResponse) {
	// Create channels to handle requests
	initializeTaskC := make(chan initializeTaskRequest)
	processSignatureC := make(chan processSignatureRequest)

	aggResponsesC := make(chan BlsAggregationServiceResponse)

	go func() {
		a.run(initializeTaskC, processSignatureC, aggResponsesC)
		close(aggResponsesC)
	}()

	// Note: here we return the channel instead of a type with a receive function (different than rust
	// implementation) because for users is more useful to use select on the channel directly instead
	// of running a thread on a goroutine and using other channel
	return ServiceHandler{initializeTaskC, processSignatureC}, aggResponsesC
}

// Here is executed the main loop, where the requests are received from the initialize task and process signature
// channels, executing the corresponding logic in each case. The aggregated responses channel is passed to the
// single task aggregator goroutine, where the initialization of the task is done an notified.
// The loop ends if one of the request channels is closed.
func (a *BlsAggregatorBuilder) run(
	initializeTaskChannel <-chan initializeTaskRequest,
	processSignatureChannel <-chan processSignatureRequest,
	aggResponsesC chan<- BlsAggregationServiceResponse,
) {
	taskChannels := make(map[types.TaskIndex]chan types.SignedTaskResponseDigest)

	for {
		select {
		case taskInitReq, ok := <-initializeTaskChannel:
			if !ok {
				return
			}
			taskIndex := taskInitReq.metadata.taskIndex

			_, taskExists := taskChannels[taskIndex]
			if taskExists {
				taskInitReq.errC <- TaskAlreadyInitializedErrorFn(taskInitReq.metadata.taskIndex)
				continue
			}
			signedTaskRespC := make(chan types.SignedTaskResponseDigest)
			taskChannels[taskIndex] = signedTaskRespC

			go a.singleTaskAggregatorGoroutineFunc(
				taskInitReq.metadata,
				signedTaskRespC,
				aggResponsesC,
			)
			taskInitReq.errC <- nil

		case signatureReq, ok := <-processSignatureChannel:
			if !ok {
				return
			}
			taskIndex := signatureReq.metadata.taskIndex

			signedTaskRespsC, taskExists := taskChannels[taskIndex]
			if taskExists {
				signedDigest := types.SignedTaskResponseDigest{
					TaskResponse:                signatureReq.metadata.taskResponse,
					BlsSignature:                signatureReq.metadata.blsSignature,
					OperatorId:                  signatureReq.metadata.operatorId,
					SignatureVerificationErrorC: signatureReq.errC,
				}

				signedTaskRespsC <- signedDigest
			} else {
				signatureReq.errC <- TaskNotFoundErrorFn(taskIndex)
			}
		}
	}
}

// InitializeNewTask sends to the service thread a request to process new signed task responses for that task
// (that are sent via ProcessNewSignature).
//
// The metadata parameter contains:
//   - taskIndex: Unique identifier for the task
//   - taskCreatedBlock: Block number at which the task was created
//   - quorumNumbers: 	Quorum numbers which should respond to the task
//   - quorumThresholdPercentages: Threshold percentage required per quorum
//   - timeToExpiry: Time before expiry of the task response aggregation
//   - windowDuration: Additional time window to collect signatures after reaching quorum (default 0)
//
// The quorumNumbers and quorumThresholdPercentages set the requirements for this task to be considered complete, which
// happens when a particular TaskResponseDigest (received via the a.taskChans[taskIndex]) has been signed by signers
// whose stake in each of the listed quorums adds up to at least quorumThresholdPercentages[i] of the total stake in
// that quorum.
// Once the quorum is reached, the task is still open for a window of `windowDuration` time (default 0) to receive more
// signatures, before sending the aggregation response through the aggregatedResponsesC channel.
// If the task expiration is reached before the window finishes, the task response will still be sent to the
// aggregatedResponsesC channel.
func (a *ServiceHandler) InitializeNewTask(
	metadata TaskMetadata,
) error {
	errChan := make(chan error)
	a.taskInitC <- initializeTaskRequest{metadata, errChan}
	err := <-errChan
	return err
}

// ProcessNewSignature sends to the service thread a request to process a new signature for a previous initialize task.
//
// The metadata parameter contains:
//   - taskIndex: Unique identifier for the task
//   - taskResponse: Response data that has been signed
//   - blsSignature: BLS cryptographic signature for the task response
//   - operatorId: id of the operator who signed the task response
//
// The last three attributes are used to make the response digest.
func (a *ServiceHandler) ProcessNewSignature(
	ctx context.Context,
	metadata TaskSignature,
) error {
	errChan := make(chan error)
	a.processSignatureC <- processSignatureRequest{metadata, errChan}

	// Doing this we let the goroutine consume the result of the operation, but allow
	// the operation to end early if the context is cancelled
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a *BlsAggregatorBuilder) singleTaskAggregatorGoroutineFunc(
	metadata TaskMetadata,
	signedTaskRespsC <-chan types.SignedTaskResponseDigest,
	aggregatedResponsesC chan<- BlsAggregationServiceResponse,
) {
	a.logger.Debug("AggregatorService goroutine processing new task",
		"taskIndex", metadata.taskIndex,
		"taskCreatedBlock", metadata.taskCreatedBlock)

	quorumThresholdPercentagesMap := make(map[types.QuorumNum]types.QuorumThresholdPercentage)
	for i, quorumNumber := range metadata.quorumNumbers {
		quorumThresholdPercentagesMap[quorumNumber] = metadata.quorumThresholdPercentages[i]
		a.logger.Debug("AggregatorService goroutine quorum threshold percentage",
			"taskIndex", metadata.taskIndex,
			"quorumNumber", quorumNumber,
			"quorumThresholdPercentage", metadata.quorumThresholdPercentages[i])
	}
	operatorsAvsStateDict, err := a.avsRegistryService.GetOperatorsAvsStateAtBlock(
		context.Background(),
		metadata.quorumNumbers,
		metadata.taskCreatedBlock,
	)
	if err != nil {
		a.logger.Error(
			"Task goroutine failed to get operators state from avs registry",
			"taskIndex",
			metadata.taskIndex,
			"err",
			err,
		)
		text := fmt.Sprintf(
			"AggregatorService failed to get operators state from avs registry at blockNum %d",
			metadata.taskCreatedBlock,
		)
		aggregatedResponsesC <- BlsAggregationServiceResponse{
			Err:       TaskInitializationErrorFn(utils.WrapError(text, err), metadata.taskIndex),
			TaskIndex: metadata.taskIndex,
		}
		return
	}
	quorumsAvsStakeDict, err := a.avsRegistryService.GetQuorumsAvsStateAtBlock(
		context.Background(),
		metadata.quorumNumbers,
		metadata.taskCreatedBlock,
	)
	if err != nil {
		a.logger.Error(
			"Task goroutine failed to get quorums state from avs registry",
			"taskIndex",
			metadata.taskIndex,
			"err",
			err,
		)
		aggregatedResponsesC <- BlsAggregationServiceResponse{
			Err:       TaskInitializationErrorFn(utils.WrapError("aggregator failed to get quorums state from avs registry", err), metadata.taskIndex),
			TaskIndex: metadata.taskIndex,
		}
		return
	}
	totalStakePerQuorum := make(map[types.QuorumNum]*big.Int)
	for quorumNum, quorumAvsState := range quorumsAvsStakeDict {
		totalStakePerQuorum[quorumNum] = quorumAvsState.TotalStake
		a.logger.Debug("Task goroutine quorum total stake",
			"taskIndex", metadata.taskIndex,
			"quorumNum", quorumNum,
			"totalStake", quorumAvsState.TotalStake)
	}
	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range metadata.quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}

	// TODO(samlaf): instead of taking a TTE, we should take a block as input
	// and monitor the chain and only close the task goroutine when that block is reached
	taskExpiredTimer := time.NewTimer(metadata.timeToExpiry)

	aggregatedOperatorsDict := map[types.TaskResponseDigest]aggregatedOperators{}
	// The windowTimer is initialized to be longer than the taskExpiredTimer as it will
	// be overwritten once the stake threshold is met
	windowTimer := time.NewTimer(metadata.timeToExpiry + 1*time.Second)
	openWindow := false
	var lastSignedTaskResponseDigest types.SignedTaskResponseDigest
	var lastDigestAggregatedOperators aggregatedOperators
	var lastTaskResponseDigest types.TaskResponseDigest
	for {
		select {
		case signedTaskResponseDigest := <-signedTaskRespsC:
			a.logger.Debug(
				"Task goroutine received new signed task response digest",
				"taskIndex",
				metadata.taskIndex,
				"signedTaskResponseDigest",
				signedTaskResponseDigest,
			)

			// compute the taskResponseDigest using the hash function
			taskResponseDigest, err := a.hashFunction(signedTaskResponseDigest.TaskResponse)
			if err != nil {
				// this error should never happen, because we've already hashed the taskResponse in verifySignature,
				// but keeping here in case the verifySignature implementation ever changes or some catastrophic bug
				// happens..
				continue
			}

			// check if the operator has already signed for this digest
			digestAggregatedOperators, ok := aggregatedOperatorsDict[taskResponseDigest]
			if ok {
				if digestAggregatedOperators.signersOperatorIdsSet[signedTaskResponseDigest.OperatorId] {
					a.logger.Info(
						"Duplicate signature received",
						"operatorId", fmt.Sprintf("%x", signedTaskResponseDigest.OperatorId),
						"taskIndex", metadata.taskIndex,
					)
					signedTaskResponseDigest.SignatureVerificationErrorC <- fmt.Errorf("duplicate signature from operator %x for task %d", signedTaskResponseDigest.OperatorId, metadata.taskIndex)
					continue
				}
			}

			err = a.verifySignature(metadata.taskIndex, signedTaskResponseDigest, operatorsAvsStateDict)
			// return the err (or nil) to the operator, and then proceed to do aggregation logic asynchronously (when no
			// error)
			signedTaskResponseDigest.SignatureVerificationErrorC <- err
			if err != nil {
				continue
			}

			// after verifying signature we aggregate its sig and pubkey, and update the signed stake amount
			if !ok {
				// first operator to sign on this digest
				signersApkG2 := bls.NewZeroG2Point()
				signersAggSigG1 := bls.NewZeroSignature()
				// for each quorum the operator has stake in, the signature is aggregated
				// see
				// https://github.com/Layr-Labs/eigenlayer-middleware/blob/7d49b5181b09198ed275783453aa082bb3766990/src/BLSSignatureChecker.sol#L161-L168
				for range operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum {
					signersApkG2 = signersApkG2.Add(
						operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey,
					)
					signersAggSigG1 = signersAggSigG1.Add(signedTaskResponseDigest.BlsSignature)
				}
				digestAggregatedOperators = aggregatedOperators{
					// we've already verified that the operator is part of the task's quorum, so we don't need checks
					// here
					signersApkG2:          signersApkG2,
					signersAggSigG1:       signersAggSigG1,
					signersOperatorIdsSet: map[types.OperatorId]bool{signedTaskResponseDigest.OperatorId: true},
					signersTotalStakePerQuorum: cloneStakePerQuorumMap(
						operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum,
					),
				}
			} else {
				a.logger.Debug("Task goroutine updating existing aggregated operator signatures",
					"taskIndex", metadata.taskIndex,
					"taskResponseDigest", taskResponseDigest)

				digestAggregatedOperators.signersOperatorIdsSet[signedTaskResponseDigest.OperatorId] = true
				// for each quorum the operator has stake in, the signature is aggregated
				for quorumNum, stake := range operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum {
					digestAggregatedOperators.signersAggSigG1.Add(signedTaskResponseDigest.BlsSignature)
					digestAggregatedOperators.signersApkG2.Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey)
					if _, ok := digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum]; !ok {
						// if we haven't seen this quorum before, initialize its signed stake to 0
						// possible if previous operators who sent us signatures were not part of this quorum
						digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
					}
					digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum].Add(digestAggregatedOperators.signersTotalStakePerQuorum[quorumNum], stake)
				}
			}

			// update the buffer variables to be used when the window timer fires
			lastDigestAggregatedOperators = digestAggregatedOperators
			lastTaskResponseDigest = taskResponseDigest
			lastSignedTaskResponseDigest = signedTaskResponseDigest

			// update the aggregatedOperatorsDict. Note that we need to assign the whole struct value at once,
			// because of https://github.com/golang/go/issues/3117
			aggregatedOperatorsDict[taskResponseDigest] = digestAggregatedOperators

			if !openWindow && checkIfStakeThresholdsMet(
				a.logger,
				digestAggregatedOperators.signersTotalStakePerQuorum,
				totalStakePerQuorum,
				quorumThresholdPercentagesMap,
			) {
				a.logger.Debug("Task goroutine stake threshold reached",
					"taskIndex", metadata.taskIndex,
					"taskResponseDigest", taskResponseDigest)

				openWindow = true
				windowTimer = time.NewTimer(metadata.windowDuration)
				a.logger.Debug("Window timer started")
			}
		case <-taskExpiredTimer.C:
			if openWindow {
				a.sendAggregatedResponse(
					operatorsAvsStateDict,
					metadata.taskIndex,
					metadata.taskCreatedBlock,
					lastSignedTaskResponseDigest,
					lastDigestAggregatedOperators,
					metadata.quorumNumbers,
					lastTaskResponseDigest,
					quorumApksG1,
					aggregatedResponsesC,
				)
			}

			aggregatedResponsesC <- BlsAggregationServiceResponse{
				Err:       TaskExpiredErrorFn(metadata.taskIndex),
				TaskIndex: metadata.taskIndex,
			}
			return
		case <-windowTimer.C:
			a.logger.Debug("Window timer expired")
			a.sendAggregatedResponse(
				operatorsAvsStateDict,
				metadata.taskIndex,
				metadata.taskCreatedBlock,
				lastSignedTaskResponseDigest,
				lastDigestAggregatedOperators,
				metadata.quorumNumbers,
				lastTaskResponseDigest,
				quorumApksG1,
				aggregatedResponsesC,
			)
			return
		}
	}
}

func (a *BlsAggregatorBuilder) sendAggregatedResponse(
	operatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState,
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	signedTaskResponseDigest types.SignedTaskResponseDigest,
	digestAggregatedOperators aggregatedOperators,
	quorumNumbers types.QuorumNums,
	taskResponseDigest types.TaskResponseDigest,
	quorumApksG1 []*bls.G1Point,
	aggregatedResponsesC chan<- BlsAggregationServiceResponse,
) {
	nonSignersOperatorIds := []types.OperatorId{}
	for operatorId := range operatorsAvsStateDict {
		if _, operatorSigned := digestAggregatedOperators.signersOperatorIdsSet[operatorId]; !operatorSigned {
			nonSignersOperatorIds = append(nonSignersOperatorIds, operatorId)
		}
	}

	// the contract requires a sorted nonSignersOperatorIds
	sort.SliceStable(nonSignersOperatorIds, func(i, j int) bool {
		iOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[i][:])
		jOprInt := new(big.Int).SetBytes(nonSignersOperatorIds[j][:])
		return iOprInt.Cmp(jOprInt) == -1
	})

	nonSignersG1Pubkeys := []*bls.G1Point{}
	for _, operatorId := range nonSignersOperatorIds {
		operator := operatorsAvsStateDict[operatorId]
		nonSignersG1Pubkeys = append(nonSignersG1Pubkeys, operator.OperatorInfo.Pubkeys.G1Pubkey)
	}

	indices, err := a.avsRegistryService.GetCheckSignaturesIndices(
		&bind.CallOpts{},
		taskCreatedBlock,
		quorumNumbers,
		nonSignersOperatorIds,
	)
	if err != nil {
		aggregatedResponsesC <- BlsAggregationServiceResponse{
			Err:       utils.WrapError(errors.New("failed to get check signatures indices"), err),
			TaskIndex: taskIndex,
		}
		return
	}

	blsAggregationServiceResponse := BlsAggregationServiceResponse{
		Err:                          nil,
		TaskIndex:                    taskIndex,
		TaskResponse:                 signedTaskResponseDigest.TaskResponse,
		TaskResponseDigest:           taskResponseDigest,
		NonSignersPubkeysG1:          nonSignersG1Pubkeys,
		QuorumApksG1:                 quorumApksG1,
		SignersApkG2:                 digestAggregatedOperators.signersApkG2,
		SignersAggSigG1:              digestAggregatedOperators.signersAggSigG1,
		NonSignerQuorumBitmapIndices: indices.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             indices.QuorumApkIndices,
		TotalStakeIndices:            indices.TotalStakeIndices,
		NonSignerStakeIndices:        indices.NonSignerStakeIndices,
	}
	aggregatedResponsesC <- blsAggregationServiceResponse
}

// verifySignature verifies that a signature is valid against the operator pubkey stored in the
// operatorsAvsStateDict for that particular task
func (a *BlsAggregatorBuilder) verifySignature(
	taskIndex types.TaskIndex,
	signedTaskResponseDigest types.SignedTaskResponseDigest,
	operatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState,
) error {
	_, ok := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId]
	if !ok {
		a.logger.Warnf("Operator %#v not found. Skipping message.", signedTaskResponseDigest.OperatorId)
		return OperatorNotPartOfTaskQuorumErrorFn(signedTaskResponseDigest.OperatorId, taskIndex)
	}

	taskResponseDigest, err := a.hashFunction(signedTaskResponseDigest.TaskResponse)
	if err != nil {
		a.logger.Error(
			"Failed to hash task response, skipping.",
			"taskIndex",
			taskIndex,
			"signedTaskResponseDigest",
			signedTaskResponseDigest,
			"err",
			err,
		)
		return HashFunctionError(err)
	}

	// verify that the msg actually came from the correct operator
	operatorG2Pubkey := operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].OperatorInfo.Pubkeys.G2Pubkey
	if operatorG2Pubkey == nil {
		a.logger.Error(
			"Operator G2 pubkey not found",
			"operatorId",
			signedTaskResponseDigest.OperatorId,
			"taskId",
			taskIndex,
		)
		return fmt.Errorf(
			"taskId %d: Operator G2 pubkey not found (operatorId: %x)",
			taskIndex,
			signedTaskResponseDigest.OperatorId,
		)
	}
	a.logger.Debug("Verifying signed task response digest signature",
		"operatorG2Pubkey", operatorG2Pubkey,
		"taskResponseDigest", taskResponseDigest,
		"blsSignature", signedTaskResponseDigest.BlsSignature,
	)

	// if the operator signs a digest that is not the digest of the TaskResponse submitted in ProcessNewTask
	// then the signature will not be verified
	signatureVerified, err := signedTaskResponseDigest.BlsSignature.Verify(operatorG2Pubkey, taskResponseDigest)
	if err != nil {
		return SignatureVerificationError(err)
	}
	if !signatureVerified {
		return ErrIncorrectSignature
	}
	return nil
}

// checkIfStakeThresholdsMet checks at least quorumThresholdPercentage of stake
// has signed for each quorum.
func checkIfStakeThresholdsMet(
	logger logging.Logger,
	signedStakePerQuorum map[types.QuorumNum]*big.Int,
	totalStakePerQuorum map[types.QuorumNum]*big.Int,
	quorumThresholdPercentagesMap map[types.QuorumNum]types.QuorumThresholdPercentage,
) bool {
	logger.Debug("Checking if stake thresholds are met.")
	for quorumNum, quorumThresholdPercentage := range quorumThresholdPercentagesMap {
		signedStakeByQuorum, ok := signedStakePerQuorum[quorumNum]
		if !ok {
			// signedStakePerQuorum not contain the quorum,
			// this case means signedStakePerQuorum has not signed for each quorum.
			// even the total stake for this quorum is zero.
			return false
		}

		totalStakeByQuorum, ok := totalStakePerQuorum[quorumNum]
		if !ok {
			// Note this case should not happen.
			// The `totalStakePerQuorum` is got from the contract, so if we not found the
			// totalStakeByQuorum, that means the code have a bug.
			logger.Errorf("TotalStake not found for quorum %d.", quorumNum)
			return false
		}

		logger.Debug("Stakes for quorum",
			"quorumNum", quorumNum,
			"totalStakeByQuorum", totalStakeByQuorum,
			"signedStakeByQuorum", signedStakeByQuorum)

		// we check that signedStake >= totalStake * quorumThresholdPercentage / 100
		// to be exact (and do like the contracts), we actually check that
		// signedStake * 100 >= totalStake * quorumThresholdPercentage
		signedStake := big.NewInt(0).Mul(signedStakeByQuorum, big.NewInt(100))
		thresholdStake := big.NewInt(0).Mul(totalStakeByQuorum, big.NewInt(int64(quorumThresholdPercentage)))

		logger.Debug("Checking if signed stake is greater than threshold",
			"signedStake", signedStake,
			"thresholdStake", thresholdStake)
		if signedStake.Cmp(thresholdStake) < 0 {
			return false
		}
	}
	return true
}

func cloneStakePerQuorumMap(stakes map[types.QuorumNum]types.StakeAmount) map[types.QuorumNum]types.StakeAmount {
	out := make(map[types.QuorumNum]types.StakeAmount, len(stakes))
	for k, v := range stakes {
		out[k] = new(big.Int).Set(v)
	}
	return out
}
