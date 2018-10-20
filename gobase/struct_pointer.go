package main

import (
	"log"
)

type userInfo struct {
	id   int
	name string
}

//返回实例地址，userInfo类型的指针
func getInfo() (user *userInfo) {
	user = &userInfo{} //是需要实例化
	user.id = 1
	user.name = "daheige"
	return
}

//返回的是一个实例的值，并非地址
//在函数里面，不需要实例化userInfo
func getInfo2() (user userInfo) {
	user.id = 1
	user.name = "daheige"
	return
}

func main() {
	log.Println(getInfo())
	log.Println(getInfo2())
}

/**
2018/10/20 11:19:44 &{1 daheige}
2018/10/20 11:19:44 {1 daheige}
*/
