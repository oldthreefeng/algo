/*
Copyright 2019 louis.
@Time : 2019/9/28 1:02
@Author : louis
@File : sumintslice
@Software: GoLand

*/

package utils

func SumSlice(res []int) int {
	tmp := 0
	for _, v := range res {
		tmp += v
	}
	return tmp
}
