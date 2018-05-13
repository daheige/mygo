package main

import (
	"fmt"
)

const LIM = 40

//预分配数组个数和内存
var fiboRes [LIM]uint64

type funcType func(int, int) int

func add(a, b int) int {
	return a + b
}

func reduce(a, b int) int {
	return a - b
}

func ride(a, b int) int {
	return a * b
}

//函数类型作为参数传递给另一个函数
//该函数可以实现多个形式调用
func callc(a, b int, f funcType) int {
	return f(a, b) //回调函数f
}

func main() {
	res := my("daheige")
	fmt.Println(res)
	a, b, c := myFun()
	fmt.Println(a, b, c)
	x, y, z := myFun2()
	fmt.Println(x, y, z)
	fmt.Println(sumInt(5))
	fmt.Println("=======fibo 40 times====")
	/*var arr [LIM]uint64
	  for i := 0; i < LIM; i++ {
	      arr[i] = fibo(i)
	  }

	  fmt.Println(arr)*/

	fmt.Println(fiboArr(40))

	//函数类型作为返回值
	var fress funcType
	fress = add
	fmt.Println(fress(1, 3))

	fmt.Println("===调用callc 实现加法====")
	fmt.Println(callc(1, 2, add))

	fmt.Println("===调用callc实现减法===")
	fmt.Println(callc(20, 13, reduce))

	fmt.Println("===调用callc实现乘法===")
	fmt.Println(callc(2, 3, ride))
}

//go推荐写法
func my(name string) (res string) {
	res = "name: " + name
	if name == "daheige" {
		fmt.Println("haha")
	}

	//if可以支持初始化
	if a := 12; a > 10 {
		fmt.Println(a)
	}

	return
}

func myFun() (int, int, int) {
	return 1, 2, 3
}

//给函数返回值定义名字
func myFun2() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return
}

//1+2+3+4...累加
func sumInt(i int) (sum int) {
	if i == 1 {
		sum = 1
		return
	}

	sum = i + sumInt(i-1)
	return
}

func fibo(n int) (res uint64) {
	if fiboRes[n] != 0 {
		res = fiboRes[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibo(n-1) + fibo(n-2)
	}

	//将结果缓存起来
	fiboRes[n] = res
	return
}

func fiboArr(num int) []int {
	farr := make([]int, num) //预先分配切片长度
	farr[0], farr[1] = 1, 1
	for i := 2; i < num; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}

	return farr
}
