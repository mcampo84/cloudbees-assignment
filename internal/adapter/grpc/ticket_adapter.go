package grpc

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
	"github.com/mcampo84/cloudbees-assignment/internal/grpc"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type TicketAdapter struct {
	ticket service.Ticket
}

func NewTicketAdapter(ticket service.Ticket) grpc.Ticket {
	return &TicketAdapter{ticket: ticket}
}

func (a *TicketAdapter) GenerateReceipt() *pb.Receipt {
	user := a.ticket.GetUser()
	userAdapter := NewUserAdapter(user)

	return &pb.Receipt{
		User:  userAdapter.ToProto(),
		From:  a.ticket.GetFrom(),
		To:    a.ticket.GetTo(),
		Price: a.ticket.GetPurchasePrice(),
	}
}
