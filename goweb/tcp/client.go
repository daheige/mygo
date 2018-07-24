package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

var address string

func init() {
	flag.StringVar(&address, "address", ":1338", "tcp ip:port")
	flag.Parse()
}

func main() {
	// tcpAddr, err := net.ResolveTCPAddr("tcp4", address) //绑定ip,port
	// checkError(err)

	// conn, err := net.DialTCP("tcp", nil, tcpAddr)
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	checkError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	checkError(err)

	result, err := ioutil.ReadAll(conn) //读取响应结果
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", e)
		os.Exit(1)
	}
}
