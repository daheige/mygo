package main

import (
	"fmt"
	"time"
)

//1.对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：
//如果通道中没有数据，接收者就阻塞了

//2.对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。
//所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）

//当没有接收者话，就不能给通道传入值
//接收操作在没有准备好之前，是阻塞的。
func main() {
	ch := make(chan string)

	//如果2个协程需要通信，你必须给他们同一个通道作为参数才行
	go setData(ch)
	go getData(ch)
	time.Sleep(1e9) //停顿1s
}

//通道赋值
func setData(ch chan string) {
	ch <- "heige"
	ch <- "zhuwei"
	ch <- "beijing"
}

//接收数据
func getData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Println("value is", input)
	}
}
