package main

import (
	"fmt"
)

func main() {
	m := map[int]string{
		1: "heige",
		2: "3333",
	}
	fmt.Println("m is", m)

	//for...range
	for k, v := range m {
		fmt.Printf("%d====%s\n", k, v)
	}

	//判断k是否需存在map中
	if v, ok := m[1]; ok {
		fmt.Println(v)
	}

}
