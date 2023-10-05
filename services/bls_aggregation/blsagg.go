package blsagg

import (
	"context"
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
	taskExpiredError = fmt.Errorf("task expired")
)

// BlsAggretationServiceResponse is the response from the bls aggregation service
// it's half of the data needed to build the NonSignerStakesAndSignature struct
// needed as input to the CheckSignatures function in the BLSSignatureChecker contract
type BlsAggretationServiceResponse struct {
	err                 error
	NonSignersPubkeysG1 []*bls.G1Point
	QuorumApksG1        []*bls.G1Point
	SignersApkG2        *bls.G2Point
	SignersAggSigG1     *bls.Signature
}

type BlsAggregationService struct {
	// ResponseC is the channel which all goroutines share to send their responses back to the
	// main thread after they are done aggregating (either they reached the threshold, or timeout expired)
	ResponseC chan BlsAggretationServiceResponse
	// taskChans are the channels to send the signed tasks to the goroutines processing them
	taskChans map[types.TaskIndex]chan types.SignedTaskResponseDigest
	// we add chans to taskChans from the main thread (InitializeNewTask) when we create new tasks,
	// we read them in ProcessNewSignature from the main thread when we receive new signed tasks,
	// and remove them from its respective goroutine when the task is completed or reached timeout
	// we thus need a mutex to protect taskChans
	taskChansMutex     sync.Mutex
	avsRegistryService avsregistry.AvsRegistryService
	logger             logging.Logger
}

func NewBlsAggregationService(avsRegistryService avsregistry.AvsRegistryService, logger logging.Logger) *BlsAggregationService {
	return &BlsAggregationService{
		ResponseC:          make(chan BlsAggretationServiceResponse),
		taskChans:          make(map[types.TaskIndex]chan types.SignedTaskResponseDigest),
		avsRegistryService: avsRegistryService,
		logger:             logger,
	}
}

func (a *BlsAggregationService) InitializeNewTask(
	ctx context.Context,
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
) {
	taskC := make(chan types.SignedTaskResponseDigest)
	a.taskChansMutex.Lock()
	a.taskChans[taskIndex] = taskC
	a.taskChansMutex.Unlock()
	go a.singleTaskAggregatorGoroutineFunc(taskIndex, taskCreatedBlock, quorumNumbers, quorumThresholdPercentages, taskC)
}

// TODO: this currently assumes that the signature is correct.
// we should also add function that actually checks the signature
func (a *BlsAggregationService) ProcessNewSignature(
	taskIndex types.TaskIndex,
	taskResponseDigest types.TaskResponseDigest,
	blsSignature *bls.Signature,
	operatorId bls.OperatorId,
) error {
	a.taskChansMutex.Lock()
	taskC, taskInitialized := a.taskChans[taskIndex]
	a.taskChansMutex.Unlock()
	if !taskInitialized {
		return fmt.Errorf("task %d not initialized or already completed", taskIndex)
	}
	// TODO(samlaf): should we check that the signature is correct here?
	// send the task to the goroutine processing this task
	taskC <- types.SignedTaskResponseDigest{
		TaskResponseDigest: taskResponseDigest,
		BlsSignature:       blsSignature,
		OperatorId:         operatorId,
	}
	return nil
}

func (a *BlsAggregationService) singleTaskAggregatorGoroutineFunc(
	taskIndex types.TaskIndex,
	taskCreatedBlock uint32,
	quorumNumbers []types.QuorumNum,
	quorumThresholdPercentages []types.QuorumThresholdPercentage,
	taskRespC <-chan types.SignedTaskResponseDigest,
) {
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
	// TODO(samlaf): change this to the real TTR (get it from onchain or passed when task created)
	taskExpiredTimer := time.NewTimer(5 * time.Second)

	signersApkG2 := bls.NewZeroG2Point()
	signersAggSigG1 := bls.NewZeroSignature()
	signersTotalStakePerQuorum := make(map[types.QuorumNum]*big.Int)
	signersOperatorIdsSet := make(map[types.OperatorId]bool)
F:
	for {
		select {
		case signedTaskResponseDigest := <-taskRespC:
			signersAggSigG1.Add(signedTaskResponseDigest.BlsSignature)
			signersApkG2.Add(operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].Pubkeys.G2Pubkey)
			signersOperatorIdsSet[signedTaskResponseDigest.OperatorId] = true

			for quorumNum, stake := range operatorsAvsStateDict[signedTaskResponseDigest.OperatorId].StakePerQuorum {
				if _, ok := signersTotalStakePerQuorum[quorumNum]; !ok {
					signersTotalStakePerQuorum[quorumNum] = big.NewInt(0)
				}
				signersTotalStakePerQuorum[quorumNum].Add(signersTotalStakePerQuorum[quorumNum], stake)
			}
			if checkIfStakeThresholdsMet(signersTotalStakePerQuorum, totalStakePerQuorum, quorumThresholdPercentagesMap) {
				break F
			}
		case <-taskExpiredTimer.C:
			a.ResponseC <- BlsAggretationServiceResponse{
				err:                 taskExpiredError,
				NonSignersPubkeysG1: nil,
				QuorumApksG1:        nil,
				SignersApkG2:        nil,
				SignersAggSigG1:     nil,
			}
			a.taskChansMutex.Lock()
			delete(a.taskChans, taskIndex)
			a.taskChansMutex.Unlock()
			return
		}
	}

	quorumApksG1 := []*bls.G1Point{}
	for _, quorumNumber := range quorumNumbers {
		quorumApksG1 = append(quorumApksG1, quorumsAvsStakeDict[quorumNumber].AggPubkeyG1)
	}
	a.ResponseC <- BlsAggretationServiceResponse{
		err:                 nil,
		NonSignersPubkeysG1: getG1PubkeysOfNonSigners(signersOperatorIdsSet, operatorsAvsStateDict),
		QuorumApksG1:        quorumApksG1,
		SignersApkG2:        signersApkG2,
		SignersAggSigG1:     signersAggSigG1,
	}
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
