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

func (c *circleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("queue full")
	}
	// c.tail不含队尾元素
	c.array[c.tail] = val
	c.tail = (c.tail + 1) % maxSize
	return
}
func (c *circleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	//取出数据,head包含队首元素
	val = c.array[c.head]
	c.head = (c.head + 1) % maxSize
	return
}

func (c *circleQueue) ListQueue() {
	size := c.Size()
	if size == 0 {
		fmt.Println("empty queue")
	}
	//设计辅助变量
	temphead := c.head
	for i := c.head; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", temphead, c.array[temphead])
		temphead = (temphead + 1) % maxSize
	}
}

//判断环形队列是否满
func (c *circleQueue) IsFull() bool {
	return (c.tail+1)%maxSize == c.head
}

//判断环形队列是否空
func (c *circleQueue) IsEmpty() bool {
	return c.head == c.tail
}

//取出环形队列有多少元素
func (c *circleQueue) Size() int {
	return (c.tail + maxSize - c.head) % maxSize
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
