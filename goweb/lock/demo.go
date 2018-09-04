//lock用来保护局部范围内的数据安全,比如说全局变量,对单个数据读写可以用rW LOCK
package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex //不能递归枷锁或反复枷锁,将mutex作为匿名字段时候,方法接收者必须是一个指针
}

func (d *data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 10; i++ {
		fmt.Println(s, i)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	var d data
	var wg sync.WaitGroup
	wg.Add(2) //通过wg保证下面的goroutine执行完毕
	go func() {
		defer wg.Done()
		d.test("daheige")
	}()

	go func() {
		defer wg.Done()
		d.test("xiaoming")
	}()

	wg.Wait()
	fmt.Println("has done")
}
