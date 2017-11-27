package main

import (
	"fmt"
	hg "gobase/hgMap" //调用自定义的包并重命名包名
	"regexp"
	"strconv"
)

func main() {
	searchStr := "daheige 123.45 sss 345.43 shfe"
	pat := "\\d+.\\d+"

	if ok, _ := regexp.Match(pat, []byte(searchStr)); ok {
		fmt.Println("ok")
	}
	f, _ := strconv.ParseFloat("123.45", 32)
	fmt.Println(f)                                  //将字符串数字转化为浮点数
	fmt.Println(strconv.FormatFloat(f, 'f', 3, 32)) //格式化浮点数

	//将匹配的字符替换为***
	re, _ := regexp.Compile(pat)               //预处理正则
	s := re.ReplaceAllString(searchStr, "***") //替换所有的字符串
	fmt.Println(s)
	var bar = map[string]int{"x": 1, "y": 2}
	fmt.Println(hg.Fclip(bar))
}
