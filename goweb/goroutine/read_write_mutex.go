package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//读写锁可以让多个读操作同时并发，同时读取，但是对于写操作是完全互斥的。也就是说，当一个goroutine进行写操作的时候
// 其他goroutine既不能进行读操作，也不能进行写操作。

var (
	count int
	wg    sync.WaitGroup
	rw    sync.RWMutex
)

func main() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		go read(i)
	}
	for i := 0; i < 5; i++ {
		go write(i)
	}
	wg.Wait()
}

func read(n int) {
	rw.RLock()
	fmt.Printf("读goroutine %d 正在读取。。。\n", n)

	v := count //对共享变量进行读取操作
	fmt.Printf("读goroutine %d 读取结束，值为： %d \n", n, v)

	wg.Done()
	rw.RUnlock()
}

func write(n int) {
	rw.Lock()
	fmt.Printf("写goroutine %d 正在写入。。。\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("写goroutine %d 写入结束，新值是%d \n", n, v)

	wg.Done()   //计数器减去1
	rw.Unlock() //释放锁
}

/**运行结果;
读goroutine 0 正在读取。。。
读goroutine 1 正在读取。。。
读goroutine 0 读取结束，值为： 0
读goroutine 1 读取结束，值为： 0
写goroutine 4 正在写入。。。
写goroutine 4 写入结束，新值是81
读goroutine 2 正在读取。。。
读goroutine 2 读取结束，值为： 81
读goroutine 3 正在读取。。。
读goroutine 3 读取结束，值为： 81
读goroutine 4 正在读取。。。
读goroutine 4 读取结束，值为： 81
写goroutine 0 正在写入。。。
写goroutine 0 写入结束，新值是887
写goroutine 1 正在写入。。。
写goroutine 1 写入结束，新值是847
写goroutine 2 正在写入。。。
写goroutine 2 写入结束，新值是59
写goroutine 3 正在写入。。。
写goroutine 3 写入结束，新值是81
*/
