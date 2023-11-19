package service

import (
	"context"

	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type TicketRepositoryAdapter struct {
	ticketRepo *in_memory.TicketRepository
	trainRepo  *in_memory.TrainRepository
	userRepo   *in_memory.UserRepository
}

func NewTicketRepositoryAdapter(ticketRepo *in_memory.TicketRepository, userRepo *in_memory.UserRepository, trainRepo *in_memory.TrainRepository) service.TicketRepository {
	return &TicketRepositoryAdapter{ticketRepo: ticketRepo, userRepo: userRepo, trainRepo: trainRepo}
}

func (a *TicketRepositoryAdapter) Create(ctx context.Context, user service.User, train service.Train, purchasePrice float32) (service.Ticket, error) {
	userModel, err := a.userRepo.GetByID(user.GetID())
	if err != nil {
		return nil, err
	}

	trainModel, err := a.trainRepo.GetByID(train.GetID())
	if err != nil {
		return nil, err
	}

	ticket, err := a.ticketRepo.Create(ctx, userModel, trainModel, purchasePrice)
	if err != nil {
		return nil, err
	}

	return NewTicketModelAdapter(ticket), nil
}

func (a *TicketRepositoryAdapter) Delete(ctx context.Context, id uint) {
	a.ticketRepo.Delete(ctx, id)
}
