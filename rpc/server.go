package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	fmt.Println("client req data", request)
	*reply = "hello:" + request
	return nil
}

var port int
var ip string

func init() {
	flag.StringVar(&ip, "ip", "127.0.0.1", "rpc server ip")
	flag.IntVar(&port, "port", 1339, "rpc server port")
	flag.Parse()

	//注册服务名称
	rpc.RegisterName("HelloService", &HelloService{})
}

func main() {
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatal("listen server error")
	}

	log.Println("server has run on", port)
	//接收客户端的请求
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		//对于每个请求,开启goroutine的方式处理
		//通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务
		go rpc.ServeConn(conn)
	}
}

/**
运行结果
先运行server.go
$ go run server.go
2018/08/21 21:37:52 server has run on 1339

再运行客户端go run client.go
$ go run client.go
2018/08/21 21:38:29 hello:hello
*/
