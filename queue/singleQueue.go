package main

import (
	"errors"
	"fmt"
	"os"
)

const maxSize = 5

type singleQueue struct {
	array [maxSize]int //模拟队列
	front int          //指向队列队首
	rear  int          //指向队列队尾
}

func help() string {
	return `
- add:   "add val to queue"
- get:   "get val from queue"
- show:  "show val from queue"
- exit:  "exit func"
`
}

func main() {
	var queue = &singleQueue{
		front: -1,
		rear:  -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(" queue.AddQueue(val) err = ", err)
			} else {
				fmt.Println(" add queue ok")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("get val = %d from queue\n", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("your input is valid~")
		}
	}

}

//在队列存入数据
func (this *singleQueue) AddQueue(val int) (err error) {
	//判断队列是否已满
	if this.rear == maxSize-1 { //rear是队列尾部,含最后一个元素
		return errors.New("queue full")
	}
	this.rear++ //rear 后移
	this.array[this.rear] = val
	return
}

//在队列取出数据
func (this *singleQueue) GetQueue() (val int, err error) {
	if this.rear == this.front {
		//队列满了之后将front和rear初始化
		this.rear = -1
		this.front = -1
		return -1, errors.New("queue is empty")
	}
	this.front++
	val = this.array[this.front]
	return
}

//显示队列内容
func (this *singleQueue) ShowQueue() { //front指向队首,不含第一个元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("singleQueue: array[%d] = %d\t", i, this.array[i])
		fmt.Println()
	}
}
