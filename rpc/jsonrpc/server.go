package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	flag.IntVar(&port, "port", 1336, "rpc server port")
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
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //传入的参数是针对服务端的json编解码器
	}
}

/**
先启动go run server.go

然后调用
$echo -e '{"method":"HelloService.Hello","params":["daheige"],"id":1}' | nc localhost 1336
返回结果{"id":1,"result":"hello:daheige","error":null}

返回的json数据也是对应内部的两个结构体：客户端是clientResponse，服务端是serverResponse。两个结构体的内容同样也是类似的：
type clientResponse struct {
    Id     uint64           `json:"id"`
    Result *json.RawMessage `json:"result"`
    Error  interface{}      `json:"error"`
}

type serverResponse struct {
    Id     *json.RawMessage `json:"id"`
    Result interface{}      `json:"result"`
    Error  interface{}      `json:"error"`
}
因此无论采用何种语言，只要遵循同样的json结构，以同样的流程就可以和Go语言编写的RPC服务进行通信

*/
