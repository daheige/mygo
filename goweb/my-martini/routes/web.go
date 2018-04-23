package routes

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"goweb/my-martini/diydata"
	"net/http"
)

func WebRoute(m *martini.ClassicMartini) {
	m.Get("/", func() string {
		return "daheige,this is martini!"
	})

	m.Get("/info", func() (int, string) {
		return 404, "not found!"
	})

	//兼容http的rw,req签名的处理函数
	m.Get("/test", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json") //设置header头

		//设置cookie
		myCookie := http.Cookie{
			Name:     "daheige",
			Value:    "dddd",
			HttpOnly: true,
			MaxAge:   60,
		}
		http.SetCookie(res, &myCookie)

		data := &diydata.UserInfo{
			Name: "daheige",
			Job:  "phper",
			Age:  28,
		}

		json_data, _ := json.Marshal(data)
		res.Write(json_data)
	})

	//绑定参数
	m.Get("/hello/:name", func(params martini.Params) string {
		return "hello," + params["name"]
	})

	//http://localhost:3000/api/sss
	//路由分组,添加前缀
	m.Group("/api", func(r martini.Router) {
		r.Get("/:id", func(params martini.Params) string {
			return params["id"] + "fefefe"
		})
		r.Get("/info", func(params martini.Params) string {
			return "123456,fefefe"
		})
	}, apiMiddleWare) //为分组路由添加中间件

}

func apiMiddleWare(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("X-token", "123456") //设置header头
}
