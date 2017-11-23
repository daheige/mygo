package main

import "fmt"

//map是key/val的类型hash
//声明格式 map[keyType]valType
/**三种格式
 * 1 var x map[keyType]valType 初始化，并赋值为零值
 * 2 var x = make(map[keyType]valType) //分配内存，赋值零值
 * 3 x := map[keyType]valType{}声明并初始化
 */
func main() {
	var mapList = make(map[string]int) //采用make初始化map
	var mapSign map[string]int
	mapList["x"] = 1
	mapList["y"] = 2
	mapList["z"] = 3

	fmt.Println(mapList)
	fmt.Println(mapList["x"])

	//底层同时指向同一个map对象，当一个发生改变，另一个也跟着发生改变
	mapSign = mapList

	mapSign["p"] = 333

	fmt.Println(mapList, mapSign)

	//func作为map的值
	map_f := map[int]func() int{
		1: func() int { return 1 },
		2: func() int { return 10 },
		3: func() int { return 12 },
	}

	fmt.Println(map_f[1])
	fmt.Println(map_f) //map[3:0x488610 1:0x4885f0 2:0x488600] //整形都被映射到函数地址,val被映射到函数的地址

	//val为浮点类型的map
	yif := map[string]float32{
		"a0": 12.3,
		"b0": 12.4,
		"c1": 8.9,
	}

	fmt.Println(yif)
	fmt.Println(yif["a0"]) //访问字典中的值

	//val为slice切片的map
	map_sli := map[int][]int{
		1: {1, 2, 3},
		2: {3, 4, 5},
		3: {12, 32, 32},
	}

	fmt.Println(map_sli[1]) //通过key来访问
}
