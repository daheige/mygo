package main

import (
	"fmt"
	"testing"
)

func Test_postfile(t *testing.T) {
	targetUrl := "http://localhost:8080/upload"
	filename := "./test/263e73650c2adad32726800e76322a3e.jpg"
	postFile(filename, targetUrl)
	str := "heige黑哥"
	m := []byte(str) // to byte type
	fmt.Println(m)
	fmt.Println(string(m))
	t.Log("ok")
}
