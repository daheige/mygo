package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseController
}

//前置操作,返回处理器函数(中间件)
func (ctrl *HomeController) Before() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		fmt.Print(uri)
		ctx.Next()
		return
	}
}

// action
func (ctrl *HomeController) Index(ctx *gin.Context) {
	ctx.JSON(HTTP_SUCCESS_CODE, gin.H{
		"code":    200,
		"message": "ok",
		"data":    "this is test",
	})
}

func (ctrl *HomeController) Test(ctx *gin.Context) {
	ctx.String(200, "this is test page %s", "daheige")
}

type res struct {
	Id    int
	Books []string
	Desc  string
}

//返回路由列表设置
func (ctrl *HomeController) Router(router *gin.Engine) {
	router.GET("/json", ctrl.Index)

	//设置分组路由
	v1 := router.Group("v1").Use(ctrl.Before()) //采用中间件
	v1.GET("/test", ctrl.Test)

	// http://mygo.com/v1/hg
	//{"data":[{"Id":1,"Books":["bo","php"],"Desc":"program"},{"Id":2,"Books":["golang","php"],"Desc":"study notes"}]}
	v1.GET("/hg", func(ctx *gin.Context) {
		data := []res{
			{
				Id:    1,
				Books: []string{"bo", "php"},
				Desc:  "program",
			},
			{
				Id:    2,
				Books: []string{"golang", "php"},
				Desc:  "study notes",
			},
		}

		ctx.JSON(HTTP_SUCCESS_CODE, gin.H{
			"data": data,
		})

	})

	v1.GET("/user-info", func(ctx *gin.Context) {
		ctrl.Success(ctx, "", []string{"golang", "javascript"})
	})

	router.GET("/error", func(ctx *gin.Context) {
		ctrl.Error(ctx, HTTP_ERROR_CODE, "参数错误,错误的请求")
	})

}
