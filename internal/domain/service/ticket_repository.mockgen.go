// Code generated by MockGen. DO NOT EDIT.
// Source: ticket_repository.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTicketRepository is a mock of TicketRepository interface.
type MockTicketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTicketRepositoryMockRecorder
}

// MockTicketRepositoryMockRecorder is the mock recorder for MockTicketRepository.
type MockTicketRepositoryMockRecorder struct {
	mock *MockTicketRepository
}

// NewMockTicketRepository creates a new mock instance.
func NewMockTicketRepository(ctrl *gomock.Controller) *MockTicketRepository {
	mock := &MockTicketRepository{ctrl: ctrl}
	mock.recorder = &MockTicketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketRepository) EXPECT() *MockTicketRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTicketRepository) Create(ctx context.Context, userID uint, from, to string, purchasePrice float32) (Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID, from, to, purchasePrice)
	ret0, _ := ret[0].(Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTicketRepositoryMockRecorder) Create(ctx, userID, from, to, purchasePrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTicketRepository)(nil).Create), ctx, userID, from, to, purchasePrice)
}
