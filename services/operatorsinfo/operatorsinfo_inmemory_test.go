package operatorsinfo

import (
	"context"
	"log/slog"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/chainio/mocks"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/mock/gomock"

	apkregistrybindings "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	regcoordbindings "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
)

type testOperator struct {
	operatorAddr     common.Address
	operatorInfo     types.OperatorInfo
	contractG1Pubkey apkregistrybindings.BN254G1Point
	contractG2Pubkey apkregistrybindings.BN254G2Point
}

func TestGetOperatorPubkeys(t *testing.T) {
	logger := logging.NewTextSLogger(os.Stdout, &logging.SLoggerOptions{Level: slog.LevelDebug})
	operator1Pubkeys := types.OperatorPubkeys{
		G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
		G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
	}
	contractG1Pubkey, contractG2Pubkey := operator1Pubkeys.ToContractPubkeys()
	testOperator1 := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		operatorInfo: types.OperatorInfo{
			Pubkeys: operator1Pubkeys,
			Socket:  "localhost:8080",
		},
		contractG1Pubkey: contractG1Pubkey,
		contractG2Pubkey: contractG2Pubkey,
	}

	// Define tests
	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*mocks.MockAvsRegistrySubscriber, *mocks.MockAvsRegistryReader, *mocks.MockSubscription)
		queryOperatorAddr       common.Address
		wantOperatorFound       bool
		wantOperatorInfo        types.OperatorInfo
	}{
		{
			name: "should return false if operator not found",
			mocksInitializationFunc: func(mockAvsRegistrySubscriber *mocks.MockAvsRegistrySubscriber, mockAvsReader *mocks.MockAvsRegistryReader, mockSubscription *mocks.MockSubscription) {
				errC := make(chan error)
				mockSubscription.EXPECT().Err().AnyTimes().Return(errC)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToNewPubkeyRegistrations().Return(nil, mockSubscription, nil)
				mockAvsReader.EXPECT().QueryExistingRegisteredOperatorPubKeys(gomock.Any(), nil, nil).Return(nil, nil, nil)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToOperatorSocketUpdates().Return(nil, mockSubscription, nil)
				mockAvsReader.EXPECT().QueryExistingRegisteredOperatorSockets(gomock.Any(), nil, nil).Return(nil, nil)
			},
			queryOperatorAddr: testOperator1.operatorAddr,
			wantOperatorFound: false,
			wantOperatorInfo:  types.OperatorInfo{},
		},
		{
			name: "should return operator pubkeys found via query",
			mocksInitializationFunc: func(mockAvsRegistrySubscriber *mocks.MockAvsRegistrySubscriber, mockAvsReader *mocks.MockAvsRegistryReader, mockSubscription *mocks.MockSubscription) {
				errC := make(chan error)
				mockSubscription.EXPECT().Err().AnyTimes().Return(errC)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToNewPubkeyRegistrations().Return(nil, mockSubscription, nil)
				mockAvsReader.EXPECT().QueryExistingRegisteredOperatorPubKeys(gomock.Any(), nil, nil).
					Return([]common.Address{testOperator1.operatorAddr}, []types.OperatorPubkeys{testOperator1.operatorInfo.Pubkeys}, nil)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToOperatorSocketUpdates().Return(nil, mockSubscription, nil)
				mockAvsReader.EXPECT().QueryExistingRegisteredOperatorSockets(gomock.Any(), nil, nil).
					Return(map[types.OperatorId]types.Socket{
						types.OperatorIdFromG1Pubkey(testOperator1.operatorInfo.Pubkeys.G1Pubkey): testOperator1.operatorInfo.Socket,
					}, nil)
			},
			queryOperatorAddr: testOperator1.operatorAddr,
			wantOperatorFound: true,
			wantOperatorInfo:  testOperator1.operatorInfo,
		},
		{
			name: "should return operator pubkeys found via subscription",
			mocksInitializationFunc: func(mockAvsRegistrySubscriber *mocks.MockAvsRegistrySubscriber, mockAvsReader *mocks.MockAvsRegistryReader, mockSubscription *mocks.MockSubscription) {
				errC := make(chan error)
				pubkeyRegistrationEventC := make(chan *apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration, 1)
				pubkeyRegistrationEvent := &apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration{
					Operator: testOperator1.operatorAddr,
					PubkeyG1: testOperator1.contractG1Pubkey,
					PubkeyG2: testOperator1.contractG2Pubkey,
					Raw:      gethtypes.Log{},
				}
				pubkeyRegistrationEventC <- pubkeyRegistrationEvent
				operatorSocketUpdateEventC := make(chan *regcoordbindings.ContractRegistryCoordinatorOperatorSocketUpdate, 1)
				operatorSocketUpdateEvent := &regcoordbindings.ContractRegistryCoordinatorOperatorSocketUpdate{
					OperatorId: types.OperatorIdFromG1Pubkey(testOperator1.operatorInfo.Pubkeys.G1Pubkey),
					Socket:     string(testOperator1.operatorInfo.Socket),
					Raw:        gethtypes.Log{},
				}
				operatorSocketUpdateEventC <- operatorSocketUpdateEvent
				mockSubscription.EXPECT().Err().AnyTimes().Return(errC)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToNewPubkeyRegistrations().Return(pubkeyRegistrationEventC, mockSubscription, nil)
				mockAvsReader.EXPECT().QueryExistingRegisteredOperatorPubKeys(gomock.Any(), nil, nil).
					Return([]common.Address{}, []types.OperatorPubkeys{}, nil)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToOperatorSocketUpdates().Return(operatorSocketUpdateEventC, mockSubscription, nil)
				mockAvsReader.EXPECT().QueryExistingRegisteredOperatorSockets(gomock.Any(), nil, nil).Return(nil, nil)
			},
			queryOperatorAddr: testOperator1.operatorAddr,
			wantOperatorFound: true,
			wantOperatorInfo:  testOperator1.operatorInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistrySubscriber := mocks.NewMockAvsRegistrySubscriber(mockCtrl)
			mockAvsReader := mocks.NewMockAvsRegistryReader(mockCtrl)
			mockSubscription := mocks.NewMockSubscription(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistrySubscriber, mockAvsReader, mockSubscription)
			}
			// Create a new instance of the operatorpubkeys service
			service := NewOperatorsInfoServiceInMemory(context.Background(), mockAvsRegistrySubscriber, mockAvsReader, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorsInfo, gotOperatorFound := service.GetOperatorInfo(context.Background(), tt.queryOperatorAddr)
			if tt.wantOperatorFound != gotOperatorFound {
				t.Fatalf("GetOperatorPubkeys returned wrong ok. Got: %v, want: %v.", gotOperatorFound, tt.wantOperatorFound)
			}
			if tt.wantOperatorFound == true && !reflect.DeepEqual(tt.wantOperatorInfo, gotOperatorsInfo) {
				t.Fatalf("GetOperatorPubkeys returned wrong operator pubkeys. Got: %v, want: %v.", gotOperatorsInfo, tt.wantOperatorInfo)
			}
		})
	}
}
