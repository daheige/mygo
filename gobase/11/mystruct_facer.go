package main

import (
	"fmt"
)

//接口是一组抽象方法或接口内嵌
type myfacer interface {
	setName(name string)
	getName(name string) string
	notify(info string)
}

type subfacer interface {
	myfacer //把setName,getName,notify都嵌入进来了,实现了接口继承
	say()
}

type mystruct struct {
	name string
	sex  int
	age  int
}

// person结构体内嵌了mystruct
type Person struct {
	mystruct
	hobby string
}

//mystruct实现接口myfacer,也实现了subfacer
//接收者是*mystruct
func (this *mystruct) setName(name string) {
	this.name = name
}

func (this *mystruct) getName(name string) string {
	return "hello" + this.name + name
}

func (this *mystruct) notify(info string) {
	fmt.Println("my name is "+this.name, " age is ", this.age, " sex is ", this.sex)
	fmt.Println("notice: ", info)
}

func (this *mystruct) say() {
	fmt.Println("haha", this.name)
}

func main() {
	var user = mystruct{
		name: "daheige",
		sex:  1,
		age:  28,
	}

	user.say()

	var m myfacer
	m = &user //接口变量必须是mystruct的实例指针(地址)
	m.setName("ddd")
	fmt.Println(m.getName("333"))
	m.notify("456")
	user.notify("3333")

	var m2 subfacer
	m2 = &user
	m2.notify("sss")

	m2.setName("zhuwei") //user的name发生了改变
	m2.say()
	fmt.Println(m2.getName("333"))

	fmt.Println(m.getName("1111"))
	user.say()

	//person结构体内嵌了mystruct

	// p := &Person{ //mystruct的接收者是指针类型或值类型都可以
	p := Person{
		mystruct: mystruct{
			name: "xiaoming",
			sex:  1,
			age:  28,
		},
		hobby: "php go",
	}

	p.mystruct.notify("sss")
	p.mystruct.say()
	// (&p.mystruct).say()
	p.say()

	p1 := &Person{
		mystruct{
			name: "ge3333",
			sex:  0,
			age:  23,
		},
		"go",
	}

	fmt.Println(p1)
	p1.mystruct.say()
	p1.say()            //调用mystruct的say方法可以省略匿名字段mystruct
	p1.setName("heige") //改变p1的mystruct上的name属性
	p1.say()

}
