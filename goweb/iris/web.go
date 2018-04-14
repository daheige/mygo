package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()              //创建一个app实例
	app.Logger().SetLevel("debug") //设置日志level

	//使用中间件
	app.Use(recover.New())
	app.Use(logger.New())

	//请求方法

	app.Handle("GET", "/", func(ctx iris.Context) { //上下文
		ctx.HTML("welcome to iris")
	})

	//get method
	app.Get("/test", func(ctx iris.Context) {
		ctx.WriteString("this is test page")
	})

	app.Get("/foo", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "ok",
			"code":    200,
		})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

//访问http://localhost:8080/foo
