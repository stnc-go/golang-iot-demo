// Code generated by MockGen. DO NOT EDIT.
// Source: iot-demo/pkg/metrics/add-metrics (interfaces: Inserter,ConfigGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	alert "iot-demo/pkg/metrics/alert"
	ingestion "iot-demo/pkg/metrics/ingestion"
	reflect "reflect"
)

// MockInserter is a mock of Inserter interface
type MockInserter struct {
	ctrl     *gomock.Controller
	recorder *MockInserterMockRecorder
}

// MockInserterMockRecorder is the mock recorder for MockInserter
type MockInserterMockRecorder struct {
	mock *MockInserter
}

// NewMockInserter creates a new mock instance
func NewMockInserter(ctrl *gomock.Controller) *MockInserter {
	mock := &MockInserter{ctrl: ctrl}
	mock.recorder = &MockInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInserter) EXPECT() *MockInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockInserter) Insert(arg0 int, arg1 []*ingestion.DecimalMetricValue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockInserterMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockInserter)(nil).Insert), arg0, arg1)
}

// MockConfigGetter is a mock of ConfigGetter interface
type MockConfigGetter struct {
	ctrl     *gomock.Controller
	recorder *MockConfigGetterMockRecorder
}

// MockConfigGetterMockRecorder is the mock recorder for MockConfigGetter
type MockConfigGetterMockRecorder struct {
	mock *MockConfigGetter
}

// NewMockConfigGetter creates a new mock instance
func NewMockConfigGetter(ctrl *gomock.Controller) *MockConfigGetter {
	mock := &MockConfigGetter{ctrl: ctrl}
	mock.recorder = &MockConfigGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigGetter) EXPECT() *MockConfigGetterMockRecorder {
	return m.recorder
}

// GetThreshold mocks base method
func (m *MockConfigGetter) GetThreshold() alert.Threshold {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThreshold")
	ret0, _ := ret[0].(alert.Threshold)
	return ret0
}

// GetThreshold indicates an expected call of GetThreshold
func (mr *MockConfigGetterMockRecorder) GetThreshold() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThreshold", reflect.TypeOf((*MockConfigGetter)(nil).GetThreshold))
}
