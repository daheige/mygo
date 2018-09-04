//等待多个task执行完毕
//指定多个goroutine任务
//同步实现: wg或chan机制
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		//goroutine和defer一样会延迟执行,这里就直接把i传入了func
		go func(n int) {
			runtime.Gosched() //让出cpu给其他goroutine执行
			defer wg.Done()
			// time.Sleep(1 * time.Second)
			fmt.Println("goroutine", n, "has done.")
			runtime.GC() //强制进行垃圾回收,一般在收发chan会发生携程资源泄露
		}(i)
	}

	fmt.Println("main is running...")
	wg.Wait() //阻塞,直到计数归零

	//采用chan实现同步执行
	//chan用于解决逻辑层次的并发操作
	ch := make(chan int)
	go func() {
		fmt.Println("this goroutine has done")
		ch <- 1
	}()

	<-ch

	fmt.Println("main will exit")
}

/*
goroutine 884242 has done.
main will exit
[Finished in 4.5s]
*/
