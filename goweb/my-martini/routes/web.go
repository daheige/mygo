package routes

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render" //模板渲染
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
	//martini 具有http res,req签名方法ServeHTTP
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

	//模板渲染
	m.Get("/foo", func(r render.Render) {
		r.HTML(200, "hello", "daheige")
	})
	m.Get("/hg-foo", func(r render.Render) {
		r.HTML(200, "foo", "heige313")
	})

	m.Get("/json", func(r render.Render) {
		r.JSON(200, map[string]interface{}{
			"code":    200,
			"message": "success",
			"data":    []string{"php", "go", "nodejs"},
		})
	})
}
