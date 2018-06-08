package main

import (
	"fmt"
)

func main() {
	fmt.Println("daheige33")
	s1 := make([]int, 3)

	fmt.Println("切片的容量是", cap(s1))
	s1 = append(s1, 1, 3, 4, 5)
	fmt.Println(s1)

	s2 := s1[:] //当我们改变s2的时候会将底层的s1也发生了改变
	s2[0] = 12
	fmt.Println(s1, s2)
	fmt.Println(s1[1:3]) //从第1个位置开始,到第3个(不包含3index位置的元素)

	var a []int
	//	a[1] = 2
	fmt.Println(a)         //panic: runtime error: index out of range
	a = append(a, 1, 2, 3) //在原来的切片末尾添加元素
	fmt.Println(a)

	b := []int{}
	//	b[1] = 1 //panic: runtime error: index out of range
	//	fmt.Println(b)
	b = append(b, 12, 13)
	fmt.Println(b)

	c := make([]int, 1, 3)
	c[0] = 1
	fmt.Println(c)
	//	c[1] = 2 //越界了会引发error
	c = append(c, 2, 3)

	fmt.Println("当前c的容量是", cap(c)) //3
	c = append(c, 1, 2, 3)
	fmt.Println("扩容后的容量", cap(c)) //6

}
