package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

// go build -race会打印资源竞争的情况
func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count) //如果不枷锁，可能为2，4存在资源竞争，对同一个变量进行读写
}
func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		// mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		// mutex.Unlock()
	}
}
