/*
@Time : 2019/8/27 22:17
@Author : Administrator
@File : test_paxos
@Software: GoLand
*/

package paxos

import (
	"fmt"
	"testing"
	"time"
)

type node struct {
	key  int
	next *node
}

func TestN(t *testing.T) {
	first := &node{}
	cur := &node{}
	lastSeq := 1
	id := 0
	// 保证生成的数不相同
	// 生成编号 N 的提案  测试了1亿级别的数据,耗时15.53s
	// --- PASS: TestN (15.53s)
	const num = 10000000

	for id = 0; id <= num; id++ {
		n := &node{key: lastSeq<<16 | id}
		if id == 0 {
			first = n
			cur = n
		} else {
			cur.next = n
			if cur.key == cur.next.key {
				t.Errorf("有重复元素: %+v", cur.key)
			}
			cur = n
			lastSeq++
		}
	}
	t.Log(first)
}

func TestSingleProposer(t *testing.T) {
	// 1,2,3,4,5 acceptors
	// 1001 proposer
	// 2001 learner
	pn := newPaxNet(1, 2, 3, 4, 5, 1001, 2001)
	ac := make([]*acceptor, 0)
	for i := 1; i <= 5; i++ {
		ac = append(ac, newAcceptor(i, pn.agentNet(i), 2001))
	}

	for _, a := range ac {
		go a.run()
	}
	wantValue := "hello world"
	p := newPropose(1001, wantValue, pn.agentNet(1001), 1, 2, 3, 4, 5)
	go p.run()

	l := newLearner(2001, pn.agentNet(2001), 1, 2, 3, 4, 5)
	go l.learn()
	v := l.GetValue()
	if v != wantValue {
		t.Errorf("value = %s wanted value = %s", v, wantValue)
	}
}

func TestTwoPropose(t *testing.T) {
	// 1,2,3 acceptors
	// 1001,1002 proposer
	// 2001 learner
	pn := newPaxNet(1, 2, 3, 1001, 1002, 2001)
	ac := make([]*acceptor, 0)
	for i := 1; i <= 3; i++ {
		ac = append(ac, newAcceptor(i, pn.agentNet(i), 2001))
	}
	for _, a := range ac {
		go a.run()
	}

	wantV1 := "hello world"
	p1 := newPropose(1001, wantV1, pn.agentNet(1001), 1, 2, 3)
	go p1.run()

	wantV2 := "hello world v2"
	// 提出提案N 此时lastSeq++;
	p2 := newPropose(1002, wantV2, pn.agentNet(1002), 1, 2, 3)
	go p2.run()

	l := newLearner(2001, pn.agentNet(2001), 1, 2, 3)
	go l.learn()
	va := l.GetValue()
	if va != wantV2 {
		t.Errorf("value = %s,wantValue = %s", va, wantV2)
	}

}

func TestNPropose(t *testing.T) {
	pn := newPaxNet(1, 2, 3, 1001, 1002, 1003, 2001, 2002)
	ac := make([]*acceptor, 0)
	for i := 1; i <= 3; i++ {
		ac = append(ac, newAcceptor(i, pn.agentNet(i), 2001, 2002))
	}
	for _, a := range ac {
		go a.run()
	}
	pp := make([]*proposer, 0)
	for i := 1001; i <= 1003; i++ {
		wantStr := "hello world v" + fmt.Sprint(i)
		pp = append(pp, newPropose(i, wantStr, pn.agentNet(i), 1, 2, 3))
	}

	for _, p := range pp {
		go p.run()
	}
	//这里模拟两个learner
	ln := make([]*learner, 0)
	for i := 2001; i <= 2002; i++ {
		ln = append(ln, newLearner(i, pn.agentNet(i), 1, 2, 3))
	}
	var v [2]string
	for k, l := range ln {
		go l.learn()
		v[k] = l.GetValue()
	}
	if v[0] != v[1] {
		t.Errorf("value = %s,wantValue = %s", v[0], v[1])
	}
	time.Sleep(500 * time.Millisecond)
}
