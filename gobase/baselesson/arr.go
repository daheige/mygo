package main

import (
	"fmt"
)

func main()  {
	var a [5]int = [5]int{1,2,3}
	fmt.Println(a)

	b := [5]int{1,2,3}
	fmt.Println(b) //[1 2 3 0 0]

}
