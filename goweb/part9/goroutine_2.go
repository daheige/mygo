package main

import (
	"fmt"
	"time"
)

func main() {
	goPrint1()                   //当我们执行运行的时候，发现什么也没有输出来。因为在执行它之前，已经结束了。
	time.Sleep(time.Millisecond) //停顿1ms后可以输出
}

func printNum1() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Microsecond)
		fmt.Println(i)
	}
}

func printWord1() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(time.Microsecond) //停顿1us
		fmt.Printf("%c", i)
	}
}

func goPrint1() {
	go printNum1() //在函数前面加go创建一个协程
	go printWord1()
}
