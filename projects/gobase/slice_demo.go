package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	//声明切片，指定长度
	hg_slice := make([]int, 5)
	hg_slice[0] = 1
	hg_slice[1] = 23
	hg_slice[2] = 35
	//未赋值的元素值是0,如果不指定元素的字面值，默认是底层数组的元素的零值
	fmt.Println(hg_slice) //[1 23 35 0 0]

	//使用字面量，声明切片
	hg_s := []int{1, 2, 3}
	//字面量的切片会自动推断长度和容量一样
	fmt.Println(hg_s, len(hg_s), cap(hg_s))

	// 指定某个元素的值，其他默认采用零值
	hg_s2 := []int{4: 1} //第五个值是1，其他都是0
	fmt.Println(hg_s2)

	//nil切片意味着指向底层数组的指针为nil，而空切片对应的指针是个地址
	var nilSlice []int //nil切片
	slice := []int{}   //空切片
	fmt.Println(nilSlice, slice)

	// 空切片不能再赋值
	// slice[0] = 1
	// fmt.Println(slice)
	slice = append(slice, 1)
	fmt.Println(slice) //[1]

	d := []byte{'h', 'e', 'i', 'g', 'e'}

	fmt.Println(d)

	s := d[:cap(d)-3] //切片的截取
	fmt.Println(d, s)

	//slice copy 它只复制较短切片的长度个元素
	hg_x := []int{1, 3}
	hg_y := []int{3, 4, 5}

	copy_num := copy(hg_y, hg_x)
	fmt.Println(copy_num, hg_x, hg_y)

	// 扩大hg_y的容量
	// hg_y := []int{3, 4, 5}
	hg_z := make([]int, len(hg_y), (cap(hg_y)+1)*2) //自定义扩容
	copy(hg_z, hg_y)
	hg_y = hg_z
	fmt.Println(cap(hg_y)) //8
	//增加hg_y的元素
	//append的签名是func append(s []T, x ...T) []T
	//追加元素的时候，s,x必须是同一中数据类型
	hg_y = append(hg_y, 1, 3, 4) //追加的元素是一个可变参数，即可以是一个或多个
	fmt.Println(hg_y)
	hg_y = append(hg_y, []int{1, 2}...) //通过展开的方式追加元素到切片中
	fmt.Println(hg_y)

	// fmt.Println(hg_y)
	strArr := copyDigits("digit.md")
	fmt.Println(strArr)

	//删除第5个索引的值
	index := 5
	strArr = append(strArr[:index], strArr[index+1:]...) //先截取到index位置的值，然后把index后面的元素追加到后面，形成了一个切片
	fmt.Println(strArr)

	//在切片中间插入元素
	//index = 5
	tmp := append([]string{}, strArr[index:]...) //保留index后面的元素
	strArr = append(strArr[:index], "890")       // 在index后面加入一个值
	strArr = append(strArr, tmp...)              //把剩余的切片加前面一部分上
	fmt.Println(strArr)

	//对比追加后的切片
	// [123 234 234 555 444 23 67 334]
	// [123 234 234 555 444 890 23 67 334]

}

func copyDigits(filename string) []string {
	b, _ := ioutil.ReadFile(filename) //读取整个文件内容到字节数组中
	// fmt.Println("content", b)
	reg := regexp.MustCompile(`\d+`) //查找字符串中所有的数字
	return reg.FindAllString(string(b), -1)
}

/**
* 1 当我们使用make初始化切片的时候，必须给出size。go语言的书上一般都会告诉我们，当切片有足够大小的时候，append操作是非常快的。
* 但是当给出初始大小后，我们得到的实际上是一个含有这个size数量切片类型的空元素
*
* 2 当我们用append追加元素到切片时，如果容量不够，go就会创建一个新的切片变量
*
* 3如果，在make初始化切片的时候给出了足够的容量，append操作不会创建新的切片
*
* 4 可以结合make预先分配好切片大小，然后采用copy的方式实现扩容处理
*  // hg_y := []int{3, 4, 5}
   hg_z := make([]int, len(hg_y), (cap(hg_y)+1)*2) //自定义扩容
   copy(hg_z, hg_y)
   hg_y = hg_z
   fmt.Println(cap(hg_y)) //8
*/
