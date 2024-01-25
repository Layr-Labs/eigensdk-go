package aggregation

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	sdkecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/ecdsa/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// TestEcdsaAgg is a suite of test that tests the main aggregation logic of the aggregation service
// it don't check any of the indices fields because those are just provided as a convenience to the caller
// and aren't related to the main logic which we actually need to test
// they are gotten from a call to the chain at the end of the aggregation so we should test that elsewhere
func TestEcdsaAgg(t *testing.T) {

	// we hardcode this for now, until we implement this feature properly
	// 1 second seems to be enough for tests to pass. Currently takes 5s to run all tests
	tasksTimeToExpiry := 1 * time.Second

	t.Run("1 quorum 1 operator 1 correct signature", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponseDigest := types.TaskResponseDigest{123}
		ecdsaSig, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err = ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSig, testOperator1.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          taskIndex,
			TaskResponseDigest: taskResponseDigest,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId},
			Signatures:         [][]byte{ecdsaSig},
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 3 operator 3 correct signatures", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator3 := newTestOperator(
			newEcdsaPrivKeyPanics("0x3"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(300), 1: big.NewInt(100)},
		)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponseDigest := types.TaskResponseDigest{123}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1, testOperator2, testOperator3})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err := ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		ecdsaSigOp1, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)
		ecdsaSigOp2, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator2.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSigOp2, testOperator2.OperatorId)
		require.Nil(t, err)
		ecdsaSigOp3, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator3.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSigOp3, testOperator3.OperatorId)
		require.Nil(t, err)

		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          taskIndex,
			TaskResponseDigest: taskResponseDigest,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId, testOperator2.OperatorId, testOperator3.OperatorId},
			Signatures:         [][]byte{ecdsaSigOp1, ecdsaSigOp2, ecdsaSigOp3},
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("2 quorums 2 operators 2 correct signatures", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0, 1}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
		taskResponseDigest := types.TaskResponseDigest{123}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err := ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		ecdsaSigOp1, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)
		ecdsaSigOp2, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator2.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSigOp2, testOperator2.OperatorId)
		require.Nil(t, err)

		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          taskIndex,
			TaskResponseDigest: taskResponseDigest,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId, testOperator2.OperatorId},
			Signatures:         [][]byte{ecdsaSigOp1, ecdsaSigOp2},
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("2 concurrent tasks 2 quorums 2 operators 2 correct signatures", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		quorumNumbers := []types.QuorumNum{0, 1}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		// initialize 2 concurrent tasks
		task1Index := types.TaskIndex(1)
		task1ResponseDigest := types.TaskResponseDigest{123}
		err := ecdsaAggServ.InitializeNewTask(task1Index, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		task2Index := types.TaskIndex(2)
		task2ResponseDigest := types.TaskResponseDigest{230}
		err = ecdsaAggServ.InitializeNewTask(task2Index, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)

		ecdsaSigTask1Op1, err := sdkecdsa.SignMsg(task1ResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), task1Index, task1ResponseDigest, ecdsaSigTask1Op1, testOperator1.OperatorId)
		require.Nil(t, err)
		ecdsaSigTask2Op1, err := sdkecdsa.SignMsg(task2ResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), task2Index, task2ResponseDigest, ecdsaSigTask2Op1, testOperator1.OperatorId)
		require.Nil(t, err)
		ecdsaSigTask1Op2, err := sdkecdsa.SignMsg(task1ResponseDigest[:], testOperator2.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), task1Index, task1ResponseDigest, ecdsaSigTask1Op2, testOperator2.OperatorId)
		require.Nil(t, err)
		ecdsaSigTask2Op2, err := sdkecdsa.SignMsg(task2ResponseDigest[:], testOperator2.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), task2Index, task2ResponseDigest, ecdsaSigTask2Op2, testOperator2.OperatorId)
		require.Nil(t, err)

		wantAggregationServiceResponseTask1 := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          task1Index,
			TaskResponseDigest: task1ResponseDigest,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId, testOperator2.OperatorId},
			Signatures:         [][]byte{ecdsaSigTask1Op1, ecdsaSigTask1Op2},
		}
		wantAggregationServiceResponseTask2 := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          task2Index,
			TaskResponseDigest: task2ResponseDigest,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId, testOperator2.OperatorId},
			Signatures:         [][]byte{ecdsaSigTask2Op1, ecdsaSigTask2Op2},
		}

		// we don't know which of task1 or task2 responses will be received first
		gotAggregationServiceResponseTaskFirstReceived := <-ecdsaAggServ.aggregatedResponsesC
		gotAggregationServiceResponseTaskSecondReceived := <-ecdsaAggServ.aggregatedResponsesC

		if gotAggregationServiceResponseTaskFirstReceived.TaskIndex == task1Index {
			require.EqualValues(t, wantAggregationServiceResponseTask1, gotAggregationServiceResponseTaskFirstReceived)
			require.EqualValues(t, wantAggregationServiceResponseTask2, gotAggregationServiceResponseTaskSecondReceived)
		} else {
			require.EqualValues(t, wantAggregationServiceResponseTask2, gotAggregationServiceResponseTaskFirstReceived)
			require.EqualValues(t, wantAggregationServiceResponseTask1, gotAggregationServiceResponseTaskSecondReceived)
		}
	})

	t.Run("1 quorum 1 operator 0 signatures - task expired", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err := ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err: TaskExpiredError,
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 1 correct signature quorumThreshold 50% - verified", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{50}
		taskResponseDigest := types.TaskResponseDigest{123}

		ecdsaSig, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err = ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSig, testOperator1.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          taskIndex,
			TaskResponseDigest: taskResponseDigest,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId},
			Signatures:         [][]byte{ecdsaSig},
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 1 correct signature quorumThreshold 60% - task expired", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{60}
		taskResponseDigest := types.TaskResponseDigest{123}
		ecdsaSig, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err = ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSig, testOperator1.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err: TaskExpiredError,
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("send signature of task that isn't initialized - task not found error", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		taskResponseDigest := types.TaskResponseDigest{123}
		ecdsaSig, err := sdkecdsa.SignMsg(taskResponseDigest[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest, ecdsaSig, testOperator1.OperatorId)
		require.Equal(t, TaskNotFoundErrorFn(taskIndex), err)
	})

	// this is an edge case as typically we would send new tasks and listen for task responses in a for select loop
	// but this test makes sure the context deadline exceeded can get us out of a deadlock
	t.Run("send new signedTaskDigest before listen on responseChan - context timeout cancels the request to prevent deadlock", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err := ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		taskResponseDigest1 := types.TaskResponseDigest{1}
		ecdsaSigOp1, err := sdkecdsa.SignMsg(taskResponseDigest1[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest1, ecdsaSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)

		taskResponseDigest2 := types.TaskResponseDigest{2}
		ecdsaSigOp2, err := sdkecdsa.SignMsg(taskResponseDigest2[:], testOperator2.EcdsaPrivKey)
		require.Nil(t, err)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		err = ecdsaAggServ.ProcessNewSignature(ctx, taskIndex, taskResponseDigest2, ecdsaSigOp2, testOperator2.OperatorId)
		// this should timeout because the task goroutine is blocked on the response channel (since we only listen for it below)
		require.Equal(t, context.DeadlineExceeded, err)

		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err:                nil,
			TaskIndex:          taskIndex,
			TaskResponseDigest: taskResponseDigest1,
			SignerIds:          []types.EcdsaOperatorId{testOperator1.OperatorId},
			Signatures:         [][]byte{ecdsaSigOp1},
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 2 signatures on 2 different msgs - task expired", func(t *testing.T) {
		testOperator1 := newTestOperator(
			newEcdsaPrivKeyPanics("0x1"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		testOperator2 := newTestOperator(
			newEcdsaPrivKeyPanics("0x2"),
			map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		)
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := []types.QuorumNum{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestEcdsaOperator{testOperator1, testOperator2})
		noopLogger := logging.NewNoopLogger()
		ecdsaAggServ := NewEcdsaAggregatorService(fakeAvsRegistryService, noopLogger)

		err := ecdsaAggServ.InitializeNewTask(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		require.Nil(t, err)
		taskResponseDigest1 := types.TaskResponseDigest{1}

		ecdsaSigOp1, err := sdkecdsa.SignMsg(taskResponseDigest1[:], testOperator1.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest1, ecdsaSigOp1, testOperator1.OperatorId)
		require.Nil(t, err)
		taskResponseDigest2 := types.TaskResponseDigest{2}
		ecdsaSigOp2, err := sdkecdsa.SignMsg(taskResponseDigest2[:], testOperator2.EcdsaPrivKey)
		require.Nil(t, err)
		err = ecdsaAggServ.ProcessNewSignature(context.Background(), taskIndex, taskResponseDigest2, ecdsaSigOp2, testOperator2.OperatorId)
		require.Nil(t, err)
		wantAggregationServiceResponse := EcdsaAggregationServiceResponse{
			Err: TaskExpiredError,
		}
		gotAggregationServiceResponse := <-ecdsaAggServ.aggregatedResponsesC
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})
}

func newEcdsaPrivKeyPanics(hexKey string) *ecdsa.PrivateKey {
	keypair, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	return keypair
}

func newTestOperator(ecdsaPrivKey *ecdsa.PrivateKey, stakePerQuorum map[types.QuorumNum]types.StakeAmount) types.TestEcdsaOperator {
	return types.TestEcdsaOperator{
		OperatorId:     crypto.PubkeyToAddress(ecdsaPrivKey.PublicKey),
		StakePerQuorum: stakePerQuorum,
		EcdsaPrivKey:   ecdsaPrivKey,
	}
}
