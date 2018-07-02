package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

/**
// set 10000 the size of goroutine pool
p, _ := ants.NewPool(10000)
// submit a task
p.Submit(func() {})
pool.ReSize(1000) // Readjust its capacity to 1000
*/

var sum int32

func main() {
	defer ants.Release() //释放资源

	runtimes := 100

	//使用通用协成
	var wg sync.WaitGroup
	for i := 0; i < runtimes; i++ {
		wg.Add(1) //计数器用来保证goroutine执行完毕
		fmt.Println("start at ", i)
		//异步执行,submit提交任务task
		ants.Submit(func() error {
			testFunc()
			wg.Done()
			return nil
		})
	}
	wg.Wait() //等待任务执行完毕

	//通过协成函数的方式运行 use pool with a func
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) error {
		myFunc(i) //依次叠加
		wg.Done()
		return nil
	})

	defer p.Release()

	for i := 0; i < runtimes; i++ {
		wg.Add(1)
		p.Serve(int32(i)) // Serve submit a task to pool 将提交的任务放入pool任务池
	}
	wg.Wait()
	fmt.Println("running goroutines:", p.Running()) //正在运行的协程数
	fmt.Printf("finish all tasks,res is %d", sum)
}

func myFunc(i interface{}) error {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
	return nil
}

func testFunc() error {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("hello")
	return nil
}
