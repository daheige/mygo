package main

import (
	"testing"
)

//1.单元测试的函数名必须以Test开头，是可导出公开的函数
//2.文件必须_test.go接尾
//3.单元测试文件名_test.go前面的部分最好是被测试的方法所在go文件的文件名
//4.测试函数的签名必须接收一个指向testing.T类型的指针，并且不能返回任何值
// t是一个指向testing.T类型的指针
func TestAdd(t *testing.T) {
	sum := Add(1, 3)

	if sum == 4 {
		t.Log("1 + 3 = 4 ,res is ok~")
	} else {
		t.Log("res is error")
	}

	//测试多次
	sum = Add(3, 5)
	if sum == 8 {
		t.Log("3 +5 =8")
	} else {
		t.Log("3+5 != 8")
	}
}

/**
t$ go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
    demo_test.go:16: 1 + 3 = 4 ,res is ok~
    demo_test.go:23: 3 +5 =8
PASS
ok      goweb/gotest    0.002s
*/
