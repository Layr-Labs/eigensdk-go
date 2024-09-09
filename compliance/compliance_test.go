package compliance

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/internal/fakes"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type fakeOperatorInfoService struct {
	operatorInfo types.OperatorInfo
}

func newFakeOperatorInfoService(operatorInfo types.OperatorInfo) *fakeOperatorInfoService {
	return &fakeOperatorInfoService{
		operatorInfo: operatorInfo,
	}
}

func (f *fakeOperatorInfoService) GetOperatorInfo(
	ctx context.Context,
	operator common.Address,
) (operatorInfo types.OperatorInfo, operatorFound bool) {
	return f.operatorInfo, true
}

type TestData struct {
	Input struct {
		QuorumNumbers types.QuorumNums `json:"quorum_numbers"`
		BlockNum      int32            `json:"block_num"`
	} `json:"input"`
}

func TestCompliance_GetOperatorsAvsState(t *testing.T) {
	test_data_path := os.Getenv("TEST_DATA_PATH")
	data, err := os.ReadFile(test_data_path)
	if err != nil {
		t.Fatalf("Failed to read JSON file for test %s: %v", t.Name(), err)
	}

	var testData TestData
	if err := json.Unmarshal(data, &testData); err != nil {
		t.Fatalf("Failed to unmarshal JSON data for test %s: %v", t.Name(), err)
	}

	logger := testutils.GetTestLogger()

	testOperator1 := fakes.TestOperator{
		OperatorAddr: common.HexToAddress("0x1"),
		OperatorId:   types.OperatorId{1},
		OperatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point(
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
				),
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
		operator                  *fakes.TestOperator
	}{
		{
			name:               "should return operatorsAvsState",
			queryQuorumNumbers: testData.Input.QuorumNumbers,
			operator:           &testOperator1,
			queryBlockNum:      uint32(testData.Input.BlockNum),
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.OperatorId]types.OperatorAvsState{
				testOperator1.OperatorId: {
					OperatorId:     testOperator1.OperatorId,
					OperatorInfo:   testOperator1.OperatorInfo,
					StakePerQuorum: map[types.QuorumNum]types.StakeAmount{testData.Input.QuorumNumbers[0]: big.NewInt(123)},
					BlockNumber:    uint32(testData.Input.BlockNum),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := fakes.NewFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.OperatorInfo)

			// Create a new instance of the avsregistry service
			service := avsregistry.NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorsAvsStateDict, gotErr := service.GetOperatorsAvsStateAtBlock(
				context.Background(),
				tt.queryQuorumNumbers,
				tt.queryBlockNum,
			)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorsAvsStateDict, gotOperatorsAvsStateDict) {
				t.Fatalf(
					"GetOperatorsAvsState returned wrong operatorsAvsStateDict. Got: %v, want: %v.",
					gotOperatorsAvsStateDict,
					tt.wantOperatorsAvsStateDict,
				)
			}
		})
	}
}
