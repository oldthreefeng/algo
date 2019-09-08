/*
@Time : 2019/8/22 15:47
@Author : louis
@File : main.v1
@Software: GoLand
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world v1"))
	})
	http.HandleFunc("/bye", sayBye)
	log.Println("start serve..v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye... v1"))
}
