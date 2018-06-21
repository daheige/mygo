/**
读取 r 中的所有数据，返回读取的数据和遇到的错误。
如果读取成功，则 err 返回 nil，而不是 EOF，因为 ReadAll 定义为读取
所有数据，所以不会把 EOF 当做错误处理。

func ReadAll(r io.Reader) ([]byte, error)
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	//读取返回的内容
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
