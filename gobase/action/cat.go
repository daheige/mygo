package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

//是否要在打印的前面加行号
var numberFlag = flag.Bool("n", false, "number each line")

//查看文件内容
func cat(fp *bufio.Reader) {
	i := 1
	for {
		buf, err := fp.ReadBytes('\n') //读取文件内容到缓冲区
		if err == io.EOF {
			return
		}

		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
}

//运行结果
//go run cat.go ../demo/out.md ../demo/input.md
// go run cat.go -n ../demo/out.md
// 1 fefefe
// 2 ll
// 3 123
// 4 345

//可以查看多个文件内容
func main() {
	flag.Parse() //flag.Parse() 扫描参数列表（或者常量列表）并设置 flag, flag.Arg(i) 表示第i个参数

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			continue
		}
		cat(bufio.NewReader(f))
	}

}
