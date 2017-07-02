package main

import (
	"errors"
	"fmt"
	"mymath"
)

//定义全局变量
var username string = "heige"

func main() {
	fmt.Printf("hello,world.Sqrt(2) = %v\n", mymath.Sqrt(2))
	fmt.Println("daheige")
	fmt.Println(username)
	age := 25
	fmt.Println(age)
	fmt.Println(age)
	fmt.Println("fefefe")

	_, b := 23, 35
	fmt.Println(b)

	var i int = 12
	fmt.Println(i)

	var hg_a, hg_b, hg_c int = 1, 2, 3
	fmt.Println(hg_a, hg_b, hg_c) //一次声明多个变量

	//定义常量
	const (
		Pi        float64 = 3.1415
		hg_i      int16   = 120
		MaxThread         = 10
		prefix            = "hg_"
	)
	fmt.Println(Pi, hg_i, MaxThread, prefix)

	avail := true //简短声明，:=只能在函数里面声明变量
	fmt.Println(avail)
	//定义整形变量和浮点型变量（float32,float64)
	var ax int8 = 12
	var ay int = 13
	var az float32 = 123.09
	fmt.Println(ax, ay, az)

	var french string                        //定义字符串，但未赋值
	var emptyStr string = "fefefe,hei heige" //定义并赋值
	french = "hello"
	no, yes, maybe := "no", "yes", "mybe" //简短声明（只能在函数中声明）
	fmt.Println(french, emptyStr, no, yes, maybe)

	//修改字符串，在go中字符串不可变，可以用slice方式来修改
	var hg_s string = "hello"
	c := []byte(hg_s) //单字符切片slice
	c[0] = 'a'
	s2 := string(c)                  //转回字符串类型
	fmt.Printf("%s\n", s2)           //格式化打印
	fmt.Printf("%s\n", "heige"+s2)   //字符串拼接
	fmt.Printf("%s\n", "c"+hg_s[1:]) //字符串可以做切片操作

	//	多行字符操作,用反引号`
	m := `fefefefefefefefe
		fefefefe
			fefefe`
	fmt.Println(m)

	//错误类型处理
	err := errors.New("emit macho err")
	if err != nil {
		fmt.Print(err) //打印错误到终端
	}

	//	iota使用 ,枚举实现
	const (
		hg_x = iota //0
		hg_y = iota //1
		hg_z = iota //2
		hg_w        //默认和之前一个值的字面相同,这里隐式地说hg_w = iota，因此hg_w == 3
	)
	const hg_v = iota // 0,每遇到一个const时候iota就会重置
	fmt.Println(hg_x, hg_y, hg_z, hg_w, hg_v)

	//iota在同一行值相同
	const (
		o, p, q = iota, iota, iota
		s       = iota //1
	)
	fmt.Println(o, p, q, s)

	//array定义
	var arr [10]int //声明了一个int类型的数组
	arr[0] = 1
	arr[1] = 2

	fmt.Println(arr[0], arr[2], arr[3])

	hg_arr := [3]int{1, 2} //长度是3,初始化了前两个元素
	fmt.Println(hg_arr)    //[1,2,0]
	//采用`...`的方式，Go会自动根据元素个数来计算长度
	hg_arr_2 := [...]int{1, 2, 3} //[1,2,3]
	fmt.Println(hg_arr_2)

	//二维数组声明
	hg_arr_3 := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(hg_arr_3[0][1], hg_arr_3) //2 [[1 2 3 4] [5 6 7 8]]

	//slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度
	//	而声明slice时，方括号内没有任何字符
	//切片slice定义，不是真正的数组，而是一个引用类型，底层指向一个数组，声明不需要指定长度
	var fslice []int
	//	fslice[0] = 1 //非法操作
	//	fmt.Println(fslice)
	fslice = hg_arr[1:] //底层指向数组hg_arr的第1个位置到末尾的一段
	fmt.Println(fslice) //[2,0]
	hg_arr_4 := [6]byte{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println(hg_arr_4) //[97 98 99 100 101 102]

	hg_a_slice := hg_arr_4[2:5] //从第三个元素开始到第五个元素结束
	hg_b_slice := hg_arr_4[3:5] //不包含结束位置5

	fmt.Println(hg_a_slice, hg_b_slice, len(hg_arr_4)) //[99 100 101] [100 101]
	//slice从概念上讲包含了三个元素，一个指针，指向数组中指定的开始位置
	//长度，slice的长度
	//最大长度，就是slice开始位置到结束位置的长度

	fmt.Println("hg_a_slice最大容量是：", cap(hg_a_slice)) //从hg_a_slice开始位置到hg_arr_4的结束位置的长度
	//	追加元素到slice
	hg_a_slice = append(hg_a_slice, 'g')
	fmt.Println(hg_a_slice, hg_arr_4) //结果hg_arr_4也发生了改变，因为slice是引用类型，指向hg_arr_4

	//ｍap字典　格式为map[keyType]valueType　键的类型，值的类型
	//长度是不固定的，也就是和slice一样，也是一种引用类型
	numbers := make(map[string]string)
	numbers["user"] = "heige"
	numbers["age"] = "26"
	numbers["job"] = "php"
	fmt.Println(numbers) //map[user:heige age:26 job:php] 字典是无顺序的，每次打印出来都不同
	//获取字典的值
	fmt.Println(numbers["user"], numbers["age"])
	fmt.Println("key的长度是", len(numbers))
	//判断key是否存在
	//	hg_user, ok := numbers["user"]
	if hg_user, ok := numbers["user"]; ok { //if语句中可以定义语句和变量
		fmt.Println("username", hg_user)
	} else {
		fmt.Println("user is not exist")
	}

	hg_m := make(map[string]string)
	hg_m["he"] = "1111"
	hg_m["x"] = "fefefe"
	hg_m1 := hg_m
	hg_m1["y"] = "fefessss"
	fmt.Println(hg_m, hg_m1) //map[he:1111 x:fefefe y:fefessss] map[he:1111 x:fefefe y:fefessss]

	//make or new 操作
	//	make用于内建类型（map、slice 和channel）的内存分配

}
