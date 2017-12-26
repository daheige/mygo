package main

import (
	"fmt"
)

func main() {
	//值是一个字符串切片的map
	dict := map[int][]string{}
	dict[0] = []string{"ab", "cd", "ef"}
	fmt.Println(dict)
	dict[0] = []string{"heige", "php", "go"}
	fmt.Println(dict)

	colors := map[string]string{}
	colors["x"] = "123"
	colors["y"] = "heige"
	colors["z"] = "golang"
	fmt.Println(colors)

	//遍历map
	for key, value := range colors {
		fmt.Println(key + " is " + value)
	}

	//采用make方式声明map
	dist_b := make(map[string]int, 3)
	dist_b["x"] = 3
	dist_b["y"] = 4
	dist_b["z"] = 5
	dist_b["a"] = 123
	dist_b["b"] = 123
	fmt.Println(dist_b)
	fmt.Println("长度为", len(dist_b))
}
