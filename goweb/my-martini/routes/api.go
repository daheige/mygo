package routes

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"net/http"
)

func ApiRoute(m *martini.ClassicMartini) {
	//http://localhost:3000/api/sss
	//路由分组,添加前缀
	m.Group("/api", func(r martini.Router) {

		r.Get("/info", func(params martini.Params) string {
			return "123456,fefefe"
		})

		// http://localhost:3000/api/json
		// {"code":200,"data":["php","go","nodejs"],"message":"success"}
		m.Get("/json", func(res http.ResponseWriter, req *http.Request) {
			json_data, _ := json.Marshal(map[string]interface{}{
				"code":    200,
				"message": "success",
				"data":    []string{"php", "go", "nodejs"},
			})

			res.Write(json_data)
		})

		//最后匹配这个路由
		r.Get("/:id", func(params martini.Params) string {
			return params["id"] + "fefefe"
		})

	}, apiMiddleWare) //为分组路由添加中间件
}

//原生http服务
func apiMiddleWare(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("X-token", "hg-martini") //设置header头
}
