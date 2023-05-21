package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/arjungandhi/magellan/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var port = 5601

type server struct {
	pb.UnimplementedMagellanServer
}

func (s *server) Hello(ctx context.Context, _ *emptypb.Empty) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello from Magellan"}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMagellanServer(s, &server{})
	log.Printf("Starting server on port %d", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
