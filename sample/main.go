package main

import (
	"fmt"
	"os"
	"strings"
)

//init在main之前调用
func init() {
	fmt.Println("hello")
}
func main() {
	who := "world"
	if len(os.Args) > 1 {
		who = strings.Join(os.Args[1:], ",") //字符串连接
	}
	fmt.Println("hello~", who)
	fmt.Println("this is test")
}

/*
$ sample heige
hello
hello~ heige
this is test
*/
