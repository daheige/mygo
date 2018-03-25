package main

import (
	"fmt"
)

type hello interface {
	sayHello(name string)
}

//嵌入类型（复用已经定义好的结构体或接口），可以实现别的语言的继承特性

type User struct {
	name  string
	email string
}

func (u User) say() {
	fmt.Println("i am user inner method")
	fmt.Println(u.name, u.email)
}

//User类型实现了接口hello
func (u User) sayHello(name string) {
	fmt.Println(u.name)
	fmt.Println("传入的name", name)
}

func sayhaha(h hello, name string) {
	h.sayHello(name)
}

type Admin struct {
	User  //匿名字段进行嵌套
	level int
}

//重写User类型的方法
func (a Admin) say() {
	fmt.Println("i am admin")
	fmt.Println("my name is", a.name, "my email is", a.email)
}

// 如果内部类型实现了某个接口，那么外部类型也被认为实现了这个接口

func main() {
	//使用字段名初始化赋值
	manager := Admin{
		User: User{
			name:  "heige",
			email: "zhuwei313@hotmail.com",
		},
		level: 1,
	}

	fmt.Println(manager)

	//省略字段名赋值
	m2 := Admin{
		User{
			name:  "daheige",
			email: "122@hotmail.com",
		},
		1,
	}
	fmt.Println(m2)
	fmt.Println(m2.name)      //先在m2外层中查找是否有字段name,再向里面的User类型的实例上查找name
	fmt.Println(m2.User.name) //通过User去访问name字段
	m2.say()                  //调用重写后的say方法
	m2.User.say()             //调用User上的say方法
	m2.sayHello("fefe")
	m2.User.sayHello("fefe")

	sayhaha(m2, "ssss")
	sayhaha(m2.User, "User作为参数") //m2.User作为参数

}
