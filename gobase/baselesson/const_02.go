package main

import (
	"fmt"
)

const a = 12 //自行推导为int
const b int = 1

//一次多个值定义
const (
	c = 123.34 //自行推导为float64类型
	d = "daheige"
)

var (
	m = "dd"
	n = true
)

//枚举
const (
	a1 = iota
	b1
	c1
)

//遇到const,重置为0
const (
	p = iota
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("a type is %T,b type is %T\n", a, b) //int,int类型
	fmt.Println(c, d)
	fmt.Printf("c type is %T d type is %T\n", c, d)

	fmt.Println(m, n)
	fmt.Println(a1, b1, c1, p) //0 1 2 0

}
