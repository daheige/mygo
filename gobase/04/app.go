package main

import (
	"fmt"
)

type IZ int

const PI = 3.14
const HG = "heige"

//常量枚举
const (
	HG_A = iota
	HG_B
	HG_C
)

func main() {
	var a IZ = 5
	c := int(a)
	fmt.Println("c is", c)
	d := IZ(c)
	fmt.Println("d is", d)
	fmt.Println("pi is", PI)
	fmt.Println("HG is", HG)
	fmt.Println("pi/3 is", PI/3) //浮点数
	//声明多行字符串
	s := `feefef
    fefes
    sss
    ss`
	fmt.Println(s)
	fmt.Println(HG_A, HG_B, HG_C) //0,1,2

}
