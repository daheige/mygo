# 跨平台编译

    默认我们go build的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，
    那需要怎么做呢？只需要指定目标操作系统的平台和处理器架构即可：

    SET CGO_ENABLED=0  // 禁用CGO
    SET GOOS=linux  // 目标平台是linux
    SET GOARCH=amd64  // 目标处理器架构是amd64

    使用了cgo的代码是不支持跨平台编译的

    然后再执行go build命令，得到的就是能够在Linux平台运行的可执行文件了。

    Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

    Linux 下编译 Mac 和 Windows 平台64位可执行程序：

    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build

    Windows下编译Mac平台64位可执行程序：

    SET CGO_ENABLED=0
    SET GOOS=darwin
    SET GOARCH=amd64
    go build
