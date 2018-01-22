package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	User string `json:"user"`
	Name string `json:"name"`
}

func upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("upload")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {

			fmt.Fprintln(w, string(data))
			w.WriteHeader(200)
			w.Write([]byte("heige")) //往ResponseWriter进行写入内容

		}
	}
}

func demo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.baidu.com") //设置请求头
	w.WriteHeader(302)                                 //重定向设置,调用WriterHeader后不许在对首部写入内容
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	jsonData := &Post{
		User: "heige",
		Name: "fefesss",
	}
	if json, err := json.Marshal(jsonData); err == nil {
		w.Write(json)
	} else {
		w.Write([]byte("error params"))

	}

}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.HandleFunc("/upload", upload)
	http.HandleFunc("/demo", demo)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
