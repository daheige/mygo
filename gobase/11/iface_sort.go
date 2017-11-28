package main

import (
	"fmt"
	"gobase/hgIface"
)

func days() {
	SunDay := hgIface.Day{0, "SUN", "SunDay"}
	MonDay := hgIface.Day{1, "MON", "MonDay"}
	TuesDay := hgIface.Day{2, "TUE", "TuesDay"}
	WednesDay := hgIface.Day{3, "WED", "WednesDay"}
	ThursDay := hgIface.Day{4, "THU", "ThursDay"}
	FriDay := hgIface.Day{5, "FRI", "FriDay"}
	SaturDay := hgIface.Day{6, "SAT", "Saturday"}
	data := []*hgIface.Day{&TuesDay, &ThursDay, &WednesDay, &SunDay, &MonDay, &FriDay, &SaturDay}
	a := hgIface.DayList{data}
	hgIface.SortData(&a)
	if !hgIface.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.LongName)
	}
	fmt.Printf("\n")
}

func main() {
	data := []int{1, 2, 3, 89, -90, -1, 2}
	a := hgIface.IntArray(data)
	hgIface.SortData(a)
	fmt.Println(a)

	//实现整数切片排序
	b := []int{1, 2, 3, 34, 12}
	hgIface.SortInts(b)
	fmt.Println(b)

	fmt.Println(hgIface.IntsAreSorted(b))

	//实现字符串切片排序
	s := []string{"a", "bc", "ab", "de", "cd"}
	hgIface.SortStrings(s)
	fmt.Println(s)
	fmt.Println(hgIface.StringsAreSorted(s))

	//实现星期排序
	days()

}
