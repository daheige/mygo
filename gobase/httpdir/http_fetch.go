package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//请求远端数据，写入文件中
func main() {
	res, err := http.Get("http://www.baidu.com")
	checkError(err)

	defer func() {
		res.Body.Close()
		fmt.Println("请求结果写入文件成功")
	}()

	data, err := ioutil.ReadAll(res.Body) //读取请求到的内容
	// str := fmt.Sprintf("get data is %q", string(data))
	ioutil.WriteFile("./req_data.md", data, 0644)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("get err %v", err)
	}
}
