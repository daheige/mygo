package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type myJsonConf struct {
	Enable bool   `json:"enable"`
	Name   string `json:"name"`
	Path   string `json:"path"`
}

func main() {
	file, err := os.Open("conf/conf.json")
	checkErr(err)
	defer file.Close()

	//创建一个从file读取并解码json对象的*Decoder,解码器有自己的缓冲，并可能超前读取部分json数据
	decoder := json.NewDecoder(file)
	conf := myJsonConf{}
	//从输入流中读取下一个json编码值并保存在conf指向的值里
	err = decoder.Decode(&conf)
	checkErr(err)
	//打印读到的数据
	fmt.Println(conf)
	fmt.Println(conf.Enable, conf.Name)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

	return
}
