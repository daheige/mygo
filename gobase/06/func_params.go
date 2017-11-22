package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	callback(13, Add) //函数作为参数使用

	// 匿名函数作为闭包中的函数
	func(i int) {
		for j := 0; j < i; j++ {
			fmt.Println(j)
		}
	}(3)

	fmt.Println("f 调用后，ret", f()) //11

	hg_f := HgAdd(12) //返回值是一个函数
	fmt.Println(hg_f(123))

	//在闭包外面声明变量
	var g int
	func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
		fmt.Println(1, g)
	}(10000)

	//使用闭包做调试
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s : %d", file, line)
	}

	where()
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}

	where()

	//检测函数调用的时间
	start := time.Now()
	for j := 0; j < 10000; j++ {
		fmt.Println(j)
	}
	end := time.Now()
	fmt.Println(end) //当前时间
	const_time := end.Sub(start)

	fmt.Printf("cost time: %s\n", const_time)

}

func f() (ret int) {
	//在函数结束之前调用，具有延迟调用
	defer func() {
		ret++
	}()

	return 10 //return会先更新ret,再调用defer后面的函数
}

//将函数作为返回值
func HgAdd(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func Add(a, b int) {
	fmt.Println("a and b is", a, b)
}

func callback(y int, f func(int, int)) {
	f(y, 22)
}
