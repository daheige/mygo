package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test") //hash实现了io.writer
	b := []byte{}
	fmt.Printf("res is %x\n", hasher.Sum(b)) //加密后的字符串
	fmt.Printf("res is %d\n", hasher.Sum(b))

	hasher.Reset() //重置
	data := []byte("123456")
	if n, err := hasher.Write(data); n != len(data) || err != nil {
		log.Printf("hash write err is %s", err)
	}

	checksum := hasher.Sum(b)
	fmt.Printf("res is %x\n", checksum)

	hg_hasher := md5.New()
	io.WriteString(hg_hasher, "123456") //将字符串123456写入
	fmt.Printf("res is %x\n", hg_hasher.Sum(b))
	x := fmt.Sprintf("%x", hg_hasher.Sum(b)) //形成16进制的32位字符串
	fmt.Println(x)

}
