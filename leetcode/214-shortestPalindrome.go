/*
Copyright 2019 louis.
@Time : 2019/10/9 21:54
@Author : louis
@File : 214-shortestPalindrome
@Software: GoLand

*/

package leetcode

import "github.com/oldthreefeng/algo/utils"

func ShortestPalindrome(s string) string  {
	if len(s) <= 1 {
		return s
	}
	rs := utils.Reverse(s)
	var k int
	for k = range s {
		if s[:len(s)-k] == rs[k:] {
			break
		}
	}
	return rs[:k] + s
}

func ShortestPalindromeX(s string) string  {
	revStr := utils.Reverse(s)
	newStr := s + "#" + revStr
	next := make([]int, len(newStr), len(newStr))

	for i := 1; i < len(newStr); i++ {

		t := next[i - 1]
		for ; t > 0 && newStr[i] != newStr[t]; {
			t = next[t - 1]
		}

		if newStr[i] == newStr[t] {
			t++
		}

		next[i] = t
	}

	return revStr[0 : len(s) -  next[len(newStr) - 1]] + s
}