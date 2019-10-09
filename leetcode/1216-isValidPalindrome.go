/*
@Time : 2019/10/9 10:49
@Author : louis
@File : 1216-isValidPalindrome
@Software: GoLand
*/

package leetcode

import "gogs.wangke.co/go/algo/utils"

/*
IsValidPalindrome
给出一个字符串 s 和一个整数 k，请你帮忙判断这个字符串是不是一个「K 回文」。
所谓「K 回文」：如果可以通过从字符串中删去最多 k 个字符将其转换为回文，那么这个字符串就是一个「K 回文」。
*/
func IsValidPalindrome(s string, k int) bool {
	return len(s) - Lcs(s, utils.Reverse(s)) <= k
}

