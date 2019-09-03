/*
@Time : 2019/9/3 0:25
@Author : louis
@File : poly
@Software: GoLand
*/

package main

import "fmt"

// p,q 为当前计算节点;  rear为指向和多项式链表的尾节点
// p.exp < q.exp ; p后移
// p.exp = q.exp ; 和为0,A中删除p,释放p,q;和不为零,修改p的数据域,释放q节点
// p.exp > q.exp ; q节点插入p之前,q节点的指针在原来的链表上后移

type PolyNode struct {
	coef int
	exp  int
	next *PolyNode
}

func CreatePoly() *PolyNode {
	rear := &PolyNode{}
	head := rear
	s := &PolyNode{}
	var c, e int
	fmt.Scanf("%d,%d", &c, &e)
	for c != 0 {
		s.coef = c
		s.exp = e
		rear.next = s
		rear = s
		fmt.Scanf("%d,%d", &c, &e)
	}
	rear.next = nil
	return head
}
