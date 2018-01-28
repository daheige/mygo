package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id   int
	Name string
	Job  string
}

func storeData(data interface{}, fileName string) {
	buf := new(bytes.Buffer) //创建写入缓冲区
	//创建gob编码器
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(data) //将data数据编码到缓冲区
	if err != nil {
		panic(err)
	}

	//将缓冲区中已编码的数据写入文件中
	err = ioutil.WriteFile(fileName, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

//将gob写入的内容，载入到data中
func loadData(data interface{}, fileName string) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	//根据这些原始数据，创建缓冲区
	buf := bytes.NewBuffer(raw)
	//将数据解码到缓冲区 (为缓冲区创建解码器)
	dec := gob.NewDecoder(buf)
	//解码数据到data中
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("gob文件读写")
	post := Post{
		Id:   1,
		Name: "heige",
		Job:  "goer",
	}
	storeData(post, "post_gob.md")

	//从文件中载入gob写入的文件内容到post
	var postData Post

	loadData(&postData, "post_gob.md") //第一个参数接受postData的内存地址，因为loadData载入的数据会存入data中
	fmt.Println(postData)
	fmt.Println(postData.Id, postData.Name)

	//实现字符串的存取
	storeData("fefefe", "test_gob.md")
	var str string
	loadData(&str, "test_gob.md")
	fmt.Println(str)
}
