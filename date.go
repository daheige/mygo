package main

// 时间相关的操作，加减操作

import (
	"log"
	"math"
	"time"
)

var tmFormat = "2006-01-02 15:04:05"

func main() {
	t := time.Now()
	d := t.AddDate(0, 0, 3) // 往后3天

	log.Println(d.Format("2006-01-02"))

	d2 := t.AddDate(0, 0, -3) //往前3天
	str := d2.Format("2006-01-02")
	log.Println("str: ", str)

	m, _ := time.ParseDuration("-1m")

	t3 := t.Add(10 * m) //10min之前

	log.Println("10min before: ", t3.Format("2006-01-02 15:04:05"))

	h, _ := time.ParseDuration("-1h")
	log.Println(t.Add(1 * h).Format(tmFormat)) //1小时之前

	now := time.Now()
	day, _ := time.ParseDuration("-24h")
	d1 := now.Add(1 * day) //一天前
	log.Println("one day before: ", d1.Format(tmFormat))

	log.Println("sub hours: ", math.Floor(t.Sub(d1).Hours())) //相隔多个小时
}

//对于月份操作
/**
2018-05-30 加一个月变成了2018-07-01

AddDate会将结果规范化，类似Date函数的做法。因此，举个例子，给时间点October 31添加一个月，会生成时间点December 1。（从时间点November 31规范化而来）

再看一下源码
image.png
所以当你给month加 1，day 是不会变的。5-31变成 6-31，最后转化为 7-1。
所以大家在用任何官方、非官方的接口，都一定要仔细阅读接口文档呀，不然很容易出问题
*/

/**
2019/10/29 23:52:37 2019-11-01
2019/10/29 23:52:37 str:  2019-10-26
2019/10/29 23:52:37 10min before:  2019-10-29 23:42:37
2019/10/29 23:52:37 2019-10-29 22:52:37
2019/10/29 23:52:37 one day before:  2019-10-28 23:52:37
2019/10/29 23:52:37 sub hours:  23
*/
