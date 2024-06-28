package avsregistry

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	chainiomocks "github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	opstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	servicemocks "github.com/Layr-Labs/eigensdk-go/services/mocks"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/mock/gomock"
)

type testOperator struct {
	operatorAddr common.Address
	operatorId   types.OperatorId
	operatorInfo types.OperatorInfo
}

func TestAvsRegistryServiceChainCaller_getOperatorPubkeys(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorId:   types.OperatorId{1},
		operatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
			},
			Socket: "localhost:8080",
		},
	}

	// TODO(samlaf): add error test cases
	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*chainiomocks.MockAVSReader, *servicemocks.MockOperatorsInfoService)
		queryOperatorId         types.OperatorId
		wantErr                 error
		wantOperatorInfo        types.OperatorInfo
	}{
		{
			name: "should return operator info",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAVSReader, mockOperatorsInfoService *servicemocks.MockOperatorsInfoService) {
				mockAvsRegistryReader.EXPECT().GetOperatorFromId(gomock.Any(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockOperatorsInfoService.EXPECT().GetOperatorInfo(gomock.Any(), testOperator.operatorAddr).Return(testOperator.operatorInfo, true)
			},
			queryOperatorId:  testOperator.operatorId,
			wantErr:          nil,
			wantOperatorInfo: testOperator.operatorInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryReader := chainiomocks.NewMockAVSReader(mockCtrl)
			mockOperatorsInfoService := servicemocks.NewMockOperatorsInfoService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockOperatorsInfoService)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorInfo, gotErr := service.getOperatorInfo(context.Background(), tt.queryOperatorId)
			if tt.wantErr != gotErr {
				t.Fatalf("GetOperatorPubkeys returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorInfo, gotOperatorInfo) {
				t.Fatalf("GetOperatorPubkeys returned wrong operator pubkeys. Got: %v, want: %v.", gotOperatorInfo, tt.wantOperatorInfo)
			}
		})
	}
}

func TestAvsRegistryServiceChainCaller_GetOperatorsAvsState(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorId:   types.OperatorId{1},
		operatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
			},
			Socket: "localhost:8080",
		},
	}

	var tests = []struct {
		name                      string
		mocksInitializationFunc   func(*chainiomocks.MockAVSReader, *servicemocks.MockOperatorsInfoService)
		queryQuorumNumbers        types.QuorumNums
		queryBlockNum             types.BlockNum
		wantErr                   error
		wantOperatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAVSReader, mockOperatorsInfoService *servicemocks.MockOperatorsInfoService) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(gomock.Any(), types.QuorumNums{1}, types.BlockNum(1)).Return([][]opstateretrievar.OperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
				mockAvsRegistryReader.EXPECT().GetOperatorFromId(gomock.Any(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockOperatorsInfoService.EXPECT().GetOperatorInfo(gomock.Any(), testOperator.operatorAddr).Return(testOperator.operatorInfo, true)
			},
			queryQuorumNumbers: types.QuorumNums{1},
			queryBlockNum:      1,
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.OperatorId]types.OperatorAvsState{
				testOperator.operatorId: {
					OperatorId:     testOperator.operatorId,
					OperatorInfo:   testOperator.operatorInfo,
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
			mockAvsRegistryReader := chainiomocks.NewMockAVSReader(mockCtrl)
			mockOperatorsInfoService := servicemocks.NewMockOperatorsInfoService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockOperatorsInfoService)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

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
		operatorId:   types.OperatorId{1},
		operatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
			},
			Socket: "localhost:8080",
		},
	}

	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*chainiomocks.MockAVSReader, *servicemocks.MockOperatorsInfoService)
		queryQuorumNumbers      types.QuorumNums
		queryBlockNum           types.BlockNum
		wantErr                 error
		wantQuorumsAvsStateDict map[types.QuorumNum]types.QuorumAvsState
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAVSReader, mockOperatorsInfoService *servicemocks.MockOperatorsInfoService) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(gomock.Any(), types.QuorumNums{1}, types.BlockNum(1)).Return([][]opstateretrievar.OperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
				mockAvsRegistryReader.EXPECT().GetOperatorFromId(gomock.Any(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockOperatorsInfoService.EXPECT().GetOperatorInfo(gomock.Any(), testOperator.operatorAddr).Return(testOperator.operatorInfo, true)
			},
			queryQuorumNumbers: types.QuorumNums{1},
			queryBlockNum:      1,
			wantErr:            nil,
			wantQuorumsAvsStateDict: map[types.QuorumNum]types.QuorumAvsState{
				1: {
					QuorumNumber: types.QuorumNum(1),
					TotalStake:   big.NewInt(123),
					AggPubkeyG1:  bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
					BlockNumber:  1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryReader := chainiomocks.NewMockAVSReader(mockCtrl)
			mockOperatorsInfoService := servicemocks.NewMockOperatorsInfoService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockOperatorsInfoService)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

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
