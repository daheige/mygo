package main

import (
	"errors" //处理错误包
	"fmt"
)

//布尔值赋值
func test() {
	var avilable bool
	valid := false
	avilable = true
	fmt.Println("bool is", avilable, valid)
}
func main() {
	//同时声明多个int 类型变量
	var v1, v2, v3 int = 1, 2, 3
	fmt.Println("v1,v2,v3", v1, v2, v3)

	//简单的声明一个变量，让编译器自动判断类型为int
	b := 34
	fmt.Println("b value is", b)

	//定义常量
	const PI = 3.1415926
	const MAX_TH = 10

	fmt.Println("PI is", PI, "MAX_TH is", MAX_TH)

	// 内置基础类型
	//boolean
	var isActive bool = true
	var is_display = false //可以不显示指定变量类型
	fmt.Println(isActive, is_display)

	test()

	//浮点数　float32 float64,默认是float64
	price := 1.23
	fmt.Println("price is", price)

	//整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
	// Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。其中rune是int32的别称，byte是uint8的别称

	//字符串用双引号包起来,字符串是原始值，不能更改里面的字符

	name := "daheige"
	fmt.Println("name is ", name)

	char_name := []byte(name)
	fmt.Println(char_name) //[100 97 104 101 105 103 101]

	char_name[0] = 'D'
	name = string(char_name)
	fmt.Println(name)

	//通过切片方式更改字符串里的字符(字符串可以做切片操作)
	name = "hello," + name[1:]
	fmt.Println(name)

	// 错误类型
	err := errors.New("error!")
	fmt.Println(err)

	//一次声明多个变量
	const (
		LIFE  = 100
		MAX_J = 12.3
	)
	fmt.Println(LIFE, MAX_J)

	//声明枚举
	const (
		l_x = iota //默认从0开始，遇到下一个const iota重新为0
		l_y
		l_z
	)

	fmt.Println("l_x,l_y,l_z ", l_x, l_y, l_z)

	//声明数组 var arr [n]type 数组实质上一段连续的相同类型的元素组成
	var arr = [3]int{1, 2, 3} //声明并初始化
	fmt.Println("arr 数组的元素是", arr)

	var arr2 = [4]int{1, 3} //[1 3 0 0] 只初始化２个值，其他为０
	fmt.Println("arr2 ", arr2)

	var arr3 = [2][3]int{{1, 2}, {3, 4, 5}}
	fmt.Println("arr3二维数组 ", arr3)

	//切片声明　var slice []type
	var slice_1 = []string{"fefe", "sss"} //切片类型是string
	fmt.Println(slice_1)

	//从数组中创建切片arr2[m:n] //从m开始，到n-1位置的一段，切片三要素，开始位置m,结束位置n,容量cap
	// 开始元素指向底层数组的开始的那个元素
	slice_2 := arr2[0:2]
	fmt.Println(slice_2)

	slice_3 := []byte{'a', 'b', 'c'}
	fmt.Println(slice_3)

	fmt.Println([]byte("fef大黑哥"))         //转换为单个字节，是一个字节数组
	fmt.Println(string([]byte("fef大黑哥"))) //将字节数组转换为字符串

	//map　map[keyType]valueType key/val类型，类似js中的对象
	var obj = map[string]int{"x": 1, "y": 2}
	fmt.Println(obj["x"], obj["y"])
	//简短声明
	stu := map[int]string{0: "fefe", 1: "ssss"}
	fmt.Println(stu, stu[0], stu[1])
	//先声明后赋值
	obj_2 := make(map[string]string) //make主要用于map,slice,chan内存分配,new(T)可以任意类型
	obj_2["name"] = "daheige"
	obj_2["sex"] = "1"
	fmt.Println(obj_2)
	fmt.Printf("name is %s and sex is %s\n", obj_2["name"], obj_2["sex"]) //格式化打印%s表示字符串转换说明

	//判断key是否存在
	if _, ok := obj_2["name"]; ok { //通过value-ok方式判断map是否存在某个key
		fmt.Println("name is exist")
	} else {
		fmt.Println("name is not exist")
	}

	var slice_4 = make([]string, 3, 4) //一个字符串切片，长度为３，容量是4,切片只能通过索引来获取元素的值
	slice_4[0] = "fefee"
	slice_4[1] = "2222"
	slice_4[2] = "22sss"
	fmt.Println(slice_4)
	//append追加切片元素
	// slice_4[3] = "ssss333"
	// slice_4[4] = "4sss4444" //超过容量后，底层数组会指向一个新的数组(这里不能这样赋值，会报错，只能用append)
	slice_5 := append(slice_4, "4sss4444") //超过容量后，指向slice_5底层数组
	fmt.Println(slice_5)

	var l_num = new([3]int) //new(T)返回一个地址,一个*T类型的值
	l_num[0] = 12
	l_num[1] = 13
	l_num[2] = 14
	fmt.Println(l_num, l_num[0])

	//make返回初始化的非零值，new返回指针
	// 内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。
	// 本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化

}
