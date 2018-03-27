package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 对于并发来说，就是Go语言本身自己实现的调度，对于并行来说，是和运行的电脑的物理处理器的核数有关的
// 多核就可以并行并发，单核只能并发了
func main() {
	fmt.Println(runtime.NumCPU()) //打印cpu核数
	//强制使用一个逻辑处理器
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("A:", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}
