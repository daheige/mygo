package main

import (
	"fmt"
)

//结构体可以指定key的不同类型
type User struct {
	name string
	age  int
	job  string
}

type Foo map[string]int

func main() {
	// 声明方式1：结构体作为一个指针的方式声明
	ms := new(User) //采用new分配内存，返回User类型的指针
	ms.age = 27
	ms.name = "heige"
	ms.job = "nodejs"
	fmt.Println(ms)       //&{heige 27 nodejs}
	fmt.Println(ms.name)  //访问结构体上的属性
	fmt.Println(&ms.name) //属性的内存地址

	// 声明方式2：结构体作为 value type
	var hg_ms = User{name: "daheige", age: 27, job: "php"}
	fmt.Println(hg_ms)

	//声明方式3：字面量方式声明 struct as a literal
	//初始化一个结构体实例采用取地址返回结构体的内存地址
	//字面量方式声明，返回结构体实例的指针 底层仍然会调用 new ()
	hg_s := &User{name: "dheige", age: 23, job: "golang"} //指向User实例值的指针
	fmt.Printf("地址是%p\n", hg_s)
	// 无论变量是结构体类型还是结构体类型指针都是用.来访问属性
	fmt.Println(hg_s.age, (*hg_s).name) //结构体是引用类型，所以间接引用的方式访问属性

	var p *User
	p = &User{name: "heige"}
	p.name = "zhuwei"
	p.job = "php"
	// (*p).name = "zhuwei313" //可以省略*
	p.name = "fefesss"
	fmt.Println(p, p.job)

	//对结构体不能使用make分配内存
	// //试图用new来创建一个实例指针，对于map类型返回一个nil
	// mf := new(Foo)
	// mf["x"] = 1 //type *Foo does not support indexing
	// // (*mf)["x"] = 1 //panic: assignment to entry in nil map
	// fmt.Println(mf)

}
