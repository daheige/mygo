package main

import (
	"fmt"
)

func main() {
	demo(1, 3, 4)
	//函数作为值使用
	tmp := func() {}
	fmt.Printf("%T\n", tmp) //func()
	callback(1, demo)
}

func demo(argv ...int) {
	fmt.Printf("%T\n", argv) //[]int切片类型
	fmt.Println(argv[0], argv)
}

func callback(x int, f func(...int)) {
	fmt.Println(x)
	f(x, 1, 2, 2) //f是一个可变参数的函数

}
