package grpc

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type ServerSuite struct {
	suite.Suite
	controller    *gomock.Controller
	server        *Server
	mockTicketSvc *MockTicketService
}

func (suite *ServerSuite) SetupTest() {
	suite.controller = gomock.NewController(suite.T())
	suite.mockTicketSvc = NewMockTicketService(suite.controller)
	suite.server = NewServer(suite.mockTicketSvc)
}

// TestPurchaseTicket_Success tests the PurchaseTicket method on successful ticket purchase.
func (suite *ServerSuite) TestPurchaseTicket_Success() {
	ticket := NewMockTicket(suite.controller)
	expectedResponse := &pb.Receipt{}
	suite.mockTicketSvc.EXPECT().PurchaseTicket(gomock.Any(), "John", "Doe", "jdoe@email.co", "from", "to", float32(20.0)).Return(ticket, nil)
	ticket.EXPECT().GenerateReceipt().Return(expectedResponse)

	request := &pb.Ticket{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jdoe@email.co",
		From:      "from",
		To:        "to",
		Price:     20.0,
	}

	response, err := suite.server.PurchaseTicket(context.Background(), request)

	suite.NoError(err)
	suite.Equal(expectedResponse, response)
}

// TestPurchaseTicket_Failure tests the PurchaseTicket method on a failed ticket purchase.
func (suite *ServerSuite) TestPurchaseTicket_Failure() {
	// Mock TicketService behavior for a failure scenario
	err := errors.New("User not found")
	suite.mockTicketSvc.EXPECT().PurchaseTicket(gomock.Any(), "John", "Doe", "jdoe@email.co", "from", "to", float32(20.0)).Return(nil, err)

	request := &pb.Ticket{
		FirstName: "John",
		LastName: "Doe",
		Email: "jdoe@email.co",
		From:  "from",
		To:    "to",
		Price: 20.0,
	}

	_, err = suite.server.PurchaseTicket(context.Background(), request)

	st, _ := status.FromError(err)

	suite.Equal(codes.Unknown, st.Code())
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
