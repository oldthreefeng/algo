/*
@Time : 2019/9/19 10:06
@Author : louis
@File : main
@Software: GoLand
*/

package main

import (
	"github.com/oldthreefeng/algo/rpc/client"
	"log"
)

func main() {
	//client, err := rpc.Dial("tcp", "localhost:1234")
	//if err != nil {
	//	log.Fatal("rpc.Dial:", err)
	//}
	//
	//var reply string
	//err = client.Call(server.HelloServiceName+".Hello", "hello", &reply)
	//if err != nil {
	//	log.Fatal("client.Call:", err)
	//}

	c, err := client.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("DialHelloService:", err)
	}
	var reply string
	err = c.Hello("hello", &reply)
	if err != nil {
		log.Fatal("client.Call:", err)
	}
	log.Println(reply)

}
