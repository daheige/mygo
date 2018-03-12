package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	// "log"
	"net/http"
)

//定义表单提交字段映射关系和json数据映射关系的tag
type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	// Password interface{} `form:"password" json:"password" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func main() {
	r := gin.Default()
	r.POST("/login", func(ctx *gin.Context) {
		var user User
		var err error
		contentType := ctx.Request.Header.Get("Content-Type")

		//根据请求类型处理
		switch contentType {
		case "application/x-www-form-urlencoded":
			err = ctx.BindWith(&user, binding.Form)
		case "application/json":
			err = ctx.BindJSON(&user)
		}
		if err != nil {
			fmt.Println(err)
			// log.Fatal(err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    500,
				"message": "参数提交不合法",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user.Username,
			"pwd":  user.Password,
			"age":  user.Age,
		})
	})
	//采用ctx.Bind自行推断
	r.POST("/user-login", func(ctx *gin.Context) {
		var user User
		var err error
		//自行推断
		err = ctx.Bind(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":    500,
				"message": "参数提交不合法",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"user": user.Username,
			"pwd":  user.Password,
			"age":  user.Age,
		})

	})

	//通常响应会有html，text，plain，json和xml等
	r.GET("/test", func(ctx *gin.Context) {
		contentType := ctx.DefaultQuery("contentType", "json")
		var data = []string{"fefe", "ssss"}
		// var data []interface{}
		// res := map[string]interface{}{
		//  "code":    200,
		//  "message": "ok",
		// }

		if contentType == "json" {
			ctx.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "ok",
				"data":    data,
			})
		} else if contentType == "xml" {
			ctx.XML(http.StatusOK, gin.H{
				"code":    200,
				"message": "ok",
				"data":    data,
			})
		}
	})
	r.Run(":8000")
}

/**
 * 请求格式是form表单提交
 * curl -X POST http://127.0.0.1:8000/login -H "Content-Type:application/x-www-form-urlencoded" -d "username=heige&password=123&age=21"
{"age":21,"pwd":"123","user":"heige"}
//请求数据格式是json
curl -X POST httpe:application/json" -d '{"username":"heige","age":28}'
{"code":500,"message":"参数提交不合法"}
curl -X POST httpe:application/json" -d '{"username":"heige","password":"sss","age":28}'
{"age":28,"pwd":"sss","user":"heige"}
password必须是字符串类型
curl -X POST http:pe:application/json" -d '{"username":"heige","password":123,"age":28}'
{"code":500,"message":"参数提交不合法"}

//自行推断
curl -X POST http://127.0.0.1:8000/user-login -H "Content-Type:application/json" -d '{"username":"heige","password":"sss","age":28}'
{"age":28,"pwd":"sss","user":"heige"}

curl -X GET http://127.0.0.1:8000/test
{"code":200,"data":["fefe","ssss"],"message":"ok"}
*/
