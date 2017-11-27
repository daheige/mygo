package main

import (
	"fmt"
)

//实现排序的接口
type hgSorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//实现接口hgSorter
type IntArray []int

func (this IntArray) Len() int {
	return len(this)
}

func (this IntArray) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

//实现数据切片的排序，交换位置
func SortData(data hgSorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data hgSorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

func SortInts(a []int) {
	SortData(IntArray(a))
}

func IntsAreSorted(a []int) bool {
	return IsSorted(IntArray(a))
}

func main() {
	data := []int{1, 2, 3, 89, -90, -1, 2}
	a := IntArray(data)
	SortData(a)
	fmt.Println(a)

	//实现整数切片排序
	b := []int{1, 2, 3, 34, 12}
	SortInts(b)
	fmt.Println(b)

	fmt.Println(IntsAreSorted(b))
}
