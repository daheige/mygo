package main

import (
	"log"
)

//日志操作

func main() {
	log.Println("大黑哥", "heige313") //2018/03/28 22:13:53 大黑哥 heige313
	log.Printf("heige %s", "gopher")

	//通过log.SetFlags设置文件名+行号
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("fefe") //2018/03/28 22:16:48 log_demo.go:15: fefe

	//设置前缀
	log.SetPrefix("[hg] ")
	//绝对路径和行号
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("1234") //[hg] 2018/03/28 22:18:23 /web/mygo/goweb/log/log_demo.go:19: 1234

}
