/**
 * 一个（web）服务器应用需要响应众多客户端的并发请求：go会为每一个客户端产生一个协程用来处理请求。
 * 我们需要使用net包中网络通信的功能。它包含了用于TCP/IP以及UDP协议、域名解析等方法
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting server...")
	server, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("error listening", err.Error())
		return
	}

	//监听客户端的连接
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
			return
		}
		go doSomeThing(conn) //go会为每一个客户端产生一个协程用来处理请求

	}
}

func doSomeThing(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		l, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading")
			return
		}
		fmt.Printf("received data %v\n", string(buf[:l]))

	}
}

/**
 * 先启动server.go
 * $ go run server.go
starting server...

再启动client.go
go run client.go
irst,what is your name
sss
name is sss

what to send data:
ll
your input is ll

what to send data:
daheige
your input is daheige

what to send data:
诸位
your input is 诸位

我们在server.go运行的终端可以看到
starting server...
received data sss say ll
received data sss say daheige
received data sss say 诸位
*/
