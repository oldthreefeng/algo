/*
@Time : 2019/9/23 20:42
@Author : louis
@File : catchthatcow
@Software: GoLand
*/

package main

import (
	"fmt"
	"github.com/oldthreefeng/algo/queue"
)

var (
	K       int
	N       int
	visited [MaxN + 10]int
	q       queue.Queue
)

const MaxN = 100000

type Step struct {
	x     int //位置
	steps int //步数
}

func NewStep(xx, s int) Step {
	return Step{x: xx, steps: s}
}

func main() {
	fmt.Scanln(&N,&K)
	q.Enqueue(Step{N, 0})
	visited[N] = 1

	for !q.IsEmpty() {
		s := q.Peek().(Step)
		if s.x == K {
			fmt.Println(s.steps)
			return
		} else {
			if s.x-1 >= 0 && visited[s.x-1] == 0 {
				//向左走一步,左边的点塞入队列.
				q.Enqueue(Step{s.x - 1, s.steps + 1})
				visited[s.x-1] = 1
			} //向右走一步,
			if s.x+1 <= MaxN && visited[s.x+1] == 0 {
				q.Enqueue(Step{s.x + 1, s.steps + 1})
				visited[s.x+1] = 1
			} // 跳步, 跳两倍
			if s.x*2 <= MaxN && visited[s.x*2] == 0 {
				q.Enqueue(Step{s.x * 2, s.steps + 1})
				visited[s.x*2] = 1
			}
			q.Dequeue()
		}
	}
}
