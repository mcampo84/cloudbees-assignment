package grpc

import (
	"context"

	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
	"github.com/mcampo84/cloudbees-assignment/internal/grpc"
)

type TicketServiceAdapter struct {
	ticketService *service.TicketService
}

// PurchaseTicket implements grpc.TicketService.
func (a *TicketServiceAdapter) PurchaseTicket(ctx context.Context, userID uint, from string, to string, purchasePrice float32) (grpc.Ticket, error) {
	ticket, err := a.ticketService.PurchaseTicket(ctx, userID, from, to, purchasePrice)
	if err != nil {
		return nil, err
	}

	return NewTicketAdapter(ticket), nil
}

func NewTicketServiceAdapter(ticketService *service.TicketService) grpc.TicketService {
	return &TicketServiceAdapter{ticketService: ticketService}
}
