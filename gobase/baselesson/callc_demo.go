package main

import (
	"fmt"
	"gobase/baselesson/callc"
)

func main() {
	res := callc.Add(1, 2)

	fmt.Println(res)

	fmt.Println(callc.Add(3, 4))
}
