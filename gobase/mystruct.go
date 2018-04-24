package main

import (
	"fmt"
)

type IndexController struct {
}

func (c *IndexController) Index() {
	fmt.Println(111)
}

func main() {
	c := (&IndexController{}).Index
	// f := c.Index

	c()

}
