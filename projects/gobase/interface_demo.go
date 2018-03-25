package main

import (
	"fmt"
)

/**
 * 接口是用来定义行为的类型，它是抽象的，这些定义的行为不是由接口直接实现，而是通过方法由用户定义的类型实现。
 * 1 如果用户定义的类型，实现了接口类型声明的所有方法，那么这个用户定义的类型就实现了这个接口
 * 所以这个用户定义类型的值就可以赋值给接口类型的值。
 * 2 接口是一组方法的抽象，只要定义的类型，实现了接口类型声明的所有方法，就认为定义的类型就实现了该接口
 * 3 实体类型以值接收者实现接口的时候，不管是实体类型的值，还是实体类型值的指针，都实现了该接口
 * 4 实体类型以指针接收者实现接口的时候，只有指向这个类型的指针才被认为实现了该接口
 * 5 如果是值接收者，实体类型的值和指针都可以实现对应的接口；
 * 6 如果是指针接收者，那么只有类型的指针能够实现对应的接口
 * 7 一个接口，可以被多个不同类型的定义的方法实现
 */

type animal interface {
	printInfo(name string)
}

//person类型实现了animal的printInfo方法，所以person实现了接口animal
type person struct {
	name string
	age  int
}

func (p *person) printInfo(name string) {
	fmt.Println(p.name, p.age)
	fmt.Println("recived name is ", name)
}

func invoke(a animal, name string) {
	a.printInfo(name)
}

type cat int

func (c cat) printInfo(name string) {

	fmt.Println("fefe")
}

type dog int

//dog类型的指针作为方法接收者
//方法接收者是一个类型指针
func (d *dog) printInfo(name string) {
	fmt.Println(&d, d, name)
}

func main() {
	// var a animal

	var tea = person{
		name: "heige",
		age:  23,
	}

	// a = &tea

	// a.printInfo("daheige")
	invoke(&tea, "sss")

	//实体类型以值接收者实现接口的时候，不管是实体类型的值，还是实体类型值的指针，都实现了该接口
	var c cat
	invoke(c, "345")
	invoke(&c, "345")

	var d dog
	d.printInfo("234")
	// invoke(d, "234") //不能将值类型作为指针类型传递给方法
	invoke(&d, "234")
}
