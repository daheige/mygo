package ware

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
)

func ApiWare(m *martini.ClassicMartini) {
	//添加自定义的中间件
	m.Use(func(c martini.Context, res http.ResponseWriter, req *http.Request) {
		res.Header().Set("x-power", "daheige&111111")
		c.Next()

		fmt.Println("after req!")
	})
}
