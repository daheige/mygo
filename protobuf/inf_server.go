package main

// server.go
import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "protobuf/helloworld"
)

const (
	port = ":50051"
)

//struct
type server struct{}

//method
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
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

// go run inf_client.go  daheige
// $ go run inf_client.go  daheige321
// 2018/05/19 23:00:30 Greeting: Hello daheige321
