package main

import (
	"fmt"
	"math"
)

func Uint8FromInt(n int) (uint8, error) {
	//在uint8的范围中
	if n >= 0 && n <= math.MaxInt8 {
		return uint8(n), nil
	}

	return 0, fmt.Errorf("%d is out of uint8 range", n)
}

//将浮点数转化为int32类型
func float64Toint(f float64) (int, error) {
	if math.MinInt32 <= f && f <= math.MaxUint32 {
		whole, frac := math.Modf(f)
		if frac >= 0.5 {
			whole++
		}

		return int(whole), nil
	}

	return 0, fmt.Errorf("%f out of int32 range", f)

}

type ByteSize float64

//字节单位
const (
	_           = iota //0开始
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	var a int = 45
	var b int = 400000000
	var hg_a uint8 = 123

	c, ok := Uint8FromInt(a) //返回多个值
	// fmt.Println(c, ok)
	if ok != nil {
		fmt.Println(ok)
	}

	fmt.Println(c)

	if d, ok := Uint8FromInt(b); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(d)
	}

	// fmt.Println(d, ok)
	fmt.Println(hg_a + c) //类型转化为同一个可以做运算

	var f_num float64 = 123.445
	if f_int, ok := float64Toint(f_num); ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println("转化后的值是", f_int)
	}

	fmt.Println(KB, MB, GB, TB, PB)
}
