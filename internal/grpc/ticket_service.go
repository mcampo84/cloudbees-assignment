//go:generate mockgen -source=ticket_service.go -destination=ticket_service.mockgen.go -package=grpc

package grpc

import (
	"context"

	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type TicketService interface {
	PurchaseTicket(ctx context.Context, user *pb.User, from string, to string, purchasePrice float32) (Ticket, error)
}
