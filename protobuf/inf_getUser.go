package main

import (
	"fmt"
	"protobuf/library"
)

func main() {
	greeterService := library.Greeter{
		Address: "localhost:50051",
	}
	for i := 0; i < 10; i++ {
		userInfo := greeterService.GetUserInfo(1, 2)
		fmt.Println("reply data: ", userInfo)
	}
}
