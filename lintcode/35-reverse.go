/*
Copyright 2019 louis.
@Time : 2019/9/28 16:42
@Author : louis
@File : 35-reverse
@Software: GoLand

*/

package lintcode

type LinkNode struct {
	Val int
	Next *LinkNode
}

func ReverseLink(l *LinkNode) *LinkNode  {
	if l == nil {
		return l
	}
	var pre,cur *LinkNode
	for l != nil {
		//l.Next,l,pre = pre,l.Next,l
		cur = l.Next  	//  cur的指针指向l.Next,记录l.next位置
		//尾插法插入pre,pre一直处于头指针位置
		l.Next = pre  	//  l.next的指针pre
		pre = l			//  pre的指针指向l
		l = cur			//  l的指针指向cur ==> l已经在链表的尾部,l.next
	}
	return pre
}

func ReverseLinkX(l *LinkNode) *LinkNode  {
	if l == nil {
		return l
	}
	var pre *LinkNode
	for l != nil {
		l.Next,l,pre = pre,l.Next,l
	}
	return pre
}