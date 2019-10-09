/*
@Time : 2019/10/9 12:13
@Author : louis
@File : 392-isSubsequence
@Software: GoLand
*/

package leetcode

// IsSubsequence 给定字符串 s 和 t ，判断 s 是否为 t 的子序列.

func IsSubsequence(s, t string) bool {
	ls := len(s)
	if ls == 0 {
		return true
	}
	var i, j int
	for i < ls && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == ls
}
