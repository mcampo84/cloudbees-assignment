package service

import (
	"context"

	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type TicketRepositoryAdapter struct {
	repo *in_memory.TicketRepository
}

// Create implements service.TicketRepository.
func (a *TicketRepositoryAdapter) Create(ctx context.Context, userID uint, from string, to string, purchasePrice float32) (service.Ticket, error) {
	ticket, err := a.repo.Create(ctx, userID, from, to, purchasePrice)
	if err != nil {
		return nil, err
	}

	return NewTicketModelAdapter(ticket), nil
}

func NewTicketRepositoryAdapter(repo *in_memory.TicketRepository) service.TicketRepository {
	return &TicketRepositoryAdapter{repo: repo}
}
