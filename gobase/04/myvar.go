package main

import (
	"fmt"
	"os"
	"runtime"
)

//一次声明多个变量
var (
	a int    = 1
	b bool   = true
	c string = "heige"
)

//编译器自行推断变量类型
var d = 12
var e = false
var f = "i am heige"
var n int32 = 23

//获取系统级别的常量
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOPATH = os.Getenv("GOPATH")
	UUID   = os.Getuid()
)

func test() {
	a := 1 //在函数里面的局部变量用简写方式:=声明
	fmt.Println(a)
}

func main() {
	fmt.Println("a,b,c", a, b, c)
	fmt.Println(d, e, f, n)
	fmt.Println(HOME, USER, GOPATH, UUID)
	test()

	//获取运行时候的环境变量
	var goos string = runtime.GOOS
	fmt.Printf("The system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("The system path is %s\n", path)
}
