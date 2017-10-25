package main

import (
	"fmt"
	"reflect"
)

//反射可以检测运行时候的状态和类型
//要去反射是一个类型的值(这些值都实现了空interface)，
// 首先需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)。
// t := reflect.TypeOf(i) //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
// v := reflect.ValueOf(i) //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Printf("x type is %s\n", v.Type())
	fmt.Println("x kind is float64 : ", v.Kind() == reflect.Float64)
	fmt.Println("value is", v.Float()) //获取里面的值
	fmt.Println(reflect.TypeOf(x))     //float64

	//修改值
	p := reflect.ValueOf(&x)
	pv := p.Elem() //获取里面的元素
	pv.SetFloat(2.21)
	fmt.Println(pv.Float())
	fmt.Println(reflect.TypeOf(p)) //reflect.Value
	fmt.Println("type is ", pv.Kind(), pv.Type())
}
