package main

import (
	"fmt"
	"gobase/hgMap" //调用自定义的包
)

func main() {
	var bar = hgMap.Map_s_i{
		"x":     123,
		"y":     234,
		"bravo": 32,
	}

	map_i := hgMap.Fclip(bar)
	fmt.Println(map_i)
}
