package main

import (
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

//方法的定义
//1 要么是属于一个结构体的，要么属于一个新定义的类型的

//2 方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者
//这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法
//3 unc和方法名之间增加的参数(p person),这个就是接收者

func (p Person) String() {
	fmt.Println("name is", p.Name, "age is", p.Age)
}

//使用指针作为接收者
//指针接收者传递的是一个指向原值指针的副本，指针的副本，指向的还是原来类型的值
//所以修改时，同时也会影响原来类型变量的值
func (p *Person) change(name string) {
	p.Name = name
}

func (p Person) modify(name string) {
	p.Name = name
}

func main() {
	var p = Person{
		Name: "daheige",
		Age:  28,
	}

	p.String() //调用结构体上的方法，通过类型的变量.方法名的方式调用

	p.modify("heige313")
	fmt.Println(p.Name) //daheige 通过值类型传递，并不会结构体上的成员属性
	p.String()

	p.change("heige313")
	p.String() //name is heige313 age is 28

	//关于方法的调用
	//方法的调用，既可以使用值，也可以使用指针，我们不必要严格的遵守这些
	//Go语言编译器会帮我们进行自动转义的，这大大方便了我们开发者
	(&p).String()

	//多返回值
	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("error open")
	}

	fmt.Println(file)
	content := make([]byte, 10)
	read_num, err := file.Read(content)
	fmt.Println(read_num, content, string(content))

	res, err := add(1, 3)
	fmt.Println(res, err)

	printStr(1, 3, 4, "fefe", true)
}

func add(a, b int) (int, error) {
	return a + b, nil
}

//可变参数的定义，在类型前加上省略号…
//可变参数本质上是一个数组，所以可以用for...range来遍历
func printStr(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}
