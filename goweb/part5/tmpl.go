package main

import (
	"fmt"
	"html/template"
	"math/rand" //随机数
	"net/http"
	"time" //时间处理包
)

// {{ if arg }}{{ else }}{{ end }}条件动作
// {{ range .}}{{ . }}{{ end }} 迭代数组，切片，map
// {{ range $key,$val := .}}{{ $key }}{{ $val }}{{ else }} default value {{ end }} 迭代动作中设置变量
func parseHtml(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html") //解析html模板
	if err != nil {
		fmt.Fprintln(w, "error html")
	}
	t.Execute(w, "daheige this is test!") //第二个参数将结果传递到模板中
}
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("process.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(100) > 20)
}

func rangeData(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("range.html")
	t.Execute(w, []string{"fefe", "php", "js", "go"}) //传递一个字符串切片到模板中
	// t.Execute(w, nil) //传递一个字符串切片到模板中
}

//{{ template "t2.html" .}} 模板t1包含模板t2
func parseDemo2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t1.html", "t2.html")
	t.Execute(w, "fefesss")
}

func getDate(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"fdate": diyDate,
	}
	//创建一个模板
	t := template.New("diyfunc.html").Funcs(funcMap)
	t, _ = t.ParseFiles("diyfunc.html")
	t.Execute(w, time.Now())
}

//自定义函数
func diyDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

//模板内容可以上下文感知 可以转义html,css,js 防止xss攻击
func showHtml(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl_show.html")
	// content := `this is test <i>fefe</i><script>alert(1);</script>`
	// t.Execute(w, content)
	t.Execute(w, r.FormValue("content")) //http://127.0.0.1:8081/showHtml?content=%3Cscript%3Ealert(1)%3C/script%3E
	// t.Execute(w, template.HTML(r.FormValue("content"))) //template.HTML方法不会对html进行转义
}

//模板嵌套
func testDemo(w http.ResponseWriter, r *http.Request) {
	//ParseFiles默认会用第一个模板的名称作为语法分析的模板名称
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(10)
	content_name := "demo_content.html"
	if rnd > 5 {
		content_name = "test_content.html"
	}
	t, err := template.ParseFiles("layout/base.html", content_name) //必须在demo_content.html中用define显示指定模板名称content
	fmt.Println(err)
	t.ExecuteTemplate(w, "base", "hello,heige") //显示执行模板解析
}

//采用block块动作定义默认布局文件
func testDemoLayout(w http.ResponseWriter, r *http.Request) {
	//ParseFiles默认会用第一个模板的名称作为语法分析的模板名称
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(10)
	var t *template.Template //t是一个模板类型的指针
	var err error
	if rnd > 5 {
		t, err = template.ParseFiles("layout/index.html", "test_content.html")
	} else {
		t, err = template.ParseFiles("layout/index.html") //采用默认的模板
	}

	if err != nil {
		fmt.Fprintln(w, "error layout!")
		return
	}

	t.ExecuteTemplate(w, "index", "hello,heige") //显示执行模板解析
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}

	http.HandleFunc("/tmpl", parseHtml)
	http.HandleFunc("/rnd", process)
	http.HandleFunc("/range", rangeData)
	http.HandleFunc("/demo", parseDemo2)
	http.HandleFunc("/getDate", getDate)
	http.HandleFunc("/showHtml", showHtml)
	http.HandleFunc("/testDemo", testDemo)
	http.HandleFunc("/test", testDemoLayout)

	server.ListenAndServe()
}
