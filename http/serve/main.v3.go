/*
@Time : 2019/8/22 16:03
@Author : louis
@File : main.v3
@Software: GoLand
*/

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var m map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	server := http.Server{
		Addr:         ":4002",
		WriteTimeout: 5 * time.Second,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	mux := http.NewServeMux()
	m = make(map[string]func(w http.ResponseWriter, r *http.Request))
	m["/bye"] = sayByeV3
	mux.Handle("/", &myHandlerV3{}) //将自己的实现的Handler添加到路由
	//mux.HandleFunc("/bye",sayByeV3)
	go func() {
		<-quit //阻塞再这里
		if err := server.Close(); err != nil {
			log.Fatal("close server:", err)
		}
	}()

	server.Handler = mux
	log.Println("start serve v3...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("server listen err:", err)
	}

}

func sayByeV3(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(3 * time.Second)
	//w.Write([]byte("bye bye v3..."))
	io.WriteString(w, "bye bye v3...")
}

type myHandlerV3 struct { //实现Handler这个接口,

}

func (mh *myHandlerV3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m[r.URL.String()]; ok { //能匹配到map里面的url,则调用map里面对应的函数
		h(w, r)
		return //成功调用就返回.不用调用本身函数
	}
	io.WriteString(w, "hello my handler v3,the request  URL is :"+r.URL.String())
	//w.Write([]byte("hello my handler v3,the request  URL is : " + r.URL.String()))
}
