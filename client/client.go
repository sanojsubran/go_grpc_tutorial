package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sanojsubran/pipe"
	"google.golang.org/grpc"
)

func main() {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:8793", opts)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer cc.Close()

	client := pipe.NewGreeterClient(cc)

	request := &pipe.HelloRequest{
		Name: "John Lark",
	}

	resp, err := client.SayHello(context.Background(), request)
	if nil != err {
		log.Fatal("No response from server")
		os.Exit(1)
	}
	fmt.Printf("Received response: %v", resp.Message)

}
