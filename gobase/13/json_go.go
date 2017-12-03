package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//定义结构体，设置json标签
type Address struct {
	Type    string `json:"type"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	//将数据写入json文件
	vc := &Address{
		Type:    "zhuwei222",
		City:    "beijing",
		Country: "china",
	}

	tmp, _ := os.OpenFile("tmp.json", os.O_CREATE|os.O_WRONLY, 0644)
	defer tmp.Close()
	fmt.Println(vc)
	enc := json.NewEncoder(tmp)
	hg_err := enc.Encode(vc)
	if hg_err != nil {
		log.Println("Error in encoding json")
	}

}
