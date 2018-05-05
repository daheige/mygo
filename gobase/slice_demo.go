package main

import (
	"fmt"
)

func main() {
	// 创建一个整型切片
	// 其长度为 3 个元素,容量为 5 个元素
	slice := make([]int, 3, 5)

	fmt.Printf("slice address is %p\n", slice)

	slice = append(slice, []int{1, 3, 4, 5, 6, 7}...) //扩容为10,原来的2倍

	fmt.Printf("slice value is %v\n", slice)
	fmt.Printf("slice address is %p,cap is %d\n", slice, cap(slice)) //地址发生了改变

	slice[1] = 345 //修改切片的值
	fmt.Println(slice)
	hgp(slice)

	fmt.Println(slice) //值没有发生改变
	// 如果传递的是切片的指针就会发生改变
	hgp2(&slice)
	fmt.Println("改变之后的slice", slice)
}

//由于与切片关联的数据包含在底层数组里,不属于切片本身,所以将切片
// 复制到任意函数的时候,对底层数组大小都不会有影响。复制时只会复制切片本身
// 不会涉及底层数组
func hgp(a []int) { //传的是值
	a = append(a, 1, 2, 3)
	fmt.Println(a)
}

func hgp2(a *[]int) {
	*a = append(*a, 1, 2, 3)
	fmt.Println(a) //a是一个切片指针
	fmt.Println((*a)[1])
}

/**
运行结果:
slice address is 0xc42007e030
slice value is [0 0 0 1 3 4 5 6 7]
slice address is 0xc420084050,cap is 10
关于扩容
在切片的容量小于 1000 个元素时,总是
会成倍地增加容量。一旦元素个数超过 1000,容量的增长因子会设为 1.25,也就是会每次增加 25%
的容量。随着语言的演化,这种增长算法可能会有所改变

切片所占内存大小:
在 64 位架构的机器上,一个切片需要 24 字节的内存:指针字段需要 8 字节,长度和容量字段分别需要 8 字节

在函数间传递 24 字节的数据会非常快速、简单。这也是切片效率高的地方。不需要传递指
针和处理复杂的语法,只需要复制切片,按想要的方式修改数据,然后传递回一份新的切片副本。
*/
