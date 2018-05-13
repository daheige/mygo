package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 123.34
	c := 'a'
	d := "heige"
	fmt.Printf("%T,%T,%T,%T\n", a, b, c, d) //int,float64,int32,string
	fmt.Printf("%v,%v,%v,%v", a, b, c, d)   //自动转换值 1,123.34,97,heige
	fmt.Printf("c is %c", c)
	fmt.Printf("a address %p", &a)
}

/**
格式化说明符
    %d 整型
    %c 字符
    %s 字符串
    %f 浮点型
    %T 变量类型打印
    %v 自动转换值
    %p 打印内存地址 比如&a
**/
