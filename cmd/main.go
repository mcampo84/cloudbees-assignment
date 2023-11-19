package main

import (
	"context"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc"

	"github.com/mcampo84/cloudbees-assignment/cmd/server"
	"github.com/mcampo84/cloudbees-assignment/internal/domain"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
	localRPC "github.com/mcampo84/cloudbees-assignment/internal/grpc"
	"github.com/mcampo84/cloudbees-assignment/internal/proto/pb"
)

func main() {
	app := fx.New(
		fx.Provide(
			context.Background,
		),
		localRPC.Module,
		domain.Module,
		server.Module,
		fx.Invoke(seedTrains),
		fx.Invoke(startServer),
	)

	ctx := context.Background()

	if err := app.Start(ctx); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	defer app.Stop(ctx)

	select {}
}

func seedTrains(svc service.TrainService) {
	svc.InitTrains()
}

func startServer(localServer pb.TrainTicketServiceServer) {
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	pb.RegisterTrainTicketServiceServer(server, localServer)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
