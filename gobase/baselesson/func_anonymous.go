package main

import (
	"fmt"
)

func main() {
	f := func(a, b int) int {
		return a + b
	}

	fmt.Println(f(1, 2))

	f2 := bibo() //结果是一个匿名函数
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}

//匿名函数形成闭包环境
//闭包本质是保存了函数执行环境(内部函数执行的生命周期)
func bibo() func() int {
	var a = 1
	return func() int {
		a = a * 2
		return a
	}
}
