package main

import (
	"fmt"
	"mymath"
)

//函数声明
/*
		func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
	//这里是处理逻辑代码
	//返回多个值
	return value1, value2}
*/
//函数一般定义在最外部
func Hgtest(x int, y int) int {
	return x + y
}

func getMax(a, b int) (c int) { //指定返回值的名称和类型
	if a > b {
		c = a
	} else {
		c = b
	}
	return
}

//多个返回值
func getAll(x int, y int) (int, int) {
	return x*1 + 2, y*3 + 1
}

func getListRec(arg ...int) {
	for _, v := range arg { //通过for...range来遍历slice
		fmt.Println(v)
	}
}

//传值和传地址的区别
func add1(x int) int {
	x = x + 1
	return x
}

func hg_add(x *int) int { //将x的地址传入，*int是一个指针类型，参数类型是一个指针类型
	*x = *x + 1
	return *x
}

//函数作为值、类型
type CheckInt func(int) bool //声明了一个类型是CheckInt是一个函数，返回值是布尔值

func isOdd(x int) bool {
	if x%2 != 0 {
		return true
	}
	return false
}

func isEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

//过滤slice中的奇数和偶数

func filter(slice []int, f CheckInt) []int { //f是一个函数
	var res []int //第一个一个数字类型的切片，并未初始化
	for _, v := range slice {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

type person struct {
	name string
	age  int
}

func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	fmt.Println("黑哥，你好")
	//for用法
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
	//for 实现while
	x := 1
	for x < 5 {
		fmt.Println(x)
		x += 1
	}

	//switch用法
	y := 3
	switch y {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("this ok")
	default:
		fmt.Println("no found")
	}

	fmt.Println(Hgtest(1, 2))
	fmt.Println("2,3", getMax(2, 3))
	p, q := getAll(2, 1)
	fmt.Println(p, q)
	getListRec(1, 2, 3, 4)
	hg_x := 3
	hg_x1 := add1(hg_x)
	fmt.Println(hg_x1)
	fmt.Println(hg_x)          //依旧是3
	fmt.Println(hg_add(&hg_x)) //取出x的地址，作为参数传递给hg_add
	fmt.Println(hg_x)          //4

	//小结
	/*
			传指针使得多个函数能操作同一个对象。
		传指针比较轻量级 (8bytes),只是传内存地址，我们可以用指针传递体积大的结构体。如果用参数值传
		递的话, 在每次copy上面就会花费相对较多的系统开销（内存和时间）。所以当你要传递大的结构体的
		时候，用指针是一个明智的选择。
		Go语言中 channel ， slice ， map 这三种类型的实现机制类似指针，所以可以直接传递，而不用
		取地址后传递指针。
	*/

	//函数作为值、类型
	hg_slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(filter(hg_slice, isOdd))
	fmt.Println(filter(hg_slice, isEven))
	mymath.GetList()

	var user1 person
	user1.age = 24
	user1.name = "heige"
	fmt.Println(user1.age, user1.name)

	user2 := person{age: 23, name: "fefe"} //简短写法
	fmt.Println(user2)

	toler, diff := older(user1, user2)
	fmt.Println(toler, diff) //{heige 24} 1

}
