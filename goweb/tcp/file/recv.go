package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listenner, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.listen err: ", err)
		return
	}
	defer listenner.Close()

	for {
		//阻塞用户连接
		conn, err := listenner.Accept() //建立连接
		if err != nil {
			fmt.Println("net.Accept err: ", err)
			return
		}

		//接收客户端文件内容
		//每一个客户端连接都是一个独立的携程处理
		go recvFile(conn)
	}
}

func recvFile(conn net.Conn) {
	defer conn.Close() //关闭连接

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("connect.read err: ", err)
		return
	}

	//接收到文件名回复ok
	filename := string(buf[:n])
	conn.Write([]byte("ok"))

	f, err := os.Create("recv_" + filename)
	if err != nil {
		fmt.Println("os.Create file err: ", err)
		return
	}
	defer f.Close()

	buf = make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println(filename, "file has recv end")
				return
			} else {
				fmt.Println("read content err: ", err)
				return
			}
		}

		f.Write(buf[:n])
	}
}
