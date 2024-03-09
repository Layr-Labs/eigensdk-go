// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/services/bls_aggregation (interfaces: BlsAggregationService)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/blsagg/blsaggregation.go -package=mocks github.com/Layr-Labs/eigensdk-go/services/bls_aggregation BlsAggregationService
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	bls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	types "github.com/Layr-Labs/eigensdk-go/types"
	gomock "go.uber.org/mock/gomock"
)

// MockBlsAggregationService is a mock of BlsAggregationService interface.
type MockBlsAggregationService struct {
	ctrl     *gomock.Controller
	recorder *MockBlsAggregationServiceMockRecorder
}

// MockBlsAggregationServiceMockRecorder is the mock recorder for MockBlsAggregationService.
type MockBlsAggregationServiceMockRecorder struct {
	mock *MockBlsAggregationService
}

// NewMockBlsAggregationService creates a new mock instance.
func NewMockBlsAggregationService(ctrl *gomock.Controller) *MockBlsAggregationService {
	mock := &MockBlsAggregationService{ctrl: ctrl}
	mock.recorder = &MockBlsAggregationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlsAggregationService) EXPECT() *MockBlsAggregationServiceMockRecorder {
	return m.recorder
}

// GetResponseChannel mocks base method.
func (m *MockBlsAggregationService) GetResponseChannel() <-chan blsagg.BlsAggregationServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseChannel")
	ret0, _ := ret[0].(<-chan blsagg.BlsAggregationServiceResponse)
	return ret0
}

// GetResponseChannel indicates an expected call of GetResponseChannel.
func (mr *MockBlsAggregationServiceMockRecorder) GetResponseChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseChannel", reflect.TypeOf((*MockBlsAggregationService)(nil).GetResponseChannel))
}

// InitializeNewTask mocks base method.
func (m *MockBlsAggregationService) InitializeNewTask(arg0, arg1 uint32, arg2 types.QuorumNums, arg3 types.QuorumThresholdPercentages, arg4 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeNewTask", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeNewTask indicates an expected call of InitializeNewTask.
func (mr *MockBlsAggregationServiceMockRecorder) InitializeNewTask(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeNewTask", reflect.TypeOf((*MockBlsAggregationService)(nil).InitializeNewTask), arg0, arg1, arg2, arg3, arg4)
}

// ProcessNewSignature mocks base method.
func (m *MockBlsAggregationService) ProcessNewSignature(arg0 context.Context, arg1 uint32, arg2 types.Bytes32, arg3 *bls.Signature, arg4 [32]byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessNewSignature", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessNewSignature indicates an expected call of ProcessNewSignature.
func (mr *MockBlsAggregationServiceMockRecorder) ProcessNewSignature(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNewSignature", reflect.TypeOf((*MockBlsAggregationService)(nil).ProcessNewSignature), arg0, arg1, arg2, arg3, arg4)
}
