package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

//定义和配置文件中一样的结构体
type mySec struct {
	Section struct { //Section必须和conf.ini的块名字一样
		Enabled bool
		Path    string
	}
}

func main() {
	sec_config := mySec{}
	err := gcfg.ReadFileInto(&sec_config, "conf/conf.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println(sec_config)
	fmt.Println(sec_config.Section.Enabled, sec_config.Section.Path)
}
