package main

import (
	"fmt"
	"os"
	"strings"
)

//从命令行获取参数
// 命令行参数会放置在切片 os.Args[] 中（以空格分隔），从索引1开始（os.Args[0] 放的是程序本身的名字，在本例中是 os_args）
//go run os_args.go heige
//hello,heige
func main() {
	welcome := "hello,"

	if len(os.Args) > 1 {
		welcome += strings.Join(os.Args[1:], " ")
	}

	fmt.Println(welcome)

}
