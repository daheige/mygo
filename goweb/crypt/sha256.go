package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello world"))
	fmt.Printf("%x\n", h.Sum(nil))
}
