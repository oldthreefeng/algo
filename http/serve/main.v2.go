/*
@Time : 2019/8/22 15:53
@Author : louis
@File : main.v2
@Software: GoLand
*/

package main

import (
	"log"
	"net/http"
)

type myHandler struct { //实现Handler这个接口,

}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello my handler v2,the URL is : " + r.URL.String()))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{}) //将自己的实现的Handler添加到路由
	mux.HandleFunc("/bye", sayByeV2)

	log.Println("start serve..v2")
	log.Fatal(http.ListenAndServe(":4001", mux)) //将handler传入listener
}
func sayByeV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye v2..." + r.URL.String()))
}
