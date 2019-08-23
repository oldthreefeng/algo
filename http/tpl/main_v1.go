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

const tplName  = "d:/text/login.html"

func login(w http.ResponseWriter,r *http.Request)  {
	if r.Method == "GET" {
		ctime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(ctime, 10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		t,err:= template.ParseFiles(tplName)
		if err != nil {
			panic(err)
		}
		log.Println(t.Execute(w,token)) //增加token认证
	} else {
		_ = r.ParseForm()
		token:=r.Form.Get("token")
		if token != "" {
			fmt.Println(token)
		} else {

		}
		fmt.Println("username length",len(r.Form["username"][0]))
		fmt.Println("username:",template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("username:",template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter,r *http.Request)  {
	if r.Method == "GET" {
		ctime := time.Now().Unix()
		h := md5.New()
		_, _ = io.WriteString(h, strconv.FormatInt(ctime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, err := template.ParseFiles("d:/text/upload.html")
		if err != nil {
			panic(err)
		}
		log.Println(t.Execute(w, token)) //增加token认证
	} else {
		r.ParseMultipartForm(32<<20)  //将文件存入内存32M
		file, hanlder,err := r.FormFile("uploadfile") //使用formFile方法获取文件句柄
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w,"%v",hanlder.Header)
		f ,err := os.OpenFile("d:/text/"+hanlder.Filename,os.O_CREATE|os.O_WRONLY,os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, _ = io.Copy(f, file) //将file句柄复制到f,生成文件
	}
}

func helloWorld(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"Hello World! URL: "+ r.URL.Path)
}

func main() {
	http.HandleFunc("/login",login)
	http.HandleFunc("/upload",upload)
	http.HandleFunc("/",helloWorld)
	err := http.ListenAndServe(":9999",nil)
	if err != nil {
		panic(err)
	}
}