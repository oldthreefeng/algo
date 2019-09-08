package main

import (
	"fmt"
)

//定义一个节点

type HeroNode struct {
	no       int
	name     string
	nickName string
	pre      *HeroNode //指向前一个节点
	next     *HeroNode //指向下一个节点
}

//给链表插入一个节点.在双链表的最后加入
func InsertHeadNode(head *HeroNode, newHeroNode *HeroNode) {
	// 1. 找到该链表的最后一个节点,[利用中间变量找到最后一个节点]
	temp := head
	for {
		if temp.next == nil { //找到最后节点
			break
		}
		temp = temp.next // 指向下一个链表节点
	}

	// 2. 将newHeroNode加入节点
	temp.next = newHeroNode
	newHeroNode.pre = temp

}

// 根据no的编号从小到大排序
func InsertHeroNodeByNo(head *HeroNode, newHeroNode *HeroNode) {
	// 1. [利用中间变量找到最后合适节点]
	temp := head
	flag := true
	for {
		if temp.next == nil { //找到最后节点
			break
		} else if temp.next.no >= newHeroNode.no { // 说明newHeroNode应该在temp后面
			break
		} else if temp.next.no == newHeroNode.no { // 已存在这个表,不插入,
			flag = false
			break
		}
		temp = temp.next
	}
	// 2. 将newHeroNode加入节点
	if !flag {
		fmt.Println("已存在no", newHeroNode.no)
	} else {
		newHeroNode.next = temp.next //ok
		newHeroNode.pre = temp       //ok
		if temp.next != nil {
			temp.next.pre = newHeroNode //ok 让temp的下一个节点的前节点指向新节点
		}
		temp.next = newHeroNode //ok让新节点的下一个节点指向新节点
	}

}

//删除双向链表一个节点
func DelHeroNode(head *HeroNode, id int) {
	// 1. [利用中间变量找到最后合适节点]
	temp := head
	flag := false
	for {
		if temp.next == nil { //找到最后节点
			break
		} else if temp.next.no == id { // 已存在这个表,不插入,
			flag = true
			break
		}
		temp = temp.next
	}
	// 2. 将newHeroNode加入节点
	if flag { //直接把当前指针指向下下个指针
		temp.next = temp.next.next //ok
		if temp.next != nil {
			// 如果要删除的是最后一个节点,那么temp即将成为最后一个节点.
			// 即:将temp.next置为nil,那么temp.next.pre没有意义.
			temp.next.pre = temp //ok
		}

	} else {
		fmt.Println("sorry,删除的不存在", id)
	}
}

//显示链表的所有节点信息,和单链表一样
func ListHeroNode(head *HeroNode) {
	temp := head
	//if temp.next == nil {
	//	fmt.Println("empty link")
	//	return
	//}
	for {
		fmt.Printf("[%d,%s,%s,%p,%p,%p] <==>\n",
			temp.no, temp.name, temp.nickName, temp.pre, temp, temp.next) //因为
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

}

func ListHeroNode2(head *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for {
		fmt.Printf("[%d,%s,%s,%p,%p,%p] <==>\n",
			temp.no, temp.name, temp.nickName, temp.pre, temp, temp.next)
		if temp.pre == nil {
			break
		}
		temp = temp.pre
	}

}

func main() {
	//1. 头结点
	head := &HeroNode{name: "head"}

	//2. 创建新节点
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickName: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       4,
		name:     "林冲",
		nickName: "豹子头",
	}
	hero4 := &HeroNode{
		no:       3,
		name:     "无用",
		nickName: "智多星",
	}

	//3. 加入节点
	InsertHeadNode(head, hero1)
	InsertHeadNode(head, hero2)
	InsertHeadNode(head, hero3)
	InsertHeadNode(head, hero4)
	//InsertHeroNodeByNo(head, hero2)
	//InsertHeroNodeByNo(head, hero1)
	//InsertHeroNodeByNo(head, hero4)
	//InsertHeroNodeByNo(head, hero3) //如果id相同,后加在前,

	//4. 列出节点信息
	ListHeroNode(head)
	fmt.Println("逆序打印")
	ListHeroNode2(head)

	//5.删除
	//DelHeroNode(head, 2)
	//DelHeroNode(head, 3)

	//6. 继续显示
	//ListHeroNode(head)
}
