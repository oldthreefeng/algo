/*
@Time : 2019/9/19 10:14
@Author : louis
@File : helloclient
@Software: GoLand
*/

package client

import (
	"gogs.wangke.co/go/algo/rpc/server"
	"net/rpc"
)

//HelloServiceClient
type HelloServiceClient struct {
	*rpc.Client
}

//ensure HelloServiceClient implement HelloServiceInterface
var _ server.HelloServiceInterface = (*HelloServiceClient)(nil)

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(server.HelloServiceName+".Hello", request, reply)
}

//DialHelloService
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

