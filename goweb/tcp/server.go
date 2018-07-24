package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var address string

func init() {
	flag.StringVar(&address, "address", ":1338", "tcp ip:port")
	flag.Parse()
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", address)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept() //接收客户端请求
		if err != nil {
			continue
		}

		go handleClient(conn)
		runtime.Gosched() //让出cpu,防止goroutine阻塞
	}
}

//处理客户端请求
func handleClient(conn net.Conn) {
	//设置请求超时,下面的for循环即会因为连接已关闭而跳出（即socket断开)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	request := make([]byte, 512) //每次处理512个字节
	defer conn.Close()           //关闭连接
	for {
		read_len, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte("req time:" + daytime))
		}

		fmt.Println("------")
		a := len(request)
		fmt.Println(string(request))
		fmt.Println("request len", a)
		fmt.Println(string(request[:read_len]))
		fmt.Println("read len", read_len)

		request = make([]byte, 128) // clear last read content

	}

}

func checkError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", e)
		os.Exit(1)
	}
}
