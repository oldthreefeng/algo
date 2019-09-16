/*
@Time : 2019/9/11 22:27
@Author : louis
@File : linkqueue
@Software: GoLand
*/

package main

type node struct {
	date int
	next *node
}

type queue struct {
	rear *node
}

func (q *queue) EnterQueue(i int) {
	n := &node{date: i}
	//如果q.rear为空,则n为第一个节点;
	if q.rear == nil {
		q.rear = n
	} else {
		// 如果q的尾节点不为空.再把rear节点指向刚加的节点n.
		n.next = q.rear
		q.rear = n
	}
}

func (q *queue) DeleteQueue() (i int, b bool) {
	if q.rear == nil {
		return 0, false
	}
	// only one data
	if q.rear.next == nil {
		i = q.rear.date
		q.rear = nil //最后一个出队,空队列了.
		return i, true
	}
	cur := q.rear
	// 大于两个节点时先找到第一个节点.
	for  {
		// 如果下下一个节点为空.说明是最后一个
		if cur.next.next == nil {
			i = cur.next.date
			cur.next = nil
			return i, true
		}
		cur = cur.next
	}
}

func (q *queue) Peek() (i int, b bool) {
	if q.rear ==nil {
		return 0,false
	}
	return q.rear.date,true
}

func (q *queue) Get() (a []int) {
	if q.rear == nil {
		return
	}
	cur := q.rear
	for cur.next != nil {
		a = append(a,cur.date)
		cur = cur.next

	}
}
