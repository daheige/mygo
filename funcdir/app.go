package main

import (
	"fmt"
)

//函数作为返回值
//defer在函数退出之前执行，调用遵循先进后出
func test(x int) func() {
	x = x + 1
	return func() {
		defer fmt.Println("x已经发生改变") //退出函数之前调用
		fmt.Println("改变之后的x", x)
		defer fmt.Println("我是后面执行的defer")
	}
}
func main() {
	fmt.Println(1)
	x := 100
	f := test(x)
	f()

	slice_1 := []int{1, 2, 3} //声明切片，未指定容量
	slice_1 = append(slice_1, 10)
	fmt.Println(slice_1)

	//创建容量为5的切片
	slice_2 := make([]int, 0, 5) //创建一个切片，长度初始化为0，容量为5
	for i := 0; i < 5; i++ {
		slice_2 = append(slice_2, i+1)
	}
	fmt.Println(slice_2)

	// obj := map[string]string{} //声明一个空map
	obj := make(map[string]interface{})
	obj["name"] = "heige"
	obj["sex"] = 1
	obj["desc"] = "fefe"
	fmt.Println(obj)

	_, ok := obj["sex"]
	if ok {
		fmt.Println("obj存在sex键")
	}

}
