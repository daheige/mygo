package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "protobuf/helloworld"
	"strconv"
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

	var ua int
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	if len(os.Args) > 2 {
		//字符串转换为int
		if ua, err = strconv.Atoi(os.Args[2]); err != nil {
			log.Fatal("错误的ua")
			return
		}
	}

	//ua必须转换为int32
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name, Ua: int32(ua)})
	if err != nil {
		fmt.Println(err)
		log.Fatalf("could not greet: %v", err)
	}

	//reply data
	fmt.Println("reply data: ", r.GetMessage())
	log.Printf("Greeting: %s", r.Message)
}
