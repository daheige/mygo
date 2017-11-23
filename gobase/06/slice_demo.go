package main

//声明切片
//一个切片三要素 1.指向底层数组的指针，长度，容量
//  new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址
//  这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体，它相当于 &T{}。
//  make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。
//  换言之，new 函数分配内存，make 函数初始化

import (
	"fmt"
)

func main() {
	// 1 字面量方式声明一个切片
	var s = []int{1, 3, 4}
	fmt.Println(s)

	// 2 通过make方式声明
	var s2 = make([]int, 2) //长度为2 初始化为[0 0]
	fmt.Println(s2)

	//3 指定容量
	var s3 = make([]int, 3, 5) //该切片的底层是一个数组[5]int
	fmt.Println(s3)

	//4 通过从一个底层数组的方式来声明一个切片
	var arr = [...]int{1, 3, 4, 5, 6, 8}
	var s4 = arr[0:4]

	fmt.Println(s4)

	//5 通过new()函数来分配一个底层数组，然后截取子数组形成一个切片

	var s5 = new([5]int)[0:4] //new用来分配内存
	fmt.Printf("s5是%v s5内存地址是%p\n", s5, &s5)

	//使用new来声明一个数组
	var s6 = new([6]int)
	fmt.Println("new函数返回一个地址", s6, "s6这个地址存放的值是", *s6)
}
