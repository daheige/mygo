//fibo内存缓存结果，通过内存缓存来提升性能
package main

import (
	"fmt"
	"time"
)

const LIM = 1000

var fiboRes [LIM]uint64 //存放fibo函数执行结果

func main() {
	var res uint64 = 0 //无符号的64位int
	start := time.Now()
	for i := 0; i < LIM; i++ {
		res = fibo(i)
		fmt.Printf("fibo(%d) is %d\n", i, res)
	}
	end := time.Now()
	cost_time := end.Sub(start)
	fmt.Println("cost time is", cost_time) //cost time is 0.195s
}

func fibo(n int) (res uint64) {
	if fiboRes[n] != 0 {
		res = fiboRes[n]
		return
	}

	if n <= 1 {
		res = 1
		return
	}

	res = fibo(n-1) + fibo(n-2)
	fiboRes[n] = res
	return
}

// func fibo(n int) uint64 {
//  if fiboRes[n] != 0 {
//      return fiboRes[n]
//  }

//  if n <= 1 {
//      return 1
//  }

//  fiboRes[n] = fibo(n-1) + fibo(n-2)
//  return fiboRes[n]
// }
