package main

import (
	"fmt"
)

var i = 5
var str = "abc"

type Person struct {
	name string
	age  int
}

//定义一个空接口
//任何其他类型都实现了空接口（它不仅仅像 Java/C# 中 Object 引用类型），any 或 Any 是空接口一个很好的别名或缩写
//可以给一个空接口类型的变量 var val interface {} 赋任何类型的值
//每个 interface {} 变量在内存中占据两个字长：
//一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针
type Any interface{}

//定义一个容器可以存放任意类型的数据
type Vector struct {
	data []Any
}

func (this *Vector) Set(i int, val Any) {
	this.data[i] = val
}

func (this *Vector) Get(i int) Any {
	return this.data[i]
}

func main() {
	var hgval Any
	hgval = 1
	fmt.Println(hgval)

	hgval = 12.3
	fmt.Println(hgval)

	hgper := &Person{
		name: "daheige",
		age:  27,
	}

	fmt.Println(hgper)

	hgval = hgper
	fmt.Println(hgval)

	//空接口实现类型断言判断
	switch hgval.(type) {
	case int:
		fmt.Println("int type")
	case bool:
		fmt.Println("bool type")
	case string:
		fmt.Println("string type")
	case float32:
		fmt.Println("float32 type")
	case float64:
		fmt.Println("float64 type")
	case *Person:
		fmt.Println("Person pointer type", hgval)
	default:
		fmt.Println("unkown type")

	}

	//容器中存放的是任意类型
	vdata := &Vector{
		data: []Any{
			[]int{1, 3, 4},
			1,
			2,
			3,
			5,
			12.3,
			"daheige",
			[]string{"a", "b", "c"},
		},
	}

	fmt.Println(vdata)

	fmt.Println(vdata.Get(2))
	vdata.Set(3, "golang")
	fmt.Println(vdata.Get(3))

}
