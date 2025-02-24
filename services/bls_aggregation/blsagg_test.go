package blsagg

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"log/slog"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/utils"
	avssm "github.com/Layr-Labs/eigensdk-go/contracts/bindings/MockAvsServiceManager"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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

	hashFunction := func(taskResponse types.TaskResponse) (types.TaskResponseDigest, error) {
		taskResponseBytes, err := json.Marshal(taskResponse)
		if err != nil {
			return types.TaskResponseDigest{}, err
		}
		return types.TaskResponseDigest(sha256.Sum256(taskResponseBytes)), nil
	}

	wrongHashFunction := func(taskResponse types.TaskResponse) (types.TaskResponseDigest, error) {
		taskResponseBytes, err := json.Marshal(taskResponse)
		if err != nil {
			return types.TaskResponseDigest{}, err
		}
		// append something to the taskResponseBytes to make it different
		taskResponseBytes = append(taskResponseBytes, []byte("something")...)
		return types.TaskResponseDigest(sha256.Sum256(taskResponseBytes)), nil
	}

	type mockTaskResponse struct {
		Value int
	}

	t.Run("1 quorum 1 operator 1 correct signature", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponse := mockTaskResponse{123} // Initialize with appropriate data

		// Compute the TaskResponseDigest as the SHA-256 sum of the TaskResponse
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1:        []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1()},
			SignersApkG2:        testOperator1.BlsKeypair.GetPubKeyG2(),
			SignersAggSigG1:     testOperator1.BlsKeypair.SignMessage(taskResponseDigest),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
	})

	t.Run("1 quorum 3 operator 3 correct signatures", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		testOperator3 := types.TestOperator{
			OperatorId:     types.OperatorId{3},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(300)},
			BlsKeypair:     newBlsKeyPairPanics("0x3"),
			Socket:         "localhost:8082",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponse := mockTaskResponse{123}

		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2, testOperator3},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)
		blsSigOp3 := testOperator3.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature3 := NewTaskSignature(taskIndex, taskResponse, blsSigOp3, testOperator3.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature3,
		)
		require.Nil(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
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
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
	})

	t.Run("2 quorums 2 operators 2 correct signatures", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0, 1}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
		taskResponse := mockTaskResponse{123}
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)

		op1G2Key := testOperator1.BlsKeypair.GetPubKeyG2()
		op2G2Key := testOperator2.BlsKeypair.GetPubKeyG2()
		op1Signature := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		op2Signature := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1: []*bls.G1Point{
				bls.NewZeroG1Point().
					Add(testOperator1.BlsKeypair.GetPubKeyG1()).
					Add(testOperator2.BlsKeypair.GetPubKeyG1()),
				bls.NewZeroG1Point().
					Add(testOperator1.BlsKeypair.GetPubKeyG1()).
					Add(testOperator2.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2:    op1G2Key.Add(op1G2Key).Add(op2G2Key).Add(op2G2Key),
			SignersAggSigG1: op1Signature.Add(op1Signature).Add(op2Signature).Add(op2Signature),
			// each key is added twice because both operators stake on two quorums
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("2 concurrent tasks 2 quorums 2 operators 2 correct signatures", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		quorumNumbers := types.QuorumNums{0, 1}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		// initialize 2 concurrent tasks
		task1Index := types.TaskIndex(1)
		task1Response := mockTaskResponse{123}
		task1ResponseDigest, err := hashFunction(task1Response)
		require.Nil(t, err)
		metadata1 := NewTaskMetadata(task1Index, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata1)
		require.Nil(t, err)
		task2Index := types.TaskIndex(2)
		task2Response := mockTaskResponse{234}
		task2ResponseDigest, err := hashFunction(task2Response)
		require.Nil(t, err)
		metadata2 := NewTaskMetadata(task2Index, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata2)
		require.Nil(t, err)

		blsSigTask1Op1 := testOperator1.BlsKeypair.SignMessage(task1ResponseDigest)
		taskSignature1Op1 := NewTaskSignature(task1Index, task1Response, blsSigTask1Op1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1Op1,
		)
		require.Nil(t, err)
		blsSigTask2Op1 := testOperator1.BlsKeypair.SignMessage(task2ResponseDigest)
		taskSignature2Op1 := NewTaskSignature(task2Index, task2Response, blsSigTask2Op1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2Op1,
		)
		require.Nil(t, err)
		blsSigTask1Op2 := testOperator2.BlsKeypair.SignMessage(task1ResponseDigest)
		taskSignature1Op2 := NewTaskSignature(task1Index, task1Response, blsSigTask1Op2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1Op2,
		)
		require.Nil(t, err)
		blsSigTask2Op2 := testOperator2.BlsKeypair.SignMessage(task2ResponseDigest)
		taskSignature2Op2 := NewTaskSignature(task2Index, task2Response, blsSigTask2Op2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2Op2,
		)
		require.Nil(t, err)

		wantAggregationServiceResponseTask1 := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           task1Index,
			TaskResponse:        task1Response,
			TaskResponseDigest:  task1ResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1: []*bls.G1Point{
				bls.NewZeroG1Point().
					Add(testOperator1.BlsKeypair.GetPubKeyG1()).
					Add(testOperator2.BlsKeypair.GetPubKeyG1()),
				bls.NewZeroG1Point().
					Add(testOperator1.BlsKeypair.GetPubKeyG1()).
					Add(testOperator2.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: bls.NewZeroG2Point().
				Add(testOperator1.BlsKeypair.GetPubKeyG2()).
				Add(testOperator1.BlsKeypair.GetPubKeyG2()).
				Add(testOperator2.BlsKeypair.GetPubKeyG2()).
				Add(testOperator2.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(task1ResponseDigest).
				Add(testOperator1.BlsKeypair.SignMessage(task1ResponseDigest)).
				Add(testOperator2.BlsKeypair.SignMessage(task1ResponseDigest)).
				Add(testOperator2.BlsKeypair.SignMessage(task1ResponseDigest)),
		}
		wantAggregationServiceResponseTask2 := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           task2Index,
			TaskResponse:        task2Response,
			TaskResponseDigest:  task2ResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1: []*bls.G1Point{
				bls.NewZeroG1Point().
					Add(testOperator1.BlsKeypair.GetPubKeyG1()).
					Add(testOperator2.BlsKeypair.GetPubKeyG1()),
				bls.NewZeroG1Point().
					Add(testOperator1.BlsKeypair.GetPubKeyG1()).
					Add(testOperator2.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().
				Add(testOperator1.BlsKeypair.GetPubKeyG2()).
				Add(testOperator2.BlsKeypair.GetPubKeyG2()).
				Add(testOperator2.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(task2ResponseDigest).
				Add(testOperator1.BlsKeypair.SignMessage(task2ResponseDigest)).
				Add(testOperator2.BlsKeypair.SignMessage(task2ResponseDigest)).
				Add(testOperator2.BlsKeypair.SignMessage(task2ResponseDigest)),
		}

		// we don't know which of task1 or task2 responses will be received first
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponseTaskFirstReceived := <-responseChannel
		responseChannel = blsAggServ.GetResponseChannel()
		gotAggregationServiceResponseTaskSecondReceived := <-responseChannel

		if gotAggregationServiceResponseTaskFirstReceived.TaskIndex == task1Index {
			require.EqualValues(t, wantAggregationServiceResponseTask1, gotAggregationServiceResponseTaskFirstReceived)
			require.EqualValues(t, wantAggregationServiceResponseTask2, gotAggregationServiceResponseTaskSecondReceived)
		} else {
			require.EqualValues(t, wantAggregationServiceResponseTask2, gotAggregationServiceResponseTaskFirstReceived)
			require.EqualValues(t, wantAggregationServiceResponseTask1, gotAggregationServiceResponseTaskSecondReceived)
		}
	})

	t.Run("1 quorum 2 operators operator 1 double sign", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}

		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		taskResponse := mockTaskResponse{123}

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		logger.Info("Initializing new task", "taskIndex", taskIndex)
		timeToExpire := 10 * time.Second
		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, timeToExpire)
		err := blsAggServ.InitializeNewTask(metadata)
		require.NoError(t, err)

		taskResponseDigest, err := hashFunction(taskResponse)
		require.NoError(t, err)

		logger.Info("Processing first signature", "operatorId", testOperator1.OperatorId)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.NoError(t, err)

		logger.Info("Processing second signature (Operator 1 double sign)", "operatorId", testOperator1.OperatorId)
		blsSigOp1Dup := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1Dup := NewTaskSignature(taskIndex, taskResponse, blsSigOp1Dup, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1Dup,
		)

		if err != nil {
			logger.Info("Received error from second signature", "error", err)
			require.Contains(t, err.Error(), "duplicate signature")
		} else {
			t.Fatal("Expected an error for duplicate signature, but got nil")
		}

		logger.Info("Processing second signature", "operatorId", testOperator2.OperatorId)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.NoError(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1: []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1().
				Add(testOperator2.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().
				Add(testOperator2.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).
				Add(testOperator2.BlsKeypair.SignMessage(taskResponseDigest)),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)

	})
	t.Run("1 quorum 1 operator 0 signatures - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err := blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredErrorFn(taskIndex),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 1 correct signature quorumThreshold 50% - verified", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{50}
		taskResponse := mockTaskResponse{123}
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{testOperator2.BlsKeypair.GetPubKeyG1()},
			QuorumApksG1: []*bls.G1Point{
				testOperator1.BlsKeypair.GetPubKeyG1().Add(testOperator2.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2:    testOperator1.BlsKeypair.GetPubKeyG2(),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("1 quorum 2 operator 1 correct signature quorumThreshold 60% - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{60}
		taskResponse := mockTaskResponse{123}
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredErrorFn(taskIndex),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("2 quorums 2 operators which just stake one quorum; 2 correct signature - verified", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId: types.OperatorId{1},
			// Note the quorums is {0, 1}, but operator id 1 just stake 0.
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId: types.OperatorId{2},
			// Note the quorums is {0, 1}, but operator id 1 just stake 0.
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0, 1}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
		taskResponse := mockTaskResponse{123}
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{},
			QuorumApksG1: []*bls.G1Point{
				bls.NewZeroG1Point().Add(testOperator1.BlsKeypair.GetPubKeyG1()),
				bls.NewZeroG1Point().Add(testOperator2.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().Add(testOperator2.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).
				Add(testOperator2.BlsKeypair.SignMessage(taskResponseDigest)),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
	})

	t.Run("2 quorums 3 operators which just stake one quorum; 2 correct signature quorumThreshold 50% - verified",
		func(t *testing.T) {
			testOperator1 := types.TestOperator{
				OperatorId: types.OperatorId{1},
				// Note the quorums is {0, 1}, but operator id 1 just stake 0.
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
				BlsKeypair:     newBlsKeyPairPanics("0x1"),
				Socket:         "localhost:8080",
			}
			testOperator2 := types.TestOperator{
				OperatorId: types.OperatorId{2},
				// Note the quorums is {0, 1}, but operator id 2 just stake 1.
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{1: big.NewInt(200)},
				BlsKeypair:     newBlsKeyPairPanics("0x2"),
				Socket:         "localhost:8081",
			}
			testOperator3 := types.TestOperator{
				OperatorId:     types.OperatorId{3},
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
				BlsKeypair:     newBlsKeyPairPanics("0x3"),
				Socket:         "localhost:8082",
			}
			taskIndex := types.TaskIndex(0)
			quorumNumbers := types.QuorumNums{0, 1}
			quorumThresholdPercentages := []types.QuorumThresholdPercentage{50, 50}
			taskResponse := mockTaskResponse{123}
			taskResponseDigest, err := hashFunction(taskResponse)
			require.Nil(t, err)
			blockNum := uint32(1)

			fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
				blockNum,
				[]types.TestOperator{testOperator1, testOperator2, testOperator3},
			)
			logger := testutils.GetTestLogger()
			blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

			metadata := NewTaskMetadata(
				taskIndex,
				blockNum,
				quorumNumbers,
				quorumThresholdPercentages,
				tasksTimeToExpiry,
			)
			err = blsAggServ.InitializeNewTask(metadata)
			require.Nil(t, err)
			blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
			taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature1,
			)
			require.Nil(t, err)
			blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
			taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature2,
			)
			require.Nil(t, err)

			wantAggregationServiceResponse := BlsAggregationServiceResponse{
				Err:                nil,
				TaskIndex:          taskIndex,
				TaskResponse:       taskResponse,
				TaskResponseDigest: taskResponseDigest,
				NonSignersPubkeysG1: []*bls.G1Point{
					testOperator3.BlsKeypair.GetPubKeyG1(),
				},
				QuorumApksG1: []*bls.G1Point{
					bls.NewZeroG1Point().
						Add(testOperator1.BlsKeypair.GetPubKeyG1()).
						Add(testOperator3.BlsKeypair.GetPubKeyG1()),
					bls.NewZeroG1Point().
						Add(testOperator2.BlsKeypair.GetPubKeyG1()).
						Add(testOperator3.BlsKeypair.GetPubKeyG1()),
				},
				SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().Add(testOperator2.BlsKeypair.GetPubKeyG2()),
				SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).
					Add(testOperator2.BlsKeypair.SignMessage(taskResponseDigest)),
			}
			responseChannel := blsAggServ.GetResponseChannel()
			gotAggregationServiceResponse := <-responseChannel
			require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		},
	)

	t.Run("2 quorums 3 operators which just stake one quorum; 2 correct signature quorumThreshold 60% - task expired",
		func(t *testing.T) {
			testOperator1 := types.TestOperator{
				OperatorId: types.OperatorId{1},
				// Note the quorums is {0, 1}, but operator id 1 just stake 0.
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
				BlsKeypair:     newBlsKeyPairPanics("0x1"),
				Socket:         "localhost:8080",
			}
			testOperator2 := types.TestOperator{
				OperatorId: types.OperatorId{2},
				// Note the quorums is {0, 1}, but operator id 2 just stake 1.
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{1: big.NewInt(200)},
				BlsKeypair:     newBlsKeyPairPanics("0x2"),
				Socket:         "localhost:8081",
			}
			testOperator3 := types.TestOperator{
				OperatorId:     types.OperatorId{3},
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
				BlsKeypair:     newBlsKeyPairPanics("0x3"),
				Socket:         "localhost:8082",
			}
			taskIndex := types.TaskIndex(0)
			quorumNumbers := types.QuorumNums{0, 1}
			quorumThresholdPercentages := []types.QuorumThresholdPercentage{60, 60}
			taskResponse := mockTaskResponse{123}
			taskResponseDigest, err := hashFunction(taskResponse)
			require.Nil(t, err)
			blockNum := uint32(1)

			fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
				blockNum,
				[]types.TestOperator{testOperator1, testOperator2, testOperator3},
			)
			logger := testutils.GetTestLogger()
			blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

			metadata := NewTaskMetadata(
				taskIndex,
				blockNum,
				quorumNumbers,
				quorumThresholdPercentages,
				tasksTimeToExpiry,
			)
			err = blsAggServ.InitializeNewTask(metadata)
			require.Nil(t, err)
			blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
			taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature1,
			)
			require.Nil(t, err)
			blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
			taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature2,
			)
			require.Nil(t, err)

			wantAggregationServiceResponse := BlsAggregationServiceResponse{
				Err: TaskExpiredErrorFn(taskIndex),
			}
			responseChannel := blsAggServ.GetResponseChannel()
			gotAggregationServiceResponse := <-responseChannel
			require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
			require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		},
	)

	t.Run("2 quorums 1 operators which just stake one quorum; 1 signatures - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId: types.OperatorId{1},
			// Note the quorums is {0, 1}, but operator id 1 just stake 0.
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0, 1}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
		taskResponse := mockTaskResponse{123}
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blockNum := uint32(1)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
		require.Nil(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredErrorFn(taskIndex),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
	})

	t.Run("2 quorums 2 operators, 1 operator which just stake one quorum; 1 signatures - task expired",
		func(t *testing.T) {
			testOperator1 := types.TestOperator{
				OperatorId: types.OperatorId{1},
				// Note the quorums is {0, 1}, but operator id 1 just stake 0.
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
				BlsKeypair:     newBlsKeyPairPanics("0x1"),
				Socket:         "localhost:8080",
			}
			testOperator2 := types.TestOperator{
				OperatorId:     types.OperatorId{2},
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{1: big.NewInt(200)},
				BlsKeypair:     newBlsKeyPairPanics("0x2"),
				Socket:         "localhost:8081",
			}
			taskIndex := types.TaskIndex(0)
			quorumNumbers := types.QuorumNums{0, 1}
			quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}
			taskResponse := mockTaskResponse{123}
			taskResponseDigest, err := hashFunction(taskResponse)
			require.Nil(t, err)
			blockNum := uint32(1)

			fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
				blockNum,
				[]types.TestOperator{testOperator1, testOperator2},
			)
			logger := testutils.GetTestLogger()
			blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

			metadata := NewTaskMetadata(
				taskIndex,
				blockNum,
				quorumNumbers,
				quorumThresholdPercentages,
				tasksTimeToExpiry,
			)
			err = blsAggServ.InitializeNewTask(metadata)
			require.Nil(t, err)
			blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
			taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature,
			)
			require.Nil(t, err)

			wantAggregationServiceResponse := BlsAggregationServiceResponse{
				Err: TaskExpiredErrorFn(taskIndex),
			}
			responseChannel := blsAggServ.GetResponseChannel()
			gotAggregationServiceResponse := <-responseChannel
			require.EqualValues(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
			require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		},
	)

	t.Run("send signature of task that isn't initialized - task not found error", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		taskResponse := mockTaskResponse{123}
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(blockNum, []types.TestOperator{testOperator1})
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
		require.Equal(t, TaskNotFoundErrorFn(taskIndex), err)
	})

	// this is an edge case as typically we would send new tasks and listen for task responses in a for select loop
	// but this test makes sure the context deadline exceeded can get us out of a deadlock
	t.Run(
		"send new signedTaskDigest before listen on responseChan - context timeout cancels the request to prevent deadlock",
		func(t *testing.T) {
			testOperator1 := types.TestOperator{
				OperatorId:     types.OperatorId{1},
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
				BlsKeypair:     newBlsKeyPairPanics("0x1"),
				Socket:         "localhost:8080",
			}
			testOperator2OperatorId := types.OperatorId{2}
			testOperator2BlsKeypair := newBlsKeyPairPanics("0x2")

			blockNum := uint32(1)
			taskIndex := types.TaskIndex(0)
			quorumNumbers := types.QuorumNums{0}
			quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}

			fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
				blockNum,
				[]types.TestOperator{testOperator1},
			)
			logger := testutils.GetTestLogger()
			blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

			metadata := NewTaskMetadata(
				taskIndex,
				blockNum,
				quorumNumbers,
				quorumThresholdPercentages,
				tasksTimeToExpiry,
			)
			err := blsAggServ.InitializeNewTask(metadata)
			require.Nil(t, err)
			taskResponse1 := mockTaskResponse{1}
			taskResponseDigest1, err := hashFunction(taskResponse1)
			require.Nil(t, err)
			blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest1)
			taskSignature1 := NewTaskSignature(taskIndex, taskResponse1, blsSigOp1, testOperator1.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature1,
			)
			require.Nil(t, err)

			taskResponse2 := mockTaskResponse{2}
			taskResponseDigest2, err := hashFunction(taskResponse2)
			require.Nil(t, err)
			blsSigOp2 := testOperator2BlsKeypair.SignMessage(taskResponseDigest2)
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			taskSignature2 := NewTaskSignature(taskIndex, taskResponse2, blsSigOp2, testOperator2OperatorId)
			err = blsAggServ.ProcessNewSignature(
				ctx,
				taskSignature2,
			)
			// this should timeout because the task goroutine is blocked on the response channel (since we only listen
			// for it below)
			require.Equal(t, context.DeadlineExceeded, err)

			wantAggregationServiceResponse := BlsAggregationServiceResponse{
				Err:                 nil,
				TaskIndex:           taskIndex,
				TaskResponse:        taskResponse1,
				TaskResponseDigest:  taskResponseDigest1,
				NonSignersPubkeysG1: []*bls.G1Point{},
				QuorumApksG1:        []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1()},
				SignersApkG2:        testOperator1.BlsKeypair.GetPubKeyG2(),
				SignersAggSigG1:     testOperator1.BlsKeypair.SignMessage(taskResponseDigest1),
			}
			responseChannel := blsAggServ.GetResponseChannel()
			gotAggregationServiceResponse := <-responseChannel
			require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
			require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		},
	)

	t.Run("1 quorum 2 operator 2 signatures on 2 different msgs - task expired", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages, tasksTimeToExpiry)
		err := blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)
		taskResponse1 := mockTaskResponse{1}
		taskResponseDigest1, err := hashFunction(taskResponse1)
		require.Nil(t, err)
		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest1)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse1, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		taskResponse2 := mockTaskResponse{2}
		taskResponseDigest2, err := hashFunction(taskResponse2)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest2)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse2, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)
		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err: TaskExpiredErrorFn(taskIndex),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
	})

	t.Run("1 quorum 1 operator 1 invalid signature (TaskResponseDigest does not match TaskResponse)",
		func(t *testing.T) {
			testOperator1 := types.TestOperator{
				OperatorId:     types.OperatorId{1},
				StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
				BlsKeypair:     newBlsKeyPairPanics("0x1"),
			}
			blockNum := uint32(1)
			taskIndex := types.TaskIndex(0)
			quorumNumbers := types.QuorumNums{0}
			quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
			taskResponse := mockTaskResponse{123} // Initialize with appropriate data

			taskResponseDigest, err := wrongHashFunction(taskResponse)
			require.Nil(t, err)

			blsSig := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)

			fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
				blockNum,
				[]types.TestOperator{testOperator1},
			)
			logger := testutils.GetTestLogger()
			blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

			metadata := NewTaskMetadata(
				taskIndex,
				blockNum,
				quorumNumbers,
				quorumThresholdPercentages,
				tasksTimeToExpiry,
			)
			err = blsAggServ.InitializeNewTask(metadata)
			require.Nil(t, err)
			taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, testOperator1.OperatorId)
			err = blsAggServ.ProcessNewSignature(
				context.Background(),
				taskSignature,
			)
			require.EqualError(t, err, "signature verification failed. incorrect signature")
		},
	)

	t.Run("signatures are processed during window after quorum", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(200)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		testOperator3 := types.TestOperator{
			OperatorId:     types.OperatorId{3},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x3"),
			Socket:         "localhost:8082",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{67}
		taskResponse := mockTaskResponse{123}

		timeToExpiry := 5 * time.Second
		windowDuration := 1 * time.Second

		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2, testOperator3},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		start := time.Now()
		metadata := NewTaskMetadata(
			taskIndex,
			blockNum,
			quorumNumbers,
			quorumThresholdPercentages,
			tasksTimeToExpiry,
		).WithWindowDuration(windowDuration)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)

		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		// quorum has already been reached, but window should still be open
		require.Nil(t, err)
		blsSigOp3 := testOperator3.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature3 := NewTaskSignature(taskIndex, taskResponse, blsSigOp3, testOperator3.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature3,
		)
		require.Nil(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
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
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		elapsed := time.Since(start)
		t.Log("elapsed: ", elapsed.Seconds())
		require.True(
			t,
			elapsed.Seconds() >= windowDuration.Seconds(),
			"The aggregation response should be sent after the window finishes",
		)
		require.True(t, elapsed.Seconds() < timeToExpiry.Seconds())
	})

	t.Run("if quorum has been reached and the task expires during window, the response is sent", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		testOperator3 := types.TestOperator{
			OperatorId:     types.OperatorId{3},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x3"),
			Socket:         "localhost:8082",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{50}
		taskResponse := mockTaskResponse{123}

		timeToExpiry := 5 * time.Second
		windowDuration := 6 * time.Second

		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2, testOperator3},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		start := time.Now()
		metadata := NewTaskMetadata(
			taskIndex,
			blockNum,
			quorumNumbers,
			quorumThresholdPercentages,
			timeToExpiry,
		).WithWindowDuration(windowDuration)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)

		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)

		// quorum has already been reached, window will be open and receiving signature until the task expires

		blsSigOp3 := testOperator3.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature3 := NewTaskSignature(taskIndex, taskResponse, blsSigOp3, testOperator3.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature3,
		)
		require.Nil(t, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
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
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		elapsed := time.Since(start)
		t.Log("elapsed: ", elapsed.Seconds())
		require.True(t, elapsed.Seconds() >= timeToExpiry.Seconds())
		require.True(t, elapsed.Seconds() < windowDuration.Seconds())
	})

	t.Run("if window duration is zero, no signatures are aggregated after reaching quorum", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
		}
		testOperator3 := types.TestOperator{
			OperatorId:     types.OperatorId{3},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x3"),
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{50}
		taskResponse := mockTaskResponse{123}

		timeToExpiry := 5 * time.Second
		windowDuration := 0 * time.Second

		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2, testOperator3},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(
			taskIndex,
			blockNum,
			quorumNumbers,
			quorumThresholdPercentages,
			tasksTimeToExpiry,
		).WithWindowDuration(windowDuration)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)

		start := time.Now()

		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)

		time.Sleep(1 * time.Millisecond)
		// quorum has already been reached, next signatures should not be aggregated
		// this should timeout as the task goroutine is blocked on the response channel
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		blsSigOp3 := testOperator3.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature3 := NewTaskSignature(taskIndex, taskResponse, blsSigOp3, testOperator3.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			ctx,
			taskSignature3,
		)
		require.Equal(t, context.DeadlineExceeded, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{testOperator3.BlsKeypair.GetPubKeyG1()},
			QuorumApksG1: []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1().
				Add(testOperator2.BlsKeypair.GetPubKeyG1()).
				Add(testOperator3.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().
				Add(testOperator2.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).
				Add(testOperator2.BlsKeypair.SignMessage(taskResponseDigest)),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		elapsed := time.Since(start)
		t.Log("elapsed: ", elapsed.Seconds())
		require.True(t, elapsed.Seconds() < timeToExpiry.Seconds())
	})

	t.Run("no signatures are aggregated after window", func(t *testing.T) {
		testOperator1 := types.TestOperator{
			OperatorId:     types.OperatorId{1},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x1"),
			Socket:         "localhost:8080",
		}
		testOperator2 := types.TestOperator{
			OperatorId:     types.OperatorId{2},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x2"),
			Socket:         "localhost:8081",
		}
		testOperator3 := types.TestOperator{
			OperatorId:     types.OperatorId{3},
			StakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100)},
			BlsKeypair:     newBlsKeyPairPanics("0x3"),
			Socket:         "localhost:8082",
		}
		blockNum := uint32(1)
		taskIndex := types.TaskIndex(0)
		quorumNumbers := types.QuorumNums{0}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{50}
		taskResponse := mockTaskResponse{123}

		timeToExpiry := 10 * time.Second
		windowDuration := 1 * time.Second

		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)

		fakeAvsRegistryService := avsregistry.NewFakeAvsRegistryService(
			blockNum,
			[]types.TestOperator{testOperator1, testOperator2, testOperator3},
		)
		logger := testutils.GetTestLogger()
		blsAggServ := NewBlsAggregatorService(fakeAvsRegistryService, hashFunction, logger)

		metadata := NewTaskMetadata(
			taskIndex,
			blockNum,
			quorumNumbers,
			quorumThresholdPercentages,
			tasksTimeToExpiry,
		).WithWindowDuration(windowDuration)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)

		start := time.Now()

		blsSigOp1 := testOperator1.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature1 := NewTaskSignature(taskIndex, taskResponse, blsSigOp1, testOperator1.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature1,
		)
		require.Nil(t, err)
		blsSigOp2 := testOperator2.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature2 := NewTaskSignature(taskIndex, taskResponse, blsSigOp2, testOperator2.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature2,
		)
		require.Nil(t, err)

		time.Sleep(2 * time.Second)

		// quorum has already been reached, next signatures should not be aggregated
		// this should timeout as the task goroutine is blocked on the response channel
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		blsSigOp3 := testOperator3.BlsKeypair.SignMessage(taskResponseDigest)
		taskSignature3 := NewTaskSignature(taskIndex, taskResponse, blsSigOp3, testOperator3.OperatorId)
		err = blsAggServ.ProcessNewSignature(
			ctx,
			taskSignature3,
		)
		require.Equal(t, context.DeadlineExceeded, err)

		wantAggregationServiceResponse := BlsAggregationServiceResponse{
			Err:                 nil,
			TaskIndex:           taskIndex,
			TaskResponse:        taskResponse,
			TaskResponseDigest:  taskResponseDigest,
			NonSignersPubkeysG1: []*bls.G1Point{testOperator3.BlsKeypair.GetPubKeyG1()},
			QuorumApksG1: []*bls.G1Point{testOperator1.BlsKeypair.GetPubKeyG1().
				Add(testOperator2.BlsKeypair.GetPubKeyG1()).
				Add(testOperator3.BlsKeypair.GetPubKeyG1()),
			},
			SignersApkG2: testOperator1.BlsKeypair.GetPubKeyG2().
				Add(testOperator2.BlsKeypair.GetPubKeyG2()),
			SignersAggSigG1: testOperator1.BlsKeypair.SignMessage(taskResponseDigest).
				Add(testOperator2.BlsKeypair.SignMessage(taskResponseDigest)),
		}
		responseChannel := blsAggServ.GetResponseChannel()
		gotAggregationServiceResponse := <-responseChannel
		require.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		require.EqualValues(t, taskIndex, gotAggregationServiceResponse.TaskIndex)
		elapsed := time.Since(start)
		t.Log("elapsed: ", elapsed.Seconds())
		require.True(t, elapsed.Seconds() < timeToExpiry.Seconds())
	})
}

