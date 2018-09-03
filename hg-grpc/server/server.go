package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	pb "hg-grpc/server/pb"

	"encoding/json"

	"google.golang.org/grpc"
)

var ip = "127.0.0.1"
var port = 50051

type HelloServer struct{}

type User struct {
	Uid  int    `json:"uid"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

//实现方法SayHello
func (h *HelloServer) SayHello(ctx context.Context, in *pb.HelloReq) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "hello," + in.Name,
	}, nil
}

func (h *HelloServer) GetInfo(ctx context.Context, in *pb.HelloReq) (*pb.Reply, error) {
	u := []User{
		{
			Uid:  1,
			Name: "xiaohei",
			Age:  23,
		},
	}

	json_data, _ := json.Marshal(u)
	return &pb.Reply{
		Code:    200,
		Message: "ok",
		Data:    string(json_data),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port)) //listen port
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	fmt.Println("service has run in ", port)
	s := grpc.NewServer()
	//注册服务
	pb.RegisterHelloServiceServer(s, &HelloServer{})
	s.Serve(lis)
}
