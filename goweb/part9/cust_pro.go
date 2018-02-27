package main

//利用chan实现消息传递
//生产者，将消息放入通道中
//消费者，将消息从通道中取出

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go cust(ch) //必须先产生ch,才可以被消费取出ch
	go production(ch)
	time.Sleep(1 * time.Millisecond)
}

func cust(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i //将i放入通道ch中
		fmt.Println("send data >> ", i)
	}
}

func production(ch chan int) {
	for i := 0; i < 5; i++ {
		num := <-ch //取出是<-ch
		fmt.Println("caught << ", num)
	}
}
