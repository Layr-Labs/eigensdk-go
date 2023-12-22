package operatorpubkeys

import (
	"context"
	"math/big"
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
)

type testOperator struct {
	operatorAddr     common.Address
	pubkeys          types.OperatorPubkeys
	contractG1Pubkey apkregistrybindings.BN254G1Point
	contractG2Pubkey apkregistrybindings.BN254G2Point
}

func TestGetOperatorPubkeys(t *testing.T) {
	logger := logging.NewNoopLogger()
	testOperator1 := testOperator{
		operatorAddr: common.HexToAddress("0x1"),
		pubkeys: types.OperatorPubkeys{
			G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
			G2Pubkey: bls.NewG2Point([2]*big.Int{big.NewInt(1), big.NewInt(1)}, [2]*big.Int{big.NewInt(1), big.NewInt(1)}),
		},
		contractG1Pubkey: apkregistrybindings.BN254G1Point{
			X: big.NewInt(1),
			Y: big.NewInt(1),
		},
		contractG2Pubkey: apkregistrybindings.BN254G2Point{
			X: [2]*big.Int{big.NewInt(1), big.NewInt(1)},
			Y: [2]*big.Int{big.NewInt(1), big.NewInt(1)},
		},
	}

	// Define tests
	var tests = []struct {
		name                    string
		mocksInitializationFunc func(*mocks.MockAvsRegistrySubscriber, *mocks.MockAvsRegistryReader, *mocks.MockSubscription)
		queryOperatorAddr       common.Address
		wantOperatorFound       bool
		wantOperatorPubkeys     types.OperatorPubkeys
	}{
		{
			name: "should return false if operator not found",
			mocksInitializationFunc: func(mockAvsRegistrySubscriber *mocks.MockAvsRegistrySubscriber, mockElReader *mocks.MockAvsRegistryReader, mockSubscription *mocks.MockSubscription) {
				errC := make(chan error)
				mockSubscription.EXPECT().Err().AnyTimes().Return(errC)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToNewPubkeyRegistrations().Return(nil, mockSubscription, nil)
				mockElReader.EXPECT().QueryExistingRegisteredOperatorPubKeys(gomock.Any(), nil, nil).Return(nil, nil, nil)
			},
			queryOperatorAddr:   testOperator1.operatorAddr,
			wantOperatorFound:   false,
			wantOperatorPubkeys: types.OperatorPubkeys{},
		},
		{
			name: "should return operator pubkeys found via query",
			mocksInitializationFunc: func(mockAvsRegistrySubscriber *mocks.MockAvsRegistrySubscriber, mockElReader *mocks.MockAvsRegistryReader, mockSubscription *mocks.MockSubscription) {
				errC := make(chan error)
				mockSubscription.EXPECT().Err().AnyTimes().Return(errC)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToNewPubkeyRegistrations().Return(nil, mockSubscription, nil)
				mockElReader.EXPECT().QueryExistingRegisteredOperatorPubKeys(gomock.Any(), nil, nil).
					Return([]common.Address{testOperator1.operatorAddr}, []types.OperatorPubkeys{testOperator1.pubkeys}, nil)
			},
			queryOperatorAddr:   testOperator1.operatorAddr,
			wantOperatorFound:   true,
			wantOperatorPubkeys: testOperator1.pubkeys,
		},
		{
			name: "should return operator pubkeys found via subscription",
			mocksInitializationFunc: func(mockAvsRegistrySubscriber *mocks.MockAvsRegistrySubscriber, mockElReader *mocks.MockAvsRegistryReader, mockSubscription *mocks.MockSubscription) {
				errC := make(chan error)
				pubkeyRegistrationEventC := make(chan *apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration, 1)
				pubkeyRegistrationEvent := &apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration{
					Operator: testOperator1.operatorAddr,
					PubkeyG1: testOperator1.contractG1Pubkey,
					PubkeyG2: testOperator1.contractG2Pubkey,
					Raw:      gethtypes.Log{},
				}
				pubkeyRegistrationEventC <- pubkeyRegistrationEvent
				mockSubscription.EXPECT().Err().AnyTimes().Return(errC)
				mockAvsRegistrySubscriber.EXPECT().SubscribeToNewPubkeyRegistrations().Return(pubkeyRegistrationEventC, mockSubscription, nil)
				mockElReader.EXPECT().QueryExistingRegisteredOperatorPubKeys(gomock.Any(), nil, nil).
					Return([]common.Address{}, []types.OperatorPubkeys{}, nil)
			},
			queryOperatorAddr:   testOperator1.operatorAddr,
			wantOperatorFound:   true,
			wantOperatorPubkeys: testOperator1.pubkeys,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockCtrl := gomock.NewController(t)
			mockAvsRegistrySubscriber := mocks.NewMockAvsRegistrySubscriber(mockCtrl)
			mockElReader := mocks.NewMockAvsRegistryReader(mockCtrl)
			mockSubscription := mocks.NewMockSubscription(mockCtrl)

			if tt.mocksInitializationFunc != nil {
				tt.mocksInitializationFunc(mockAvsRegistrySubscriber, mockElReader, mockSubscription)
			}
			// Create a new instance of the pubkeycompendium service
			service := NewPubkeyCompendiumInMemory(context.Background(), mockAvsRegistrySubscriber, mockElReader, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorPubkeys, gotOperatorFound := service.GetOperatorPubkeys(context.Background(), tt.queryOperatorAddr)
			if tt.wantOperatorFound != gotOperatorFound {
				t.Fatalf("GetOperatorPubkeys returned wrong ok. Got: %v, want: %v.", gotOperatorFound, tt.wantOperatorFound)
			}
			if tt.wantOperatorFound == true && !reflect.DeepEqual(tt.wantOperatorPubkeys, gotOperatorPubkeys) {
				t.Fatalf("GetOperatorPubkeys returned wrong operator pubkeys. Got: %v, want: %v.", gotOperatorPubkeys, tt.wantOperatorPubkeys)
			}
		})
	}
}
