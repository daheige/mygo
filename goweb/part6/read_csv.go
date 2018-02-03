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

func main() {
	csvFile, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1       //字段数量和记录中的字段数量，不一致时，读取进程不会中断
	allPosts, err := reader.ReadAll() //读取结果到一个切片中
	if err != nil {
		panic(err)
	}

	var posts []Post
	//读取后遍历
	for key, value := range allPosts {
		if key == 0 {
			continue
		}

		id, _ := strconv.ParseInt(value[0], 0, 0)
		line := Post{Id: int(id), Name: value[1], Job: value[2]}
		fmt.Println("读取的数据是", line)
		posts = append(posts, line)
	}

	fmt.Println(posts)

}
