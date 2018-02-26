```
###需要安装的工具###
go get github.com/beego/bee
go get github.com/astaxie/beego
go get github.com/smartystreets/goconvey/convey

###参考文档###
https://beego.me/docs/install/bee.md

###app.conf###
    beego.BConfig.AppName="beepkg" 可以在控制器中设置配置项
    在app.conf的配置项目，放在一个结构体上，https://godoc.org/github.com/astaxie/beego#Config
    支持section块配置，对于不同模式runmode 下的配置
    [dev]
    httpport = 1337

    [prod]
    httpport = 8080

    [test]
    httpport = 1336

    支持include "database.conf" 导入配置文件


```
