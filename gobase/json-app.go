package main

/**
关于json发序列化需要注意的点：
golang按照map解析json结构时，int类型会被升级为float64
json中字符串，数字型，默认就是float（对应golang是float64）

这是由于 JSON 里的数字默认都会转成 Golang 的 float64 类型引起的，
使用 Golang 解析 JSON  格式数据时，若以 interface{} 接收数据，则会按照下列规则进行解析：
	go数据格式          json数据格式
    bool, for JSON booleans

    float64, for JSON numbers

    string, for JSON strings

    []interface{}, for JSON arrays

    map[string]interface{}, for JSON objects

    nil for JSON null
而浮点数打印时的默认规则是超过一定长度后会换成科学计数法打印。

因此，只要在打印时指定打印格式，或者（按照LZ示例里是整数的情况时），转换为整数打印

fmt.Println( int( a["id"].(float64) ) ) // 将 “id” 键申明为 float64 类型，再转换为 int 型
*/

import (
	"encoding/json"
	"log"
)

type Reply struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func main() {
	str := `{"code":0,"message":"ok","data":{"id":2,"name":"daheige"}}`
	res := &Reply{}

	err := json.Unmarshal([]byte(str), res)
	if err != nil {
		log.Println("json decode error: ", err)
		return
	}

	log.Println(res.Data)
	log.Println(res.Data["id"])
	num := res.Data["id"]
	log.Printf("%T", num) //float64

	//先将数字转换为float64,然后转换为int/int64就可以
	id := int64(res.Data["id"].(float64))

	log.Println(id, id == 2)

}

/**
$ go run app.go
2019/08/03 11:42:03 map[id:2 name:daheige]
2019/08/03 11:42:03 2
2019/08/03 11:42:03 float64
2019/08/03 11:42:03 2 true
*/
