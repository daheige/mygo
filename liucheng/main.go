package main

import (
	"fmt"
)

func main() {
	x := 12
	if x < 12 {
		fmt.Println("x > 12")
	} else if x < 100 {
		fmt.Println("x >=12 and x < 100")
	} else {
		fmt.Println("x >=100")
	}
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("1-9的和是", sum)

	for i := 10; i > 0; i-- {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
	fmt.Println("fefe")

	flag := 1
	for flag < 5 {
		fmt.Println(flag)
		flag++
	}

	y := 123
	switch y {
	case 12:
		fmt.Println("y = 12")
	case 123:
		fmt.Println("y = 123")
	default:
		fmt.Println("haha")
	}

	arr := []int{1, 2, 4, 5}
	for item := range arr {
		fmt.Println(item)
	}
}
