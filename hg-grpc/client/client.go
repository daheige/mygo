package main

import (
	"fmt"
	pb "hg-grpc/server/pb"
	"log"
	"time"

	"encoding/json"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type User struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	//建立grpc连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(10*time.Second)) //10s超时限制
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewHelloServiceClient(conn) //create Greet service Client

	name := "xiaoming"
	r, err := client.SayHello(context.Background(), &pb.HelloReq{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	//reply data
	fmt.Println("reply data: ", r.GetMessage())

	reply, err := client.GetInfo(context.Background(), &pb.HelloReq{Name: "heige"})
	fmt.Printf("%T\n", reply)
	fmt.Println("reply data:", reply)
	fmt.Println(reply.GetCode())
	fmt.Println(reply.GetMessage())
	fmt.Println(reply.GetData())

	u := []User{}
	//将data解析到结构体中
	err = json.Unmarshal([]byte(reply.Data), &u)
	fmt.Println(u, err)
	fmt.Println(u[0])
}
