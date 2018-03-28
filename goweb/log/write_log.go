package main

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	//写入文件中
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
func main() {
	Info.Println("heige:", "333")
	Warning.Printf("daheige git %s\n", "daheige")
	Error.Println("欢迎访问")
}

/**
Info:2018/03/28 22:31:57 write_log.go:25: heige: 333
Warning:2018/03/28 22:31:57 write_log.go:26: daheige git daheige
Error:2018/03/28 22:31:57 write_log.go:27: 欢迎访问
*/
