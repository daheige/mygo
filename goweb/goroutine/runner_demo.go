package main

import (
	"goweb/goroutine/common"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("======开始执行任务=======")
	timeout := 3 * time.Second     //任务超时为3s
	r := common.NewRunner(timeout) //创建一个runner
	r.Add(createTask(), createTask(), createTask())

	//开始执行任务
	if err := r.Start(); err != nil {
		switch err {
		case common.ErrInterrupt: //中断
			log.Println(err)
			os.Exit(2)
		case common.ErrorTimeout: //超时
			log.Println(err)
			os.Exit(1)
		}
	}
	log.Println("任务执行完毕")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

/**
 *运行结果
 *$ go run runner_demo.go
2018/03/27 22:21:31 ======开始执行任务=======
2018/03/27 22:21:31 正在执行任务0
2018/03/27 22:21:31 正在执行任务1
^C2018/03/27 22:21:32 执行者被中断
exit status 2
$ ^C 执行期间按下ctrl+c中断
$ go run runner_demo.go
2018/03/27 22:21:37 ======开始执行任务=======
2018/03/27 22:21:37 正在执行任务0
2018/03/27 22:21:37 正在执行任务1
^C^C2018/03/27 22:21:38 执行者被中断
exit status 2
$ ^C
$ go run runner_demo.go
2018/03/27 22:21:45 ======开始执行任务=======
2018/03/27 22:21:45 正在执行任务0
2018/03/27 22:21:45 正在执行任务1
2018/03/27 22:21:46 正在执行任务2
2018/03/27 22:21:48 执行任务超时
exit status 1
*/
//此外这个执行者也是一个很不错的模式，比如我们写好之后，交给定时任务去执行即可，比如cron，这个模式我们还可以扩展
//更高效率的并发，更多灵活的控制程序的生命周期，更高效的监控等
