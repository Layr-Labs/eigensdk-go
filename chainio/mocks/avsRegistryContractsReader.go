// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry (interfaces: Reader)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avsRegistryContractsReader.go -package=mocks github.com/Layr-Labs/eigensdk-go/chainio/clients/avsregistry Reader
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	big "math/big"
	reflect "reflect"

	contractOperatorStateRetriever "github.com/Layr-Labs/eigensdk-go/contracts/bindings/OperatorStateRetriever"
	types "github.com/Layr-Labs/eigensdk-go/types"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// GetCheckSignaturesIndices mocks base method.
func (m *MockReader) GetCheckSignaturesIndices(arg0 *bind.CallOpts, arg1 uint32, arg2 types.QuorumNums, arg3 []types.Bytes32) (contractOperatorStateRetriever.OperatorStateRetrieverCheckSignaturesIndices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckSignaturesIndices", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(contractOperatorStateRetriever.OperatorStateRetrieverCheckSignaturesIndices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckSignaturesIndices indicates an expected call of GetCheckSignaturesIndices.
func (mr *MockReaderMockRecorder) GetCheckSignaturesIndices(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckSignaturesIndices", reflect.TypeOf((*MockReader)(nil).GetCheckSignaturesIndices), arg0, arg1, arg2, arg3)
}

// GetOperatorAddrsInQuorumsAtCurrentBlock mocks base method.
func (m *MockReader) GetOperatorAddrsInQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.QuorumNums) ([][]common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorAddrsInQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([][]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorAddrsInQuorumsAtCurrentBlock indicates an expected call of GetOperatorAddrsInQuorumsAtCurrentBlock.
func (mr *MockReaderMockRecorder) GetOperatorAddrsInQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorAddrsInQuorumsAtCurrentBlock", reflect.TypeOf((*MockReader)(nil).GetOperatorAddrsInQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorFromId mocks base method.
func (m *MockReader) GetOperatorFromId(arg0 *bind.CallOpts, arg1 types.Bytes32) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorFromId", arg0, arg1)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorFromId indicates an expected call of GetOperatorFromId.
func (mr *MockReaderMockRecorder) GetOperatorFromId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorFromId", reflect.TypeOf((*MockReader)(nil).GetOperatorFromId), arg0, arg1)
}

// GetOperatorId mocks base method.
func (m *MockReader) GetOperatorId(arg0 *bind.CallOpts, arg1 common.Address) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorId", arg0, arg1)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorId indicates an expected call of GetOperatorId.
func (mr *MockReaderMockRecorder) GetOperatorId(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorId", reflect.TypeOf((*MockReader)(nil).GetOperatorId), arg0, arg1)
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockReader) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.Bytes32) (map[types.QuorumNum]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].(map[types.QuorumNum]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockReaderMockRecorder) GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockReader)(nil).GetOperatorStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsAtBlock mocks base method.
func (m *MockReader) GetOperatorsStakeInQuorumsAtBlock(arg0 *bind.CallOpts, arg1 types.QuorumNums, arg2 uint32) ([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtBlock indicates an expected call of GetOperatorsStakeInQuorumsAtBlock.
func (mr *MockReaderMockRecorder) GetOperatorsStakeInQuorumsAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtBlock", reflect.TypeOf((*MockReader)(nil).GetOperatorsStakeInQuorumsAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsAtCurrentBlock mocks base method.
func (m *MockReader) GetOperatorsStakeInQuorumsAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.QuorumNums) ([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperatorsStakeInQuorumsAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsAtCurrentBlock.
func (mr *MockReaderMockRecorder) GetOperatorsStakeInQuorumsAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsAtCurrentBlock", reflect.TypeOf((*MockReader)(nil).GetOperatorsStakeInQuorumsAtCurrentBlock), arg0, arg1)
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock mocks base method.
func (m *MockReader) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0 *bind.CallOpts, arg1 types.Bytes32, arg2 uint32) (types.QuorumNums, [][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.QuorumNums)
	ret1, _ := ret[1].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtBlock.
func (mr *MockReaderMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtBlock(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtBlock", reflect.TypeOf((*MockReader)(nil).GetOperatorsStakeInQuorumsOfOperatorAtBlock), arg0, arg1, arg2)
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock mocks base method.
func (m *MockReader) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0 *bind.CallOpts, arg1 types.Bytes32) (types.QuorumNums, [][]contractOperatorStateRetriever.OperatorStateRetrieverOperator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", arg0, arg1)
	ret0, _ := ret[0].(types.QuorumNums)
	ret1, _ := ret[1].([][]contractOperatorStateRetriever.OperatorStateRetrieverOperator)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock indicates an expected call of GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock.
func (mr *MockReaderMockRecorder) GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock", reflect.TypeOf((*MockReader)(nil).GetOperatorsStakeInQuorumsOfOperatorAtCurrentBlock), arg0, arg1)
}

// GetQuorumCount mocks base method.
func (m *MockReader) GetQuorumCount(arg0 *bind.CallOpts) (byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuorumCount", arg0)
	ret0, _ := ret[0].(byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuorumCount indicates an expected call of GetQuorumCount.
func (mr *MockReaderMockRecorder) GetQuorumCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuorumCount", reflect.TypeOf((*MockReader)(nil).GetQuorumCount), arg0)
}

// IsOperatorRegistered mocks base method.
func (m *MockReader) IsOperatorRegistered(arg0 *bind.CallOpts, arg1 common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOperatorRegistered", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOperatorRegistered indicates an expected call of IsOperatorRegistered.
func (mr *MockReaderMockRecorder) IsOperatorRegistered(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOperatorRegistered", reflect.TypeOf((*MockReader)(nil).IsOperatorRegistered), arg0, arg1)
}

// QueryExistingRegisteredOperatorPubKeys mocks base method.
func (m *MockReader) QueryExistingRegisteredOperatorPubKeys(arg0 context.Context, arg1, arg2, arg3 *big.Int) ([]common.Address, []types.OperatorPubkeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryExistingRegisteredOperatorPubKeys", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].([]types.OperatorPubkeys)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryExistingRegisteredOperatorPubKeys indicates an expected call of QueryExistingRegisteredOperatorPubKeys.
func (mr *MockReaderMockRecorder) QueryExistingRegisteredOperatorPubKeys(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryExistingRegisteredOperatorPubKeys", reflect.TypeOf((*MockReader)(nil).QueryExistingRegisteredOperatorPubKeys), arg0, arg1, arg2, arg3)
}

// QueryExistingRegisteredOperatorSockets mocks base method.
func (m *MockReader) QueryExistingRegisteredOperatorSockets(arg0 context.Context, arg1, arg2, arg3 *big.Int) (map[types.Bytes32]types.Socket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryExistingRegisteredOperatorSockets", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(map[types.Bytes32]types.Socket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryExistingRegisteredOperatorSockets indicates an expected call of QueryExistingRegisteredOperatorSockets.
func (mr *MockReaderMockRecorder) QueryExistingRegisteredOperatorSockets(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryExistingRegisteredOperatorSockets", reflect.TypeOf((*MockReader)(nil).QueryExistingRegisteredOperatorSockets), arg0, arg1, arg2, arg3)
}
