package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go setChData(ch) //只有一个协程
	getChData(ch)    //getChData和main构成同一个线程
}

func setChData(ch chan string) {
	ch <- "heige"
	ch <- "php"
	ch <- "go"
	close(ch)
}

func getChData(ch chan string) {
	// for {
	//  input, open := <-ch //通过v-ok模式判断
	//  if !open {
	//      break
	//  }
	//  fmt.Println(input)
	// }

	//for...range会自动判断通道是否已关闭
	for v := range ch {
		fmt.Printf("%s\n", v)
	}
}
