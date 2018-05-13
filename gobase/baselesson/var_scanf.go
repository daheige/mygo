package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Printf("请输入a: ")

	//接收键盘输入的值
	// fmt.Scanf("%d", &a)
	fmt.Scan(&a)
	fmt.Printf("a = %d", a)
}
