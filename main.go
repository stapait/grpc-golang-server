package main

import (
	"context"
	"grpc-rest-benchmark/calculator"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

var wg sync.WaitGroup

// Sum implements calculator.Sum
func (s *server) Sum(ctx context.Context, in *calculator.Numbers) (*calculator.Result, error) {
	sum := in.GetNumber1() + in.GetNumber2()
	return &calculator.Result{Result: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
package = "calculator";
package = "calculator";
package = "calculator";
package = "calculator";
package = "calculator";
package = "calculator";
package = "calculator";
package = "calculator";
package = "calculator";
	wg.Add(1)

	go func() {
		reflection.Register(s)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("gRPC server listening at port %s", port)
	wg.Wait()
}
