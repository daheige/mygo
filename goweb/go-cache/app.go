package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	//设置默认5分钟过期，每隔10分钟，做一次清理工作
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("foo", "mytest", cache.DefaultExpiration)

	//get data
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	fmt.Println(found)

}
