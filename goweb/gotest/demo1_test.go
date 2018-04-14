package main

import (
	"goweb/gotest/common"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//初始化路由设置
func init() {
	common.Routes()
}

//利用httptest.NewRecorder()创建一个http.ResponseWriter，模拟了真实服务端的响应
// 这种响应时通过调用http.DefaultServeMux.ServeHTTP方法触发的
func TestSendJson(t *testing.T) {

	//创建一个请求
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建请求req失败")
	}

	//响应
	res := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(res, req)

	log.Println(res.Header().Get("Content-Type"))
	log.Println("code: ", res.Code)
	log.Println("body: ", res.Body.String())

}

/**
=== RUN   TestSendJson
2018/04/14 12:30:27 applicaion/json
2018/04/14 12:30:27 code:  200
2018/04/14 12:30:27 body:  {"name":"heige大黑哥"}
--- PASS: TestSendJson (0.00s)
*/
