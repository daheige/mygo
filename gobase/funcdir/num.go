package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		fmt.Println("请输入一个数字:")
		fmt.Scan(&i) //接受终端输入的数字
		if i > 999 && i < 10000 {
			break
		}

		fmt.Println("输入的数字不合法")
	}

	fmt.Println("大黑哥")
}
