package sortable

import (
	"sort"
)

/**
sort包中实现了３种基本的排序算法：插入排序．快排和堆排序．和其他语言中一样，这三种方式都是不公开的
他们只在sort包内部使用．所以用户在使用sort包进行排序时无需考虑使用那种排序方式，sort.Interface定义的三个方法：
获取数据集合长度的Len()方法、比较两个元素大小的Less()方法和交换两个元素位置的Swap()方法，
就可以顺利对数据集合进行排序。sort包会根据实际数据自动选择高效的排序算法
*/

//定义可排序的接口
//任何实现了Sortable接口的数据类型都可进行冒泡排序, 下面是对Bubblesort方法的改造
//任何实现了 Sortable 的类型（一般为集合），均可使用该包中的方法进行排序。这些方法要求集合内列出元素的索引为整数
type Sortable interface {
	Len() int           //返回类型的长度
	Swap(int, int)      //交换位置
	Less(int, int) bool //比较两个值的大小
}

//实现冒泡排序
func BubbleSort(s Sortable) {
	length := s.Len()
	for i := 0; i < length; i++ {
		for j := i; j < length-i-1; j++ {
			if s.Less(j, j+1) {
				s.Swap(j, j+1)
			}
		}
	}
}

//参数是sort.Interface
//通用类型的冒泡排序
func CommonBubbleSort(s sort.Interface) {
	length := s.Len()
	for i := 0; i < length; i++ {
		for j := i; j < length-i-1; j++ {
			if s.Less(j, j+1) {
				s.Swap(j, j+1)
			}
		}
	}
}

//int类型的切片
/*
在go中, 实现接口不需要像java那样显式的说明对某个接口的implements
只需要为类型提供所有interface中定义的方法即可. 此例中, 我们给IntArr提供了所有Sortable中定义的方法, 所以IntArr已经实现了Sortable接口.
接下来要做的是将IntArr类型的数据传递给BubbleSortable函数就可以了
*/

/*========================int类型的切片，实现冒泡排序=============*/
type IntArr []int

//返回类型的长度
func (arr IntArr) Len() int {
	return len(arr)
}

//交换位置
func (arr IntArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr IntArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

/*========================int64类型的切片，实现冒泡排序=============*/
type Int64Arr []int

//返回类型的长度
func (arr Int64Arr) Len() int {
	return len(arr)
}

//交换位置
func (arr Int64Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr Int64Arr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

/*========================string类型的切片，实现冒泡排序=============*/
type StrArr []string

func (arr StrArr) Len() int {
	return len(arr)
}

func (arr StrArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr StrArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

/*========================float64类型的切片，实现冒泡排序=============*/
type Float64Arr []float64

//返回类型的长度
func (arr Float64Arr) Len() int {
	return len(arr)
}

//交换位置
func (arr Float64Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr Float64Arr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

/*========================float类型的切片，实现冒泡排序=============*/
type FloatArr []float64

//返回类型的长度
func (arr FloatArr) Len() int {
	return len(arr)
}

//交换位置
func (arr FloatArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr FloatArr) Less(i, j int) bool {
	return arr[i] < arr[j]
}
