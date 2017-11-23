package main

import (
	"fmt"
)

func main() {
	arr := [...]float64{3, 4, 5.6, 5}
	fmt.Println(Sum(arr[:]))
	fmt.Println(HgSum(&arr)) //传递一个数组的指针
}

//把一个大数组传递给函数会消耗很多内存。
//有两种方法可以避免这种现象：1.传递数组的指针 2. 使用数组的切片
//把数组的切片传递给函数
func Sum(a []float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

//传递数组的指针作为参数
func HgSum(a *[4]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}
