/*
Copyright 2019 louis.
@Time : 2019/10/9 22:54
@Author : louis
@File : 5-longestPalindrome
@Software: GoLand

*/

package leetcode

import (
	"gogs.wangke.co/go/algo/utils"
	"strings"
)

func LongestPalindromeX(s string) string {
	return StrLcs(s,utils.Reverse(s))
}

func LongestPalindrome(s string) string {
	var sb strings.Builder
	n := len(s)
	if n == 0 {
		return ""
	}

	for i := 0; i < n; i++ {
		sb.WriteByte('#')
		sb.WriteByte(s[i])
	}
	sb.WriteByte('#')

	ms := sb.String()

	N := len(ms)

	// fmt.Println(ms, N)

	dp := make([]int, N)
	right := 0
	var center int

	var left int
	var maxi int
	var maxy int
	for i := 0; i < N; i++ {
		// fmt.Println("start ->", i, right, center, 2*center-i, dp[left])
		if i < right {
			left = 2*center - i
			dp[i] = dp[left]
			if i+dp[i] < right {
				continue
			}
		}

		for right <= N-1 && 2*i-right >= 0 && ms[right] == ms[2*i-right] {
			right++
			center = i
		}

		// fmt.Println("i -> ", i, right, center)

		dp[i] = right - i
		if dp[i] > maxy {
			maxy = dp[i]
			maxi = i
		}
	}

	// fmt.Println(dp)
	l, r := (maxi-dp[maxi]+1)/2, (maxi+dp[maxi]-1)/2
	return string(s[l:r])
}
