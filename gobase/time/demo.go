package main

import (
	"fmt"
	"time"
)

const tmMissMs = "2006-01-02 15:04:05"

func main() {
	//对于time.time 我们可以通过使用函数Before，After，Equal来比较两个time.Time时间
	t1 := time.Now()
	t2 := t1.Add(-1 * time.Minute) //t1是当前时间，t2是当前时间的前一分钟
	fmt.Println(t1.Before(t2))     //false

	fmt.Println(t1 == t2)    //false
	fmt.Println(t1.IsZero()) //false

	// 常常会将time.Time格式化为字符串形式：
	loc, _ := time.LoadLocation("PRC")
	t := time.Now().In(loc)

	str_t := t.Format(tmMissMs)

	fmt.Println(str_t) //2018-06-04 23:06:07

	//parset str to time.time
	t3, _ := time.Parse(tmMissMs, str_t) //time.Parse 解析出来的时区却是 time.UTC

	fmt.Println("time", t3)

	//一般的，我们应该总是使用 time.ParseInLocation 来解析时间
	// 并给第三个参数传递 time.Local

	// t4, _ := time.ParseInLocation(tmMissMs, str_t, loc)
	t4, _ := time.ParseInLocation(tmMissMs, str_t, time.Local)

	fmt.Println(t4)
}
