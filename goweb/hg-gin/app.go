package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	//待完成

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "welcome hg-gin page",
			"data": []string{
				"php",
				"go",
			},
		})
	})

	app.Run(":8080")
}
