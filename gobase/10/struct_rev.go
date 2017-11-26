package main

import (
	"fmt"
	"unsafe"
)

type TwoInts struct {
	a int
	b int
}

//结构体内嵌在别的结构体中
type SubInts struct {
	TwoInts
	c int
	d int
	e int
}

//非结构体类型上的方法
//类型在其他的，或是非本地的包里定义，在它上面定义方法都会报错

type IntVec []int //通过定义类型别名的方式定义类型IntVec

func main() {
	two1 := TwoInts{a: 1, b: 2}
	two1.AddPar(123)
	fmt.Println(two1.AddThem())

	hg_vec := IntVec{1, 2, 3, 4, 5}
	fmt.Println("hg_vec整型切片的元素之和为", hg_vec.Sum())

	//改变two1上属性a的值
	two1.changeA(123)
	fmt.Println(two1)
	two1.say()
	fmt.Println(two1)

	two2 := &SubInts{
		TwoInts{a: 1, b: 3},
		13,
		12,
	}
	fmt.Println(two2)
}

//定义结构体上的方法
//在 Go 中，（接收者）类型关联的方法不写在类型结构里面，就像类那样；
//耦合更加宽松；类型和方法之间的关联由接收者来建立。
//recv接收者就像是面向对象语言中的 this 或 self
//this指向TwoInts类型的一个指针
func (this *TwoInts) AddThem() int {
	return this.a + this.b
}

func (this *TwoInts) AddPar(a int) int {
	fmt.Println(this.a, this.b, a)
	return this.a + this.b + a
}

//非结构体上定义方法
func (this IntVec) Sum() (res int) {
	for val := range this {
		res += val
	}
	return
}

//###############指针或值作为接收者###############
//对值的拷贝
func (this TwoInts) say() {
	fmt.Println("占据的内存大小:", unsafe.Sizeof(this)) //发生了值的拷贝
	this.a = 1111
	fmt.Println("this.a is", this.a)
}

//当接收者是一个指针的时候，我们通过指针间接访问实例上的属性和方法，可以改变实例上的属性
func (this *TwoInts) changeA(a int) {
	fmt.Println("占据的内存大小:", unsafe.Sizeof(this))
	defer fmt.Println("this.a", this.a, "将发生改变！")
	this.a = a
}
