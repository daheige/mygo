package main

import (
	"time"
	"fmt"
)

func main()  {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3 *time.Second) // sleep 3 seconds
		timeout <- true
	}()

	//超时处理
	ch := make(chan int)
	go func(){
		ch <-1
	}()

	//这里一定不能用default，否则3s超时还未到直接执行default，case2便不会执行
	// 超时机制便不会实现。timeout会在3s超时后读取到数据
	select {
	case v := <-ch:
		fmt.Println(v)
	case <-timeout:
		fmt.Println("timeout!")
	}

	//采用select判断是否存满
	ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)

	ch1 <- 1
	//ch2 <- 2

	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	//case <-ch2:
	//	fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}

	test()
}

func add(ch chan int){
	for i := 0;i<10;i++{
		ch <- i //发送i到ch
	}
}
/**
0
sleep 1s
1s exec end
1
sleep 1s
1s exec end
2
sleep 1s
1s exec end
timeout
 */
// 当case <- ch 和 case <- ticker.C 同时成立时
// Select会随机公平地选出一个执行，有可能选择到前者，导致超时随机行失败
func test()  {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	ch := make(chan int,10)
	go add(ch) //异步往ch填充数据

	for{
		select{
		case v := <- ch:
			fmt.Println(v)
			fmt.Println("sleep 1s")
			//time.Sleep(1 * time.Second) //如果不停顿,当ch满了后,就会依次读取,当无法获取到ch后,会走到case2
			fmt.Println("1s exec end")
		case <- ticker.C:
			fmt.Println("timeout")
			return
		}
	}

}