package main

import (
	"testing"
)

/*
 go test -v
 func TestGoprintNum1(t *testing.T) {
    printNum1() //测试goroutine_2.go中的printNum1
}*/

//go test -run x -bench . -cpu 1
/*
300000         35471 ns/op
PASS
ok      goweb/part9 10.934s
*/
func BenchmarkGoprintNum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1() //测试携程运行的效率
	}
}

/*func BenchmarkprintNum1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        goPrint1()
    }
}*/
