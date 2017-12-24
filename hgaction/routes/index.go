package routes

import (
	"hgaction/controllers"
	"net/http"
)

func SetRouters() {
	// 监听事件
	http.HandleFunc("/", controllers.Handler)
}
