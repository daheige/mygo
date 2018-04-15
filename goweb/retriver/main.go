package main

import (
	"fmt"
	"goweb/retriver/mock"
)

func download(r mock.MyRetriver) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r mock.MyRetriver
	//抽象接口,必须由某个具体类型来实现,成为一个实例才可以调用接口中的方法
	r = &mock.Retry{
		Content: "daheige",
	}

	fmt.Println(download(r))

}
