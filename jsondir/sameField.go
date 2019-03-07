package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	// Id   int `json:"id"`
	Id   int `json:"person_id"`
	Name string
}

type PersonConfig struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"`
	Job  string `json:"job,omitempty"` //这个值为空，就忽略它
	City string `json:"-"`             //忽略字段
}

//匿名结构体作为字段，如果存在同名的字段无法采用json导出，需要重新指定json tag
type Info struct {
	Person
	// Person Person
	PersonConfig
}

func main() {
	info := &Info{}

	p := Person{
		Id:   1,
		Name: "fefe",
	}

	personConfig := PersonConfig{
		Id:   12,
		Age:  123,
		City: "beijing",
	}

	log.Println(personConfig)

	info.Person = p
	info.PersonConfig = personConfig
	log.Println(info)
	jsonData, err := json.Marshal(info)

	log.Println(err, string(jsonData))
	//<nil> {"person_id":1,"Name":"fefe","id":12,"age":123}

	c := &Info{} //c必须是一个指针，因为会把解析结果保存在c中
	err = json.Unmarshal([]byte(`{"person_id":1,"Name":"fefe","id":12,"age":123}`), c)
	log.Println(c)
	log.Println(c.Person, c.Person.Id)

}
