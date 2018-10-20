package main

import (
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close() //关闭连接

	//从键盘输入内容，给服务器发送内容
	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				log.Println("os.Stdin err: ", err)
				return
			}

			//把输入的内容给服务器发送
			conn.Write(str[:n])
		}
	}()

	//主携程不停接受服务器端发送过来的数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read err: ", err)
			return
		}

		log.Println(string(buf[:n])) //打印接受到数据
	}

}
