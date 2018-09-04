package main

import "fmt"

type pool chan []string

func NewPool(c int) pool {
	return make(chan []string, c)
}

func (p pool) get() []string {
	var v []string
	select {
	case v = <-p:
	default:
		v = make([]string, 10)
	}

	return v
}

func (p pool) put(b []string) {
	select {
	case p <- b:
	default:
	}
}

func main() {
	p := NewPool(3)
	b := p.get()
	b[0] = "sss"

	//追加第4个元素
	b = append(b, "ddd")
	fmt.Println(b)

}
