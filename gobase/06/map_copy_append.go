package main

import (
	"fmt"
)

func main() {
	sl_src := []int{1, 2, 3}
	sl_dist := make([]int, 10)
	//复制一个切片到新的切片中，需要创建的切片比原有的切片大
	n := copy(sl_dist, sl_src) //返回复制的个数
	fmt.Println(n)
	fmt.Println(sl_dist) //[1 2 3 0 0 0 0 0 0 0]

	sl2 := []int{1, 3, 4}
	sl2 = append(sl2, 2, 3, 4, 5) //append追加的类型需要和切片的类型一样
	fmt.Println(sl2)

	var sl3 = []byte("大黑哥")
	sl3 = append(sl3, ([]byte("daheige"))...)
	fmt.Println(sl3)
	fmt.Println(string(sl3))

}
