package main

import (
	"encoding/json"
	"fmt"
)

//同一个结构体中，成员属性和方法不能相同
//在 Go 中，类型就是类（数据和关联的方法）。Go 不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。
// 在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫 组件编程（Component Programming）。
// 许多开发者说相比于类继承，Go 的接口提供了更强大、却更简单的多态行为
type Base struct {
	Id int `json:"id"`
}

func (this *Base) GetId() int {
	return this.Id
}
func (this *Base) SetId(id int) {
	this.Id = id
}

type Person struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employee struct {
	Person
	Salary float32 `json:"salary"`
}

func main() {
	//简写模式，可以省略字段名称
	emp := &Employee{
		Person{
			Base{12},
			"heige",
			"zuwei",
		},
		123.45,
	}

	//写全所有的字段所对应的值
	emp2 := &Employee{
		Person: Person{
			Base:      Base{Id: 12},
			FirstName: "heige",
			LastName:  "zuwei",
		},
		Salary: 123.45,
	}

	fmt.Println(emp, emp.FirstName, emp.Id, emp.Salary, emp.LastName)
	fmt.Println(emp2, emp2.LastName)
	emp2.SetId(123) //改变id的值
	fmt.Println(emp2.GetId())

	//输出json
	json_byte, _ := json.Marshal(emp2)
	fmt.Println(string(json_byte))
}
