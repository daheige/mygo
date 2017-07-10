package main

/*
JSON对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型(T是Go语言中任意的类型)
Channel, complex和function是不能被编码成JSON的
嵌套的数据是不能编码的，不然会让JSON编码进入死循环
指针在编码的时候会输出指针指向的内容，而空指针会输出null
*/

import (
	"encoding/json"
	"fmt"
	"os"
)

//导出的json字段名必须大写，但可以用struct tag来实现
type server struct {
	ServerName string `json:"serverName"`
	ServerIp   string `json:"serverIp"`
}
type serverList struct {
	Servers []server `json:"servers"`
}

//struct tag用法
type hgServer struct {
	Id         int    `json:"-"` //不输出到页面
	ServerName string `json:"serverName"`
	Info       string `json:"info,string"` //转换为string

	//如果ip为空就不输出
	Ip string `json:"ip,omitempty"`
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
	m := f.(map[string]interface{}) //通过断言访问
	fmt.Println(m)
	fmt.Println(m["name"])

	//生成json
	var s serverList
	s.Servers = append(s.Servers, server{ServerName: "shanghai", ServerIp: "127.0.0.1"})
	s.Servers = append(s.Servers, server{ServerName: "xianggang", ServerIp: "127.0.0.2"})
	res, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err", err)
	}
	fmt.Println(res)         //结果是一个字节类型，需要转换为json字符串
	fmt.Println(string(res)) //{"servers":[{"serverName":"shanghai","serverIp":"127.0.0.1"},{"serverName":"xianggang","serverIp":"127.0.0.2"}]}

	//struct tag用法
	hgs := hgServer{Id: 3, ServerName: `go "1.0"`, Info: "Go 1.0 fefe", Ip: ``}
	hgres, err := json.Marshal(hgs)
	fmt.Println(string(hgres)) //{"serverName":"go \"1.0\"","info":"\"Go 1.0 fefe\""}
	os.Stdout.Write(hgres)
}
