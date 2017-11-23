package main

import (
	"fmt"
)

const (
	WIDTH  = 10
	HEIGHT = 2
)

var screen [WIDTH][HEIGHT]int

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = x * y
		}
	}

	fmt.Println(screen)
}
