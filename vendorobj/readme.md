# govendor使用
    使用vendor管理Golang项目依赖

# 背景说明
    Golang 官方并没有推荐最佳的包管理方案。到了1.5版本时代，官方引入包管理的设计，加了 vendor 目录来支持本地包管理依赖。官方 wiki 推荐了多种支持这种特性的包管理工具，如：Godep、gv、gvt、glide、govendor等。

    我们在项目中除了大量的使用Python外，也大量的使用了Golang构建高效基础运行服务。在使用Golang过程中，我们发现Golang程序缺少依赖库版本功能是一个非常令人头大的问题：某些依赖在某个commit之后发生了API变更之后，如果不修改代码很难兼容，然而开发者之间很有可能因为参与的时间不同，导致执行 go get 命令获取的版本不同，而导致在不同电脑上出现编译不通过问题。同时，在多个程序中，如果使用的commit版本不同，也可能会导致程序编译过程中出现不同的问题。

    在之前，我们解决这个问题有两个方案，一种是拆解 go get 命令的执行，首先创建对应依赖目录，利用git命令切换至指定的commit，然后执行 go install 命令。另外一种比较省事的方法是使用 godep 工具，这里就不做过多介绍了，具体可以参考文档或者搜索中文教程。

# 安装govendor
    下面简要介绍一个我在项目中用到的 -- govendor。
    该工具将项目依赖的外部包拷贝到项目下的 vendor 目录下，并通过 vendor.json 文件来记录依赖包的版本，方便用户使用相对稳定的依赖。
    对于 govendor 来说，依赖包主要有以下多种类型:

    状态      缩写状态    含义
    +local      l   本地包，即项目自身的包组织
    +external   e   外部包，即被 $GOPATH 管理，但不在 vendor 目录下
    +vendor     v   已被 govendor 管理，即在 vendor 目录下
    +std        s   标准库中的包
    +unused     u   未使用的包，即包在 vendor 目录下，但项目并没有用到
    +missing    m   代码引用了依赖包，但该包并没有找到
    +program    p   主程序包，意味着可以编译为执行文件
    +outside    外部包和缺失的包
    +all        所有的包
    Installation
    go get -u github.com/kardianos/govendor
    命令行执行 govendor，若出现以下信息，则说明安装成功。
    备注:
         1. If using go1.5, ensure GO15VENDOREXPERIMENT=1 is set
         2. 需要把 $GOPATH/bin/ 加到 PATH 中

# 关于vendor
    在Golang1.5之后，Go提供了 GO15VENDOREXPERIMENT 环境变量，用于将go build时的应用路径搜索调整成为 当前项目目录/vendor 目录方式。通过这种形式，我们可以实现类似于 godep 方式的项目依赖管理。不过起码在程序编译过程中，再也无需在其他端部署一个 godep 工具。

    在使用之前，需要安装一个辅助工具（如果Golang自改一个就好了）： go get -u -v github.com/kardianos/govendor

