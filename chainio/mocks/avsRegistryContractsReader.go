// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/avsregistry (interfaces: AvsRegistryReader)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/avsregistry AvsRegistryReader
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractBLSOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSOperatorStateRetriever"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsRegistryReader is a mock of AvsRegistryReader interface.
type MockAvsRegistryReader struct {
	ctrl     *gomock.Controller
	recorder *MockAvsRegistryReaderMockRecorder
}

// MockAvsRegistryReaderMockRecorder is the mock recorder for MockAvsRegistryReader.
type MockAvsRegistryReaderMockRecorder struct {
	mock *MockAvsRegistryReader
}

// NewMockAvsRegistryReader creates a new mock instance.
func NewMockAvsRegistryReader(ctrl *gomock.Controller) *MockAvsRegistryReader {
	mock := &MockAvsRegistryReader{ctrl: ctrl}
	mock.recorder = &MockAvsRegistryReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsRegistryReader) EXPECT() *MockAvsRegistryReaderMockRecorder {
	return m.recorder
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockAvsRegistryReader) GetCheckSignaturesIndices(arg0 context.Context, arg1 uint32, arg2 []byte, arg3 [][32]byte) (contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockAvsRegistryReaderMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetOperatorId mocks base method.
func (m *MockAvsRegistryReader) GetOperatorId(arg0 context.Context, arg1 common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorId", arg0, arg1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorId indicates an expected call of GetOperatorId.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorId", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorId), arg0, arg1)
}

// GetOperatorsStakeInQuorumsAtBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsAtBlock(arg0 context.Context, arg1 []byte, arg2 uint32) ([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtBlock indicates an expected call of GetOperatorsStakeInQuorumsAtBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0 context.Context, arg1 [32]byte, arg2 uint32) ([]byte, [][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsOfOperatorAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockAvsRegistryReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0 context.Context, arg1 [32]byte) ([]byte, [][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([][]contractBLSOperatorStateRetriever.BLSOperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockAvsRegistryReaderMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockAvsRegistryReader)(nil).GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// IsOperatorRegistered mocks base method.
func (m *MockAvsRegistryReader) IsOperatorRegistered(arg0 context.Context, arg1 common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperatorRegistered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperatorRegistered indicates an expected call of IsOperatorRegistered.
func (mr *MockAvsRegistryReaderMockRecorder) IsOperatorRegistered(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperatorRegistered", reflect.TypeOf((*MockAvsRegistryReader)(nil).IsOperatorRegistered), arg0, arg1)
}
