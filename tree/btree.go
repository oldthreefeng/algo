/*
@Time : 2019/9/19 21:52
@Author : louis
@File : btree
@Software: GoLand
*/

package tree

import (
	"fmt"
)

const MaxSize = 100

type btree struct {
	data int
	l    *btree
	r    *btree
}

// NLR
// f(b) 是大问题 --> f(b.l) f(b.r)
func (b *btree) PreOrder() {
	if b != nil {
		fmt.Printf("%c ", b.data)
		b.l.PreOrder()
		b.r.PreOrder()
	}
}

// LNR
func (b *btree) InOrder() {
	if b != nil {
		b.l.InOrder()
		fmt.Printf("%c ", b.data)
		b.r.InOrder()
	}
}

// LRN
func (b *btree) PostOrder() {
	if b != nil {
		b.l.PostOrder()
		b.r.PostOrder()
		fmt.Printf("%c ", b.data)
	}
}

// 对于一个二叉树

func (b *btree) LevelOrder() {
	var p *btree
	var qu [MaxSize]*btree
	var front, rear int
	rear++
	qu[rear] = b
	for front != rear {
		front = (front + 1) % MaxSize
		p = qu[front]
		fmt.Printf("%c ", p.data)
		if p.l != nil {
			rear = (rear + 1) % MaxSize
			qu[rear] = p.l
		}
		if p.r != nil {
			rear = (rear + 1) % MaxSize
			qu[rear] = p.r
		}
	}
}

//Nodes 计算所有节点个数
func (b *btree) Nodes() int {
	if b == nil {
		return 0
	} else {
		return b.l.Nodes() + b.r.Nodes() + 1
	}
}

//Layers
func (b *btree) Layers() int {
	if b == nil {
		return 0
	}
	left := b.l.Layers()
	right := b.r.Layers()
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// LeafNodes 计算所有叶子节点个数
func (b *btree) LeafNodes() int {
	if b == nil {
		return 0
		//叶子节点的度为0,说明没有左右节点,左右节点为空
	} else if b.l == nil && b.r == nil {
		return 1
	} else {
		return b.l.LeafNodes() + b.r.LeafNodes()
	}
}

func (b *btree) Copy(t *btree) {
	if b == nil {
		return
	} else {
		t = &btree{}
		t.data = b.data
		b.l.Copy(t.l)
		b.r.Copy(t.r)
	}
}

func (b *btree) Swap(t *btree) {
	if b == nil {
		return
	} else {
		t = &btree{}
		t.data = b.data
		b.l.Swap(t.r)
		b.r.Swap(t.l)
	}

}

// 求根节点到叶子节点的逆路径

// 构造二叉树  任何n个不同节点的二叉树,都可以有他的中序序列和后序序列唯一的确认

func (b *btree) NewBtree() {
	b.data = 'A'
	b.l = &btree{data: 'B'}
	b.r = &btree{data: 'C'}
	b.l.l = &btree{data: 'D'}
	b.l.r = &btree{data: 'E'}
	b.l.r.l = new(btree)
	b.l.r.l.data = 'H'
	b.l.r.l.r = &btree{data: 'J'}
	b.r.l = &btree{data: 'F'}
	b.r.l.r = &btree{data: 'I'}
	b.r.r = &btree{data: 'G'}
}