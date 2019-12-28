package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// log使用方式
	log.Println("fefe")
	log.SetPrefix("[info]")
	log.Println(1111)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println(12345)

	// 通过通道实现同步，这里的ch是一个非阻塞的缓冲通道
	ch := make(chan struct{}, 1)
	go func() {
		log.Println(123)
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()

	<-ch

	// goroutine同步计数器
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done() // goroutine执行完毕后，计数器减1

		for i := 0; i < 10; i++ {
			log.Println("hello,index: ", i)
		}
	}()

	// 当计数为0，这里就解除了阻塞
	// Counter is 0, no need to wait.
	wg.Wait()

	log.Println("main will exit")
}
