package main

import (
	"fmt"
)

//采用自动增长定义枚举常量
const (
	b = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

func main() {
	fmt.Println(b, kb, mb, gb, tb, pb)
	//1 1024 1048576 1073741824 1099511627776 1125899906842624

}
