package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"goweb/my-martini/routes"
	"goweb/my-martini/ware"
	"os"
)

//初始化环境判断
func init() {
	//默认是开发环境development
	if _, err := os.Stat("/etc/go.env.production"); err == nil { //生产环境
		martini.Env = "production"
	} else if _, err := os.Stat("/etc/go.env.testing"); err == nil {
		martini.Env = "testing"
	} else if _, err := os.Stat("/etc/go.env.staging"); err == nil {
		martini.Env = "staging"
	} else {
		martini.Env = "development"
	}
}

func main() {
	m := martini.Classic() //返回*martini.ClassicMartini

	//加入中间件
	ware.ApiWare(m)
	ware.TplConfig(m)

	//导入路由
	routes.WebRoute(m)
	routes.ApiRoute(m)

	//静态服务器
	//http://localhost:3000/index.html
	m.Use(martini.Static("public"))

	fmt.Println(martini.Env)
	fmt.Println("martini has run")
	m.RunOnAddr(":3000")

}
