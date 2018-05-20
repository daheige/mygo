package main

import (
	"fmt"
	"io"
	"os"
	pb "protobuf/exam"

	proto "github.com/golang/protobuf/proto"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func main() {
	path := "test.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed: %s\n", err)
		return
	}

	defer file.Close()
	fi, err := file.Stat()
	CheckError(err)

	//创建缓存区
	buffer := make([]byte, fi.Size())
	_, err = io.ReadFull(file, buffer)
	CheckError(err)

	msg := &pb.HelloWorld{}

	//发序列化pb为msg 结构体
	err = proto.Unmarshal(buffer, msg)
	CheckError(err)

	fmt.Printf("read: %s\n", msg.String())
}
