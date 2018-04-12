package main

//go中rune(int32)就是char
import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	s := "大黑哥222"
	fmt.Println(len(s)) //12
	fmt.Printf("%x\n", s)

	//一个汉字占据3个字节
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}

	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %x) ", i, ch) //ch是一个字节
	}

	//字符串中rune的个数
	fmt.Println("Rune count: ",utf8.RuneCountInString(s)) //6个rune
	bytes := []byte(s)
	for len(bytes) > 0{
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}

	fmt.Println()
	//转换为rune进行遍历
	for k,v := range []rune(s){
		fmt.Printf("(%d %c) ",k,v)
		//fmt.Println(string(v))
	}

	//字符串相关的操作包strings
	pos := strings.Index("heige313","he") //子字符串出现的位置
	fmt.Println(pos)

	fmt.Println(strings.TrimLeft("        fefefe"," "))
	fmt.Println(strings.ToLower("UploadFile"))
	fmt.Println(strings.Repeat("heige",2))

}
