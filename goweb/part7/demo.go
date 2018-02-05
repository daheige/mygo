package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().String()
	fmt.Println(t)
	cur_time := time.Now().Unix()
	fmt.Println(cur_time)
	fmt.Println(time.Unix(cur_time, 0)) //根据时间戳返回当前本地时间
	t1 := time.Now()
	fmt.Println(t1.UTC())   //utc时间表示
	fmt.Println(t1.Local()) //RPC时间
	//格式化时间
	fmt.Println(t1.Format(time.RFC3339))
	fmt.Println(t1.Year()) //year
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())        //5
	fmt.Println(t1.Clock())      //返回hour,min.sec
	fmt.Println(t1.Hour())       //返回当前小时
	fmt.Println(t1.Minute())     //分
	fmt.Println(t1.Second())     //秒
	fmt.Println(t1.Nanosecond()) //纳秒
	fmt.Println(t1.Location())   //时间的时区
	fmt.Println(t1.Unix())       //当前时间戳
	fmt.Println(t1.UnixNano())   //当前时间转换为纳秒数
	year, month, day := t1.Date()
	fmt.Println(year, month, day)
	h, m, s := t1.Clock() //时分秒
	fmt.Println(h, m, s)
}
