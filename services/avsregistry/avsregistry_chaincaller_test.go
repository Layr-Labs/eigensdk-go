package avsregistry

import (
	"context"
	"math/big"
	"reflect"
	"testing"

	chainiomocks "github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	blsoperatorstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
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

func TestAvsRegistryServiceChainCaller_GetOperatorPubkeys(t *testing.T) {
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
		mocksInitializationFunc func(*chainiomocks.MockAvsRegistryReader, *chainiomocks.MockELReader, *servicemocks.MockPubkeyCompendiumService)
		queryOperatorId         types.OperatorId
		wantErr                 error
		wantOperatorPubkeys     types.OperatorPubkeys
	}{
		{
			name: "should return operatorpubkeys",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsRegistryReader, mockElReader *chainiomocks.MockELReader, mockPubkeyCompendiumService *servicemocks.MockPubkeyCompendiumService) {
				mockElReader.EXPECT().GetOperatorAddressFromPubkeyHash(context.Background(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockPubkeyCompendiumService.EXPECT().GetOperatorPubkeys(context.Background(), testOperator.operatorAddr).Return(testOperator.pubkeys, true)
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
			mockElReader := chainiomocks.NewMockELReader(mockCtrl)
			mockPubkeyCompendium := servicemocks.NewMockPubkeyCompendiumService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockElReader, mockPubkeyCompendium)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockElReader, mockPubkeyCompendium, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorPubkeys, gotErr := service.GetOperatorPubkeys(context.Background(), tt.queryOperatorId)
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
		mocksInitializationFunc   func(*chainiomocks.MockAvsRegistryReader, *chainiomocks.MockELReader, *servicemocks.MockPubkeyCompendiumService)
		queryQuorumNumbers        []types.QuorumNum
		queryBlockNum             types.BlockNumber
		wantErr                   error
		wantOperatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState
		wantAggG1PubkeyPerQuorum  map[types.QuorumNum]*bls.G1Point
	}{
		{
			name: "should return operatorsAvsState",
			mocksInitializationFunc: func(mockAvsRegistryReader *chainiomocks.MockAvsRegistryReader, mockElReader *chainiomocks.MockELReader, mockPubkeyCompendiumService *servicemocks.MockPubkeyCompendiumService) {
				mockAvsRegistryReader.EXPECT().GetOperatorsStakeInQuorumsAtBlock(context.Background(), []types.QuorumNum{1}, types.BlockNumber(1)).Return([][]blsoperatorstateretrievar.BLSOperatorStateRetrieverOperator{
					{
						{
							OperatorId: testOperator.operatorId,
							Stake:      big.NewInt(123),
						},
					},
				}, nil)
				mockElReader.EXPECT().GetOperatorAddressFromPubkeyHash(context.Background(), testOperator.operatorId).Return(testOperator.operatorAddr, nil)
				mockPubkeyCompendiumService.EXPECT().GetOperatorPubkeys(context.Background(), testOperator.operatorAddr).Return(testOperator.pubkeys, true)
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
			wantAggG1PubkeyPerQuorum: map[types.QuorumNum]*bls.G1Point{
				1: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistryReader := chainiomocks.NewMockAvsRegistryReader(mockCtrl)
			mockElReader := chainiomocks.NewMockELReader(mockCtrl)
			mockPubkeyCompendium := servicemocks.NewMockPubkeyCompendiumService(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistryReader, mockElReader, mockPubkeyCompendium)
			}
			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockElReader, mockPubkeyCompendium, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorsAvsStateDict, aggG1PubkeyPerQuorum, gotErr := service.GetOperatorsAvsStateAtBlock(context.Background(), tt.queryQuorumNumbers, tt.queryBlockNum)
			if tt.wantErr != gotErr {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorsAvsStateDict, gotOperatorsAvsStateDict) {
				t.Fatalf("GetOperatorsAvsState returned wrong operatorsAvsStateDict. Got: %v, want: %v.", gotOperatorsAvsStateDict, tt.wantOperatorsAvsStateDict)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantAggG1PubkeyPerQuorum, aggG1PubkeyPerQuorum) {
				t.Fatalf("GetOperatorsAvsState returned wrong aggG1PubkeyPerQuorum. Got: %v, want: %v.", aggG1PubkeyPerQuorum, tt.wantAggG1PubkeyPerQuorum)
			}
		})
	}
}
