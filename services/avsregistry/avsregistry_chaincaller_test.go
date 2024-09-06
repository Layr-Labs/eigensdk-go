package avsregistry

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"math/big"
	"reflect"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/internal/fakes"
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

func TestAvsRegistryServiceChainCaller_getOperatorPubkeys(t *testing.T) {
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

	// TODO(samlaf): add error test cases
	var tests = []struct {
		name             string
		operator         *fakes.TestOperator
		queryOperatorId  types.OperatorId
		wantErr          error
		wantOperatorInfo types.OperatorInfo
	}{
		{
			name:             "should return operator info",
			operator:         &testOperator1,
			queryOperatorId:  testOperator1.OperatorId,
			wantErr:          nil,
			wantOperatorInfo: testOperator1.OperatorInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := fakes.NewFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.OperatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorInfo, gotErr := service.getOperatorInfo(context.Background(), tt.queryOperatorId)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorPubkeys returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorInfo, gotOperatorInfo) {
				t.Fatalf(
					"GetOperatorPubkeys returned wrong operator pubkeys. Got: %v, want: %v.",
					gotOperatorInfo,
					tt.wantOperatorInfo,
				)
			}
		})
	}
}

type G2Point struct {
	X [2]int64 `json:"x"`
	Y [2]int64 `json:"y"`
}
type G1Point struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type BlsPublicKey struct {
	G1Pubkey G1Point `json:"g1_pub_key"`
	G2Pubkey G2Point `json:"g2_pub_key"`
}

type TestOperator struct {
	OperatorAddr string `json:"operator_addr"`
	OperatorId   string `json:"operator_id"`
	OperatorInfo struct {
		Pubkeys BlsPublicKey `json:"pubkeys"`
		Socket  string       `json:"socket"`
	} `json:"operator_info"`
}

type TestData struct {
	Input  Input  `json:"input"`
	Output Output `json:"output"`
}

type Input struct {
	QueryQuorumNumbers []int        `json:"query_quorum_numbers"`
	QueryBlockNum      int32        `json:"query_block_num"`
	Operator           TestOperator `json:"operator"`
}

type Output struct {
	OperatorsAvsState []struct {
		OperatorAddr   string `json:"operator_addr"`
		OperatorId     string `json:"operator_id"`
		StakePerQuorum struct {
			Quorum int   `json:"quorum"`
			Stake  int64 `json:"stake"`
		} `json:"stake_per_quorum"`
		BlockNumber int32 `json:"block_number"`
	} `json:"operators_avs_state"`
}

// converts a hex string (starting with "0x") into a Bytes32.
func NewBytes32FromString(hexString string) [32]byte {
	var b32 [32]byte

	// Remove the "0x" prefix if it's present
	if len(hexString) >= 2 && hexString[:2] == "0x" {
		hexString = hexString[2:]
	}

	// Decode the hex string
	bytes, _ := hex.DecodeString(hexString)

	// Copy the bytes into the Bytes32 array
	copy(b32[:], bytes)
	return b32
}

func TestAvsRegistryServiceChainCaller_GetOperatorsAvsState(t *testing.T) {
	test_data_path := os.Getenv("TEST_DATA_PATH")
	data, err := ioutil.ReadFile(test_data_path)
	if err != nil {
		t.Fatalf("Failed to read JSON file for test %s: %v", t.Name(), err)
	}

	var testData TestData
	if err := json.Unmarshal(data, &testData); err != nil {
		t.Fatalf("Failed to unmarshal JSON data for test %s: %v", t.Name(), err)
	}

	logger := testutils.GetTestLogger()
	pubkeys := testData.Input.Operator.OperatorInfo.Pubkeys
	testOperator1 := fakes.TestOperator{
		OperatorAddr: common.HexToAddress(testData.Input.Operator.OperatorAddr),
		OperatorId:   NewBytes32FromString(testData.Input.Operator.OperatorId),
		OperatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(pubkeys.G1Pubkey.X), big.NewInt(pubkeys.G1Pubkey.Y)),
				G2Pubkey: bls.NewG2Point(
					[2]*big.Int{big.NewInt(pubkeys.G2Pubkey.X[0]), big.NewInt(pubkeys.G2Pubkey.X[1])},
					[2]*big.Int{big.NewInt(pubkeys.G2Pubkey.Y[0]), big.NewInt(pubkeys.G2Pubkey.Y[1])},
				),
			},
			Socket: types.Socket(testData.Input.Operator.OperatorInfo.Socket),
		},
	}

	outputOperator := testData.Output.OperatorsAvsState[0]
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
			queryQuorumNumbers: types.QuorumNums{types.QuorumNum(testData.Input.QueryQuorumNumbers[0])},
			operator:           &testOperator1,
			queryBlockNum:      uint32(testData.Input.QueryBlockNum),
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.OperatorId]types.OperatorAvsState{
				NewBytes32FromString(outputOperator.OperatorId): {
					OperatorId:     NewBytes32FromString(outputOperator.OperatorId),
					OperatorInfo:   testOperator1.OperatorInfo,
					StakePerQuorum: map[types.QuorumNum]types.StakeAmount{types.QuorumNum(outputOperator.StakePerQuorum.Quorum): big.NewInt(outputOperator.StakePerQuorum.Stake)},
					BlockNumber:    uint32(outputOperator.BlockNumber),
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
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

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

func TestAvsRegistryServiceChainCaller_GetQuorumsAvsState(t *testing.T) {
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
		name                    string
		queryQuorumNumbers      types.QuorumNums
		queryBlockNum           types.BlockNum
		wantErr                 error
		wantQuorumsAvsStateDict map[types.QuorumNum]types.QuorumAvsState
		operator                *fakes.TestOperator
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
			mockAvsRegistryReader := fakes.NewFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.OperatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			aggG1PubkeyPerQuorum, gotErr := service.GetQuorumsAvsStateAtBlock(
				context.Background(),
				tt.queryQuorumNumbers,
				tt.queryBlockNum,
			)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantQuorumsAvsStateDict, aggG1PubkeyPerQuorum) {
				t.Fatalf(
					"GetOperatorsAvsState returned wrong aggG1PubkeyPerQuorum. Got: %v, want: %v.",
					aggG1PubkeyPerQuorum,
					tt.wantQuorumsAvsStateDict,
				)
			}
		})
	}
}
