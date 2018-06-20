package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func countNum(index int, stmt *sql.Stmt, done chan<- bool) {
	myNum := 0
	for num := (100 * (index - 1)); num < (100 * index); num++ {
		myNum = num + 1
		// fmt.Println(myNum)
		stmt.Exec(myNum)
	}

	done <- true //执行完毕后,就将bool发送到通道中
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/laravel")
	checkErr(err)

	fmt.Println(db)
	// 插入数据
	stmt, err := db.Prepare("insert into performance (num) values (?)")
	fmt.Println(reflect.TypeOf(stmt))

	// 计算程序运行时间
	start := time.Now()
	done := make(chan bool) //采用done无缓冲通道
	// 开启10个协程
	for i := 1; i <= 10; i++ {
		go countNum(i, stmt, done)
	}

	//当通道已经有值,就立即取出来(在未将值放入通道中,这里一直会阻塞,等待有值就取出来)
	// channel在主线程中被赋值后，主线程就会阻塞，直到channel的值在子线程中被取出
	for i := 1; i <= 10; i++ {
		<-done
	}

	elapsed := time.Since(start)

	fmt.Println("执行时间: ", elapsed)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
