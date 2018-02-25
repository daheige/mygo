package controllers //同一目录下的包名一样

import (
	"github.com/astaxie/beego"
)

//注册路由
type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//分配数据到模板中
	this.Data["Website"] = "daheige beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["title"] = "heige"
	this.TplName = "index.tpl"
}

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Data["name"] = "heige"
	this.Data["title"] = "测试模板"
	this.TplName = "Index/index.tpl" //指定模板路径
}
