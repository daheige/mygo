package main

import (
	"fmt"
)

type Retrier interface {
	Get(name string) string
}

type R struct {
	contents string
}

//实现标准接口上的String方法
func (r *R) String() string {
	return fmt.Sprintf("Contents is %s", r.contents)
}

func (r *R) Get(name string) string {
	return r.contents
}

func main() {
	r := &R{
		contents: "daheige",
	}

	var myretrier Retrier
	myretrier = r

	fmt.Println(myretrier.Get("daheige"))
	fmt.Printf("%T\n", myretrier) //*main.R
	fmt.Printf("%v", myretrier)   //Contents is daheige
}
