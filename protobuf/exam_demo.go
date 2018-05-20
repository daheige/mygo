package main

import (
	"fmt"
	"os"
	pb "protobuf/exam"

	proto "github.com/golang/protobuf/proto"
)

func main() {

	msg := &pb.HelloWorld{
		Id:  proto.Int32(996),
		Str: proto.String("daheige"),
	}

	path := "./test.txt"
	f, err := os.Create(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer f.Close()

	//序列化成pb数据格式
	buffer, err := proto.Marshal(msg)
	f.Write(buffer)
	fmt.Println("write success")
}
