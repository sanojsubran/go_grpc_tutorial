package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sanojsubran/pipe"
	"google.golang.org/grpc"
)

type server struct {
	pipe.UnimplementedGreeterServer
}

func (*server) SayHello(ctx context.Context, request *pipe.HelloRequest) (*pipe.HelloReply, error) {
	name := request.Name
	response := &pipe.HelloReply{
		Message: "Hello " + name,
	}
	fmt.Printf("\nReceived request with name: %v", name)
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8793")
	if nil != err {
		log.Fatalf("\nError: %v", err)
	}
	fmt.Printf("\nServer is listening on port 127.0.0.1:8793")
	s := grpc.NewServer()
	pipe.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}
