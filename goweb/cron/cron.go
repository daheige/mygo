package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()

	// spec := "*/3 * * * *" //每3s执行一次
	spec := "0 */1 * * * *" //每3s执行一次
	// spec := "2,10 9-12 * * *" // 每天上午9点到12点的第2和第10分钟执行
	c.AddFunc(spec, func() {
		i++
		log.Println("exec task: ", i)
	})
	c.Start() //开始计划任务
	select {}
}
