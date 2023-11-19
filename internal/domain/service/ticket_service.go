package service

import "context"

type TicketService struct {
	repo TicketRepository
}

func NewTicketService(repo TicketRepository) *TicketService {
	return &TicketService{repo: repo}
}

func (s *TicketService) PurchaseTicket(ctx context.Context, user User, from string, to string, purchasePrice float32) (Ticket, error) {
	return s.repo.Create(ctx, user, from, to, purchasePrice)
}
