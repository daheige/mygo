package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type User struct {
	Name string
	age  int
}

func (u *User) say(info string) {
	fmt.Println(u.Name, "say ", info)
}

// notify使用指针接收者实现了接口notifier
func (u *User) notify() {
	fmt.Printf("sending message from %s", u.Name)
}

//参数是实现了接口notifier接口值
// 函数可以针对任意实体类型的值来执行 notifier 方法
//实现多态,只要接口值是一个实现了notifier的类型实体就可以调用该方法
func sendMess(n notifier) {
	n.notify()
}

type Admin struct {
	User
	hobby string
}

func (a *Admin) notify() {
	fmt.Println("this message from ", a.Name, " his hobby is", a.hobby)
}

func main() {
	u := User{
		Name: "daheige",
		age:  28,
	}
	u.say("ddd") //自动解引用为传值的方式给函数say

	u.notify()

	var n notifier
	n = &u
	fmt.Println()
	sendMess(n)
	sendMess(&u)

	fmt.Println("")

	lim := Admin{
		User: User{
			Name: "heige",
			age:  27,
		},
		hobby: "php golang",
	}

	sendMess(&lim) //通过传递一个实现了notifier的实体类型的接口值

	sendMess(&lim.User) //这里User也实现了notifier接口,所以也可以调用

}
