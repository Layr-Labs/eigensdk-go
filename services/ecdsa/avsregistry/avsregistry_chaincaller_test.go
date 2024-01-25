package avsregistry

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	chainiomocks "github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	opstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/ECDSAOperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/mock/gomock"
)

type testOperator struct {
	operatorAddr common.Address
	operatorId   types.EcdsaOperatorId
}

func TestAvsRegistryServiceChainCaller_GetOperatorsAvsState(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorId:   types.EcdsaOperatorId{1},
	}

	var tests = []struct {
		name                      string
		mocksInitializationFunc   func(*chainiomocks.MockAvsEcdsaRegistryReader)
		queryQuorumNumbers        []types.QuorumNum
		queryBlockNum             types.BlockNum
		wantErr                   error
		wantOperatorsAvsStateDict map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsEcdsaRegistryReader) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(gomock.Any(), []types.QuorumNum{1}, types.BlockNum(1)).Return([][]opstateretrievar.ECDSAOperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
			},
			queryQuorumNumbers: []types.QuorumNum{1},
			queryBlockNum:      1,
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.EcdsaOperatorId]types.OperatorEcdsaAvsState{
				testOperator.operatorId: {
					OperatorId:     testOperator.operatorId,
					StakePerQuorum: map[types.QuorumNum]types.StakeAmount{1: big.NewInt(123)},
					BlockNumber:    1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryReader := chainiomocks.NewMockAvsEcdsaRegistryReader(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorsAvsStateDict, gotErr := service.GetOperatorsAvsStateAtBlock(context.Background(), tt.queryQuorumNumbers, tt.queryBlockNum)
			if tt.wantErr != gotErr {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorsAvsStateDict, gotOperatorsAvsStateDict) {
				t.Fatalf("GetOperatorsAvsState returned wrong operatorsAvsStateDict. Got: %v, want: %v.", gotOperatorsAvsStateDict, tt.wantOperatorsAvsStateDict)
			}
		})
	}
}

func TestAvsRegistryServiceChainCaller_GetQuorumsAvsState(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorId:   types.EcdsaOperatorId{1},
	}

	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*chainiomocks.MockAvsEcdsaRegistryReader)
		queryQuorumNumbers      []types.QuorumNum
		queryBlockNum           types.BlockNum
		wantErr                 error
		wantQuorumsAvsStateDict map[types.QuorumNum]types.QuorumEcdsaAvsState
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsEcdsaRegistryReader) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(gomock.Any(), []types.QuorumNum{1}, types.BlockNum(1)).Return([][]opstateretrievar.ECDSAOperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
			},
			queryQuorumNumbers: []types.QuorumNum{1},
			queryBlockNum:      1,
			wantErr:            nil,
			wantQuorumsAvsStateDict: map[types.QuorumNum]types.QuorumEcdsaAvsState{
				1: types.QuorumEcdsaAvsState{
					QuorumNumber: types.QuorumNum(1),
					TotalStake:   big.NewInt(123),
					OperatorIds:  []types.EcdsaOperatorId{testOperator.operatorId},
					BlockNumber:  1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryReader := chainiomocks.NewMockAvsEcdsaRegistryReader(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			aggG1PubkeyPerQuorum, gotErr := service.GetQuorumsAvsStateAtBlock(context.Background(), tt.queryQuorumNumbers, tt.queryBlockNum)
			if tt.wantErr != gotErr {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantQuorumsAvsStateDict, aggG1PubkeyPerQuorum) {
				t.Fatalf("GetOperatorsAvsState returned wrong aggG1PubkeyPerQuorum. Got: %v, want: %v.", aggG1PubkeyPerQuorum, tt.wantQuorumsAvsStateDict)
			}
		})
	}
}
