package main

import (
	"context"
	"fmt"
	"grpc-golang-server/calculator"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"google.golang.org/grpc"
)

const (
	grpcPort  = ":50051"
	htttpPort = 8080
)

var wg sync.WaitGroup

type server struct{}

// Sum implements calculator.Sum
func (s *server) Sum(ctx context.Context, in *calculator.Numbers) (*calculator.Result, error) {
	sum := in.GetNumber1() + in.GetNumber2()
	return &calculator.Result{Result: sum}, nil
}

func handleSum(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		number1, _ := strconv.Atoi(request.URL.Query().Get("number1"))
		number2, _ := strconv.Atoi(request.URL.Query().Get("number2"))
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(strconv.Itoa(number1 + number2)))
	} else {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func startHTTPServer() {
	log.Printf("HTTP Server started at port %d", htttpPort)

	http.HandleFunc("/sum", handleSum)

	err := http.ListenAndServe(fmt.Sprintf(":%d", htttpPort), nil)
	if err != nil {
		log.Fatalf("Error starting HTTP Server %v", err)
	}
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServer(s, &server{})

	log.Printf("gRPC server listening at port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	wg.Add(1)
	go startHTTPServer()
	go startGrpcServer()
	wg.Wait()
}
