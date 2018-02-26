package main

import (
	"fmt"
	"sync"
	"time"
)

//输出0-9
func printnum3(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done() //等待组计数器减少1
}

//输出A-J
func printLetters(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	wg.Done() //等待组计数器减少1 如果我们忘记将计数器减去1，就会引发错误error: all goroutines are asleep
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printLetters(&wg)
	go printnum3(&wg)
	wg.Wait() //一直等待计数器变成0（直到每个goroutine运行完毕）
}