func TestIntegrationBlsAgg(t *testing.T) {

	tasksTimeToExpiry := 10 * time.Second

	hashFunction := func(taskResponse types.TaskResponse) (types.TaskResponseDigest, error) {
		taskResponseBytes, err := json.Marshal(taskResponse)
		if err != nil {
			return types.TaskResponseDigest{}, err
		}
		return types.TaskResponseDigest(sha256.Sum256(taskResponseBytes)), nil
	}
	type mockTaskResponse struct {
		Value int
	}

	anvilStateFileName := "contracts-deployed-anvil-state.json"
	anvilC, err := testutils.StartAnvilContainer(anvilStateFileName)
	require.NoError(t, err)
	anvilHttpEndpoint, err := anvilC.Endpoint(context.Background(), "http")
	require.NoError(t, err)
	anvilWsEndpoint, err := anvilC.Endpoint(context.Background(), "ws")
	require.NoError(t, err)
	contractAddrs := testutils.GetContractAddressesFromContractRegistry(anvilHttpEndpoint)

	t.Run("1 quorums 1 operator", func(t *testing.T) {
		// read input from JSON if available, otherwise use default values
		var defaultInput = struct {
			QuorumNumbers              types.QuorumNums                 `json:"quorum_numbers"`
			QuorumThresholdPercentages types.QuorumThresholdPercentages `json:"quorum_threshold_percentages"`
			BlsPrivKey                 string                           `json:"bls_key"`
		}{
			QuorumNumbers:              types.QuorumNums{0},
			QuorumThresholdPercentages: types.QuorumThresholdPercentages{100},
			BlsPrivKey:                 "0x1",
		}
		testData := testutils.NewTestData(defaultInput)

		// define operator ecdsa and bls private keys
		ecdsaPrivKey, err := crypto.HexToECDSA(testutils.ANVIL_FIRST_PRIVATE_KEY)
		require.NoError(t, err)
		blsPrivKeyHex := testData.Input.BlsPrivKey
		blsKeyPair := newBlsKeyPairPanics(blsPrivKeyHex)
		operatorId := types.OperatorIdFromG1Pubkey(blsKeyPair.GetPubKeyG1())

		// create avs clients to interact with contracts deployed on anvil
		ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)
		logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})
		avsClients, err := clients.BuildAll(clients.BuildAllConfig{
			EthHttpUrl:                 anvilHttpEndpoint,
			EthWsUrl:                   anvilWsEndpoint,
			RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
			OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
			AvsName:                    "avs",
			PromMetricsIpPortAddress:   "localhost:9090",
			ServiceManagerAddress:      contractAddrs.ServiceManager.String(),
		}, ecdsaPrivKey, logger)
		require.NoError(t, err)
		avsWriter := avsClients.AvsRegistryChainWriter

		// create aggregation service
		operatorsInfoService := operatorsinfo.NewOperatorsInfoServiceInMemory(
			context.TODO(),
			avsClients.AvsRegistryChainSubscriber,
			avsClients.AvsRegistryChainReader,
			nil,
			operatorsinfo.Opts{},
			logger,
		)
		avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(
			avsClients.AvsRegistryChainReader,
			operatorsInfoService,
			logger,
		)
		blsAggServ := NewBlsAggregatorService(avsRegistryService, hashFunction, logger)

		// register operator
		quorumNumbers := testData.Input.QuorumNumbers
		_, err = avsWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivKey,
			blsKeyPair,
			quorumNumbers,
			"socket",
			true,
		)
		require.NoError(t, err)

		// create the task related parameters: RBN, quorumThresholdPercentages, taskIndex and taskResponse
		curBlockNum, err := ethHttpClient.BlockNumber(context.Background())
		require.NoError(t, err)
		referenceBlockNumber := uint32(curBlockNum)
		// need to advance chain by 1 block because of the check in signatureChecker where RBN must be < current block
		// number
		testutils.AdvanceChainByNBlocksExecInContainer(context.TODO(), 1, anvilC)
		taskIndex := types.TaskIndex(0)
		taskResponse := mockTaskResponse{123} // Initialize with appropriate data
		quorumThresholdPercentages := testData.Input.QuorumThresholdPercentages

		// initialize the task
		metadata := NewTaskMetadata(
			taskIndex,
			uint32(curBlockNum),
			quorumNumbers,
			quorumThresholdPercentages,
			tasksTimeToExpiry,
		)
		err = blsAggServ.InitializeNewTask(metadata)
		require.Nil(t, err)

		// compute the signature and send it to the aggregation service
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blsSig := blsKeyPair.SignMessage(taskResponseDigest)
		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, operatorId)
		err = blsAggServ.ProcessNewSignature(
			context.Background(),
			taskSignature,
		)
		require.Nil(t, err)

		// wait for the response from the aggregation service and check the signature
		responseChannel := blsAggServ.GetResponseChannel()
		blsAggServiceResp := <-responseChannel

		avsServiceManager, err := avssm.NewContractMockAvsServiceManager(contractAddrs.ServiceManager, ethHttpClient)
		require.NoError(t, err)

		_, _, err = avsServiceManager.CheckSignatures(
			&bind.CallOpts{},
			taskResponseDigest,
			quorumNumbers.UnderlyingType(),
			uint32(referenceBlockNumber),
			blsAggServiceResp.toNonSignerStakesAndSignature(),
		)
		require.NoError(t, err)
	})

	t.Run("2 quorums 1 operator staking on both", func(t *testing.T) {
		// define operator ecdsa and bls private keys
		ecdsaPrivKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		require.NoError(t, err)
		blsPrivKeyHex := "0x1"
		blsKeyPair := newBlsKeyPairPanics(blsPrivKeyHex)
		operatorId := types.OperatorIdFromG1Pubkey(blsKeyPair.GetPubKeyG1())

		// create avs clients to interact with contracts deployed on anvil
		ethHttpClient, err := ethclient.Dial(anvilHttpEndpoint)
		require.NoError(t, err)
		logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})
		avsClients, err := clients.BuildAll(clients.BuildAllConfig{
			EthHttpUrl:                 anvilHttpEndpoint,
			EthWsUrl:                   anvilWsEndpoint,
			RegistryCoordinatorAddr:    contractAddrs.RegistryCoordinator.String(),
			OperatorStateRetrieverAddr: contractAddrs.OperatorStateRetriever.String(),
			AvsName:                    "avs",
			PromMetricsIpPortAddress:   "localhost:9090",
			ServiceManagerAddress:      contractAddrs.ServiceManager.String(),
		}, ecdsaPrivKey, logger)
		require.NoError(t, err)
		avsWriter := avsClients.AvsRegistryChainWriter
		avsServiceManager, err := avssm.NewContractMockAvsServiceManager(contractAddrs.ServiceManager, ethHttpClient)
		require.NoError(t, err)

		// create aggregation service
		operatorsInfoService := operatorsinfo.NewOperatorsInfoServiceInMemory(
			context.TODO(),
			avsClients.AvsRegistryChainSubscriber,
			avsClients.AvsRegistryChainReader,
			nil,
			operatorsinfo.Opts{},
			logger,
		)
		avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(
			avsClients.AvsRegistryChainReader,
			operatorsInfoService,
			logger,
		)
		blsAggServ := NewBlsAggregatorService(avsRegistryService, hashFunction, logger)

		// create quorum
		registryCoordinator, _ := regcoord.NewContractRegistryCoordinator(
			contractAddrs.RegistryCoordinator,
			ethHttpClient,
		)
		operatorSetParam := regcoord.ISlashingRegistryCoordinatorTypesOperatorSetParam{
			MaxOperatorCount:        10,
			KickBIPsOfOperatorStake: 1,
			KickBIPsOfTotalStake:    1,
		}
		strategyParam := []regcoord.IStakeRegistryTypesStrategyParams{
			{
				Strategy:   contractAddrs.Erc20MockStrategy,
				Multiplier: big.NewInt(1),
			},
		}
		noSendTxOpts, err := avsClients.TxManager.GetNoSendTxOpts()
		require.NoError(t, err)
		tx, err := registryCoordinator.CreateTotalDelegatedStakeQuorum(
			noSendTxOpts,
			operatorSetParam,
			big.NewInt(0),
			strategyParam,
		)
		require.NoError(t, err)
		_, err = avsClients.TxManager.Send(context.TODO(), tx, true)
		require.NoError(t, err)

		tx, err = registryCoordinator.CreateTotalDelegatedStakeQuorum(
			noSendTxOpts,
			operatorSetParam,
			big.NewInt(0),
			strategyParam,
		)
		require.NoError(t, err)
		_, err = avsClients.TxManager.Send(context.TODO(), tx, true)
		require.NoError(t, err)

		// register operator
		quorumNumbers := types.QuorumNums{1, 2}
		quorumThresholdPercentages := []types.QuorumThresholdPercentage{100, 100}

		_, err = avsWriter.RegisterOperator(
			context.Background(),
			ecdsaPrivKey,
			blsKeyPair,
			quorumNumbers,
			"socket",
			true,
		)
		require.NoError(t, err)

		// create the task related parameters: RBN, quorumThresholdPercentages, taskIndex and taskResponse
		curBlockNum, err := ethHttpClient.BlockNumber(context.Background())
		require.NoError(t, err)
		referenceBlockNumber := uint32(curBlockNum)
		// need to advance chain by 1 block because of the check in signatureChecker where RBN must be < current block
		// number
		testutils.AdvanceChainByNBlocksExecInContainer(context.TODO(), 1, anvilC)
		taskIndex := types.TaskIndex(0)
		taskResponse := mockTaskResponse{123} // Initialize with appropriate data

		newTaskMetadata := NewTaskMetadata(taskIndex,
			uint32(referenceBlockNumber),
			quorumNumbers,
			quorumThresholdPercentages,
			tasksTimeToExpiry,
		)
		// initialize the task
		err = blsAggServ.InitializeNewTask(newTaskMetadata)
		require.Nil(t, err)

		// compute the signature and send it to the aggregation service
		taskResponseDigest, err := hashFunction(taskResponse)
		require.Nil(t, err)
		blsSig := blsKeyPair.SignMessage(taskResponseDigest)
		taskSignature := NewTaskSignature(taskIndex, taskResponse, blsSig, operatorId)
		err = blsAggServ.ProcessNewSignature(context.Background(), taskSignature)
		require.Nil(t, err)

		// wait for the response from the aggregation service and check the signature
		responseChannel := blsAggServ.GetResponseChannel()
		blsAggServiceResp := <-responseChannel
		_, _, err = avsServiceManager.CheckSignatures(
			&bind.CallOpts{},
			taskResponseDigest,
			quorumNumbers.UnderlyingType(),
			uint32(referenceBlockNumber),
			blsAggServiceResp.toNonSignerStakesAndSignature(),
		)
		require.NoError(t, err)
	})
}

