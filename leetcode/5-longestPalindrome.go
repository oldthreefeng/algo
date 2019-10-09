/*
Copyright 2019 louis.
@Time : 2019/10/9 22:54
@Author : louis
@File : 5-longestPalindrome
@Software: GoLand

*/

package leetcode

import "gogs.wangke.co/go/algo/utils"

func LongestPalindrome(s string) string {
	return StrLcs(s,utils.Reverse(s))
}
