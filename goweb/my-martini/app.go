package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"goweb/my-martini/routes"
	"goweb/my-martini/ware"
)

func main() {

	m := martini.Classic() //返回*martini.ClassicMartini
	//导入路由
	routes.WebRoute(m)

	//加入中间件
	ware.ApiWare(m)

	//静态服务器
	//http://localhost:3000/index.html
	m.Use(martini.Static("public"))

	fmt.Println("martini has run")
	m.RunOnAddr(":3000")

}
