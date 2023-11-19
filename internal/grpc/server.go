package grpc

import (
	"context"

	"google.golang.org/grpc/status"

	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type Server struct {
	ticketService TicketService
}

func NewServer(ticketService TicketService) *Server {
	return &Server{ticketService: ticketService}
}

func (s *Server) PurchaseTicket(ctx context.Context, req *pb.PurchaseTicketRequest) (*pb.PurchaseTicketResponse, error) {
	userID := uint(req.GetUserId())
	from := req.GetFrom()
	to := req.GetTo()
	purchasePrice := req.GetPrice()

	ticket, err := s.ticketService.PurchaseTicket(ctx, userID, from, to, purchasePrice)
	st, ok := status.FromError(err)
	if !ok {
		return nil, st.Err()
	}

	return ticket.ToResponse(), nil
}
