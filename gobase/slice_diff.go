package main

//对两个slice做diff操作
import "log"

func main() {
	arr := []int64{
		1, 2, 3, 4,
	}

	arr2 := []int64{
		1, 2, 4, 5,
	}

	arrLen := len(arr)
	arr2Len := len(arr2)
	c1 := make(map[int64]bool, arrLen)
	c2 := make(map[int64]bool, arr2Len)
	for _, v := range arr {
		c1[v] = true
	}

	for _, v := range arr2 {
		c2[v] = true
	}

	arrCommon := make([]int64, 0, arrLen+arr2Len)
	addArr := make([]int64, 0, arr2Len)
	deleteArr := make([]int64, 0, 2)
	for _, v := range arr2 {
		if _, ok := c1[v]; ok { //c1,c2共同部分
			arrCommon = append(arrCommon, v)
		} else { //不在c1中，在c2中,就表示新增的
			addArr = append(addArr, v)
		}
	}

	for _, v := range arr {
		if _, ok := c2[v]; !ok { //c1中的元素不在c2中，就表示需要删除的元素
			deleteArr = append(deleteArr, v)
		}
	}

	log.Println("add arr: ", addArr)
	log.Println("common arr: ", arrCommon)
	log.Println("delete arr: ", deleteArr)
}

/**
$ go run arr_diff.go
2019/07/31 22:02:07 add arr:  [5]
2019/07/31 22:02:07 common arr:  [1 2 4]
2019/07/31 22:02:07 delete arr:  [3]
*/
