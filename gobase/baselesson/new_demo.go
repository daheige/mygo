package main

import (
	"fmt"
)

func main() {
	var p *int //p是一个指针变量,类型是*int,指向int类型的指针

	p = new(int) //new分配内存空间,必须有指向的指针变量才可以赋值

	*p = 12
	fmt.Printf("p type is %T \n value is %v \n", p, p)
	// p type is *int
	//  value is 0xc420014128

	fmt.Println(*p)
	a := 12
	p = &a //存放的整型int变量的内存地址
	a++
	fmt.Println(a, *p) //13,13
}
