// Code generated by MockGen. DO NOT EDIT.
// Source: ticket_service.go

// Package grpc is a generated GoMock package.
package grpc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pb "github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

// MockTicketService is a mock of TicketService interface.
type MockTicketService struct {
	ctrl     *gomock.Controller
	recorder *MockTicketServiceMockRecorder
}

// MockTicketServiceMockRecorder is the mock recorder for MockTicketService.
type MockTicketServiceMockRecorder struct {
	mock *MockTicketService
}

// NewMockTicketService creates a new mock instance.
func NewMockTicketService(ctrl *gomock.Controller) *MockTicketService {
	mock := &MockTicketService{ctrl: ctrl}
	mock.recorder = &MockTicketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicketService) EXPECT() *MockTicketServiceMockRecorder {
	return m.recorder
}

// PurchaseTicket mocks base method.
func (m *MockTicketService) PurchaseTicket(ctx context.Context, user *pb.User, from, to string, purchasePrice float32) (Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PurchaseTicket", ctx, user, from, to, purchasePrice)
	ret0, _ := ret[0].(Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PurchaseTicket indicates an expected call of PurchaseTicket.
func (mr *MockTicketServiceMockRecorder) PurchaseTicket(ctx, user, from, to, purchasePrice interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurchaseTicket", reflect.TypeOf((*MockTicketService)(nil).PurchaseTicket), ctx, user, from, to, purchasePrice)
}
