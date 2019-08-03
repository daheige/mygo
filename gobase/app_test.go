package hg_test

import (
	"log"
	"strconv"
	"testing"
)

func BenchmarkAdd(t *testing.B) {
	t.Log("start test")
	for i := 0; i < t.N; i++ {
		Add(10000)
	}

}

func BenchmarkStr(t *testing.B) {
	t.Log("start test")
	var s = ""
	for i := 0; i < t.N; i++ {
		s += "fefe" + strconv.Itoa(i)
	}

	log.Println(s)

}

func Add(n int) {
	var s int
	for i := 0; i < n; i++ {
		s += i
	}

	log.Println("sum: ", s)
}

var maxLoop = 10000

func BenchmarkLoopSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		total := 0
		for j := 0; j <= maxLoop; j++ {
			total += j
		}
	}
}

//安装分析工具
// $ sudo apt install graphviz

//$ go test -bench=Add -trace trace.out
/**
2019/08/03 10:09:36 sum:  49995000
   50000             25055 ns/op
--- BENCH: BenchmarkAdd-4
    app_test.go:9: start test
    app_test.go:9: start test
    app_test.go:9: start test
    app_test.go:9: start test
    app_test.go:9: start test
PASS
ok      _/web/go/hg-test        2.481s

//查看trace.out
$ go tool trace trace.out
2019/08/03 10:11:36 Parsing trace...
2019/08/03 10:11:36 Splitting trace...
2019/08/03 10:11:36 Opening browser. Trace viewer is listening on http://127.0.0.1:42459

会默认打开浏览器
http://127.0.0.1:42459/


//测试Str
测试100000次，即test.B提供的N, ns/op表示每一个操作耗费多少时间(纳秒)
fefe99998fefe99999
  100000            179792 ns/op
--- BENCH: BenchmarkStr-4
    app_test.go:18: start test
    app_test.go:18: start test
    app_test.go:18: start test
    app_test.go:18: start test
PASS
ok      _/web/go/hg-test        18.192s

/*
$ go test -v -bench=LoopSum  -trace trace.out
goos: linux
goarch: amd64
BenchmarkLoopSum-4        300000              4154 ns/op
PASS
ok      _/web/go/hg-test        1.301s
*/
/**
看benchtest的参数: go help testflag
-bench grep
通过正则表达式过滤出需要进行benchtest的用例
-count n
跑n次benchmark，n默认为1
-benchmem 可以提供每次操作分配内存的次数，以及每次操作分配的字节数
打印内存分配的信息
-benchtime=5s
自定义测试时间，默认为1s
其中B/op表示每次调用需要分配16个字节。allocs/op表示每次调用有多少次分配

测试demo
一共执行4组压力测试
$ go test -v -bench=LoopSum -count=4 -benchtime=5s -benchmem -trace trace.out
goos: linux
goarch: amd64
BenchmarkLoopSum-4       2000000              4077 ns/op               0 B/op          0 allocs/op
BenchmarkLoopSum-4       2000000              4266 ns/op               0 B/op          0 allocs/op
BenchmarkLoopSum-4       2000000              4247 ns/op               0 B/op          0 allocs/op
BenchmarkLoopSum-4       2000000              4316 ns/op               0 B/op          0 allocs/op
PASS
ok      _/web/go/hg-test        51.106s

查看内存
go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
然后就可以用输出的文件使用pprof
go tool pprof profile.out

小结：
bench常用的flag
-bench regexp:性能测试，支持表达式对测试函数进行筛选。-bench .则是对所有的benchmark函数测试
-benchmem:性能测试的时候显示测试函数的内存分配的统计信息
－count n:运行测试和性能多少此，默认一次
-run regexp:只运行特定的测试函数， 比如-run ABC只测试函数名中包含ABC的测试函数
-timeout t:测试时间如果超过t, panic,默认10分钟
-v:显示测试的详细信息，也会把Log、Logf方法的日志显示出来
*/
