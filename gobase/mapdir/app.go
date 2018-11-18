package main

import (
	"log"
)

func main() {
	/**
	* 和数组不同，map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。
	* 但是你也可以选择标明 map 的初始容量 capacity，就像这样：make(map[keytype]valuetype, cap)。例如：
	map2 := make(map[string]float32, 100)
	当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1
	所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
	*/
	var m = make(map[string]string, 10)
	m["x"] = "111fe"
	m["y"] = "111fe"
	m["z"] = "111fe"
	m["p"] = "111fe"
	log.Println(m)
}
