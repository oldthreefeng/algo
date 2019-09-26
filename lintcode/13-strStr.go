/*
@Time : 2019/9/26 16:07
@Author : louis
@File : 13-strStr
@Software: GoLand
*/

package lintcode

func StrStr(source, target string) int {
	a, b := len(target), len(source)
	if a == 0 {
		return 0
	}
	if b == 0 || a > b {
		return -1
	}
	for i := 0; i < b-a+1; i++ {
		tmp := source[i:a+i]
		if tmp == target {
			return i
		}
	}
	return -1
}