func newBlsKeyPairPanics(hexKey string) *bls.KeyPair {
	keypair, err := bls.NewKeyPairFromString(hexKey)
	if err != nil {
		panic(err)
	}
	return keypair
}

func (blsAggServiceResp *BlsAggregationServiceResponse) toNonSignerStakesAndSignature() avssm.IBLSSignatureCheckerTypesNonSignerStakesAndSignature {
	nonSignerPubkeys := []avssm.BN254G1Point{}
	for _, nonSignerPubkey := range blsAggServiceResp.NonSignersPubkeysG1 {
		nonSignerPubkeys = append(nonSignerPubkeys, avssm.BN254G1Point(utils.ConvertToBN254G1Point(nonSignerPubkey)))
	}
	quorumApks := []avssm.BN254G1Point{}
	for _, quorumApk := range blsAggServiceResp.QuorumApksG1 {
		quorumApks = append(quorumApks, avssm.BN254G1Point(utils.ConvertToBN254G1Point(quorumApk)))
	}
	nonSignerStakesAndSignature := avssm.IBLSSignatureCheckerTypesNonSignerStakesAndSignature{
		NonSignerPubkeys: nonSignerPubkeys,
		QuorumApks:       quorumApks,
		ApkG2:            avssm.BN254G2Point(utils.ConvertToBN254G2Point(blsAggServiceResp.SignersApkG2)),
		Sigma: avssm.BN254G1Point(
			utils.ConvertToBN254G1Point(blsAggServiceResp.SignersAggSigG1.G1Point),
		),
		NonSignerQuorumBitmapIndices: blsAggServiceResp.NonSignerQuorumBitmapIndices,
		QuorumApkIndices:             blsAggServiceResp.QuorumApkIndices,
		TotalStakeIndices:            blsAggServiceResp.TotalStakeIndices,
		NonSignerStakeIndices:        blsAggServiceResp.NonSignerStakeIndices,
	}
	return nonSignerStakesAndSignature
}
