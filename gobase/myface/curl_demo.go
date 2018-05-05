package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("curl argv count must be 2")
		os.Exit(1)
	}
}

func main() {
	//请求web端,获得响应
	// http.Get 函数会返回一个 http.Response 类型的指针。
	// http.Response 类型包含一个名为 Body 的字段,这个字段是一个 io.ReadCloser 接口类型的值
	res, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	//从body复制响应结果到Stdout
	//io.Copy 会把服务器的数据分成小段,源源不断地传给终端窗口,直到最后一个片段读取并写入
	// 终端, io.Copy 函数才返回
	io.Copy(os.Stdout, res.Body)
}

//go run myface/curl_demo.go http://wwwbaidu.com
//请求结果将会输出到终端
// $ go run myface/curl_demo.go http://www.baidu.com
// <!DOCTYPE html>
// <!--STATUS OK-->
