package main

import (
	"fmt"
	"time"
)

func main() {
	// stream := sumpFactory()
	// go reviceSump(stream)

	// stream_2 := setMoreCh(2)
	// getMoreCh(stream_2)

	//两个协程之间通信
	getMoreCh(setMoreCh(3))
	time.Sleep(1e9) //停顿1s让通道的收发在1s之内完成
}

//通过工厂模式生成一个通道
func sumpFactory() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func reviceSump(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

//用函数来生成一个通道并返回（工厂角色）
func setMoreCh(buf int) chan int {
	ch := make(chan int, buf)
	go func() {
		for i := 0; i < buf; i++ {
			ch <- i
		}
	}()

	return ch
}

func getMoreCh(ch chan int) {
	// fmt.Println(ch) //是一个地址
	// for {
	//  fmt.Println(<-ch)
	// }

	//利用for...range来遍历通道中的值
	//for 循环的 range 语句可以用在通道 ch 上，便可以从通道中获取值
	/*for v := range ch {
	      fmt.Println(v)
	  }
	*/

	//for它从指定通道中读取数据直到通道关闭，才继续执行下边的代码
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
