package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//定义配置文件结构体
type myYaml struct {
	Enabled bool   `yaml:"enabled"` //指定yaml标签tag
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
}

//读取配置文件
func (c *myYaml) getConf(path string) *myYaml {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("read yaml file error: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, c) //第二个参数是一个任意类型的interface
	if err != nil {
		log.Fatalf("Unmarshal yaml err: %v", err)
	}

	return c
}

func main() {
	var c myYaml
	c.getConf("conf/conf.yaml")
	fmt.Println(c)
	fmt.Println(c.Enabled, c.Name, c.Path)
	/**运行结果
	      {true heige /usr/local/go}
	  true heige /usr/local/go
	*/
}
