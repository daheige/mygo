package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//等待组WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(1 * time.Millisecond)
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("hello")
}
