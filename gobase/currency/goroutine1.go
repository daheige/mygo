package main

import (
	"fmt"
	"time"
)

//协程会随着main的结束而消亡
func main() {
	fmt.Println("in main")

	//创建一个协程
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main")
	time.Sleep(10 * 1e9)
	// time.Sleep(3 * 1e9)
	fmt.Println("at end of main")
}

func longWait() {
	fmt.Println("begin longwait")
	time.Sleep(5 * 1e9)
	fmt.Println("end of longwait")
}

func shortWait() {
	fmt.Println("begin shortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("end of shortwait")
}
