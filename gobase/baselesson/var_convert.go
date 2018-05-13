//类型转换
package main

import (
	"fmt"
)

func main() {
	a := true
	fmt.Println(a)

	// fmt.Printf("a = %d", true)   //a = %!d(bool=true)无法转换
	// fmt.Printf("a = %d", int(a)) //cannot convert a (type bool) to type int

	//不同类型不能转换,兼容的类型才可以转换
	var c = 'a' //本质上就是数字,整型
	var d int

	d = int(c)
	fmt.Println(d) //97
}
