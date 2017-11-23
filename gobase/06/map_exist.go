package main

import (
	"fmt"
)

func main() {
	mapList := map[string]int{
		"x": 12,
		"y": 13,
		"z": 2322,
	}

	//判断key是否存在map中 val,ok语法，其中val可以忽略
	if val, ok := mapList["y"]; ok {
		fmt.Println("mapList中存在y", "对应的值是", val)
	}

	//删除某个key
	delete(mapList, "y")
	fmt.Println(mapList)

	if _, ok := mapList["a"]; ok {
		fmt.Println("存在a")
	} else {
		fmt.Println("不存在a")
	}
}
