package main

import (
	"fmt"
	"log"
	// 辅助库
	"github.com/golang/protobuf/proto"

	// test.pb.go 的路径
	"protobuf/example"
)

func main() {
	// 创建一个消息 Test
	test := &example.Test{
		// 使用辅助函数设置域的值
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	fmt.Printf("Label:%s Type:%d\n", test.GetLabel(), test.GetType())

	*(test.Label) = "hello go"
	*(test.Type) = 18

	// 进行编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Printf("Binary Len:%d\n", len(data))

	// 进行解码
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Printf("Label:%s Type:%d\n", test.GetLabel(), test.GetType())

	// 测试结果
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
