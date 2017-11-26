package main

import (
	"fmt"
)

//定义一个通用接口
type Valuable interface {
	getValue() float32
}

//StockPos Car都实现了接口Valuable
type StockPos struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (this *StockPos) getValue() float32 {
	return this.sharePrice * this.count
}

type Car struct {
	imake string
	model string
	price float32
}

func (this *Car) getValue() float32 {
	return this.price * 2
}

//显示价格，调用接口实例上的getValue方法
func showValue(iface Valuable) {
	fmt.Printf("Value of the price is %f\n", iface.getValue())
}

func main() {
	var obj = &StockPos{
		ticker:     "jeep",
		sharePrice: 123.21,
		count:      1.2,
	}
	showValue(obj)

	var obj2 = &Car{
		imake: "China",
		model: "aiguozhe",
		price: 1235.23,
	}

	showValue(obj2)
}
