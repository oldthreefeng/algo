package main

import (
	"fmt"
	"os"
)

/*
- google上机题
有新人来的时候,要求将该员工的信息加入(id,性别,年龄,住址....) 当输入id时,要求查找到该员工的所以信息
1.不允许使用数据库,要求速度越快越来,尽量节省内存.
2. 添加是,保证雇员的id从低到高插入


散列表(Hash table, 也叫哈希表), 通过K-v直接访问的数据结构,这个通过关键码值来映射查找,
这个映射函数就叫散列函数,存放记录的数组叫散列表.

- 维护用户在线一亿条数据
struct userState {
	id int64
	state byte
}
*/

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("Link %d emp id = %d name =%s ->",
		this.Id%7, this.Id, this.Name)
}

func (this *Emp) IsEmpty() bool {
	return this.Next == nil
}

type EmoLink struct {
	Head *Emp
}

func (this *EmoLink) IsEmpty() bool {
	return this.Head == nil
}

func (this *EmoLink) Insert(emp *Emp) {
	cur := this.Head
	var pre *Emp = nil //辅助指针
	if cur == nil {
		this.Head = emp
		return
	}
	// 不是空链表,给emp插入合适的位置
	for {
		if cur == nil {
			break
		}
		if cur.Id > emp.Id { //从小到大
			break
		}
		pre = cur
		cur = cur.Next
	}

	pre.Next = emp
	emp.Next = cur
}

func (this *EmoLink) ShowLink(num int) {
	if this.Head == nil {
		fmt.Printf("Link %d is Empty\n", num) //第num个链表
	}
	cur := this.Head
	for {
		if cur == nil {
			break
		}
		fmt.Printf("Link %d emp id = %d name =%s ->", num, cur.Id, cur.Name)
		cur = cur.Next
	}
	fmt.Println()
}
func (this *EmoLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur == nil {
			return nil
		}
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}
}

type HashTable struct {
	Link [7]EmoLink
}

func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数,确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	// 使用对应的链表进行添加
	this.Link[linkNo].Insert(emp)
}

func (this *HashTable) HashFun(id int) int { //可以考虑二级链表
	return id % 7
}

func (this *HashTable) FindById(id int) *Emp {
	linkNO := this.HashFun(id)
	return this.Link[linkNO].FindById(id)
}

// list所以的雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.Link); i++ {
		this.Link[i].ShowLink(i)
	}
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("---------雇员系统菜单---------")
		fmt.Println("add------表示添加雇员---------")
		fmt.Println("show-----表示查找雇员---------")
		fmt.Println("find-----表示查找雇员---------")
		fmt.Println("del------表示删除雇员---------")
		fmt.Println("exit-----表示退出系统---------")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入雇员id:")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员姓名")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "del":
		//	TODO func Delete
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Printf("id =%d 的雇员不存在\n", id)
			} else {
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}

	}
}
