package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(res http.ResponseWriter, req *http.Request) {
	//将响应结果写入res
	fmt.Fprintf(res, "request url is %s and path is %s", req.URL, req.URL.Path[1:])
}

func main() {
	fmt.Println("request start")
	http.HandleFunc("/", sayHello) //监听事件
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("listen error", err.Error())
	}
}
