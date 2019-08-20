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

func NQueen(k int) {
	var queenPos [N]int
	var i int
	if k == N {
		//for i = 0; i < N-1; i++ {
		//	fmt.Printf("%d\n", queenPos[i]+1)
		//}
		fmt.Println(queenPos)
		return
	}
	for i = 0; i < N; i++ { //遍历第k皇后的为止,
		var j int
		for j = 0; j < k; j++ {
			if queenPos[j] == i ||queenPos[j]-i == k-j ||
				queenPos[j]-i == j-k{
				break
			}
		}
		if j == k {
			queenPos[k] = i //将k皇后摆法好i位置
			NQueen(k + 1)
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
		if safe(boards,col) {
			put(boards,col+1)
		}
	}
}

func safe(boards []int, col int) bool {
	for c:=0;c<col ;c++{
		if isAttack(boards,c,col) {
			return false
		}
	}
	return true
}

func isAttack(boards []int, c,col int) bool {
	if boards[c] == boards[col]{
		return true
	}
	if boards[col] - boards[c] == c-col {
		return true
	}
	if boards[col] - boards[c] == col-c{
		return true
	}
	return false
}