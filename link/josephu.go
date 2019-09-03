package main

import "fmt"

type Person struct {
	no   int
	next *Person
}

// 输入总人数n,返回head指针的环形列表
func AddPerson(n int) *Person {
	first := &Person{}
	curPer := &Person{}
	if n < 1 {
		return first
	}
	for i := 1; i <= n; i++ {
		person := &Person{
			no: i,
		}

		if i == 1 {
			first = person //第一个用户,构成循环列表
			curPer = person
			curPer.next = first
		} else {
			curPer.next = person
			curPer = person
			curPer.next = first
		}
	}

	return first
}

func ListPer(first *Person) int {
	temp := first
	index := 0
	if temp.next == nil {
		return -1
	}
	for {
		index++
		fmt.Printf("id=%d -> ", temp.no)
		if temp.next == first {
			break
		}
		temp = temp.next
	}
	fmt.Println()
	return index
}

/*
设编号为1，2，...n的n个人围坐一圈，约定编号为k（1<=k<=n）的人从1开始报数，
数到m的那个人出列，它的下一位又从1开始报数，数到m的那个人又出列，依次类推，直到所有人出列为止，
由此产生一个出队编号的序列。

1. 编写函数,playGame(first *Person ,startNum ,countNum int)
2. 按照要求在环形链表中留下1人即可
*/

func playGame(first *Person, startNum, countNum int) {
	if first.next == nil || startNum > ListPer(first) {
		return
	}
	//first指针和tail指针,
	tail := first
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}

	//移动指针到startNum,
	for i := 1; i < startNum; i++ {
		first = first.next
		tail = tail.next
	}

	//数countNum下,删除当前first指针
	for {
		//数countNum下,删除当前first指针,即出圈
		for j := 1; j < countNum; j++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("id=%d出列 ->\n", first.no)
		first = first.next
		tail.next = first

		if tail == first { //如果只有一个人,就退出
			break
		}
	}
	fmt.Printf("id=%d出列 ->", first.no)

}

//编写函数,构成单向环形列表
func main() {
	first := AddPerson(7)
	//ListPer(first)
	playGame(first, 1, 3)
}
