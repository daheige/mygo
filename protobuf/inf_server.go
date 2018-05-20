package main

// server.go
import (
	"errors"
	"fmt"
	"log"
	"net"
	pb "protobuf/helloworld"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

//struct
type server struct{}

//method
func (s *server) SayHello(ctx context.Context, input *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("request data:", input)
	if errUa := s.CheckUa(input.Ua); errUa.Code != 200 {
		return nil, errors.New(errUa.Message)
	}

	return &pb.HelloReply{Message: "Hello " + input.Name}, nil
}

func (s *server) GetUserInfo(ctx context.Context, input *pb.GetUserReq) (*pb.GetUserReply, error) {
	fmt.Println("get user info input data: ", input)
	return &pb.GetUserReply{
		UserInfo: &pb.User{
			Uid:  1,
			Name: "daheige",
			Job:  "golang",
			Age:  28,
		},
	}, nil
}

func (s *server) CheckUa(ua int32) (errUa pb.ErrorUa) {
	if ua == 0 {
		errUa.Code = 500
		errUa.Message = "ua error"
		return
	}
	if _, ok := pb.Ua_name[ua]; !ok {
		errUa.Code = 500
		errUa.Message = "ua error2"
		return
	}

	errUa.Code = 200
	return
}

func main() {
	fmt.Println("service has run in ", port)
	lis, err := net.Listen("tcp", port) //listen port
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{}) //register service Greeter
	s.Serve(lis)

}

//run
//go run inf_server.go
// service has run in  :50051

// go run inf_client.go  daheige 1
// $ go run inf_client.go  daheige321 2
// 2018/05/19 23:00:30 Greeting: Hello daheige321
