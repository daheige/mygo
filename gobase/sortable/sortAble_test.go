package sortable

import (
	"fmt"
	"log"
	"sort"
	"testing"
)

func TestSortAble(t *testing.T) {
	var intSlice = IntArr{
		1, 2, 3, 32, 2,
	}

	BubbleSort(intSlice)
	log.Println(intSlice)

	sort.Ints(intSlice)

	log.Println(intSlice)
	a := []int{1, 2, 5, 3, 4}
	fmt.Println(a)                            // [1 2 5 3 4]
	sort.Sort(sort.Reverse(sort.IntSlice(a))) //倒叙

	log.Println(a)

}
