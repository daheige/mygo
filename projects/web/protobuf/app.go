package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"projects/web/protobuf/server/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	u := &pb.UserInfoReq{
		Name: "hello",
		Age:  28,
	}

	fmt.Println(u)
	json_data, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}

	json_str := string(json_data)
	fmt.Println(json_str)

	//将json字符串转换为结构体对应的类型
	userInfo := &pb.UserInfoReq{}
	BufUnmarshal("json", []byte(json_str), userInfo)
	fmt.Println("json decode data: ", userInfo)
	fmt.Println(userInfo.Name)

	user := &pb.UserInfoReq{
		Name: "hello",
		Age:  28,
	}

	// 进行编码
	buf, err := proto.Marshal(user)
	fmt.Println(buf, err)

	//pb 解码到pb.Message
	obj := &pb.UserInfoReq{}
	// proto.Unmarshal(buf, obj)
	BufUnmarshal("pb", buf, obj)
	fmt.Println("pb decode data: ", obj)
}

func BufUnmarshal(t string, buf []byte, data interface{}) error {
	if t == "json" {
		return json.Unmarshal(buf, data)
	}

	if t == "pb" {
		return proto.Unmarshal(buf, data.(proto.Message)) //类型断言
	}

	return errors.New("message type error")
}
