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
	pubkeys      types.OperatorPubkeys
}

func TestAvsRegistryServiceChainCaller_getOperatorPubkeys(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorId:   types.OperatorId{1},
		pubkeys: types.OperatorPubkeys{
			G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
			G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
		},
	}

	// TODO(samlaf): add error test cases
	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*chainiomocks.MockAvsRegistryReader, *servicemocks.MockOperatorPubkeysService)
		queryOperatorId         types.OperatorId
		wantErr                 error
		wantOperatorPubkeys     types.OperatorPubkeys
	}{
		{
			name: "should return operatorpubkeys",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsRegistryReader, mockOperatorPubkeysService *servicemocks.MockOperatorPubkeysService) {
				mockAvsRegistryReader.EXPECT().GetOperatorFromId(gomock.Any(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockOperatorPubkeysService.EXPECT().GetOperatorPubkeys(gomock.Any(), testOperator.operatorAddr).Return(testOperator.pubkeys, true)
			},
			queryOperatorId:     testOperator.operatorId,
			wantErr:             nil,
			wantOperatorPubkeys: testOperator.pubkeys,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryReader := chainiomocks.NewMockAvsRegistryReader(mockCtrl)
			mockOperatorPubkeysService := servicemocks.NewMockOperatorPubkeysService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockOperatorPubkeysService)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorPubkeysService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorPubkeys, gotErr := service.getOperatorPubkeys(context.Background(), tt.queryOperatorId)
			if tt.wantErr != gotErr {
				t.Fatalf("GetOperatorPubkeys returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorPubkeys, gotOperatorPubkeys) {
				t.Fatalf("GetOperatorPubkeys returned wrong operator pubkeys. Got: %v, want: %v.", gotOperatorPubkeys, tt.wantOperatorPubkeys)
			}
		})
	}
}

func TestAvsRegistryServiceChainCaller_GetOperatorsAvsState(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorId:   types.OperatorId{1},
		pubkeys: types.OperatorPubkeys{
			G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
			G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
		},
	}

	var tests = []struct {
		name                      string
		mocksInitializationFunc   func(*chainiomocks.MockAvsRegistryReader, *servicemocks.MockOperatorPubkeysService)
		queryQuorumNumbers        []types.QuorumNum
		queryBlockNum             types.BlockNum
		wantErr                   error
		wantOperatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsRegistryReader, mockOperatorPubkeysService *servicemocks.MockOperatorPubkeysService) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(gomock.Any(), []types.QuorumNum{1}, types.BlockNum(1)).Return([][]opstateretrievar.OperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
				mockAvsRegistryReader.EXPECT().GetOperatorFromId(gomock.Any(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockOperatorPubkeysService.EXPECT().GetOperatorPubkeys(gomock.Any(), testOperator.operatorAddr).Return(testOperator.pubkeys, true)
			},
			queryQuorumNumbers: []types.QuorumNum{1},
			queryBlockNum:      1,
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.OperatorId]types.OperatorAvsState{
				testOperator.operatorId: {
					OperatorId:     testOperator.operatorId,
					Pubkeys:        testOperator.pubkeys,
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
			mockAvsRegistryReader := chainiomocks.NewMockAvsRegistryReader(mockCtrl)
			mockOperatorPubkeysService := servicemocks.NewMockOperatorPubkeysService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockOperatorPubkeysService)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorPubkeysService, logger)

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
		pubkeys: types.OperatorPubkeys{
			G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
			G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
		},
	}

	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*chainiomocks.MockAvsRegistryReader, *servicemocks.MockOperatorPubkeysService)
		queryQuorumNumbers      []types.QuorumNum
		queryBlockNum           types.BlockNum
		wantErr                 error
		wantQuorumsAvsStateDict map[types.QuorumNum]types.QuorumAvsState
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsRegistryReader, mockOperatorPubkeysService *servicemocks.MockOperatorPubkeysService) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(gomock.Any(), []types.QuorumNum{1}, types.BlockNum(1)).Return([][]opstateretrievar.OperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
				mockAvsRegistryReader.EXPECT().GetOperatorFromId(gomock.Any(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockOperatorPubkeysService.EXPECT().GetOperatorPubkeys(gomock.Any(), testOperator.operatorAddr).Return(testOperator.pubkeys, true)
			},
			queryQuorumNumbers: []types.QuorumNum{1},
			queryBlockNum:      1,
			wantErr:            nil,
			wantQuorumsAvsStateDict: map[types.QuorumNum]types.QuorumAvsState{
				1: types.QuorumAvsState{
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
			mockAvsRegistryReader := chainiomocks.NewMockAvsRegistryReader(mockCtrl)
			mockOperatorPubkeysService := servicemocks.NewMockOperatorPubkeysService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockOperatorPubkeysService)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorPubkeysService, logger)

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
