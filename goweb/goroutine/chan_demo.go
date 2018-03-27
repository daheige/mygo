package main

import (
	"fmt"
)

// 使用同步通道来等待
func main() {
	ch := make(chan int) //创建无缓存通道
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum //将sum发送到ch
	}()

	// main主goroutine就不会终止，只有当计算和的goroutine完成后
	// 并且发送到ch通道的操作准备好后，同时<-ch就会接收计算好的值，然后打印出来
	fmt.Println(<-ch) //在太通道中没有接收到值之前，会一直阻塞等待

	one := make(chan int)
	two := make(chan int)
	go func() {
		one <- 100
	}()

	//将one通道中的值，作为two的输入
	go func() {
		two <- (<-one) //如果我们不将<-one发送给two,后面读取通道two中的值，会一直阻塞
	}()
	fmt.Println(<-two)

	//有缓冲通道，其实是一个队列，这个队列的最大容量就是我们使用make函数创建通道时
	//通过第二个参数指定的

	res := getRes()
	fmt.Println(res)

	// 定义单向通道也很简单，只需要在定义的时候，带上<-即可
	var send chan<- int //只能发送
	go func() {
		send <- 2
	}()
	// fmt.Println(<-send) //报错

}

func getRes() string {
	str_list := make(chan string, 3)
	go func() {
		str_list <- fmt.Sprintf("i am %d", 1)
	}()
	go func() {

		str_list <- fmt.Sprintf("i am %d", 2)
	}()
	go func() {

		str_list <- fmt.Sprintf("i am %d", 3)
	}()

	return <-str_list
}
