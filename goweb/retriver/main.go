package main

import (
	"fmt"
	"goweb/retriver/mock"
	"time"
)

func download(r mock.MyRetriver) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r mock.MyRetriver
	//抽象接口,必须由某个具体类型来实现,成为一个实例才可以调用接口中的方法
	r = &mock.Retry{
		Content:   "daheige",
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	fmt.Printf("%T %v\n", r, r) //*mock.Retry &{daheige Mozilla/5.0 1m0s}
	//interface有两个东西,一个是对象的值,一个是对象的指针
	switch v := r.(type) {
	case mock.MyRetriver:
		fmt.Println(v)
		fmt.Println("contents ", v.Get("http://www.baidu.com"))
	case *mock.Retry:
		fmt.Println("UserAgent", v.UserAgent)
	}

	// fmt.Println(download(r))

}
