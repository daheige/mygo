package main

import (
	"fmt"
)

func main() {
	//先声明后使用赋值
	var m1 map[string]string
	//再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	m1["a"] = "fefe"
	m1["b"] = "sss"
	fmt.Println(m1)

	//2 直接创建map
	m2 := make(map[string]string)
	m2["a"] = "daheige"
	m2["b"] = "zhuwei"
	fmt.Println(m2)

	//3 初始化+赋值一体化
	m3 := map[string]string{
		"a": "sss",
		"b": "dahei",
		"c": "heige313",
	}
	fmt.Println(m3)

	//=============map应用=============
	// 查找键值是否存在
	if v, ok := m3["a"]; ok {
		fmt.Println("m3中a的值是", v)
	} else {
		fmt.Println("m3中不存在a")
	}

	//遍历map 每次遍历的值顺序是不同的
	for k, v := range m3 {
		fmt.Println(k, v)
	}
}
