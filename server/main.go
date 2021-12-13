// Package main implements a server for Counter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "sync/atomic"
  pb "github.com/ritmatter/gocounter/counter"
)

var (
	port = flag.Int("port", 50051, "The server port")
  counter = int64(0)
)

type server struct {
	pb.UnimplementedCounterServer
}

func (s *server) Increment(ctx context.Context, in *pb.IncrementRequest) (*pb.IncrementResponse, error) {
  amount := in.GetAmount()
  atomic.AddInt64(&counter, amount)
	return &pb.IncrementResponse{NewTotal: counter}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCounterServer(s, &server{})

  // Register reflection service on gRPC server.
  reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
