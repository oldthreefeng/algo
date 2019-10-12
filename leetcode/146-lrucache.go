/*
@Time : 2019/10/12 10:04
@Author : louis
@File : 146-lrucache
@Software: GoLand
*/

package leetcode

import "container/list"

type LRUCache struct {
	cap int                   // capacity
	l   *list.List            // doubly linked list
	m   map[int]*list.Element // hash table for checking if list node exists
}

type pair struct {
	key   int
	value int
}

func NewLRU(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.m[key]; ok {
		val := node.Value.(*list.Element).Value.(pair).value
		//move the node to front
		lru.l.MoveToFront(node)
		return val
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.m[key]; ok {
		// 存在,则直接移动node到链表首部
		lru.l.MoveToFront(node)
		node.Value.(*list.Element).Value = pair{key: key, value: value}

	} else {
		// list 满了之后 ; 删除链表的最后一个节点
		if lru.l.Len() == lru.cap {
			delIndex := lru.l.Back().Value.(*list.Element).Value.(pair).key
			// 删除hashMap里面的最后的元素
			delete(lru.m, delIndex)
			// 删除list最后一个节点
			lru.l.Remove(lru.l.Back())
		}
		// 初始化list 节点
		node := &list.Element{
			Value: pair{
				key:   key,
				value: value,
			},
		}
		// 将新的node节点放入list
		p := lru.l.PushFront(node)
		lru.m[key] = p
	}
}
