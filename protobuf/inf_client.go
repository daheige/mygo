package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "protobuf/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	//create a connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn) //create Greet service Client

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
