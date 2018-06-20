package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func countNum(index int, stmt *sql.Stmt, wg *sync.WaitGroup) {
	defer wg.Done()
	myNum := 0
	for num := (100 * (index - 1)); num < (100 * index); num++ {
		myNum = num + 1
		// fmt.Println(myNum)
		stmt.Exec(myNum)
	}

}

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
	// 开启10个协程
	var wg sync.WaitGroup //信号计数器,保证协程执行完毕
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		//这里必须传递wg的指针,在协程内部执行才可以改变wg计数器
		//否则wg传递过去的是一个副本,无法保证协程执行
		go countNum(i, stmt, &wg)
	}
	wg.Wait() //等待协程执行完毕
	elapsed := time.Since(start)
	fmt.Println("执行时间: ", elapsed)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
