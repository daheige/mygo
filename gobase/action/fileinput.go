package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	inputfile, err := os.Open("../demo/input.md")
	if err != nil {
		fmt.Println("file is exists!")
		return
	}

	defer inputfile.Close()
	inputReader := bufio.NewReader(inputfile) //获得一个读取器变量
	for {
		inputStr, readerError := inputReader.ReadString('\n') //将文件的内容逐行（行结束符 '\n'）读取出来
		if readerError == io.EOF {
			return
		}

		fmt.Printf("the input was %s", inputStr)
	}

}
