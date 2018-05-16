package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main()  {
	n := 10
	var arr [10]int
	for i:=0;i<n;i++{
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)

	for i :=0;i<n-1;i++{
		for j := 0;j <n-1-i;j++{
			if arr[j] > arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	fmt.Println(arr)
}
/**
运行结果:
[9 10 80 36 35 13 60 28 10 16]
[9 10 10 13 16 28 35 36 60 80]
 */
