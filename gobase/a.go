package main

import (
	"fmt"
)

//map传递到函数中，是引用传递，并非值的副本
func delColor(colrs map[string]string, key string) {
	delete(colrs, key)
}

//定义结果体
type User struct {
	username string
	email    string
	age      int
}

type Admin struct {
	person User
	level  int
}

type euser struct {
	name  string
	email string
}

//以值传递给接收者，是值的一个副本
func (u euser) notify() {
	fmt.Printf("send user email %s<%s>\n", u.name, u.email)
}

//使用指针作为方法的接收者
func (u *euser) changeEmail(email string) {
	u.email = email
}

func main() {
	colors := map[string]string{
		"alice": "#fff",
		"coral": "#ff7",
		"dar":   "#a99",
		"fore":  "#ffsss", //在map的末尾必须有,
	}

	for key, val := range colors {
		fmt.Printf("key is %s Value is %s\n", key, val) //每次打印key,val都不同
	}
	//删除key
	delete(colors, "fore")
	fmt.Println("colors is", colors)
	delColor(colors, "dar")
	fmt.Println("colors is", colors)

	stu := User{
		username: "heige",
		email:    "zhuwei313@hotmail.com",
		age:      27, //逗号必须有
	}
	fmt.Println("username:"+stu.username, "email: "+stu.email)

	user_1 := Admin{
		person: User{
			username: "zhuwei",
			age:      23,
			email:    "111@qq.com",
		},
		level: 1,
	}

	fmt.Println("user1 info\nusername:"+user_1.person.username,
		"\nage:", user_1.person.age, "level:", user_1.level)

	bill := euser{
		name:  "heige",
		email: "111@qq.com",
	}
	bill.notify()
	bill.changeEmail("zhuwei@com")
	bill.notify()

}
