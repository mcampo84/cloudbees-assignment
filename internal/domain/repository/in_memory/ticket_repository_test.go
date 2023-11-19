package in_memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TicketRepositorySuite struct {
	suite.Suite
	userRepo   *UserRepository
	ticketRepo *TicketRepository
}

func (suite *TicketRepositorySuite) SetupTest() {
	suite.userRepo = NewUserRepository()
	suite.ticketRepo = NewTicketRepository(suite.userRepo)
}

func (suite *TicketRepositorySuite) TestCreateTicket_Success() {
	// Create a user
	user := suite.userRepo.FindOrCreate("John", "Doe", "john@example.com")

	// Create ticket
	ticket, err := suite.ticketRepo.Create(context.Background(), user.FirstName, user.LastName, user.Email, "London", "France", 20.0)

	// Assertions
	suite.NotNil(ticket)
	suite.NoError(err)
	suite.Equal("London", ticket.From)
	suite.Equal("France", ticket.To)

	secondTicket, _ := suite.ticketRepo.Create(context.Background(), user.FirstName, user.LastName, user.Email, "London", "France", 20.0)
	suite.Equal(uint(2), secondTicket.ID)
}

func TestTicketRepositorySuite(t *testing.T) {
	suite.Run(t, new(TicketRepositorySuite))
}
