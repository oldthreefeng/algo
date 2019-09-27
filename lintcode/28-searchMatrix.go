/*
@Time : 2019/9/27 10:59
@Author : louis
@File : 28-searchMatrix
@Software: GoLand
*/

package lintcode

func SearchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	col := len(matrix[0])   //行
	row := len(matrix) 		//列
	lens := col * row
	l := 0
	r := lens -1
	for l <= r {
		mid := l + (r-l)/2
		tmp := matrix[mid/col][mid%col] // mid对应的第几行第几列
		if target > tmp {
			l = mid+1
		} else {
			r = mid-1
		}
	}
	if matrix[l/col][l%col] == target {
		return true
	}
	return false

}
