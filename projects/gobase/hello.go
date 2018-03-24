package main

//1.以包作为管理单位，每个文件必须先声明包名
//2.程序要运行，必须有一个main包

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println("heige")

	fmt.Printf("%.2f\n", 12.356) //格式化处理
	fmt.Printf("%2.2f", 1.234)
}
