package main

import (
	"encoding/json"
	"fmt"
	"gobase/mystructs"
	"reflect"
	"unsafe"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	hgf := mystructs.GetFileInstance(1, "test.md") //工厂模式创建结构体
	fmt.Println(unsafe.Sizeof(hgf))                //占据8个字节内存
	fmt.Println(hgf)
	hgf.Fd = 12
	fmt.Println(hgf.Fd, hgf.Name)

	//判断结构体上属性的类型和访问属性标签
	fmt.Println(reflect.TypeOf(hgf.Fd))
	s := reflect.TypeOf(hgf).Elem() //通过反射获取type定义
	fmt.Println(s)
	fmt.Println(s.Field(0).Type)
	// for i := 0; i < s.NumField(); i++ {
	fmt.Println(s.Field(0).Tag) //将tag输出出来
	// }

	tt := reflect.TypeOf(hgf.Fd)
	fmt.Println(tt)

	//struct转换为json字符串
	hg_u := &User{Name: "heige", Age: 23}
	u_b, _ := json.Marshal(hg_u) //u_b是一个字节数组
	fmt.Println("转换后的字节数组是", u_b)
	fmt.Println("转换后的json是", string(u_b)) //{"name":"heige","age":23}

	//解析json字符串
	json_str := `{"name":"heige","age":23}`
	var data interface{}
	err := json.Unmarshal([]byte(json_str), &data)

	if err != nil {
		fmt.Println("json error")
	}

	// fmt.Println(data["name"]) //这样的方式错误的
	m := data.(map[string]interface{}) //通过断言访问

	fmt.Println(m["name"])

}
