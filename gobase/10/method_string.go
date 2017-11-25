package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

//如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
//还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法
//自定义String
//当你广泛使用一个自定义类型时，最好为它定义 String()方法。
//格式化描述符 %T 会给出类型的完全规格，%#v 会给出实例的完整输出，包括它的字段（在程序自动生成 Go 代码时也很有用）

func (this *TwoInts) String() string {
	return "(" + strconv.Itoa(this.a) + "/" + strconv.Itoa(this.b) + ")"
}

func main() {
	t := &TwoInts{
		a: 1,
		b: 12,
	}

	fmt.Println("t is", t) //自动调用String
	fmt.Printf("t is %v\n", t)
	fmt.Printf("t is %T\n", t) //自动调用String()
	fmt.Printf("t is %#v", t)  //t is &main.TwoInts{a:1, b:12} 把实例的定义全部打印出来

}
