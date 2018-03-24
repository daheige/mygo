package main

import (
	"fmt"
)

//go 语言
//1 变量类型写在变量之后
//2 编译器可以推断变量类型
//3 没有char,只有rune
//4 原生支持复数类型

func varZeroVal() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

//在包里面定义全局变量
var hg_name = "heige"
var hg_age = 23

func main() {
	var a int = 12 //变量声明了不用会编译不通过
	fmt.Println(a)

	a = 13
	fmt.Println(a)

	//变量的初始化，声明并赋值
	var b int = 10
	b = 35
	fmt.Println(b)

	//自动推断类型
	var c = 12
	fmt.Println(c)
	hg_c := 333
	fmt.Printf("hg_c type is %T\n", hg_c) //int类型
	fmt.Println(hg_c)

	//格式化输出
	fmt.Printf("num is %d\n", 345)

	varZeroVal()

	//同时定义多个变量
	var hg_a, hg_b int = 2, 3
	fmt.Println(hg_a, hg_b)

	//变量自行推断
	var hg_x, hg_y = 1, "fefe"
	fmt.Println(hg_x, hg_y)

	//简写模式（只能在函数中使用）
	hg_m, hg_n := 3, true
	fmt.Println(hg_m, hg_n)

	fmt.Println(hg_name, hg_age)

}
