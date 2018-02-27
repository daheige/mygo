package main

import (
	"fmt"
	"time"
)

func main() {
	//创建两个通道chan，没有第二个参数的通道是一个无缓存的通道，实时同步
	w1, w2 := make(chan bool), make(chan bool)
	go showNum1(w1)
	go showWord(w2)
	//在w1,w2未放入true之前，一直是阻塞状态，当goroutine执行完毕后，true才放入
	//当取出w1,w2后，main的阻塞才得到释放
	a := <-w1
	b := <-w2
	fmt.Println(a, b)
}

func showNum1(ch chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d ", i)
	}

	ch <- true //将true放入通道ch中
}

func showWord(ch chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(time.Millisecond)
		fmt.Printf("%c ", i)
	}
	//结束for后将true放入ch
	ch <- true
}
