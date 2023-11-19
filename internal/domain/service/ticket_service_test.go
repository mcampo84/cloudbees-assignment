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
	ctrl          *gomock.Controller
	mockRepo      *MockTicketRepository
	ticketService *TicketService
}

func (suite *TicketServiceSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockRepo = NewMockTicketRepository(suite.ctrl)
	suite.ticketService = NewTicketService(suite.mockRepo)
}

func (suite *TicketServiceSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *TicketServiceSuite) TestPurchaseTicket_Success() {
	// Mock behavior for successful ticket purchase
	expectedTicket := new(MockTicket)
	mockUser := new(MockUser)
	from := "from"
	to := "to"
	price := float32(20.0)
	suite.mockRepo.EXPECT().Create(gomock.Any(), mockUser, from, to, price).Return(expectedTicket, nil)

	// Invoke method
	ticket, err := suite.ticketService.PurchaseTicket(context.Background(), mockUser, from, to, price)

	// Assert expectations and results
	suite.NoError(err)
	suite.Equal(expectedTicket, ticket)
}

func (suite *TicketServiceSuite) TestPurchaseTicket_Failure() {
	// Mock behavior for failed ticket purchase
	mockUser := new(MockUser)
	from := "from"
	to := "to"
	price := float32(20.0)
	suite.mockRepo.EXPECT().Create(gomock.Any(), mockUser, from, to, price).Return(nil, errors.New("oops"))

	// Invoke method
	_, err := suite.ticketService.PurchaseTicket(context.Background(), mockUser, from, to, price)

	// Assert expectations and results
	suite.Error(err)
}

func TestTicketServiceSuite(t *testing.T) {
	suite.Run(t, new(TicketServiceSuite))
}
