# golang1.11包管理
    go 1.11 有了对模块的实验性支持，大部分的子命令都知道如何处理一个模块
    比如 run build install get list mod 子命令，第三方工具可能会支持的晚一些。到 go 1.12 会删除对 GOPATH
    的支持， go get
    命令也会变成只能获取模块，不能像现在这样直接获取一个裸包。

# 可以用环境变量 GO111MODULE开启或关闭模块支持
    它有三个可选值： off,on,auto默认值是 auto

    GO111MODULE=off
    无模块支持，go 会从 GOPATH 和 vendor 文件夹寻找包。
    GO111MODULE=on
    模块支持，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod
    下载依赖。
    GO111MODULE=auto
        在 $GOPATH/src外面且根目录有 go.mod 文件时，开启模块支持。
        在使用模块的时候， GOPATH是无意义的，不过它还是会把下载的依赖储存在 $GOPATH/src/mod中，也会把 go install的结果放在 $GOPATH/bin中。

# 定义模块
    模块根目录和其子目录的所有包构成模块，在根目录下存在 go.mod
    文件，子目录会向着父目录、爷目录一直找到 go.mod
    文件。

    模块路径指模块根目录的导入路径，也是其他子目录导入路径的前缀。 go.mod
    文件第一行定义了模块路径，有了这一行才算作是一个模块。

# go.mod
    文件接下来的篇幅用来定义当前模块的依赖和依赖版本，也可以排除依赖和替换依赖。

    module example.com/m 

    require (
        golang.org/x/text v0.3.0
        gopkg.in/yaml.v2 v2.1.0 
    )

    replace (
        golang.org/x/text => github.com/golang/text v0.3.0
    )
    这个文件不用手写，可以用 go mod init example.com/m
    生成 go.mod的第一行，文件的剩余部分也不用担心，
    在执行 go build,go test,go list命令时会根据需要的依赖
    自动生成 require语句。

    官方建议经常维护这个文件，保持依赖项是干净的。对于国内用户来说，手动维护这个文件是必然的，因为你需要把 golang.org/x/text
    替换成 github.com/golang/text啊。不需要像以前那样以 hack 的方式替换 GOPATH
    中的依赖，我一开始还保持着老思维，居然想要去替换模块的下载缓存，不过如果用 GOPROXY 功能也确实可以做到替换。

# go list 操作命令
    go list 命令
    go list -m
    可以查看当前的依赖和版本
# go mod 命令
    这个子命令用来处理 go.mod
    文件，上一小节我们已经见过 go mod -init
    了。

    go mod fmt
    格式化 go.mod文件。
    
    go mod sync
    从 go.mod
    删除不需要的依赖、新增需要的依赖，这个操作不会改变依赖版本。
    
    go mod require=path@version
    添加依赖或修改依赖版本，这里支持模糊匹配版本号，详情可以看下文 go get
    的用法。

    go mod vendor生成 vendor 文件夹。
    其他的自行 go help mod查看。

# go get 命令
    获取依赖的特定版本，用来升级和降级依赖。可以自动修改 go.mod
    文件，而且依赖的依赖版本号也可能会变。在 go.mod 中使用 exclude
    排除的包，不能 go get下来。

    与以前不同的是，新版 go get
    可以在末尾加 @符号，用来指定版本。

    它要求仓库必须用 v1.2.0格式打 tag，像 v1.2少个零都不行的，必须是 语义化
    的、带 v前缀的版本号。

    go get github.com/gorilla/mux    # 匹配最新的一个 tag
    go get github.com/gorilla/mux@latest    # 和上面一样
    go get github.com/gorilla/mux@v1.6.2    # 匹配 v1.6.2
    go get github.com/gorilla/mux@e3702bed2 # 匹配 v1.6.2
    go get github.com/gorilla/mux@c856192   # 匹配 c85619274f5d
    go get github.com/gorilla/mux@master    # 匹配 master 分支


    go get可以模糊匹配版本号，但 go.mod
    文件只体现完整的版本号，即 v1.2.0
    、 v0.0.0-20180517173623-c85619274f5d
    ，只不过不需要手写这么长的版本号，用 go get
    或上文的 go mod -require
    模糊匹配即可，它会把匹配到的完整版本号写进 go.mod
    文件。

# go build 命令
    go build getmode=vendor
    在开启模块支持的情况下，用这个可以退回到使用 vendor 的时代。
