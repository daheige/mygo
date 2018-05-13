package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type resData map[string]interface{}

func main() {
	myfunc(1, 2, 4)
	myfunc()
}

//可变参数的应用,必须放在最后一个参数,args实际上是一个索引切片类型
func myfunc(args ...int) {
	fmt.Println(args)
	res := success(200)
	fmt.Println(string(res))
	fmt.Println(string(success(200)))
	fmt.Println(string(success(200, "ok")))
	fmt.Println(string(success(200, "ok", []string{"php", "go"})))
}

func success(data ...interface{}) []byte {
	args_len := len(data)

	if args_len < 1 {
		panic("code,message必选参数")
	}

	if _, ok := (data[0]).(int); !ok {
		panic("code必须是一个int类型")
	}

	res := resData{}
	res["code"] = data[0]
	res["req_time"] = time.Now().Unix()

	if args_len == 1 { //code,message
		res["message"] = ""
		res["data"] = nil
	}

	if args_len == 2 { //code,message
		res["message"] = data[1]
		res["data"] = nil
	}

	if args_len == 3 {
		res["message"] = data[1]
		res["data"] = data[2]
	}

	json_data, _ := json.Marshal(res)
	return json_data
}
