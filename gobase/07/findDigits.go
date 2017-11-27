package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digReg = regexp.MustCompile("[0-9]+")

func main() {
	var c = findDig("demo.md")
	fmt.Println(string(c))
}

func findDig(filename string) []byte {
	b, _ := ioutil.ReadFile(filename) //读取文件内容到一个字节切片中
	b = digReg.Find(b)                //根据正则查找数字
	c := make([]byte, len(b))         //创建一个切片，长度为读取的数字字节长度
	copy(c, b)                        //将读取的内容复制到新的切片中
	return c
}
