package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("../demo/input.md")
	if err != nil {
		fmt.Println("文件读取失败，error:", err)
	}

	inputReader := bufio.NewReader(inputFile) //创建一个读取器
	// 将内容读取到缓冲中
	buf := make([]byte, 1024)
	for {
		n, _ := inputReader.Read(buf)
		if n == 0 {
			break
		}
	}

	fmt.Println(string(buf))

}
