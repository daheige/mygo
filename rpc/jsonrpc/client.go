//jsonrpc客户端
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type serverResponse struct {
	Id     *json.RawMessage `json:"id"`
	Result interface{}      `json:"result"`
	Error  interface{}      `json:"error"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1336")
	if err != nil {
		log.Fatal("net.dial error", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)) //创建一个jsonrpc的客户端
	var reply string
	err = client.Call("HelloService.Hello", "daheige", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("返回的结果", reply)

}

//模拟请求
//$echo -e '{"method":"HelloService.Hello","params":["daheige"],"id":1}' | nc localhost 1336
