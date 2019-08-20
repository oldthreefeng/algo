/*
@Time : 2019/8/16 17:58
@Author : louis
@File : migongL
@Software: GoLand
*/

package recursion

import (
	"errors"
	"fmt"
	"os"
)

// 读取文件中的数据，使用循环生成迷宫图
func ReadMazeFile(filename string) [][]int {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type Point struct {
	I, J int
}

// 移动步骤顺序: 上、左、下、右
var Direction = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

// 执行上左下右移动操作
func (p Point) Move(r Point) Point {
	return Point{p.I + r.I, p.J + r.J}
}

// 判断移动之后的位置是否有越界或者撞墙
func (p Point) At(grid [][]int) (int, error) {
	if p.I < 0 || p.I >= len(grid) {
		return 0, errors.New("raw out of range")
	}

	if p.J < 0 || p.J >= len(grid[0]) {
		return 0, errors.New("col out of range")
	}

	return grid[p.I][p.J], nil
}

func Walk(maze [][]int, start, end Point) [][]int {

	// 创建一个全0的二维数组，填行走的步骤
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []Point{start}

	for len(Q) > 0 {
		index := Q[0] // 每次取队列中的第一个元素进行上左下右移动

		// 发现终点，直接退出
		if index == end {
			break
		}

		Q = Q[1:]

		for _, d := range Direction {
			next := index.Move(d)

			val, err := next.At(maze)
			// 遇到墙或者越界，跳过本次循环
			if err != nil || val == 1 {
				continue
			}

			// 新的二维数组中移动的下一个点如果值不是0的话，说明已经走过这个点，直接跳过
			val, err = next.At(steps)
			if err != nil || val != 0 {
				continue
			}

			// 回到原点,也要跳过
			if next == start {
				continue
			}

			// steps二维数组初始值都是0，每走一步，原来的地方填上行走的步数上加一
			stepsNum, _ := index.At(steps)
			steps[next.I][next.J] = stepsNum + 1

			// 每走一步，先加入到队列中
			Q = append(Q, next)

		}
	}

	return steps
}