# demo
    面，我们用一个例子来说明。首先有一个名为 vendorproj 的项目。假如只有一个文件：app.go
    package main

    import (
        "github.com/gin-gonic/gin"
        "net/http"
    )

    func main() {
        router := gin.New()
        router.GET("/", func(ctx *gin.Context) {
            ctx.JSON(http.StatusOK, gin.H{
                "code":    200,
                "message": "this is test",
                "data":    []string{"php", "go", "js"},
            })
        })
        router.Run(":8081")
    }

    执行一下命令就可以生成vendor文件夹：
    $ govendor init
    app.go  vendor
    $ cd vendor/
    $ ls #查看文件
    vendor.json
    这个 vendor.json 会类似 godep 工具中的描述文件版本的功能。接下来，需要执行命令将当前应用必须的文件包含进来

    $ govendor add +external  #或者 govendor add +e
    如果需要更新或移除，可以参考一下工具的具体文档使用 update 或者 remove 命令即可。这样会在vendor目录下将必须的编译文件移入进来（注意：测试所需依赖并不包含，依赖项目的测试文件也不会包含）。
    $ ls
        github.com  gopkg.in  vendor.json
    $ cat vendor.json 
        {
            "comment": "",
            "ignore": "test",
            "package": [
                {
                    "checksumSHA1": "QeKwBtN2df+j+4stw3bQJ6yO4EY=",
                    "path": "github.com/gin-contrib/sse",
                    "revision": "22d885f9ecc78bf4ee5d72b937e4bbcdc58e8cae",
                    "revisionTime": "2017-01-09T09:34:21Z"
                },
                {
                    "checksumSHA1": "9nXAGrHM0OvuVPmoVnLsZCezKmQ=",
                    "path": "github.com/gin-gonic/gin",
                    "revision": "814ac9490a38bb955a742d8ea15fa140489ee658",
                    "revisionTime": "2018-04-22T07:04:38Z"
                },
                {
                    "checksumSHA1": "qEoVYhFT/TTaIfa9qI249vU9yHo=",
                    "path": "github.com/gin-gonic/gin/binding",
                    "revision": "814ac9490a38bb955a742d8ea15fa140489ee658",
                    "revisionTime": "2018-04-22T07:04:38Z"
                },
                {
                    "checksumSHA1": "woO1qIxIeQ1bcbPSiMfAKk3r4xg=",
                    "path": "github.com/gin-gonic/gin/json",
                    "revision": "814ac9490a38bb955a742d8ea15fa140489ee658",
                    "revisionTime": "2018-04-22T07:04:38Z"
                },
                {
                    "checksumSHA1": "O+2Q4Uc9Vidvc/ctORliL315ZVI=",
                    "path": "github.com/gin-gonic/gin/render",
                    "revision": "814ac9490a38bb955a742d8ea15fa140489ee658",
                    "revisionTime": "2018-04-22T07:04:38Z"
                },
                {
                    "checksumSHA1": "WX1+2gktHcBmE9MGwFSGs7oqexU=",
                    "path": "github.com/golang/protobuf/proto",
                    "revision": "e09c5db296004fbe3f74490e84dcd62c3c5ddb1b",
                    "revisionTime": "2018-03-28T16:31:53Z"
                },
                {
                    "checksumSHA1": "w5RcOnfv5YDr3j2bd1YydkPiZx4=",
                    "path": "github.com/mattn/go-isatty",
                    "revision": "6ca4dbf54d38eea1a992b3c722a76a5d1c4cb25c",
                    "revisionTime": "2017-11-07T05:05:31Z"
                },
                {
                    "checksumSHA1": "d5ty9+omxFH0EI8uZc58dvnLy6c=",
                    "path": "github.com/ugorji/go/codec",
                    "revision": "f3cacc17c85ecb7f1b6a9e373ee85d1480919868",
                    "revisionTime": "2018-04-07T10:30:00Z"
                },
                {
                    "checksumSHA1": "P/k5ZGf0lEBgpKgkwy++F7K1PSg=",
                    "path": "gopkg.in/go-playground/validator.v8",
                    "revision": "5f1438d3fca68893a817e4a66806cea46a9e4ebf",
                    "revisionTime": "2017-07-30T05:02:35Z"
                },
                {
                    "checksumSHA1": "ZSWoOPUNRr5+3dhkLK3C4cZAQPk=",
                    "path": "gopkg.in/yaml.v2",
                    "revision": "5420a8b6744d3b0345ab293f6fcba19c978f1183",
                    "revisionTime": "2018-03-28T19:50:20Z"
                }
            ],
            "rootPath": "vendorobj"
        }

    查看包依赖
        $ cd vendorobj
        $ govendor list
         v  github.com/gin-contrib/sse             
         v  github.com/gin-gonic/gin               
         v  github.com/gin-gonic/gin/binding       
         v  github.com/gin-gonic/gin/json          
         v  github.com/gin-gonic/gin/render        
         v  github.com/golang/protobuf/proto       
         v  github.com/mattn/go-isatty             
         v  github.com/ugorji/go/codec             
         v  gopkg.in/go-playground/validator.v8    
         v  gopkg.in/yaml.v2                       
        pl  vendorobj                              
          m github.com/json-iterator/go            
          m golang.org/x/sys/unix   

# 编译并执行
    切换到vendorobj
    通过设置环境变量 GO15VENDOREXPERIMENT=1 使用vendor文件夹构建文件。
    将$GOBIN加入环境变量中
    我的golang开发环境变量设置如下: vim ~/.bashrc
        export GOROOT=/usr/local/go
        export GOPATH=/mygo
        export GOBIN=$GOPATH/bin
        export GOPKG=$GOPATH/pkg
        export GOSRC=$GOPATH/src
        export GO15VENDOREXPERIMENT=1

        export PATH=$PATH:$GOROOT/bin:$GOBIN

    可以选择 export GO15VENDOREXPERIMENT=1 或者干脆 GO15VENDOREXPERIMENT=1 go build 执行编译。

    通过这种方式就可以保证程序能够实现类似Python中Virtualenv的模式，实现不同程序使用不同版本依赖的目的。
    查看生成的执行文件vendorobj
    #执行脚本
    heige@daheige:/web/mygo/vendorobj$ ./vendorobj 
    [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
     - using env:   export GIN_MODE=release
     - using code:  gin.SetMode(gin.ReleaseMode)

    [GIN-debug] GET    /                         --> main.main.func1 (1 handlers)
    [GIN-debug] Listening and serving HTTP on :8081

    生成的vendorob执行文件可以复制到线上环境部署就可以

# 相关govendor命令参数说明
    init     创建 vendor 文件夹和 vendor.json 文件
    list     列出已经存在的依赖包
    add      从 $GOPATH 中添加依赖包，会加到 vendor.json
    update   从 $GOPATH 升级依赖包
    remove   从 vendor 文件夹删除依赖
    status   列出本地丢失的、过期的和修改的package
    fetch   从远端库增加新的，或者更新 vendor 文件中的依赖包
    sync     Pull packages into vendor folder from remote repository with revisions
    migrate  Move packages from a legacy tool to the vendor folder with metadata.
    get     类似 go get，但是会把依赖包拷贝到 vendor 目录
    license  List discovered licenses for the given status or import paths.
    shell    Run a "shell" to make multiple sub-commands more efficient for large projects.

    go tool commands that are wrapped:
          `+<status>` package selection may be used with them
        fmt, build, install, clean, test, vet, generate, tool

