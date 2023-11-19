//go:generate mockgen -source=ticket_repository.go -destination=ticket_repository.mockgen.go -package=service

package service

import "context"

type TicketRepository interface {
	Create(ctx context.Context, userID int64, from string, to string, purchasePrice float32) (Ticket, error)
}
