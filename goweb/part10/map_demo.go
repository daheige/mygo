package main

import (
	"fmt"
)

func main() {
	//字面量map
	m1 := map[string]string{
		"name": "heige",
		"job":  "php",
	}

	fmt.Println(m1)
	fmt.Println(m1["name"], m1["job"])

	//通过make创建map,初始化是empty map
	m2 := make(map[string]int)

	var m3 map[string]int //m3 == nil 初始化值是nil

	fmt.Println(m2, m2 == nil, m3 == nil, m3)

	//map的遍历是无序的
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	name, ok := m1["name"] //判断k是否在map中
	fmt.Println(name, ok)

	delete(m1, "name") //删除某个key
	fmt.Println(m1)

}
