package main

import (
	"fmt"
)

//常量在编译的时候已经确定了值和范围
type week int

const (
	Sunday week = iota //第一行iota初始值是0，以后依次加1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	a = iota //0
	b        //1
	c = iota //2
	d        //3
	e = 12
	f = iota //5，因为e自增了iota
	g        //6
	h int    = 456
	i        //当没有遇到iota，下一行的i和上一行的h类型和值都是一样，456
	m = iota //9
)

//在遇到下一个const后，iota初始化的值又置为0
const (
	hg = iota
	hg_x
)

func main() {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(a, b, c, d, e, f, g, h, i) //0 1 2 3 12 5 6 456 456
	fmt.Println(m)
	fmt.Println(hg, hg_x)
}
