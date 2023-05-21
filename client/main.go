package main

import (
	"context"
	"log"
	"time"

	pb "github.com/arjungandhi/magellan/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	addr = "localhost:5601"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMagellanClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.GetMessage())
}
