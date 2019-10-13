/*
@Time : 2019/10/12 10:04
@Author : louis
@File : 146-lrucache
@Software: GoLand
*/

package leetcode

import "container/list"

// 内存缓存LRU算法: 最近最少使用

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

// 如果key存在缓存中, 则获取key的value
// 每次数据项被查询到时，都将此数据项移动到链表头部（O(1)的时间复杂度）
// 这样，在进行过多次查找操作后，最近被使用过的内容就向链表的头移动，而没有被使用的内容就向链表的后面移动
// 当需要替换时，链表最后的位置就是最近最少被使用的数据项，我们只需要将最新的数据项放在链表头部，
// 当Cache满时，淘汰链表最后的位置就是了。
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.m[key]; ok {
		val := node.Value.(*list.Element).Value.(pair).value
		//move the node to front
		lru.l.MoveToFront(node)
		return val
	}
	return -1
}

// 写入数据, 如果key不存在, 则写入其数据值, 当缓存容量达到上限时,
// 它应该在写入新数据之前删除最近最少使用的value.
func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.m[key]; ok {
		// 存在,则直接移动node到链表首部,并更新value
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

		// 初始化node节点
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
