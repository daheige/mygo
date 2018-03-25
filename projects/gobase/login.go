package main

import (
	"fmt"
	"projects/gobase/common"
)

func main() {
	fmt.Println("使用Loginer实现登录")
	l := common.NewLoginer()
	l.Login()
}
