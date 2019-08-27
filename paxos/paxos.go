/*
@Time : 2019/8/27 18:39
@Author : louis
@File : paxos
@Software: GoLand
*/

package paxos

import "time"

const (
	Prepare = iota + 1
	Propose
	Promise
	Accept
)

type promise interface {
	number() int
}

type accept interface {
	proposalValue() string
	proposalNumber() int
}

type mes struct {
	from, to int
	mesT     int
	n        int
	pren     int
	value    string
}

func (m *mes) proposalValue() string {
	//返回value需要什么条件呢？
	switch m.mesT {
	case Accept,Promise:
		return m.value
	default:
		panic("unexpect proposalValue")
	}

}

func (m *mes) proposalNumber() int {
	switch m.mesT {
	case Promise:
		return m.pren
	case Accept:
		return m.n
	default:
		panic("unexpect proposalNumber")
	}
}

func (m *mes) number() int {
	return m.n
}



type network interface {
	send(m mes)
	recv(timeout time.Duration) (mes, bool)
}


type paxosNet struct {
	recv map[int]chan mes
}

func (n *paxosNet) send(m mes) {
	n.recv[m.to] <- m
}

func (n *paxosNet) rec(from int, timeout time.Duration) (mes, bool) {
	select {
	case m := <-n.recv[from]:
		return m, true
	case <-time.After(timeout):
		return mes{}, false
	}
}

type agentNet struct {
	id int
	pn *paxosNet
}

func (an *agentNet) send(m mes) {
	an.pn.send(m)
}

func (an *agentNet) recv(timeout time.Duration) (mes, bool) {
	return an.pn.rec(an.id, timeout)
}
