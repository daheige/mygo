package main

import "fmt"

//通过传递指针改变a,b的值
func chageValue(a,b *int){
	*b,*a = *a,*b
}

func printSlice(s []int)  {
	fmt.Printf("len=%d,cap=%d\n",len(s),cap(s))
}

func main()  {
	a,b := 13,12
	fmt.Println(a,b)
	chageValue(&a,&b)
	fmt.Println(a,b)

	//array
	arr := [3]int{1,3,4}
	fmt.Println(arr)

	var arr2 [5]string

	arr2[0]="afefe"
	arr2[1]="x"
	fmt.Println(arr2)

	arr3 := [...]int{1,3,45}
	fmt.Println(arr3)

	//定义一个任意类型的数组
	arr4 := [3]interface{}{1,"a",23.4}
	fmt.Println(arr4)

	for k,v := range arr4{
		fmt.Println(k,v)
	}

	//_省略k
	for _,v := range arr4{
		fmt.Println(v)
	}

	arr5 := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr5[2:6] //s1底层数组是arr5,s1的长度是4,cap容量是6
	s2 := s1[3:5]
	fmt.Println(s1,s2) //[2 3 4 5] [5 6]
	fmt.Println(len(s1),cap(s1)) //4,6
	fmt.Println("s2的len",len(s2),cap(s2)) //2,3 所以可取得s1[3:5

	//添加元素到s2,第二个参数是一个s2类型的可变参数
	s2 = append(s2,[]int{1,3,4}...)
	fmt.Println(s2)
	fmt.Println(cap(s2)) //6 容量增加了一倍,之前的底层数组如果没有使用就会gc掉
	//slice切片三要素：ptr指向底层数组的指针,len,cap

	var s []int //s的值默认是nil(零值)

	for i := 0;i<10;i++{
		printSlice(s)
		s = append(s,2 * i +1)
	}

	fmt.Println(s)
	/**容量成倍数增加
	len=0,cap=0
	len=1,cap=1
	len=2,cap=2
	len=3,cap=4
	len=4,cap=4
	len=5,cap=8
	len=6,cap=8
	len=7,cap=8
	len=8,cap=8
	len=9,cap=16
	 */

	 //指定长度和容量（底层数组长度)
	 s3 := make([]int,8,8)
	 s3[1]= 12
	 fmt.Println(s3) //[0 12 0]
	 copy(s3,[]int{1,3,4})
	 fmt.Println(s3)

	 //删除slice中的值
	 //删除第三个元素
	 s3 = append(s3[:2],s3[3:]...)
	 fmt.Println(s3)
	 front := s3[0]
	 tail := s3[len(s3)-1]
	 s3 = s3[:len(s3)-1] //去掉尾部的元素
	 fmt.Println(front,tail,s3)

}
