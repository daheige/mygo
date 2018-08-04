package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
)

//想要用go写接口,api加密形式在php中是hash_hmac('sha1',$string,$key);
func main() {
	//sha1
	h := sha1.New()
	io.WriteString(h, "aaaaaa")
	fmt.Printf("%x\n", h.Sum(nil))

	//hmac ,use sha1
	key := []byte("123456")
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", mac.Sum(nil))
}
