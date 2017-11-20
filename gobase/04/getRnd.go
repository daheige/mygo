package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := rand.Int()
	fmt.Println(a)

	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d /", a)
	}

	r := rand.Intn(9) //[0-9)随机数生成
	fmt.Println(r)

	times := int64(time.Now().Nanosecond()) //当前时间的纳秒级数字

	rand.Seed(times) //生成随机种子

	fmt.Println(times)

	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f /", 100*rand.Float32())
	}
}
