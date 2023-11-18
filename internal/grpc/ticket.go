//go:generate mockgen -source=ticket.go -destination=ticket.mockgen.go -package=grpc

package grpc

import "github.com/mcampo84/cloudbees-assignment/internal/proto/pb"

type Ticket interface {
	ToResponse() *pb.PurchaseTicketResponse
}
