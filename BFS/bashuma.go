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
	status int
	father int
	move   byte
}

var (
	result [MaxS]byte
	qHead  = 0
	qTail  = 1
	moves  = []byte{'u', 'd', 'r', 'l'}
	myQ    = [MaxS]Node{}
)

func NewStatus(status int, move byte) int {
	var tmp []byte
	var zeroPos int
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
	expanded[status] = 1
	for qHead != qTail {
		status = myQ[qHead].status
		if status == goalStatus {
			return true
		}
		for i := 0; i < 4; i++ {
			newStatus = NewStatus(status, moves[i])
			if newStatus == -1 {
				continue
			} else if _, ok := expanded[newStatus]; ok {
				continue
			}
			expanded[newStatus] = 1
			myQ[qTail] = Node{newStatus, qHead, moves[i]}
			qTail++
		}
		qHead++
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
			ms++
			pos = myQ[pos].father
			if pos == -1 {
				break
			}
		}
		fmt.Print("x")
		for i := ms - 1; i >= 0; i-- {
			fmt.Printf("%c->",result[i])
		}
	} else {
		fmt.Println("unsolvable")
	}
}
