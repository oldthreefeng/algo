/*
@Time : 2019/8/27 22:17
@Author : Administrator
@File : test_paxos
@Software: GoLand
*/

package paxos

import (
	"log"
	"testing"
)

type node struct {
	key int
	next *node
}

func TestN(t *testing.T) {
	first :=&node{}
	cur := &node{}
	lastSeq := 1
	id := 0
	// 保证生成的数不相同
	// 生成编号 N 的提案  测试了1亿级别的数据,耗时15.53s
	// --- PASS: TestN (15.53s)
	const num  = 100000000

	for id = 0; id <= num; id++ {
		n :=&node{key:lastSeq<<8|id}
		if id == 0 {
			first = n
			cur = n
		} 
		cur.next = n
		if cur.key == cur.next.key {
				log.Fatalln("有重复元素")
		}
		cur = n
		lastSeq++
	}

	log.Println(first)


}
