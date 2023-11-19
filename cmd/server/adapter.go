package server

import (
	"context"

	"github.com/mcampo84/cloudbees-assignment/internal/grpc"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

type ServerAdapter struct {
	pb.UnimplementedTrainTicketServiceServer

	server *grpc.Server
}

func NewServerAdapter(server *grpc.Server) pb.TrainTicketServiceServer {
	return &ServerAdapter{server: server}
}

// PurchaseTicket implements pb.TrainTicketServiceServer.
func (a *ServerAdapter) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Ticket, error) {
	return a.server.PurchaseTicket(ctx, req)
}
