//ServeMux结构包含了一个映射，这个映射UI将url映射到对应的处理器handler上
//其他多路复用器
//  /   indexHandler
//  /test   testHandler
//ServeMux是一个结构，defaultServeMux是它的一个实例
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//请求第三个参数
func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello ,%s\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

//http://localhost:8080/hello/fefe
