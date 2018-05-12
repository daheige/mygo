package main

import (
	"fmt"
)

func main() {
	var x = 1 //自动推断变量类型
	fmt.Println(x)
	fmt.Printf("%T\n", x) //int

	y := 123.45           //自动推导类型
	fmt.Printf("%T\n", y) //float64

	//同事声明多个变量
	m, n := 1, 12.4
	fmt.Println(m, n)
	m = 23 //重新赋值

	fmt.Println(m)
	fmt.Printf("this is %s\n", "heige")

	//多重变量和匿名变量
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)

	//变量交换值
	b, a = a, b
	fmt.Println(a, b)

	c, _ = a, c //_表示丢弃变量,不要使用的情况下可以用_

	fmt.Println(a, b, c)
}
