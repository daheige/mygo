package main

import (
	"fmt"

	"github.com/wulijun/go-php-serialize/phpserialize"
)

func main() {
	//serialize序列化
	data2 := make(map[interface{}]interface{})
	data2["x"] = "fefefe"
	res, err := phpserialize.Encode(data2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
	fmt.Printf("%T\n", res) //string

	fmt.Println(phpserialize.Encode("fefeef"))

	//反序列化
	// fmt.Println(phpserialize.Decode(res))

	decodeRes, err := phpserialize.Decode(res)

	if err != nil {
		fmt.Printf("decode data fail %v, %v", err, res)
		return
	}

	decodeData, ok := decodeRes.(map[interface{}]interface{})
	if !ok {
		fmt.Println("unserialize error")
	}

	fmt.Println(decodeData)

}
