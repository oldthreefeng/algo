/*
@Time : 2019/8/16 17:58
@Author : louis
@File : migongL
@Software: GoLand
*/

package main

import (
	"errors"
	"fmt"
	"os"
)

// 读取文件中的数据，使用循环生成迷宫图
func readMazeFile(filename string) [][]int {
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

type point struct {
	i, j int
}

// 移动步骤顺序: 上、左、下、右
var direction = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

// 执行上左下右移动操作
func (p point) move(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 判断移动之后的位置是否有越界或者撞墙
func (p point) at(grid [][]int) (int, error) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, errors.New("raw out of range")
	}

	if p.j < 0 || p.j >= len(grid[0]) {
		return 0, errors.New("col out of range")
	}

	return grid[p.i][p.j], nil
}

func walk(maze [][]int, start, end point) [][]int {

	// 创建一个全0的二维数组，填行走的步骤
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		index := Q[0] // 每次取队列中的第一个元素进行上左下右移动

		// 发现终点，直接退出
		if index == end {
			break
		}

		Q = Q[1:]

		for _, d := range direction {
			next := index.move(d)

			val, err := next.at(maze)
			// 遇到墙或者越界，跳过本次循环
			if err != nil || val == 1 {
				continue
			}

			// 新的二维数组中移动的下一个点如果值不是0的话，说明已经走过这个点，直接跳过
			val, err = next.at(steps)
			if err != nil || val != 0 {
				continue
			}

			// 回到原点,也要跳过
			if next == start {
				continue
			}

			// steps二维数组初始值都是0，每走一步，原来的地方填上行走的步数上加一
			stepsNum, _ := index.at(steps)
			steps[next.i][next.j] = stepsNum + 1

			// 每走一步，先加入到队列中
			Q = append(Q, next)

		}
	}

	return steps
}

func main() {
	maze := readMazeFile("d:/text/maze.txt")
	/*  maze.txt 的内容如下:
	6 5
	0 1 0 0 0
	0 0 0 1 0
	0 1 0 1 0
	1 1 1 0 0
	0 1 0 0 1
	0 1 0 0 0
	*/
	fmt.Println("source maze: ")
	for i := range maze {
		for j := range maze[i] {
			fmt.Printf("%3d", maze[i][j])
		}
		fmt.Println()
	}

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	fmt.Println("\nnew steps: ")
	for i := range steps {
		for j := range steps[i] {
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}

	fmt.Println("")
	i := len(maze) - 1
	j := len(maze[0]) - 1

	fmt.Printf("走出迷宫总共移动的%d步\n",steps[i][j])


	// 计算最优路径
	last := steps[i][j]
	lookupPath := []point{point{i,j}}

	for last > 0 {
		last = last - 1
		index := lookupPath[len(lookupPath)-1]

		for _, d := range direction {
			next := index.move(d)

			val, _ := next.at(steps)
			if val == last {
				lookupPath = append(lookupPath,next)
				//fmt.Println(last,index,next)
				break
			}
		}


	}

	// 反转lookupPath
	newPath := []point{}
	for index,_ := range lookupPath {
		n := len(lookupPath) - index - 1
		newPath = append(newPath, lookupPath[n])
	}

	fmt.Printf("最优路径:%d", newPath)


}

