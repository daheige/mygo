//多重继承指的是类型获得多个父类型行为的能力，它在传统的面向对象语言中通常是不被实现的（C++ 和 Python 例外）。
//因为在类继承层次中，多重继承会给编译器引入额外的复杂度。
//但是在 Go 语言中，通过在类型中嵌入所有必要的父类型，可以很简单的实现多重继承。

package main

import (
	"fmt"
)

type Camera struct{}
type Phone struct{}

//通过内嵌组合的方式实现一个结构体有多个父结构体
type CameraPhone struct {
	name string
	Camera
	Phone
}

func (this *Camera) takepic() string {
	return "can take pic"
}

func (this *Phone) call() string {
	return "it can call"
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic() //调用实例上的方法，这里的实例指向Base
	self.Magic()
}

type Voodoo struct {
	Base
}

//这里Voodoo实现了自己的Magic方法
func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	came := &CameraPhone{
		name: "iphone7",
	}
	fmt.Println("iphone name is", came.name, came.call(), came.takepic())

	//方法的重载
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
