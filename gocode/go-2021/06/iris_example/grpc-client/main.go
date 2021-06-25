package main

import (
	"context"
	"gocode-chb/gocode/go-2021/06/iris_example/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:443"
	defaultName = "world"
)

func main() {
	cred, err := credentials.NewClientTLSFromFile("../server.crt", "localhost")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(cred), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
