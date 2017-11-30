package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (this *Page) save() {
	outFile, err := os.OpenFile(this.Title, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer outFile.Close()
	outFile.WriteString(string(this.Body))
	fmt.Println("write success")
}

func (this *Page) load(title string) {
	if this.Title != title {
		fmt.Printf("error: 文件%s不存在", title)
	}

	buf, err := ioutil.ReadFile(this.Title)
	if err != nil {
		fmt.Println("读取文件失败")
		return
	}

	fmt.Println(string(buf))

}

func main() {
	outFile, err := os.OpenFile("hg.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	// fmt.Fprintf(outFile, "fefesss\n") //简单的内容直接用Fprintf写入文件

	outWriter := bufio.NewWriter(outFile) //创建一个写入器（缓冲区）对象
	outWriter.WriteString("daheige")
	outWriter.Flush() //将缓冲区中的内容全部写入文件中

	//简单的内容直接写入
	fp, err := os.OpenFile("hgdemo.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fp.Close()
	fp.Write([]byte("zhuwei"))
	fp.WriteString("ddss")

	hgp := &Page{
		Title: "test.md",
		Body:  []byte("zhuwei,daheige,sss123"),
	}
	hgp.save()
	// hgp.load(hgp.Title)

}
