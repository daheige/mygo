package main

//读取请求首部

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header //得到是一个map
	fmt.Fprintln(w, h)
	fmt.Fprintf(w, r.Header.Get("User-Agent")) //获取某个header
}

func body(w http.ResponseWriter, r *http.Request) {
	body_len := r.ContentLength
	body := make([]byte, body_len)
	r.Body.Read(body) //读取body中的内容到body中
	fmt.Fprintln(w, string(body))

}

func getForm(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username") //url和表单中的字段值
	pwd := r.FormValue("pwd")
	fmt.Fprintln(w, "username:"+username+"pwd:"+pwd)
}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/header", headers)
	http.HandleFunc("/body", body)    //curl -id "user=heige&age=27" 127.0.0.1:8080/body
	http.HandleFunc("/form", getForm) //http://localhost:8080/form?username=sfefe&pwd=sss
	server.ListenAndServe()
}
