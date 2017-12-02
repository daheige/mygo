package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool) //无缓冲通道
	go func() {
		fmt.Println("go go")
		ch <- true //执行完毕通知
		close(ch)  //关闭通道
	}()

	//等待通道中的值返回，这里会阻塞
	<-ch
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("执行完毕")

}
