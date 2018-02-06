package main

import (
	"fmt"
)

func main() {
	fmt.Println("变参的本质")
	f(1, 3, "fefefe", 'x', true, []int{1, 2})
	res := sum(1, 23, 4)
	fmt.Println("这几个数字的和", res)
}

//Go语言中，空接口可以指向任何数据对象，所以可以使用interface{}定义任意类型变量，同时interface{}也是类型安全的
func f(a ...interface{}) { //变参实际内部是一个切片
	// fmt.Println(a)
	other := make([]interface{}, 0)
	for _, arg := range a {
		switch v := arg.(type) { //类型检测
		case int:
			fmt.Println(arg, "类型是int")
		case string:
			fmt.Println(arg, "类型是string")
		case float64:
			fmt.Println(arg, "类型是float")
		case bool:
			fmt.Println(arg, "类型是boolean")
		default:
			// fmt.Println(v) //具体的值
			other = append(other, v)
		}
	}
	fmt.Println("其他类型:", other)
}

//变参的声明a...typeName 类型名称
//指定了函数返回值，就不需要用return显式返回
func sum(a ...int) (res int) {
	for _, v := range a {
		res += v
	}
	return
}
