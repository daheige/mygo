package main

import (
	"github.com/gin-gonic/gin"
	"goweb/hg-gin/application/routes"
)

func main() {
	router := gin.New()

	//待完成
	routes.WebRoute(router)

	router.Run(":8080")
}
