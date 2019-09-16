/*
@Time : 2019/8/20 15:28
@Author : louis
@File : lightout
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/imroc/biu"
	"gogs.wangke.co/go/algo/search"
)

func main() {
	var a byte
	a = 30 //00011110
	b := search.GetBit(a, 2)
	fmt.Printf("%d二进制为:%v\n", a, biu.ToBinaryString(a))
	fmt.Printf("a的第2位为:%v\n", b)
	b = search.SetBit(a, 5, 1)
	fmt.Printf("%d的第6位,置为1后的二进制为:%v\n", a, biu.ToBinaryString(b))
	b = search.SetBit(a, 6, 1)
	fmt.Printf("%d的第7位,置为1后的二进制为:%v\n", a, biu.ToBinaryString(b))
	b = search.SetBit(a, 7, 1)
	fmt.Printf("%d的第8位,置为1后的二进制为:%v\n", a, biu.ToBinaryString(b))
	b = search.SetBit(a, 0, 1)
	fmt.Printf("%d的第1位,置为1后的二进制为:%v\n", a, biu.ToBinaryString(b))
	b = search.SetBit(a, 1, 0)
	fmt.Printf("%d的第2位,置为0后的二进制为:%v\n", a, biu.ToBinaryString(b))

	for i := 0; i < 20; i++ {
		if search.IsPowerOfTwo(i) {
			fmt.Printf("%d \tisPowerOfTwo\n", i)
		}
	}

}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		n := target - v

		if x, ok := m[n]; ok {
			return []int{i, x}
		}

		m[v] = i
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var c = 0

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil && c == 0 {
		return nil
	}
	if l1 != nil {
		c += l1.Val
		l1 = l1.Next
	} else {
		c += 0
	}
	if l2 != nil {
		c += l2.Val
		l2 = l2.Next
	} else {
		c += 0
	}
	cur := &ListNode{}
	cur.Val = c % 10
	c = c / 10
	cur.Next = addTwoNumbers(l1, l2)
	return cur
}
