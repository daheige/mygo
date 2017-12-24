package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//get请求
func main() {
	// res, err := http.Get("https://www.baidu.com")
	// res, err := http.Get("http://localhost:1337/")
	// res, err := http.Get("http://localhost:1337/info")
	res, err := http.Get("http://localhost:1337/user")
	if err != nil {
		log.Fatalf("error is %s", err)
		return
	}

	defer res.Body.Close()                //请求结束后，关闭连接
	body, err := ioutil.ReadAll(res.Body) //请求的结果
	if err != nil {
		log.Fatalf("error:%s", err)
		return
	}

	fmt.Printf("result is %s", string(body))
}

/**
 * /info
 * result is {"code":200,"message":"ok","data":{"id":1,"name":"heige","sex":1}}
 * /user
 * result is heige
 */
