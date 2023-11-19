package service

import (
	"github.com/mcampo84/cloudbees-assignment/internal/domain/model"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

type TicketModelAdapter struct {
	ticket *model.Ticket
}

func (a *TicketModelAdapter) GetID() uint {
	return a.ticket.ID
}

// GetFrom implements service.Ticket.
func (a *TicketModelAdapter) GetFrom() string {
	return a.ticket.Train.From
}

// GetPurchasePrice implements service.Ticket.
func (a *TicketModelAdapter) GetPurchasePrice() float32 {
	return a.ticket.PurchasePrice
}

// GetTo implements service.Ticket.
func (a *TicketModelAdapter) GetTo() string {
	return a.ticket.Train.To
}

// GetUser implements service.Ticket.
func (a TicketModelAdapter) GetUser() service.User {
	user := a.ticket.User

	return NewUserModelAdapter(user)
}

func NewTicketModelAdapter(ticket *model.Ticket) service.Ticket {
	return &TicketModelAdapter{ticket: ticket}
}
