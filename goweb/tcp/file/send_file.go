package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	fmt.Println("请输入需要传入的文件: ")
	var path string
	fmt.Scan(&path)

	//获取文件名 info.Name()
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err: ", err)
		return
	}

	//连接服务端
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close() //关闭连接

	// 发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("conn.write err: ", err)
	}

	//接收服务器的回复
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if "ok" == string(buf[:n]) {
		var wg sync.WaitGroup
		wg.Add(1)
		go sendFile(path, conn, &wg)
		wg.Wait()
		return
	}
}

func sendFile(path string, conn net.Conn, wg *sync.WaitGroup) {
	fp, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer fp.Close()
	defer wg.Done()

	//读取文件内容，读多少就发送多少
	buf := make([]byte, 1024)
	for {
		n, err := fp.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("file has read finished")
				return
			} else {
				fmt.Println("read file err: ", err)
				return
			}
		}

		conn.Write(buf[:n])
		fmt.Println(string(buf[:n]))
		fmt.Println("send file content success")
	}

}
