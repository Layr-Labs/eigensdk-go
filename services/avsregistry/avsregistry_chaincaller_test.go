package avsregistry

import (
	"context"
	"errors"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	opstateretrievar "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type fakeAVSRegistryReader struct {
	filterQueryRange *big.Int
	opAddress        []types.OperatorAddr
	opPubKeys        []types.OperatorPubkeys
	operatorId       types.OperatorId
	socket           types.Socket
	err              error
}

func newFakeAVSRegistryReader(
	opr *testOperator,
	err error,
) *fakeAVSRegistryReader {
	if opr == nil {
		return &fakeAVSRegistryReader{}
	}
	return &fakeAVSRegistryReader{
		opAddress:  []common.Address{opr.operatorAddr},
		opPubKeys:  []types.OperatorPubkeys{opr.operatorInfo.Pubkeys},
		socket:     opr.operatorInfo.Socket,
		operatorId: opr.operatorId,
		err:        err,
	}
}

func (f *fakeAVSRegistryReader) GetOperatorFromId(
	opts *bind.CallOpts,
	operatorId types.OperatorId,
) (common.Address, error) {
	return f.opAddress[0], f.err
}

func (f *fakeAVSRegistryReader) GetOperatorsStakeInQuorumsAtBlock(
	opts *bind.CallOpts,
	quorumNumbers types.QuorumNums,
	blockNumber types.BlockNum,
) ([][]opstateretrievar.OperatorStateRetrieverOperator, error) {
	return [][]opstateretrievar.OperatorStateRetrieverOperator{
		{
			{
				OperatorId: f.operatorId,
				Stake:      big.NewInt(123),
			},
		},
	}, nil
}

func (f *fakeAVSRegistryReader) GetCheckSignaturesIndices(
	opts *bind.CallOpts,
	referenceBlockNumber uint32,
	quorumNumbers types.QuorumNums,
	nonSignerOperatorIds []types.OperatorId,
) (opstateretrievar.OperatorStateRetrieverCheckSignaturesIndices, error) {
	return opstateretrievar.OperatorStateRetrieverCheckSignaturesIndices{}, nil
}

type fakeOperatorInfoService struct {
	operatorInfo types.OperatorInfo
}

func newFakeOperatorInfoService(operatorInfo types.OperatorInfo) *fakeOperatorInfoService {
	return &fakeOperatorInfoService{
		operatorInfo: operatorInfo,
	}
}

func (f *fakeOperatorInfoService) GetOperatorInfo(ctx context.Context, operator common.Address) (operatorInfo types.OperatorInfo, operatorFound bool) {
	return f.operatorInfo, true
}

type testOperator struct {
	operatorAddr common.Address
	operatorId   types.OperatorId
	operatorInfo types.OperatorInfo
}

func TestAvsRegistryServiceChainCaller_getOperatorPubkeys(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator1 := testOperator{
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
		name             string
		operator         *testOperator
		queryOperatorId  types.OperatorId
		wantErr          error
		wantOperatorInfo types.OperatorInfo
	}{
		{
			name:             "should return operator info",
			operator:         &testOperator1,
			queryOperatorId:  testOperator1.operatorId,
			wantErr:          nil,
			wantOperatorInfo: testOperator1.operatorInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := newFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.operatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorInfo, gotErr := service.getOperatorInfo(context.Background(), tt.queryOperatorId)
			if !errors.Is(gotErr, tt.wantErr) {
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
	testOperator1 := testOperator{
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
		queryQuorumNumbers        types.QuorumNums
		queryBlockNum             types.BlockNum
		wantErr                   error
		wantOperatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState
		operator                  *testOperator
	}{
		{
			name:               "should return operatorsAvsState",
			queryQuorumNumbers: types.QuorumNums{1},
			operator:           &testOperator1,
			queryBlockNum:      1,
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.OperatorId]types.OperatorAvsState{
				testOperator1.operatorId: {
					OperatorId:     testOperator1.operatorId,
					OperatorInfo:   testOperator1.operatorInfo,
					StakePerQuorum: map[types.QuorumNum]types.StakeAmount{1: big.NewInt(123)},
					BlockNumber:    1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := newFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.operatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorsAvsStateDict, gotErr := service.GetOperatorsAvsStateAtBlock(context.Background(), tt.queryQuorumNumbers, tt.queryBlockNum)
			if !errors.Is(gotErr, tt.wantErr) {
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
	testOperator1 := testOperator{
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
		queryQuorumNumbers      types.QuorumNums
		queryBlockNum           types.BlockNum
		wantErr                 error
		wantQuorumsAvsStateDict map[types.QuorumNum]types.QuorumAvsState
		operator                *testOperator
	}{
		{
			name:               "should return operatorsAvsState",
			queryQuorumNumbers: types.QuorumNums{1},
			operator:           &testOperator1,
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
			mockAvsRegistryReader := newFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.operatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			aggG1PubkeyPerQuorum, gotErr := service.GetQuorumsAvsStateAtBlock(context.Background(), tt.queryQuorumNumbers, tt.queryBlockNum)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantQuorumsAvsStateDict, aggG1PubkeyPerQuorum) {
				t.Fatalf("GetOperatorsAvsState returned wrong aggG1PubkeyPerQuorum. Got: %v, want: %v.", aggG1PubkeyPerQuorum, tt.wantQuorumsAvsStateDict)
			}
		})
	}
}
