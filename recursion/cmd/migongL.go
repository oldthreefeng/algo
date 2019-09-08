/*
@Time : 2019/8/20 10:27
@Author : louis
@File : migongL
@Software: GoLand
*/

package main

import (
	"fmt"
	. "gogs.wangke.co/go/algo/recursion"
)

func main() {
	maze := ReadMazeFile("d:/text/maze.txt")
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

	steps := Walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})

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

	fmt.Printf("走出迷宫总共移动的%d步\n", steps[i][j])

	// 计算最优路径
	last := steps[i][j]
	lookupPath := []Point{Point{i, j}}

	for last > 0 {
		last = last - 1
		index := lookupPath[len(lookupPath)-1]

		for _, d := range Direction {
			next := index.Move(d)

			val, _ := next.At(steps)
			if val == last {
				lookupPath = append(lookupPath, next)
				//fmt.Println(last,index,next)
				break
			}
		}

	}

	// 反转lookupPath
	newPath := []Point{}
	for index, _ := range lookupPath {
		n := len(lookupPath) - index - 1
		newPath = append(newPath, lookupPath[n])
	}

	fmt.Printf("最优路径:%d", newPath)

}
