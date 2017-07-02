package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	height float32
}

//结构体嵌套
type Student struct {
	Human   //匿名字段
	spec    string
	address string
	int            //内置类型作为匿名字段
	name    string //相同字段
}

//method继承
//如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该method
func (h *Human) sayHi() { //传递指针，开销小，直接对变量的地址操作，可以改变变量的值
	fmt.Println("your name is ", h.name, "your age is ", h.age, "your height is ", h.height)
}

//method方法的重载
type Employee struct {
	Human
	company string
}

func (e *Employee) sayHi() {
	fmt.Println("your height is", e.height, "haha")
}

const Pi = 3.14

// method的语法如下：
// func (r ReceiverType) funcName(parameters) (results)
type Rectangle struct {
	width, height float64
}

//area的发出者是Rectangle，在结构体上实现了方法area，此时area作为Rectangle的方法
//方法调用xxx.funcname,接收者不同，结果不同
func (r Rectangle) area() float64 { //第一个参数是一个接收器,接着是函数名
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * Pi
}

//自定义类型map 格式type typeName typeLiteral(go内置的类型)
type months map[string]string //字典

func main() {
	mark := Student{Human{name: "heige", age: 26, height: 1.70}, "php", "shenzhen", 12, "heige313"}
	john := Student{Human{name: "maoge", age: 27, height: 1.65}, "java", "shanghai", 13, "maowencai"}
	fmt.Println(mark.Human.name, mark.age, john.name, john.height, john.address)
	fmt.Println("内置类型", john.int)
	fmt.Println("外部的name", john.name, "内部的name", john.Human.name)
	fmt.Println(john.name)

	rect := Rectangle{12, 10} //创建结构体
	fmt.Println("长方形面积", rect.area())

	cir := Circle{10}
	fmt.Println("r = 10的圆面积是", cir.area())

	m := months{
		"one":   "01",
		"two":   "02",
		"three": "03",
		"four":  "04",
	}

	fmt.Println(m["one"])

	//method继承
	mark.sayHi() //your name is  heige your age is  26 your height is  1.7
	john.sayHi()

	emp := Employee{Human{
		name:   "xiaoming",
		age:    27,
		height: 1.72}, "qq company"}
	emp.sayHi() //your height is 1.72 haha

}
