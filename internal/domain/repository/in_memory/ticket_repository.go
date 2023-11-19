package in_memory

import (
	"context"

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

func (r *TicketRepository) Create(ctx context.Context, userID uint, from string, to string, purchasePrice float32) (*model.Ticket, error) {
	user, err := r.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	ticket := model.NewTicket(user, from, to, purchasePrice)
	ticket.ID = r.generateID()

	r.tickets[ticket.ID] = ticket
	r.maxID = ticket.ID

	return ticket, nil
}

func (r *TicketRepository) generateID() uint {
	return r.maxID + 1
}
