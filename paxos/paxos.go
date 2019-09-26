/*
@Time : 2019/8/27 18:39
@Author : louis
@File : paxos
@Software: GoLand
*/

package paxos

import (
	"fmt"
	"log"
	"time"
)

const (
	//Prepare 准备发送的消息
	Prepare = iota + 1
	//Propose 待批准的消息-提案消息,并保证小于n的提案不在接收
	Propose
	//Promise accept返回的消息
	Promise
	//Accept 已经接收的提案
	Accept
)

//promise 接口只能获取消息的提案号
type promise interface {
	number() int
}

type accept interface {
	//accept接口获取提案的值
	proposalValue() string
	//accept接口获取提案的号
	proposalNumber() int
}

//mes 消息体
type mes struct {
	from, to int    //消息发送\接收
	typ      int    //消息类型
	n        int    //提案号
	pren     int    //前一个提案号
	value    string //提案值value
}

//proposalValue 只有是accept和promise类型的消息,才具备提案value值
func (m mes) proposalValue() string {
	//返回value需要什么条件呢？
	switch m.typ {
	case Accept, Promise:
		return m.value
	default:
		panic("unexpect proposalValue")
	}

}

func (m mes) proposalNumber() int {
	switch m.typ {
	case Promise:
		return m.pren
	case Accept:
		return m.n
	default:
		panic("unexpect proposalNumber")
	}
}

func (m mes) number() int {
	return m.n
}

//network to send and recv mes
type network interface {
	send(m mes)
	recv(timeout time.Duration) (mes, bool)
}

// paxosNet paxos 的消息管道,recv[i]消息的接收i接收的信息
type paxosNet struct {
	recv map[int]chan mes
}

// newPaxNet 生成paxosNet,根据agent的数量生成
func newPaxNet(agents ...int) *paxosNet {
	pn := &paxosNet{recv: make(map[int]chan mes, 0)}
	for _, a := range agents {
		pn.recv[a] = make(chan mes, 1024)
	}
	return pn
}

//send 发送消息mes至pn中的接收者i
func (pn *paxosNet) send(m mes) {
	log.Printf("nt send message :%+v", m)
	pn.recv[m.to] <- m
}

//rec 从agent接收mes,并输出,返回信息mes和bool.
func (pn *paxosNet) rec(from int, timeout time.Duration) (mes, bool) {
	select {
	case m := <-pn.recv[from]:
		log.Printf("nt recv message :%+v", m)
		return m, true
	case <-time.After(timeout):
		return mes{}, false
	}
}

//agentNet 根据agent的id生成AgentNet结构体
func (pn *paxosNet) agentNet(id int) *agentNet {
	return &agentNet{id: id, pn: pn}
}

//Empty 判断pn是不是空
func (pn *paxosNet) empty() bool {
	var n int
	for i, q := range pn.recv {
		log.Printf("nt %+v left %d", i, len(q))
		n += len(q)
	}
	return n == 0
}

//AgentNet 代理网络结构体,存放代理的id号和pn消息网络结构体
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

//proposer 提案提出者
type proposer struct {
	id        int    //提案号
	lastSeq   int    //上一条Seq消息好
	value     string //提案的value
	valueN    int
	acceptors map[int]promise //接受者的消息map
	nt        network //paxos的网络
}

//newPropose 生成新的提案提出者proposer
func newPropose(id int, value string, nt network, acceptors ...int) *proposer {
	p := &proposer{
		id:        id,
		lastSeq:   0,
		nt:        nt,
		value:     value,
		acceptors: make(map[int]promise),
	}
	for _, a := range acceptors { //遍历acceptors,生成proposer
		p.acceptors[a] = mes{}
	}
	return p
}

