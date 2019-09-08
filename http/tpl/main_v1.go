/**
 * @time: 2019-08-23 14:03
 * @author: louis
 */
package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	loginTpl  = "http/tpl/login.html"
	uploadTpl = "http/tpl/upload.html"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(ctime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, err := template.ParseFiles(loginTpl)
		if err != nil {
			panic(err)
		}
		log.Println(t.Execute(w, token)) //增加token认证
	} else {
		_ = r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("token", token)
		} else {

		}
		fmt.Println("username length", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte("Welcome "+r.Form.Get("username")+" login"))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(ctime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, err := template.ParseFiles(uploadTpl)
		if err != nil {
			panic(err)
		}
		log.Println(t.Execute(w, token)) //增加token认证
	} else {
		//application/x-www-form-urlencoded   表示在发送前编码所有字符（默认）
		//multipart/form-data	  不对字符编码。在使用包含文件上传控件的表单时，必须使用该值。
		//text/plain	  空格转换为 "+" 加号，但不对特殊字符编码。
		r.ParseMultipartForm(32 << 20)                 //将文件存入内存32M
		file, handler, err := r.FormFile("uploadfile") //使用formFile方法获取文件句柄
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		// A MIMEHeader represents a MIME-style header mapping
		// keys to sets of values.
		fmt.Fprintf(w, "%v", handler.Header)
		// map[Content-Disposition:[form-data; name="uploadfile"; filename="尚硅谷_韩顺平_Go语言核心编程.pdf"] Content-Type:[application/pdf]]
		f, err := os.OpenFile("d:/text/"+handler.Filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, _ = io.Copy(f, file) //将file句柄复制到f,生成文件
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! URL: "+r.URL.Path)
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}
