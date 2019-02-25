package library

import (
	"fmt"
	"log"
	pb "mygo/protobuf/helloworld"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

type Greeter struct {
	Address string
}

func (this *Greeter) getClient() (c pb.GreeterClient) {
	//create a connection
	conn, err := grpc.Dial(this.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// defer conn.Close() //二次利用这里不用defer close
	c = pb.NewGreeterClient(conn) //create Greet service Client
	return
}

func (this *Greeter) GetUserInfo(Uid int64, Level int32) (u *pb.User) {
	conn, err := grpc.Dial(this.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn) //create Greet service Client

	//ua必须转换为int32
	r, err := c.GetUserInfo(context.Background(), &pb.GetUserReq{
		Uid:   Uid,
		Level: Level,
	})

	if err != nil {
		fmt.Println(err)
		log.Fatalf("getUserinfo error: %v", err)
		return
	}

	u = r.GetUserInfo()
	return
}
