package main

import (
	"fmt"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Printf("测试: %d + %d = %d\n", i, i+1, Add(i, i+1))
	}
}

/**
go test -bench=. -run=none
测试: 299994 + 299995 = 599989
测试: 299995 + 299996 = 599991
测试: 299996 + 299997 = 599993
测试: 299997 + 299998 = 599995
测试: 299998 + 299999 = 599997
测试: 299999 + 300000 = 599999
  300000          7687 ns/op
PASS
ok      goweb/gotest    2.367s

300000次 每次需要7687ns(纳秒)


benchtime可以指定测试时间
go test -bench=. -benchtime=3s -run=none
测试: 999994 + 999995 = 1999989
测试: 999995 + 999996 = 1999991
测试: 999996 + 999997 = 1999993
测试: 999997 + 999998 = 1999995
测试: 999998 + 999999 = 1999997
测试: 999999 + 1000000 = 1999999
 1000000          7641 ns/op
PASS
ok      goweb/gotest    7.705s

benchmem可以查看内存分配 go test -bench=. -benchmem -benchtime=3s -run=none
 1000000          7645 ns/op          24 B/op          3 allocs/op
PASS
ok      goweb/gotest    7.717s
每次都需要3 allocs/op (3次内存分配) 分配的内存大小24b
*/
