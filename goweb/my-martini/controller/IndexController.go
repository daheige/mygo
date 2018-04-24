package controller

import (
	"encoding/json"
	"net/http"
)

type IndexController struct{}

//处理器 任何接口只需要有一个ServeHTTP方法，并具有w,r签名，就是一个处理器
func (c *IndexController) Index(res http.ResponseWriter, req *http.Request) {
	json_data, _ := json.Marshal(map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    []string{"php", "go", "nodejs"},
	})

	res.Write(json_data)
}

func (c *IndexController) Test() (int, string) {
	return 403, "hello"
}
