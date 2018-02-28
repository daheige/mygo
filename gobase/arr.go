package main

//数组声明和遍历
//数组传参
import (
	"fmt"
)

func main() {
	//先申明后赋值 中括号后面的是类型type
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	//指定长度并初始化赋值
	var arr_b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr_b)

	//不指定长度
	var arr_c = [...]int{1, 3, 4, 2}
	fmt.Println(arr_c)

	//函数里申明数组
	arr_d := [3]string{"a", "b", "c"}
	fmt.Println(arr_d)

	//打印数组value
	for i := 0; i < len(arr_b); i++ {
		fmt.Println(arr_b[i])
	}

	//for...range遍历
	for k, v := range arr_b {
		fmt.Printf("第%d个值是%v\n", k, v)
	}

	//指针数组和数组本身差不多，只不过元素类型是指针
	point_arr := [5]*int{1: new(int), 2: new(int)}
	fmt.Println(point_arr)
	//索引1和2都创建了内存空间，其他索引是指针的零值nil
	//这时候我们要修改指针变量的值也很简单
	*point_arr[1] = 12
	fmt.Println(point_arr, point_arr[1], *point_arr[1])

	//将数组的指针传递给函数
	hg_arr := [5]int{1, 2, 3, 5, 6}
	modify(&hg_arr)
	fmt.Println(*(&hg_arr), hg_arr[1])
	fmt.Println(hg_arr)
}

//参数是数组的指针 *[5]int
func modify(arr *[5]int) {
	arr[1] = 3
	fmt.Println(arr)
}

/**
运行结果：
[1 2 3 0 0]
[1 2 3 4 5]
[1 3 4 2]
[a b c]
1
2
3
4
5
第0个值是1
第1个值是2
第2个值是3
第3个值是4
第4个值是5
[<nil> 0xc04204c118 0xc04204c120 <nil> <nil>]
[<nil> 0xc04204c118 0xc04204c120 <nil> <nil>] 0xc04204c118 12
&[1 3 3 5 6]
[1 3 3 5 6] 3
[1 3 3 5 6]
*/

