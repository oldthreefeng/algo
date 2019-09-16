/*
@Time : 2019/9/3 0:25
@Author : louis
@File : poly
@Software: GoLand
*/

package main

import (
	"fmt"
)

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

type LinkPolyNode struct {
	head *PolyNode
}

func (l *LinkPolyNode) InsertFirst(c, e int) {
	date := &PolyNode{coef: c, exp: e}
	if l.head != nil {
		date.next = l.head
	}
	l.head = date
}

func (l *LinkPolyNode) GetOpposite() *LinkPolyNode {
	l1 := l
	p := l1.head
	if p == nil {
		return nil
	}
	for p != nil {
		p.coef *= -1
		p = p.next
	}
	return l1
}

func (l *LinkPolyNode) InsertLast(c, e int) {
	date := &PolyNode{coef: c, exp: e}
	if l.head == nil {
		// first one PolyNode
		date.next = l.head
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = date
}

func (l *LinkPolyNode) PrintItems() {
	p := l.head
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
		fmt.Printf("%dx^%d\n", p.coef, p.exp)
	}
}

func (l *LinkPolyNode) GetSize() (count int) {
	cur := l.head
	for cur != nil{
		count++
		cur = cur.next
	}
	return count
}

// p,q 为当前计算节点;  rear为指向和多项式链表的尾节点
// p.exp < q.exp ; 节点p所指的节点是"和多项式"中的一项,p后移
// p.exp = q.exp ; 合并同类项.
// p.exp > q.exp ; q节点插入p之前,q节点的指针在原来的链表上后移
// a= 7+3x+9x^8+5x^17
// b= 8x+22x^7-9x^8
func (l *LinkPolyNode) Add(lb *LinkPolyNode) (lc *LinkPolyNode) {
	lc = &LinkPolyNode{}
	lc.head = &PolyNode{} //初始化头节点的PolyNode
	// 记住head头指针
	pc := lc.head
	p, q := l.head.next, lb.head.next //采用尾插法,头结点没有元素,跳过
	for p != nil && q != nil {
		var tmp = &PolyNode{}
		/*1. 节点p所指的节点是"和多项式"中的一项,p后移 */
		if p.exp < q.exp {
			tmp.coef = p.coef
			tmp.exp = p.exp
			p = p.next
			/*2. 合并同类项,*/
		} else if p.exp == q.exp {
			tmp.coef = p.coef + q.coef
			tmp.exp = p.exp
			p = p.next
			q = q.next
			/*3. q节点插入rear节点,q节点的指针在原来的链表上后移*/
		} else {
			tmp.coef = q.coef
			tmp.exp = q.exp
			q = q.next
		}
		// 和不为零,将和插入至rear
		if tmp.coef != 0 {
			// 将rear.next指向tmp
			pc.next = tmp
			// 将rear指向tmp
			pc = tmp
		}
	}

	// 当p,q中有空的时候,说明其中的一个多项式已经算完了
	for p != nil {
		var tmp = &PolyNode{}
		tmp.coef = p.coef
		tmp.exp = p.exp
		p = p.next
		pc.next = tmp
		pc = tmp
	}

	for q != nil {
		var tmp = &PolyNode{}
		tmp.coef = q.coef
		tmp.exp = q.exp
		q = q.next
		pc.next = tmp //尾结点插入
		pc = tmp
	}

	return lc
}

// 创建头节点LinkPolyNode
func NewLinkPoly() (l *LinkPolyNode) {
	l = &LinkPolyNode{}
	l.head = &PolyNode{}
	rear := l.head
	var coef, exp int
	fmt.Scanln(&coef, &exp)
	for coef != 0 {
		s := &PolyNode{}
		s.coef = coef
		s.exp = exp
		rear.next = s
		rear = s
		fmt.Scanln(&coef, &exp)
	}
	return l
}

func main() {
	// a= 7+3x+9x^8+5x^17
	// b= 8x+22x^7-9x^8


	//b1 := Sub(b)
	//d := Add(a,b1)
	//ShowPoly(d)
	fmt.Println("请输入A多项式1*x^2==>(1 2),系数为0表示输入结束")
	a := NewLinkPoly()
	a.PrintItems()
	fmt.Println("请输入B多项式,系数为0表示输入结束")
	b := NewLinkPoly()
	b.PrintItems()
	fmt.Println("加法(0),减法(1):请输入:(默认为加法)")
	var key byte
	fmt.Scanf("%1d",&key)
	switch key {
	case 1:
		//取相反数,再相加就好了
		b1 := b.GetOpposite()
		d := a.Add(b1)
		d.PrintItems()
	case 0:
		fallthrough
	default:
		c := a.Add(b)
		c.PrintItems()
	}



}

/*
请输入A多项式1*x^2==>(1 2),系数为0表示输入结束
7 0
3 1
9 8
5 17
0
7x^0+3x^1+9x^8+5x^17
请输入B多项式,系数为0表示输入结束
8 1
22 7
-9 8
0
8x^1+22x^7-9x^8
加法(0),减法(1):请输入:(默认为加法)

7x^0+11x^1+22x^7+5x^17
*/
