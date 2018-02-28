package main
//数组的声明和遍历
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
}
