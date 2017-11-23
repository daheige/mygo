package main

import (
	"bytes"
	"fmt"
)

//通过 buffer 串联字符串
//实现方式比使用 += 要更节省内存和 CPU，尤其是要串联的字符串数目特别多的时候
func main() {
	var buf bytes.Buffer

	buf.WriteString("daheige黑哥") //往buf中写入字符串

	fmt.Println(buf.String()) //将buf中的字符转化为字符串

	//声明一个字节类型的切片
	var s1 []byte
	s1 = append(s1, 'a')
	fmt.Println(s1)
	s1 = AppendByteToSl(s1, []byte("dagegge黑哥"))
	fmt.Println(s1, string(s1))
	fmt.Println("\n")
	//获取buf分片的前5个字节
	fmt.Println([]byte(buf.String())[0:5])

	fmt.Println([]byte("黑哥")) //一个汉字占三个字节

}

func AppendByteToSl(s []byte, data []byte) []byte {
	//追加字符到切片s中
	for _, v := range data {
		s = append(s, v)
	}

	return s
}
