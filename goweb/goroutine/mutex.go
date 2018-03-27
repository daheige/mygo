package main

//sync提供了互斥锁，保证临界区的代码只能有一个goroutine访问
//sync包里提供了一种互斥型的锁，可以让我们自己灵活的控制哪些代码
//同时只能有一个goroutine访问，被sync互斥锁控制的这段代码范围，被称之为临界区
//临界区的代码，同一时间，只能又一个goroutine访问

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup //等待组
	mutex sync.Mutex     //互斥锁变量 互斥锁有两个方法，一个是mutex.Lock(),一个是mutex.Unlock()
	// 这两个之间的区域就是临界区，临界区的代码是安全的
)

func main() {
	wg.Add(2)
	go incNum()
	go incNum()
	wg.Wait()          //会一直等待incNum运行结束
	fmt.Println(count) //4
}

func incNum() {
	defer wg.Done() //退出之前wg计数器减去1
	for i := 0; i < 2; i++ {
		mutex.Lock()      //枷锁
		value := count    //count初始化值是0
		runtime.Gosched() //让当前goroutine暂停的意思，退回执行队列，让其他等待的goroutine运行
		value++
		count = value
		mutex.Unlock() //释放锁
	}
}
