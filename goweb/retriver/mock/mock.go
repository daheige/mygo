package mock

import (
	"net/http"
	"net/http/httputil"
	"time"
)

//定义接口
type MyRetriver interface {
	Get(url string) string
}

//定义结构体Retriver
type Retry struct {
	Content   string
	UserAgent string
	TimeOut   time.Duration
}

//Retry实现了MyRetriver
func (r *Retry) Get(url string) string {
	resp, err := http.Get(url) //发送get请求
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	//从响应中获取结果
	result, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	return string(result)
}
