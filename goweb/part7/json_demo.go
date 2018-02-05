package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name    string
	Age     int
	Classes []string
	Price   float32
}

//通过标签映射字段
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	st := &Stu{
		Name:    "heige",
		Age:     27,
		Classes: []string{"main", "go", "php"},
		Price:   12.09,
	}

	//将结构体数据转换为json字节数组
	json_data, _ := json.Marshal(st)
	fmt.Println(string(json_data)) //{"Name":"heige","Age":27,"Classes":["main","go","php"],"Price":12.09}
	str := string(json_data)

	//将json字符串解析到指定结构体中
	var res Stu
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Name, res.Age, res.Classes[0], res.Price)

	person := &Person{
		Name: "daheige313",
		Age:  28,
	}

	json_person, _ := json.Marshal(person)
	fmt.Println(string(json_person)) //{"name":"daheige313","age":28}
}
