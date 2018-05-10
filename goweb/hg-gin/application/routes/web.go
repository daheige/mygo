package routes

import (
	"github.com/gin-gonic/gin"
	"goweb/hg-gin/application/controller"
)

func WebRoute(router *gin.Engine) {
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "welcome hg-gin page",
			"data": []string{
				"php",
				"go",
			},
		})
	})

	// 将相同控制器的路由放在一个Router方法中,方便管理
	homeController := controller.HomeController{}
	homeController.Router(router)

}
