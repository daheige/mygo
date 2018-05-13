package main

import (
	"fmt"
)

func main() {
	var a bool = true
	fmt.Println(a)

	var b byte = 'a' //97 ascii码 0-255
	var c byte = 127 //0-255
	fmt.Println(b, c)
	fmt.Printf("格式化打印:%c %d\n", b, b)
	fmt.Printf("格式化打印:%d %d\n", 'A', 'a') //大写字母和小写字母相差32
	fmt.Printf("小写a转换为 %c\n", 'a'-32)

	var a1 int = 123
	var b1 float64 = 123.345
	var c1 float64 = 3.14
	fmt.Println(a1, b1)
	fmt.Println(c1)

	f2 := 2.345
	fmt.Printf("%T\n", f2) //类型 //float64

	s := "daheige"

	for k, v := range []rune(s) {
		fmt.Println(k, v)
		fmt.Println(string(s[k]))
	}

	fmt.Println("字符个数", len("大黑哥"))             //9一个汉字含有3个字符
	fmt.Println("字符串长度", len([]rune("大黑哥313"))) //6
	s_rune := []rune(s)
	l := len(s_rune) //7个字符
	fmt.Println(string(s_rune[1]), l)

	t1 := 10 + 3i
	fmt.Printf("t1 type is %T\n", t1) //complex128
	fmt.Println("t1 real is ", real(t1), "t1 imag ", imag(t1))
}
