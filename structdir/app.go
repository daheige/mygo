package main

import (
	"fmt"
)

type user struct {
	name string
	age  byte
}

//给user结构体定义一个方法
func (u user) Tostring() string {
	return fmt.Sprintf("%+v", u)
}

//结构体嵌套在另一个结构体中
type manager struct {
	user
	title string
	job   string
}

type HGInt int

//方法，给HGInt自定义方法
func (x *HGInt) inc() {
	*x++
}

func main() {
	var m manager
	m.user.name = "heige"
	m.user.age = 23
	m.title = "haha"
	m.job = "php"
	fmt.Println(m)

	hgM2 := manager{
		user: user{
			name: "daheige",
			age:  23,
		},
		title: "ddd",
		job:   "golang",
	}

	fmt.Println(hgM2)

	//顺序必须一样
	m3 := manager{
		user{
			"fefe",
			24,
		},
		"fefesss",
		"python",
	}
	fmt.Println(m3)
	//可以省略字段名
	m4 := manager{user: user{"ss", 23}, title: "ssfefe", job: "c++"}
	fmt.Println(m4)

	var x HGInt = 123
	x.inc()
	println(x)

	fmt.Println(m3.Tostring()) //调用匿名字段user上的方法

}
