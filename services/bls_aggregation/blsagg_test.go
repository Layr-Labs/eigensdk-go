package blsagg

import (
	"context"
	"math/big"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/mocks"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testOperator struct {
	operatorAddr   common.Address
	operatorId     types.OperatorId
	stakePerQuorum map[types.QuorumNum]types.StakeAmount
	blsKeypair     *bls.KeyPair
}

func TestBlsAgg(t *testing.T) {
	testOperator1 := testOperator{
		operatorAddr:   common.HexToAddress("0x1"),
		operatorId:     types.OperatorId{1},
		stakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
		blsKeypair:     newBlsKeyPairPanics("0x1"),
	}
	// testOperator2 := testOperator{
	// 	operatorAddr:   common.HexToAddress("0x2"),
	// 	operatorId:     types.OperatorId{2},
	// 	stakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(100), 1: big.NewInt(200)},
	// 	blsKeypair:     newBlsKeyPairPanics("0x2"),
	// }
	// testOperator3 := testOperator{
	// 	operatorAddr:   common.HexToAddress("0x3"),
	// 	operatorId:     types.OperatorId{3},
	// 	stakePerQuorum: map[types.QuorumNum]types.StakeAmount{0: big.NewInt(300), 1: big.NewInt(100)},
	// 	blsKeypair:     newBlsKeyPairPanics("0x3"),
	// }

	// blockNum can be shared between tests because it doesn't affect anything
	blockNum := uint32(1)

	var tests = []struct {
		name                            string
		mocksInitializationFunc         func(*mocks.MockAvsRegistryService)
		performOperationsAndReturnWants func(blsAggServ *BlsAggregationService) (wantAggregationServiceResponse BlsAggretationServiceResponse)
	}{
		{
			name: "1 quorum 1 operator 1 correct signature",
			mocksInitializationFunc: func(mockAvsRegistryService *mocks.MockAvsRegistryService) {
				mockAvsRegistryService.EXPECT().GetOperatorsAvsStateAtBlock(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					map[[32]byte]types.OperatorAvsState{
						testOperator1.operatorId: {
							OperatorId: testOperator1.operatorId,
							Pubkeys: types.OperatorPubkeys{
								G1Pubkey: testOperator1.blsKeypair.GetPubKeyG1(),
								G2Pubkey: testOperator1.blsKeypair.GetPubKeyG2(),
							},
							StakePerQuorum: testOperator1.stakePerQuorum,
							BlockNumber:    blockNum,
						},
					},
					nil,
				)
				mockAvsRegistryService.EXPECT().GetQuorumsAvsStateAtBlock(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					map[types.QuorumNum]types.QuorumAvsState{
						0: {
							QuorumNumber: 0,
							AggPubkeyG1:  testOperator1.blsKeypair.GetPubKeyG1(),
							TotalStake:   testOperator1.stakePerQuorum[0],
							BlockNumber:  blockNum,
						},
						1: {
							QuorumNumber: 1,
							AggPubkeyG1:  testOperator1.blsKeypair.GetPubKeyG1(),
							TotalStake:   testOperator1.stakePerQuorum[1],
							BlockNumber:  blockNum,
						},
					},
					nil,
				)
			},
			performOperationsAndReturnWants: func(blsAggServ *BlsAggregationService) (wantAggregationServiceResponse BlsAggretationServiceResponse) {
				taskIndex := types.TaskIndex(0)
				quorumNumbers := []types.QuorumNum{0}
				quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
				taskResponseHash := types.TaskResponseDigest{123}
				blsSig := testOperator1.blsKeypair.SignMessage(taskResponseHash)

				blsAggServ.InitializeNewTask(context.Background(), taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages)
				err := blsAggServ.ProcessNewSignature(taskIndex, taskResponseHash, blsSig, testOperator1.operatorId)
				assert.Nil(t, err)
				return BlsAggretationServiceResponse{
					err:                 nil,
					NonSignersPubkeysG1: []*bls.G1Point{},
					QuorumApksG1:        []*bls.G1Point{testOperator1.blsKeypair.GetPubKeyG1()},
					SignersApkG2:        testOperator1.blsKeypair.GetPubKeyG2(),
					SignersAggSigG1:     testOperator1.blsKeypair.SignMessage(taskResponseHash),
				}
			},
		},
		{
			name: "1 quorum 1 operator 0 signatures - task expired",
			mocksInitializationFunc: func(mockAvsRegistryService *mocks.MockAvsRegistryService) {
				mockAvsRegistryService.EXPECT().GetOperatorsAvsStateAtBlock(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					map[[32]byte]types.OperatorAvsState{
						testOperator1.operatorId: {
							OperatorId: testOperator1.operatorId,
							Pubkeys: types.OperatorPubkeys{
								G1Pubkey: testOperator1.blsKeypair.GetPubKeyG1(),
								G2Pubkey: testOperator1.blsKeypair.GetPubKeyG2(),
							},
							StakePerQuorum: testOperator1.stakePerQuorum,
							BlockNumber:    blockNum,
						},
					},
					nil,
				)
				mockAvsRegistryService.EXPECT().GetQuorumsAvsStateAtBlock(gomock.Any(), gomock.Any(), gomock.Any()).Return(
					map[types.QuorumNum]types.QuorumAvsState{
						0: {
							QuorumNumber: 0,
							AggPubkeyG1:  testOperator1.blsKeypair.GetPubKeyG1(),
							TotalStake:   testOperator1.stakePerQuorum[0],
							BlockNumber:  blockNum,
						},
						1: {
							QuorumNumber: 1,
							AggPubkeyG1:  testOperator1.blsKeypair.GetPubKeyG1(),
							TotalStake:   testOperator1.stakePerQuorum[1],
							BlockNumber:  blockNum,
						},
					},
					nil,
				)
			},
			performOperationsAndReturnWants: func(blsAggServ *BlsAggregationService) (wantAggregationServiceResponse BlsAggretationServiceResponse) {
				taskIndex := types.TaskIndex(0)
				quorumNumbers := []types.QuorumNum{0}
				quorumThresholdPercentages := []types.QuorumThresholdPercentage{100}
				blsAggServ.InitializeNewTask(context.Background(), taskIndex, blockNum, quorumNumbers, quorumThresholdPercentages)
				return BlsAggretationServiceResponse{
					err:                 taskExpiredError,
					NonSignersPubkeysG1: nil,
					QuorumApksG1:        nil,
					SignersApkG2:        nil,
					SignersAggSigG1:     nil,
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryService := mocks.NewMockAvsRegistryService(mockCtrl)
			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryService)
			}

			noopLogger := logging.NewNoopLogger()
			blsAggServ := NewBlsAggregationService(mockAvsRegistryService, noopLogger)

			wantAggregationServiceResponse := tt.performOperationsAndReturnWants(blsAggServ)
			gotAggregationServiceResponse := <-blsAggServ.ResponseC
			assert.Equal(t, wantAggregationServiceResponse, gotAggregationServiceResponse)
		})
	}
}

func newBlsKeyPairPanics(hexKey string) *bls.KeyPair {
	keypair, err := bls.NewKeyPairFromString(hexKey)
	if err != nil {
		panic(err)
	}
	return keypair
}

// TODO:
// - test tasks are deleted after TTR expired
