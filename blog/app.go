package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayHell(w http.ResponseWriter, req *http.Request) {
	req.ParseForm() //parse args
	fmt.Println(req.Form)
	fmt.Println("path", req.URL.Path)
	fmt.Fprintf(w, "heige")
}

func login(w http.ResponseWriter, req *http.Request) {
	var method string = strings.ToLower(req.Method)
	fmt.Println("method", method)
	if method == "get" {
		cur_time := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cur_time, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println(token)
		var str string = `<!DOCTYPE html>
                                <html>
                                <head>
                                    <meta charset="utf-8">
                                    <meta http-equiv="X-UA-Compatible" content="IE=edge">
                                    <title>login</title>
                                </head>
                                <body>
                                    <form action="/login" method="post" accept-charset="utf-8">
                                        username : <input type="text" name="username"/>
                                        pwd : <input type="password" name="password"/>
                                        age : <input type="text" name="age"/>
                                        sex : F<input type="radio" name="sex" value="1"/>
                                        M<input type="radio" name="sex" value="2"/>` +
			"<input type='hidden' name='token' value='" + token + "'" + "/>" +
			`<input type="submit" value="login"/>
                                    </form>
                                    <form action="/upload" method="post" enctype="multipart/form-data">
                                        <input type="file" name="upload_file"/>
                                        <input type="submit" value="upload"/>
                                    </form>
                                </body>
                         </html>`
		fmt.Fprint(w, str)
	} else {
		//request data is login data
		req.ParseForm() //解析提交的参数，get/post都可以，只执行一次就可以
		if len(strings.TrimSpace(req.Form["username"][0])) == 0 {
			fmt.Fprintf(w, "username is empty")
			return
		}
		if m, _ := regexp.MatchString("^[a-zA-z]+$", req.Form.Get("username")); !m {
			fmt.Fprint(w, "username must be an english str")
			return
		}
		//regexp match age
		if m, _ := regexp.MatchString("^[0-9]+$", req.Form.Get("age")); !m {
			fmt.Fprint(w, "please age in number")
			return
		}
		age, err := strconv.Atoi(req.Form.Get("age")) //conv str to int
		if err != nil {
			fmt.Fprint(w, "age must be a number")
			return
		}

		if age > 100 {
			fmt.Fprint(w, "age > 100 !!")
			return
		}
		sex, serr := strconv.Atoi(req.Form.Get("sex"))
		if serr != nil {
			fmt.Fprint(w, "sex error")
			return
		}

		//check sex
		sexArr := []int{1, 2}
		var flag bool = false
		for _, item := range sexArr {
			if item == sex {
				flag = true
			}
		}
		if !flag {
			fmt.Fprint(w, "sex must be m or f")
			return
		}
		fmt.Println("username is ", strings.Join(req.Form["username"], ""))
		fmt.Println("username", string(req.PostFormValue("username")))
		fmt.Println("password", string(req.PostFormValue("password")))
		fmt.Fprintf(w, "username is "+string(req.PostFormValue("username"))+string(req.PostFormValue("password")))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 把上传的文件存储在内存和临时文件中
		r.ParseMultipartForm(32 << 20)                  //upload file storage in mem
		file, handler, err := r.FormFile("upload_file") //获取文件句柄，然后对文件进行存储等处理
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "upload err")
			return
		}

		defer file.Close()
		// fmt.Fprintf(w, "%v", handler.Header) //file mime
		//创建文件存放的目录
		var storage_path string = "./test/"
		if _, err := os.Stat(storage_path); err != nil {
			os.Mkdir(storage_path, 0755)
		}
		// fmt.Println(handler.Filename)

		cur_time := time.Now().Format("2006-01-02 15:04:05")
		//md5保存文件名
		var filename string = storage_path + md5_str(cur_time+strconv.Itoa(getRndNum(1, 1000))) + path.Ext(handler.Filename)
		// f, err := os.OpenFile(storage_path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666) //存放在服务器上的文件名
		if err != nil {
			fmt.Println(err)
			fmt.Fprint(w, "upload error")
			return
		}
		defer f.Close()
		io.Copy(f, file) //将临时文件复制到目标文件
		fmt.Fprintf(w, "%s", "upload success,filename is "+filename)
		return
	} else {
		fmt.Fprint(w, "error!")
	}
}

func md5_str(text string) string {
	cx := md5.New()
	cx.Write([]byte(text))
	return hex.EncodeToString(cx.Sum(nil))
}
func getRndNum(min, max int) (randNum int) {
	rand.Seed(time.Now().Unix()) //给一个时间戳的种子，否则每次生成都是一样的值
	randNum = rand.Intn(max-min) + min
	return
}
func foo(w http.ResponseWriter, r *http.Request) {
	targetUrl := "http://localhost:8080/upload"
	filename := "./test/263e73650c2adad32726800e76322a3e.jpg"
	postFile(filename, targetUrl)
	fmt.Println("ok")
}
func main() {

	http.HandleFunc("/", sayHell)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/foo", foo)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
