package main

import (
	"fmt"
	"reflect"
)

type HgT struct {
	a, b, c string
}

//发射应用

func main() {
	var x float64 = 3.3
	fmt.Println("type :", reflect.TypeOf(x))
	fmt.Println("val :", reflect.ValueOf(x)) //通过传递一个 x 拷贝创建了 v
	v := reflect.ValueOf(x)
	fmt.Println("val type is", v.Type())
	fmt.Println("v kind", v.Kind())
	fmt.Println("value is ", v.Float())
	fmt.Printf("value is %5.2e\n", v.Interface()) //还原值
	y := v.Interface().(float64)
	fmt.Println(y)

	m := reflect.ValueOf(&x) //得到x的地址
	fmt.Println(m)
	m = m.Elem()
	fmt.Println(m)
	//是否能设置值
	fmt.Println(m.CanSet())
	m.SetFloat(12.3) //改变值
	fmt.Println(m)

	//反射一个结构类型。NumField() 方法返回结构内的字段数量；
	//通过一个 for 循环用索引取得每个字段的值 Field(i)

	var hgt = &HgT{"ab", "cd", "de"}
	fmt.Println(hgt)
	s := reflect.ValueOf(hgt).Elem()

	tt := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(i, tt.Field(i).Name, f.Type(), f)
	}

}
