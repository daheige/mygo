package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id   int
	Name string
	Job  string
}

//用于记录写入csv文件成功和失败的id
type HandlerRes struct {
	success_ids []int
	error_ids   []int
}

func main() {
	allPosts := []Post{
		{
			Id:   1,
			Name: "heige",
			Job:  "php",
		},
		{
			Id:   2,
			Name: "maogege",
			Job:  "java",
		},
		{
			Id:   3,
			Name: "xiaoming",
			Job:  "go",
		},
		{
			Id:   4,
			Name: "ssssxiaoming",
			Job:  "go-go-go",
		},
	}
	//创建csv文件
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile) //创建一个写入实例
	writer.Write([]string{"Id", "Name", "Job"})
	res := HandlerRes{}
	for _, post := range allPosts {
		//Id需要转换为字符串
		line := []string{strconv.Itoa(post.Id), post.Name, post.Job} //转换为字符串切片
		fmt.Println(line)
		if err := writer.Write(line); err != nil {
			res.error_ids = append(res.error_ids, post.Id)
			panic(err)
		}
		res.success_ids = append(res.success_ids, post.Id)
	}
	writer.Flush() //必须调用，保证缓冲区中的所有数据正确的写入文件中

	if len(res.error_ids) > 0 {
		fmt.Println("写入失败的id", res.error_ids)
		return
	}

	fmt.Println("写入成功的id", res.success_ids)
}
