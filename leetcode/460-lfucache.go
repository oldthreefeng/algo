/*
@Time : 2019/10/12 11:24
@Author : louis
@File : 460-lfucache
@Software: GoLand
*/

package leetcode

import (
	"container/list"
)

/*
用2个hash-map，CntM的key为访问次数，value为缓存元素组成的链表；
KMap的key为缓存元素的key，value为CntM的链表的节点指针。
关键还有一个int型的Min，记录当前访问次数最小值
*/
type LFUCache struct {
	Cap  int
	kMap map[int]*list.Element //通过Key查找
	cntM map[int]*list.List    //通过次数查找
	min  int                   //最小的访问次数，只有在put需要删除项目，且缓存容量达上限时才用的上，新的项目cnt必然是1，所以只要记录最小的即可
}

type node struct {
	key   int
	value int
	cnt   int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		Cap:  capacity,
		kMap: make(map[int]*list.Element, capacity),
		cntM: make(map[int]*list.List),
		min:  0,
	}
}

func (lfu *LFUCache) Get(key int) int {
	e, ok := lfu.kMap[key]
	if ok { //存在Key，则更新Key对应的访问次数，并更新cntM和Min
		lfu.cntM[e.Value.(*node).cnt].Remove(e)
		if lfu.cntM[e.Value.(*node).cnt].Len() <= 0 { //对应访问次数的列表被清空时，应删除cntM的元素
			delete(lfu.cntM, e.Value.(*node).cnt)
			if lfu.min == e.Value.(*node).cnt {
				lfu.min++ //如果被删除的元素正好是最小次数，则最小次数++
			}
		}
		e.Value.(*node).cnt++
		e2, ok := lfu.cntM[e.Value.(*node).cnt]
		if ok {
			lfu.kMap[key] = e2.PushFront(e.Value)
		} else {
			l := list.New()
			lfu.kMap[key] = l.PushFront(e.Value)
			lfu.cntM[e.Value.(*node).cnt] = l
		}
		return e.Value.(*node).value
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.Cap <= 0 {
		return
	}
	e, ok := lfu.kMap[key]
	if ok { //存在Key，则更新Key对应的访问次数，并更新cntM和Min
		lfu.cntM[e.Value.(*node).cnt].Remove(e)
		if lfu.cntM[e.Value.(*node).cnt].Len() <= 0 { //对应访问次数的列表被清空时，应删除cntM的元素
			delete(lfu.cntM, e.Value.(*node).cnt)
			if lfu.min == e.Value.(*node).cnt {
				lfu.min++ //如果被删除的元素正好是最小次数，则最小次数++
			}
		}
		e.Value.(*node).cnt++
		e.Value.(*node).value = value
		e2, ok := lfu.cntM[e.Value.(*node).cnt]
		if ok {
			lfu.kMap[key] = e2.PushFront(e.Value)
		} else {
			l := list.New()
			lfu.kMap[key] = l.PushFront(e.Value)
			lfu.cntM[e.Value.(*node).cnt] = l
		}
	} else { //不存在Key，则新增Key，并更新CntM和Min
		for lfu.Cap <= len(lfu.kMap) { //容量不足，得删除掉最不经常使用的
			e := lfu.cntM[lfu.min].Back()
			lfu.cntM[lfu.min].Remove(e)
			if lfu.cntM[lfu.min].Len() <= 0 { //&& lfu.min != 1 {
				delete(lfu.cntM, lfu.min) //lfu.Min为1时不需要删除，因为下面马上就要添加了
			}
			delete(lfu.kMap, e.Value.(*node).key)
		}
		lfu.min = 1 //毋庸置疑
		n := node{key, value, 1}
		e2, ok := lfu.cntM[1]
		if ok {
			lfu.kMap[key] = e2.PushFront(&n)
		} else {
			l := list.New()
			lfu.kMap[key] = l.PushFront(&n)
			lfu.cntM[1] = l
		}
	}
}
