package grpc

import (
	"context"

	serviceAdapter "github.com/mcampo84/cloudbees-assignment/internal/domain/adapter/service"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
	"github.com/mcampo84/cloudbees-assignment/internal/grpc"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type TicketServiceAdapter struct {
	ticketService *service.TicketService
}

// PurchaseTicket implements grpc.TicketService.
func (a *TicketServiceAdapter) PurchaseTicket(ctx context.Context, user *pb.User, from string, to string, purchasePrice float32) (grpc.Ticket, error) {
	userAdapter := serviceAdapter.NewUserProtoAdapter(user)

	ticket, err := a.ticketService.PurchaseTicket(ctx, userAdapter, from, to, purchasePrice)
	if err != nil {
		return nil, err
	}

	return NewTicketAdapter(ticket), nil
}

func NewTicketServiceAdapter(ticketService *service.TicketService) grpc.TicketService {
	return &TicketServiceAdapter{ticketService: ticketService}
}