func (p *proposer) run() {
	var ok bool
	var m mes
	//c.2 acceptor里返回的promise达到半数以上
	for !p.quorumCheck() {
		if !ok {
			//a. Proposer准备prepare,生成提案编号N，然后向半数以上的Acceptor发送编号为N的Prepare请求。
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
		switch m.typ {
		case Promise:
			//c.1 更新proposer里面的acceptor,保证proposer存入的acceptor提案N为最新
			p.receivePromise(m)
		default:
			log.Panicf("proposer: %d unexpected message type: %v", p.id, m.typ)
		}
	}
	log.Printf("%d promise %d reached quorum %d", p.id, p.n(), p.quorum())
	//c.3 如果Proposer收到半数以上Acceptor对其发出的编号为N的Prepare请求的响应，
	//   那么它就会发送一个针对[N,V]提案的Accept请求给半数以上的Acceptor
	ms := p.propose()
	for i := range ms {
		fmt.Printf("proposer %d: ", p.id)
		p.nt.send(ms[i])
	}
}

//大多数,半数+1即为大多数
func (p *proposer) quorum() int {
	return len(p.acceptors)/2 + 1
}

//c.2 acceptor返回的promise大于半数以上同意
func (p *proposer) quorumCheck() bool {
	m := 0
	for _, promise := range p.acceptors {
		// promise里面的提案N 必须和 proposer的提案N相同.
		if promise.number() == p.n() {
			m++
		}
	}
	if m >= p.quorum() {
		return true
	}
	return false
}

//在一个paxos实例中，每个提案需要有不同的编号，且编号间要存在全序关系
func (p *proposer) n() int {
	// 把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0
	// 再把结果和p.id 进行或运算,将1位合并置1
	return p.lastSeq<<16 | p.id
}

//c.3 如果Proposer收到半数以上Acceptor对其发出的编号为N的Prepare请求的响应，把proposer的value携带上.
//   那么它就会发送一个针对[N,V]提案的Accept请求给半数以上的Acceptor
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
				typ:   Propose,
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

//a. Proposer选择一个提案编号N，然后向半数以上的Acceptor发送编号为N的Prepare请求。
func (p *proposer) prepare() []mes {
	p.lastSeq++
	m := make([]mes, p.quorum())
	i := 0
	// 只取acceptors里面的index,即acceptors
	for to := range p.acceptors {
		m[i] = mes{
			from: p.id,
			to:   to,
			typ:  Prepare,
			n:    p.n(),
		}
		i++
		if i == p.quorum() {
			break
		}
	}
	return m
}

//c.1 更新proposer里面的acceptor,保证proposer存入的acceptor提案N为最新
// 从acceptor接收的promise消息.
func (p *proposer) receivePromise(promise mes) {
	prePromise := p.acceptors[promise.from]
	// 从acceptor接收的promise和propose里面的number进行比对
	// propose.acceptors[promise.from].number() < promise.number()
	// 返回的promise大,则更新proposer里面的promise为最新版.
	if prePromise.number() < promise.number() {
		log.Printf("proposer: %d received a new promise %+v", p.id, promise)
		p.acceptors[promise.from] = promise
		//acceptors返回的提案preN > proposer的N
		if promise.proposalNumber() > p.valueN {
			log.Printf("proposer: %d updated the value [%s] to %s",
				p.id, p.value, promise.proposalValue())
			//promise.pren = p.n()
			p.valueN = promise.proposalNumber()
			p.value = promise.proposalValue()
		}
	}
}

type acceptor struct {
	id int
	//一旦acceptor接收提案propose;
	// 便需要通知所有learners,通信总次数为 M * N
	//TODO 通知learner集合,再由learner集合通知剩下的learner
	learners []int
	accept   mes
	promised promise
	nt       network
}

func newAcceptor(id int, nt network, learners ...int) *acceptor {
	return &acceptor{
		id:       id,
		nt:       nt,
		promised: mes{},
		learners: learners,
	}
}

func (a *acceptor) run() {
	for {
		m, ok := a.nt.recv(time.Hour)
		if !ok {
			continue
		}
		switch m.typ {
		case Propose:
			//d. 接收proposer的提案N,如果已接收提案N>返回的propose的提案N,则忽略
			//    如果已接收提案N< 返回的propose提案N,说明错误;违反了P2c原则
			//    如果相同,则接收提案N,将propose提案存入accept里面,并将accept.typ改为接受.
			accepted := a.receivePropose(m)
			if accepted {
				//e. 将accept信息发送至learner,通知learner我接受提案N
				for _, l := range a.learners {
					m = a.accept
					m.from = a.id
					m.to = l
					a.nt.send(m)
				}
			}
			//b. acceptor返回promise,并保证小于n的提案不在接收
		case Prepare:
			promised, ok := a.receivePrepare(m)
			if ok {
				a.nt.send(promised)
			}
		default:
			log.Panicf("accepted : %d message tpye unknwon: %d", a.id, m.typ)
		}
	}
}

// d. 接收proposer的提案N,如果已接收提案N>返回的propose的提案N,则忽略
//    如果已接收提案N< 返回的propose提案N,说明错误;违反了P2c原则
//    如果相同,则接收提案N,将propose提案存入accept里面,并将accept.typ改为接受.
//P2c：如果一个编号为 n 的提案具有 value v，那么存在一个多数派，要么他们中所有人都没有接受（accept）编号小于 n
//的任何提案，要么他们已经接受（accept）的所有编号小于 n 的提案中编号最大的那个提案具有 value v。
func (a *acceptor) receivePropose(propose mes) bool {
	// 已接收提案N > propose的mes.n;不接收这个提案
	if a.promised.number() > propose.number() {
		log.Printf("acceptor %d [promised: %+v] ignored propose mes: %+v", a.id, a.promised, propose)
		return false
	}
	//已接收提案N < propose的mes.n;
	if a.promised.number() < propose.number() {
		log.Panicf("acceptor %d [promised: %+v] received unexpected proposal mes: %+v",
			a.id, a.promised, propose)
	}
	log.Printf("acceptor %d [promised: %+v, accept: %+v]  accepted propose: %+v",
		a.id, a.promised, a.accept, propose)
	a.accept = propose
	a.accept.typ = Accept
	return true
}

//b. acceptor返回promise,并保证小于n的提案不在接收
func (a *acceptor) receivePrepare(prepare mes) (promised mes, b bool) {
	// 如果获取的m.n大于提案N,Promised提案接收,承诺不再接收任何小于N的提案
	// P1:一个 acceptor 必须接受（accept）第一次收到的提案。
	if a.promised.number() < prepare.number() {
		log.Printf("acceptor %d [promised: %+v]  promised %+v", a.id, a.promised, prepare)
		a.promised = prepare
		//把消息返回
		promised = mes{
			typ:   Promise,
			from:  a.id,
			to:    prepare.from,
			n:     a.promised.number(),
			pren:  a.accept.n,
			value: a.accept.value,
		}
		return promised, true
	}
	log.Printf("acceptor %d [promised: %+v] ignored prepare mes: %+v", a.id, a.promised, prepare)
	return mes{}, false
}

type learner struct {
	id        int
	acceptors map[int]accept
	nt        network
	value     chan string //测试数据比对,learner学习后得到的提案N对应的V[N,V]
}

func newLearner(id int, nt network, acceptors ...int) *learner {
	l := &learner{id: id, nt: nt, acceptors: make(map[int]accept), value: make(chan string)}
	for _, a := range acceptors {
		l.acceptors[a] = mes{typ: Accept}
	}
	return l
}

func (l *learner) GetValue() (v string) {
	select {
	case v := <-l.value:
		return v
	case <-time.After(time.Second):
		return
	}
}

func (l *learner) learn() {
	for {
		//f. 等待acceptor发送accept mes,
		m, ok := l.nt.recv(time.Hour)
		if !ok {
			continue
		}
		//f.1 如果消息类型不为Accept, 返回错误
		if m.typ != Accept {
			log.Panicf("learner :%d receive an unexpected proposal mes: %+v", l.id, m)
		}
		//f.2 从accepted消息中来进行比对,如果接收的提案N > learner存入的提案N,需要重新学习;否则就忽略
		l.receiveAccept(m)
		//g. 如果半数以上的learner选择了提案N,则说明选择完毕
		accept, ok := l.chosen()
		if !ok {
			continue
		}
		log.Printf("learner :%d has chosen proposal : %v ", l.id, accept)
		l.value <- accept.proposalValue()
		return
	}
}

//g. 如果半数的learner选择了提案N,则说明选择完毕
func (l *learner) chosen() (accept, bool) {

	counts := make(map[int]int)
	accepts := make(map[int]accept)

	for _, accepted := range l.acceptors {
		// 统计learner接收提案的次数;为0说明没有接收过提案
		if accepted.proposalNumber() != 0 {
			counts[accepted.proposalNumber()]++
			accepts[accepted.proposalNumber()] = accepted
		}
	}

	for n, count := range counts {
		// quorum达到即返回
		if count >= l.quorum() {
			return accepts[n], true
		}
	}

	return mes{}, false
}

func (l *learner) quorum() int {
	return len(l.acceptors)/2 + 1
}

//f.2 从accepted消息中来进行比对,如果接收的提案N > learner存入的提案N,需要重新学习;否则就忽略
func (l *learner) receiveAccept(accepted mes) {
	a := l.acceptors[accepted.from]
	// 提案N < 接收的 N; 需要接收大于N的提案
	if a.proposalNumber() < accepted.n {
		log.Printf("learner %d has learned a new proposal mes: %+v", l.id, accepted)
		l.acceptors[accepted.from] = accepted
	}
}
