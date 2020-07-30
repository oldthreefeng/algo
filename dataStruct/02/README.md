## 线性表

线性表是有n个类型相同的数据a1,a2,..ai,..,an组成的有限序列.记做(a1,a2,..,ai,...,an).
数据元素是一对一的关系,每个数据元素最多有一个直接前驱和直接后继.

- 特点

```cgo
同一性: 每个ai必须属于同一数据对象
有穷性:有限数据元素组成
有序性:<ai,ai+1>
```

### 线性表的顺序存储

每个元素占k个单元.第一个元素地址为: loc(a1);则第i个元素的地址为:
loc(ai) = loc(a1) + (i-1)*k
- 查找操作; L.elem[i] , locate(L,x); O(n/2)
- 插入操作: 从an倒序操作, an --> an+1;然后插入ai = x.

### 单链表

采用链式结构来存放数据

单链表包括两个域:数据域,用来存储节点的数据值;指针域,用来存储数据元素的直接后继的地址. 

```cgo
type LinkList struct {
    n int   //数据域
    next *Link // 指针域
}
// 不能丢失链的信息
```

插入i元素
```go
func InsLink(head *LinkList,new *LinkList) {
	pre := head
	k:=0
	for pre.next != nil && k < new.n-1 {
		pre=pre.next
		k=k+1
	}
	if k!=new.n-1 {
		return 
	}
	//不能换位置
	new.next=pre.next
	pre.next=new
}

func DelLink(head *LinkList, i int) {
	pre := head
	k:=0
	for pre.next != nil && k < i-1  {
    		pre=pre.next
    		k=k+1
    }
	if k!=i-1 {
		return 
    }
	pre.next = pre.next.next
	
}
```

有时候为了方便操作,在单链表的第一个节点之前附加一个节点,称为头结点.

- 建立单链表

头插法建表.在链表的头结点插入节点

```go
func createFromHead() {
	// L头结点 , s待插入节点
	L :=&Link
    for x!=-1 {
    	s.data = x  //赋值
    	if L.next == nil {
    		L.next = s  // 将L.next 指向s
    		fmt.scanf("%d",x)
    		continue
    	}
    	s.next = L.next  // 将s.next指向L.next
    	L.next = s  // 将L.next 指向s
    	fmt.scanf("%d",x)
    }
}
```

尾插法建表,在尾部插入数据.在循环单链表中,尾指针会更好.

```go
func createFromTail() {
	H := &Link{}
	s := &Link{}
	r := H
	var x int
	fmt.scanf("%d",x)
	for x!=-1 {
		s.data =x
		s.next = r.next
		r.next = s
		r = s;
		fmt.scanf("%d",x)
	}
}
```

- 循环链表

```go
// 合并两个循环链表,带头指针,合并算法时间为O(n)
func merge01(La *LinkList,Lb *LinkList) *LinkList {
	p := La
	q := Lb
	for p.next != La {
		p=p.next
	}
	for q.next != Lb {
		q=q.next
    }
	p.next=Lb.next
	q.next=La
	return La
}

//带尾指针合并循环表. 算法的提升为O(1)
func merge02(Ra *LinkList,Rb *LinkList) *LinkList {
	//保存Ra的头结点
	p:= Ra.next
	//把Ra的next节点指向Rb的第一个节点.
	Ra.next = Rb.next.next
	//把Rb的next节点指向Ra的头结点
	Rb.next = p
	return Rb
}

```

### 双向链表

在单链表的链表的指针域,增加一个前驱的指针域pre

```cgo
type DoubleLink struct {
    n int
    pre *DoubleLink
    next *DoubleLink
}

```
- 插入操作

```go
//在p节点前插入s节点
// 先改p的前驱节点与插入节点
p.pre.next = s
s.pre = p.pre 
// 后改p的节点与要插入的前驱节点.
p.pre = s
s.next = p 
```

- 删除操作

```go
// 删除p节点前的s节点
p.pre.next = p.next
p.next.pre = p.re

```

### 一元多项式的表示和运算

P = p0+p1X+p2X^2+...+PnX^n
线性表(p0,p1,..,pn)

```cgo
type PolyNode struct {
    coef int
    exp int
    next *PolyNode
}

// A(X)=7+3X+9X^8+5X^17
// B(X)=8X+22X^7-9X^8
// p.exp < q.exp ; p后移
// p.exp = q.exp ; 和为0,A中删除p,释放p,q;和不为零,修改p的数据域,释放q节点
//  p.exp > q.exp ; q节点插入p之前,q节点的指针在原来的链表上后移

```

