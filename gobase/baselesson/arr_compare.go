package main

import "fmt"

func main() {
	//类型和值要一样才可以
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = b", a == b)

	c := [5]int{1, 2}
	fmt.Println(a == c)
}

