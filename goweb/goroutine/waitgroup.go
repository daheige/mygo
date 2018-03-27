package main

import (
	"fmt"
	"sync"
)

//让程序交替打印A0-100,B0-100
//sync.WaitGroup其实是一个计数的信号量，使用它的目的是要main函数等待两个goroutine执行完成后再结束
//不然这两个goroutine还在运行的时候，程序就结束了，看不到想要的结果
func main() {
	var wg sync.WaitGroup
	wg.Add(2) //设置计数器为2
	//开启一个协程
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("A:", i)
		}
	}()

	//开启另一个goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait() //等待两个goroutine执行完毕
	fmt.Println("success")

}
