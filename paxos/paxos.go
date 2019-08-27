/*
@Time : 2019/8/27 18:39
@Author : louis
@File : paxos
@Software: GoLand
*/

package paxos

import (
	"log"
	"time"
)

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
	case Accept, Promise:
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

type proposer struct {
	id        int
	lastSeq   int
	value     string
	valueN    int
	acceptors map[int]promise
	nt        network
}

func NewPropose(id int, value string, nt network, acceptors ...int) *proposer {
	p := &proposer{
		id:        id,
		lastSeq:   0,
		nt:        nt,
		value:     value,
		acceptors: make(map[int]promise),
	}
	for _, a := range acceptors { //遍历acceptors,生成proposer
		p.acceptors[a] = &mes{}
	}
	return p
}

func (p *proposer) run() {
	var ok bool
	var m mes
	//1. Proposer选择一个提案编号N，然后向半数以上的Acceptor发送编号为N的Prepare请求。
	for !p.quorumCheck() {
		if !ok {
			ms := p.prepare()
			for i := range ms {
				p.nt.send(ms[i])
			}
		}
		m, ok = p.nt.recv(time.Second)
		// 返回数据失败,说明此次的prepare失败,重新prepare
		if !ok {
			continue
		}
		// prepare成功
		switch m.mesT {
		case Promise:
			p.receivePromise(&m)
		default:
			log.Panicf("proposer: %d unexpected message type: %v", p.id, m.mesT)
		}
	}
	log.Printf("%d promise %d reached quorum %d", p.id, p.n(), p.quorum())
	//2. 如果Proposer收到半数以上Acceptor对其发出的编号为N的Prepare请求的响应，
	//   那么它就会发送一个针对[N,V]提案的Accept请求给半数以上的Acceptor
	ms := p.propose()
	for i := range ms {
		p.nt.send(ms[i])
	}
}

//大多数,半数+1即为大多数
func (p *proposer) quorum() int {
	return len(p.acceptors)/2 + 1
}

func (p *proposer) quorumCheck() bool {
	m := 0
	for _, promise := range p.acceptors {
		if promise.number() == p.n() {
			m++
		}
	}
	if m >= p.quorum() {
		return true
	}
	return false
}

//向Proposer承诺保证不再接受任何编号小于N的提案。
func (p *proposer) n() int {
	// 把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0
	// 再把结果和p.id 进行或运算,将1位合并置1
	return p.lastSeq<<16 | p.id
}

//发送一个针对[N,V]提案的Accept请求给半数以上的Acceptor;生成一个消息切片,准备发送
func (p *proposer) propose() []mes {
	m := make([]mes, p.quorum())
	i := 0
	// 取acceptors里面的index,promise,只能取promise里面的mes.n
	for to, promise := range p.acceptors {
		//如果acceptors的n,和proposer的提案n相等,即acceptor接收proposer的提案n.
		if promise.number() == p.n() {
			m[i] = mes{
				from:  p.id,
				to:    to,
				mesT:  Propose,
				n:     p.n(),
				value: p.value,
			}
			i++
		}
		if i == p.quorum() {
			break
		}
	}
	return m
}

//(a) 向Proposer承诺保证不再接受任何编号小于N的提案。
//(b) 如果Acceptor已经接受过提案，那么就向Proposer响应已经接受过的编号小于N的最大编号的提案。
func (p *proposer) prepare() []mes {
	p.lastSeq++
	m := make([]mes, p.quorum())
	i := 0
	// 只取acceptors里面的index,即acceptors
	for to := range p.acceptors {
		m[i] = mes{
			from: p.id,
			to:   to,
			mesT: Prepare,
			n:    p.n(),
		}
		i++
		if i == p.quorum() {
			break
		}
	}
	return m
}

//暂未理解
func (p *proposer) receivePromise(promise *mes) {
	prePromise := p.acceptors[promise.number()]
	// 待接收的promise
	if prePromise.number() < promise.number() {
		log.Printf("proposer: %d received a new promise %+v", p.id, promise)
		p.acceptors[promise.number()] = promise

		if promise.proposalNumber() > p.valueN {
			//promise的preN大于proposer的valueN
			log.Printf("proposer: %d updated the value [%s] to %s",
				p.id, p.value, promise.proposalValue())
			p.valueN = promise.proposalNumber()
			p.value = promise.proposalValue()
		}
	}
}

