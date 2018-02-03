package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入mysql驱动
	"strconv"
	"time"
)

var db *sql.DB

type UserInfo struct {
	id   int
	name string
}

func init() {
	var err error
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	db, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/test?charset=utf8")
	checkError(err)
}

func main() {
	//插入数据
	// stmt, err := db.Prepare("insert into test (name) values(?)") //Prepare预处理
	// checkError(err)
	// res, err := stmt.Exec("heige")
	// checkError(err)
	// insert_id, err := res.LastInsertId()
	// fmt.Println("插入的id", insert_id)
	// //执行执行sql
	// rs, err := db.Exec("insert into test (name) values('daheige313')")
	// checkError(err)
	// ins_id, _ := rs.LastInsertId()
	// fmt.Println("插入的id", ins_id)

	getData()

	//更新数据可以用db.Exec来处理,受到影响的函数获取,使用RowsAffected方法
	ret3, _ := db.Exec("update test set name ='haha3' where id = ?", 12)
	aff_rows, _ := ret3.RowsAffected()
	fmt.Println("影响的行数是", aff_rows)
	//删除数据也可用db.Exec来处理，返回处理受到影响的行数RowsAffected

	//事务处理
	t := time.Now().Unix()
	rnd_str := strconv.Itoa(int(t)) //将时间戳转换为int类型,最后转换为字符串
	tx, _ := db.Begin()
	r4, _ := tx.Exec("update test set name = 'heige3333"+rnd_str+"' where id = ?", 13)
	r5, _ := tx.Exec("update test set name = 'heige3314"+rnd_str+"' where id = ?", 14)
	up_num1, _ := r4.RowsAffected()
	up_num2, _ := r5.RowsAffected()
	if up_num1 > 0 && up_num2 > 0 {
		tx.Commit() //提交事务
		fmt.Println("执行成功")
	} else {
		tx.Rollback()
		fmt.Println("执行失败！")
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getData() {
	rows, err := db.Query("select id,name from test where id = ?", 12)
	checkError(err)
	//Next()是用来迭代数据的，无论是一行，还是多行，执行完毕后，必须调用
	for rows.Next() { //row.Next() 当所有的行迭代完毕后，返回io.EOF
		var name string
		var id int
		//Scan将数据解析到id,name,Scan映射指定字段
		if err := rows.Scan(&id, &name); err != nil {
			panic(err)
		}
		fmt.Printf("%s is %d\n", name, id)
	}
	rows.Close()

	//查询所有字段
	res, err := db.Query("select * from test")
	checkError(err)
	//获取查询后的字段
	cols, _ := res.Columns()
	fmt.Println(cols)                       //cols是一个[]byte [id name]
	vals := make([][]byte, len(cols))       //这里一行所有列的值，放在[][]byte中
	scans := make([]interface{}, len(cols)) //这里一行填充的数据
	//scans引用vals,把数据填充到[]byte中
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)
	for res.Next() {
		res.Scan(scans...) //将数据映射到scans中,当scans发生了改变，vals也相应发生改变
		r := make(map[string]string)
		//把vals中的数据复制到r
		for k, v := range vals {
			key := cols[k]
			r[key] = string(v)
		}
		result[i] = r
		// fmt.Println(r)
		i++
	}
	fmt.Println("第一行数据是", result[0]["id"], result[0]["name"])

	//查询一条记录
	row := db.QueryRow("select id,name from test limit ?", 1)
	user := UserInfo{} //user的结构是一个结构体的一个实例
	row.Scan(&user.id, &user.name)
	fmt.Println(user.id, user.name)

}
