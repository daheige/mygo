package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		//gin.H封装了生成json的方式
		//c.JSON则返回json数据
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "ok testing",
		})
	})
	//冒号:加上一个参数名组成路由参数。可以使用c.Params的方法读取其值
	r.GET("/foo/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello %s", name) //c.String返回响应，顾名思义则返回string类型
	})
	//http://localhost:8080/test/fefe/fess
	r.GET("/test/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.String(http.StatusOK, message)
	})
	//默认查询字符串 http://localhost:8080/info?name=ssss
	r.GET("/info", func(ctx *gin.Context) {
		// ctx.String(200, "fefefe")

		name := ctx.DefaultQuery("name", "daheige") //设置默认参数
		ctx.String(200, "your name is %s", name)
	})
	//在终端模拟post请求
	// curl -X POST http://127.0.0.1:8080/form_post -H "Content-application/x-www-form-urlencoded" -d "message=hello&nick=daheige"
	// {"code":200,"message":"hello","nick":"daheige"}
	r.POST("/form_post", func(ctx *gin.Context) {
		message := ctx.PostForm("message")
		nick := ctx.DefaultPostForm("nick", "heige")
		ctx.JSON(200, gin.H{
			"code":    200,
			"nick":    nick,
			"message": message,
		})
	})

	//上传单个文件
	/**
	 * multipart/form-data转用于文件上传。
	 * gin文件上传也很方便，和原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中了
	 */
	//curl -X POST http://127.0.0.1:8080/upload -F "upload=@/web/pic/heige.jpeg" -H "Content-Type: multipart/form-data"
	r.POST("/upload", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		fmt.Println(name)
		file, header, err := ctx.Request.FormFile("upload")
		if err != nil {
			ctx.String(http.StatusBadRequest, "bad request!")
			return
		}
		filename := header.Filename

		fmt.Println(file, err, filename)

		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
			return
		}

		ctx.String(http.StatusCreated, "upload success")
	})

	//上传多个文件
	//curl -X POST http://127.0.0.1:8080/multi/upload -F "upload=@/web/pic/heige.jpeg" -H "Content-Type: multipart/form-data"
	r.POST("/multi/upload", func(ctx *gin.Context) {
		err := ctx.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
			return
		}

		formdata := ctx.Request.MultipartForm
		files := formdata.File["upload"]
		for i := range files {
			file, err := files[i].Open()
			defer file.Close()

			if err != nil {
				log.Fatal(err)
			}

			//将文件放入out
			out, err := os.Create(files[i].Filename)
			defer out.Close()

			if err != nil {
				log.Fatal(err)
			}
			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
			ctx.String(http.StatusCreated, "upload2 success")
		}
	})
	r.Run(":8080") //监听8080端口运行
}
