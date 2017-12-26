package main

import (
	"fmt"
)

type User struct {
	Name     string
	Email    string
	Level    byte
	Privated bool
}

//方法 this是一个指向User结构体的指针类型 方法体可以看做一个函数Say(this *User,word String)第一个参数是默认的结构体的一个实例对象
func (this *User) Say(word string) {
	fmt.Println("Username is", this.Name, "email is ", this.Email, " say "+word)
}

//结构体嵌套在另一个结构体中
type Admin struct {
	Person *User //指针类型占据4-8个字节
	Sex    byte  //0-255
}

//以匿名字段的形式嵌套到另一个结构体中
type Teacher struct {
	User
	job string
}

func main() {
	//通过指针的方式赋值 user1指向一个结构体类型的指针
	user1 := &User{
		Name:     "heige",
		Email:    "zhuwei313@hotmail.com",
		Level:    1,
		Privated: true,
	}

	fmt.Println(user1)
	fmt.Println(user1.Name)    //heige
	fmt.Println((*user1).Name) //可以省略* 取值操作

	//调用结构体上的方法
	user1.Say("daheige ,go go...")

	superMan := &Admin{
		Person: &User{
			Name:     "daheige",
			Email:    "zhuwei313@hotmail.com",
			Level:    2,
			Privated: false,
		},
		Sex: 1,
	}
	superMan.Person.Say("fefe")

	tea := &Teacher{
		User: User{
			Name:     "xiaohei",
			Email:    "zhuwei313@hotmail.com",
			Level:    2,
			Privated: false,
		},
		job: "golang",
	}

	tea.Say("haha")
}
