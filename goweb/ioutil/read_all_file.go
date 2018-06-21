package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dirname := "../log"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	//遍历读取到的文件
	for _, v := range files {
		fmt.Println(v.Name())
	}

	//读取文件内容
	con, err := ioutil.ReadFile("../log/errors.log")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(con))
}

/**
errors.log
log_demo.go
mylog
slog_demo.go
test
write_log.go
*/
