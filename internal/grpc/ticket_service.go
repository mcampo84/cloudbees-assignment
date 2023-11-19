//go:generate mockgen -source=ticket_service.go -destination=ticket_service.mockgen.go -package=grpc

package grpc

import (
	"context"
)

type TicketService interface {
	PurchaseTicket(ctx context.Context, firstName string, lastName string, email string, from string, to string, purchasePrice float32) (Ticket, error)
}
