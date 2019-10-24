package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"strings"
)

type info struct {
	Id   int
	Name string
}

func main() {
	s := []int{1, 2, 3}
	log.Println("cap(s) = ", cap(s)) //3

	s = append(s, []int{4, 5, 6, 7, 8}...)
	log.Println("cap(s) = ", cap(s)) //8

	s = append(s, 123)
	log.Println("cap(s) = ", cap(s)) //16

	s2 := make([]int, 0, 10)
	s2 = append(s2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...) //追加元素没有发生扩容
	log.Println("cap(s2) = ", cap(s2))                   //cap(s2) =  10

	s3 := make([]int, 1024)

	log.Println("cap(s3) = ", cap(s3), "len = ", len(s3)) //cap(s3) =  1024 len =  1024
	//当元素超过1024个之后，就按照1/4的方式扩容
	// 1024/4   + 1024 = 1280
	s3 = append(s3, []int{1, 2}...) //cap(s3) =  1280 len =  1026

	log.Println("cap(s3) = ", cap(s3), "len = ", len(s3))

	bf := strings.Builder{}
	bf.WriteString("fefe")
	bf.WriteString("daheige")
	log.Println(bf.String())

	bf2 := bytes.Buffer{}
	bf2.WriteString("fefe,fefeefessss")
	log.Println(bf2.String())

	a := []string{"fefe", "feef", "sss", "sfe", "abc", "eeee"}
	log.Printf("a pointer is %p\n", a)

	log.Println("cap = ", cap(a))

	//删除一个元素后，内存地址指向没有发生改变
	a = append(a[:0], a[1:]...)

	log.Printf("a pointer is %p\n", a)
	log.Println("cap = ", cap(a))

	log.Println("a = ", a, len(a))
	// b := make([]string, 0)
	b := make([]string, len(a))
	// 拷贝的长度为两个slice中长度较小的长度值
	copy(b, a) //复制a,b中最小长度的

	log.Println("b = ", b)
	log.Printf("pointer of b is %p\n", b)

	log.Println(len(b))

	c := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	log.Printf("c pointer: %p", c)
	log.Println(remove(c, 1))

	log.Printf("c pointer: %p", c)

	x := &info{
		Id:   1,
		Name: "fefe",
	}

	p := x

	p.Name = "fefe123"
	y := new(info)

	log.Printf("x pointer is:%p", x)
	log.Println(deepCopy(y, x))
	log.Println("y = ", y)
	log.Printf("y pointer is: %p", y)

	/**
		2019/10/24 21:56:27 x pointer is:0xc00000c940
	2019/10/24 21:56:27 <nil>
	2019/10/24 21:56:27 y =  &{1 fefe123}
	2019/10/24 21:56:27 y pointer is: 0xc00000c960
	*/

}

// 移除切片中的元素
func remove(data []int, idx int) []int {
	copy(data[idx:], data[idx+1:]) //目标地址，源地址
	return data[:len(data)-1]
}

// 深度复制对象
// 需要深拷贝的变量必须首字母大写才可以被拷贝
/* 深拷贝
对于深拷贝就比较好了解了，任何对象都会被完完整整的拷贝一份
拷贝对象与被拷贝对象不存在如何联系，也就不会互相影响。
如果你需要拷贝的对象中没有引用类型，那么对于Golang而言使用浅拷贝就可以了。

基于序列化和反序列化来实现对象的深度拷贝: */
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		log.Println("gob encode err: ", err)
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

/**
$ go run buf.go
$ go run mytest.go
2019/10/24 22:10:06 cap(s) =  3
2019/10/24 22:10:06 cap(s) =  8
2019/10/24 22:10:06 cap(s) =  16
2019/10/24 22:10:06 cap(s2) =  10
2019/10/24 22:10:06 cap(s3) =  1024 len =  1024
2019/10/24 22:10:06 cap(s3) =  1280 len =  1026
2019/10/24 22:10:06 fefedaheige
2019/10/24 22:10:06 fefe,fefeefessss
2019/10/24 22:10:06 a pointer is 0xc000054240
2019/10/24 22:10:06 cap =  6
2019/10/24 22:10:06 a pointer is 0xc000054240
2019/10/24 22:10:06 cap =  6
2019/10/24 22:10:06 a =  [feef sss sfe abc eeee] 5
2019/10/24 22:10:06 b =  [feef sss sfe abc eeee]
2019/10/24 22:10:06 pointer of b is 0xc0000a2050
2019/10/24 22:10:06 5
2019/10/24 22:10:06 c pointer: 0xc00001c200
2019/10/24 22:10:06 [1 3 4 5 6 7]
2019/10/24 22:10:06 c pointer: 0xc00001c200
2019/10/24 22:10:06 x pointer is:0xc00000c940
2019/10/24 22:10:06 <nil>
2019/10/24 22:10:06 y =  &{1 fefe123}
2019/10/24 22:10:06 y pointer is: 0xc00000c960
*/
