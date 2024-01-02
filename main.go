package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/ehakan/dummy-grpc-server/greeter"
	"google.golang.org/grpc"
)

type greeterServer struct {
	greeter.UnimplementedDummyGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	return &greeter.HelloResponse{
		Greeting: "Hello, " + req.GetName(),
	}, nil
}

func (s *greeterServer) StreamGreetings(req greeter.DummyGreeter_StreamGreetingsServer) error {
	return errors.New("unimplemented")
}

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	greeter.RegisterDummyGreeterServer(grpcServer, &greeterServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
