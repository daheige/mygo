package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//创建临时目录,并写入文件
	content := []byte("dddd")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "file")
	if err := ioutil.WriteFile(tmpfn, content, 0644); err != nil {
		log.Fatal(err)
	}
	log.Println("write success")
}
