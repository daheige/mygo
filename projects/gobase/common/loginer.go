package common

//导出还是未导出，是通过名称首字母的大小写决定的
//它们决定了是否可以访问，也就是标志符的可见性
import (
	"fmt"
)

type Loginer interface {
	Login()
}

type defaultLogin int

func (d defaultLogin) Login() {
	fmt.Println("login in....")
}

//对外提供了Loginer接口的Login
func NewLoginer() Loginer {
	return defaultLogin(0)
}
