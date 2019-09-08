/**
 * @time: 2019-08-20 00:27
 * @author: louis
 */
package recursion

import (
	"fmt"
)

/*
输入一个正整数N,输出N皇后的全部摆法
*/

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

const N = 4

func NQueen(col int) {
	var queenPos [N]int
	var pos int
	if col == N {
		//for i = 0; i < N-1; i++ {
		//	fmt.Printf("%d\n", queenPos[i]+1)
		//}
		fmt.Println(queenPos)
		return
	}
	for pos = 0; pos < N; pos++ { //逐个尝试第k个皇后的位置,遍历第k皇后的为止,
		var j int
		for j = 0; j < col; j++ {
			if queenPos[j] == pos || queenPos[j]-pos == col-j ||
				queenPos[j]-pos == j-col {
				break
			}
		}
		if j == col {
			queenPos[col] = pos //将k皇后摆法好i位置
			NQueen(col + 1)
		}

	}
}

func Queen(size int) {
	boards := make([]int, size)
	put(boards, 0)
}

func put(boards []int, col int) {
	size := len(boards)
	if col == size {
		fmt.Println(boards)
		return
	}
	for pos := 0; pos < size; pos++ {
		boards[col] = pos
		if safe(boards, col) {
			put(boards, col+1)
		}
	}
}

func safe(boards []int, col int) bool {
	for c := 0; c < col; c++ {
		if isAttack(boards, c, col) {
			return false
		}
	}
	return true
}

func isAttack(boards []int, c, col int) bool {
	if boards[c] == boards[col] {
		return true
	}
	if boards[col]-boards[c] == c-col {
		return true
	}
	if boards[col]-boards[c] == col-c {
		return true
	}
	return false
}