type acceptor struct {
	id       int
	learners []int
	accept   mes
	promised promise
	nt       network
}

func NewAcceptor(id int, nt network, learners ...int) *acceptor {
	return &acceptor{
		id:       id,
		nt:       nt,
		promised: &mes{},
		learners: learners,
	}
}

// P1. An acceptor must accept the first proposal that it receives.
// If a proposal with value v is chosen, then every higher-numbered proposal
// accepted by any acceptor has value v.
func (a *acceptor) run() {
	for {
		m, ok := a.nt.recv(time.Hour)
		if !ok {
			continue
		}
		switch m.mesT {
		case Propose:
			accepted := a.receivePropose(m)
			if accepted {
				for _, l := range a.learners {
					m = a.accept
					m.from = a.id
					m.to = l
					a.nt.send(m)
				}
			}
		case Prepare:
			promised, ok := a.receivePrepare(m)
			if ok {
				a.nt.send(promised)
			}
		default:
			log.Panicf("accepted : %d message tpye unknwon: %d", a.id, m.mesT)
		}
	}
}

// If an acceptor receives an accept request for a proposal numbered
// n, it accepts the proposal unless it has already responded to a prepare
// request having a number greater than n.
func (a *acceptor) receivePropose(propose mes) bool {
	// 提案N > mes.n;不接收这个提案
	if a.promised.number() > propose.number() {
		log.Printf("acceptor %d [promised: %+v] ignored propose mes: %+v", a.id, a.promised, propose)
		return false
	}
	// 提案N < mes.n;
	if a.promised.number() < propose.number() {
		log.Panicf("acceptor %d [promised: %+v] received unexpected proposal mes: %+v",
			a.id, a.promised, propose)
	}
	log.Printf("acceptor %d [promised: %+v, accept: %+v]  accepted propose: %+v",
		a.id, a.promised, a.accept, propose)
	a.accept = propose
	a.accept.mesT = Accept
	return true
}

// If an acceptor receives a prepare request with number n greater
// than that of any prepare request to which it has already responded,
// then it responds to the request with a promise not to accept any more
// proposals numbered less than n and with the highest-numbered proposal
// (if any) that it has accepted.
func (a *acceptor) receivePrepare(m mes) (promised mes, b bool) {
	//如果获取的m.n大于提案N,Promised提案接收,承诺不再接收任何小于N的提案
	if a.promised.number() < m.number() {
		log.Printf("acceptor %d [promised: %+v]  promised mes: %+v", a.id, a.promised, m)
		a.promised = &m
		//把消息返回
		promised = mes{
			mesT:  Promise,
			from:  a.id,
			to:    m.from,
			n:     m.number(),
			pren:  a.accept.n,
			value: a.accept.value,
		}
		return
	}
	log.Printf("acceptor %d [promised: %+v] ignored prepare mes: %+v", a.id, a.promised, m)
	return mes{}, false
}

type learner struct {
	id        int
	acceptors map[int]accept
	nt        network
}

func NewLearner(id int, nt network, acceptors ...int) *learner {
	l := &learner{id: id, nt: nt, acceptors: make(map[int]accept)}
	for _, a := range acceptors {
		l.acceptors[a] = &mes{mesT: Accept}
	}
	return l
}

func (l *learner) learn() string {
	for {
		m,ok := l.nt.recv(time.Hour)
		if !ok {
			continue
		}
		if m.mesT != Accept {
			log.Panicf("learner :%d receive an unexpected proposal mes: %+v",l.id,m)
		}
		l.receiveAccept(m)
		accept,ok := l.chosen()
		if !ok {
			continue
		}
		log.Printf("learner :%d has chosen proposal : %v ",l.id,accept)
		return accept.proposalValue()
	}
}

func (l *learner) chosen() (accept,bool)  {

}

func (l *learner) quorum() int {
	return len(l.acceptors)/2 +1
}
func (l *learner) receiveAccept(accept mes) bool  {

}