package main

import (
	"io/ioutil"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", ":8080", 3*time.Second)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	//发送数据给server
	_, err = conn.Write([]byte("daheige hello"))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("send data success")
	//读取服务端发送过来的响应结果
	result, err := ioutil.ReadAll(conn) //读取响应结果
	log.Println("recevice server data: ", string(result))
}
