package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3} //未初始化的默认为0
	fmt.Println(arr)
	arr[2] = 23
	fmt.Println(arr)

	//声明一个指针数组
	//array := [5]*int{0: new(int), 1: new(int)}
	//*array[0] = 1
	//*array[2] = 2
	//fmt.Println(array)

	str_arr := [3]string{"abc", "fefe"}
	str_arr[2] = "heige" //如果下标超过3就会报错
	fmt.Println(str_arr)

	muli_arr := [3][2]int{{1, 2}, {3, 4}} //第三组默认为[0,0]
	fmt.Println(muli_arr)

	//切片的声明
	slice_a := make([]int, 2) //长度为2的整型切片，内存存放是连续的
	//因为切片的底层内存也是在连续块中分配的,所以切片还能获得索引、迭代以及为垃圾回收优化的好处
	slice_a[0] = 12
	slice_a[1] = 10
	slice_a = append(slice_a, 123)
	fmt.Println(slice_a)
	//字面量声明
	slice_b := []string{"ab", "cd"}
	slice_b = append(slice_b, "fefe", "heige") //自动扩容
	fmt.Println(slice_b)

	for k, v := range slice_b { //k是索引，v表示切片元素的值
		fmt.Println(k, v)
	}

}
