package main

import (
	"errors"
	"fmt"
)

//自定义类型
var errNotFound error = errors.New("not found error") //通过errors.New建立一个错误提示符

func main() {
	fmt.Println("error", errNotFound)
	fmt.Printf("error %s\n", errNotFound) //fmt.Printf 会自动调用 String() 方法
	panic(1 / 0)                          //抛出异常，终止程序
	fmt.Println("ok")                     //这里不会运行

}
