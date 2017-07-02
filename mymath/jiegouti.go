package mymath

import (
	"fmt"
)

func GetList() {
	var user person
	user.name = "heige"
	user.age = 26
	user.sex = 1
	fmt.Println("test user" + user.name)
}

type person struct {
	name string
	age  int
	sex  int
}
