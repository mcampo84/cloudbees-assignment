package grpc

import (
	"context"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type Server struct {
	ticketService TicketService
}

func NewServer(ticketService TicketService) *Server {
	return &Server{ticketService: ticketService}
}

func (s *Server) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Receipt, error) {
	firstName := req.GetFirstName()
	lastName := req.GetLastName()
	email := req.GetEmail()
	from := req.GetFrom()
	to := req.GetTo()
	purchasePrice := req.GetPrice()

	ticket, err := s.ticketService.PurchaseTicket(ctx, firstName, lastName, email, from, to, purchasePrice)
	st, ok := status.FromError(err)
	if !ok {
		return nil, st.Err()
	}

	return ticket.GenerateReceipt(), nil
}

func (a *Server) ViewReceipt(context.Context, *pb.User) (*pb.Receipt, error) {
	panic("unimplemented")
}

func (a *Server) ViewPassengerManifest(context.Context, *emptypb.Empty) (*pb.Manifest, error) {
	panic("unimplemented")
}

func (a *Server) CancelTicket(context.Context, *pb.User) (*emptypb.Empty, error) {
	panic("unimplemented")
}

func (a *Server) ChangeSeat(context.Context, *pb.ChangeSeatRequest) (*pb.Receipt, error) {
	panic("unimplemented")
}
