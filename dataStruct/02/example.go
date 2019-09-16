/*
@Time : 2019/9/13 0:03
@Author : louis
@File : example
@Software: GoLand
*/

package main

import "fmt"

type link struct {
	val  int
	next *link
}

type headLink struct {
	head *link
}

func exchange(arr []int) {
	a := 0
	b := len(arr) - 1
	for a < b {
		arr[a], arr[b] = arr[b], arr[a]
		a++
		b--
	}
}

func (l *headLink) insert(val int) {
	date := &link{val: val}
	if l != nil {
		date.next = l.head
	}
	l.head = date
}

func (l *headLink) reserve()  {
	// 空节点或者只有一个节点就直接return,不计算
	if l.head == nil || l.head.next == nil || l.head.next.next == nil {
		return
	}
	var pre *link
	cur := l.head.next
	for cur !=nil {
		// 记录cur的下一个节点位置
		tmp := cur.next
		// 将cur以尾插法加入pre
		cur.next = pre
		pre = cur
		// 将cur指向cur.next
		cur = tmp
	}
	l.head.next = pre
}

func exchangeLink(l *link) *link {
	cur := l
	var pre *link
	//l.head.next = nil
	for cur != nil {
		pre, cur, cur.next = cur, cur.next, pre
	}
	return pre
}

func list(l *link) (arr []int) {
	cur := l
	if cur == nil {
		return nil
	}
	for {
		arr = append(arr, cur.val)
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	return arr
}

func main() {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//exchange(arr)
	//fmt.Println(arr)

	var l = &headLink{}
	l.insert(1)
	l.insert(2)
	l.insert(3)
	l.insert(4)
	l.insert(5)
	fmt.Println(list(l.head))
	pre := exchangeLink(l.head)
	fmt.Println(list(pre))
	l1 := l
	if b := getNewLink(l1.head, l.head, 1, 3, 2); b {

	}
}

/*
已知两个单链表A和B，其头指针分别为la和lb，
编写一个过程从单链表A中删除自第i个元素起的共len个元素，
然后将单链表A插入到单链表B的第j个元素之前。
*/
func getNewLink(la, lb *link, i, j, lens int) bool {
	if i < 1 || j < 1 || lens < 1 { //参数错误
		return false
	}
	var p, q *link
	p = la //p为链表A的工作指针,初始化为A的头指针

	//查找第i个节点,p指向i-1
	k := 1
	for k < i-1 && p != nil {
		k++
		p = p.next
	}
	// i超过了链表A,退出.
	if p == nil {
		return false
	}
	q = p.next // q也为工作指针, 初始化指向A链表的第一个被删除的节点,也就是i节点
	k = 1
	for k <= lens && q != nil { //指针后移,移到i+lens
		k++
		q = q.next
	}
	if k <= lens {
		return false
	}

	p.next = q //A 链表删除 len个元素,将p的next指到

	if la.next != nil { //链表节点全部删除,无须插入
		for p.next != nil {
			p = p.next
		} // 找到A的尾指针
		q = lb
		k = 0
		for q != nil && k < j-1 { //查到第j个节点,
			k++
			q = q.next
		}
		if q == nil { // 找不到j节点,说明j过大
			return false
		}
		p.next = q.next // 将A链表链入
		q.next = la.next
		// 将A的第一个元素链在B的第j-1个节点.
	}

	return true
}

// 设指针la和lb分别指向两个无头结点单链表中的首元结点，
// 试设计从表la中删除自第i个元素起共len个元素，
// 并将它们插入到表lb的第j个元素之后的算法。

func DelInset(la, lb *link, i, j, lens int) bool {
	if i < 1 || j < 1 || lens < 1 { //参数错误
		return false
	}
	p := la
	for k := 0; k < i && p != nil; k++ {
		p = p.next
		//la, 找到 i节点;
	}
	// 循环结束,p为nil,说明i超出限制
	if p == nil {
		return false
	}
	q := lb
	for k := 0; k < j && q != nil; k++ {
		q = q.next
		//lb, 找到 j节点
	}
	// 循环结束,q为nil,说明j超出限制
	if q == nil {
		return false
	}
	// 进行插入操作
	tmp := q.next //记录当前q.next的指针.
	for p !=nil { //p==nil时,说明插入完成
		q.next = p
		p = p.next
		q = q.next
	}
	q.next = tmp
	return true
}
