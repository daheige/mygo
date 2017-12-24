package main

//log的一些用法  Golang's log模块主要提供了3类接口。
//分别是 “Print 、Panic 、Fatal ”，对每一类接口其提供了3中调用方式，分别是 "Xxxx 、 Xxxxln 、Xxxxf"，基本和fmt中的相关函数类似
// 对于 log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1 。
// 但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用
import (
	"log"
)

func main() {
	arr := []int{1, 2, 3}
	log.Println(arr) //2017/12/24 20:57:17 [1 2 3]
	log.Printf("arr is %v", arr)
	log.Printf("item1 is %d", arr[0]) //2017/12/24 20:59:05 item1 is 1
	test()
}

func test() {
	//defer不会执行
	defer func() {
		log.Print("1111")
	}()
	log.Println("i am exit")
	log.Fatal("exit 1")
}
