package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

//beego.Controller 拥有很多方法，其中包括 Init、Prepare、Post、Get、Delete、Head 等方法。
//我们可以通过重写的方式来实现这些方法，而我们上面的代码就是重写了 Get 方法
//beego 是一个 RESTful 的框架，所以我们的请求默认是执行对应 req.Method 的方法。
//例如浏览器的是 GET 请求，那么默认就会执行 MainController 下的 Get 方法
type HomeController struct {
	beego.Controller //嵌套了beego.Controller所有方法和属性
}

func (this *HomeController) Get() {
	this.Data["name"] = "fefessss"
	this.TplExt = "html"
	fmt.Println(beego.AppConfig.String("mysqluser")) //root可以动态获取app.conf的配置项
	this.TplName = "Home/index.html"
}
