/*
@Time : 2019/9/3 19:56
@Author : louis
@File : josephu
@Software: GoLand
*/

package main

import (
	"fmt"
)

type Jose struct {
	no int
	//password int
	next *Jose
}

//创建数为n的环,返回尾结点.
func CreateJo(n int) *Jose {
	first := &Jose{}
	p := first
	if n < 1 {
		return first
	}
	for i := 1; i <= n; i++ {
		q := &Jose{
			no: i,
		}
		//第一个元素,
		if i == 1 {
			first = q
			p = q
			p.next = first
		} else {
			// 将p的next指针指向新加入的q,
			p.next = q
			// 将p的指针指向q,即向后移动
			p = q
			// 将p的next指针指向first,形成环
			p.next = first
		}
	}
	// 返回tail指针
	return p
}

func ListJo(tail *Jose) (index int) {
	p := tail
	index = 0
	// 传进来的tail指针不能为空,
	for p.next != nil {
		index++
		fmt.Printf("id = %d -> ", p.next.no)
		//并且下一个不能是自己.
		if p.next == tail {
			break
		}
		p = p.next
	}
	fmt.Println()
	return index
}

func Game(tail *Jose, start, password int) {
	// 如果tail为空,或者开始的数大于约瑟夫环的总数
	if tail.next == nil || start > ListJo(tail) {
		return
	}
	// 获取first指针和 tail指针
	first := tail.next

	// 移动到开始数数的那个位置
	for i := 1; i < start; i++ {
		first = first.next
		tail = tail.next
	}

	for {
		// 开始数password下,
		for i := 1; i < password; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("id = %d,出列 ->\n", first.no)
		// 移除first
		first = first.next
		tail.next = first
		if tail == first {
			break
		}
	}
	// 最后一个出列
	fmt.Printf("id = %d,出列 ->\n", first.no)
}

func main() {
	tail := CreateJo(7)
	Game(tail, 1, 3)
}
