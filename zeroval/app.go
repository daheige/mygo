package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var a [1]int
	var b uint8 = 12

	fmt.Println("int arr 0值", a)

	fmt.Println("max int 32范围", math.MaxInt32)

	fmt.Println(b)

	var flag = true
	fmt.Println(flag)

	//go 自行推断变量的类型
	c := 12 //最简单的方式 只能在函数中声明
	fmt.Println(c)

	d, e, f := 1, 2, 3
	fmt.Println(d, e, f)

	var x int
	x = 123
	var y float64 = 1.0
	z := y / 3

	fmt.Println(x, y, y/3, z)

	var hg_a = 65
	fmt.Println(string(hg_a)) //转化为字符A
	// 如果真的需要转化为字符串
	fmt.Println(strconv.Itoa(hg_a))
	hg_b, _ := strconv.Atoi("65") //将字符串数字转换为数字
	fmt.Println("65 to int 65", hg_b)
	//将浮点数转换为字符串
	hg_c := strconv.FormatFloat(123.9899, 'f', 2, 32)

	fmt.Println(hg_c)

}
