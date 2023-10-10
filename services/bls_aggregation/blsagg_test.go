package blsagg

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/stretchr/testify/require"
)

// TestBlsAgg is a suite of test that tests the main aggregation logic of the aggregation service
// it don't check any of the indices fields because those are just provided as a convenience to the caller
// and aren't related to the main logic which we actually need to test
// they are gotten from a call to the chain at the end of the aggregation so we should test that elsewhere
func TestBlsAgg(t *testing.T) {

	// we hardcode this for now, until we implement this feature properly
	// 1 second seems to be enough for tests to pass. Currently takes 5s to run all tests
	tasksTimeToExpiry := 1 * time.Second

	t.Run("1 quorum 1 operator 1 correct signature", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponseDigest := types.TaskResponseDigest{123}
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSig, testOperator1.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1:        []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1()},
			SignersApkG2:        testOperator1.BlsKeypair.GetPubKeyG2(),
			SignersAggSigG1:     testOperator1.BlsKeypair.SignMessage(taskResponseDigest),
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 3 operator 3 correct signatures", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
		}
		testOperator3 := types.TestOperator{
			OperatorId:     types.OperatorId{3},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(300), 1: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x3"),
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponseDigest := types.TaskResponseDigest{123}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1, testOperator2, testOperator3})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSigOp2, testOperator2.OperatorId)
		require.Nil(t, err)
		blsSigOp3 := testOperator3.BlsKeypair.SignMessage(taskResponseDigest)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSigOp3, testOperator3.OperatorId)
		require.Nil(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1: []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1().
				Add(testOperator2.BlsKeypair.GetPubKeyG1()).
				Add(testOperator3.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().
				Add(testOperator2.BlsKeypair.GetPubKeyG2()).
				Add(testOperator3.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).
				Add(testOperator2.BlsKeypair.SignMessage(taskResponseDigest)).
				Add(testOperator3.BlsKeypair.SignMessage(taskResponseDigest)),
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 1 operator 0 signatures - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredError,
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 1 correct signature quorumThreshold 50% - verified", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{50}
		taskResponseDigest := types.TaskResponseDigest{123}
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSig, testOperator1.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{testOperator2.BlsKeypair.GetPubKeyG1()},
			QuorumApksG1:        []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1().Add(testOperator2.BlsKeypair.GetPubKeyG1())},
			SignersApkG2:        testOperator1.BlsKeypair.GetPubKeyG2(),
			SignersAggSigG1:     testOperator1.BlsKeypair.SignMessage(taskResponseDigest),
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 1 correct signature quorumThreshold 60% - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{60}
		taskResponseDigest := types.TaskResponseDigest{123}
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSig, testOperator1.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredError,
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("send signature of task that isn't initialized - task not found error", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		taskResponseDigest := types.TaskResponseDigest{123}
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, blsSig, testOperator1.OperatorId)
		require.Equal(t, TaskNotFoundErrorF(taskIndex), err)
	})

	// this is an edge case as typically we would send new tasks and listen for task responses in a for select loop
	// but this test makes sure the context deadline exceeded can get us out of a deadlock
	t.Run("send new signedTaskDigest before listening  - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		taskResponseDigest1 := types.TaskResponseDigest{1}
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest1)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest1, blsSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)

		taskResponseDigest2 := types.TaskResponseDigest{2}
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest2)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = blsAggServ.ProcessNewSignature(ctx, taskIndex, taskResponseDigest2, blsSigOp2, testOperator2.OperatorId)
		// this should timeout because the task goroutine is blocked on the response channel (since we only listen for it below)
		require.Equal(t, context.DeadlineExceeded, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponseDigest:  taskResponseDigest1,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1:        []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1()},
			SignersApkG2:        testOperator1.BlsKeypair.GetPubKeyG2(),
			SignersAggSigG1:     testOperator1.BlsKeypair.SignMessage(taskResponseDigest1),
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 2 signatures on 2 different msgs - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, noopLogger)

		err := blsAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		taskResponseDigest1 := types.TaskResponseDigest{1}
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest1)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest1, blsSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)
		taskResponseDigest2 := types.TaskResponseDigest{2}
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest2)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest2, blsSigOp2, testOperator2.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredError,
		}
		gotAggregationServiceResponse := <-blsAggServ.ResponseC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})
}

func newBlsKeyPairPanics(hexKey string) *bls.KeyPair {
	keypair, err := bls.NewKeyPairFromString(hexKey)
	if err != nil {
		panic(err)
	}
	return keypair
}
