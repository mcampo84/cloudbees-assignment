package in_memory

import (
	"context"
	"testing"

	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/stretchr/testify/suite"
)

type TicketRepositorySuite struct {
	suite.Suite
	userRepo   *UserRepository
	ticketRepo *TicketRepository
}

func (suite *TicketRepositorySuite) SetupSuite() {
	suite.userRepo = NewUserRepository()
	suite.ticketRepo = NewTicketRepository(suite.userRepo)
}

func (suite *TicketRepositorySuite) SetupTest() {
	suite.ticketRepo.tickets = map[uint]*model.Ticket{}
	suite.userRepo.users = map[uint]*model.User{}
}

func (suite *TicketRepositorySuite) TestCreateTicket_Success() {
	train := &model.Train{From: "London", To: "France"}
	user := &model.User{FirstName: "John", LastName: "Doe", Email: "john@example.com"}

	// Create ticket
	ticket, err := suite.ticketRepo.Create(context.Background(), user, train, 20.0)

	// Assertions
	suite.NotNil(ticket)
	suite.NoError(err)
	suite.Equal("London", ticket.Train.From)
	suite.Equal("France", ticket.Train.To)

	secondTicket, _ := suite.ticketRepo.Create(context.Background(), user, train, 20.0)
	suite.Equal(uint(2), secondTicket.ID)
	suite.Equal(ticket, user.Tickets[0])
	suite.Equal(secondTicket, user.Tickets[1])
}

func (suite *TicketRepositorySuite) TestDeleteTicket() {
	suite.Empty(suite.ticketRepo.tickets)

	// delete nonexistent ticket
	suite.ticketRepo.Delete(context.Background(), uint(1))
	suite.Empty(suite.ticketRepo.tickets)

	// delete ticket that exists for a key
	fakeTicket := &model.Ticket{}
	suite.ticketRepo.tickets[1] = fakeTicket
	suite.ticketRepo.Delete(context.Background(), uint(1))
	suite.Empty(suite.ticketRepo.tickets)
}

func TestTicketRepositorySuite(t *testing.T) {
	suite.Run(t, new(TicketRepositorySuite))
}
