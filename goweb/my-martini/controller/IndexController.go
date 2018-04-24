package controller

import (
	"encoding/json"
	"net/http"
)

type IndexController struct{}

func (c *IndexController) Index(res http.ResponseWriter, req *http.Request) {
	json_data, _ := json.Marshal(map[string]interface{}{
		"code":    200,
		"message": "success",
		"data":    []string{"php", "go", "nodejs"},
	})

	res.Write(json_data)
}
