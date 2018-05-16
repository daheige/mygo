package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	//如果种子一样,产生的随机数字一样
	//rand.Seed(100)
	rand.Seed(time.Now().Unix()) //设置种子
	rnd := rand.Int() //产生一个大的数字

	fmt.Println(rnd)

	fmt.Println(rand.Intn(100)) //生成100内数字
	fmt.Println(makeRndNum(10,120))

}

//生成[n,m]范围的随机数字
func makeRndNum(min,max int64) int64{
	if min > max || max == 0{
		return max
	}

	rand.Seed(time.Now().UnixNano())
	return rand.Int63n((max - min)) +min
}
