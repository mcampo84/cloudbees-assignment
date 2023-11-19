package in_memory

import (
	"context"
	"fmt"

	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
)

type TicketRepository struct {
	maxID    uint
	tickets  map[uint]*model.Ticket
	userRepo *UserRepository
}

func NewTicketRepository(userRepo *UserRepository) *TicketRepository {
	return &TicketRepository{tickets: map[uint]*model.Ticket{}, userRepo: userRepo}
}

func (r *TicketRepository) Create(ctx context.Context, user *model.User, train *model.Train, purchasePrice float32) (*model.Ticket, error) {
	ticket := model.NewTicket(user, train, purchasePrice)
	ticket.ID = r.generateID()

	r.tickets[ticket.ID] = ticket
	r.maxID = ticket.ID

	fmt.Printf("booked ticket to %s from %s for %s %s at %f with id %d\n", train.To, train.From, user.FirstName, user.LastName, purchasePrice, ticket.ID)

	user.Tickets = append(user.Tickets, ticket)

	return ticket, nil
}

func (r *TicketRepository) Delete(ctx context.Context, id uint) {
	fmt.Printf("deleting ticket %d\n", id)
	delete(r.tickets, id)
}

func (r *TicketRepository) generateID() uint {
	return r.maxID + 1
}
