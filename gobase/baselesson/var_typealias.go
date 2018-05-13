//类型别名
package main

import (
	"fmt"
)

type bigInt int64 //自定义类型bigInt,别名
type char byte

func main() {
	var a bigInt = 1234
	fmt.Printf("a type is %T\n", a) //main.bigInt
	var c char = 'b'
	fmt.Println(c)
}
