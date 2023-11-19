// Code generated by MockGen. DO NOT EDIT.
// Source: train_service.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTrainService is a mock of TrainService interface.
type MockTrainService struct {
	ctrl     *gomock.Controller
	recorder *MockTrainServiceMockRecorder
}

// MockTrainServiceMockRecorder is the mock recorder for MockTrainService.
type MockTrainServiceMockRecorder struct {
	mock *MockTrainService
}

// NewMockTrainService creates a new mock instance.
func NewMockTrainService(ctrl *gomock.Controller) *MockTrainService {
	mock := &MockTrainService{ctrl: ctrl}
	mock.recorder = &MockTrainServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrainService) EXPECT() *MockTrainServiceMockRecorder {
	return m.recorder
}

// AutoAssignSeat mocks base method.
func (m *MockTrainService) AutoAssignSeat(train Train, passenger User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoAssignSeat", train, passenger)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoAssignSeat indicates an expected call of AutoAssignSeat.
func (mr *MockTrainServiceMockRecorder) AutoAssignSeat(train, passenger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoAssignSeat", reflect.TypeOf((*MockTrainService)(nil).AutoAssignSeat), train, passenger)
}

// GetPassengers mocks base method.
func (m *MockTrainService) GetPassengers(train Train) map[string][]User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPassengers", train)
	ret0, _ := ret[0].(map[string][]User)
	return ret0
}

// GetPassengers indicates an expected call of GetPassengers.
func (mr *MockTrainServiceMockRecorder) GetPassengers(train interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassengers", reflect.TypeOf((*MockTrainService)(nil).GetPassengers), train)
}

// InitTrains mocks base method.
func (m *MockTrainService) InitTrains() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitTrains")
}

// InitTrains indicates an expected call of InitTrains.
func (mr *MockTrainServiceMockRecorder) InitTrains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitTrains", reflect.TypeOf((*MockTrainService)(nil).InitTrains))
}

// NewTrain mocks base method.
func (m *MockTrainService) NewTrain(from, to string) Train {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTrain", from, to)
	ret0, _ := ret[0].(Train)
	return ret0
}

// NewTrain indicates an expected call of NewTrain.
func (mr *MockTrainServiceMockRecorder) NewTrain(from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTrain", reflect.TypeOf((*MockTrainService)(nil).NewTrain), from, to)
}