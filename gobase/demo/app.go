package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := &Person{
		Id:   1,
		Name: "daheige",
		Age:  28,
	}

	//单条json数据返回
	json_data, err := json.Marshal(person)
	log.Println(string(json_data), err)

	//采用[]type   切片类型模拟多条数据返回
	perList := []*Person{
		person,
		&Person{
			Id:   2,
			Name: "XIAOMING",
			Age:  23,
		},
	}

	log.Println(perList[0])

	json_data2, err := json.Marshal(perList)
	log.Println(string(json_data2), err)

	//采用map[string]interface{}方式返回json数据
	var info = map[string]interface{}{
		"user": perList,
	}

	json_data3, err := json.Marshal(info)
	log.Println(info)
	log.Println(string(json_data3))
}
