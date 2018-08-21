package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
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
	flag.IntVar(&port, "port", 1339, "rpc server port")
	flag.Parse()

	//注册服务名称
	rpc.RegisterName("HelloService", &HelloService{})
}

func main() {
	http.HandleFunc("/json-rpc", jsonTest)

	address := fmt.Sprintf("%s:%d", ip, port)
	fmt.Println("server has run on ", address)
	http.ListenAndServe(address, nil)
}

func jsonTest(w http.ResponseWriter, r *http.Request) {
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}

	rpc.ServeRequest(jsonrpc.NewServerCodec(conn)) //发送请求
}

//发送curl请求
//$ curl localhost:1339/json-rpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
// {"id":0,"result":"hello:hello","error":null}
//id不一定需要
// curl localhost:1339/json-rpc -X POST --data '{"method":"HelloService.Hello","params":["daheige"]}'
