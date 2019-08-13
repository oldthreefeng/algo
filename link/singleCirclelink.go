package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertNode(head *CatNode, newNode *CatNode)  {

	// 如果只有一个元素,构成单元数链表
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.next = head
		return
	}

	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	newNode.next = head

	return
}

// 必须要返回一个*CatNode,因为如果删除头结点,那么链表就丢失了
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	tail := head
	for { //取得最后的指针
		if tail.next == head {
			break
		}
		tail = tail.next
	}

	if temp.next == nil { //链表为空,或者
		return head
	}
	if temp.next == head && temp.no == id { //链表已到头部且要删除的就是这个,将head指针置空
		temp.next = nil
		return head
	}

	flag := true
	for {
		if temp.next == head { // 找到最后都没找到
			break
		}
		if temp.no == id { //比当前节点
			if temp == head { //删除头节点
				head = temp.next //把下一个节点当初头节点
			}
			tail.next = temp.next // 并将尾节点结点指向新的头结点
			flag = false
			break
		}
		temp = temp.next //移动至下一个  这两个是同步移动
		tail = tail.next // 移动至下一个
	}

	if flag {
		if temp.no == id { //最后一个节点还未比较
			tail.next = temp.next
		} else {
			fmt.Println("del no is not exist")
		}
	}

	return head
}

func ListSingleCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("empty singleCircleLink")
		return
	}
	for {
		fmt.Printf("%d,%s,%p,%p ==>\n", temp.no, temp.name, temp, temp.next)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	head := &CatNode{}
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "jerry",
	}
	cat3 := &CatNode{
		no:   3,
		name: "mike",
	}
	InsertNode(head, cat1)
	InsertNode(head, cat2)
	InsertNode(head, cat3)
	ListSingleCircleLink(head)

	head = DelCatNode(head, 1)
	ListSingleCircleLink(head)

}
