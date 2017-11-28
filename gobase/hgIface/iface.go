package hgIface

//实现排序的接口
type HgSorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//实现接口HgSorter
type IntArray []int

type StringArray []string

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
func SortData(data HgSorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(data HgSorter) bool {
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

func (p StringArray) Len() int           { return len(p) }
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func SortStrings(a []string)             { SortData(StringArray(a)) }
func StringsAreSorted(a []string) bool   { return IsSorted(StringArray(a)) }

//实现对星期的排序
type Day struct {
	Num       int
	ShortName string
	LongName  string
}

type DayList struct {
	Data []*Day
}

func (p *DayList) Len() int {
	return len(p.Data)
}
func (p *DayList) Less(i, j int) bool {
	return p.Data[i].Num < p.Data[j].Num
}
func (p *DayList) Swap(i, j int) {
	p.Data[j], p.Data[i] = p.Data[i], p.Data[j]
}
