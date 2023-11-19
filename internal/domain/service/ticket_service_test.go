package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type TicketServiceSuite struct {
	suite.Suite
	ctrl             *gomock.Controller
	mockTicketRepo   *MockTicketRepository
	mockUserRepo     *MockUserRepository
	mockTrainRepo    *MockTrainRepository
	mockTrainService *MockTrainService
	ticketService    *TicketService
}

func (suite *TicketServiceSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockTicketRepo = NewMockTicketRepository(suite.ctrl)
	suite.mockUserRepo = NewMockUserRepository(suite.ctrl)
	suite.mockTrainRepo = NewMockTrainRepository(suite.ctrl)
	suite.mockTrainService = NewMockTrainService(suite.ctrl)
	suite.ticketService = NewTicketService(suite.mockTicketRepo, suite.mockUserRepo, suite.mockTrainRepo, suite.mockTrainService)
}

func (suite *TicketServiceSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *TicketServiceSuite) TestPurchaseTicket_Success() {
	// Mock behavior for successful ticket purchase
	mockTrain := NewMockTrain(suite.ctrl)
	expectedTicket := NewMockTicket(suite.ctrl)
	firstName := "John"
	lastName := "Doe"
	email := "jdoe@email.co"
	user := NewMockUser(suite.ctrl)
	from := "from"
	to := "to"
	price := float32(20.0)
	suite.mockTrainRepo.EXPECT().Find(from, to).Return(mockTrain, nil)
	suite.mockUserRepo.EXPECT().FindOrCreate(firstName, lastName, email).Return(user)
	suite.mockTicketRepo.EXPECT().Create(gomock.Any(), user, mockTrain, price).Return(expectedTicket, nil)
	expectedTicket.EXPECT().GetUser().Return(user)
	suite.mockTrainService.EXPECT().AutoAssignSeat(mockTrain, user).Return(nil)

	// Invoke method
	ticket, err := suite.ticketService.PurchaseTicket(context.Background(), firstName, lastName, email, from, to, price)

	// Assert expectations and results
	suite.NoError(err)
	suite.Equal(expectedTicket, ticket)
}

func (suite *TicketServiceSuite) TestPurchaseTicket_TicketCreationFailure() {
	mockTrain := NewMockTrain(suite.ctrl)
	// Mock behavior for failed ticket purchase
	firstName := "John"
	lastName := "Doe"
	email := "jdoe@email.co"
	user := NewMockUser(suite.ctrl)
	from := "from"
	to := "to"
	price := float32(20.0)
	suite.mockTrainRepo.EXPECT().Find(from, to).Return(mockTrain, nil)
	suite.mockUserRepo.EXPECT().FindOrCreate(firstName, lastName, email).Return(user)
	suite.mockTicketRepo.EXPECT().Create(gomock.Any(), user, mockTrain, price).Return(nil, errors.New("oops"))

	// Invoke method
	_, err := suite.ticketService.PurchaseTicket(context.Background(), firstName, lastName, email, from, to, price)

	// Assert expectations and results
	suite.Error(err)
}

func (suite *TicketServiceSuite) TestPurchaseTicket_SeatAssignFailure() {
	// Mock behavior for successful ticket purchase
	mockTrain := NewMockTrain(suite.ctrl)
	expectedTicket := NewMockTicket(suite.ctrl)
	firstName := "John"
	lastName := "Doe"
	email := "jdoe@email.co"
	user := NewMockUser(suite.ctrl)
	from := "from"
	to := "to"
	price := float32(20.0)
	suite.mockTrainRepo.EXPECT().Find(from, to).Return(mockTrain, nil)
	suite.mockUserRepo.EXPECT().FindOrCreate(firstName, lastName, email).Return(user)
	suite.mockTicketRepo.EXPECT().Create(gomock.Any(), user, mockTrain, price).Return(expectedTicket, nil)
	expectedTicket.EXPECT().GetUser().Return(user)
	suite.mockTrainService.EXPECT().AutoAssignSeat(mockTrain, user).Return(errors.New("oops"))
	expectedTicket.EXPECT().GetID().Return(uint(1))
	suite.mockTicketRepo.EXPECT().Delete(gomock.Any(), uint(1))

	// Invoke method
	_, err := suite.ticketService.PurchaseTicket(context.Background(), firstName, lastName, email, from, to, price)

	// Assert expectations and results
	suite.EqualError(err, "oops")
}

func TestTicketServiceSuite(t *testing.T) {
	suite.Run(t, new(TicketServiceSuite))
}
