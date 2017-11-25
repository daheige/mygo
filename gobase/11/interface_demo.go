package main

import (
	"fmt"
	"strconv"
)

const PI = 3.14

//定义接口，接口由一系列的抽象方法组成，不会自己具体实现，接口中不能包含变量
//只要某个类型实现了接口中的方法，就实现了该接口，实现接口的类型可以有自己的方法
type Shaper interface {
	Area() string
}

//Square 实现了接口Shaper
type Square struct {
	radis float32
}

//结构体上的方法
//定义了一个接收者类型是 Square 方法的 Area()
func (this *Square) Area() string {
	return strconv.FormatFloat(float64(this.radis*this.radis*PI), 'f', 2, 32)
}

func (this *Square) Info() string {
	return "radis is " + strconv.FormatFloat(float64(this.radis), 'f', 2, 32) + " area is " + this.Area()
}

func main() {
	s1 := &Square{
		radis: 2.12,
	}

	fmt.Println(s1.Area())

	//声明接口
	var shap Shaper
	shap = s1
	fmt.Println(shap.Area())
	// fmt.Println(shap.Info()) //接口不能调用类型的其他方法，因为接口上没有定义抽象方法Info
	fmt.Println(s1.Info())
	r := &Square{
		radis: 2.12,
	}
	q := &Square{
		radis: 2.12,
	}

	fmt.Println(r.radis)

	//定义一个接口类型的切片
	shapes := []Shaper{r, q}
	// fmt.Println(shapes)

	for n, _ := range shapes {
		fmt.Println(shapes[n].Area())
	}

}
