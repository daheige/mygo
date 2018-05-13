package main

import (
	"fmt"
	"os"
)

func main() {
	argv := os.Args
	fmt.Println(argv) //切片类型
	fmt.Println(len(argv))

	fmt.Printf("%T", argv[1]) //string

}
