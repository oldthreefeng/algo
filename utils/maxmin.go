/*
@Time : 2019/9/19 0:34
@Author : louis
@File : maxmin
@Software: GoLand
*/

package utils

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
