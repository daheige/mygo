package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {　//第一个参数是相应客户端res,第二个参数请求req
	r.ParseForm()       //解析参数
	fmt.Println(r.Form) //输出到服务端 ,参数保存在r.Form中
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello heige") //写入到w是输出到客户端的，输出到客户端的内容
}
func main() {
	http.HandleFunc("/", sayHello)           //设置访问的路由，路由规则，回调函数
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("listenAndServe", err)
	}
}
