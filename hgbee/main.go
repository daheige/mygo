package main

import (
	"github.com/astaxie/beego"
	_ "hgbee/routers"
)

func main() {
	beego.SetStaticPath("/assets", "assets") //添加静态资源目录,必须在beego.Run()执行执行加入
	beego.Run()
}
