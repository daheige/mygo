package main

import (
	"fmt"
)

var a = 1 // 全局变量在所有的地方有效

func main() {
	//块作用域
	{
		fmt.Println(1)
	}

	if a := 23; a < 12 {
		fmt.Println("a < 12")
	} else if a > 20 {
		fmt.Println("a > 20")
	}

	fmt.Println(a) //全局变量
	test()

}

func test() {
	fmt.Println(a)
}
