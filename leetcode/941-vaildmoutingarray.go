/*
Copyright 2019 louis.
@Time : 2019/10/1 21:36
@Author : louis
@File : 941-vaildmoutingarray
@Software: GoLand

*/

package leetcode

// ValidMountingArray 判断是否是山脉数组.
//
// 思路: 找到递增数列的下标i,找到递减数列的下标j; i==j时说明为山脉数组.
// i=0 || j=len(arr)-1不符合山脉数组的要求.
func ValidMountingArray(m []int) bool {

	i, j, l := 0, len(m)-1, len(m)
	if j < 2 {
		return false
	}
	for i+1 < l && m[i] < m[i+1] {
		i++
	}
	for j > 0 && m[j-1] > m[j] {
		j--
	}
	if i == 0 || j == l-1 {
		return false
	}
	return i == j
}
