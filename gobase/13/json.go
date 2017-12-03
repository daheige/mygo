package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//定义结构体，设置json标签
type Address struct {
	Type    string `json:"type"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	json_str := `{"name":"fefefefe","age":25,"pars":["fefefe","heige"]}`
	b := []byte(json_str)

	//解析到interface中
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("error json")
	}
	fmt.Println(f)

	var hg_empty interface{}
	json.Unmarshal([]byte(`{"name":"heige","age":27}`), &hg_empty)
	fmt.Println(hg_empty)

	m := f.(map[string]interface{}) //通过断言访问
	fmt.Println(m)
	fmt.Println(m["name"])

	pa := &Address{
		Type:    "heige",
		City:    "shenzhen",
		Country: "China",
	}

	json, _ := json.Marshal(pa) //转换为字节切片
	fmt.Println(string(json))

	//将json写入文件中
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(string(json))
}
