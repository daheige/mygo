package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"protobuf/Im"
)

func main() {
	//create a msg
	test := &Im.Helloworld{
		Str: "hello daheige",
		Opt: 12345,
	}

	test.Id = 1

	//encode proto data
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshal pb error:", err)
	}

	fmt.Println(data) //[8 1 18 13 104 101 108 108 111 32 100 97 104 101 105 103 101 24 185 96]

	//decode proto data
	newTest := &Im.Helloworld{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshal error: ", err)
	}

	log.Printf("id:%d;opt:%d;str:%s\n", newTest.Id, newTest.Opt, newTest.Str)

	//testing result
	if test.String() != newTest.String() {
		log.Fatalf("data mismatch %q != %q", test.String(), newTest.String())
	}
}
