package main

import (
	"context"
	"grpc-golang-server/calculator"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

// Sum implements calculator.Sum
func (s *server) Sum(ctx context.Context, in *calculator.Numbers) (*calculator.Result, error) {
	sum := in.GetNumber1() + in.GetNumber2()
	return &calculator.Result{Result: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServer(s, &server{})

	log.Printf("gRPC server listening at port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
