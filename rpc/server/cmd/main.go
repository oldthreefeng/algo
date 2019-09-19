/*
@Time : 2019/9/19 10:27
@Author : louis
@File : main
@Software: GoLand
*/

package main

import (
	"gogs.wangke.co/go/algo/rpc/server"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "HelloService:" + request
	return nil
}

func main() {

	//rpc.RegisterName("HelloService", new(HelloService))
	//
	//listener, err := net.Listen("tcp", ":1234")
	//if err != nil {
	//	log.Fatal("ListenTCP error:", err)
	//}
	//
	//conn, err := listener.Accept()
	//if err != nil {
	//	log.Fatal("Accept error:", err)
	//}
	//rpc.ServeConn(conn)

	_ = server.RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("net.Listen:", err)
	}
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("listener.Accept:", err)
		}

		go rpc.ServeConn(conn)
	}

}
