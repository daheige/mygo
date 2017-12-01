package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	res, err := copyFile("../demo/target.md", "../demo/input.md")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res, "copy done!")
}

//复制文件到另一个文件
func copyFile(distName, srcName string) (w int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dist, err := os.OpenFile(distName, os.O_CREATE|os.O_WRONLY, 0644) //打开文件，如果不存在就创建一个
	if err != nil {
		return
	}
	defer dist.Close() //关闭文件句柄

	return io.Copy(dist, src)
}
