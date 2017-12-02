package main

import (
	"fmt"
	"time"
)

func mygo() {
	fmt.Println("gogo")
}
func main() {
	go mygo()
	time.Sleep(2 * time.Second) //不建议使用暂停
}
