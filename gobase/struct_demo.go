package main

import (
	"fmt"
)

//struct结构体相当于其他语言的类class
type User struct {
	Name  string
	age   int
	email string
}

//Go 语言里有两种类型的接收者:值接收者和指针接收者
func (u User) changeEmail(email string) {
	u.email = email
}

//方法的接收者是User类型的值的指针
// 接收者的类型是指向 user 类型值的指针,而不是 user 类型的值
func (u *User) setEmail(email string) {
	u.email = email
	fmt.Println(u)
}

//创建一个变量并初始化为其零值,习惯是使用关键字 var 。这种用法是为了更明
// 确地表示一个变量被设置为零值。如果变量被初始化为某个非零值,就配合结构字面量和短变量
// 声明操作符来创建变量

func main() {
	var bill User
	fmt.Println(bill)
	bill.Name = "heige"
	bill.age = 28
	bill.email = "heige@hotmail.com"

	fmt.Println(bill, bill.Name)

	bill.changeEmail("ddd@hotmail.com")
	fmt.Println(bill)

	// 使用指针接收者声明的方法
	//go内部会解引用,指针接收者被解引用为值
	bill.setEmail("ddd@hotmail.com")
	fmt.Println(bill)

	u := User{
		Name:  "daheige",
		age:   27,
		email: "11@hotmail.com",
	}

	u.setEmail("ddd222@hotmail.com")

	fmt.Println(u)
}
