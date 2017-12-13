package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开连接
	//在网络编程中net.Dial函数是非常重要的，一旦你连接到远程系统，就会返回一个Conn类型接口，我们可以用它发送和接收数据。
	//Dial函数巧妙的抽象了网络结构及传输。所以IPv4或者IPv6，TCP或者UDP都可以使用这个公用接口
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		//由于目标主机拒绝连接而无法创建连接
		fmt.Println("error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin) //创建一个缓冲读取器
	fmt.Println("first,what is your name")
	clientName, _ := inputReader.ReadString('\n')
	fmt.Println("name is", clientName)
	trimName := strings.Trim(clientName, "\r\n") //去掉换行符

	//发送消息给服务器直到退出
	for {
		fmt.Println("what to send data:")
		input, _ := inputReader.ReadString('\n')
		fmt.Println("your input is", input)
		trimInput := strings.Trim(input, "\r\n")
		if trimInput == "quit" {
			fmt.Println("quit success")
			return
		}
		_, err = conn.Write([]byte(trimName + " say " + trimInput))
	}
}
