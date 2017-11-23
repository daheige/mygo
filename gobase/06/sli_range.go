package main

import (
	"fmt"
)

func main() {
	var sli = make([]int, 4)
	sli[0] = 12
	sli[1] = 23
	sli[2] = 3
	sli[3] = 56

	//遍历切片
	for k, val := range sli {
		fmt.Println(k, val)
	}

	var arr = [...]int{1, 3, 4}
	//切片里的val是对切片元素值的拷贝
	for _, item := range arr {
		item *= 2
	}

	fmt.Println(arr)

	var af = [...]float32{1, 3, 4.2, 3.2}
	fmt.Println(Fsum(af[:]))

	var hg_af = [...]float32{1, 2.1, 3}
	fmt.Println(hgSum(&hg_af)) //将数组的指针作为参数传递给函数

	fmt.Println(hg_af[3:3], len(hg_af[3:3])) //hg_af[n:n]长度为0

	//改变切片元素的值，底层数组也发生了改变
	hg_sli := hg_af[:]
	hg_sli[0] = 211
	fmt.Println(hg_af, hg_sli)
	hg_sli = hg_sli[0:2] //切片重组，如果没有超过容量，底层数组不发生改变
	fmt.Println(hg_af, hg_sli)
}

func Fsum(a []float32) float32 {
	var res float32
	for _, v := range a {
		res += v
	}

	return res
}

//将数组的指针作为参数传递
func hgSum(a *[3]float32) (res float32) {
	for _, v := range a {
		res += v
	}

	return
}
