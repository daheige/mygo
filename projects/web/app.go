package main

import (
    "fmt"
    "projects/web/jwt"
    "sync"
)

var str string
var once sync.Once

//如果没有多个s执行不同的初始化,就可以采用唯一单例模式
//once是一个结构体,内部上面有一个Done字段,执行一次,后面就不会执行
//atomic.StoreUint32 原子性操作
func initStr(s string) string {
    once.Do(func() {
        fmt.Println(111)
        str = s
    })

    return str
}

//允许改变配置的单例模式
//通过这样的模式,可以解除唯一单例模式
type Single struct {
    once   sync.Once
    config string
}

func (this *Single) initConfig(config string) {
    this.once.Do(func() {
        fmt.Println("init once config:", config)
        this.config = config
    })
}

func main() {
    j := &jwt.Jwt{}
    jwt.Sign = "ssf33"
    j.Test("sss")
    fmt.Println("fefe")

    //调用了函数后,内部的Once上的Done执行了一次加1,后面的不会再执行
    //所以s2的值是sss
    s := initStr("sss")
    s2 := initStr("heige")
    fmt.Println(s, s2)

    obj := &Single{}
    obj.initConfig("daheige")
    obj.initConfig("sss") //这里不会发生改变
    fmt.Println(obj.config)

    obj2 := &Single{}
    obj2.initConfig("ssss123456")
    obj2.initConfig("ssss123456") //这里不会执行fmt.Println
    fmt.Println(obj2.config)
}
