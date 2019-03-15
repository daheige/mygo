/**
 * logrus无缝兼容了log包的方法
 * 可以自定义日志格式，日志输出的位置
 */
package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("fefefe") //默认输出到终端中

	//设置日志为json格式
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)   //日志输出到终端中
	log.WithFields(log.Fields{ //设置字段,log.Fields 其实一个map[string]interface{}
		"animal": "walrus",
	}).Info("A walrus appears")

	log.Info(111)
	log.Warn(222)
	log.Println("fefe", 222)
	log.Printf("%s", "daheige")

	//将日志记录到文件中
	// file, err := os.Create("test.log")

	//日志追加的方式写入
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println("create log file error: ", err)
		return
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println(11111)
	log.Printf("%s", "daheige")
}
