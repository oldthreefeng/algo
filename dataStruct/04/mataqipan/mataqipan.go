/*
Copyright 2019 louis.
@Time : 2019/9/25 22:30
@Author : louis
@File : mataqipan
@Software: GoLand

*/

package mataqipan

import (
	"fmt"
	"github.com/oldthreefeng/algo/stack"
)

var (
	chess     [8][8]int //初始化棋盘
	direction = [2][9]int{{0, -2, -1, 1, 2, 2, 1, -1, -2},
		{0, 1, 2, 2, 1, -1, -2, -2, -1}} // 马可以走的8个方向
	cur, next Spot        //当前步数和下一步
	s         stack.Stack //栈
	weight    [8][8]int   //表示该位置周围可行点的数目,比如weight[0][0] = 2 只有两步可走.
)

//Spot 保存当前点的位置及可行走方向是否走过
type Spot struct {
	x int    //行
	y int    //列
	d [9]int //d[i]记录了第i号方向是否已经走过,1表示走过
}
//Feasible 该点是不是可以走,超出棋盘界限或者已经走过,都不能走.
func Feasible(x, y int) bool {
	if 0 <= x && x < 8 && 0 <= y && y < 8 && chess[x][y] == 0 {
		return true
	}
	return false
}


//NextDirection 每次走下一步,选择下一步最少的权值,进行贪心算法
func NextDirection(c Spot) int {
	var MinDirection, Min int
	var x, y int
	Min = 9
	for i := 1; i <= 8; i++ {
		//访问过则不考虑
		if c.d[i] != 0 {
			continue
		}
		x = c.x + direction[0][i]
		y = c.y + direction[1][i]
		//选择最小的权值
		if Feasible(x, y) && weight[x][y] < Min {
			Min = weight[x][y]
			MinDirection = i
		}
	}
	return MinDirection
}


// InitWeight 初始化每个点的权值
// 初始为0; 对每个点进行探索,若方向可以行,则weight[i][j]++
func InitWeight() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			for k := 1; k <= 8; k++ {
				x := i + direction[0][k]
				y := j + direction[1][k]
				if Feasible(x, y) {
					weight[i][j]++
				}
			}
		}
	}
}

//SetWeight 当(x,y)点被占用的时候,当前节点权值设置为9,位置(x,y)周围所有的可行点权值减1
func SetWeight(x, y int) {
	for k := 1; k <= 8; k++ {
		weight[x][y] = 9
		i := x + direction[0][k]
		j := y + direction[1][k]
		if Feasible(i, j) {
			weight[i][j]--
		}
	}
}

// UnsetWeight 回退操作,需要重新计算weight[i][j]的权值,
// 依次探测周围,若可行,则weight[i][j]++; 其周围可行点的权值+1
func UnsetWeight(x, y int) {
	for k := 1; k <= 8; k++ {
		weight[x][y] = 0
		i := x + direction[0][k]
		j := y + direction[1][k]
		if Feasible(i, j) {
			weight[x][y]++
			weight[i][j]++
		}
	}
}
//output 输出棋盘
func output() {
	for j := 0; j < 8; j++ {
		fmt.Printf("%2d\n", chess[j])
	}
	//for j := 0; j < 8; j++ {
	//	fmt.Printf("%2d\n", weight[j])
	//}
}

func outWeight() {
	for j := 0; j < 8; j++ {
		fmt.Printf("%2d\n", weight[j])
	}
}

// 当找不到下一个位置时,即NextDirection返回值为0,要进行回退
// 为了回退方便,使用栈来存储, 能进时,当前的位置入栈, 向i走一步
// 回退操作, 在棋牌的cur点置0,Step--; 出栈一个点,设置为当前的cur
// 回退操作, 不能重复探测,去重操作.

func Solution(x,y int) int {
	InitWeight()
	cur.x,cur.y= x,y
	backup := 0
	Step := 1
	SetWeight(cur.x, cur.y)
	chess[cur.x][cur.y] = Step
	for Step < 64 {
		k := NextDirection(cur)
		if k != 0 {
			next.x = cur.x + direction[0][k]
			next.y = cur.y + direction[1][k]

			cur.d[k] = 1
			s.Push(cur)
			cur = next
			Step++
			chess[cur.x][cur.y] = Step
			SetWeight(cur.x, cur.y)
			//回退
		} else {
			chess[cur.x][cur.y] = 0
			backup++
			Step--
			UnsetWeight(cur.x, cur.y)
			cur = s.Pop().(Spot)
		}
	}
	output()
	return backup
}

//func main() {
//
//	InitWeight()
//	outWeight()
//	fmt.Scanln(&cur.x,&cur.y)
//	backup := 0
//	Step := 1
//	SetWeight(cur.x, cur.y)
//	chess[cur.x][cur.y] = Step
//	for Step < 64 {
//		k := NextDirection(cur)
//		if k != 0 {
//			next.x = cur.x + direction[0][k]
//			next.y = cur.y + direction[1][k]
//
//			cur.d[k] = 1
//			s.Push(cur)
//			cur = next
//			Step++
//			chess[cur.x][cur.y] = Step
//			SetWeight(cur.x, cur.y)
//			//回退
//		} else {
//			chess[cur.x][cur.y] = 0
//			backup++
//			Step--
//			UnsetWeight(cur.x, cur.y)
//			cur = s.Pop().(Spot)
//		}
//	}
//	output()
//	fmt.Print(backup)
//}
