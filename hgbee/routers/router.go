package routers

import (
	"github.com/astaxie/beego"
	"hgbee/controllers"
)

func init() {
	//注册路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test", &controllers.IndexController{})
}
