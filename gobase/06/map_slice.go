//map类型的切片
//先分配切片，再分配val为切片
//使用两次 make() 函数，第一次分配切片，第二次分配 切片中每个 map 元素
package main

import (
	"fmt"
)

func main() {
	var mapSli = make([]map[string]int, 5)
	mapSli[0] = map[string]int{
		"x": 12,
		"y": 23,
	}
	/*
	   mapSli[0] = make(map[string]int,2)
	   mapSli[0]["x"] = 1
	*/

	//利用make第二次分配切片中的每个元素
	mapSli[1] = make(map[string]int)
	mapSli[1]["x"] = 89
	mapSli[1]["y"] = 33

	fmt.Println(mapSli)
	//[map[x:12 y:23] map[x:89 y:33] map[] map[] map[]]
}
