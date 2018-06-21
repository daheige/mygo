/**
WriteFile
WriteFile 向文件中写入数据，写入前会清空文件。
如果文件不存在，则会以指定的权限创建该文件。
返回遇到的错误。
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)
	fmt.Println("write success")
}
