package main

/*
go test测试函数
于Go的单元测试，它也是接受一个包名作为参数，如果没有指定，使用当前目录。
go test运行的单元测试必须符合go的测试要求。
    1 写有单元测试的文件名，必须以_test.go结尾。
    2 测试文件要包含若干个测试函数。
    3 这些测试函数要以Test为前缀，还要接收一个*testing.T类型的参数
*/
import (
	"fmt"
	"testing"
)

func Add(a, b int) int {
	return a + b
}

//go test -v
func TestAdd(t *testing.T) {
	fmt.Println(1)
	if Add(1, 2) == 3 {
		t.Log("1 + 2 = 3")
	}

	if Add(1, 1) == 3 {
		t.Error("1 + 1 = 3")
	}

	fmt.Println(222)
}

/**
$ go test -v
=== RUN   TestAdd
1
222
--- PASS: TestAdd (0.00s)
        hello_test.go:15: 1 + 2 = 3
PASS
ok      goweb/hello     0.060s
*/
