package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//采用ioutil工具进行文件的写入和读取
	fmt.Println("开始写入文件")
	data := []byte("heige,golang")
	if err := ioutil.WriteFile("test.md", data, 0644); err != nil {
		panic(err)
	}

	read_data, _ := ioutil.ReadFile("test.md")
	fmt.Println("data is", string(read_data))

	//方式2 ： 采用os.write 更加灵活地读取和写入数据
	str := "fefsss"
	file1, _ := os.Create("demo.md")
	defer file1.Close()
	file1.WriteString(str)

	read2, _ := os.Open("demo.md")
	defer read2.Close()

	//读取数据到data2中
	data2 := make([]byte, len([]byte(str))-1)
	bytes, _ := read2.Read(data2)
	fmt.Println("读取的字节数是", bytes, string(data2))
	os.Exit(0)
}
