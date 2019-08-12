package main

import (
	"errors"
	"fmt"
	"os"
)

//使用结构体管理环形队列,要求默写

const maxSize = 5

type circleQueue struct {
	array [maxSize]int
	head  int //队首
	tail  int //队尾
}

func help() string {
	return `
- add:   "add val to queue"
- get:   "get val from queue"
- show:  "show val from queue"
- exit:  "exit func"
`
}

func (this *circleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	// this.tail不含队尾元素
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % maxSize
	return
}
func (this *circleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	//取出数据,head包含队首元素
	val = this.array[this.head]
	this.head = (this.head + 1) % maxSize
	return
}

func (this *circleQueue) ListQueue() {
	size := this.Size()
	if size == 0 {
		fmt.Println("empty queue")
	}
	//设计辅助变量
	temphead := this.head
	for i := this.head; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", temphead, this.array[temphead])
		temphead = (temphead + 1) % maxSize
	}
}

//判断环形队列是否满
func (this *circleQueue) IsFull() bool {
	return (this.tail+1)%maxSize == this.head
}

//判断环形队列是否空
func (this *circleQueue) IsEmpty() bool {
	return this.head == this.tail
}

//取出环形队列有多少元素
func (this *circleQueue) Size() int {
	return (this.tail + maxSize - this.head) % maxSize
}

func main() {
	var queue = &circleQueue{
		tail: 0,
		head: 0,
	}

	var key string
	var val int
	for {
		fmt.Println(help())
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("你要添加什么数到queue")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(" queue.AddQueue(val) err = ", err)
			} else {
				fmt.Println(" Push queue ok")
			}
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Pop val = %d from queue\n", val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("your input is valid~")
		}
	}

}
