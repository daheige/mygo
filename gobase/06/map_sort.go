package main

import (
	"fmt"
	"sort"
)

func main() {
	var map_s = map[string]int{
		"x": 1,
		"y": 3,
		"a": 12,
		"s": 23,
		"b": 233,
	}

	fmt.Println("unsorted")
	for k, val := range map_s {
		fmt.Println("key is", k, "val is", val)
	}

	//获取map_f所有的key
	keys := make([]string, len(map_s))
	i := 0
	for k, _ := range map_s {
		keys[i] = k
		i++
	}

	sort.Strings(keys) //对key进行排序
	fmt.Println("sorted")
	for _, k := range keys {
		fmt.Println("key is", k, "val is", map_s[k])
	}

	// key is a val is 12
	// key is b val is 233
	// key is s val is 23
	// key is x val is 1
	// key is y val is 3
}
