package main

import (
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

const HelloServiceName = "HelloService"

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

//优化后的client,更加安全
func main() {
	client, err := DialHelloService("tcp", "localhost:1339")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply) //直接传递req,res
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reply)
}

func test() {
	client, err := rpc.Dial("tcp", "localhost:1339")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string                                           //响应结果
	err = client.Call("HelloService.Hello", "daheige", &reply) //请求结果放在reply中
	if err != nil {
		log.Println("dialing:", err)
		return
	}

	log.Println(reply)

}
