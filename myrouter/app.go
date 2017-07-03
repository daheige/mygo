package main

import (
	"fmt"
	"net/http"
	"strings"
)

//自定义路由功能
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		sayHell(w, req)
		return
	}
	req.ParseForm()
	fmt.Println(req.Form)
	//	var id = req.Form["id"]
	if id, ok := req.Form["id"]; ok {
		fmt.Println(req.Form["id"]) //是一个字典
		fmt.Println("id = " + strings.Join(id, ""))
	} else {
		fmt.Fprintf(w, "empty id")
		return
	}

	fmt.Println(req.Method)
	http.NotFound(w, req)
	return
}
func sayHell(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "daheige")
}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8080", mux)
}
