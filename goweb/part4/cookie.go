package main

import (
	"fmt"
	"net/http"
)

//将cookie发至浏览器

func setCookieDemo(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "heige",
		MaxAge:   3600,
		Value:    "123456",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "test",
		Value:    "ssfefelllll",
		HttpOnly: true, //浏览器无法改变cookie
	}
	c3 := http.Cookie{
		Name:     "test_c3",
		Value:    "345",
		HttpOnly: true, //浏览器无法改变cookie
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String()) //再加一个header头

	//通过http.Setcookie(w,&c1)方式来设置cookie 是go简单方法
	http.SetCookie(w, &c3)
	fmt.Fprintf(w, "test cookie")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("heige") //快速获取cookie value c1是一个*http.Cookie类型的指针
	if err == nil {
		fmt.Fprintln(w, "设置的cookie为", c1.Value)
	}

	all_cookies := r.Cookies() //是一个切片 [heige=123456 test=ssfefelllll test_c3=345]
	fmt.Fprintln(w, all_cookies)
	fmt.Fprintln(w, "第一个cookie heige值是:"+all_cookies[0].Value) //获取第一个值

}

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	http.HandleFunc("/test-cookie", setCookieDemo)
	http.HandleFunc("/getCookie", getCookie)
	server.ListenAndServe()
}
