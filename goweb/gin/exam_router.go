package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	//重定向功能
	r.GET("/baidu", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
		return
	})

	//路由分组功能http://localhost:8080/v2/login
	v2 := r.Group("/v2")
	v2.GET("/login", func(ctx *gin.Context) {
		//设置cookie
		cookie := &http.Cookie{
			Name:     "username",
			Value:    "123",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(ctx.Writer, cookie)
		ctx.String(http.StatusOK, "heige login %s", "mygo")
		return
	})
	//获取cookie
	v2.GET("/get-cookie", func(ctx *gin.Context) {
		name, err := ctx.Request.Cookie("username")
		if err != nil {
			ctx.String(http.StatusBadRequest, "获取cookie失败")
			return
		}
		ctx.String(http.StatusOK, name.Value)
	})

	//http://localhost:8080/v2/test
	v2.GET("/test", MiddleWare(), func(ctx *gin.Context) {
		request := ctx.MustGet("request").(string) //获取上一个中间件中的值
		ctx.JSON(http.StatusOK, gin.H{
			"middile_request": request,
		})
	})

	//异步携程处理
	/**
	  * [GIN] 2018/03/12 - 22:20:17 | 200 |       12.99µs |             ::1 | GET      /async
	    Done! in path/async
	*/
	r.GET("/async", func(c *gin.Context) {
		//因为涉及异步过程，请求的上下文需要copy到异步的上下文，并且这个上下文是只读的
		cCp := c.Copy()
		//并发处理
		//异步的逻辑则看到响应返回了，然后程序还在后台的协程处理
		go func() {
			time.Sleep(1 * time.Millisecond)
			log.Println("Done! in path" + cCp.Request.URL.Path)
		}()
		c.String(http.StatusOK, "success")
	})

	//静态文件服务
	//http://localhost:8080/assets/heige.jpeg
	// r.Static("/assets", "./static")

	//模板渲染 Using LoadHTMLGlob() or LoadHTMLFiles()
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "heige313",
		})
	})
	r.Run(":8080")

	//自定义路由机制
	// http.ListenAndServe(":8080", r) //利用原生http方式注册路由
	//或者使用如下方式注册路由
	/*s := &http.Server{
	      Addr:           ":8080",
	      Handler:        r,
	      ReadTimeout:    10 * time.Second,
	      WriteTimeout:   10 * time.Second,
	      MaxHeaderBytes: 1 << 20,
	  }
	  s.ListenAndServe()*/
}

//注册中间件
//一般使用中间件，用来授权或注册cookie等等
//返回值是gin.HandlerFunc
func MiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("before middleware")
		ctx.Set("request", "clinet_request 313")
		ctx.Next()
		fmt.Println("before middleware")
	}
}

/**
 * http://localhost:8080/baidu
 */
