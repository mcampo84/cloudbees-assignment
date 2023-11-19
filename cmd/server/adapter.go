package server

import (
	"context"

	"github.com/mcampo84/cloudbees-assignment/internal/grpc"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServerAdapter struct {
	pb.UnimplementedTrainTicketServiceServer

	server *grpc.Server
}

func NewServerAdapter(server *grpc.Server) pb.TrainTicketServiceServer {
	return &ServerAdapter{server: server}
}

// PurchaseTicket implements pb.TrainTicketServiceServer.
func (a *ServerAdapter) PurchaseTicket(ctx context.Context, req *pb.Ticket) (*pb.Receipt, error) {
	return a.server.PurchaseTicket(ctx, req)
}

func (a *ServerAdapter) ViewReceipt(context.Context, *pb.User) (*pb.Receipt, error) {
	panic("unimplemented")
}

func (a *ServerAdapter) ViewPassengerManifest(context.Context, *emptypb.Empty) (*pb.Manifest, error) {
	panic("unimplemented")
}

func (a *ServerAdapter) CancelTicket(context.Context, *pb.User) (*emptypb.Empty, error) {
	panic("unimplemented")
}

func (a *ServerAdapter) ChangeSeat(context.Context, *pb.ChangeSeatRequest) (*pb.Receipt, error) {
	panic("unimplemented")
}
