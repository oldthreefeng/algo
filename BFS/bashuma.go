/*
@Time : 2019/9/24 0:45
@Author : louis
@File : bashuma
@Software: GoLand
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	goalStatus = 123456780
	MaxS       = 400000
)

type Node struct {
	status int  //当前状态数
	father int  //父节点
	move   byte //移动
}

var (
	result [MaxS]byte                   //移动的结果存放
	qHead  = 0                          //队列头部
	qTail  = 1                          //队列尾部
	moves  = []byte{'u', 'd', 'r', 'l'} //移动的方向,上下左右
	myQ    = [MaxS]Node{}               //队列
	//q      queue.Queue
)

func NewStatus(status int, move byte) int {
	var tmp []byte
	var zeroPos int // 记录x也就是0的位置
	date := []byte(fmt.Sprint(status))
	tmp = date[:9]
	for i := 0; i < 9; i++ {
		if tmp[i] == '0' {
			zeroPos = i
			break
		}
	}
	switch move {
	case 'u':
		if zeroPos-3 < 0 {
			return -1 //空格在第一行
		} else {
			tmp[zeroPos] = tmp[zeroPos-3]
			tmp[zeroPos-3] = '0'
		}
	case 'd':
		if zeroPos+3 > 8 {
			return -1 // 空格在第三行
		} else {
			tmp[zeroPos] = tmp[zeroPos+3]
			tmp[zeroPos+3] = '0'
		}
	case 'l':
		if zeroPos%3 == 0 {
			return -1 //空格在第一列
		} else {
			tmp[zeroPos] = tmp[zeroPos-1]
			tmp[zeroPos-1] = '0'
		}
	case 'r':
		if zeroPos%3 == 2 {
			return -1 //空格在第三列
		} else {
			tmp[zeroPos] = tmp[zeroPos+1]
			tmp[zeroPos+1] = '0'
		}
	}
	a, _ := strconv.Atoi(string(tmp[:9]))
	return a
}

func Bfs(status int) bool {
	var newStatus int
	expanded := make(map[int]int)
	myQ[qHead] = Node{status: status, father: -1, move: 0}
	//q.Enqueue(Node{status: status, father: -1, move: 0})
	expanded[status] = 1 //将头结点加入到队列,采用map来进行去重
	for qHead != qTail {
		//for !q.IsEmpty() {
		status = myQ[qHead].status
		//status = q.Peek().(Node).status
		//如果队列头部的状态为目标状态,说明可以走通
		if status == goalStatus {
			return true
		}
		for i := 0; i < 4; i++ { //分四个方向走,movies[0-3]-->u, d, r, l
			newStatus = NewStatus(status, moves[i])
			if newStatus == -1 { //不能朝某个方向,就直接换方向
				continue
				//如果这个状态我曾经达到过,也不用拓展,说明在同一级或者上一级
			} else if _, ok := expanded[newStatus]; ok {
				continue
			}
			//新状态,先记录到map里面
			expanded[newStatus] = 1
			//将新状态加入到队列尾部,同事,记录自己的父节点,以及移动的方向.
			//q.Enqueue(Node{newStatus, qHead, moves[i]})
			myQ[qTail] = Node{newStatus, qHead, moves[i]}
			qTail++
		}
		qHead++ //将头指针向后移一位,如果4步都不能加入尾结点,说明该节点不可拓展
		//q.Dequeue()
	}
	return false
}

func main() {
	var line [10]byte
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadBytes('\n')
	if err != nil {
		return
	}
	var i int
	for _, v := range input {
		if v != 32 {
			if v == 'x' {
				line[i] = '0'
			} else {
				line[i] = v
			}
			i++
		}

	}

	sl, _ := strconv.Atoi(string(line[:9]))
	if Bfs(sl) {
		var ms = 0
		var pos = qHead
		for {
			result[ms] = myQ[pos].move
			//result[ms] = q.PeekBack().(Node).move
			ms++
			pos = myQ[pos].father
			if pos == 0 {
				break
			}
		}
		//fmt.Print("x")
		for i := ms - 1; i >= 0; i-- {
			fmt.Printf("%c->", result[i])
		}
	} else {
		fmt.Println("unsolvable")
	}
}
