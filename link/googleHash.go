package main

import (
	"errors"
	"fmt"
	"os"
)

/*
- google上机题
有新人来的时候,要求将该员工的信息加入(ID,性别,年龄,住址....) 当输入ID时,要求查找到该员工的所以信息
1.不允许使用数据库,要求速度越快越来,尽量节省内存.
2. 添加是,保证雇员的ID从低到高插入


散列表(Hash table, 也叫哈希表), 通过K-v直接访问的数据结构,这个通过关键码值来映射查找,
这个映射函数就叫散列函数,存放记录的数组叫散列表.

- 维护用户在线一亿条数据
struct userState {
	ID int64
	state byte
}
*/

// Emp to store employee status
type Emp struct {
	ID     int
	UserID string
	Name   string
	Mail   string
	Next   *Emp
}

// ShowMe to show  Emp status
func (e *Emp) ShowMe() {
	fmt.Printf("Link %d emp ID = %d name =%s ->\n",
		e.ID%7, e.ID, e.Name)
}

// EmoLink to store Emp Link
type EmoLink struct {
	Head *Emp
}

// Insert to add Emp to EmoLink
func (e *EmoLink) Insert(emp *Emp) {
	cur := e.Head
	var pre *Emp //辅助指针
	if cur == nil {
		e.Head = emp
		return
	}
	// 不是空链表,给emp插入合适的位置
	for {
		if cur == nil {
			break
		}
		if cur.ID > emp.ID { //从小到大
			break
		}
		pre = cur
		cur = cur.Next
	}

	pre.Next = emp
	emp.Next = cur
}

// ShowLink is to show EmoLink status
func (e *EmoLink) ShowLink(num int) {
	if e.Head == nil {
		fmt.Printf("Link %d is Empty", num) //第num个链表
	}
	cur := e.Head
	for {
		if cur == nil {
			break
		}
		fmt.Printf("Link %d emp ID = %d name =%s ->", num, cur.ID, cur.Name)
		cur = cur.Next
	}
	fmt.Println()
}

//FindByID to find a *Emp on id
func (e *EmoLink) FindByID(id int) (emp *Emp, err error) {
	cur := e.Head
	for {
		if cur == nil {
			return nil, errors.Errorf("%d is not exist", id)
		}
		if cur.ID == id {
			return cur, nil
		}
		cur = cur.Next
	}
}

//DelEmp to delete Emp in EmoLink by id
func (e *EmoLink) DelEmp(id int) (err error) {
	cur := e.Head
	if cur == nil {
		return errors.Errorf("%d is not exist", id)
	}
	if cur.Next == nil && cur.ID == id { //要删的是第一个用户,且后续没有,需要把头指针置空
		e.Head = cur.Next
		return
	}
	for {
		if cur.Next == nil {
			break
		} else if cur.Next.ID == id {
			cur.Next = cur.Next.Next
			return
		}
		cur = cur.Next
	}
	return errors.Errorf("%d is not exist", id)
}

//HashTable store EmoLink Slice
type HashTable struct {
	Link [7]EmoLink
}

// Insert Emp to HashTable
func (e *HashTable) Insert(emp *Emp) {
	//使用散列函数,确定将该雇员添加到哪个链表
	linkNo := e.HashFun(emp.ID)
	_, err := e.Link[linkNo].FindByID(emp.ID)
	if err != nil {
		// 使用对应的链表进行添加
		e.Link[linkNo].Insert(emp)
	} else {
		fmt.Println(emp.ID, "emp is exist")
	}
}

// HashFun  a algo to store emp
func (e *HashTable) HashFun(ID int) int { //TODO 可以考虑二级链表
	return ID % 7
}

//FindByID to find a *Emp on id
func (e *HashTable) FindByID(ID int) *Emp {
	linkNO := e.HashFun(ID)
	emp, _ := e.Link[linkNO].FindByID(ID)
	return emp
}

// ShowAll Emp
func (e *HashTable) ShowAll() {
	for i := 0; i < len(e.Link); i++ {
		e.Link[i].ShowLink(i)
	}
}

//DelEmp to del Emp
func (e *HashTable) DelEmp(id int) (err error) {
	linkId := e.HashFun(id)
	return e.Link[linkId].DelEmp(id)
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
		_, _ = fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入雇员ID:")
			_, _ = fmt.Scanln(&id)
			fmt.Println("请输入雇员姓名")
			_, _ = fmt.Scanln(&name)
			emp := &Emp{
				ID:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "del":
			fmt.Println("请输入雇员ID:")
			_, _ = fmt.Scanln(&id)
			err := hashTable.DelEmp(id)
			if err != nil {
				fmt.Println(err)
			}
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入ID号:")
			_, _ = fmt.Scanln(&id)
			emp := hashTable.FindByID(id)
			if emp == nil {
				fmt.Printf("ID =%d 的雇员不存在\n", id)
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
