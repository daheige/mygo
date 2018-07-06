//对于大量的rune,byte转为字符串压力测试
//相对来说内存不拷贝
// str和bytes共用一片内存
// str := (*string)(unsafe.Pointer(&bytes))
// 效率高
// 执行命令go test -test.bench=".*" -count=5，可以看到如下结果： （使用-count可以指定执行多少次）
package hello

import (
	"fmt"
	"testing"
)

func makeBytes() []rune {
	var bytes []rune
	for i := 0; i < 100; i++ {
		// fmt.Println([]byte(strconv.Itoa(i)))
		for a := 'A'; a < 'Z'; a++ {
			bytes = append(bytes, a)
			bytes = append(bytes, 'x')
		}

	}

	return bytes
}

/*func Test_byte_to_str(t *testing.T) {
    t.Log(1)
    bytes := makeBytes()
    fmt.Println(string(bytes))

}

func Test_byte_tostring(t *testing.T) {
    bytes := makeBytes()
    //内存不拷贝
    // str和bytes共用一片内存
    str := (*string)(unsafe.Pointer(&bytes))
    fmt.Println(*str)

}*/

/*func Benchmark_ByteToStr1(b *testing.B) {
    b.StopTimer() //调用该函数停止压力测试的时间计数

    //做一些初始化的工作,例如读取文件数据,数据库连接之类的,
    //这样这些时间不影响我们测试函数本身的性能

    b.StartTimer() //重新开始时间
    for i := 0; i < 1024; i++ {
        bytes := makeBytes()
        //内存不拷贝
        // str和bytes共用一片内存
        str := (*string)(unsafe.Pointer(&bytes))
        fmt.Println(*str)
    }

}*/

/**
1000000000           0.26 ns/op
PASS
ok      hello/t 2.440s

real    0m2.740s
user    0m0.780s
sys 0m0.446s


*/

func Benchmark_ByteToStr2(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < 1024; i++ {
		bytes := makeBytes()
		fmt.Println(string(bytes))
		// fmt.Println(fmt.Sprintf("%x", bytes))
	}

}

/**
2000000000           0.32 ns/op
PASS
ok      hello/t 16.790s

real    0m17.115s
user    0m4.097s
sys 0m1.303s


*/
