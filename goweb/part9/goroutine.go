package main

import (
	"fmt"
	// "time"
)

// func main() {
//  printNum()
//  printWord()
//  fmt.Println("==============")
//  goPrint()                    //当我们执行运行的时候，发现什么也没有输出来。因为在执行它之前，已经结束了。
//  time.Sleep(time.Millisecond) //停顿1ms后可以输出
// }

func printNum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printWord() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c", i)
	}
}

func goPrint() {
	go printNum() //在函数前面加go创建一个协程
	go printWord()
}
