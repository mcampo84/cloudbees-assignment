// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// GetEmail mocks base method.
func (m *MockUser) GetEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockUserMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockUser)(nil).GetEmail))
}

// GetFirstName mocks base method.
func (m *MockUser) GetFirstName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFirstName indicates an expected call of GetFirstName.
func (mr *MockUserMockRecorder) GetFirstName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstName", reflect.TypeOf((*MockUser)(nil).GetFirstName))
}

// GetID mocks base method.
func (m *MockUser) GetID() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(uint)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockUserMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUser)(nil).GetID))
}

// GetLastName mocks base method.
func (m *MockUser) GetLastName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLastName indicates an expected call of GetLastName.
func (mr *MockUserMockRecorder) GetLastName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastName", reflect.TypeOf((*MockUser)(nil).GetLastName))
}
