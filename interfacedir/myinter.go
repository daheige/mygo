package main

import (
	"fmt"
	"strconv" //字符串转换
)

//interface是一组method签名的组合，我们通过interface来定义对象的一组行为
//interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能
// 自我实现， Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来
// 也像鸭子，那么这只鸟就可以被称为鸭子"
// interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口

type Human struct {
	name  string
	phone string
	age   int
}

type Employ struct {
	Human
	company string
}

func (h Human) SayHi(intro string) {
	fmt.Println("my name is", h.name, "i can say ", intro)
}

func (h Human) Song(ly string) {
	fmt.Println("this song is", ly)
}

//重写了SayHi
func (e Employ) SayHi(intro string) {
	fmt.Println("my name is ", e.name, "i can say ", intro, "my age ", e.age, "my phone is", e.phone)
}

type Men interface { //接口是一组抽象方法
	SayHi(intro string)
	Song(ly string)
}

type Element interface{} //定义一个空接口，用来存放任意类型的数值
type HgList []Element    //切片类型，存放的是一组空接口
type Person struct {
	name string
	age  int
	sex  int
}

func (p Person) getUserInfo() string {
	return "your name is " + p.name + " age is " + strconv.Itoa(p.age) + "sex is " + strconv.Itoa(p.sex)
}

//接口嵌入
type ConnectUsb interface {
	Iusb
	getInfo() string
	getName(name string) string
}

type Iusb interface {
	push(intro string) bool
	add()
}

type CellPhoneUSB struct {
	name string
}

func (cusb CellPhoneUSB) getInfo() string {
	return cusb.name + "fefe"
}
func (cusb CellPhoneUSB) getName(intro string) string {
	return intro + cusb.name + "test"
}

func (cusb CellPhoneUSB) push(intro string) bool {
	if intro == "heige" {
		return true
	}
	return false
}

func (cusb CellPhoneUSB) add() {
	fmt.Println(cusb.name + " is add!")
}

func main() {
	mark := Human{"heige", "124", 27}
	lucy := Human{"lucy", "1234446", 26}

	heige := Employ{Human{"daheige", "188", 26}, "shenzhen it"}
	tom := Employ{Human{"tom", "168", 26}, "shenzhen php"}

	var hg_i Men

	hg_i = mark
	hg_i.SayHi("hello")
	hg_i.Song("good is girl")

	lucy.SayHi("ssss")
	/*
	       you can say  hello
	   this song is good is girl
	   you can say  ssss

	*/
	heige.SayHi("haha")
	hg_i = heige
	hg_i.SayHi("hello world")
	hg_i.Song("just like me")

	hg_list := make([]Men, 4) //初始化一个切片，类型是Men接口
	hg_list[0], hg_list[1], hg_list[2], hg_list[3] = mark, lucy, heige, tom

	for _, value := range hg_list {
		value.SayHi("hello")
		value.Song("good is girl")
	}

	//空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface
	// 对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数
	// 值的时候相当有用，因为它可以存储任意类型的数
	var emptyInterface interface{}
	var i int = 5
	var s string = "heige"
	emptyInterface = i
	fmt.Println(emptyInterface)

	emptyInterface = s
	fmt.Println(emptyInterface)

	//接口存储的类型判断
	//interface的变量里面可以存储任意类型的数值(该类型实现了interface)。

	per_list := make(HgList, 3)
	per_list[0] = 1
	per_list[1] = "heige"
	per_list[2] = Person{name: "heige", age: 26, sex: 1}

	for index, element := range per_list {
		switch value := element.(type) { //通过element.(T) 方式判断类型
		case int:
			fmt.Printf("list[%d]type is int and value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] type is string and value is %d\n", index, value)
		case Person:
			fmt.Printf("list[%d] type is Person and base info is %s\n", index, value.getUserInfo())
		default:
			fmt.Printf("list[%d] type is unknown", index)
		}

		if value, ok := element.(Person); ok { //Comma-ok断言
			fmt.Printf("list[%d] type is Person and base info is %s\n", index, value.getUserInfo())
		}
	}

	phoneUsb := CellPhoneUSB{"iphone"}
	var conUsb Iusb = phoneUsb
	conUsb.push("haha")
	conUsb.add()
}
