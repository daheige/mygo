package main

import (
	"fmt"
)

func main() {
	var arr = [5]int{1, 2, 3, 4, 5} //定义一个数组（整型数组）
	fmt.Println(arr)
	array := [5]int{1, 3, 2, 4, 5} //字面量形式定义，长度为5
	fmt.Println(array)

	arr2 := [...]int{3, 4, 5}
	fmt.Println("不固定长度，由初始化值决定", len(arr2), arr2)

	//声明指针数组
	arr3 := [2]*int{0: new(int), 1: new(int)}

	//赋值
	*arr3[0] = 10
	*arr3[1] = 12
	fmt.Println(arr3) //[0xc420012150 0xc420012158] 数组的每个元素是一个指针
}
