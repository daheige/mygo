# 关于golang1.11版本module研究
    什么是modules
    https://tip.golang.org/cmd/go/#hdr-Modules__module_versions__and_more
    A module is a collection of related Go packages. Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules. Modules replace the old GOPATH-based approach to specifying which source files are used in a given build.
    
    翻译过来：
    模块是相关Go包的集合。modules是源代码交换和版本控制的单元。 go命令直接支持使用modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。
    可以得到两个重要信息：

        1. Go命令行支持modules操作
        2. modules用来替换GOPATH的
        大家不需要太担心了，golang 1.11版本仅仅是指对modules的初步支持，之前老的GOPATH还是可以继续使用的，有人说是在golang 1.12去除，但是我觉得有点早了，毕竟人的惯性不是这么容易改变的。
# 如何使用modules
    modules是一个新的特性，那么就需要新的Golang版本进行支持了，可以到官网下载，一定要是go 1.11及以上的版本。

    还有人记得vendor刚刚出来时候golang提供的环境变量GO15VENDOREXPERIMENT吗？现在modules出来，按照惯例也提供了一个环境变量GO111MODULE，这个变量的三个1太有魔性了。

# golang1.11版本的module支持
    GO111MODULE
    GO111MODULE可以设置为三个字符串值之一：off，on或auto（默认值）。

    off，则go命令从不使用新模块支持。它查找vendor 目录和GOPATH以查找依赖关系;也就是继续使用“GOPATH模式”。
    
    on，则go命令需要使用模块，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod下载依赖。
    
    auto或未设置，则go命令根据当前目录启用或禁用模块支持。
    仅当当前目录位于GOPATH/src之外并且其本身包含go.mod文件或位于包含go.mod文件的目录下时，才启用模块支持。
# 初始化
    go mod init
    会在当前目录新建go.mod文件
# 下载依赖包
    go mod vendor
# 查看包
    go list -m
    细心的同学一定可以发现，执行go mod init [module]使用go.mod只有一行信息hg-mux，在执行 go build、 go test
    go list命令时会根据需要的依赖自动生成 require语句。
# 如何定义module
    如何定义一个modules，modules是由Go源文件目录结构定义的，如果目录下含有go.mod文件，该目录称为模块根目录（module root）。模块根目录及其子目录所有的Go包
    都是属于该modules的，但是如果子目录包含有了自己的go.mod文件就隶属于该modules。

# 主模块和构建列表
    The main module and the build list 暂且翻译为主模块和构建列表。
    “主模块”是包含运行go命令的目录的模块。 go命令通过查找当前目录中的go.mod或者当前目录的父目录，或者祖父目录，依次递归查找。

    go.mod文件可以通过require，replace和exclude语句使用的精确软件包集。

    require语句指定的依赖项模块
    replace语句可以替换依赖项模块
    exclude语句可以忽略依赖项模块
    go list，可以查看当前的依赖和版本.

# go mod用法
    go mod init      初始化modules
    go mod download 下载modules到本地cache 放在$GOPATH/pkg/mod
    go mod edit     编辑go.mod文件，选项有-json、-require和-exclude 
    go mod graph    以文本模式打印模块需求图
    go mod tidy     删除错误或者不使用的modules
    go mod vendor   生成vendor目录
    go mod verify   验证依赖是否正确
    go mod why      查找依赖
    go help mod edit 查看命令的帮助说明
# go的 mod与get
    go get这个命令大家应该不会陌生，这是下载go依赖包的根据，下载Go 1.11出来了，go get命令也与时俱进，支持了modules。
    go get 来更新 module:

    运行 go get -u 将会升级到最新的次要版本或者修订版本
    运行 go get -u=patch 将会升级到最新的修订版本（比如说，将会升级到 1.0.1 版本，但不会升级到 1.1.0 版本）
    运行 go get package@version将会升级到指定的版本号
    运行go get如果有版本的更改，那么go.mod文件也会更改

# go mod生产环境说明
    go1.11版本下go程序，建议用govendor来进行加载模块
    会在当前根目录下生成vendor目录和vendor/vendor.json

# 生成依赖vendor
    1.设置环境变量 
        go1.11版本后，当我们设置了GO111MODULE=auto
        vim ~/.bashrc
        #在 $GOPATH/src 外面且根目录有 go.mod 文件时，开启模块支持
        export GO111MODULE=auto

        source ~/.bashrc

    2.生成依赖vendor 
        cd /web/projects/hg-mux
        go mod vendor
# 本地缓存module
    cd /web/projects/hg-mux
    go mod download
# 查找依赖
    $ go mod why
    # hg-mux
    hg-mux

# go build构建项目
    下面我们使用go build来编译我们的代码：
    go build
    值得注意的是，新增了一个编译选项“-mod”，它有如下的可选值：

    go build -mod=readonly
    在这个模式下任何会导致依赖关系变动的情况都将导致build失败，前面提到过build能查找并更新依赖关系，使用这个选项可以检查依赖关系的变动。

    go build -mod=vendor
    意思是忽略cache里的包，只使用vendor目录里的版本。
    查看构建目录信息：
        $ tree -L 1
        .
        ├── apps
        ├── app.yaml
        ├── bin
        ├── docs
        ├── go.mod
        ├── go.sum
        ├── hg-mux
        ├── logs
        ├── main.go
        ├── README.md
        └── vendor
        我们的代码成功构建了，包管理都由go modules替我们完成了。

# 包的版本控制
    go.mod文件
    module hg-mux

    require (
        github.com/daheige/thinkgo v0.0.0-20181103012329-7c6c02264ad2
        github.com/gomodule/redigo v0.0.0-20180627144507-2cd21d9966bf
        github.com/gorilla/mux v0.0.0-20181012153151-deb579d6e030
        gopkg.in/yaml.v2 v2.2.1 // indirect
    )
    go.mod文件的前面部分是包的名字，也就是import时需要写的部分，而空格之后的是版本号，版本号遵循如下规律：

    vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef
    vX.0.0-yyyymmddhhmmss-abcdefabcdef
    vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef
    vX.Y.Z
    也就是版本号+时间戳+hash，我们自己指定版本时只需要制定版本号即可，没有版本tag的则需要找到对应commit的时间和hash值。

    默认使用最新版本的package。
# 修改包的版本号
    现在我们要修改依赖关系了，我们想使用chromedp 的v0.1.0版本，怎么办呢？
    参考（demo)只需要如下命令：
    go mod edit -require="github.com/chromedp/chromedp@v0.1.0"
    @后面加上你需要的版本号。go.mod已经修改了：

    module test

    require github.com/chromedp/chromedp v0.1.0
# go mod后续
    go modules是一个很大的主题，以后我还将进一步介绍它。
    因为go1.11刚发布不久，这篇文件作为探路，必定会有错误和疏漏，欢迎大家指正！
