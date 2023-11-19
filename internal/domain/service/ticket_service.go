package service

import "context"

type TicketService struct {
	ticketRepo   TicketRepository
	trainRepo    TrainRepository
	trainService TrainService
	userRepo     UserRepository
}

func NewTicketService(ticketRepo TicketRepository, userRepo UserRepository, trainRepo TrainRepository, trainService TrainService) *TicketService {
	return &TicketService{ticketRepo: ticketRepo, userRepo: userRepo, trainRepo: trainRepo, trainService: trainService}
}

func (s *TicketService) PurchaseTicket(ctx context.Context, firstName string, lastName string, email string, from string, to string, purchasePrice float32) (Ticket, error) {
	train, err := s.trainRepo.Find(from, to)
	if err != nil {
		return nil, err
	}

	user := s.userRepo.FindOrCreate(firstName, lastName, email)

	ticket, err := s.ticketRepo.Create(ctx, user, train, purchasePrice)
	if err != nil {
		return nil, err
	}

	err = s.trainService.AutoAssignSeat(train, ticket.GetUser())
	if err != nil {
		// if we can't assign the seat, we can't book the ticket - in a relational database, this would be handled in a transaction
		// but since we're impleenting this with an in-memory solution, I'm just deleting the ticket we just created.
		// since I don't have a mechanism to determine if we created or found the user, I'll just leave the user record
		// in memory since I don't foresee too many users being created for this exercise to the pont that we'll run out of
		// memory on the local machine
		s.ticketRepo.Delete(ctx, ticket.GetID())
		return nil, err
	}

	return ticket, nil
}
