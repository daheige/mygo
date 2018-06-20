package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/laravel")
	checkErr(err)
	defer db.Close()

	fmt.Println(db)
	// 插入数据
	stmt, err := db.Prepare("insert into performance (num) values (?)")
	fmt.Println(reflect.TypeOf(stmt))

	// 计算程序运行时间
	start := time.Now()
	done := make(chan int, 10) //采用done缓冲通道
	product(stmt, done)
	close(done) //通道写完后，必须关闭通道，否则range遍历会出现死锁
	cust(done)
	elapsed := time.Since(start)
	fmt.Println("执行时间: ", elapsed)

}

func product(stmt *sql.Stmt, done chan<- int) {
	var wg sync.WaitGroup //它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成
	for index := 1; index <= 10; index++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Printf("批量执行第%d组操作\n", i)
			defer wg.Done() //当前协程执行完毕,计数器减去1
			for num := (100 * (i - 1)); num < (100 * i); num++ {
				myNum := num + 1
				// fmt.Println(myNum)
				stmt.Exec(myNum)
			}
		}(index, &wg)

		done <- index //执行完毕后,往通道中流入数据
	}
	wg.Wait() //等待协程执行完毕

	fmt.Println("批量插入计划任务成功!")
}

func cust(done <-chan int) { //只读通道
	for ch := range done {
		fmt.Printf("第%d组数据执行完毕通知!\n", ch)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
//测试db table
create database laravel default charset utf8;
CREATE TABLE `performance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `num` int(11) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
*/
