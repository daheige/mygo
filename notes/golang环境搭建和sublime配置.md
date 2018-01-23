```
1、linux下安装好go1.8.3
下载go1.8.3.linux-amd64.tar.gz
cd /usr/local/
sudo wget https://www.golangtc.com/static/go/1.8.3/go1.8.3.linux-amd64.tar.gz
创建目录
    sudo mkdir /mygo
    sudo mkdir /mygo/bin
    sudo mkdir /mygo/src
    sudo mkdir /mygo/pkg
2、设置环境变量vim ~/.bashrc 或者sudo vim /etc/profile
export GOROOT=/usr/local/go
export GOOS=linux
export GOPATH=/mygo
export GOSRC=$GOPATH/src
export GOBIN=$GOPATH/bin
export GOPKG=$GOPATH/pkg
export PATH=$GOROOT/bin:$GOBIN:$PATH
3、source ~/.bashrc 生效配置

###开发设置如下#################
########gocode自动补全安装######
1、安装Git
sudo apt-get install git
2、下载gocode
go get github.com/nsf/gocode
3、安装gocode 进入/mygo/src执行如下操作
go install github.com/nsf/gocode
把生成的gocode复制到$GOROOT/bin下
sudo cp /mygo/bin/gocode $GOROOT/bin/

#####gosublime配置####
安装好sublime text3 并安装gosublime，配置如下
{
    "env": {
        "GOPATH": "/mygo",
        "GOROOT": "/usr/local/go"
    },
    "fmt_tab_indent": true,
    "fmt_tab_width": 4
}

安装gofmt插件进行格式化代码即可进行开发，配置gofmt
{
  "cmds": [
    ["gofmt", "-e", "-s"]
  ],
  "format_on_save": true,
  "fmt_tab_indent": true,
  "fmt_tab_width": 4,
  "tab_size": 4,
  "translate_tabs_to_spaces": false,
  "trim_automatic_white_space": true,
  "trim_trailing_white_space_on_save": true
}

Ctrl+B自动运行go
    Tools->Build System->New Build System
    新增一个golang的编译
     {
        "shell_cmd": "/usr/local/go/bin/go run $file",
        "encoding": "utf-8"
     }
```
