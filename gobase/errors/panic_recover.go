package main

import (
	"fmt"
)

func badCall() {
	panic("oo bad end")
}

func test() {
	defer func() {
		//用来捕获异常
		if e := recover(); e != nil {
			fmt.Printf("panic %s\n", e)
		}
	}()
	badCall()
	fmt.Println("after badCall")
}

func main() {
	fmt.Println("calling test")
	test()
	fmt.Printf("test complete\n")
}

/**
 * calling test
panic oo bad end
test complete
*/
