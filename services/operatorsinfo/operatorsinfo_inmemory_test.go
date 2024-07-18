package operatorsinfo

import (
	"context"
	"github.com/ethereum/go-ethereum/event"
	"log/slog"
	"math/big"
	"os"
	"reflect"
	"testing"
	"time"

	apkregistrybindings "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	blsapkreg "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	regcoord "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

type fakeAVSRegistryReader struct {
	opAddress []types.OperatorAddr
	opPubKeys []types.OperatorPubkeys
	socket    types.Socket
	err       error
}

func newFakeAVSRegistryReader(
	opr *testOperator,
	err error,
) *fakeAVSRegistryReader {
	if opr == nil {
		return &fakeAVSRegistryReader{}
	}
	return &fakeAVSRegistryReader{
		opAddress: []common.Address{opr.operatorAddr},
		opPubKeys: []types.OperatorPubkeys{opr.operatorInfo.Pubkeys},
		socket:    opr.operatorInfo.Socket,
		err:       err,
	}
}

func (f *fakeAVSRegistryReader) QueryExistingRegisteredOperatorPubKeys(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) ([]types.OperatorAddr, []types.OperatorPubkeys, error) {
	return f.opAddress, f.opPubKeys, f.err
}

func (f *fakeAVSRegistryReader) QueryExistingRegisteredOperatorSockets(
	ctx context.Context,
	startBlock *big.Int,
	stopBlock *big.Int,
	blockRange *big.Int,
) (map[types.OperatorId]types.Socket, error) {
	if len(f.opPubKeys) == 0 {
		return nil, nil
	}

	return map[types.OperatorId]types.Socket{
		types.OperatorIdFromG1Pubkey(f.opPubKeys[0].G1Pubkey): f.socket,
	}, nil
}

type fakeAVSRegistrySubscriber struct {
	pubkeyRegistrationEventC   chan *apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration
	operatorSocketUpdateEventC chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate
	eventSubscription          *fakeEventSubscription
}

func newFakeAVSRegistrySubscriber(
	eventSubscription *fakeEventSubscription,
	pubkeyRegistrationEventC chan *apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration,
	operatorSocketUpdateEventC chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate,
) *fakeAVSRegistrySubscriber {
	return &fakeAVSRegistrySubscriber{
		pubkeyRegistrationEventC:   pubkeyRegistrationEventC,
		operatorSocketUpdateEventC: operatorSocketUpdateEventC,
		eventSubscription:          eventSubscription,
	}
}

func (f *fakeAVSRegistrySubscriber) SubscribeToNewPubkeyRegistrations() (chan *blsapkreg.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error) {
	return f.pubkeyRegistrationEventC, f.eventSubscription, nil
}

func (f *fakeAVSRegistrySubscriber) SubscribeToOperatorSocketUpdates() (chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate, event.Subscription, error) {
	return f.operatorSocketUpdateEventC, f.eventSubscription, nil
}

type fakeEventSubscription struct {
	errC chan error
}

func newFakeEventSubscription(
	errC chan error) *fakeEventSubscription {
	return &fakeEventSubscription{
		errC: errC,
	}
}

func (f *fakeEventSubscription) Err() <-chan error {
	return f.errC
}

func (f *fakeEventSubscription) Unsubscribe() {

}

type testOperator struct {
	operatorAddr     common.Address
	operatorInfo     types.OperatorInfo
	contractG1Pubkey apkregistrybindings.BN254G1Point
	contractG2Pubkey apkregistrybindings.BN254G2Point
}

func TestGetOperatorInfo(t *testing.T) {
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

	pubkeyRegistrationEventC := make(chan *apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration, 1)
	pubkeyRegistrationEvent := &apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration{
		Operator: testOperator1.operatorAddr,
		PubkeyG1: testOperator1.contractG1Pubkey,
		PubkeyG2: testOperator1.contractG2Pubkey,
		Raw:      gethtypes.Log{},
	}
	pubkeyRegistrationEventC <- pubkeyRegistrationEvent
	operatorSocketUpdateEventC := make(chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate, 1)
	operatorSocketUpdateEvent := &regcoord.ContractRegistryCoordinatorOperatorSocketUpdate{
		OperatorId: types.OperatorIdFromG1Pubkey(testOperator1.operatorInfo.Pubkeys.G1Pubkey),
		Socket:     string(testOperator1.operatorInfo.Socket),
		Raw:        gethtypes.Log{},
	}
	operatorSocketUpdateEventC <- operatorSocketUpdateEvent

	// Define tests
	var tests = []struct {
		name                       string
		operator                   *testOperator
		pubkeyRegistrationEventC   chan *apkregistrybindings.ContractBLSApkRegistryNewPubkeyRegistration
		operatorSocketUpdateEventC chan *regcoord.ContractRegistryCoordinatorOperatorSocketUpdate
		eventErrC                  chan error
		queryOperatorAddr          common.Address
		wantOperatorFound          bool
		wantOperatorInfo           types.OperatorInfo
	}{
		{
			name:              "should return false if operator not found",
			queryOperatorAddr: testOperator1.operatorAddr,
			wantOperatorFound: false,
			wantOperatorInfo:  types.OperatorInfo{},
		},
		{
			name:              "should return operator info found via query",
			operator:          &testOperator1,
			queryOperatorAddr: testOperator1.operatorAddr,
			wantOperatorFound: true,
			wantOperatorInfo:  testOperator1.operatorInfo,
		},
		{
			name:                       "should return operator info found via subscription",
			queryOperatorAddr:          testOperator1.operatorAddr,
			pubkeyRegistrationEventC:   pubkeyRegistrationEventC,
			operatorSocketUpdateEventC: operatorSocketUpdateEventC,
			wantOperatorFound:          true,
			wantOperatorInfo:           testOperator1.operatorInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockSubscription := newFakeEventSubscription(tt.eventErrC)
			mockAvsRegistrySubscriber := newFakeAVSRegistrySubscriber(mockSubscription, tt.pubkeyRegistrationEventC, tt.operatorSocketUpdateEventC)
			mockAvsReader := newFakeAVSRegistryReader(tt.operator, nil)

			// Create a new instance of the operatorpubkeys service
			service := NewOperatorsInfoServiceInMemory(context.Background(), mockAvsRegistrySubscriber, mockAvsReader, nil, logger)
			time.Sleep(2 * time.Second) // need to give it time to process the subscription events.. not sure if there's a better way to do this.

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
