package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writeFile("test.md", "daheige你好")
}

func writeFile(filename string, str string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() //从缓存区把内容写入文件中
	writer.WriteString(str)

	fmt.Println("写入成功!")
}
