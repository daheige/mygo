//将文件内容读取到一个字符串中，然后写入文件中
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "../demo/input.md"
	outFile := "../demo/out.md"
	buf, err := ioutil.ReadFile(inputFile) //读取文件内容到buf中
	if err != nil {
		fmt.Fprintf(os.Stdout, "file error: %s\n", err)
		return
	}

	fmt.Printf("读取到的字符串是:%s\n", string(buf))
	err = ioutil.WriteFile(outFile, buf, 0644) //将buf写入文件中
	if err != nil {
		fmt.Printf("write error is %s\n", err)
	}
	fmt.Println("write success")
}
