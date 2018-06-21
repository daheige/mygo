package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("i am heige")
}

func taskWithParams(a, b int) {
	fmt.Println(a, b)
}
func job() {
	fmt.Println(222)
}

func main() {
	gocron.Every(1).Second().Do(task)
	gocron.Every(1).Second().Do(taskWithParams, 1, 2)
	gocron.Every(1).Minute().Do(job)
	gocron.Every(2).Seconds().Do(func() {
		fmt.Println("每3s执行一次")
	})

	gocron.Remove(task)               //去除某个计划任务
	gocron.Clear()                    //清除所有的计划任务
	gocron.Every(1).Second().Do(task) //重新开始task

	<-gocron.Start()

	//开启另一个计划任务
	// s := gocron.NewScheduler()
	// s.Every(3).Seconds().Do(func() {
	// 	fmt.Println("haha")
	// })
	// <-s.Start()
}
