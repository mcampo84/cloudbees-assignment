package in_memory

import (
	"context"
	"testing"

	domainError "github.com/mcampo84/cloudbees-assignment/internal/domain/error"
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
	user := suite.userRepo.Create("John", "Doe", "john@example.com")

	// Create ticket
	ticket, err := suite.ticketRepo.Create(context.Background(), user.ID, "London", "France", 20.0)

	// Assertions
	suite.NotNil(ticket)
	suite.NoError(err)
	suite.Equal("London", ticket.From)
	suite.Equal("France", ticket.To)

	secondTicket, _ := suite.ticketRepo.Create(context.Background(), user.ID, "London", "France", 20.0)
	suite.Equal(uint(2), secondTicket.ID)
}

func (suite *TicketRepositorySuite) TestCreateTicket_UserDNE() {
	// Create ticket
	_, err := suite.ticketRepo.Create(context.Background(), 10, "London", "France", 20.0)

	suite.EqualError(err, domainError.NotFound("user %d", 10).Error())
}

func TestTicketRepositorySuite(t *testing.T) {
	suite.Run(t, new(TicketRepositorySuite))
}
