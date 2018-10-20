package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //等待用户链接
		if err != nil {
			log.Println("err: ", err)
			return
		}

		//接受客户端请求
		go handlerData(conn)
	}
}

func handlerData(conn net.Conn) {
	defer conn.Close() //关闭连接

	buf := make([]byte, 1024) //缓冲区1024byte
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("recevice data: ", string(buf[:n]))
	//发送信息给客户端
	conn.Write([]byte("hello client"))
}
