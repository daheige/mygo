//处理函数
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// 处理器函数 w底层实际是是一个指针
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello heige")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s ,go go...", "heige")
}

//处理器 任何接口只需要有一个ServeHTTP方法，并具有w,r签名，就是一个处理器
type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is handler")
}

//串联处理器函数 接受一个处理器函数作为参数，返回一个新的处理器函数
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fucName := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name() //得到调用的函数
		fmt.Println("this call handleFunc is " + fucName)                 //this call handleFunc is main.world
		h(w, r)
	}
}
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil, //可以绑定处理器，默认采用多路处理器
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/test", world)
	http.HandleFunc("/world", log(world))
	hello := helloHandler{}
	http.Handle("/foo", &hello)
	server.ListenAndServe() //DefaultServeMux是ServeMux结构的一个实例
}
