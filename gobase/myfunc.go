package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//返回多个值
func sumAndPro(a, b int) (int, int) {
	return a + b, a * b
}

// 变参，多个参数 在函数体中，变量arg是一个int的slice
func restVal(arg ...int) {
	//for---range遍历
	for key, val := range arg {
		fmt.Printf("key is %d and value is %d\n", key, val)
	}
}

//将指针作为函数的参数
// 变量在内存中是存放于一定地址上的，修改变量实际是修改变量地址处的内存
//defer在函数退出之前调用 defer是采用后进先出模式
func updateVal(a *int) {
	defer fmt.Println(*a, "将发生改变")
	*a = *a + 1
}

//将函数作为值，类型
type testFn func(int) bool //定义testFn只一个函数运行之后的结果类型

func isEven(a int) bool {
	return a%2 == 0
}

func filterNum(arg []int, f testFn) (res []int) {
	for _, val := range arg {
		if f(val) {
			res = append(res, val)
		}
	}
	return
}
