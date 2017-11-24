```
# mygo
go study notes
gobase目录是基于the-way-to-go_ZH_CN https://github.com/daheige/the-way-to-go_ZH_CN 做的demo
## 开发环境搭建
    请参考：notes/golang环境搭建和sublime配置.md
    
go开发常见的工具包
    fiximports
    gocode
    go-contrib-init
    godef
    godex
    go-find-references
    goimports
    gomodifytags
    gomvpkg
    go-outline-parser
    gopkgs
    gorename
    guru
    go-outline

上面的gocode安装方式(-u表示update更新包)
    go get -u github.com/nsf/gocode
    go install github.com/nsf/gocode
其他的包，我们通过克隆https://github.com/golang/tools.git后安装，参考如下：
    cd $GOPATH/src
    mkdir -p $GOPATH/src/golang.org/x
    cd $GOPATH/src/golang.org/x
    git clone https://github.com/golang/tools.git
    cd godoc
    go install

go-outline安装(用于vs code编辑器)
克隆https://github.com/ramya-rao-a/go-outline.git到$GOPATH/src/github.com目录下
    cd $GOPATH/src
    git clone https://github.com/ramya-rao-a/go-outline.git
    cd go-outline
    go install
    sudo cp $GOPATH/bin/go-outline $GOROOT/bin/
```
