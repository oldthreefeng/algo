/*
@Time : 2019/9/19 10:00
@Author : louis
@File : helloservice
@Software: GoLand
*/

package server

import "net/rpc"

//HelloServiceName name, path is not likely to golang
// 1.declare Name and interface
const HelloServiceName = "test/HelloService"

//HelloServiceInterface method declare & assignment
//2. list all interface method
type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

//RegisterHelloService is register service rpc
//3. register the interface service
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

