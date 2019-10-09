/*
@Time : 2019/10/9 11:52
@Author : louis
@File : 680-validpalindrome2
@Software: GoLand
*/

package leetcode

import "gogs.wangke.co/go/algo/utils"

func ValidPalindrome(s string) bool {
	return len(s) - Lcs(s,utils.Reverse(s)) <= 1
}
