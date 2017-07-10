package main

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_postfile(t *testing.T) {
	str := "i am learing go lang,hello world!"
	re, _ := regexp.Compile("[a-z]{2,4}")
	str_b := []byte(str)
	//first regexp
	one := re.Find(str_b)
	fmt.Println("find", string(one))

	all := re.FindAll(str_b, -1)           //find all match str
	fmt.Println("find all ", all)          //生成多个字节类型的切片
	allIndex := re.FindAllIndex(str_b, -1) //小于0表示返回全部符合的字符串
	for k, v := range allIndex {
		fmt.Println(k, v)
	}

	re2, _ := regexp.Compile("am(.*)lang(.*)")
	all_str := re2.FindSubmatch(str_b) //查找Submatch,返回数组
	fmt.Println(all_str)
	for k, v := range all_str {
		fmt.Println(k, string(v))
	}

	//FindAllSubmatch,查找所有符合条件的子匹配
	subAll := re2.FindAllSubmatch(str_b, -1)
	// fmt.Println(subAll)
	fmt.Println("sub all\n")
	for _, v := range subAll {
		for _k, _v := range v {
			fmt.Println(_k, string(_v))
		}
	}
	t.Log("ok")
}
