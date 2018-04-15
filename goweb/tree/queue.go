package main

import (
	"fmt"
	"goweb/queue"
)

func main() {
	q := queue.Queue{1, 3}

	q.Push([]int{12, 34}...)
	q.Push(1, 3, 4, 5)
	fmt.Println(q)

	v := q.Pop()
	fmt.Println(v)
	fmt.Println(q)

	q.Shift()
	fmt.Println(q)

	fmt.Println("----开始清空----")
	q.Clear()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	//清空
	q = []int{}
	fmt.Println(q)
	fmt.Println(q.IsEmpty())

	q.Push(1, 3)
	fmt.Println(q)

}
