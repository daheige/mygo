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

		//开启独立携程处理接受客户端请求
		go handlerData(conn)
	}
}

func handlerData(conn net.Conn) {
	defer conn.Close() //关闭连接

	log.Println("client connection success")
	buf := make([]byte, 1024) //缓冲区1024byte
	//循环读取客户端发送过来的数据
	for {
		n, err := conn.Read(buf) //这里包含了EOF
		if err != nil {
			log.Println(err)
			return
		}

		if "exit" == string(buf[:n-1]) {
			log.Println("client has closed!")
			return
		}

		log.Println("recevice data: ", string(buf[:n-1]))
		//发送信息给客户端
		conn.Write([]byte("hello client"))
	}
}
