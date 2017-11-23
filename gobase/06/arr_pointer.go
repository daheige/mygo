//引用类型和传递指针
package main

import (
	"fmt"
)

func main() {
	var arr1 = new([5]int) //变量类型是一个指针，返回new(T)返回T的指针，返回一个内存地址&[0 0 0 0 0],T表示这种类型的变量
	// make操作分配内存空间，并进行初始化零值操作(赋值)
	// make跟new都可以用于内存分配。
	// new用于各种类型的内存分配，本质上跟java的new的功能是一样的。new(T)返回了一个指针，指向新分配的类型T的零值。
	// make只能用于给slice、map和channel类型的内存分配，并且返回一个有初始值(非零)的T类型，而不是指针T
	hg_arr := make([]int, 3) //返回一个切片类型
	fmt.Println(arr1, hg_arr)
	fmt.Println(arr1[0]) //访问值，不需要用*arr1,本身可以引用访问值
	fmt.Println("第一个元素的内存地址", &arr1[0], arr1[0])
	fmt.Println("arr1的值", *arr1) //*arr1表示取值，对内存地址前面加*表示取值

	//对数组内存的拷贝
	arr2 := *arr1
	arr2[0] = 12 //这里不会改变arr1
	fmt.Printf("arr1 内存地址是%p arr2的内存地址是%p\n", arr1, &arr2)
	fmt.Println(*arr1, arr2)

	var arr3 [5]int //是一个int类型的数组，返回变量的值
	fmt.Println(arr3)
	fmt.Printf("arr3 内存地址是%p\n", &arr3) //打印内存地址，转换说明是%p

	arr4 := [5]int{1, 3}
	fmt.Println(arr4)

	arr5 := [3]int{1, 2, 3}
	fp(&arr5) //arr5的值存放在内存地址中，&arr5内存地址是内存中一个唯一标识

	var p *string
	var s string = "fefefe"
	p = &s
	fmt.Println("指针变量p是一个内存地址，指向变量s的内存地址", p)
	fmt.Println("对p指针进行取值操作，实际上获取的是s变量的值", *p)
	fmt.Printf("s变量的内存地址是%p", p)
	fmt.Println("指针是一种特殊的变量类型，p指针变量的值是一个地址，它是", p, "自己的内存地址是", &p)

}

//1 &用来获取一个变量的内存地址，*p表示获取指针p内存地址中存放的值，slice,map,chan,array都是引用类型
//2 指针其实一个特殊的变量类型，值是另外一个变量的内存地址，其本身是一个变量，当然也有内存地址
//函数的参数是指针数组，a类型是一个指针,在传递参数的时候必须传入一个变量的内存地址

//a是指向int类型数组的指针
//数组指针是数组变量在内存中的地址
func fp(a *[3]int) {
	fmt.Printf("指针a的值是 %p", a)              //a是一个指针，存放的是另一个变量的内存地址
	fmt.Println("指针a存放的值是", *a)             //对a指针进行取值操作
	fmt.Println("访问指针a存放变量的第一个元素", (*a)[1]) //可以省略* 因为a[1]通过内存地址间接引用
	fmt.Println("指针a的内存地址是", &a)            //取地址

}
