package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) //给通道分配内存
	go pump(ch)
	go suck(ch)
	time.Sleep(1e9)
}

//通道提供数据，我们一般会称为生产者
//接收数据，称为接收者，消费者
//无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i //把i传入通道ch
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
