package main

import (
	"fmt"
)

func main() {
	var s = "daheige黑哥3321a"
	fmt.Println(s[0:10])

	c := []byte(s)
	c[0] = 'D'

	s2 := string(c)
	fmt.Println(s2)
}
