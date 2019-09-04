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
// a= 7+3x+9x^8+5x^17
// b= 8x+22x^7-9x^8
type PolyNode struct {
	coef int
	exp  int
	next *PolyNode
}

func (p *PolyNode) append(q *PolyNode) {
	tmp := p
	// 找到链表最后节点
	for tmp.next != nil {
		tmp = tmp.next
	}
	// 添加新节点
	tmp.next = q
}

func CreatePoly() *PolyNode {
	head := &PolyNode{}
	s := &PolyNode{}
	fmt.Scanln(&s.coef, &s.exp)
	for s.coef != 0 {
		head.append(s)
		s = &PolyNode{}
		fmt.Scanln(&s.coef, &s.exp)
	}
	s.next = nil
	return head
}

func ShowPoly(head *PolyNode) {
	p := head
	if p == nil {
		fmt.Println("empty")
	} else {
		for p.next != nil {
			// 头结点的零值去掉
			if p.coef == 0 {
				p = p.next
				continue
			}
			// 如果下一个节点的coef大于0,说明是加,否则就是减
			if p.next.coef >= 0 {
				fmt.Printf("%dx^%d+", p.coef, p.exp)
			} else {
				fmt.Printf("%dx^%d", p.coef, p.exp)
			}
			p = p.next
		}
		fmt.Printf("%dx^%d\n ", p.coef, p.exp)
	}

}

func main() {
	// a= 7+3x+9x^8+5x^17
	// b= 8x+22x^7-9x^8
	a := CreatePoly()
	ShowPoly(a)
	//b := CreatePoly()
	//ShowPoly(b)
}