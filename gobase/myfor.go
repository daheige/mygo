package main

import (
	"fmt"
)

//函数返回一个切片
func getData() []int {
	defer println("origin data")
	return []int{1, 2, 3, 4}
}

func init() {
	println("初始化。。。")
}

func main() {
	data := []string{"abc", "ffe", "123"}
	for val := range data {
		fmt.Println("val is", val)
		fmt.Println("val address is", &val) //val address is 0xc4200120f8局部变量的地址不会发生改变，会重复使用
	}

	fmt.Println("for success")

	//for遍历的目标表达式是一个函数执行一次
	//_是忽略值
	for _, val := range getData() {
		fmt.Println("val is ", val)
	}

	x := 1
	switch x {
	case 1:
		println("1")
	default:
		println("not 1")
	}

	i := 0
	for i < 10 {
		if i%2 == 0 {
			fmt.Println("x is ", i)
		}

		if i == 8 {
			break
		}

		i++
	}

}
