package main

import (
	"bufio"
	"fmt"
	"io"
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

	//逐行读取文件内容
	fp, err := os.Open("../demo/out.md")
	if err != nil {
		fmt.Println("文件读取失败", err)
		return
	}
	//创建一个缓冲读取器
	bufReader := bufio.NewReader(fp)
	for {
		s, err := bufReader.ReadString('\n')
		if err == io.EOF {
			break
		}

		fmt.Println(s)
	}
	defer fp.Close()

}
