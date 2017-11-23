package main

import (
	"fmt"
)

func main() {
	mapList := map[string]string{
		"name": "heige",
		"job":  "php js go",
		"age":  "27",
	}

	for key := range mapList {
		fmt.Println("key is", key, "val is ", mapList[key])
	}

	//val是任意类型
	map_2 := map[string]interface{}{
		"name": "heige",
		"age":  27,
		"job":  "golang",
	}

	fmt.Println(map_2)
	for k, val := range map_2 {
		fmt.Println("key", k, "val ", val)
	}
}
