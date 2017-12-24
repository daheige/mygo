package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIp   string `json:"serverIp"`
}

type ServerSlice struct {
	Servers   []Server `json:"servers"`
	ServersId string   `json:"serversId"`
}

func main() {
	var s ServerSlice
	var newServer Server
	newServer.ServerName = "heige_vpn"
	newServer.ServerIp = "127.0.0.1"
	s.Servers = append(s.Servers, newServer)
	s.Servers = append(s.Servers, Server{ServerName: "test_vpn", ServerIp: "127.0.0.1"})
	s.ServersId = "test_team"
	s_b, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("error is %s", err)
		return
	}
	// fmt.Println(string(s_b))

	body := bytes.NewBuffer([]byte(s_b)) //发送的数据是一个buf
	// fmt.Println(body)
	res, err := http.Post("http://localhost:1337/add-user", "application/json", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close() //请求结束后，关闭连接

	r, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	// fmt.Printf("%s", r)
	//将结果解析到interface
	var f interface{}
	json.Unmarshal(r, &f)
	fmt.Println(f)
	//类型断言访问
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case int:
			fmt.Println(k, "is int", vv)
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case map[string]interface{}:
			fmt.Println(k, "is array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type unknown")
		}
	}

}
