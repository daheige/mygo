//在 select 中使用发送操作并且有 default可以确保发送不被阻塞！如果没有 case，select 就会一直阻塞。
// select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出
// select 做的就是：选择处理列出的多个通信情况中的一个。

/*如果都阻塞了，会等待直到其中一个可以处理
如果多个可以处理，随机选择一个
如果没有通道操作可以处理并且写了 default 语句，它就会执行：
default 永远是可运行的（这就是准备好了，可以执行）。
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1, 3)
	go pump2(ch2, 3)
	go suck2(ch1, ch2)

	time.Sleep(1e9)

}

func pump1(ch chan int, num int) {
	for i := 0; i < num; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int, num int) {
	for i := 0; i < num; i++ {
		ch <- i + 6
	}
}

//从协程中取出一个通道,从多个通信中选择一个
func suck2(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1 value is", v)
		case v := <-ch2:
			fmt.Println("ch2 value is", v)
		}
	}
}
