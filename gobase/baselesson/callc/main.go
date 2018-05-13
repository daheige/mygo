package callc

import (
	"fmt"
)

func init() {
	fmt.Println("init callc")
}

func Add(a, b int) int {
	return a + b
}
