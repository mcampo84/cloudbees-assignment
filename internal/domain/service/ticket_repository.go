//go:generate mockgen -source=ticket_repository.go -destination=ticket_repository.mockgen.go -package=service

package service

import "context"

type TicketRepository interface {
	Create(ctx context.Context, user User, train Train, purchasePrice float32) (Ticket, error)
	Delete(ctx context.Context, id uint)
}
